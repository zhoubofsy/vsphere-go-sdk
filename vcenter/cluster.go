package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* Cluster Operations
 */
type Cluster struct {
	con *common.Connector
	uri string
}

type ClusterInfo struct {
	ClusterID  string `json:"cluster"`
	DRSEnabled bool   `json:"drs_enabled"`
	Name       string `json:"name"`
	HAEnabled  bool   `json:"ha_enabled"`
}

type ValueOfClusterInfo struct {
	Value []ClusterInfo `json:"value,omitempty"`
}

func (o *Cluster) List() ([]ClusterInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ListClusters")
		return nil, err
	}

	clusters := ValueOfClusterInfo{}
	err = json.Unmarshal(resp.Data, &clusters)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ListClusters")
		return nil, err
	}
	return clusters.Value, err
}
