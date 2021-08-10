package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* DataStore Operations
 */
type DataStore struct {
	con *common.Connector
	uri string
}

type DSInfo struct {
	DSID      string `json:"datastore"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Capacity  int    `json:"capacity"`
	FreeSpace int    `json:"free_space"`
}

type ValueOfDSInfo struct {
	Value []DSInfo `json:"value,omitempty"`
}

func (o *DataStore) List() ([]DSInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.Error("ListDSs Error: ", err)
		return nil, err
	}

	dss := ValueOfDSInfo{}
	err = json.Unmarshal(resp.Data, &dss)
	if err != nil {
		log.Error("ListDSs Response Data", string(resp.Data))
		return nil, err
	}
	return dss.Value, err
}
