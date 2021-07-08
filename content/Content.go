package content

import "vsphere-go-sdk/common"

type Connection struct {
	client *common.Client
	sid string
}

type Content struct {
	library Library
}

func (o *Content) GetSessionHandle() *Library {
	return &(o.library)
}

func NewContent(c common.Client,sid string) *Content {
	conn:= &Connection{client: c, sid: sid}
	return &Content{
		Library{
			uri:  "com/vmware/cis/session",
			conn: conn,
		},
	}
}
