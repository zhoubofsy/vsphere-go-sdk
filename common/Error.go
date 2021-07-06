package common

// Error basic error info for IAAS
type Error struct {
	Code int32
	Msg  string
}

func (e *Error) Error() string {
	return e.Msg
}

func SimpleError(msg string) *Error {
	return &Error{Code: 0, Msg: msg}
}

// general error code
var (
	EOK                  = &Error{Code: 200, Msg: "200 OK"}
	EUNAUTHORED          = &Error{Code: 401, Msg: "401 Unauthorized"}
	EJSONFORMAT          = &Error{Code: 500, Msg: "Interval Json Format Error"}
	EUNSUPPORTEDPROVIDER = &Error{Code: 599, Msg: "599 unsupported cloud provider"}
	ESERVICE_UNAVAILABLE = &Error{Code: 503, Msg: "com.vmware.vapi.std.errors.service_unavailable"}
	EUNKNOW              = &Error{Code: 1000, Msg: "unknow error"}
	ESENDREQUEST         = &Error{Code: 1001, Msg: "send request error"}
	EUNMARSHAL           = &Error{Code: 1002, Msg: "json unmarshal error"}
)
