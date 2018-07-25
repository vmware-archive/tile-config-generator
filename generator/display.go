package generator

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Displayer struct {
	PathToPivotalFile string
	Writer            io.Writer
}

func NewDisplayer(filePath string, writer io.Writer) *Displayer {
	return &Displayer{
		PathToPivotalFile: filePath,
		Writer:            writer,
	}
}

func (d *Displayer) Display() error {
	metadataBytes, err := extractMetadataBytes(d.PathToPivotalFile)
	if err != nil {
		return err
	}
	metadata, err := NewMetadata(metadataBytes)
	if err != nil {
		return err
	}

	err = d.requiredTable(metadata)
	if err != nil {
		return err
	}

	d.Writer.Write([]byte("\n"))
	err = d.defaultsTable(metadata)
	if err != nil {
		return err
	}

	d.Writer.Write([]byte("\n"))
	err = d.resourceDefaultsTable(metadata)
	if err != nil {
		return err
	}
	d.Writer.Write([]byte("\n"))
	err = d.errandDefaultsTable(metadata)
	if err != nil {
		return err
	}

	d.Writer.Write([]byte("\n"))
	opsFiles, err := CreateProductPropertiesFeaturesOpsFiles(metadata)
	if err != nil {
		return err
	}
	err = d.operationsFileTable(opsFiles, "features", "Features")
	if err != nil {
		return err
	}

	d.Writer.Write([]byte("\n"))
	opsFiles, err = CreateProductPropertiesOptionalOpsFiles(metadata)
	if err != nil {
		return err
	}
	err = d.operationsFileTable(opsFiles, "optional", "Optional")
	if err != nil {
		return err
	}

	d.Writer.Write([]byte("\n"))
	opsFiles, err = CreateResourceOpsFiles(metadata)
	if err != nil {
		return err
	}
	err = d.operationsFileTable(opsFiles, "resource", "Resource")
	if err != nil {
		return err
	}

	d.Writer.Write([]byte("\n"))
	opsFiles, err = CreateNetworkOpsFiles(metadata)
	if err != nil {
		return err
	}
	err = d.operationsFileTable(opsFiles, "network", "Network")
	if err != nil {
		return err
	}

	return nil
}

func (d *Displayer) requiredTable(metadata *Metadata) error {
	requiredProperties, err := CreateProductProperties(metadata)
	if err != nil {
		return err
	}

	var data [][]string

	var keys []string
	for k := range requiredProperties {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, propertyName := range keys {
		property := requiredProperties[propertyName]
		if !property.IsSelector() {
			parameters := d.cleanParamaters(property.Parameters())
			data = append(data, []string{propertyName, strings.Join(parameters, "\n")})
		}
	}

	d.Writer.Write([]byte("*****  Required Properties ******* (product.yml) \n"))

	table := tablewriter.NewWriter(d.Writer)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Name", "Parameter"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (d *Displayer) defaultsTable(metadata *Metadata) error {
	vars, err := CreateProductPropertiesVars(metadata)
	if err != nil {
		return err
	}

	var data [][]string
	var keys []string
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, propertyName := range keys {
		value := vars[propertyName]
		data = append(data, []string{propertyName, fmt.Sprintf("%v", value)})
	}

	d.Writer.Write([]byte("*****  Default Property Values ******* (product-default-vars.yml) \n"))

	table := tablewriter.NewWriter(d.Writer)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Parameter", "Value"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (d *Displayer) resourceDefaultsTable(metadata *Metadata) error {
	vars := CreateResourceVars(metadata)

	var data [][]string
	var keys []string
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, propertyName := range keys {
		value := vars[propertyName]
		data = append(data, []string{propertyName, fmt.Sprintf("%v", value)})
	}

	d.Writer.Write([]byte("*****  Resource Property Values ******* (resource-vars.yml) \n"))

	table := tablewriter.NewWriter(d.Writer)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Parameter", "Value"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (d *Displayer) errandDefaultsTable(metadata *Metadata) error {
	vars := CreateErrandVars(metadata)

	var data [][]string
	var keys []string
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, propertyName := range keys {
		value := vars[propertyName]
		data = append(data, []string{propertyName, fmt.Sprintf("%v", value)})
	}

	d.Writer.Write([]byte("*****  Errand Property Values ******* (errand-vars.yml) \n"))

	table := tablewriter.NewWriter(d.Writer)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Parameter", "Value"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (d *Displayer) operationsFileTable(operationsFiles map[string][]Ops, prefix, description string) error {

	var data [][]string
	var keys []string
	for k := range operationsFiles {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, fileName := range keys {
		var parameters []string
		operations := operationsFiles[fileName]
		for _, op := range operations {
			if op.Value != nil {

				parameters = append(parameters, d.cleanParamaters(op.Value.Parameters())...)
			}
		}
		data = append(data, []string{fmt.Sprintf("%s/%s.yml", prefix, fileName), strings.Join(parameters, "\n")})
	}

	d.Writer.Write([]byte(fmt.Sprintf("*****  %s Operations Files ******* \n", description)))

	table := tablewriter.NewWriter(d.Writer)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"File", "Parameters"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
	return nil
}

func (d *Displayer) cleanParamaters(parameters []string) []string {
	var rawParameters []string
	for _, parameter := range parameters {
		if strings.HasPrefix(parameter, "((") {
			rawParameters = append(rawParameters, strings.Replace(strings.Replace(parameter, "((", "", -1), "))", "", -1))
		}
	}
	return rawParameters
}
