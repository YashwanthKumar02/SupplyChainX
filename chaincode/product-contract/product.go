package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ProductContract struct {
	contractapi.Contract
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

func (pc *ProductContract) CreateProduct(ctx contractapi.TransactionContextInterface, id, name, description, timestamp string) error {
	product := Product{ID: id, Name: name, Description: description, Timestamp: timestamp}
	productBytes, _ := json.Marshal(product)
	return ctx.GetStub().PutState(id, productBytes)
}

func (pc *ProductContract) QueryProduct(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
	productBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product: %v", err)
	}
	product := new(Product)
	_ = json.Unmarshal(productBytes, product)
	return product, nil
}

func main() {
	chaincode, _ := contractapi.NewChaincode(new(ProductContract))
	chaincode.Start()
}