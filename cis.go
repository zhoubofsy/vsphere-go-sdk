package vsphere

import (
	b64 "encoding/base64"
)

type session struct {
	baseURI string
}

func (o *session) CreateSession(basic string) (string, error) {
	return "Fucking High!", nil
}

func (o *session) DeleteSession(basic string) error {
	return nil
}

type cis struct {
	s session
}

func (o *cis) GetSessionHandle() *session {
	return o.s
}

func GetCIS() *cis {
	return &cis{
		s: session{baseURI: "/rest/com/vmware/cis/session"},
	}
}

func CodeBase64(username string, password string) string {
	data := username + ":" + password
	return b64.StdEncoding.EncodeToString([]byte(data))
}
