package vcenter

import (
	"vsphere-go-sdk/common"
)

/*
* VM Operations
 */
type VM struct {
	client common.Client
	uri    string
}

type ListVMsResult struct {
}

func (o *VM) ListVMs() []ListVMsResult {
}

type VMInfo struct {
}

func (o *VM) GetVMInfo(vm string) *VMInfo {
}

func (o *VM) NewHardware(vm string) *Hardware {
	return &Hardware{
		client: o.client,
		uri:    o.uri + vm + "/hardware",
	}
}

/*
* Hardware Operations
 */
type Hardware struct {
	client common.Client
	uri    string
}

func (o *Hardware) NewDisk() *Disk {
	return &Disk{
		client: o.client,
		uri:    o.uri + "/disk",
	}
}

/*
* Disk Operations
 */
type Disk struct {
	client common.Client
	uri    string
}
