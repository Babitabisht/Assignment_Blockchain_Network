/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

/*


 */

 package main

 //WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
 //calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
 //to be modified as well with the new ID of chaincode_example02.
 //chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
 //hard-coding.
 
 import (
	"bytes"
	 "fmt"
	"encoding/json"
	"time"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/protos/peer"
	 pb "github.com/hyperledger/fabric/protos/peer"
 )
 

 var logger = shim.NewLogger("EV chaincode")
 
 type DocChaincode struct {
 }
 
 // Success HTTP 2xx with a payload
 func Success(rc int32, doc string, payload []byte) peer.Response {
	 return peer.Response{
		 Status:  rc,
		 Message: doc,
		 Payload: payload,
	 }
 }
 
 // Error HTTP 4xx or 5xx with an error message
 func Error(rc int32, doc string) peer.Response {
	 logger.Errorf("Error %d = %s", rc, doc)
	 return peer.Response{
		 Status:  rc,
		 Message: doc,
	 }
 }
 
 func (t *DocChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	 fmt.Printf("In init function !")
	 return shim.Success(nil)
 }
 
 // Invoking the chaincode
 func (t *DocChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	 logger.Info("ex02 Invoke")
	 function, args := stub.GetFunctionAndParameters()
			fmt.Println(function, args)
	 if function == "registerDoc" {
		 return t.registerDoc(stub, args)
	 }else if  function == "getAllDocs" {
		  return t.getAllDocs(stub, args)
	 }else if  function == "getfileInfoByHash" {
		  return t.getfileInfoByHash(stub, args)
	 }else{
		return shim.Error("Invalid invoke function name.")
	 }
 }

 
func (t *DocChaincode) registerDoc(stub shim.ChaincodeStubInterface, args[]string) pb.Response {
	 var docs []Document
	 docsFromArgs := args[0]
	 err := json.Unmarshal([]byte(docsFromArgs), &docs)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 TxTimestamp, _ := stub.GetTxTimestamp()
	 timestr := time.Unix(TxTimestamp.GetSeconds(), 0)
	 transactionId := stub.GetTxID()
	 for i := range docs {
		 doc := Document{
			ObjectType 		: "Document",
			Fieldname 		:  docs[i].Fieldname,
			Originalname 	: docs[i].Originalname,
			Encoding		: docs[i].Encoding,
			Mimetype		: docs[i].Mimetype,
			Destination		: docs[i].Destination,
			Filename		: docs[i].Filename,
			Path			: docs[i].Path,
			Size			: docs[i].Size,
			DocumentID		:  docs[i].DocumentID,
			FileHash        :  docs[i].FileHash,
			UploadedTime	: timestr,
			Txid			: transactionId}

		 docAsBytes, err :=json.Marshal(doc)		
		if err != nil {
			 return shim.Error(err.Error())
		 }
		stub.PutState(doc.DocumentID, docAsBytes)
	 }

	 msg := fmt.Sprintf("Successful operation !");
	return shim.Success([]byte(msg))
}
 

func (t *DocChaincode) getAllDocs(stub shim.ChaincodeStubInterface, args[]string) pb.Response {

	queryString := fmt.Sprintf("{\"selector\":{\"Document\":\"Document\"}}")
		// Print the received query on the console
		logger.Info("Query JSON=%s \n\n", queryString)

		// GetQueryResult
		resultsIterator, err := stub.GetQueryResult(queryString)
	
		if err != nil {
		return shim.Error(err.Error())
		}
		defer resultsIterator.Close()
	
		// buffer is a JSON array containing QueryResults
		var buffer bytes.Buffer
		buffer.WriteString("[")
	
		bArrayMemberAlreadyWritten := false
		for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
	
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
		}
		buffer.WriteString("]")
	
		fmt.Printf("- getAllManufacturedBattery:\n%s\n", buffer.String())
	
		return shim.Success(buffer.Bytes())

}


func (t *DocChaincode) getfileInfoByHash(stub shim.ChaincodeStubInterface, args[]string) pb.Response {

	queryString := fmt.Sprintf("{\"selector\":{\"Document\":\"Document\",\"FileHash\" : \"%s\" }}", args[0])
		// Print the received query on the console
		logger.Info("Query JSON=%s \n\n", queryString)

		// GetQueryResult
		resultsIterator, err := stub.GetQueryResult(queryString)
	
		if err != nil {
		return shim.Error(err.Error())
		}
		defer resultsIterator.Close()
	
		// buffer is a JSON array containing QueryResults
		var buffer bytes.Buffer
		buffer.WriteString("[")
	
		bArrayMemberAlreadyWritten := false
		for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
	
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
		}
		buffer.WriteString("]")
	
		fmt.Printf("- getAllManufacturedBattery:\n%s\n", buffer.String())
	
		return shim.Success(buffer.Bytes())

}


 
 func main() {
	 err := shim.Start(new(DocChaincode))
	 if err != nil {
		 fmt.Printf("Error starting Simple chaincode: %s", err)
	 }
 
 }
