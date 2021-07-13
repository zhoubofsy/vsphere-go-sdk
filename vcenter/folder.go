package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* Folder Operations
 */
type Folder struct {
	con *common.Connector
	uri string
}

type FolderInfo struct {
	FolderID string `json:"folder"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

type ValueOfFolderInfo struct {
	Value []FolderInfo `json:"value,omitempty"`
}

func (o *Folder) List() ([]FolderInfo, *common.Error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("ListFolders")
		return nil, common.ESENDREQUEST
	}
	if resp.Status != 200 {
		log.WithFields(log.Fields{"Response": resp}).Error("ListFolders")
		switch resp.Status {
		case 401:
			return nil, common.EUNAUTHORED
		}
		return nil, common.EUNKNOW
	}
	folders := ValueOfFolderInfo{}
	err = json.Unmarshal(resp.Data, &folders)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("ListFolders")
		return nil, common.EUNMARSHAL
	}
	return folders.Value, common.EOK
}
