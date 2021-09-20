// Package models defines different data structures and interfaces
// that will be used in more than 1 package
package models

// ValidateResp ...
//
// Structure that defines the fields we will be returning in our /validate handler
type ValidateResp struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

// ErrResp ...
//
// Structure that defines the fields we will be returning in our /errors handler
type ErrResp struct {
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	ErrType    string `json:"type"`
	Expression string `json:"expression"`
}
