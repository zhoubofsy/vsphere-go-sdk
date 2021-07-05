package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	xurl "net/url"
	"time"
)

type VsphereClient struct {
	httpClient *http.Client
}

// @param timeoutSec  timeout in seconds
func GetVsphereClient(timeoutSec int) *VsphereClient {
	cli := &VsphereClient{}
	cli.httpClient = &http.Client{
		Timeout: time.Duration(timeoutSec) * time.Second,
	}
	return cli
}

//Url中要把需要的参数param都拼接进去
//method需要写POST/GET/DELETE等参数
func (c *VsphereClient)  sendRequest(url string,headers map[string]string,body []byte,method string) ([]byte,error) {
	ioReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, ioReader)
	if err != nil {
		return []byte{}, err
	}

	//给请求添加header
	for key,value:=range headers {
		req.Header.Add(key,value)
	}
	//发起请求
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func getHostFromUrl(url string) string {
	u, _ := xurl.Parse(url)
	host, _, _ := net.SplitHostPort(u.Host)
	return host
}
