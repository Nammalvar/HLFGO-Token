package main
/**
 * tokenv2
 * Shows the  
 *    A) Use of Logger
 **/
import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// peer.Response is in the peer package
    "github.com/hyperledger/fabric/protos/peer"

)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// V2

// ChaincodeName - Create an instance of the Logger
const ChaincodeName = "tokenv2"
var logger = shim.NewLogger(ChaincodeName)

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	// fmt.Println("Init executed")

	logger.Debug("Init executed")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//fmt.Println("Invoke executed ")
	logger.Debug("Init executed")

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode.")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		//fmt.Printf("Error starting chaincode: %s", err)
		// V2
		logger.Error("Error starting chaincode: %s", err)
	}
}