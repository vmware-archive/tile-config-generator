package commands

import (
	"fmt"
)

var VERSION = "shitty-beta"

type Version struct {
}

//Execute - returns the version
func (c *Version) Execute([]string) error {
	fmt.Println(VERSION)
	return nil
}
