package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	SendRequest(url string,headers map[string]string,body []byte,method string) (*ResponseResult,error)
}
func  GetClient() Client {
	return GetRESTClient(TIMEOUT)
}

type HttpClient struct {
	httpClient *http.Client
}
// @param timeoutSec  timeout in seconds
func GetRESTClient(timeoutSec int) *HttpClient {
	cli := &HttpClient{}
	cli.httpClient = &http.Client{
		Timeout: time.Duration(timeoutSec) * time.Second,
	}
	return cli
}

//Url中要把需要的参数param都拼接进去
//method需要写POST/GET/DELETE等参数
func (c *HttpClient)  SendRequest(url string,headers map[string]string,body []byte,method string) (*ResponseResult,error) {
	ioReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, ioReader)
	if err != nil {
		return nil, err
	}
	//给请求添加header
	for key,value:=range headers {
		req.Header.Add(key,value)
	}
	//发起请求
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data,err:=ioutil.ReadAll(res.Body)
	result := &ResponseResult{res.StatusCode, data}
	return result,err
}

type ResponseResult struct {
	Status int
	Data []byte
}