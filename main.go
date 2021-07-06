package main

import (
	"log"
	"vsphere-sdk-go/common"
)

func main() {
	client := common.GetClient()

	url := "http://localhost:8088/v1/user/login"

	//headers:=make(map[string]string)
	//headers["session"]="test"
	method:="GET"
	//testBody:="gdhsdsdsghd"
	//body,_:=json.Marshal(testBody)
	resBody,_:=client.SendRequest(url,nil,nil,method)
	log.Println(resBody)

}

