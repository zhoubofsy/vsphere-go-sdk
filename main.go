package main

import (
	"encoding/json"
	"log"
)

func main() {
	vsphereClient := GetVsphereClient(10)

	url := "http://127.0.0.1:8088/v1/user"

	headers:=make(map[string]string)
	headers["session"]="test"
	method:="GET"
	testBody:="gdhsdsdsghd"
	body,_:=json.Marshal(testBody)
	resBody,_:=vsphereClient.sendRequest(url,headers,body,method)
	log.Println(resBody)

}

