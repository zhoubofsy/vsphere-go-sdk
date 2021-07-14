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
	log.WithFields(log.Fields{"ResponseData": string(resp.Data)}).Debug("GetByLibraryID")
	response := make(map[string][]string)
	err = json.Unmarshal(resp.Data, &response)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"response: ": response}).Debug("GetByLibraryIDResponse")
	return response["value"], err
}
