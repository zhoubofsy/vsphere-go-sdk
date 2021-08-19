package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* VM Template Operations
 */
type VMTemplate struct {
	con *common.Connector
	uri string
}

func (o *VMTemplate) NewLibraryItems() *LibraryItems {
	return &LibraryItems{
		con: o.con,
		uri: o.uri + "/library-items",
	}
}

type LibraryItems struct {
	con *common.Connector
	uri string
}

func (o *LibraryItems) NewItem(id string) *Item {
	return &Item{
		con: o.con,
		uri: o.uri + "/" + id + "?action=deploy",
	}
}

type Item struct {
	con *common.Connector
	uri string
}

type VMTemplateDeployHomeStorage struct {
	DataStore string `json:"datastore,omitempty"`
}

type VMTemplateHDCustomNIC struct {
	Key   string `json:"key"`
	Value struct {
		Network string `json:"network"`
	} `json:"value"`
}

type VMTemplateHDCustom struct {
	NICs []VMTemplateHDCustomNIC `json:"nics,omitempty"`
}

type VMTemplateDeployReqeust struct {
	Spec struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PoweredOn   bool   `json:"powered_on"`
		Placement   struct {
			ClusterID      string `json:"cluster,omitempty"`
			FolderID       string `json:"folder,omitempty"`
			Host           string `json:"host,omitempty"`
			ResourcePoolID string `json:"resource_pool,omitempty"`
		} `json:"placement"`
		VMHomeStorage  *VMTemplateDeployHomeStorage `json:"vm_home_storage,omitempty"`
		HardwareCustom *VMTemplateHDCustom          `json:"hardware_customization,omitempty"`
	} `json:"spec"`
}

type VMTemplateDeployResponse struct {
	Value string `json:"value"`
}

func (o *Item) Deploy(req *VMTemplateDeployReqeust) (string, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	header["Content-Type"] = "application/json"
	body, err := json.Marshal(*req)
	if err != nil {
		return "", err
	}
	resp, err := o.con.Invoker.SendRequest(o.uri, header, body, "POST")
	if err != nil {
		log.Error("Deploy SendRequest Error, ", err)
		return "", err
	}

	rps := VMTemplateDeployResponse{}
	err = json.Unmarshal(resp.Data, &rps)
	if err != nil {
		log.Error("Deploy Unmarshal Error, ", err)
		return "", err
	}
	return rps.Value, err
}
