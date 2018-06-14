# tile-config-generator
Tooling to help externalize the configuration of Pivotal Operations Manager and it's various tiles.

## Getting Started

## Maintainer

* [Caleb Washburn](https://github.com/pivotalservices)

## Support

tile-config-generator is a community supported cli.  Opening issues for questions, feature requests and/or bugs is the best path to getting "support".  We strive to be active in keeping this tool working and meeting your needs in a timely fashion.

## Build from the source

`tile-config-generator` is written in [Go](https://golang.org/).
To build the binary yourself, follow these steps:

* Install `Go`.
* Install [dep](https://github.com/golang/dep), a dependency management tool for Go.
* Clone the repo:
  - `mkdir -p $(go env GOPATH)/src/github.com/tile-config-generator`
  - `cd $(go env GOPATH)/src/github.com/pivotalservices`
  - `git clone git@github.com:pivotalservices/tile-config-generator.git`
* Install dependencies:
  - `cd tile-config-generator`
  - `dep ensure`
  - `go build -o tile-config-generator cmd/tile-config-generator/main.go`

To cross compile, set the `$GOOS` and `$GOARCH` environment variables.
For example: `GOOS=linux GOARCH=amd64 go build`.

## Testing

To run the unit tests, use `go test ./generator`.

## Contributing

PRs are always welcome or open issues if you are experiencing an issue and will do my best to address issues in timely fashion.

## Documentation

## `generate`

`generate` parse the metadata from a given .pivotal file and produce the following file in the target directory under in path <product_name>/<product_version>:

- a parameterized file representing required properties that can be interpolated using tools like [bosh interpolation](https://bosh.io/docs/cli-int/)
- a resource config yaml variable file with defaults
- a properties config yaml variable file with default for required properties
- for each `selector` in the tile a yaml [operations file](https://bosh.io/docs/cli-ops-files/) that allows toggling the selector to alternative values in the features directory
- for each optional property in the tile a yaml [operations file](https://bosh.io/docs/cli-ops-files/) that allows setting optional properties in the optional directory

### Command Usage

```
tile-config-generator [OPTIONS] generate [generate-OPTIONS]

Help Options:
-h, --help                   Show this help message

[generate command options]
  --pivotal-file-path= path to pivotal file
  --base-directory=    base directory to place generated config templates
```


## Example Usage

```
tile-config-generator generate --pivotal-file-path ~/Downloads/cf-2.1.5-build.1.pivotal --base-directory ~/sample
```

Produces the following output

```
└── cf
    └── 2.1.5
        ├── features
        │   ├── cc_api_rate_limit-enable.yml
        │   ├── container_networking_interface_plugin-external.yml
        │   ├── credhub_database-external.yml
        │   ├── garden_disk_cleanup-never.yml
        │   ├── garden_disk_cleanup-routine.yml
        │   ├── haproxy_forward_tls-disable.yml
        │   ├── haproxy_hsts_support-enable.yml
        │   ├── mysql_activity_logging-disable.yml
        │   ├── nfs_volume_driver-disable.yml
        │   ├── route_services-disable.yml
        │   ├── router_client_cert_validation-none.yml
        │   ├── router_client_cert_validation-require.yml
        │   ├── router_keepalive_connections-disable.yml
        │   ├── routing_log_client_ips-disable_all_log_client_ips.yml
        │   ├── routing_log_client_ips-disable_x_forwarded_for.yml
        │   ├── routing_minimum_tls_version-tls_v1_0.yml
        │   ├── routing_minimum_tls_version-tls_v1_1.yml
        │   ├── routing_tls_termination-ha_proxy.yml
        │   ├── routing_tls_termination-router.yml
        │   ├── smoke_tests-specified.yml
        │   ├── syslog_tls-enabled.yml
        │   ├── system_blobstore-external.yml
        │   ├── system_blobstore-external_azure.yml
        │   ├── system_blobstore-external_gcs.yml
        │   ├── system_blobstore-external_gcs_service_account.yml
        │   ├── system_database-external.yml
        │   ├── tcp_routing-enable.yml
        │   ├── uaa-ldap.yml
        │   ├── uaa-saml.yml
        │   └── uaa_database-external.yml
        ├── optional
        │   ├── add-1-credhub_hsm_provider_servers.yml
        │   ├── add-1-push_apps_manager_footer_links.yml
        │   ├── add-10-credhub_hsm_provider_servers.yml
        │   ├── add-10-push_apps_manager_footer_links.yml
        │   ├── add-2-credhub_hsm_provider_servers.yml
        │   ├── add-2-push_apps_manager_footer_links.yml
        │   ├── add-3-credhub_hsm_provider_servers.yml
        │   ├── add-3-push_apps_manager_footer_links.yml
        │   ├── add-4-credhub_hsm_provider_servers.yml
        │   ├── add-4-push_apps_manager_footer_links.yml
        │   ├── add-5-credhub_hsm_provider_servers.yml
        │   ├── add-5-push_apps_manager_footer_links.yml
        │   ├── add-6-credhub_hsm_provider_servers.yml
        │   ├── add-6-push_apps_manager_footer_links.yml
        │   ├── add-7-credhub_hsm_provider_servers.yml
        │   ├── add-7-push_apps_manager_footer_links.yml
        │   ├── add-8-credhub_hsm_provider_servers.yml
        │   ├── add-8-push_apps_manager_footer_links.yml
        │   ├── add-9-credhub_hsm_provider_servers.yml
        │   ├── add-9-push_apps_manager_footer_links.yml
        │   ├── add-cf_dial_timeout_in_seconds.yml
        │   ├── add-cloud_controller-encrypt_key.yml
        │   ├── add-credhub_hsm_provider_client_certificate.yml
        │   ├── add-credhub_hsm_provider_partition.yml
        │   ├── add-credhub_hsm_provider_partition_password.yml
        │   ├── add-diego_brain-static_ips.yml
        │   ├── add-diego_cell-executor_disk_capacity.yml
        │   ├── add-diego_cell-executor_memory_capacity.yml
        │   ├── add-diego_cell-insecure_docker_registry_list.yml
        │   ├── add-ha_proxy-internal_only_domains.yml
        │   ├── add-ha_proxy-static_ips.yml
        │   ├── add-ha_proxy-trusted_domain_cidrs.yml
        │   ├── add-logger_endpoint_port.yml
        │   ├── add-mysql-cluster_probe_timeout.yml
        │   ├── add-mysql_proxy-service_hostname.yml
        │   ├── add-mysql_proxy-static_ips.yml
        │   ├── add-push_apps_manager_accent_color.yml
        │   ├── add-push_apps_manager_company_name.yml
        │   ├── add-push_apps_manager_favicon.yml
        │   ├── add-push_apps_manager_footer_text.yml
        │   ├── add-push_apps_manager_global_wrapper_bg_color.yml
        │   ├── add-push_apps_manager_global_wrapper_footer_content.yml
        │   ├── add-push_apps_manager_global_wrapper_header_content.yml
        │   ├── add-push_apps_manager_global_wrapper_text_color.yml
        │   ├── add-push_apps_manager_invitations_memory.yml
        │   ├── add-push_apps_manager_logo.yml
        │   ├── add-push_apps_manager_marketplace_name.yml
        │   ├── add-push_apps_manager_memory.yml
        │   ├── add-push_apps_manager_product_name.yml
        │   ├── add-push_apps_manager_square_logo.yml
        │   ├── add-router-extra_headers_to_log.yml
        │   ├── add-router-static_ips.yml
        │   ├── add-routing_custom_ca_certificates.yml
        │   ├── add-saml_entity_id_override.yml
        │   ├── add-smtp_address.yml
        │   ├── add-smtp_crammd5_secret.yml
        │   ├── add-smtp_credentials.yml
        │   ├── add-smtp_from.yml
        │   ├── add-smtp_port.yml
        │   ├── add-syslog_host.yml
        │   ├── add-syslog_port.yml
        │   ├── add-syslog_protocol.yml
        │   ├── add-syslog_rule.yml
        │   ├── add-tcp_router-static_ips.yml
        │   ├── add-uaa-issuer_uri.yml
        │   └── add-uaa-service_provider_key_password.yml
        ├── product-default-vars.yml
        ├── product.yml
        └── resource-vars.yml
```
