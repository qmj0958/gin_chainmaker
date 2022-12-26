package main

import (
	"encoding/json"
	"fmt"
	"gin_chainmaker/chainmaker/chain"
	"gin_chainmaker/chainmaker/saleStockcontract"

	"log"
	"net/http"

	"chainmaker.org/chainmaker/sdk-go/v2/examples"

	"github.com/gin-gonic/gin"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/gin-contrib/cors"
)

const (
	// createContractTimeout = 5
	//claimContractName     = "claim002" //合约名称ddd
	// claimContractName = "saleStock002" //合约名称
	//claimVersion      = "2.0.0"
	// claimVersion = "1.0"
	//claimByteCodePath     = "./chainmaker_config/claim-wasm-demo/rust-fact-2.0.0.wasm"

	sdkConfigOrg1Client1Path = "./chainmaker_config/sdk_configs/sdk_config_org1_client1.yml"
)

var Resp *common.TxResponse

// 定义一个合约查询的参数
var queryParameter string

// 定义合约查询的结果
var queryResult *common.TxResponse

// 定义一个字符串string转map的函数,用作gin框架gin.H的map类型参数
func JsonstringToMap(str string) map[string]interface{} {
	var tempMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)

	}
	return tempMap
}
func main() {
	//连接区块链
	fmt.Println("====================== create client ======================")
	client, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
	if err != nil {
		log.Fatalln(err)
	}
	//创建一个服务
	r := gin.Default()
	r.Use(cors.Default())

	//检查这条长安链是否运行正常
	r.GET("/contact/testChainClient", func(context *gin.Context) {
		result := chain.TestChainClientCheckNewBlockChainConfig()
		context.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})
	//查询saleStock合约
	r.GET("/salestockcontact/:_queryParameter", func(context *gin.Context) {
		fmt.Println("====================== 查看数值 ======================")
		queryParameter = context.Param("_queryParameter")
		getContractResult := saleStockcontract.SaleStockEVMGet(client, queryParameter, true)

		fmt.Printf("Byte resp: %+v\n", getContractResult)

		//取出查询结果的对象的result属性
		//getContractResultByte := queryResult.ContractResult.Result
		//把字节数组转为为字符串
		//getContractResultString := string(getContractResultByte)
		//fmt.Printf("Byte resp: %+v\n", getContractResultString)

		//定义一个Map，以便gin框架输出json数据结构
		//var getContractResultMap = make(map[string]interface{})
		//getContractResultMap = JsonstringToMap(getContractResultString)
		//context.JSON(http.StatusOK, getContractResultMap)
		context.JSON(http.StatusOK, gin.H{
			"hash":           getContractResult[0],
			"factory":        getContractResult[1],
			"style":          getContractResult[2],
			"weight":         getContractResult[3],
			"logo":           getContractResult[4],
			"silver_content": getContractResult[5],
			"cert_id":        getContractResult[6],
			"cert_img":       getContractResult[7],
			"check_img":      getContractResult[8],
			"factory_add":    getContractResult[9],
			"time":           getContractResult[10],
		})
	})

	//前端给后端传递json
	r.POST("/json", func(context *gin.Context) {
		//request.body
		//[]byte
		data, _ := context.GetRawData()
		var m map[string]interface{}
		//包装为json数据[]byte
		_ = json.Unmarshal(data, &m)
		fmt.Println(m)

		saleStockcontract.SaleStockEVMSet(client, "setData", m, true)
		//***上面代码为执行合约
		context.JSON(http.StatusOK, m)
	})

	err1 := r.Run() // 监听并在8080上启动服务
	if err1 != nil {
		// log.Fatalln(err)
	}

}
