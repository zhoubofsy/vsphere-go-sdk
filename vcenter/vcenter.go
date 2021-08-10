package vcenter

import (
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
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

func (o *VCenter) NewCluster() *Cluster {
	c := &Cluster{
		con: o.con,
		uri: o.uri + "/cluster",
	}
	return c
}

func (o *VCenter) NewDC() *DC {
	c := &DC{
		con: o.con,
		uri: o.uri + "/datacenter",
	}
	return c
}

func (o *VCenter) NewResourcePool() *ResourcePool {
	r := &ResourcePool{
		con: o.con,
		uri: o.uri + "/resource-pool",
	}
	return r
}

func (o *VCenter) NewVMTemplate() *VMTemplate {
	r := &VMTemplate{
		con: o.con,
		uri: o.uri + "/vm-template",
	}
	return r
}

func (o *VCenter) NewDataStore() *DataStore {
	r := &DataStore{
		con: o.con,
		uri: o.uri + "/datastore",
	}
	return r
}

func (o *VCenter) NewHost() *Host {
	r := &Host{
		con: o.con,
		uri: o.uri + "/host",
	}
	return r
}
