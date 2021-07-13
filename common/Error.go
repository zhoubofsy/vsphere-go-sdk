package common

import (
	"encoding/json"
	"fmt"
)

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
	EMARSHAL             = &Error{Code: 1003, Msg: "json marshal error"}
)


type ResponseErrorMessages struct {
	Args           []string `json:"args,omitempty"`
	DefaultMessage string   `json:"default_message,omitempty"`
	Id             string   `json:"id,omitempty"`
}

type ResponseErrorValue struct {
	Messages []ResponseErrorMessages `json:"messages,omitempty"`
	Data     map[string]string       `json:"data,omitempty"`
}

type ResponseError struct {
	Type  string             `json:"type,omitempty"`
	Value ResponseErrorValue `json:"value,omitempty"`
}

type VsphereSDKError struct {
	Code    int
	Message *ResponseError
}

func (e *VsphereSDKError) Error() string {
	return fmt.Sprintf("[VsphereSDKError] Code=%d, Message=%s", e.Code, e.Message)
}

func (e *VsphereSDKError) GetCode() int {
	return e.Code
}

func (e *VsphereSDKError) GetMessage() *ResponseError {
	return e.Message
}

func NewVsphereSDKError(code int, message *ResponseError) error {
	return &VsphereSDKError{
		Code:    code,
		Message: message,
	}
}

func ParseErrorFromResponse(result *ResponseResult) (err error) {
	resp := &ResponseError{}
	err = json.Unmarshal(result.Data, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", result.Data, err)
		return NewVsphereSDKError(result.Status, &ResponseError{
			Value: ResponseErrorValue{
				Messages: []ResponseErrorMessages{
					{DefaultMessage: msg},
				},
			},
		})
	}
	return NewVsphereSDKError(result.Status, resp)
}