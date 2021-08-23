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
		uri: o.uri + "/" + id,
	}
}

type Item struct {
	con *common.Connector
	uri string
}

type VMTemplateHomeStorage struct {
	DataStore     string `json:"datastore,omitempty"`
	StoragePolicy string `json:"storage_policy,omitempty"`
}

type VMTemplateDisk struct {
	Key   string `json:"key"`
	Value struct {
		DiskStorage struct {
			StoragePolicy string `json:"storage_policy,omitempty"`
			DataStore     string `json:"datastore,omitempty"`
		} `json:"disk_storage"`
		Capacity int `json:"capacity"`
	} `json:"value"`
}

type VMTemplateNIC struct {
	Key   string `json:"key"`
	Value struct {
		BackingType string `json:"backing_type"`
		Network     string `json:"network"`
		MacType     string `json:"mac_type"`
	} `json:"Value"`
}

type VMTemplateMemory struct {
	SizeMiB int `json:"size_MiB"`
}

type VMTemplateCPU struct {
	Count          int `json:"count"`
	CoresPerSocket int `json:"cores_per_socket"`
}

type VMTemplateInfo struct {
	VMTemplateName string                 `json:"vm_template"`
	Disks          []VMTemplateDisk       `json:"disks,omitempty"`
	VMHomeStorage  *VMTemplateHomeStorage `json:"vm_home_storage,omitempty"`
	NICs           []VMTemplateNIC        `json:"nics,omitempty"`
	Memory         VMTemplateMemory       `json:"memory"`
	GuestOS        string                 `json:"guest_OS"`
	CPU            VMTemplateCPU          `json:"cpu"`
}

type ValueOfVMTemplateInfo struct {
	Value VMTemplateInfo `json:"value"`
}

func (o *Item) Get() (*VMTemplateInfo, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	header["Content-Type"] = "application/json"

	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.Error("Item Get SendRequest Error, ", err)
		return nil, err
	}

	vtmp := ValueOfVMTemplateInfo{}
	vtmp.Value.VMHomeStorage = &VMTemplateHomeStorage{}
	err = json.Unmarshal(resp.Data, &vtmp)
	if err != nil {
		log.Error("Item Get Unmarshal Error, ", err)
		return nil, err
	}
	return &(vtmp.Value), err
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
	uri := o.uri + "?action=deploy"
	resp, err := o.con.Invoker.SendRequest(uri, header, body, "POST")
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
