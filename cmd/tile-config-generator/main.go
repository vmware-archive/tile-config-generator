package main

import (
	"fmt"
	"os"

	"github.com/pivotalservices/tile-config-generator/commands"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&Manager, flags.HelpFlag)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

type manager struct {
	Generate commands.Generate `command:"generate" description:"generates configuration template that can be used with om configure-product command"`
}

var Manager manager
