// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC20HandlerDepositRecord struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_useContractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originChainTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceIDAndContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b5060405162002c6338038062002c638339818101604052810190620000379190620004a4565b81518351146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000673565b60405180910390fd5b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8351811015620001175762000109848281518110620000e057fe5b6020026020010151848381518110620000f557fe5b60200260200101516200016660201b60201c565b8080600101915050620000c5565b5060008090505b81518110156200015b576200014d8282815181106200013957fe5b60200260200101516200025860201b60201c565b80806001019150506200011e565b505050505062000798565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b62000269816200030660201b60201c565b620002ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002a29062000651565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000815190506200036d8162000764565b92915050565b600082601f8301126200038557600080fd5b81516200039c6200039682620006c3565b62000695565b91508181835260208401935060208101905083856020840282011115620003c257600080fd5b60005b83811015620003f65781620003db88826200035c565b845260208401935060208301925050600181019050620003c5565b5050505092915050565b600082601f8301126200041257600080fd5b8151620004296200042382620006ec565b62000695565b915081818352602084019350602081019050838560208402820111156200044f57600080fd5b60005b838110156200048357816200046888826200048d565b84526020840193506020830192505060018101905062000452565b5050505092915050565b6000815190506200049e816200077e565b92915050565b60008060008060808587031215620004bb57600080fd5b6000620004cb878288016200035c565b945050602085015167ffffffffffffffff811115620004e957600080fd5b620004f78782880162000400565b935050604085015167ffffffffffffffff8111156200051557600080fd5b620005238782880162000373565b925050606085015167ffffffffffffffff8111156200054157600080fd5b6200054f8782880162000373565b91505092959194509250565b60006200056a60248362000715565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000620005d260478362000715565b91507f6d69736d61746368206c656e677468206265747765656e20696e697469616c5260008301527f65736f7572636549447320616e6420696e697469616c436f6e7472616374416460208301527f64726573736573000000000000000000000000000000000000000000000000006040830152606082019050919050565b600060208201905081810360008301526200066c816200055b565b9050919050565b600060208201905081810360008301526200068e81620005c3565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715620006b957600080fd5b8060405250919050565b600067ffffffffffffffff821115620006db57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200070457600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620007338262000744565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200076f8162000726565b81146200077b57600080fd5b50565b62000789816200073a565b81146200079557600080fd5b50565b6124bb80620007a86000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80636ebcf607116100a2578063c8ba6c8711610071578063c8ba6c87146102dc578063d9caed121461030c578063db95e75c14610328578063e245386f14610358578063fc9539cd1461038e5761010b565b80636ebcf607146102445780637f79bea8146102745780638a025ce8146102a457806395601f09146102c05761010b565b80633af32abf116100de5780633af32abf1461019857806345274738146101c857806345a104db146101f85780636a70d081146102145761010b565b806307b7ed99146101105780630a6d55d81461012c5780632521ab411461015c578063318c136e1461017a575b600080fd5b61012a600480360381019061012591906119bc565b6103aa565b005b61014660048036038101906101419190611a5d565b61044d565b604051610153919061204e565b60405180910390f35b610164610480565b604051610171919061213f565b60405180910390f35b610182610493565b60405161018f919061204e565b60405180910390f35b6101b260048036038101906101ad91906119bc565b6104b9565b6040516101bf919061213f565b60405180910390f35b6101e260048036038101906101dd91906119bc565b61050f565b6040516101ef9190612299565b60405180910390f35b610212600480360381019061020d9190611b55565b610527565b005b61022e600480360381019061022991906119bc565b610845565b60405161023b919061213f565b60405180910390f35b61025e600480360381019061025991906119bc565b610865565b60405161026b9190612299565b60405180910390f35b61028e600480360381019061028991906119bc565b61087d565b60405161029b919061213f565b60405180910390f35b6102be60048036038101906102b99190611a86565b61089d565b005b6102da60048036038101906102d591906119e5565b610a23565b005b6102f660048036038101906102f191906119bc565b610b51565b604051610303919061215a565b60405180910390f35b610326600480360381019061032191906119e5565b610b69565b005b610342600480360381019061033d9190611b03565b610c09565b60405161034f9190612277565b60405180910390f35b610372600480360381019061036d9190611b03565b610dbf565b60405161038597969594939291906120c9565b60405180910390f35b6103a860048036038101906103a39190611ac2565b610ee6565b005b6103b3816104b9565b6103f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e9906121d7565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ae906121b7565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061062a816104b9565b610669576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066090612257565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106cb576106c6818885611156565b6106d8565b6106d781883086611263565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600760008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff160217905550604082015181600101556060820151816002015560808201518160030190805190602001906107e59291906117d6565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461093f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093690612217565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016109969190612033565b60405160208183030381529060405280519060200120826040516020016109bd9190612033565b6040516020818303038152906040528051906020012014610a13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0a90612197565b60405180910390fd5b610a1d8484611392565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401610a6593929190612069565b602060405180830381600087803b158015610a7f57600080fd5b505af1158015610a93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab79190611a34565b50610b09826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60046020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf0906121b7565b60405180910390fd5b610c048383836114d9565b505050565b610c11611856565b600760008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d4f5780601f10610d2457610100808354040283529160200191610d4f565b820191906000526020600020905b815481529060010190602001808311610d3257829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610eb05780601f10610e8557610100808354040283529160200191610eb0565b820191906000526020600020905b815481529060010190602001808311610e9357829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6d906121b7565b60405180910390fd5b60008060606020840151925060408401519150604051905083606001518082016020016040526084360360848337506000806003600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150610fee816104b9565b61102d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102490612237565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561109e57600080fd5b505af11580156110b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d69190611b2c565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561113d57611138838560601c89611605565b61114c565b61114b838560601c896114d9565b5b5050505050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b81526004016111969291906120a0565b600060405180830381600087803b1580156111b057600080fd5b505af11580156111c4573d6000803e3d6000fd5b5050505061121a82600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016112a593929190612069565b602060405180830381600087803b1580156112bf57600080fd5b505af11580156112d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f79190611a34565b50611349826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000808284019050838110156114cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c6906121f7565b60405180910390fd5b8091505092915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016115199291906120a0565b602060405180830381600087803b15801561153357600080fd5b505af1158015611547573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061156b9190611a34565b506115bd826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461173190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016116459291906120a0565b602060405180830381600087803b15801561165f57600080fd5b505af1158015611673573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116979190611a34565b506116e9826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600061177383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061177b565b905092915050565b60008383111582906117c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117ba9190612175565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061181757805160ff1916838001178555611845565b82800160010185558215611845579182015b82811115611844578251825591602001919060010190611829565b5b50905061185291906118c5565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6118e791905b808211156118e35760008160009055506001016118cb565b5090565b90565b6000813590506118f981612412565b92915050565b60008151905061190e81612429565b92915050565b60008135905061192381612440565b92915050565b600082601f83011261193a57600080fd5b813561194d611948826122e1565b6122b4565b9150808252602083016020830185838301111561196957600080fd5b6119748382846123b5565b50505092915050565b60008135905061198c81612457565b92915050565b6000813590506119a18161246e565b92915050565b6000815190506119b68161246e565b92915050565b6000602082840312156119ce57600080fd5b60006119dc848285016118ea565b91505092915050565b6000806000606084860312156119fa57600080fd5b6000611a08868287016118ea565b9350506020611a19868287016118ea565b9250506040611a2a8682870161197d565b9150509250925092565b600060208284031215611a4657600080fd5b6000611a54848285016118ff565b91505092915050565b600060208284031215611a6f57600080fd5b6000611a7d84828501611914565b91505092915050565b60008060408385031215611a9957600080fd5b6000611aa785828601611914565b9250506020611ab8858286016118ea565b9150509250929050565b600060208284031215611ad457600080fd5b600082013567ffffffffffffffff811115611aee57600080fd5b611afa84828501611929565b91505092915050565b600060208284031215611b1557600080fd5b6000611b238482850161197d565b91505092915050565b600060208284031215611b3e57600080fd5b6000611b4c848285016119a7565b91505092915050565b60008060008060808587031215611b6b57600080fd5b6000611b7987828801611992565b9450506020611b8a8782880161197d565b9350506040611b9b878288016118ea565b925050606085013567ffffffffffffffff811115611bb857600080fd5b611bc487828801611929565b91505092959194509250565b611bd981612356565b82525050565b611be881612356565b82525050565b611bf781612368565b82525050565b611c0681612374565b82525050565b611c1581612374565b82525050565b611c2c611c2782612374565b6123f7565b82525050565b6000611c3d8261230d565b611c478185612323565b9350611c578185602086016123c4565b611c6081612401565b840191505092915050565b6000611c768261230d565b611c808185612334565b9350611c908185602086016123c4565b611c9981612401565b840191505092915050565b6000611caf82612318565b611cb98185612345565b9350611cc98185602086016123c4565b611cd281612401565b840191505092915050565b6000611cea603583612345565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611d50601e83612345565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611d90602483612345565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611df6601b83612345565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611e36603783612345565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611e9c602883612345565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611f02603383612345565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600060e083016000830151611f736000860182611bd0565b506020830151611f866020860182612015565b506040830151611f996040860182611bfd565b506060830151611fac6060860182611ff7565b5060808301518482036080860152611fc48282611c32565b91505060a0830151611fd960a0860182611bd0565b5060c0830151611fec60c0860182611ff7565b508091505092915050565b6120008161239e565b82525050565b61200f8161239e565b82525050565b61201e816123a8565b82525050565b61202d816123a8565b82525050565b600061203f8284611c1b565b60208201915081905092915050565b60006020820190506120636000830184611bdf565b92915050565b600060608201905061207e6000830186611bdf565b61208b6020830185611bdf565b6120986040830184612006565b949350505050565b60006040820190506120b56000830185611bdf565b6120c26020830184612006565b9392505050565b600060e0820190506120de600083018a611bdf565b6120eb6020830189612024565b6120f86040830188611c0c565b6121056060830187612006565b81810360808301526121178186611c6b565b905061212660a0830185611bdf565b61213360c0830184612006565b98975050505050505050565b60006020820190506121546000830184611bee565b92915050565b600060208201905061216f6000830184611c0c565b92915050565b6000602082019050818103600083015261218f8184611ca4565b905092915050565b600060208201905081810360008301526121b081611cdd565b9050919050565b600060208201905081810360008301526121d081611d43565b9050919050565b600060208201905081810360008301526121f081611d83565b9050919050565b6000602082019050818103600083015261221081611de9565b9050919050565b6000602082019050818103600083015261223081611e29565b9050919050565b6000602082019050818103600083015261225081611e8f565b9050919050565b6000602082019050818103600083015261227081611ef5565b9050919050565b600060208201905081810360008301526122918184611f5b565b905092915050565b60006020820190506122ae6000830184612006565b92915050565b6000604051905081810181811067ffffffffffffffff821117156122d757600080fd5b8060405250919050565b600067ffffffffffffffff8211156122f857600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006123618261237e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156123e25780820151818401526020810190506123c7565b838111156123f1576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61241b81612356565b811461242657600080fd5b50565b61243281612368565b811461243d57600080fd5b50565b61244981612374565b811461245457600080fd5b50565b6124608161239e565b811461246b57600080fd5b50565b612477816123a8565b811461248257600080fd5b5056fea2646970667358221220dc63c3855c1704d110aa54e63df373be9d5b248c7d482196873d3db123b1e3c464736f6c63430006040033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.Balances(&_ERC20Handler.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.Balances(&_ERC20Handler.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_burnList", arg0)
	return *ret0, err
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCaller) BurnedTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_burnedTokens", arg0)
	return *ret0, err
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.BurnedTokens(&_ERC20Handler.CallOpts, arg0)
}

// BurnedTokens is a free data retrieval call binding the contract method 0x45274738.
//
// Solidity: function _burnedTokens(address ) view returns(uint256)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnedTokens(arg0 common.Address) (*big.Int, error) {
	return _ERC20Handler.Contract.BurnedTokens(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_contractWhitelist", arg0)
	return *ret0, err
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	ret := new(struct {
		OriginChainTokenAddress        common.Address
		DestinationChainID             uint8
		ResourceID                     [32]byte
		LenDestinationRecipientAddress *big.Int
		DestinationRecipientAddress    []byte
		Depositer                      common.Address
		Amount                         *big.Int
	})
	out := ret
	err := _ERC20Handler.contract.Call(opts, out, "_depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xe245386f.
//
// Solidity: function _depositRecords(uint256 ) view returns(address _originChainTokenAddress, uint8 _destinationChainID, bytes32 _resourceID, uint256 _lenDestinationRecipientAddress, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCallerSession) DepositRecords(arg0 *big.Int) (struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             uint8
	ResourceID                     [32]byte
	LenDestinationRecipientAddress *big.Int
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_tokenContractAddressToResourceID", arg0)
	return *ret0, err
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) UseContractWhitelist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "_useContractWhitelist")
	return *ret0, err
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) UseContractWhitelist() (bool, error) {
	return _ERC20Handler.Contract.UseContractWhitelist(&_ERC20Handler.CallOpts)
}

// UseContractWhitelist is a free data retrieval call binding the contract method 0x2521ab41.
//
// Solidity: function _useContractWhitelist() view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) UseContractWhitelist() (bool, error) {
	return _ERC20Handler.Contract.UseContractWhitelist(&_ERC20Handler.CallOpts)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	var (
		ret0 = new(ERC20HandlerDepositRecord)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "getDepositRecord", depositID)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerSession) GetDepositRecord(depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xdb95e75c.
//
// Solidity: function getDepositRecord(uint256 depositID) view returns(ERC20HandlerDepositRecord)
func (_ERC20Handler *ERC20HandlerCallerSession) GetDepositRecord(depositID *big.Int) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositID)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) IsWhitelisted(opts *bind.CallOpts, contractAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Handler.contract.Call(opts, out, "isWhitelisted", contractAddress)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) IsWhitelisted(contractAddress common.Address) (bool, error) {
	return _ERC20Handler.Contract.IsWhitelisted(&_ERC20Handler.CallOpts, contractAddress)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address contractAddress) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) IsWhitelisted(contractAddress common.Address) (bool, error) {
	return _ERC20Handler.Contract.IsWhitelisted(&_ERC20Handler.CallOpts, contractAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x45a104db.
//
// Solidity: function deposit(uint8 destinationChainID, uint256 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(destinationChainID uint8, depositNonce *big.Int, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, destinationChainID, depositNonce, depositer, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteDeposit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeDeposit", data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteDeposit(&_ERC20Handler.TransactOpts, data)
}

// ExecuteDeposit is a paid mutator transaction binding the contract method 0xfc9539cd.
//
// Solidity: function executeDeposit(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteDeposit(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteDeposit(&_ERC20Handler.TransactOpts, data)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResourceIDAndContractAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResourceIDAndContractAddress", resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResourceIDAndContractAddress(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResourceIDAndContractAddress is a paid mutator transaction binding the contract method 0x8a025ce8.
//
// Solidity: function setResourceIDAndContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResourceIDAndContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResourceIDAndContractAddress(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}

var RuntimeBytecode = "0x608060405234801561001057600080fd5b506004361061010b5760003560e01c80636ebcf607116100a2578063c8ba6c8711610071578063c8ba6c87146102dc578063d9caed121461030c578063db95e75c14610328578063e245386f14610358578063fc9539cd1461038e5761010b565b80636ebcf607146102445780637f79bea8146102745780638a025ce8146102a457806395601f09146102c05761010b565b80633af32abf116100de5780633af32abf1461019857806345274738146101c857806345a104db146101f85780636a70d081146102145761010b565b806307b7ed99146101105780630a6d55d81461012c5780632521ab411461015c578063318c136e1461017a575b600080fd5b61012a600480360381019061012591906119bc565b6103aa565b005b61014660048036038101906101419190611a5d565b61044d565b604051610153919061204e565b60405180910390f35b610164610480565b604051610171919061213f565b60405180910390f35b610182610493565b60405161018f919061204e565b60405180910390f35b6101b260048036038101906101ad91906119bc565b6104b9565b6040516101bf919061213f565b60405180910390f35b6101e260048036038101906101dd91906119bc565b61050f565b6040516101ef9190612299565b60405180910390f35b610212600480360381019061020d9190611b55565b610527565b005b61022e600480360381019061022991906119bc565b610845565b60405161023b919061213f565b60405180910390f35b61025e600480360381019061025991906119bc565b610865565b60405161026b9190612299565b60405180910390f35b61028e600480360381019061028991906119bc565b61087d565b60405161029b919061213f565b60405180910390f35b6102be60048036038101906102b99190611a86565b61089d565b005b6102da60048036038101906102d591906119e5565b610a23565b005b6102f660048036038101906102f191906119bc565b610b51565b604051610303919061215a565b60405180910390f35b610326600480360381019061032191906119e5565b610b69565b005b610342600480360381019061033d9190611b03565b610c09565b60405161034f9190612277565b60405180910390f35b610372600480360381019061036d9190611b03565b610dbf565b60405161038597969594939291906120c9565b60405180910390f35b6103a860048036038101906103a39190611ac2565b610ee6565b005b6103b3816104b9565b6103f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e9906121d7565b60405180910390fd5b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60016020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ae906121b7565b60405180910390fd5b60006060600080602085015193506040850151915060405192508460600151905080830160200160405260e4360360e4843760006003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061062a816104b9565b610669576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066090612257565b60405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106cb576106c6818885611156565b6106d8565b6106d781883086611263565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018a60ff1681526020018681526020018381526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600760008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff160217905550604082015181600101556060820151816002015560808201518160030190805190602001906107e59291906117d6565b5060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160050155905050505050505050505050565b60066020528060005260406000206000915054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b60056020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461093f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093690612217565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000806040516020016109969190612033565b60405160208183030381529060405280519060200120826040516020016109bd9190612033565b6040516020818303038152906040528051906020012014610a13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0a90612197565b60405180910390fd5b610a1d8484611392565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401610a6593929190612069565b602060405180830381600087803b158015610a7f57600080fd5b505af1158015610a93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab79190611a34565b50610b09826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60046020528060005260406000206000915090505481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf0906121b7565b60405180910390fd5b610c048383836114d9565b505050565b610c11611856565b600760008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d4f5780601f10610d2457610100808354040283529160200191610d4f565b820191906000526020600020905b815481529060010190602001808311610d3257829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050919050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff1690806001015490806002015490806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610eb05780601f10610e8557610100808354040283529160200191610eb0565b820191906000526020600020905b815481529060010190602001808311610e9357829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905087565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6d906121b7565b60405180910390fd5b60008060606020840151925060408401519150604051905083606001518082016020016040526084360360848337506000806003600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150610fee816104b9565b61102d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102490612237565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663beab71316040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561109e57600080fd5b505af11580156110b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d69190611b2c565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561113d57611138838560601c89611605565b61114c565b61114b838560601c896114d9565b5b5050505050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b81526004016111969291906120a0565b600060405180830381600087803b1580156111b057600080fd5b505af11580156111c4573d6000803e3d6000fd5b5050505061121a82600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016112a593929190612069565b602060405180830381600087803b1580156112bf57600080fd5b505af11580156112d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f79190611a34565b50611349826000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000808284019050838110156114cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c6906121f7565b60405180910390fd5b8091505092915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016115199291906120a0565b602060405180830381600087803b15801561153357600080fd5b505af1158015611547573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061156b9190611a34565b506115bd826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461173190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016116459291906120a0565b602060405180830381600087803b15801561165f57600080fd5b505af1158015611673573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116979190611a34565b506116e9826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461148490919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050565b600061177383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061177b565b905092915050565b60008383111582906117c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117ba9190612175565b60405180910390fd5b5060008385039050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061181757805160ff1916838001178555611845565b82800160010185558215611845579182015b82811115611844578251825591602001919060010190611829565b5b50905061185291906118c5565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600080191681526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6118e791905b808211156118e35760008160009055506001016118cb565b5090565b90565b6000813590506118f981612412565b92915050565b60008151905061190e81612429565b92915050565b60008135905061192381612440565b92915050565b600082601f83011261193a57600080fd5b813561194d611948826122e1565b6122b4565b9150808252602083016020830185838301111561196957600080fd5b6119748382846123b5565b50505092915050565b60008135905061198c81612457565b92915050565b6000813590506119a18161246e565b92915050565b6000815190506119b68161246e565b92915050565b6000602082840312156119ce57600080fd5b60006119dc848285016118ea565b91505092915050565b6000806000606084860312156119fa57600080fd5b6000611a08868287016118ea565b9350506020611a19868287016118ea565b9250506040611a2a8682870161197d565b9150509250925092565b600060208284031215611a4657600080fd5b6000611a54848285016118ff565b91505092915050565b600060208284031215611a6f57600080fd5b6000611a7d84828501611914565b91505092915050565b60008060408385031215611a9957600080fd5b6000611aa785828601611914565b9250506020611ab8858286016118ea565b9150509250929050565b600060208284031215611ad457600080fd5b600082013567ffffffffffffffff811115611aee57600080fd5b611afa84828501611929565b91505092915050565b600060208284031215611b1557600080fd5b6000611b238482850161197d565b91505092915050565b600060208284031215611b3e57600080fd5b6000611b4c848285016119a7565b91505092915050565b60008060008060808587031215611b6b57600080fd5b6000611b7987828801611992565b9450506020611b8a8782880161197d565b9350506040611b9b878288016118ea565b925050606085013567ffffffffffffffff811115611bb857600080fd5b611bc487828801611929565b91505092959194509250565b611bd981612356565b82525050565b611be881612356565b82525050565b611bf781612368565b82525050565b611c0681612374565b82525050565b611c1581612374565b82525050565b611c2c611c2782612374565b6123f7565b82525050565b6000611c3d8261230d565b611c478185612323565b9350611c578185602086016123c4565b611c6081612401565b840191505092915050565b6000611c768261230d565b611c808185612334565b9350611c908185602086016123c4565b611c9981612401565b840191505092915050565b6000611caf82612318565b611cb98185612345565b9350611cc98185602086016123c4565b611cd281612401565b840191505092915050565b6000611cea603583612345565b91507f636f6e7472616374206164647265737320616c72656164792068617320636f7260008301527f726573706f6e64696e67207265736f75726365494400000000000000000000006020830152604082019050919050565b6000611d50601e83612345565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611d90602483612345565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611df6601b83612345565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000611e36603783612345565b91507f7265736f75726365494420616c726561647920686173206120636f727265737060008301527f6f6e64696e6720636f6e747261637420616464726573730000000000000000006020830152604082019050919050565b6000611e9c602883612345565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611f02603383612345565b91507f70726f7669646564206f726967696e436861696e546f6b656e4164647265737360008301527f206973206e6f742077686974656c6973746564000000000000000000000000006020830152604082019050919050565b600060e083016000830151611f736000860182611bd0565b506020830151611f866020860182612015565b506040830151611f996040860182611bfd565b506060830151611fac6060860182611ff7565b5060808301518482036080860152611fc48282611c32565b91505060a0830151611fd960a0860182611bd0565b5060c0830151611fec60c0860182611ff7565b508091505092915050565b6120008161239e565b82525050565b61200f8161239e565b82525050565b61201e816123a8565b82525050565b61202d816123a8565b82525050565b600061203f8284611c1b565b60208201915081905092915050565b60006020820190506120636000830184611bdf565b92915050565b600060608201905061207e6000830186611bdf565b61208b6020830185611bdf565b6120986040830184612006565b949350505050565b60006040820190506120b56000830185611bdf565b6120c26020830184612006565b9392505050565b600060e0820190506120de600083018a611bdf565b6120eb6020830189612024565b6120f86040830188611c0c565b6121056060830187612006565b81810360808301526121178186611c6b565b905061212660a0830185611bdf565b61213360c0830184612006565b98975050505050505050565b60006020820190506121546000830184611bee565b92915050565b600060208201905061216f6000830184611c0c565b92915050565b6000602082019050818103600083015261218f8184611ca4565b905092915050565b600060208201905081810360008301526121b081611cdd565b9050919050565b600060208201905081810360008301526121d081611d43565b9050919050565b600060208201905081810360008301526121f081611d83565b9050919050565b6000602082019050818103600083015261221081611de9565b9050919050565b6000602082019050818103600083015261223081611e29565b9050919050565b6000602082019050818103600083015261225081611e8f565b9050919050565b6000602082019050818103600083015261227081611ef5565b9050919050565b600060208201905081810360008301526122918184611f5b565b905092915050565b60006020820190506122ae6000830184612006565b92915050565b6000604051905081810181811067ffffffffffffffff821117156122d757600080fd5b8060405250919050565b600067ffffffffffffffff8211156122f857600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006123618261237e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156123e25780820151818401526020810190506123c7565b838111156123f1576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61241b81612356565b811461242657600080fd5b50565b61243281612368565b811461243d57600080fd5b50565b61244981612374565b811461245457600080fd5b50565b6124608161239e565b811461246b57600080fd5b50565b612477816123a8565b811461248257600080fd5b5056fea2646970667358221220dc63c3855c1704d110aa54e63df373be9d5b248c7d482196873d3db123b1e3c464736f6c63430006040033"
