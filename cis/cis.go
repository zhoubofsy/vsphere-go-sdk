package cis

import (
	b64 "encoding/base64"
	"vsphere-go-sdk/common"
)

type session struct {
	addr string
	uri  string
}

type CreateSessionResponse struct {
	value string `json:"value"`
}

func (o *session) CreateSession(basic string) (string, *common.Error) {
	header := make(map[string]string)
	header["Authorization"] = "Basic " + basic
	resp, err := common.NewClient().SendRequest(o.addr+o.uri, header, nil, "POST")
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

func (o *session) DeleteSession(basic string) *common.Error {
	header := make(map[string]string)
	header["Authorization"] = "Basic " + basic
	resp, err := common.NewClient().SendRequest(o.addr+o.uri, header, nil, "DELETE")
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

type cis struct {
	s session
}

func (o *cis) GetSessionHandle() *session {
	return o.s
}

func NewCIS(host string) *cis {
	return &cis{
		s: session{
			addr: host,
			uri:  "/rest/com/vmware/cis/session",
		},
	}
}

func CodeBase64(username string, password string) string {
	data := username + ":" + password
	return b64.StdEncoding.EncodeToString([]byte(data))
}
