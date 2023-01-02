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

func SaleStockEVMGet(client *sdk.ChainClient, hash string, withSyncResult bool) ([]interface{}, string, int64, uint64) {

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

	result, resp_txid, resp_txtimestamp, resp_txblockheight, err := invokeUserContractWithResult(client, saleStockContractName, method, "", kvs, withSyncResult)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("val: %v\n", result)

	val, err := myAbi.Unpack("getData", result)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("val: %d\n", val)
	// fmt.Printf("val: %v\n", val)
	// fmt.Println(val)
	return val, resp_txid, resp_txtimestamp, resp_txblockheight
}

func invokeUserContractWithResult(client *sdk.ChainClient, contractName, method, txId string,
	kvs []*common.KeyValuePair, withSyncResult bool) ([]byte, string, int64, uint64, error) {

	resp, err := client.InvokeContract(contractName, method, txId, kvs, -1, withSyncResult)
	if err != nil {
		return nil, resp.TxId, resp.TxTimestamp, resp.TxBlockHeight, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return nil, resp.TxId, resp.TxTimestamp, resp.TxBlockHeight, fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}
	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(resp.TxId)
	fmt.Println(resp.TxBlockHeight) //TxTimestamp
	return resp.ContractResult.Result, resp.TxId, resp.TxTimestamp, resp.TxBlockHeight, nil
}
