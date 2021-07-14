package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* Resource Pool Operations
 */
type ResourcePool struct {
	con *common.Connector
	uri string
}

type ResourcePoolInfo struct {
	ResourcePoolID string `json:"resource_pool"`
	Name           string `json:"name"`
}

type ValueOfResourcePoolInfo struct {
	Value []ResourcePoolInfo `json:"value,omitempty"`
}

func (o *ResourcePool) List() ([]ResourcePoolInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ListResourcePools")
		return nil, err
	}

	rps := ValueOfResourcePoolInfo{}
	err = json.Unmarshal(resp.Data, &rps)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ListResourcePools")
		return nil, err
	}
	return rps.Value, err
}
