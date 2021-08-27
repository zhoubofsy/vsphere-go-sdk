package vcenter

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

/*
* OVF Operations
 */
type OVF struct {
	con *common.Connector
	uri string
}

func (o *OVF) NewOVFLibraryItem() *OVFLibraryItem {
	r := &OVFLibraryItem{
		con: o.con,
		uri: o.uri + "/library-item",
	}
	return r
}

type OVFLibraryItem struct {
	con *common.Connector
	uri string
}

func (o *OVFLibraryItem) NewOVFItem(id string) *OVFItem {
	if id == "" {
		return nil
	}
	r := &OVFItem{
		con: o.con,
		uri: o.uri + "/id:" + id,
	}
	return r
}

type OVFItem struct {
	con *common.Connector
	uri string
}

type OVFItemInfo struct {
	Annotation    string   `json:"annotation"`
	EULAs         []string `json:"EULAs"`
	Name          string   `json:"name"`
	StorageGroups []string `json:"storage_groups"`
	Networks      []string `json:"networks"`
}

type ValueOfOVFItemInfo struct {
	Value OVFItemInfo `json:"value,omitempty"`
}

type OVFActionRequest struct {
	Target           OVFReqeustTarget `json:"target"`
	OVFLibraryItemID string           `json:"ovf_library_item_id"`
}

func (o *OVFItem) Get(req *OVFActionRequest) (*OVFItemInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	header["Content-Type"] = "application/json"
	body, err := json.Marshal(*req)
	if err != nil {
		log.Error("OVF Get Marshal error, ", err, " req: ", *req)
		return nil, err
	}
	uri := o.uri + "?~action=filter"
	resp, err := o.con.Invoker.SendRequest(uri, header, body, "POST")
	if err != nil {
		log.Error("OVF LibraryItem Get error, ", err, " uri: ", uri)
		return nil, err
	}

	ovfitem := ValueOfOVFItemInfo{}
	err = json.Unmarshal(resp.Data, &ovfitem)
	if err != nil {
		log.Error("OVF LibraryItem Unmarshal error, ", err, " data: ", string(resp.Data))
		return nil, err
	}
	return &(ovfitem.Value), err
}

type OVFDeployResultResourceID struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type OVFDeployResultErrorMessage struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

type OVFDeployResultErrorIssue struct {
	LineNumber   int                         `json:"line_number"`
	File         string                      `json:"file"`
	Category     string                      `json:"category"`
	ColumnNumber int                         `json:"column_number"`
	Message      OVFDeployResultErrorMessage `json:"message"`
}

type OVFDeployResultErrorWAR struct {
	Message  OVFDeployResultErrorMessage `json:"message"`
	Category string                      `json:"category"`
	Issues   []OVFDeployResultErrorIssue `json:"issues"`
	Name     string                      `json:"name"`
	Value    string                      `json:"value"`
}

type OVFDeployResultErrorINF struct {
	Messages []OVFDeployResultErrorMessage `json:"messages"`
}

type OVFDeployResultErrorERR struct {
	Message  OVFDeployResultErrorMessage `json:"message"`
	Category string                      `json:"category"`
	Issues   []OVFDeployResultErrorIssue `json:"issues"`
	Name     string                      `json:"name"`
	Value    string                      `json:"value"`
}

type OVFDeployResultError struct {
	Warnings    []OVFDeployResultErrorWAR `json:"warnings,omitempty"`
	Information []OVFDeployResultErrorINF `json:"information,omitempty"`
	Errors      []OVFDeployResultErrorERR `json:"errors,omitempty"`
}

type OVFDeployResult struct {
	ResourceID OVFDeployResultResourceID `json:"resource_id"`
	Error      OVFDeployResultError      `json:"error"`
	Successded string                    `json:"successded"`
}

type ValueOfOVFDeployResult struct {
	Value OVFDeployResult `json:"value"`
}

type OVFDeployRequestSpecNetworkMapping struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OVFReqeustTarget struct {
	ResourcePoolID string `json:"resource_pool_id"`
	HostID         string `json:"host_id,omitempty"`
	FolderID       string `json:"folder_id"`
}

type OVFDeployRequest struct {
	DeploymentSpec struct {
		Name            string                               `json:"name"`
		AcceptAllEULA   bool                                 `json:"accept_all_EULA"`
		DefaultDSID     string                               `json:"default_datastore_id,omitempty"`
		NetworkMappints []OVFDeployRequestSpecNetworkMapping `json:"network_mappings,omitempty"`
	} `json:"deployment_spec"`
	Target           OVFReqeustTarget `json:"target"`
	OVFLibraryItemID string           `json:"ovf_library_item_id"`
}

func (o *OVFItem) Deploy(req *OVFDeployRequest) (*OVFDeployResult, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	header["Content-Type"] = "application/json"
	body, err := json.Marshal(*req)
	if err != nil {
		log.Error("OVF Deploy Marshal error, ", err, " req: ", *req)
		return nil, err
	}
	uri := o.uri + "?~action=deploy"
	resp, err := o.con.Invoker.SendRequest(uri, header, body, "POST")
	if err != nil {
		log.Error("OVF Deploy SendRequest Error, ", err)
		return nil, err
	}

	rps := ValueOfOVFDeployResult{}
	err = json.Unmarshal(resp.Data, &rps)
	if err != nil {
		log.Error("OVF Deploy Unmarshal Error, ", err, " data: ", string(resp.Data))
		return nil, err
	}
	return &(rps.Value), err
}
