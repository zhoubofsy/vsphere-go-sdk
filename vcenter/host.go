package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* Host Operations
 */
type Host struct {
	con *common.Connector
	uri string
}

type HostInfo struct {
	HostID          string `json:"host"`
	Name            string `json:"name"`
	ConnectionState string `json:"connection_state"`
	PowerState      string `json:"power_state"`
}

type ValueOfHostInfo struct {
	Value []HostInfo `json:"value,omitempty"`
}

func (o *Host) List() ([]HostInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.Error("ListHosts Error: ", err)
		return nil, err
	}

	hosts := ValueOfHostInfo{}
	err = json.Unmarshal(resp.Data, &hosts)
	if err != nil {
		log.Error("ListHosts Response Data", string(resp.Data))
		return nil, err
	}
	return hosts.Value, err
}
