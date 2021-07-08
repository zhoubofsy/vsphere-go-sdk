package content

import (
	"vsphere-go-sdk/common"
)
type Library struct {
	conn Connection
	uri    string
}
type LibraryResult struct {

}
func (l *Library) Get() ([]LibraryResult,*common.Error) {
	return nil,common.EOK
}

func (l *Library) NewItem() *Item {
	return &Item{
		conn: l.conn,
		uri:    l.uri  + "/item",
	}
}

type Item struct {
	conn Connection
	uri    string
}
func (i *Item) GetByLibraryID(libraryId string) (*Library,*common.Error) {
	return nil,common.EOK
}




