package main

import (
	"encoding/json"
	"fmt"
	"gin_chainmaker/chainmaker/chain"
	"gin_chainmaker/chainmaker/saleStockcontract"
	"path"

	"log"
	"net/http"

	"chainmaker.org/chainmaker/sdk-go/v2/examples"

	"github.com/gin-gonic/gin"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/gin-contrib/cors"
)

const (
	sdkConfigOrg1Client1Path = "./chainmaker_config/sdk_configs/sdk_config_org1_client1.yml"
)

var Resp *common.TxResponse

// 定义一个合约查询的参数
var queryParameter string

// 定义合约查询的结果
var queryResult *common.TxResponse

func main() {
	//连接区块链
	fmt.Println("====================== create client ======================")
	client, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
	if err != nil {
		log.Fatalln(err)
	}
	//创建一个服务
	r := gin.Default()
	//跨域
	r.Use(cors.Default())

	//检查这条长安链是否运行正常
	r.GET("/chainclient/test", _testChainClient)
	//查询saleStock合约
	r.GET("/salestockcontract/:_queryParameter", func(context *gin.Context) {
		fmt.Println("====================== 查看数值 ======================")
		queryParameter = context.Param("_queryParameter")
		getContractResult, resp_txid, resp_txtimestamp, resp_txblockheight := saleStockcontract.SaleStockEVMGet(client, queryParameter, true)
		fmt.Printf("Byte resp: %+v\n", getContractResult)
		context.JSON(http.StatusOK, gin.H{
			"hash":               getContractResult[0],
			"factory":            getContractResult[1],
			"style":              getContractResult[2],
			"weight":             getContractResult[3],
			"logo":               getContractResult[4],
			"silver_content":     getContractResult[5],
			"cert_id":            getContractResult[6],
			"cert_img":           getContractResult[7],
			"check_img":          getContractResult[8],
			"factory_add":        getContractResult[9],
			"time":               getContractResult[10],
			"resp_txid":          resp_txid,
			"resp_txtimestamp":   resp_txtimestamp,
			"resp_txblockheight": resp_txblockheight,
		})
	})

	//前端给后端传递json
	r.POST("/invoke", func(context *gin.Context) {

		data, _ := context.GetRawData()
		var m map[string]interface{}
		//包装为json数据[]byte
		_ = json.Unmarshal(data, &m)
		fmt.Println(m)

		saleStockcontract.SaleStockEVMSet(client, "setData", m, true)
		//***上面代码为执行合约
		context.JSON(http.StatusOK, m)
	})
	//上传证书图片(一张图片)
	r.POST("uploadcertimg", _uploadcertimg)
	//上传产品核对图片（多张）
	r.POST("uploadcheckimg", _uploadcheckimg)

	err1 := r.Run() // 监听并在8080上启动服务
	if err1 != nil {
		// log.Fatalln(err)
	}

}

// 检查链是否正常
func _testChainClient(c *gin.Context) {
	result := chain.TestChainClientCheckNewBlockChainConfig()
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// 上传证书图片(一张图片)
func _uploadcertimg(c *gin.Context) {
	//
	msg := c.Param("data")
	fmt.Println(msg)
	//

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("获取图片:", err)
		return
	}
	// file.Size
	// todo 修改上传图片存储路径
	var commonPath = "./images/cert_img/"
	dst := path.Join(commonPath, file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println("upload error:", err)
		return
	}
	// fmt.Println("upload success!")
	c.JSON(200, gin.H{"msg": "证书图片上传成功"})
}

// 上传产品核对图片（多张）
func _uploadcheckimg(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("get picture:", err)
		return
	}
	var commonPath = "./images/check_img/"
	files := form.File["pictures"]
	for _, file := range files {
		fmt.Println("文件名:", file.Filename)
		dst := path.Join(commonPath, file.Filename)
		c.SaveUploadedFile(file, dst)
	}
	// fmt.Println("upload success!")
	c.JSON(200, gin.H{"msg": "产品核对图片上传成功"})
}
