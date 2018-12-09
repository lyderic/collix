package main

import (
	"fmt"
)

type Epub struct {
	FileName    string `json:"FileName"`
	Directory   string `json:"Directory"`
	Title       string `json:"Title"`
	Author      string `json:"Creator"`
	Language    string `json:"Language"`
	Publisher   string `json:"Publisher"`
	Description string `json:"Description"`
}

func (epub Epub) String() string {
	return fmt.Sprintf("%#v", epub)
}
