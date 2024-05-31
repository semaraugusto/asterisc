all: bin bin/minimal bin/no_std bin/matmul

bin:
	mkdir bin

bin/minimal:
	# cd minimal && cargo +unicorn build --release --target riscv64im-unicorn-zkvm-elf && cp ./target/riscv64im-unicorn-zkvm-elf/release/minimal ../bin/minimal
	cd minimal && cargo +unicorn build --release --target riscv32im-succinct-zkvm-elf && cp ./target/riscv32im-succinct-zkvm-elf/release/minimal ../bin/minimal

bin/no_std:
	# cd no_std && cargo +unicorn build --release --target riscv64im-unicorn-zkvm-elf && cp ./target/riscv64im-unicorn-zkvm-elf/release/no_std ../bin/no_std
	cd no_std && cargo +unicorn build --release --target riscv32im-succinct-zkvm-elf && cp ./target/riscv32im-succinct-zkvm-elf/release/no_std ../bin/no_std

bin/matmul:
	cd matmul && cargo +unicorn build --release --target riscv64im-unicorn-zkvm-elf && cp ./target/riscv64im-unicorn-zkvm-elf/release/matmul ../bin/matmul
	# cd matmul && cargo +unicorn build --release --target riscv32im-succinct-zkvm-elf && cp ./target/riscv32im-succinct-zkvm-elf/release/matmul ../bin/matmul

clean:
	rm -rf bin && cd minimal && cargo clean && cd ../matmul && cargo clean


# minimal.dump: minimal
# 	riscv64-linux-gnu-objdump -D --disassemble --disassembler-options=no-aliases --wide --source -m riscv:rv64 -EL bin/minimal > bin/minimal.dump
#
