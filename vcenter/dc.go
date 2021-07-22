package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* DataCenter Operations
 */
type DC struct {
	con *common.Connector
	uri string
}

type DCInfo struct {
	DCID string `json:"datacenter"`
	Name string `json:"name"`
}

type ValueOfDCInfo struct {
	Value []DCInfo `json:"value,omitempty"`
}

func (o *DC) List() ([]DCInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ListDCs")
		return nil, err
	}

	dcs := ValueOfDCInfo{}
	err = json.Unmarshal(resp.Data, &dcs)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ListDCs")
		return nil, err
	}
	return dcs.Value, err
}
