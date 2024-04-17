// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RISCVMetaData contains all meta data concerning the RISCV contract.
var RISCVMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"contractIPreimageOracle\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPreimageOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"step\",\"inputs\":[{\"name\":\"stateData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"localContext\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051620021793803806200217983398101604081905261003191610056565b600080546001600160a01b0319166001600160a01b0392909216919091179055610086565b60006020828403121561006857600080fd5b81516001600160a01b038116811461007f57600080fd5b9392505050565b6120e380620000966000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637dc0d1d01461003b578063e14ced3214610085575b600080fd5b60005461005b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610098610093366004612039565b6100a6565b60405190815260200161007c565b6000610370565b8060005260206000fd5b905090565b600067ffffffffffffffff6100cf565b90565b909116919050565b60006100f2601f6100ed63ffffffff5b85610287565b610110565b92915050565b6000600160405b1b905090565b60006001603f6100ff565b60006001831b821680156101345767ffffffffffffffff841c841b8317915061014d565b61014a67ffffffffffffffff85603f031c841690565b91505b5092915050565b600061015e610105565b82168015610190577fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000083179150610194565b8291505b50919050565b60006101ae6101a76100f8565b8484010690565b9392505050565b60006101ae6101c26100f8565b8484030690565b60006101ae8383026100bc565b60006101ae8383046100bc565b60006101ae6101f184610154565b6101fa84610154565b056100bc565b60008282066101ae565b60006101ae61021884610154565b61022184610154565b076100bc565b60006100f282196100bc565b60008282106101ae565b60008282116101ae565b60006101ae61025584610154565b61025e84610154565b1290565b60006101ae61027084610154565b61027984610154565b1390565b60008183146101ae565b60008183166101ae565b60008183176101ae565b60008183186101ae565b60006101ae83831b6100bc565b600082821c6101ae565b60006101ae6102ca84610154565b831d6100bc565b60208110156101945760ff83168260081b1791508260081c92506001810190506102d1565b6000602060005b01905090565b600060206102fd6102f6565b600060086102fd610303565b600060086102fd61030f565b600060016102fd61031b565b600060016102fd610327565b600060086102fd610333565b600060086102fd61033f565b600060086102fd61034b565b60006101006102fd610357565b60806040511461037f57600080fd5b6084861461038c57600080fd5b610394610363565b6020870335146103a357600080fd5b60206103bd6103b0610363565b601f808216602003160190565b87010184146103cb57600080fd5b61022484146103d957600080fd5b6103e1610363565b6080016040526103ef610363565b866080376113cf565b6080015160209190910360031b1c90565b60800180517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209390930360031b92831b19169290911b919091179052565b50565b60006100b760206103f86102f6565b6104498160206104096102f6565b60006100b760086103f8610303565b610449816008610409610303565b60006100b760086103f861030f565b61044981600861040961030f565b60006100b760016103f8610327565b6104c0600180610409610327565b565b60006100b760016103f861031b565b61044981600161040961031b565b60006100b760086103f8610333565b610449816008610409610333565b60006100b760086103f861033f565b61044981600861040961033f565b60006100b760086103f861034b565b61044981600861040961034b565b6000610543601f8361023d565b156105545761055462bad4e96100ad565b6105726105626008846101c9565b61056d6100cc610357565b61019a565b608081015160c01c6101ae565b80610588575050565b610594601f5b8261023d565b156105a5576105a562bad4e96100ad565b6105b36105626008836101c9565b60808101805177ffffffffffffffffffffffffffffffffffffffffffffffff1660c085901b179052505050565b505050565b60006105ef6104a3565b600181146105ff57600391505090565b6106076104c2565b801561061e57600181146106275760029250505090565b60009250505090565b60019250505090565b610638610363565b6080a06000610645610363565b60802090506106526105e5565b60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff919091161790565b60006100f2600b6100ed8460146102b2565b6102b2565b60006100f2600b6100ed6106b3601f6106ae8760076102b2565b610287565b6106cd6106c18760196102b2565b60056102a5565b6102a5565b610291565b60006100f2600c6100ed61070e6106f46106ed87601f6102b2565b600c6102a5565b6106cd61070760016106ae8a60076102b2565b600b6102a5565b6106cd6107246106c1603f6106ae8a60196102b2565b6106cd610737600f6106ae8b60086102b2565b60016102a5565b60006100f260136100ed84600c6102b2565b60006100f260136100ed61078561077261076b87601f6102b2565b60136102a5565b6106cd61070760ff6106ae8a600c6102b2565b6106cd6107a261079b60016106ae8a60146102b2565b600a6102a5565b6106cd6103ff6106ae8a60156102b2565b60006100f2607f5b83610287565b60006100f2601f6106ae8460076102b2565b60006100f260076106ae84600c6102b2565b60006100f2601f6106ae84600f6102b2565b60006100f2601f6106ae8460146102b2565b60006100f28260196102b2565b60006100f28260146102b2565b600061083a6020610835603c856101c9565b6101c9565b90506100f26102245b8261019a565b6000610855601f6107bb565b156108675761086763bad10ad06100ad565b61087083610823565b803561087e60205b8361019a565b915061088b8460056102b2565b8160005b603b8110156108f95784356108a560208761019a565b95506108b660016106ae86856102b2565b80156108c957600181146108de576108ef565b600084815260208390526040902093506108ef565b600082815260208590526040902093505b505060010161088f565b5060805193508381146109135761091363badf00d16100ad565b509095945050505050565b61092a601f5b82610287565b1561093c5761093c63bad10ad06100ad565b61094583610823565b826109506020610878565b915061095d8360056102b2565b60005b603b8110156109cb57833561097760205b8661019a565b945061098860016106ae85856102b2565b801561099b57600181146109b0576109c1565b600085815260208390526040902094506109c1565b600082815260208690526040902094505b5050600101610960565b50506109d681608052565b5050505050565b600060088311156109f5576109f563bad512e06100ad565b610a076107bb601f610227565b610227565b610a146100cc8783610849565b610a1e82856101b5565b6000610a41610a2d601f610227565b6106ae610a3b60018b6101b5565b8961019a565b610a5388610a4e85602082565b6101b5565b6000610a66610a62848961027d565b1590565b15610aa35760ff8d03610a8057610a8063bad222206100ad565b610a8d6100cc8e85610849565b935060009150610aa08a610a4e87604082565b90505b85610ab26100cc8460036102a5565b1c955083610ac46100cc8360036102a5565b1c9350505050600091505b86821015610b5957610af7610ae4601f610227565b6106ae84610a4e6001610a4e8d8d61019a565b6000610b03868361027d565b60018114610b16578015610b2657610b32565b600886901c9560ff169150610b32565b600884901c9360ff1691505b50610b4a610b3f826100bc565b6106cd8960086102a5565b96505050600182019150610acf565b505050508315610b8157610b736001610a4e8560036102a5565b610b7d8183610110565b9150505b95945050505050565b60008060008084610b9b878761019a565b60005b6040811015610c635780610bb3602082610233565b8060018114610bc7578015610bd857610be5565b8760081b97508960081b9950610be5565b8660081b96508860081b98505b50610c07610bf38584610233565b6106ae6000610c028987610233565b61027d565b15610c595760ff8c610c266100cc610c1f8e876101b5565b60036102a5565b1c168160018114610c3c578015610c4b57610c56565b9781179760ff8b179a50610c56565b9681179660ff8a1799505b50505b5050600101610b9e565b50505093509350935093565b6020821115610c8557610c8563bad512e16100ad565b610c92610924601f610227565b610cb3610c9f601f610227565b6106ae610cad6001876101b5565b8561019a565b610cbd82846101b5565b610cc8868683610b8a565b9350610cd76100cc8b88610849565b8319168117610ce78b828961091e565b5050610cf3848661027d565b15610d045750505050505050505050565b60ff8a03610d1957610d1963bad222216100ad565b610d266100cc8b86610849565b9019169091179250610d41905087610d3b8490565b8361091e565b50505050505050565b6109d68585858585610c6f565b60008360018114610d7f5760028114610d845760038114610d9557610d7f63badc0de06100ad565b610da4565b610d8e8483610291565b9350610da4565b610da16107bb85610227565b93505b509392505050565b6000610db8601f6107bb565b610dc38160206101b5565b610dcd818661023d565b15610dd6578094505b50610dee6100cc6001610de984876101b5565b610849565b80610dfd6100cc8460036102a5565b1b905080610e126100cc610c1f8860206101b5565b1c90508460031b9150610e2361044c565b821b8117610e308161045b565b505050610e43610e3e600090565b610478565b5090919050565b600080805460405163e03110e160e01b8152846004820152856024820152620186a06040600060448460008786f1925050508015610e92576000519250602051915050610ea0565b50610ea063badf00d06100ad565b9250929050565b6000610eb161044c565b610eb9610469565b8160f81c60018103610f2157604080516000858152336020528983526060902091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001792505b50610f2c8183610e4a565b925082610f3f57600093505050506101ae565b610f4b601f5b86610287565b610f568160206101b5565b610f60818961023d565b15610f69578097505b50610f74848861023d565b15610f7d578396505b610f8b610c1f8860206101b5565b6001901b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01199350610fc26100cc8260036102a5565b93841c9382901c9150610fd5878461019a565b9250610fe083610478565b610fef6001610de983896101b5565b84191684831617925061100d600184611008848a6101b5565b61091e565b50949695505050505050565b6110236011610536565b80605d81146110b057605e81146110b05760d681146110cf5760de81146110f157603f811461116a5760408114611200576019811461125a576038811461130857607181146113235760dc81146113555760a3811461135f5761010581146113c2576101a681146113c257606581146113c2576110a360005b600a61057f565b6105e060005b600b61057f565b6110ba600a610536565b6110c660ff82166104d1565b506105e06104b2565b6110db6001601e6102a5565b6110e681600a61057f565b506105e060006110a9565b6110fb600a610536565b611105600b610536565b818061115c57611116610fff6107bb565b80156111345761113161112b826110006101b5565b8461019a565b92505b5061113d6104fc565b61114881600a61057f565b61115a611155848361019a565b61050b565b505b5050506105e06110a9600090565b611174600a610536565b61117e600b610536565b611188600c610536565b6000808480156111b857600381146111c357600581146111cd5767ffffffffffffffff5b9250604d5b91506111df565b6000925060006111b1565b83925060006111b1565b6111d8898587610ea7565b9250600091505b506111eb82600a61057f565b6111f681600b61057f565b5050505050505050565b61120a600a610536565b611214600b610536565b61121e600c610536565b60008084600181146111c357600281146111c357600481146111c357600681146112505767ffffffffffffffff6111ac565b6111d88486610dac565b611264600a610536565b61126e600b610536565b600080826003811461128e5767ffffffffffffffff9250601691506112f1565b8480156112d957600181146112e357600281146112e357600381146112d957600481146112e357600581146112d957600681146112ea5767ffffffffffffffff9350604d92506112ef565b60005b93506112ef565b60016112dc565b600193505b505b506112fd82600a61057f565b610d4181600b61057f565b61131967ffffffffffffffff61109c565b6105e0600d6110a9565b61132d600b610536565b682a0000000000000539611348600260018360105b86610c6f565b50506110a361109c600090565b6110a3600161109c565b611369600a610536565b611373600b610536565b816007811461138d57611388620f00126100ad565b6113ba565b6113a66002600169040000000000000004006010611342565b6113b0600061109c565b6113ba60006110a9565b505050505050565b6105e063f001ca116100ad565b6113d76104a3565b156113ed576113e4610630565b60005260206000f35b6114026113fd600161056d6104df565b6104ee565b61140a610486565b61141a60ff6000806004856109dd565b611423816107b3565b61142c826107c1565b611435836107d3565b61143e846107e5565b611447856107f7565b61145086610809565b85600381146114eb576023811461156157606381146115b7576013811461169c57601b81146117a4576033811461186157603b8114611a685760378114611c5b5760178114611c7b57606f8114611c975760678114611cd75760738114611d2b57602f8114611dcb57600f8114611fcc5760078114611fcc5760278114611fcc5760538114611fd9576114e663f001c0de6100ad565b611fe6565b6114f48861067d565b611503610a6260045b88610287565b61151460016106c860035b8a610287565b61151d87610536565b61152c610843600b5b86610110565b93505061153e600260018484876109dd565b9250505061154c818861057f565b506114e661155c60045b8b61019a565b610495565b61156a88610694565b6115756001876102a5565b61157e85610536565b61158787610536565b611594610843600b611526565b9350506115a660026001838587610d4a565b5050506114e661155c611556600490565b6115c084610536565b6115c984610536565b60008780156115ff57600181146116105760048114611622576005811461162c576006811461163e57600781146116485761165d565b611609838561027d565b915061165d565b61160960016106ae610a02868861027d565b6116098385610247565b61160960016106ae610a028688610247565b6116098385610233565b61165a60016106ae610a028688610233565b91505b508080156116815761166e8c6106d2565b935061167a848e61019a565b9c5061168f565b61168c60048e61019a565b9c505b505050506114e689610495565b6116a584610536565b6116ae8961067d565b60008780156116f457600181146117055760028114611713576003811461171d576004811461172757600581146117315760068114611782576007811461178c57611799565b6116fe838561019a565b9150611799565b6116fe846106c8603f610f45565b6116fe8385610247565b6116fe8385610233565b6116fe838561029b565b61173c8360066102b2565b801561174f57601081146117665761177c565b61175f8561068f603f5b87610287565b925061177c565b61177985611774603f611759565b6102bc565b92505b50611799565b6116fe8385610291565b6117968385610287565b91505b506115a6818a61057f565b6117ad84610536565b6117b68961067d565b60008780156117d457600181146117e657600581146117f757611799565b6116fe6117e1848661019a565b6100d7565b6116fe6117e1856106c8601f611759565b611802601f84610287565b61180d8460056102b2565b8015611820576020811461184357611854565b61183c601f5b6100ed61183663ffffffff61150e565b856102b2565b9350611854565b61185161182683601f6101b5565b93505b5050506115a6818a61057f565b61186a84610536565b61187384610536565b6000846001811461196a578880156118c257600181146118fa57600281146119085760038114611912576004811461191c5760058114611926576006811461195657600781146119605761177c565b8680156118d657602081146118e7576118f4565b6118e0858761019a565b93506118f4565b6118f185876101b5565b93505b5061177c565b61175f856106c8603f611759565b61175f8486610247565b61175f8486610233565b61175f848661029b565b86801561193a5760208114611948576118f4565b6118e08661068f603f6114fd565b6118f186611774603f6114fd565b61175f8486610291565b6117798486610287565b8880156119ae57600181146119bf57600281146119dd57600381146119e657600481146119f05760058114611a1b5760068114611a2c5760078114611a4657611a5c565b6119b884866101c9565b9250611a5c565b6119b86119cb85610154565b6119d487610154565b0260401c6100bc565b6119b8846119cb565b6119b884866119d4565b838015611a0857611a0185876101e3565b9350611a15565b67ffffffffffffffff5b93505b50611a5c565b838015611a0857611a0185876101d6565b838015611a3d57611a01858761020a565b85935050611a5c565b838015611a575761183c8587610200565b859350505b50506115a6818a61057f565b611a7184610536565b611a7a84610536565b60008460018114611b5b57888015611aa15760018114611af55760058114611b065761177c565b868015611ab55760208114611ad5576118f4565b6118e06117e1611ac863ffffffff6114fd565b61056d63ffffffff61150e565b6118f16117e1611ae863ffffffff6114fd565b610a4e63ffffffff61150e565b61175f6117e1866106c8601f6114fd565b611b10601f6100e7565b878015611b245760208114611b4257611b53565b611b3b601f5b6100ed61183663ffffffff8b610287565b9450611b53565b611b50611b2a83601f6101b5565b94505b505050611799565b888015611b875760048114611ba95760058114611bcd5760068114611bf95760078114611c2657611a5c565b6119b86117e1611b9a63ffffffff611759565b61083563ffffffff5b89610287565b838015611a0857611a016117e1611bbf876100d7565b611bc8896100d7565b6101e3565b838015611a0857611a016117e1611be763ffffffff6114fd565b611bf463ffffffff61150e565b6101d6565b838015611c1d57611a016117e1611c0f876100d7565b611c18896100d7565b61020a565b611a12866100d7565b838015611c525761183c6117e1611c4063ffffffff6114fd565b611c4d63ffffffff61150e565b610200565b611851866100d7565b611c648861073e565b611c6f81600c6102a5565b905061154c818861057f565b611c848861073e565b611c6f611556601f6100ed84600c6102a5565b611ca088610750565b611caa6004611556565b611cb4818961057f565b50611cd161155c611ccb60146100ed8560016102a5565b8c61019a565b50611fe6565b611ce084610536565b611ce98961067d565b611cf36004611ccb565b611cfd818a61057f565b50611d2461155c611d17600167ffffffffffffffff61029b565b6106ae610971600b611526565b5050611fe6565b848015611d9057611d3b89610816565b85611d49610a62600461150e565b15611d5a57611d5787610536565b90505b611d646003611ba3565b611d6f818385610d57565b92505050611d7d818961057f565b50611d8b61155c6004611ccb565b611cd1565b611d9b8960146102b2565b8015611db557611db061155c60045b8d61019a565b611d24565b611dbe8d611019565b611d2461155c6004611daa565b611dd66001866102a5565b611de0600861058e565b611deb600483610233565b1715611dfd57611dfd62bada706100ad565b611e0685610536565b611e118460026102b2565b8060028114611f475760038114611f6f57611e2b87610536565b611e3660048661027d565b15611e4757611e44816100d7565b90505b611e56600260018088886109dd565b80848015611eb05760018114611ec15760048114611ec95760088114611ed357600c8114611edd5760108114611ee75760148114611efe5760188114611f0857601c8114611f1257611eab630f001a706100ad565b611f25565b611eba848461019a565b9250611f25565b839250611f25565b611eba848461029b565b611eba8484610291565b611eba8484610287565b611ef18385610247565b15611eab57839250611f25565b611ef18385610262565b611ef18385610233565b611f1c838561023d565b15611f25578392505b50611f3560036001848a8a610d4a565b611f3f818e61057f565b505050611fba565b611f56600260018087876109dd565b611f60818c61057f565b50611f6a83610528565b611fba565b6001611f82611f7c610519565b8561027d565b15611fa557611f9088610536565b611f9f60026001838989610d4a565b50600090505b611faf818c61057f565b50611fba6000610528565b505050506114e661155c611556600490565b6114e661155c6004611556565b611fe661155c6004611556565b5050505050505050506113e4610630565b60008083601f84011261200957600080fd5b50813567ffffffffffffffff81111561202157600080fd5b602083019150836020828501011115610ea057600080fd5b60008060008060006060868803121561205157600080fd5b853567ffffffffffffffff8082111561206957600080fd5b61207589838a01611ff7565b9097509550602088013591508082111561208e57600080fd5b5061209b88828901611ff7565b9699959850966040013594935050505056fea2646970667358221220b56fa7ccb04f35ec52a8b6db1efaa1609dbd5aea86e180015e0cfff668730b4c64736f6c634300080f0033",
}

// RISCVABI is the input ABI used to generate the binding from.
// Deprecated: Use RISCVMetaData.ABI instead.
var RISCVABI = RISCVMetaData.ABI

// RISCVBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RISCVMetaData.Bin instead.
var RISCVBin = RISCVMetaData.Bin

// DeployRISCV deploys a new Ethereum contract, binding an instance of RISCV to it.
func DeployRISCV(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address) (common.Address, *types.Transaction, *RISCV, error) {
	parsed, err := RISCVMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RISCVBin), backend, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RISCV{RISCVCaller: RISCVCaller{contract: contract}, RISCVTransactor: RISCVTransactor{contract: contract}, RISCVFilterer: RISCVFilterer{contract: contract}}, nil
}

// RISCV is an auto generated Go binding around an Ethereum contract.
type RISCV struct {
	RISCVCaller     // Read-only binding to the contract
	RISCVTransactor // Write-only binding to the contract
	RISCVFilterer   // Log filterer for contract events
}

// RISCVCaller is an auto generated read-only Go binding around an Ethereum contract.
type RISCVCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RISCVTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RISCVTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RISCVFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RISCVFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RISCVSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RISCVSession struct {
	Contract     *RISCV            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RISCVCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RISCVCallerSession struct {
	Contract *RISCVCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RISCVTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RISCVTransactorSession struct {
	Contract     *RISCVTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RISCVRaw is an auto generated low-level Go binding around an Ethereum contract.
type RISCVRaw struct {
	Contract *RISCV // Generic contract binding to access the raw methods on
}

// RISCVCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RISCVCallerRaw struct {
	Contract *RISCVCaller // Generic read-only contract binding to access the raw methods on
}

// RISCVTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RISCVTransactorRaw struct {
	Contract *RISCVTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRISCV creates a new instance of RISCV, bound to a specific deployed contract.
func NewRISCV(address common.Address, backend bind.ContractBackend) (*RISCV, error) {
	contract, err := bindRISCV(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RISCV{RISCVCaller: RISCVCaller{contract: contract}, RISCVTransactor: RISCVTransactor{contract: contract}, RISCVFilterer: RISCVFilterer{contract: contract}}, nil
}

// NewRISCVCaller creates a new read-only instance of RISCV, bound to a specific deployed contract.
func NewRISCVCaller(address common.Address, caller bind.ContractCaller) (*RISCVCaller, error) {
	contract, err := bindRISCV(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RISCVCaller{contract: contract}, nil
}

// NewRISCVTransactor creates a new write-only instance of RISCV, bound to a specific deployed contract.
func NewRISCVTransactor(address common.Address, transactor bind.ContractTransactor) (*RISCVTransactor, error) {
	contract, err := bindRISCV(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RISCVTransactor{contract: contract}, nil
}

// NewRISCVFilterer creates a new log filterer instance of RISCV, bound to a specific deployed contract.
func NewRISCVFilterer(address common.Address, filterer bind.ContractFilterer) (*RISCVFilterer, error) {
	contract, err := bindRISCV(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RISCVFilterer{contract: contract}, nil
}

// bindRISCV binds a generic wrapper to an already deployed contract.
func bindRISCV(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RISCVMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RISCV *RISCVRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RISCV.Contract.RISCVCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RISCV *RISCVRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RISCV.Contract.RISCVTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RISCV *RISCVRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RISCV.Contract.RISCVTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RISCV *RISCVCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RISCV.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RISCV *RISCVTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RISCV.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RISCV *RISCVTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RISCV.Contract.contract.Transact(opts, method, params...)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_RISCV *RISCVCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RISCV.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_RISCV *RISCVSession) Oracle() (common.Address, error) {
	return _RISCV.Contract.Oracle(&_RISCV.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_RISCV *RISCVCallerSession) Oracle() (common.Address, error) {
	return _RISCV.Contract.Oracle(&_RISCV.CallOpts)
}

// Step is a paid mutator transaction binding the contract method 0xe14ced32.
//
// Solidity: function step(bytes stateData, bytes proof, bytes32 localContext) returns(bytes32)
func (_RISCV *RISCVTransactor) Step(opts *bind.TransactOpts, stateData []byte, proof []byte, localContext [32]byte) (*types.Transaction, error) {
	return _RISCV.contract.Transact(opts, "step", stateData, proof, localContext)
}

// Step is a paid mutator transaction binding the contract method 0xe14ced32.
//
// Solidity: function step(bytes stateData, bytes proof, bytes32 localContext) returns(bytes32)
func (_RISCV *RISCVSession) Step(stateData []byte, proof []byte, localContext [32]byte) (*types.Transaction, error) {
	return _RISCV.Contract.Step(&_RISCV.TransactOpts, stateData, proof, localContext)
}

// Step is a paid mutator transaction binding the contract method 0xe14ced32.
//
// Solidity: function step(bytes stateData, bytes proof, bytes32 localContext) returns(bytes32)
func (_RISCV *RISCVTransactorSession) Step(stateData []byte, proof []byte, localContext [32]byte) (*types.Transaction, error) {
	return _RISCV.Contract.Step(&_RISCV.TransactOpts, stateData, proof, localContext)
}
