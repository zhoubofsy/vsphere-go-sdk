package common

import (
	"bytes"
	"crypto/tls"
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
		log.WithFields(log.Fields{"err": err}).Error("http.NewRequest")
		return nil, err
	}
	//add header
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	log.WithFields(log.Fields{"Request": req}).Trace("http.NewRequest")
	//send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("http.Do")
		return nil, err
	}
	log.WithFields(log.Fields{"Response": res}).Trace("http.Do")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	result := &ResponseResult{res.StatusCode, data}
	return result, err
}

type ResponseResult struct {
	Status int
	Data   []byte
}
