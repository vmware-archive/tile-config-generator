package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/pivotalservices/tile-config-generator/commands"
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
	Generate commands.Generate        `command:"generate" description:"generates configuration template that can be used with om configure-product command"`
	Display  commands.Display         `command:"display" description:"displays information about tile"`
	Metdata  commands.MetadataCommand `command:"metadata" description:"gets metadata file"`
	Version  commands.Version         `command:"version" description:"displays version"`
}

var Manager manager
