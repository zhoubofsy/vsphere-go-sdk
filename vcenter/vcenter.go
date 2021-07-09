package vcenter

import (
	"vsphere-go-sdk/common"
)

type VCenter struct {
	con *common.Connector
	uri string
}

func NewVCenter(client common.Client, sessId string) *VCenter {
	vc := &VCenter{
		uri: "vcenter",
	}
	vc.con = &common.Connector{
		Invoker: client,
		Sid:     sessId,
	}
	return vc
}

func (o *VCenter) NewVM() *VM {
	v := &VM{
		con: o.con,
		uri: o.uri + "/vm",
	}
	return v
}

func (o *VCenter) NewFolder() *Folder {
	f := &Folder{
		con: o.con,
		uri: o.uri + "/folder",
	}
	return f
}
