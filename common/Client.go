package common

import (
	"bytes"
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Connector struct {
	Invoker Client
	Sid     string
}

type Client interface {
	SendRequest(url string, headers map[string]string, body []byte, method string) (*ResponseResult, error)
}

//@param host format eg:"https://127.0.0.1/rest/"
func NewClient(host string) Client {
	return NewRESTClient(host, TIMEOUT)
}
func NewClientWithTimeout(host string, timeout int) Client {
	return NewRESTClient(host, timeout)
}

type HttpClient struct {
	host       string
	httpClient *http.Client
}

// @param timeoutSec  timeout in seconds
func NewRESTClient(host string, timeoutSec int) *HttpClient {
	//avoid to have secure verify
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &HttpClient{}
	cli.host = host
	cli.httpClient = &http.Client{
		Timeout:   time.Duration(timeoutSec) * time.Second,
		Transport: tr,
	}
	return cli
}

//@param uri format eg:"com/vmware/cis/session"
//@param method supported eg:"GET" "POST"  "DELETE" "PATCH"
//@param headers
//@param body
func (c *HttpClient) SendRequest(uri string, headers map[string]string, body []byte, method string) (*ResponseResult, error) {
	ioReader := bytes.NewReader(body)
	url := c.host + uri
	req, err := http.NewRequest(method, url, ioReader)
	if err != nil {
		log.Error("SendRequest Error: ", err)
		return nil, err
	}
	//add header
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	//log.Debug("SendRequest Request: ", req)
	//send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Error("SendRequest Error: ", err)
		return nil, err
	}
	//log.Debug("SendRequest Response:", res)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		err = NewVsphereSDKError(res.StatusCode, &ResponseError{
			Value: ResponseErrorValue{
				Messages: []ResponseErrorMessages{
					{DefaultMessage: msg},
				},
			},
		})
	}
	result := &ResponseResult{res.StatusCode, data}
	if result.Status != 200 {
		err = ParseErrorFromResponse(result)
	}

	return result, err
}

type ResponseResult struct {
	Status int
	Data   []byte
}
