package content

import "vsphere-go-sdk/common"

type Content struct {
	library Library
}

func (o *Content) GetLibraryHandle() *Library {
	return &(o.library)
}

func NewContent(c common.Client, sid string) *Content {
	conn := common.Connector{Client: c, Sid: sid}
	return &Content{
		Library{
			uri:  "com/vmware/content/library",
			conn: conn,
		},
	}
}
