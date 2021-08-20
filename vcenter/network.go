package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* Network Operations
 */
type Network struct {
	con *common.Connector
	uri string
}

type NetworkInfo struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Network string `json:"network,omitempty"`
}

type ValueOfNetworkInfo struct {
	Value []NetworkInfo `json:"value,omitempty"`
}

func (o *Network) List() ([]NetworkInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ListNetworks")
		return nil, err
	}

	networks := ValueOfNetworkInfo{}
	err = json.Unmarshal(resp.Data, &networks)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ListNetworks")
		return nil, err
	}
	return networks.Value, err
}
