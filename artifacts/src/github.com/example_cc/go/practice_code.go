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

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_cc0")

type bank struct {
	IDBank            string
	Name              string
	EstablishmentYear int
	OriginCountry     string
	BusinessCountries []string
	ClientList        []client
}

type client struct {
	IDClient          string
	Name              string
	Surname           string
	Email             string
	MoneyOnAccount    float64
	ReceivedTransfers []float64
	AllowCredit       bool
	Credits           []credit
}

type transaction struct {
	IDTransaction string
	Date          time.Time
	IDSender      string
	IDReceiver    string
	Amount        float64
}

type credit struct {
	IDCredit        string
	ApprovalDate    time.Time
	EndDate         time.Time
	RateAmount      float64
	Interest        float64
	NumOfRates      int
	NumOfPayedRates int
	AmountOfCredit  float64
}

// Global variables for ID
var creditId int
var clientId int
var transId int

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Init ###########")

	var credit1 = credit{"cre1", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), time.Date(2010, 11, 10, 23, 0, 0, 0, time.UTC), 7700, 0.1, 10, 10, 70000}
	var credit2 = credit{"cre2", time.Date(2019, 8, 4, 26, 0, 0, 0, time.UTC), time.Date(2020, 8, 4, 26, 0, 0, 0, time.UTC), 11000, 0.2, 15, 4, 160500}
	var credit3 = credit{"cre3", time.Date(2020, 1, 17, 12, 0, 0, 0, time.UTC), time.Date(2021, 1, 17, 12, 0, 0, 0, time.UTC), 9900, 0.1, 10, 5, 90000}
	var credit4 = credit{"cre4", time.Date(2018, 6, 7, 7, 0, 0, 0, time.UTC), time.Date(2019, 6, 7, 7, 0, 0, 0, time.UTC), 11000, 0.1, 10, 9, 100000}
	creditId = 5

	var receivedTransfer1 = make([]float64, 0, 20)
	var receivedTransfer2 = make([]float64, 0, 20)
	var receivedTransfer3 = make([]float64, 0, 20)
	var receivedTransfer4 = make([]float64, 0, 20)
	receivedTransfer1 = append(receivedTransfer1, 20000, 40000, 5000)
	receivedTransfer2 = append(receivedTransfer2, 20000, 30000, 7000)
	receivedTransfer3 = append(receivedTransfer3, 20000)
	receivedTransfer4 = append(receivedTransfer4, 20000)

	var credits1 = make([]credit, 0, 20)
	credits1 = append(credits1, credit1)
	var credits2 = make([]credit, 0, 20)
	var credits3 = make([]credit, 0, 20)
	credits3 = append(credits3, credit3)
	var credits4 = make([]credit, 0, 20)
	credits4 = append(credits4, credit4)

	var client1 = client{"cli1", "Nikolina", "Tomic", "tomicN@gmail.com", 3200000, receivedTransfer1, true, credits1}
	var client2 = client{"cli2", "Nadia", "Radic", "nadia.ra@gmail.com", 54200, receivedTransfer2, true, credits2}
	var client3 = client{"cli3", "Mirko", "Ivic", "mirza96@gmail.com", 450000, receivedTransfer3, false, credits3}
	var client4 = client{"cli4", "Ivan", "Snobista", "smarac@gmail.com", 0, receivedTransfer4, false, credits4}
	clientId = 5

	var countBank1 = make([]string, 0, 20)
	var countBank2 = make([]string, 0, 20)
	var countBank3 = make([]string, 0, 20)
	countBank1 = append(countBank1, "Germany", "Serbia", "Italy")
	countBank2 = append(countBank2, "Bosnia", "Serbia", "France", "Israel")
	countBank3 = append(countBank3, "Japan", "China", "Rusia", "Montenegro")
	var cliBank1 = make([]client, 0, 20)
	var cliBank2 = make([]client, 0, 20)
	var cliBank3 = make([]client, 0, 20)
	cliBank1 = append(cliBank1, client1, client3)
	cliBank2 = append(cliBank2, client2, client4)
	cliBank3 = append(cliBank3, client2, client3, client4)

	var bank1 = bank{"bank1", "National bank", 1934, "Germany", countBank1, cliBank1}
	var bank2 = bank{"bank2", "Opportunity bank", 1957, "Serbia", countBank2, cliBank2}
	var bank3 = bank{"bank3", "Sberbank", 2003, "Japan", countBank3, cliBank3}

	var transaction1 = transaction{"trans1", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), "cli1", "cli3", 20000}
	var transaction2 = transaction{"trans2", time.Date(2019, 8, 4, 26, 0, 0, 0, time.UTC), "cli1", "cli2", 30000}
	var transaction3 = transaction{"trans3", time.Date(2020, 1, 17, 12, 0, 0, 0, time.UTC), "cli4", "cli2", 7000}
	var transaction4 = transaction{"trans4", time.Date(2018, 6, 7, 7, 0, 0, 0, time.UTC), "cli1", "cli4", 20000}
	transId = 5

	// Write the state to the ledger
	ajson, _ := json.Marshal(client1)
	err := stub.PutState(client1.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(client2)
	err = stub.PutState(client2.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(client3)
	err = stub.PutState(client3.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(client4)
	err = stub.PutState(client4.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(bank1)
	err = stub.PutState(bank1.IDBank, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(bank2)
	err = stub.PutState(bank2.IDBank, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(bank3)
	err = stub.PutState(bank3.IDBank, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(transaction1)
	err = stub.PutState(transaction1.IDTransaction, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(transaction2)
	err = stub.PutState(transaction2.IDTransaction, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(transaction3)
	err = stub.PutState(transaction3.IDTransaction, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(transaction4)
	err = stub.PutState(transaction4.IDTransaction, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(credit1)
	err = stub.PutState(credit1.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(credit2)
	err = stub.PutState(credit2.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(credit3)
	err = stub.PutState(credit3.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(credit4)
	err = stub.PutState(credit4.IDCredit, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "query" {
		return t.query(stub, args)
	}
	if function == "transfer" {
		return t.transfer(stub, args)
	}
	if function == "getCredit" {
		return t.getCredit(stub, args)
	}
	if function == "payRate" {
		return t.payRate(stub, args)
	}
	if function == "addClient" {
		return t.addClient(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) addClient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, surname, email string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expected 4 arguments")
	}
	name = args[0]
	surname = args[1]
	email = args[2]
	moneyOnAccount, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error("Error! Money on account is float64")
	}
	clientKey := "cli" + strconv.Itoa(clientId)
	clientId = clientId + 1
	var receivedTransfers = make([]float64, 0, 20)
	var credits = make([]credit, 0, 20)
	var newClient = client{clientKey, name, surname, email, moneyOnAccount, receivedTransfers, true, credits}

	ajson, errMarshal := json.Marshal(newClient)
	if errMarshal != nil {
		return shim.Error("Cannot marshal new client")
	}
	err = stub.PutState(newClient.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var IDSender, IDReceiver, minus string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expected 4 arguments")
	}

	IDSender = args[0]
	IDReceiver = args[1]
	amount, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return shim.Error("Error! Amount is float64")
	}

	minus = args[3]

	// load clients
	senderB, err := stub.GetState(IDSender)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + IDSender + "\"}"
		return shim.Error(jsonResp)
	}
	if senderB == nil || len(senderB) == 0 {
		jsonResp := "{\"Error\":\" " + IDSender + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}
	sender := client{}
	err = json.Unmarshal(senderB, &sender)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	receiverB, err := stub.GetState(IDReceiver)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + IDReceiver + "\"}"
		return shim.Error(jsonResp)
	}
	if receiverB == nil || len(receiverB) == 0 {
		jsonResp := "{\"Error\":\" " + IDReceiver + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}
	receiver := client{}
	err = json.Unmarshal(receiverB, &receiver)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	newValueOnAccountSender := sender.MoneyOnAccount - amount
	newValueOnAccountReceiver := receiver.MoneyOnAccount + amount

	if minus != "true" && newValueOnAccountSender < 0 {
		return shim.Error("You don't have enough money on your account! Transfer failed")
	}

	if minus == "true" {
		var sumTransferMoney float64
		sumTransferMoney = 0
		senderTransfers := sender.ReceivedTransfers
		numOfTransfers := len(sender.ReceivedTransfers)
		for _, transfer := range senderTransfers {
			sumTransferMoney += transfer
		}
		allowedMinus := sumTransferMoney / float64(numOfTransfers)
		if numOfTransfers < 1 {
			allowedMinus = 0
		}
		if newValueOnAccountSender*(-1) > allowedMinus {
			return shim.Error("You oversteped allowed minus! Transfer failed")
		}
	}

	sender.MoneyOnAccount = newValueOnAccountSender
	ajson, errMarshal := json.Marshal(sender)
	if errMarshal != nil {
		return shim.Error("Cannot marshal client sender")
	}
	err = stub.PutState(sender.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	receiver.ReceivedTransfers = append(receiver.ReceivedTransfers, amount)
	receiver.MoneyOnAccount = newValueOnAccountReceiver
	ajson, errMarshal = json.Marshal(receiver)
	if errMarshal != nil {
		return shim.Error("Cannot marshal client receiver")
	}
	err = stub.PutState(receiver.IDClient, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	transKey := "trans" + strconv.Itoa(transId)
	transId = transId + 1
	var trans = transaction{transKey, time.Now(), sender.IDClient, receiver.IDClient, amount}
	ajson, errMarshal = json.Marshal(trans)
	if errMarshal != nil {
		return shim.Error("Cannot marshal transaction")
	}
	err = stub.PutState(trans.IDTransaction, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) getCredit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var IDClient string
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expected 4 arguments")
	}

	IDClient = args[0]

	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error("Error! Amount is float64")
	}
	numOfRates, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return shim.Error("Error! Number of rates is float64")
	}
	interest, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error("Error! Interest is float64")
	}

	// load client
	clientB, err := stub.GetState(IDClient)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + IDClient + "\"}"
		return shim.Error(jsonResp)
	}
	if clientB == nil || len(clientB) == 0 {
		jsonResp := "{\"Error\":\" " + IDClient + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}
	clientCredit := client{}
	err = json.Unmarshal(clientB, &clientCredit)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	if clientCredit.AllowCredit == true {
		var sumTransferMoney float64
		sumTransferMoney = 0
		clientTransfers := clientCredit.ReceivedTransfers
		numOfTransfers := len(clientCredit.ReceivedTransfers)
		for _, transfer := range clientTransfers {
			sumTransferMoney += transfer
		}
		avgPaymant := sumTransferMoney / float64(numOfTransfers)
		if numOfTransfers < 1 {
			avgPaymant = 0
		}
		allowedCredit := avgPaymant * 5

		if amount > allowedCredit {
			return shim.Error("Requested loan is too high! Loan denied")
		}

		rateAmount := amount / numOfRates * (float64(1) + interest)
		// s := fmt.Sprintf("%f", rateAmount)
		// return shim.Error("Prosek" + s)
		duration, _ := time.ParseDuration("8766h")
		creditKey := "cre" + strconv.Itoa(creditId)
		creditId = creditId + 1
		var newCredit = credit{creditKey, time.Now(), time.Now().Add(duration), rateAmount, interest, int(numOfRates), 0, amount}

		clientCredit.Credits = append(clientCredit.Credits, newCredit)
		clientCredit.AllowCredit = false
		ajson, errMarshal := json.Marshal(clientCredit)
		if errMarshal != nil {
			return shim.Error("Cannot marshal client")
		}
		err = stub.PutState(clientCredit.IDClient, ajson)
		if err != nil {
			return shim.Error(err.Error())
		}

		ajson, errMarshal = json.Marshal(newCredit)
		if errMarshal != nil {
			return shim.Error("Cannot marshal credit")
		}
		errPut := stub.PutState(newCredit.IDCredit, ajson)
		if errPut != nil {
			return shim.Error(errPut.Error())
		}

		return shim.Success(nil)
	}
	return shim.Error("You already have an unpaid loan")

}

func (t *SimpleChaincode) payRate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var IDClient string
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expected 2 arguments")
	}

	IDClient = args[0]

	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error("Error! Amount is float64")
	}

	// load client
	clientB, err := stub.GetState(IDClient)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + IDClient + "\"}"
		return shim.Error(jsonResp)
	}
	if clientB == nil || len(clientB) == 0 {
		jsonResp := "{\"Error\":\" " + IDClient + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}
	clientRate := client{}
	err = json.Unmarshal(clientB, &clientRate)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	clientCredits := clientRate.Credits
	var unpayedCredit credit
	for _, credits := range clientCredits {
		if credits.NumOfPayedRates != credits.NumOfRates {
			unpayedCredit = credits //nadjen neplaceni kredit
		}
	}

	// load credit
	creditB, err := stub.GetState(unpayedCredit.IDCredit)
	if err != nil {
		jsonResp := "There is no unpayed credit"
		return shim.Error(jsonResp)
	}
	if creditB == nil || len(creditB) == 0 {
		jsonResp := "There is no unpayed credit"
		return shim.Error(jsonResp)
	}
	unpayedCredit = credit{}
	err = json.Unmarshal(creditB, &unpayedCredit)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	//loan found
	if amount != unpayedCredit.RateAmount {
		s := fmt.Sprintf("%f", unpayedCredit.RateAmount)
		js := "Rate for " + unpayedCredit.IDCredit + " is " + s
		return shim.Error(js)
	}
	//correct rate for loan
	unpayedCredit.NumOfPayedRates = unpayedCredit.NumOfPayedRates + 1
	if unpayedCredit.NumOfPayedRates == unpayedCredit.NumOfRates {
		clientRate.AllowCredit = true
	}

	ajson, errMarshal := json.Marshal(unpayedCredit)
	if errMarshal != nil {
		return shim.Error("Cannot marshal credit")
	}
	errPut := stub.PutState(unpayedCredit.IDCredit, ajson)
	if errPut != nil {
		return shim.Error(errPut.Error())
	}

	//delete old loan in client
	for i := 0; i < len(clientRate.Credits); i++ {
		if clientRate.Credits[i].IDCredit == unpayedCredit.IDCredit {
			clientRate.Credits = append(clientRate.Credits[:i], clientRate.Credits[i+1:]...)
		}
	}

	//add new loan in client with added 1 more payed rate
	clientRate.Credits = append(clientRate.Credits, unpayedCredit)

	ajson, errMarshal = json.Marshal(clientRate)
	if errMarshal != nil {
		return shim.Error("Cannot marshal client")
	}
	errPut = stub.PutState(clientRate.IDClient, ajson)
	if errPut != nil {
		return shim.Error(errPut.Error())
	}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
