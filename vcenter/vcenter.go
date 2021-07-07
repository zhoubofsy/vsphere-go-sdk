package vcenter

import (
	"vsphere-go-sdk/common"
)

type VCenter struct {
	uri string
}

func NewVCenter(client common.Client) *VCenter {
	vc := &VCenter{
		uri: "vcenter",
	}
	return vc
}

func (o *VCenter) NewVM() *VM {
	v := &VM{
		client: o.client,
		uri:    o.uri + "/vm",
	}
	return v
}
