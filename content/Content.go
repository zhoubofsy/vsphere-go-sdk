package content

import (
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

type Content struct {
	conn *common.Connector
	uri  string
}

func (o *Content) NewLibrary() *Library {
	v := &Library{
		conn: o.conn,
		uri:  o.uri + "/library",
	}
	return v
}

func NewContent(c common.Client, sid string) *Content {
	vc := &Content{
		uri: "com/vmware/content",
	}
	vc.conn = &common.Connector{
		Invoker: c,
		Sid:     sid,
	}
	return vc
}
