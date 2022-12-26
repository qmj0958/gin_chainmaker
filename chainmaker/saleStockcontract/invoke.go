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

func SaleStockEVMSet(client *sdk.ChainClient, method string, m map[string]interface{}, withSyncResult bool) {
	abiJson, err := ioutil.ReadFile(saleStockABIPath)
	if err != nil {
		log.Fatalln(err)
	}

	myAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		log.Fatalln(err)
	}

	// method := "setData"
	fmt.Println("====================== 0 ======================")
	//定义data

	//
	//定义mapdata3
	// var data3 = make(map[string]interface{})
	// data3["_hash"] = "abc77"
	// data3["_goodsId"] = "sz308o"
	// data3["_style"] = "shouzhuo"
	// data3["_weight"] = 99
	// data3["_pichash"] = "tupianjpg"
	// data3["_price"] = 650
	// data3["_shifu"] = "haisheng"
	// hash := m["hash"].(string)
	// 	goodsId := m["goodsId"].(string)
	// 	style := m["style"].(string)
	// 	weight := m["weight"].(string)
	// 	pichash := m["pichash"].(string)
	// 	price := m["price"].(string)
	// 	shifu := m["shifu"].(string)

	//
	fmt.Println("====================== 1 ======================")
	// dataByte, err := myAbi.Pack(method, data3)
	dataByte, err := myAbi.Pack(method, m["hash"], m["factory"], m["style"], m["weight"], m["logo"], m["silver_content"], m["cert_id"], m["cert_img"], m["check_img"], m["factory_add"])
	if err != nil {
		log.Fatalln(err)
	}

	dataString := hex.EncodeToString(dataByte)
	fmt.Println("====================== 2 ======================")
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	fmt.Println("====================== 3 ======================")
	err = invokeUserContract(client, saleStockContractName, method, "", kvs, withSyncResult)
	if err != nil {
		log.Fatalln(err)
	}
}

func invokeUserContract(client *sdk.ChainClient, contractName, method, txId string, kvs []*common.KeyValuePair, withSyncResult bool) error {

	resp, err := client.InvokeContract(contractName, method, txId, kvs, -1, withSyncResult)
	if err != nil {
		return err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}

	if !withSyncResult {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.ContractResult.Result)
	} else {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
	}

	return nil
}
