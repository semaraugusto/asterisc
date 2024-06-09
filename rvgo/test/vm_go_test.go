package test

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"

	preimage "github.com/ethereum-optimism/optimism/op-preimage"

	"github.com/ethereum-optimism/asterisc/rvgo/fast"
	"github.com/ethereum-optimism/asterisc/rvgo/slow"
)

type testOracle struct {
	hint        func(v []byte)
	getPreimage func(k [32]byte) []byte
}

func (t *testOracle) Hint(v []byte) {
	t.hint(v)
}

func (t *testOracle) GetPreimage(k [32]byte) []byte {
	return t.getPreimage(k)
}

func (t *testOracle) ReadPreimagePart(key [32]byte, offset uint64) (dat [32]byte, datlen uint8, err error) {
	v := t.getPreimage(key)
	if offset == uint64(len(v))+8 {
		return [32]byte{}, 0, nil // datlen==0 signals EOF
	}
	if offset > uint64(len(v))+8 {
		err = fmt.Errorf("cannot read past pre-image (%x) size: %d >= %d", key, offset, len(v))
		return
	}
	// read the size prefix
	if offset < 8 {
		var tmp [8]byte
		binary.BigEndian.PutUint64(tmp[:], uint64(len(v)))
		copy(dat[:8-offset], tmp[offset:])
		datlen += 8 - uint8(offset)
		offset = 0
	} else {
		offset -= 8
	}
	// once past the size prefix, read the actual preimage
	datlen += uint8(copy(dat[datlen:], v[offset:]))
	return
}

var _ fast.PreimageOracle = (*testOracle)(nil)
var _ slow.PreimageOracle = (*testOracle)(nil)

func fullTest(t *testing.T, vmState *fast.VMState, po *testOracle, symbols fast.SortedSymbols, runSlow, runEVM bool) {
	instState := fast.NewInstrumentedState(vmState, po, os.Stdout, os.Stderr)

	var contracts *Contracts
	var addrs *Addresses
	var env *vm.EVM
	if runEVM {
		contracts = testContracts(t)
		addrs = testAddrs
		env = newEVMEnv(t, contracts, addrs)
	}

	maxGasUsed := uint64(0)

	// var lastSym elf.Symbol
	for i := uint64(0); i < 20_000_000_000; i++ {
		// for i := uint64(0); i < 390_000; i++ {
		sym := symbols.FindSymbol(vmState.PC)

		// if sym.Name != lastSym.Name {
		// 	instr := vmState.Instr()
		// 	fmt.Printf("i: %4d  pc: 0x%x  instr: %08x  symbol name: %s size: %d\n", i, vmState.PC, instr, sym.Name, sym.Size)
		// }
		// lastSym = sym

		if sym.Name == "runtime.throw" {
			throwArg := vmState.Registers[10]
			throwArgLen := vmState.Registers[11]
			if throwArgLen > 1000 {
				throwArgLen = 1000
			}
			x := vmState.Memory.ReadMemoryRange(throwArg, throwArgLen)
			dat, _ := io.ReadAll(x)
			if utf8.Valid(dat) {
				fmt.Printf("THROW! %q\n", string(dat))
			} else {
				fmt.Printf("THROW! %016x: %x\n", throwArg, dat)
			}
			break
		}
		wit, err := instState.Step(runEVM || runSlow)
		require.NoError(t, err, "fast VM must run step")

		if runEVM || runSlow {
			fastPostState := vmState.EncodeWitness()
			fastStateHash, err := fastPostState.StateHash()
			require.NoError(t, err)

			if runSlow {
				input, err := wit.EncodeStepInput(fast.LocalContext{})
				require.NoError(t, err)
				slowPostHash, err := slow.Step(input, po)
				require.NoErrorf(t, err, "slow VM err at step %d, PC %08x: %v", i, vmState.PC, err)
				require.Equal(t, fastStateHash, slowPostHash, "fast post-state must match slow post-state")
			}

			if runEVM {
				evmPost, evmPostHash, gasUsed := stepEVM(t, env, wit, addrs, i, nil)
				if gasUsed > maxGasUsed {
					maxGasUsed = gasUsed
				}

				if evmPostHash != fastStateHash {
					t.Fatalf("evm state %x must match fast state %x\nfast:%x\nevm: %x\nat step %d\n", evmPostHash, fastStateHash, fastPostState, evmPost, i)
				}
			}
		}

		if vmState.Exited {
			break
		}
	}

	t.Logf("max gas used: %d", maxGasUsed)

	require.True(t, vmState.Exited, "ran out of steps")
	if vmState.ExitCode != 0 {
		t.Fatalf("failed with exit code %d", vmState.ExitCode)
	}
}

func TestSimple(t *testing.T) {
	programELF, err := elf.Open("../../tests/go-tests/bin/simple")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	preImages := make(map[[32]byte][]byte)
	addPreimage := func(img []byte) {
		h := crypto.Keccak256Hash(img)
		preImages[preimage.Keccak256Key(h).PreimageKey()] = img
	}
	addInput := func(i uint64, val []byte) {
		preImages[preimage.LocalIndexKey(i).PreimageKey()] = val
	}

	addInput(0, crypto.Keccak256([]byte("hello")))        // pre-state
	addInput(1, crypto.Keccak256([]byte("world")))        // input
	addInput(2, crypto.Keccak256([]byte("hello world!"))) // claim to verify
	addPreimage([]byte("hello"))                          // pre-state pre-image
	addPreimage([]byte("world"))                          // input pre-image

	po := &testOracle{
		hint: func(v []byte) {
			t.Logf("received hint: %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Logf("reading pre-image: %x", k)
			if v, ok := preImages[k]; ok {
				return v
			} else {
				t.Fatalf("unknown pre-image %x", k)
				return nil
			}
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
	})

	t.Run("slow", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, true, false)
	})

	t.Run("evm", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, true)
	})
}

func TestMinimal(t *testing.T) {
	programELF, err := elf.Open("../../tests/go-tests/bin/minimal")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
	})

	t.Run("slow", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, true, false)
	})

	t.Run("evm", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, true)
	})
}

func TestRustWithStd(t *testing.T) {
	programELF, err := elf.Open("../../tests/rust-tests/minimal/target/riscv32im-succinct-zkvm-elf/release/minimal")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
	})

	// t.Run("slow", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, true, false)
	// })
	//
	// t.Run("evm", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, false, true)
	// })
}

func TestRustMatmul(t *testing.T) {
	programELF, err := elf.Open("../../tests/rust-tests/matmul/target/riscv64im-unicorn-zkvm-elf/release/matmul")
	// programELF, err := elf.Open("../../tests/rust-tests/matmul/target/riscv32im-succinct-zkvm-elf/release/matmul")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
	})
	//
	// t.Run("slow", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, true, false)
	// })
	//
	// t.Run("evm", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, false, true)
	// })
}

func TestRustMnist(t *testing.T) {
	programELF, err := elf.Open("../../tests/rust-tests/mnist/target/riscv64im-unicorn-zkvm-elf/release/mnist")
	// programELF, err := elf.Open("../../tests/rust-tests/mnist/target/riscv32im-succinct-zkvm-elf/release/mnist")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
	})
	//
	// t.Run("slow", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, true, false)
	// })
	//
	// t.Run("evm", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, false, true)
	// })
}

func TestRustOnnx(t *testing.T) {
	// programELF, err := elf.Open("../../tests/rust-tests/onnx/target/riscv32im-succinct-zkvm-elf/release/onnx")
	programELF, err := elf.Open("../../tests/rust-tests/onnx/target/riscv64im-unicorn-zkvm-elf/release/onnx")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
		fmt.Println("after fullTest")
	})
	//
	// t.Run("slow", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, true, false)
	// })
	//
	// t.Run("evm", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, false, true)
	// })
}

func TestRustCross(t *testing.T) {
	t.Skip("rust-cross is only available on my machine")
	// programELF, err := elf.Open("../../../rust-cross/target/riscv32im-succinct-zkvm-elf/release/rust-cross")
	programELF, err := elf.Open("../../../rust-cross/target/riscv64im-unicorn-zkvm-elf/release/rust-cross")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		fullTest(t, vmState, po, symbols, false, false)
		fmt.Println("after fullTest")
	})
	//
	// t.Run("slow", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, true, false)
	// })
	//
	// t.Run("evm", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, false, true)
	// })
}

func TestRustWithoutStd(t *testing.T) {
	programELF, err := elf.Open("../../tests/rust-tests/no_std/target/riscv64gc-unknown-none-elf/release/no_std")
	require.NoError(t, err)
	defer programELF.Close()

	symbols, err := fast.Symbols(programELF)
	require.NoError(t, err)

	po := &testOracle{
		hint: func(v []byte) {
			t.Fatalf("unexpected pre-image hint %x", v)
		},
		getPreimage: func(k [32]byte) []byte {
			t.Fatalf("unexpected pre-image request %x", k)
			return nil
		},
	}

	t.Run("fast", func(t *testing.T) {
		vmState, err := fast.LoadELF(programELF)
		require.NoError(t, err, "must load test suite ELF binary")

		err = fast.PatchVM(programELF, vmState)
		require.NoError(t, err, "must patch VM")

		// get memory reader
		var oneAsBytes []byte
		var tenAsBytes []byte
		BIG_ENDIAN := false
		if BIG_ENDIAN {
			// oneAsBytes := make([]byte, 7)
			oneAsBytes = []byte{0, 0, 0, 0, 0, 0, 0, 1}
			tenAsBytes = []byte{0, 0, 0, 0, 0, 0, 0, 10}
		} else {
			oneAsBytes = []byte{1, 0, 0, 0, 0, 0, 0, 0}
			tenAsBytes = []byte{10, 0, 0, 0, 0, 0, 0, 0}
		}
		fmt.Println("oneAsBytes: ", oneAsBytes)
		fmt.Println("tenAsBytes: ", tenAsBytes)
		// r := vmState.Memory.AllocPage(0)
		if err = vmState.Memory.SetMemoryRange(0x1000, bytes.NewReader(oneAsBytes)); err != nil { // disable mem profiling, to avoid a lot of unnecessary floating point ops
			require.NoError(t, err, "Cannot set memory range for one")
		}

		if err = vmState.Memory.SetMemoryRange(0x1008, bytes.NewReader(tenAsBytes)); err != nil { // disable mem profiling, to avoid a lot of unnecessary floating point ops
			require.NoError(t, err, "Cannot set memory range for ten")
		}
		fullTest(t, vmState, po, symbols, false, false)
	})

	// t.Run("slow", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, true, false)
	// })
	//
	// t.Run("evm", func(t *testing.T) {
	// 	vmState, err := fast.LoadELF(programELF)
	// 	require.NoError(t, err, "must load test suite ELF binary")
	//
	// 	err = fast.PatchVM(programELF, vmState)
	// 	require.NoError(t, err, "must patch VM")
	//
	// 	fullTest(t, vmState, po, symbols, false, true)
	// })
}
