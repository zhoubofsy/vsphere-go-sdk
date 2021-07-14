package appliance

import (
	"liyongcool.nat300.top/iaas/vsphere-go-sdk/common"
)

type Appliance struct {
	con *common.Connector
	uri string
}

func NewAppliance(client common.Client, sessId string) *Appliance {
	app := &Appliance{
		uri: "appliance",
	}
	app.con = &common.Connector{
		Invoker: client,
		Sid:     sessId,
	}
	return app
}

func (o *Appliance) NewNetworking() *Networking {
	nw := &Networking{
		con: o.con,
		uri: o.uri + "/networking",
	}
	return nw
}
