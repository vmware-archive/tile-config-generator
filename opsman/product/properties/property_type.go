package properties

type PropertyType string

const (
	PropertyTypeBoolean                      PropertyType = "boolean"
	PropertyTypeCACertificate                PropertyType = "ca_certificate"
	PropertyTypeCollection                   PropertyType = "collection"
	PropertyTypeDropdownSelect               PropertyType = "dropdown_select"
	PropertyTypeDiskTypeDropdown             PropertyType = "disk_type_dropdown"
	PropertyTypeEmail                        PropertyType = "email"
	PropertyTypeHTTPURL                      PropertyType = "http_url"
	PropertyTypeInteger                      PropertyType = "integer"
	PropertyTypeIPRanges                     PropertyType = "ip_ranges"
	PropertyTypeLDAPURL                      PropertyType = "ldap_url"
	PropertyTypeMultiSelectOptions           PropertyType = "multi_select_options"
	PropertyTypeNetworkAddress               PropertyType = "network_address"
	PropertyTypePort                         PropertyType = "port"
	PropertyTypeRSACertCredentials           PropertyType = "rsa_cert_credentials"
	PropertyTypeRSAPKeyCredentials           PropertyType = "rsa_pkey_credentials"
	PropertyTypeSaltedCredentials            PropertyType = "salted_credentials"
	PropertyTypeSecret                       PropertyType = "secret"
	PropertyTypeSelector                     PropertyType = "selector"
	PropertyTypeServiceNetworkAZMultiSelect  PropertyType = "service_network_az_multi_select"
	PropertyTypeServiceNetworkAZSingleSelect PropertyType = "service_network_az_single_select"
	PropertyTypeSimpleCredentials            PropertyType = "simple_credentials"
	PropertyTypeString                       PropertyType = "string"
	PropertyTypeStringList                   PropertyType = "string_list"
	PropertyTypeText                         PropertyType = "text"
	PropertyTypeVMTypeDropdown               PropertyType = "vm_type_dropdown"
	PropertyTypeWildcardDomain               PropertyType = "wildcard_domain"
	PropertyTypeUUID                         PropertyType = "uuid"
)
