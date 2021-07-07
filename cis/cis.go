package cis

import (
	b64 "encoding/base64"
	"encoding/json"
	"vsphere-go-sdk/common"
)

type Session struct {
	client common.Client
	uri    string
}

type CreateSessionResponse struct {
	value string `json:"value"`
}

func (o *Session) CreateSession(basic string) (string, *common.Error) {
	header := make(map[string]string)
	header["Authorization"] = "Basic " + basic
	resp, err := o.client.SendRequest(o.uri, header, nil, "POST")
	if err != nil {
		return "", common.ESENDREQUEST
	}
	if resp.Status != 200 {
		switch resp.Status {
		case 401:
			return "", common.EUNAUTHORED
		case 503:
			return "", common.ESERVICE_UNAVAILABLE
		}
		return "", common.EUNKNOW
	}
	response := CreateSessionResponse{}
	err = json.Unmarshal(resp.Data, &response)
	if err != nil {
		return "", common.EUNMARSHAL
	}
	return response.value, common.EOK
}

func (o *Session) DeleteSession(basic string) *common.Error {
	header := make(map[string]string)
	header["Authorization"] = "Basic " + basic
	resp, err := o.client.SendRequest(o.uri, header, nil, "DELETE")
	if err != nil {
		return common.ESENDREQUEST
	}
	if resp.Status != 200 {
		switch resp.Status {
		case 401:
			return common.EUNAUTHORED
		case 503:
			return common.ESERVICE_UNAVAILABLE
		}
		return common.EUNKNOW
	}
	return common.EOK
}

type CIS struct {
	s Session
}

func (o *CIS) GetSessionHandle() *Session {
	return &(o.s)
}

func NewCIS(c common.Client) *CIS {
	return &CIS{
		s: Session{
			client: c,
			uri:    "/rest/com/vmware/cis/session",
		},
	}
}

func CodeBase64(username string, password string) string {
	data := username + ":" + password
	return b64.StdEncoding.EncodeToString([]byte(data))
}
