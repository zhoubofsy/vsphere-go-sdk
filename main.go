package main

import (
	"log"
	"vsphere-go-sdk/cis"
	"vsphere-go-sdk/common"
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
	log.Println(resBody)

}

func cis_test() {
	client := common.NewClient("https://localhost:8088/rest/")
	c := cis.NewCIS(client)

	sess, _ := c.GetSessionHandle().CreateSession(cis.CodeBase64("username", "password"))

	log.Println(sess)
}

func main() {
	cis_test()
}
