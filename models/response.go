// response ...
package models

type ValidateResp struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

type ErrResp struct {
	Endpoint   string
	Frequency  int
	ErrType    error
	Expression string
}
