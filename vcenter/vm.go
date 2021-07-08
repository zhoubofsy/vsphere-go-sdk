package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"vsphere-go-sdk/common"
)

/*
* VM Operations
 */
type VM struct {
	con *common.Connector
	uri string
}

type VMResult struct {
	Vm         string `json:"vm"`
	Name       string `json:"name"`
	PowerState string `json:"power_state"`
	CpuCount   int    `json:"cpu_count"`
	MemSizeMiB int    `json:"memory_size_MiB"`
}

type ListVMsResult struct {
	Value []VMResult `json:"value"`
}

func (o *VM) ListVMs() ([]VMResult, *common.Error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ListVMs")
		return nil, common.ESENDREQUEST
	}
	if resp.Status != 200 {
		log.WithFields(log.Fields{"Response": resp}).Error("ListVMs")
		switch resp.Status {
		case 401:
			return nil, common.EUNAUTHORED
		}
		return nil, common.EUNKNOW
	}
	vms := ListVMsResult{}
	err = json.Unmarshal(resp.Data, &vms)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ListVMs")
		return nil, common.EUNMARSHAL
	}
	return vms.Value, common.EOK
}

type VMInfo struct {
}

func (o *VM) GetVMInfo(vm string) *VMInfo {
	return nil
}

func (o *VM) NewHardware(vm string) *Hardware {
	return &Hardware{
		con: o.con,
		uri: o.uri + vm + "/hardware",
	}
}

/*
* Hardware Operations
 */
type Hardware struct {
	con *common.Connector
	uri string
}

func (o *Hardware) NewDisk() *Disk {
	return &Disk{
		con: o.con,
		uri: o.uri + "/disk",
	}
}

/*
* Disk Operations
 */
type Disk struct {
	con *common.Connector
	uri string
}
