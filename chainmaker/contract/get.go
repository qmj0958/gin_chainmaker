package contract

import (
	"fmt"
	"log"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainmaker.org/chainmaker/sdk-go/v2/examples"
)

const (
//createContractTimeout = 5
//claimContractName     = "claim002"
//claimVersion          = "2.0.0"
//claimByteCodePath     = "./chainmaker_config/claim-wasm-demo/rust-fact-2.0.0.wasm"

// sdkConfigOrg1Client1Path = "./chainmaker_config/sdk_configs/sdk_config_org1_client1.yml"
)

// 对合约执行查询
func ContractGet(_queryParameter string) (resp *common.TxResponse) {
	fmt.Println("====================== create client ======================")
	client, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println("====================== 执行合约查询接口 ======================")
	//txId := "d5224d3b7c1f435f911a6d2a1dc6c963"
	//fileHash = txId[len(txId)/2:]
	//fileHash := "d5224d3b7c1f435f911a6d2a1dc6c963"
	//fileHash := _queryParameter
	fileHash := "aa777"
	kvs := []*common.KeyValuePair{
		{
			//Key:   "file_hash",
			Key: "_hash",
			//Value: []byte(fileHash),
			Value: []byte(fileHash),
			//Value: 'd5224d3b7c1f435f911a6d2a1dc6c963',
		},
	}
	//return contractClaimQuery(client, "find_by_file_hash", kvs)
	return contractClaimQuery(client, "getData", kvs)

}

func contractClaimQuery(client *sdk.ChainClient, method string, kvs []*common.KeyValuePair) (resp *common.TxResponse) {
	resp, err := client.QueryContract(claimContractName, method, kvs, -1)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("QUERY claim contract resp: %+v\n", resp)
	// fmt.Printf("QUERY claim contract resp: %s\n", resp)
	return resp
	//queryResult = resp
	//fmt.Printf("QUERY claim contract resp: %+v\n", resp)
}
