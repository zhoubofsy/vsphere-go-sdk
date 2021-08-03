package content

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

type Library struct {
	conn *common.Connector
	uri  string
}

func (l *Library) ListLibraries() ([]string, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = l.conn.Sid
	resp, err := l.conn.Invoker.SendRequest(l.uri, header, nil, "GET")
	if err != nil {
		log.Error("SendRequestError: ", err)
		return nil, err
	}
	log.WithFields(log.Fields{"ResponseData": string(resp.Data)}).Debug("GetLibrary")
	response := make(map[string][]string)
	err = json.Unmarshal(resp.Data, &response)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"response: ": response}).Debug("GetLibraryResponse")
	return response["value"], err
}

func (l *Library) NewItem() *Item {
	return &Item{
		conn: l.conn,
		uri:  l.uri + "/item",
	}
}

type Item struct {
	conn *common.Connector
	uri  string
}

func (i *Item) GetItemByLibraryID(libraryId string) ([]string, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = i.conn.Sid
	resp, err := i.conn.Invoker.SendRequest(i.uri+"?library_id="+libraryId, header, nil, "GET")
	if err != nil {
		log.Error("SendRequestError: ", err)
		return nil, err
	}
	response := make(map[string][]string)
	err = json.Unmarshal(resp.Data, &response)
	if err != nil {
		return nil, err
	}
	return response["value"], err
}

type ValueOfItemInfo struct {
	Value ItemInfo `json:"value"`
}

type ItemInfo struct {
	LastModifiedTime string `json:"last_modified_time"`
	Cached           bool   `json:"cached"`
	Type             string `json:"type"`
	ID               string `json:"id"`
	Description      string `json:"description"`
	MetadataVersion  string `json:"metadata_version"`
	LibraryID        string `json:"library_id"`
	Version          string `json:"version"`
	Name             string `json:"name"`
	SourceID         string `json:"source_id"`
	CreationTime     string `json:"creation_time"`
	LastSyncTime     string `json:"last_sync_time"`
	ContentVersion   string `json:"content_version"`
	Size             int    `json:"size"`
}

func (i *Item) GetItemInfoByItemID(id string) (*ItemInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = i.conn.Sid
	uri := i.uri + "/id:" + id
	resp, err := i.conn.Invoker.SendRequest(uri, header, nil, "GET")
	if err != nil {
		log.Error("GetItemInfoByItemID Error: ", err)
		return nil, err
	}

	vii := ValueOfItemInfo{}
	err = json.Unmarshal(resp.Data, &vii)
	if err != nil {
		log.Error("GetItemInfoByItemID Response Data", string(resp.Data))
		return nil, err
	}
	return &(vii.Value), err
}
