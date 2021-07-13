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
	} `json:"spec"`
}

type VMTemplateDeployResponse struct {
	Value string `json:"value"`
}

func (o *Item) Deploy(req *VMTemplateDeployReqeust) (string, *common.Error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	header["Content-Type"] = "application/json"
	body, err := json.Marshal(*req)
	if err != nil {
		return "", common.EMARSHAL
	}
	resp, err := o.con.Invoker.SendRequest(o.uri, header, body, "POST")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ItemDeploy")
		return "", common.ESENDREQUEST
	}
	if resp.Status != 200 {
		log.WithFields(log.Fields{"Response Code": resp.Status, "Response Data": string(resp.Data)}).Error("ItemDeploy")
		switch resp.Status {
		case 401:
			return "", common.EUNAUTHORED
		}
		return "", common.EUNKNOW
	}
	rps := VMTemplateDeployResponse{}
	err = json.Unmarshal(resp.Data, &rps)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ItemDeploy")
		return "", common.EUNMARSHAL
	}
	return rps.Value, common.EOK
}
