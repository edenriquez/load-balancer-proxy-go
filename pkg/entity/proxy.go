package entity

import "errors"

// Proxy is the data structure that will handle information
type Proxy struct {
	ID       int    `json:"id,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Weigth   int    `json:"weight,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")
