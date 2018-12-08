package main

import (
	"fmt"
)

type Epub struct {
	FileName    string      `json:"FileName"`
	Directory   string      `json:"Directory"`
	Title       interface{} `json:"Title"`
	Author      interface{} `json:"Creator"`
	Language    interface{} `json:"Language"`
	Publisher   interface{} `json:"Publisher"`
	Description interface{} `json:"Description"`
}

func (epub Epub) String() string {
	return fmt.Sprintf("%#v", epub)
}
