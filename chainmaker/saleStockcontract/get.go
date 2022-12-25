package saleStockcontract

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"chainmaker.org/chainmaker/common/v2/evmutils/abi"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

func SaleStockEVMGet(client *sdk.ChainClient, hash string, withSyncResult bool) []interface{} {

	abiJson, err := ioutil.ReadFile(saleStockABIPath)
	if err != nil {
		log.Fatalln(err)
	}

	myAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		log.Fatalln(err)
	}

	method := "getData"
	dataByte, err := myAbi.Pack(method, hash)
	if err != nil {
		log.Fatalln(err)
	}

	dataString := hex.EncodeToString(dataByte)

	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}

	result, err := invokeUserContractWithResult(client, saleStockContractName, method, "", kvs, withSyncResult)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("val: %v\n", result)

	val, err := myAbi.Unpack("getData", result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("val: %d\n", val)
	fmt.Printf("val: %v\n", val)
	fmt.Println(val)
	return val
}

func invokeUserContractWithResult(client *sdk.ChainClient, contractName, method, txId string,
	kvs []*common.KeyValuePair, withSyncResult bool) ([]byte, error) {

	resp, err := client.InvokeContract(contractName, method, txId, kvs, -1, withSyncResult)
	if err != nil {
		return nil, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return nil, fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}

	return resp.ContractResult.Result, nil
}
