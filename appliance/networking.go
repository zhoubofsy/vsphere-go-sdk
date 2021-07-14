package appliance

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

type Networking struct {
	con *common.Connector
	uri string
}

type NwDNS struct {
	Hostname string   `json:"hostname,omitempty"`
	Mode     string   `json:"mode,omitempty"`
	Servers  []string `json:"servers,omitempty"`
}
type NwIpv4 struct {
	Address        string `json:"address,omitempty"`
	Configurable   bool   `json:"configurable,omitempty"`
	DefaultGateway string `json:"default_gateway,omitempty"`
	Mode           string `json:"mode,omitempty"`
	Prefix         int64  `json:"prefix,omitempty"`
}
type Ipv6Addresses struct {
	Address string `json:"address,omitempty"`
	Origin  string `json:"origin,omitempty"`
	Prefix  int64  `json:"prefix,omitempty"`
	Status  string `json:"status,omitempty"`
}
type NwIpv6 struct {
	Addresses      []Ipv6Addresses `json:"addresses,omitempty"`
	Autoconf       bool            `json:"autoconf,omitempty"`
	Configurable   bool            `json:"configurable,omitempty"`
	DefaultGateway string          `json:"default_gateway,omitempty"`
	Dhcp           bool            `json:"dhcp,omitempty"`
}
type InfValue struct {
	Ipv4   NwIpv4 `json:"ipv4,omitempty"`
	Ipv6   NwIpv6 `json:"ipv6,omitempty"`
	Mac    string `json:"mac,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}
type NwInterfaces struct {
	Key   string   `json:"key,omitempty"`
	Value InfValue `json:"value,omitempty"`
}
type NwValue struct {
	DNS        NwDNS        `json:"dns,omitempty"`
	Interfaces NwInterfaces `json:"interfaces,omitempty"`
}

type GetNetworkingResult struct {
	Value NwValue `json:"value,omitempty"`
}

func (o *Networking) Get() (*NwValue, error) {
	header := make(map[string]string)
	header["vmware-api-session-id"] = o.con.Sid
	resp, err := o.con.Invoker.SendRequest(o.uri, header, nil, "GET")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("GetNetworking")
		return nil, err
	}
	Nw := GetNetworkingResult{}
	err = json.Unmarshal(resp.Data, &Nw)
	if err != nil {
		log.WithFields(log.Fields{"Response Data": string(resp.Data)}).Error("GetNetworking")
		return nil, err
	}
	return &Nw.Value, err
}
