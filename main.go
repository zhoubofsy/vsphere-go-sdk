package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"vsphere-go-sdk/cis"
	"vsphere-go-sdk/common"
	"vsphere-go-sdk/content"
)

func common_test() {

	url := "http://localhost:8088/v1/user/login"

	client := common.NewClient(url)
	//headers:=make(map[string]string)
	//headers["session"]="test"
	method := "GET"
	//testBody:="gdhsdsdsghd"
	//body,_:=json.Marshal(testBody)
	resBody, _ := client.SendRequest(url, nil, nil, method)
	log.Debug(resBody)

}

func content_test()  {
	client := common.NewClient("https://128.179.0.241/rest/")
	log.Debug(client)
	sid:=""
	c:=content.NewContent(client,sid)
	log.Debug(*c)
	shandle:=c.GetLibraryHandle()
	strs,err:=shandle.Get()
	log.Debug("GetLibraryList: ", strs, err)

	i:=shandle.NewItem()
	strss,err:=i.GetByLibraryID(strs[0])
	log.Debug("GetItemByLibraryId: ", strss, err)
}

func cis_test() {
	code := cis.CodeBase64("root@vsphere.local", "Root@2021")
	log.Debug("base64:", code)

	client := common.NewClient("https://128.179.0.241/rest/")
	//client := common.NewClient("http://127.0.0.1/")
	log.Debug(client)

	c := cis.NewCIS(client)
	log.Debug(*c)

	shandle := c.GetSessionHandle()
	log.Debug(shandle)

	sess, err := shandle.CreateSession(code)
	log.Debug("CreateSession: ", sess, err)

	err = shandle.DeleteSession(sess)
	log.Debug("DeleteSession: ", err)
}

func main() {
	log.SetOutput(os.Stdout)     //设置日志的输出为标准输出
	log.SetLevel(log.DebugLevel) //设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
	log.SetReportCaller(true)    //设置日志的调用文件，调用函数

	// test cis module
	cis_test()

	//test content moudle
	content_test()
}
