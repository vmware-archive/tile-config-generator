package generator

type Executor struct {
}

func (e *Executor) Generate(metadataBytes []byte) (*Template, error) {
	metadata, err := NewMetadata(metadataBytes)
	if err != nil {
		return nil, err
	}
	template := &Template{}
	template.NetworkProperties = NewNetworkProperties(metadata)
	template.ResourceConfig = NewResourceConfig(metadata)
	return template, nil
}
