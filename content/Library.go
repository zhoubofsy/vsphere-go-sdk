package content

import (
	"fmt"
	"vsphere-go-sdk/common"
)
type Library struct {
	client common.Client
	uri    string
}
func (library *Library) Get() (*Library,*common.Error) {
	return nil,common.EOK
}


type Content struct {
	library Library
}

func (o *Content) GetSessionHandle() *Library {
	return &(o.library)
}

func NewContent(c common.Client) *Content {
	return &Content{
		library: Library{
			client: c,
			uri: "com/vmware/cis/session",
		},
	}
}



func Diaoyong() (*Library, error) {

	pp := &Library{
		//value: "test",
	}
	var lib Operation = pp
	fmt.Println("Before", pp)
	methodMap := map[string]interface{}{
		"Get": lib.Get(),
	}
	for k, v := range methodMap {
		switch k {
		case "Get":
			return v.(func(string) (*Library, error))("vcsdsh")
		case "Create":
			//v.(func() string)()
		}
	}
	fmt.Println("After", pp)
	return nil,nil
}

