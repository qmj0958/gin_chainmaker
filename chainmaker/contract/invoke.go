package contract

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"fmt"
)

// 合约执行
//func testUserContractClaim() {
//	fmt.Println("====================== create client ======================")
//	client, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	fmt.Println("====================== 调用合约 ======================")
//	fileHash, err := ContractInvoke(client, "save", true, _fileHash, _fileName, _curTime)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(fileHash)
//}

func ContractInvoke(client *sdk.ChainClient,
	method string, withSyncResult bool, _fileHash string, _fileName string, _curTime string) (string, error) {

	//curTime := strconv.FormatInt(time.Now().Unix(), 10)
	curTime := _curTime

	//fileHash := uuid.GetUUID()
	fileHash := _fileHash
	fileName := _fileName
	kvs := []*common.KeyValuePair{
		{
			Key:   "time",
			Value: []byte(curTime),
		},
		{
			Key:   "file_hash",
			Value: []byte(fileHash),
		},
		{
			Key: "file_name",
			//Value: []byte(fmt.Sprintf("file_%s", curTime)),
			Value: []byte(fileName),
		},
	}

	err := invokeUserContract(client, claimContractName, method, "", kvs, withSyncResult)
	if err != nil {
		return "", err
	}

	return fileHash, nil
}

func invokeUserContract(client *sdk.ChainClient, contractName, method, txId string,
	kvs []*common.KeyValuePair, withSyncResult bool) error {

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
