// response ...
package models

type ValidateResp struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

type ErrResp struct {
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	ErrType    string `json:"type"`
	Expression string `json:"expression"`
}
