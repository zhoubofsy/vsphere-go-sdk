package cis

import (
	b64 "encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

type Session struct {
	client common.Client
	uri    string
}

func (o *Session) CreateSession(basic string) (string, error) {
	header := make(map[string]string)
	header["Authorization"] = "Basic " + basic
	resp, err := o.client.SendRequest(o.uri, header, nil, "POST")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("CreateSession")
		return "", err
	}
	response := make(map[string]string)
	err = json.Unmarshal(resp.Data, &response)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("CreateSession")
		return "", err
	}
	return response["value"], err
}

type SessionInfo struct {
	CreateTime       string `json:"created_time"`
	LastAccessedTime string `json:"last_accessed_time"`
	User             string `json:"user"`
}

type ValueOfSessionInfo struct {
	Value SessionInfo `json:"value"`
}

func (o *Session) Update(sessid string) (*SessionInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = sessid
	resp, err := o.client.SendRequest(o.uri, header, nil, "POST")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("UpdateSession")
		return nil, err
	}

	sessInfo := &ValueOfSessionInfo{}
	err = json.Unmarshal(resp.Data, sessInfo)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("UpdateSession")
		return nil, err
	}
	return &(sessInfo.Value), err
}

func (o *Session) DeleteSession(sessid string) error {
	header := make(map[string]string)
	header["vmware-api-session-id"] = sessid
	_, err := o.client.SendRequest(o.uri, header, nil, "DELETE")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("DeleteSession")
	}
	return err
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
			uri:    "com/vmware/cis/session",
		},
	}
}

func CodeBase64(username string, password string) string {
	data := username + ":" + password
	return b64.StdEncoding.EncodeToString([]byte(data))
}
