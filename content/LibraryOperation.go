package content

import "fmt"

type LibraryOperation interface {
	Get() (*Library,error)
}

type Library struct {
	value string
}

func (library *Library) Get() (*Library,error) {
	return nil,nil
}

func Diaoyong() (*Library, error) {

	pp := &Library{
		value: "test",
	}
	var lib LibraryOperation = pp
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

