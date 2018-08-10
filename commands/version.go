package commands

import (
	"fmt"
)

var VERSION = "0.0.0-dev"

type Version struct {
}

//Execute - returns the version
func (c *Version) Execute([]string) error {
	fmt.Println(VERSION)
	return nil
}
