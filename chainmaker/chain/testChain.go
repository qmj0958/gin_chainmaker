package chain

import (
	"chainmaker.org/chainmaker/sdk-go/v2/examples"
	"log"
)

const (
	sdkConfigOrg1Client1Path = "./chainmaker_config/sdk_configs/sdk_config_org1_client1.yml"
)

// 检查长安链是否运行正常
func TestChainClientCheckNewBlockChainConfig() string {
	client, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.CheckNewBlockChainConfig()
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("check new block chain config: ok")
	return "check new block chain config: ok"
}
