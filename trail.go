package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type C1Info struct{}

func (c *C1Info) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Successsssssssssss"))
}

func (c *C1Info) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//function, args := stub.GetFunctionAndParameters()
	return shim.Success([]byte("Successsssssssssss"))
}

func main() {
	err := shim.Start(new(C1Info))
	if err != nil {
		fmt.Println("Unable to start the chaincode")
	}
}
