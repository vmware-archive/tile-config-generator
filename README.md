# tile-config-generator
Tooling to help externalize the configuration of Pivotal Operations Manager and it's various tiles.

## Deprecation Notice ##
The functionality of this tool has been migrated to [om](https://github.com/pivotal-cf/om) using `config-template` command.  Discontinue usage of this tooling and start to leverage the capabilities within `om` as well as open any issues to that repository.

## Getting Started

## Maintainer

* [Caleb Washburn](https://github.com/calebwashburn)

## Support

tile-config-generator is a community supported cli.  Opening issues for questions, feature requests and/or bugs is the best path to getting "support".  We strive to be active in keeping this tool working and meeting your needs in a timely fashion.

## Build from the source

`tile-config-generator` is written in [Go](https://golang.org/).
To build the binary yourself, follow these steps:

* Install `Go`.
* Install [dep](https://github.com/golang/dep), a dependency management tool for Go.
* Clone the repo:
  - `mkdir -p $(go env GOPATH)/src/github.com/pivotalservices`
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
  --pivotal-file-path=              path to pivotal file
  --base-directory=                 base directory to place generated config templates
  --do-not-include-product-version  flag to use a flat output folder
  --include-errands                 feature flag to include errands
```

## Example Usage

```
tile-config-generator generate --pivotal-file-path ~/Downloads/cf-2.1.3-build.1.pivotal --base-directory ~/sample
```

Produces the following output

```
└── cf
    └── 2.1.3
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
        ├── network
        │   ├── 2-az-configuration.yml
        │   └── 3-az-configuration.yml
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
        │   ├── add-push_apps_manager_logo.yml
        │   ├── add-push_apps_manager_marketplace_name.yml
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
        ├── resource
        │   ├── backup-prepare_additional_vm_extensions.yml
        │   ├── backup-prepare_elb_names.yml
        │   ├── backup-prepare_internet_connected.yml
        │   ├── clock_global_additional_vm_extensions.yml
        │   ├── clock_global_elb_names.yml
        │   ├── clock_global_internet_connected.yml
        │   ├── cloud_controller_additional_vm_extensions.yml
        │   ├── cloud_controller_elb_names.yml
        │   ├── cloud_controller_internet_connected.yml
        │   ├── cloud_controller_worker_additional_vm_extensions.yml
        │   ├── cloud_controller_worker_elb_names.yml
        │   ├── cloud_controller_worker_internet_connected.yml
        │   ├── consul_server_additional_vm_extensions.yml
        │   ├── consul_server_elb_names.yml
        │   ├── consul_server_internet_connected.yml
        │   ├── credhub_additional_vm_extensions.yml
        │   ├── credhub_elb_names.yml
        │   ├── credhub_internet_connected.yml
        │   ├── diego_brain_additional_vm_extensions.yml
        │   ├── diego_brain_elb_names.yml
        │   ├── diego_brain_internet_connected.yml
        │   ├── diego_cell_additional_vm_extensions.yml
        │   ├── diego_cell_elb_names.yml
        │   ├── diego_cell_internet_connected.yml
        │   ├── diego_database_additional_vm_extensions.yml
        │   ├── diego_database_elb_names.yml
        │   ├── diego_database_internet_connected.yml
        │   ├── doppler_additional_vm_extensions.yml
        │   ├── doppler_elb_names.yml
        │   ├── doppler_internet_connected.yml
        │   ├── ha_proxy_additional_vm_extensions.yml
        │   ├── ha_proxy_elb_names.yml
        │   ├── ha_proxy_internet_connected.yml
        │   ├── loggregator_trafficcontroller_additional_vm_extensions.yml
        │   ├── loggregator_trafficcontroller_elb_names.yml
        │   ├── loggregator_trafficcontroller_internet_connected.yml
        │   ├── mysql_additional_vm_extensions.yml
        │   ├── mysql_elb_names.yml
        │   ├── mysql_internet_connected.yml
        │   ├── mysql_monitor_additional_vm_extensions.yml
        │   ├── mysql_monitor_elb_names.yml
        │   ├── mysql_monitor_internet_connected.yml
        │   ├── mysql_proxy_additional_vm_extensions.yml
        │   ├── mysql_proxy_elb_names.yml
        │   ├── mysql_proxy_internet_connected.yml
        │   ├── nats_additional_vm_extensions.yml
        │   ├── nats_elb_names.yml
        │   ├── nats_internet_connected.yml
        │   ├── nfs_server_additional_vm_extensions.yml
        │   ├── nfs_server_elb_names.yml
        │   ├── nfs_server_internet_connected.yml
        │   ├── router_additional_vm_extensions.yml
        │   ├── router_elb_names.yml
        │   ├── router_internet_connected.yml
        │   ├── service-discovery-controller_additional_vm_extensions.yml
        │   ├── service-discovery-controller_elb_names.yml
        │   ├── service-discovery-controller_internet_connected.yml
        │   ├── syslog_adapter_additional_vm_extensions.yml
        │   ├── syslog_adapter_elb_names.yml
        │   ├── syslog_adapter_internet_connected.yml
        │   ├── syslog_scheduler_additional_vm_extensions.yml
        │   ├── syslog_scheduler_elb_names.yml
        │   ├── syslog_scheduler_internet_connected.yml
        │   ├── tcp_router_additional_vm_extensions.yml
        │   ├── tcp_router_elb_names.yml
        │   ├── tcp_router_internet_connected.yml
        │   ├── uaa_additional_vm_extensions.yml
        │   ├── uaa_elb_names.yml
        │   └── uaa_internet_connected.yml
        └── resource-vars.yml

```


## `display`

`display` shows a table of different configurations

### Command Usage

```
tile-config-generator [OPTIONS] display [generate-OPTIONS]

Help Options:
-h, --help                   Show this help message

[generate command options]
  --pivotal-file-path= path to pivotal file
```


## Example Usage

```
tile-config-generator display --pivotal-file-path ~/Downloads/cf-2.1.3-build.1.pivotal
```

Produces the following output

```
*****  Required Properties ******* (product.yml)
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
|                                           NAME                                            |                                   PARAMETER                                   |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.allow_app_ssh_access                                                    | cloud_controller/allow_app_ssh_access                                         |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.apps_domain                                                             | cloud_controller/apps_domain                                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.default_app_memory                                                      | cloud_controller/default_app_memory                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.default_app_ssh_access                                                  | cloud_controller/default_app_ssh_access                                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.default_disk_quota_app                                                  | cloud_controller/default_disk_quota_app                                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.default_quota_max_number_services                                       | cloud_controller/default_quota_max_number_services                            |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.default_quota_memory_limit_mb                                           | cloud_controller/default_quota_memory_limit_mb                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.enable_custom_buildpacks                                                | cloud_controller/enable_custom_buildpacks                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.max_disk_quota_app                                                      | cloud_controller/max_disk_quota_app                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.max_file_size                                                           | cloud_controller/max_file_size                                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.security_event_logging_enabled                                          | cloud_controller/security_event_logging_enabled                               |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.staging_timeout_in_seconds                                              | cloud_controller/staging_timeout_in_seconds                                   |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .cloud_controller.system_domain                                                           | cloud_controller/system_domain                                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .diego_brain.starting_container_count_maximum                                             | diego_brain/starting_container_count_maximum                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .doppler.message_drain_buffer_size                                                        | doppler/message_drain_buffer_size                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .ha_proxy.skip_cert_verify                                                                | ha_proxy/skip_cert_verify                                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql.cli_history                                                                        | mysql/cli_history                                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql.prevent_node_auto_rejoin                                                           | mysql/prevent_node_auto_rejoin                                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql.remote_admin_access                                                                | mysql/remote_admin_access                                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql_monitor.poll_frequency                                                             | mysql_monitor/poll_frequency                                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql_monitor.recipient_email                                                            | mysql_monitor/recipient_email                                                 |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql_monitor.write_read_delay                                                           | mysql_monitor/write_read_delay                                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql_proxy.shutdown_delay                                                               | mysql_proxy/shutdown_delay                                                    |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .mysql_proxy.startup_delay                                                                | mysql_proxy/startup_delay                                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .nfs_server.blobstore_internal_access_rules                                               | nfs_server/blobstore_internal_access_rules                                    |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.autoscale_api_instance_count                                                  | autoscale_api_instance_count                                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.autoscale_instance_count                                                      | autoscale_instance_count                                                      |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.autoscale_metric_bucket_count                                                 | autoscale_metric_bucket_count                                                 |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.autoscale_scaling_interval_in_seconds                                         | autoscale_scaling_interval_in_seconds                                         |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.cf_networking_enable_space_developer_self_service                             | cf_networking_enable_space_developer_self_service                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.dns_servers                        | container_networking_interface_plugin/silk/dns_servers                        |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.enable_log_traffic                 | container_networking_interface_plugin/silk/enable_log_traffic                 |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.iptables_accepted_udp_logs_per_sec | container_networking_interface_plugin/silk/iptables_accepted_udp_logs_per_sec |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.iptables_denied_logs_per_sec       | container_networking_interface_plugin/silk/iptables_denied_logs_per_sec       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.network_cidr                       | container_networking_interface_plugin/silk/network_cidr                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.network_mtu                        | container_networking_interface_plugin/silk/network_mtu                        |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.container_networking_interface_plugin.silk.vtep_port                          | container_networking_interface_plugin/silk/vtep_port                          |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.credhub_key_encryption_passwords                                              | credhub_key_encryption_passwords_0/name                                       |
|                                                                                           | credhub_key_encryption_passwords_0/provider                                   |
|                                                                                           | credhub_key_encryption_passwords_0/key                                        |
|                                                                                           | credhub_key_encryption_passwords_0/primary                                    |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.enable_grootfs                                                                | enable_grootfs                                                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.enable_service_discovery_for_apps                                             | enable_service_discovery_for_apps                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.gorouter_ssl_ciphers                                                          | gorouter_ssl_ciphers                                                          |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.haproxy_forward_tls.enable.backend_ca                                         | haproxy_forward_tls/enable/backend_ca                                         |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.haproxy_max_buffer_size                                                       | haproxy_max_buffer_size                                                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.haproxy_ssl_ciphers                                                           | haproxy_ssl_ciphers                                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.mysql_activity_logging.enable.audit_logging_events                            | mysql_activity_logging/enable/audit_logging_events                            |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.networking_poe_ssl_certs                                                      | networking_poe_ssl_certs_0/name                                               |
|                                                                                           | networking_poe_ssl_certs_0/certificate                                        |
|                                                                                           | networking_poe_ssl_certs_0/privatekey                                         |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.nfs_volume_driver.enable.ldap_server_host                                     | nfs_volume_driver/enable/ldap_server_host                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.nfs_volume_driver.enable.ldap_server_port                                     | nfs_volume_driver/enable/ldap_server_port                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.nfs_volume_driver.enable.ldap_service_account_password                        | nfs_volume_driver/enable/ldap_service_account_password                        |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.nfs_volume_driver.enable.ldap_service_account_user                            | nfs_volume_driver/enable/ldap_service_account_user                            |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.nfs_volume_driver.enable.ldap_user_fqdn                                       | nfs_volume_driver/enable/ldap_user_fqdn                                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.push_apps_manager_currency_lookup                                             | push_apps_manager_currency_lookup                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.push_apps_manager_display_plan_prices                                         | push_apps_manager_display_plan_prices                                         |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.push_apps_manager_enable_invitations                                          | push_apps_manager_enable_invitations                                          |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.rep_proxy_enabled                                                             | rep_proxy_enabled                                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.route_services.enable.ignore_ssl_cert_verification                            | route_services/enable/ignore_ssl_cert_verification                            |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.router_backend_max_conn                                                       | router_backend_max_conn                                                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.router_enable_proxy                                                           | router_enable_proxy                                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.routing_disable_http                                                          | routing_disable_http                                                          |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.saml_signature_algorithm                                                      | saml_signature_algorithm                                                      |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.secure_service_instance_credentials                                           | secure_service_instance_credentials                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.security_acknowledgement                                                      | security_acknowledgement                                                      |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.smtp_auth_mechanism                                                           | smtp_auth_mechanism                                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.smtp_enable_starttls_auto                                                     | smtp_enable_starttls_auto                                                     |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.syslog_metrics_to_syslog_enabled                                              | syslog_metrics_to_syslog_enabled                                              |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.syslog_use_tcp_for_file_forwarding_local_transport                            | syslog_use_tcp_for_file_forwarding_local_transport                            |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_expires_after_months                                    | uaa/internal/password_expires_after_months                                    |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_max_retry                                               | uaa/internal/password_max_retry                                               |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_min_length                                              | uaa/internal/password_min_length                                              |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_min_lowercase                                           | uaa/internal/password_min_lowercase                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_min_numeric                                             | uaa/internal/password_min_numeric                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_min_special                                             | uaa/internal/password_min_special                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa.internal.password_min_uppercase                                           | uaa/internal/password_min_uppercase                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa_session_cookie_max_age                                                    | uaa_session_cookie_max_age                                                    |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .properties.uaa_session_idle_timeout                                                      | uaa_session_idle_timeout                                                      |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.disable_insecure_cookies                                                          | router/disable_insecure_cookies                                               |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.drain_wait                                                                        | router/drain_wait                                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.enable_isolated_routing                                                           | router/enable_isolated_routing                                                |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.enable_write_access_logs                                                          | router/enable_write_access_logs                                               |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.enable_zipkin                                                                     | router/enable_zipkin                                                          |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.frontend_idle_timeout                                                             | router/frontend_idle_timeout                                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.lb_healthy_threshold                                                              | router/lb_healthy_threshold                                                   |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .router.request_timeout_in_seconds                                                        | router/request_timeout_in_seconds                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.apps_manager_access_token_lifetime                                                   | uaa/apps_manager_access_token_lifetime                                        |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.apps_manager_refresh_token_lifetime                                                  | uaa/apps_manager_refresh_token_lifetime                                       |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.cf_cli_access_token_lifetime                                                         | uaa/cf_cli_access_token_lifetime                                              |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.cf_cli_refresh_token_lifetime                                                        | uaa/cf_cli_refresh_token_lifetime                                             |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.customize_password_label                                                             | uaa/customize_password_label                                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.customize_username_label                                                             | uaa/customize_username_label                                                  |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.proxy_ips_regex                                                                      | uaa/proxy_ips_regex                                                           |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+
| .uaa.service_provider_key_credentials                                                     | uaa/service_provider_key_credentials/certificate                              |
|                                                                                           | uaa/service_provider_key_credentials/privatekey                               |
+-------------------------------------------------------------------------------------------+-------------------------------------------------------------------------------+

*****  Default Property Values ******* (product-default-vars.yml)
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
|                                   PARAMETER                                   |                                                                                                         VALUE                                                                                                         |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| autoscale_api_instance_count                                                  | 1                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| autoscale_instance_count                                                      | 3                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| autoscale_metric_bucket_count                                                 | 35                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| autoscale_scaling_interval_in_seconds                                         | 35                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cf_networking_enable_space_developer_self_service                             | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/allow_app_ssh_access                                         | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/default_app_memory                                           | 1024                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/default_app_ssh_access                                       | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/default_disk_quota_app                                       | 1024                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/default_quota_max_number_services                            | 100                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/default_quota_memory_limit_mb                                | 10240                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/enable_custom_buildpacks                                     | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/max_disk_quota_app                                           | 2048                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/max_file_size                                                | 1024                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/security_event_logging_enabled                               | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cloud_controller/staging_timeout_in_seconds                                   | 900                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| container_networking_interface_plugin/silk/enable_log_traffic                 | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| container_networking_interface_plugin/silk/iptables_accepted_udp_logs_per_sec | 100                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| container_networking_interface_plugin/silk/iptables_denied_logs_per_sec       | 1                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| container_networking_interface_plugin/silk/network_mtu                        | 1454                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| container_networking_interface_plugin/silk/vtep_port                          | 4789                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| credhub_key_encryption_passwords_0/primary                                    | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| credhub_key_encryption_passwords_0/provider                                   | internal                                                                                                                                                                                                              |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| diego_brain/starting_container_count_maximum                                  | 200                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| doppler/message_drain_buffer_size                                             | 10000                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| enable_grootfs                                                                | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| enable_service_discovery_for_apps                                             | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| gorouter_ssl_ciphers                                                          | ECDHE-RSA-AES128-GCM-SHA256:TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| ha_proxy/skip_cert_verify                                                     | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| haproxy_max_buffer_size                                                       | 16384                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| haproxy_ssl_ciphers                                                           | DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384                                                                                                           |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql/cli_history                                                             | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql/prevent_node_auto_rejoin                                                | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql/remote_admin_access                                                     | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql_activity_logging/enable/audit_logging_events                            | connect,query                                                                                                                                                                                                         |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql_monitor/poll_frequency                                                  | 30                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql_monitor/write_read_delay                                                | 20                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql_proxy/shutdown_delay                                                    | 30                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| mysql_proxy/startup_delay                                                     | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| nfs_server/blobstore_internal_access_rules                                    | allow 10.0.0.0/8;,allow 172.16.0.0/12;,allow 192.168.0.0/16;                                                                                                                                                          |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| push_apps_manager_currency_lookup                                             | { "usd": "$", "eur": "€" }                                                                                                                                                                                            |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| push_apps_manager_display_plan_prices                                         | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| push_apps_manager_enable_invitations                                          | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| rep_proxy_enabled                                                             | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| route_services/enable/ignore_ssl_cert_verification                            | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/disable_insecure_cookies                                               | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/drain_wait                                                             | 20                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/enable_isolated_routing                                                | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/enable_write_access_logs                                               | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/enable_zipkin                                                          | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/frontend_idle_timeout                                                  | 900                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/lb_healthy_threshold                                                   | 20                                                                                                                                                                                                                    |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router/request_timeout_in_seconds                                             | 900                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router_backend_max_conn                                                       | 500                                                                                                                                                                                                                   |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| router_enable_proxy                                                           | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| routing_disable_http                                                          | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| saml_signature_algorithm                                                      | SHA256                                                                                                                                                                                                                |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| secure_service_instance_credentials                                           | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| smtp_auth_mechanism                                                           | plain                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| smtp_enable_starttls_auto                                                     | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| syslog_metrics_to_syslog_enabled                                              | true                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| syslog_use_tcp_for_file_forwarding_local_transport                            | false                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/apps_manager_access_token_lifetime                                        | 1209600                                                                                                                                                                                                               |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/apps_manager_refresh_token_lifetime                                       | 1209600                                                                                                                                                                                                               |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/cf_cli_access_token_lifetime                                              | 7200                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/cf_cli_refresh_token_lifetime                                             | 1209600                                                                                                                                                                                                               |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/customize_password_label                                                  | Password                                                                                                                                                                                                              |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/customize_username_label                                                  | Email                                                                                                                                                                                                                 |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_expires_after_months                                    | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_max_retry                                               | 5                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_min_length                                              | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_min_lowercase                                           | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_min_numeric                                             | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_min_special                                             | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/internal/password_min_uppercase                                           | 0                                                                                                                                                                                                                     |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa/proxy_ips_regex                                                           | 10\.\d{1,3}\.\d{1,3}\.\d{1,3}|192\.168\.\d{1,3}\.\d{1,3}|169\.254\.\d{1,3}\.\d{1,3}|127\.\d{1,3}\.\d{1,3}\.\d{1,3}|172\.1[6-9]{1}\.\d{1,3}\.\d{1,3}|172\.2[0-9]{1}\.\d{1,3}\.\d{1,3}|172\.3[0-1]{1}\.\d{1,3}\.\d{1,3} |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa_session_cookie_max_age                                                    | 1800                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| uaa_session_idle_timeout                                                      | 1800                                                                                                                                                                                                                  |
+-------------------------------------------------------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+

*****  Resource Property Values ******* (resource-vars.yml)
+---------------------------------------------+-----------+
|                  PARAMETER                  |   VALUE   |
+---------------------------------------------+-----------+
| backup-prepare_instance_type                | automatic |
+---------------------------------------------+-----------+
| backup-prepare_instances                    | automatic |
+---------------------------------------------+-----------+
| blobstore_instance_type                     | automatic |
+---------------------------------------------+-----------+
| blobstore_instances                         | automatic |
+---------------------------------------------+-----------+
| blobstore_persistent_disk_size              | automatic |
+---------------------------------------------+-----------+
| clock_global_instance_type                  | automatic |
+---------------------------------------------+-----------+
| clock_global_instances                      | automatic |
+---------------------------------------------+-----------+
| cloud_controller_instance_type              | automatic |
+---------------------------------------------+-----------+
| cloud_controller_instances                  | automatic |
+---------------------------------------------+-----------+
| cloud_controller_worker_instance_type       | automatic |
+---------------------------------------------+-----------+
| cloud_controller_worker_instances           | automatic |
+---------------------------------------------+-----------+
| compute_instance_type                       | automatic |
+---------------------------------------------+-----------+
| compute_instances                           | automatic |
+---------------------------------------------+-----------+
| consul_server_instance_type                 | automatic |
+---------------------------------------------+-----------+
| consul_server_instances                     | automatic |
+---------------------------------------------+-----------+
| consul_server_persistent_disk_size          | automatic |
+---------------------------------------------+-----------+
| control_instance_type                       | automatic |
+---------------------------------------------+-----------+
| control_instances                           | automatic |
+---------------------------------------------+-----------+
| credhub_instance_type                       | automatic |
+---------------------------------------------+-----------+
| credhub_instances                           | automatic |
+---------------------------------------------+-----------+
| database_instance_type                      | automatic |
+---------------------------------------------+-----------+
| database_instances                          | automatic |
+---------------------------------------------+-----------+
| database_persistent_disk_size               | automatic |
+---------------------------------------------+-----------+
| diego_brain_instance_type                   | automatic |
+---------------------------------------------+-----------+
| diego_brain_instances                       | automatic |
+---------------------------------------------+-----------+
| diego_cell_instance_type                    | automatic |
+---------------------------------------------+-----------+
| diego_cell_instances                        | automatic |
+---------------------------------------------+-----------+
| diego_database_instance_type                | automatic |
+---------------------------------------------+-----------+
| diego_database_instances                    | automatic |
+---------------------------------------------+-----------+
| doppler_instance_type                       | automatic |
+---------------------------------------------+-----------+
| doppler_instances                           | automatic |
+---------------------------------------------+-----------+
| ha_proxy_instance_type                      | automatic |
+---------------------------------------------+-----------+
| ha_proxy_instances                          | automatic |
+---------------------------------------------+-----------+
| loggregator_trafficcontroller_instance_type | automatic |
+---------------------------------------------+-----------+
| loggregator_trafficcontroller_instances     | automatic |
+---------------------------------------------+-----------+
| mysql_instance_type                         | automatic |
+---------------------------------------------+-----------+
| mysql_instances                             | automatic |
+---------------------------------------------+-----------+
| mysql_monitor_instance_type                 | automatic |
+---------------------------------------------+-----------+
| mysql_monitor_instances                     | automatic |
+---------------------------------------------+-----------+
| mysql_persistent_disk_size                  | automatic |
+---------------------------------------------+-----------+
| mysql_proxy_instance_type                   | automatic |
+---------------------------------------------+-----------+
| mysql_proxy_instances                       | automatic |
+---------------------------------------------+-----------+
| nats_instance_type                          | automatic |
+---------------------------------------------+-----------+
| nats_instances                              | automatic |
+---------------------------------------------+-----------+
| nfs_server_instance_type                    | automatic |
+---------------------------------------------+-----------+
| nfs_server_instances                        | automatic |
+---------------------------------------------+-----------+
| nfs_server_persistent_disk_size             | automatic |
+---------------------------------------------+-----------+
| push-apps-manager_instance_type             | automatic |
+---------------------------------------------+-----------+
| push-apps-manager_instances                 | automatic |
+---------------------------------------------+-----------+
| router_instance_type                        | automatic |
+---------------------------------------------+-----------+
| router_instances                            | automatic |
+---------------------------------------------+-----------+
| service-discovery-controller_instance_type  | automatic |
+---------------------------------------------+-----------+
| service-discovery-controller_instances      | automatic |
+---------------------------------------------+-----------+
| syslog_adapter_instance_type                | automatic |
+---------------------------------------------+-----------+
| syslog_adapter_instances                    | automatic |
+---------------------------------------------+-----------+
| syslog_scheduler_instance_type              | automatic |
+---------------------------------------------+-----------+
| syslog_scheduler_instances                  | automatic |
+---------------------------------------------+-----------+
| tcp_router_instance_type                    | automatic |
+---------------------------------------------+-----------+
| tcp_router_instances                        | automatic |
+---------------------------------------------+-----------+
| uaa_instance_type                           | automatic |
+---------------------------------------------+-----------+
| uaa_instances                               | automatic |
+---------------------------------------------+-----------+

*****  Features Operations Files *******
+-------------------------------------------------------------+------------------------------------------------------------------------+
|                            FILE                             |                               PARAMETERS                               |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/cc_api_rate_limit-enable.yml                       | cc_api_rate_limit/enable/general_limit                                 |
|                                                             | cc_api_rate_limit/enable/unauthenticated_limit                         |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/container_networking_interface_plugin-external.yml |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/credhub_database-external.yml                      | credhub_database/external/host                                         |
|                                                             | credhub_database/external/port                                         |
|                                                             | credhub_database/external/username                                     |
|                                                             | credhub_database/external/password                                     |
|                                                             | credhub_database/external/tls_ca                                       |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/garden_disk_cleanup-never.yml                      |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/garden_disk_cleanup-routine.yml                    |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/haproxy_forward_tls-disable.yml                    |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/haproxy_hsts_support-enable.yml                    | haproxy_hsts_support/enable/max_age                                    |
|                                                             | haproxy_hsts_support/enable/include_subdomains                         |
|                                                             | haproxy_hsts_support/enable/enable_preload                             |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/mysql_activity_logging-disable.yml                 |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/nfs_volume_driver-disable.yml                      |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/route_services-disable.yml                         |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/router_client_cert_validation-none.yml             |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/router_client_cert_validation-require.yml          |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/router_keepalive_connections-disable.yml           |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/routing_minimum_tls_version-tls_v1_0.yml           |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/routing_minimum_tls_version-tls_v1_1.yml           |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/routing_tls_termination-ha_proxy.yml               |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/routing_tls_termination-router.yml                 |                                                                        |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/smoke_tests-specified.yml                          | smoke_tests/specified/org_name                                         |
|                                                             | smoke_tests/specified/space_name                                       |
|                                                             | smoke_tests/specified/apps_domain                                      |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/syslog_tls-enabled.yml                             | syslog_tls/enabled/tls_ca_cert                                         |
|                                                             | syslog_tls/enabled/tls_permitted_peer                                  |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/system_blobstore-external.yml                      | system_blobstore/external/endpoint                                     |
|                                                             | system_blobstore/external/buildpacks_bucket                            |
|                                                             | system_blobstore/external/droplets_bucket                              |
|                                                             | system_blobstore/external/packages_bucket                              |
|                                                             | system_blobstore/external/resources_bucket                             |
|                                                             | system_blobstore/external/access_key                                   |
|                                                             | system_blobstore/external/secret_key                                   |
|                                                             | system_blobstore/external/signature_version                            |
|                                                             | system_blobstore/external/region                                       |
|                                                             | system_blobstore/external/encryption                                   |
|                                                             | system_blobstore/external/encryption_kms_key_id                        |
|                                                             | system_blobstore/external/versioning                                   |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/system_blobstore-external_azure.yml                | system_blobstore/external_azure/buildpacks_container                   |
|                                                             | system_blobstore/external_azure/droplets_container                     |
|                                                             | system_blobstore/external_azure/packages_container                     |
|                                                             | system_blobstore/external_azure/resources_container                    |
|                                                             | system_blobstore/external_azure/account_name                           |
|                                                             | system_blobstore/external_azure/access_key                             |
|                                                             | system_blobstore/external_azure/environment                            |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/system_blobstore-external_gcs.yml                  | system_blobstore/external_gcs/buildpacks_bucket                        |
|                                                             | system_blobstore/external_gcs/droplets_bucket                          |
|                                                             | system_blobstore/external_gcs/packages_bucket                          |
|                                                             | system_blobstore/external_gcs/resources_bucket                         |
|                                                             | system_blobstore/external_gcs/access_key                               |
|                                                             | system_blobstore/external_gcs/secret_key                               |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/system_blobstore-external_gcs_service_account.yml  | system_blobstore/external_gcs_service_account/buildpacks_bucket        |
|                                                             | system_blobstore/external_gcs_service_account/droplets_bucket          |
|                                                             | system_blobstore/external_gcs_service_account/packages_bucket          |
|                                                             | system_blobstore/external_gcs_service_account/resources_bucket         |
|                                                             | system_blobstore/external_gcs_service_account/project_id               |
|                                                             | system_blobstore/external_gcs_service_account/service_account_email    |
|                                                             | system_blobstore/external_gcs_service_account/service_account_json_key |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/system_database-external.yml                       | system_database/external/host                                          |
|                                                             | system_database/external/port                                          |
|                                                             | system_database/external/account_username                              |
|                                                             | system_database/external/account_password                              |
|                                                             | system_database/external/app_usage_service_username                    |
|                                                             | system_database/external/app_usage_service_password                    |
|                                                             | system_database/external/autoscale_username                            |
|                                                             | system_database/external/autoscale_password                            |
|                                                             | system_database/external/ccdb_username                                 |
|                                                             | system_database/external/ccdb_password                                 |
|                                                             | system_database/external/diego_username                                |
|                                                             | system_database/external/diego_password                                |
|                                                             | system_database/external/locket_username                               |
|                                                             | system_database/external/locket_password                               |
|                                                             | system_database/external/networkpolicyserver_username                  |
|                                                             | system_database/external/networkpolicyserver_password                  |
|                                                             | system_database/external/nfsvolume_username                            |
|                                                             | system_database/external/nfsvolume_password                            |
|                                                             | system_database/external/notifications_username                        |
|                                                             | system_database/external/notifications_password                        |
|                                                             | system_database/external/routing_username                              |
|                                                             | system_database/external/routing_password                              |
|                                                             | system_database/external/silk_username                                 |
|                                                             | system_database/external/silk_password                                 |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/tcp_routing-enable.yml                             | tcp_routing/enable/reservable_ports                                    |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/uaa-ldap.yml                                       | uaa/ldap/url                                                           |
|                                                             | uaa/ldap/credentials                                                   |
|                                                             | uaa/ldap/search_base                                                   |
|                                                             | uaa/ldap/search_filter                                                 |
|                                                             | uaa/ldap/group_search_base                                             |
|                                                             | uaa/ldap/group_search_filter                                           |
|                                                             | uaa/ldap/server_ssl_cert                                               |
|                                                             | uaa/ldap/server_ssl_cert_alias                                         |
|                                                             | uaa/ldap/mail_attribute_name                                           |
|                                                             | uaa/ldap/email_domains                                                 |
|                                                             | uaa/ldap/first_name_attribute                                          |
|                                                             | uaa/ldap/last_name_attribute                                           |
|                                                             | uaa/ldap/ldap_referrals                                                |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/uaa-saml.yml                                       | uaa/saml/sso_name                                                      |
|                                                             | uaa/saml/display_name                                                  |
|                                                             | uaa/saml/sso_url                                                       |
|                                                             | uaa/saml/name_id_format                                                |
|                                                             | uaa/saml/sso_xml                                                       |
|                                                             | uaa/saml/sign_auth_requests                                            |
|                                                             | uaa/saml/require_signed_assertions                                     |
|                                                             | uaa/saml/email_domains                                                 |
|                                                             | uaa/saml/first_name_attribute                                          |
|                                                             | uaa/saml/last_name_attribute                                           |
|                                                             | uaa/saml/email_attribute                                               |
|                                                             | uaa/saml/external_groups_attribute                                     |
|                                                             | uaa/saml/entity_id_override                                            |
+-------------------------------------------------------------+------------------------------------------------------------------------+
| features/uaa_database-external.yml                          | uaa_database/external/host                                             |
|                                                             | uaa_database/external/port                                             |
|                                                             | uaa_database/external/uaa_username                                     |
|                                                             | uaa_database/external/uaa_password                                     |
+-------------------------------------------------------------+------------------------------------------------------------------------+

*****  Optional Operations Files *******
+------------------------------------------------------------------+--------------------------------------------------------+
|                               FILE                               |                       PARAMETERS                       |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-1-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-1-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-10-credhub_hsm_provider_servers.yml                 | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/host_address            |
|                                                                  | credhub_hsm_provider_servers_4/port                    |
|                                                                  | credhub_hsm_provider_servers_4/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_4/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_5/host_address            |
|                                                                  | credhub_hsm_provider_servers_5/port                    |
|                                                                  | credhub_hsm_provider_servers_5/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_5/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_6/host_address            |
|                                                                  | credhub_hsm_provider_servers_6/port                    |
|                                                                  | credhub_hsm_provider_servers_6/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_6/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_7/host_address            |
|                                                                  | credhub_hsm_provider_servers_7/port                    |
|                                                                  | credhub_hsm_provider_servers_7/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_7/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_8/host_address            |
|                                                                  | credhub_hsm_provider_servers_8/port                    |
|                                                                  | credhub_hsm_provider_servers_8/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_8/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_9/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_9/host_address            |
|                                                                  | credhub_hsm_provider_servers_9/port                    |
|                                                                  | credhub_hsm_provider_servers_9/partition_serial_number |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-10-push_apps_manager_footer_links.yml               | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
|                                                                  | push_apps_manager_footer_links_4/name                  |
|                                                                  | push_apps_manager_footer_links_4/href                  |
|                                                                  | push_apps_manager_footer_links_5/name                  |
|                                                                  | push_apps_manager_footer_links_5/href                  |
|                                                                  | push_apps_manager_footer_links_6/name                  |
|                                                                  | push_apps_manager_footer_links_6/href                  |
|                                                                  | push_apps_manager_footer_links_7/name                  |
|                                                                  | push_apps_manager_footer_links_7/href                  |
|                                                                  | push_apps_manager_footer_links_8/name                  |
|                                                                  | push_apps_manager_footer_links_8/href                  |
|                                                                  | push_apps_manager_footer_links_9/name                  |
|                                                                  | push_apps_manager_footer_links_9/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-2-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-2-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-3-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-3-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-4-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-4-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-5-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/host_address            |
|                                                                  | credhub_hsm_provider_servers_4/port                    |
|                                                                  | credhub_hsm_provider_servers_4/partition_serial_number |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-5-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
|                                                                  | push_apps_manager_footer_links_4/name                  |
|                                                                  | push_apps_manager_footer_links_4/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-6-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_4/host_address            |
|                                                                  | credhub_hsm_provider_servers_4/port                    |
|                                                                  | credhub_hsm_provider_servers_4/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_4/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_5/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_5/host_address            |
|                                                                  | credhub_hsm_provider_servers_5/port                    |
|                                                                  | credhub_hsm_provider_servers_5/partition_serial_number |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-6-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
|                                                                  | push_apps_manager_footer_links_4/name                  |
|                                                                  | push_apps_manager_footer_links_4/href                  |
|                                                                  | push_apps_manager_footer_links_5/name                  |
|                                                                  | push_apps_manager_footer_links_5/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-7-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/host_address            |
|                                                                  | credhub_hsm_provider_servers_4/port                    |
|                                                                  | credhub_hsm_provider_servers_4/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_4/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_5/host_address            |
|                                                                  | credhub_hsm_provider_servers_5/port                    |
|                                                                  | credhub_hsm_provider_servers_5/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_5/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_6/port                    |
|                                                                  | credhub_hsm_provider_servers_6/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_6/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_6/host_address            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-7-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
|                                                                  | push_apps_manager_footer_links_4/name                  |
|                                                                  | push_apps_manager_footer_links_4/href                  |
|                                                                  | push_apps_manager_footer_links_5/href                  |
|                                                                  | push_apps_manager_footer_links_5/name                  |
|                                                                  | push_apps_manager_footer_links_6/name                  |
|                                                                  | push_apps_manager_footer_links_6/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-8-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/port                    |
|                                                                  | credhub_hsm_provider_servers_4/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_4/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/host_address            |
|                                                                  | credhub_hsm_provider_servers_5/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_5/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_5/host_address            |
|                                                                  | credhub_hsm_provider_servers_5/port                    |
|                                                                  | credhub_hsm_provider_servers_6/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_6/host_address            |
|                                                                  | credhub_hsm_provider_servers_6/port                    |
|                                                                  | credhub_hsm_provider_servers_6/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_7/port                    |
|                                                                  | credhub_hsm_provider_servers_7/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_7/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_7/host_address            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-8-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
|                                                                  | push_apps_manager_footer_links_4/name                  |
|                                                                  | push_apps_manager_footer_links_4/href                  |
|                                                                  | push_apps_manager_footer_links_5/name                  |
|                                                                  | push_apps_manager_footer_links_5/href                  |
|                                                                  | push_apps_manager_footer_links_6/name                  |
|                                                                  | push_apps_manager_footer_links_6/href                  |
|                                                                  | push_apps_manager_footer_links_7/name                  |
|                                                                  | push_apps_manager_footer_links_7/href                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-9-credhub_hsm_provider_servers.yml                  | credhub_hsm_provider_servers_0/port                    |
|                                                                  | credhub_hsm_provider_servers_0/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_0/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_0/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_1/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_1/host_address            |
|                                                                  | credhub_hsm_provider_servers_1/port                    |
|                                                                  | credhub_hsm_provider_servers_2/host_address            |
|                                                                  | credhub_hsm_provider_servers_2/port                    |
|                                                                  | credhub_hsm_provider_servers_2/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_2/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_3/host_address            |
|                                                                  | credhub_hsm_provider_servers_3/port                    |
|                                                                  | credhub_hsm_provider_servers_3/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_3/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_4/host_address            |
|                                                                  | credhub_hsm_provider_servers_4/port                    |
|                                                                  | credhub_hsm_provider_servers_4/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_4/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_5/host_address            |
|                                                                  | credhub_hsm_provider_servers_5/port                    |
|                                                                  | credhub_hsm_provider_servers_5/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_5/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_6/port                    |
|                                                                  | credhub_hsm_provider_servers_6/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_6/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_6/host_address            |
|                                                                  | credhub_hsm_provider_servers_7/partition_serial_number |
|                                                                  | credhub_hsm_provider_servers_7/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_7/host_address            |
|                                                                  | credhub_hsm_provider_servers_7/port                    |
|                                                                  | credhub_hsm_provider_servers_8/hsm_certificate         |
|                                                                  | credhub_hsm_provider_servers_8/host_address            |
|                                                                  | credhub_hsm_provider_servers_8/port                    |
|                                                                  | credhub_hsm_provider_servers_8/partition_serial_number |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-9-push_apps_manager_footer_links.yml                | push_apps_manager_footer_links_0/name                  |
|                                                                  | push_apps_manager_footer_links_0/href                  |
|                                                                  | push_apps_manager_footer_links_1/name                  |
|                                                                  | push_apps_manager_footer_links_1/href                  |
|                                                                  | push_apps_manager_footer_links_2/href                  |
|                                                                  | push_apps_manager_footer_links_2/name                  |
|                                                                  | push_apps_manager_footer_links_3/name                  |
|                                                                  | push_apps_manager_footer_links_3/href                  |
|                                                                  | push_apps_manager_footer_links_4/name                  |
|                                                                  | push_apps_manager_footer_links_4/href                  |
|                                                                  | push_apps_manager_footer_links_5/name                  |
|                                                                  | push_apps_manager_footer_links_5/href                  |
|                                                                  | push_apps_manager_footer_links_6/name                  |
|                                                                  | push_apps_manager_footer_links_6/href                  |
|                                                                  | push_apps_manager_footer_links_7/name                  |
|                                                                  | push_apps_manager_footer_links_7/href                  |
|                                                                  | push_apps_manager_footer_links_8/href                  |
|                                                                  | push_apps_manager_footer_links_8/name                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-cf_dial_timeout_in_seconds.yml                      | cf_dial_timeout_in_seconds                             |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-cloud_controller-encrypt_key.yml                    | cloud_controller/encrypt_key                           |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-credhub_hsm_provider_client_certificate.yml         | credhub_hsm_provider_client_certificate/certificate    |
|                                                                  | credhub_hsm_provider_client_certificate/privatekey     |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-credhub_hsm_provider_partition.yml                  | credhub_hsm_provider_partition                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-credhub_hsm_provider_partition_password.yml         | credhub_hsm_provider_partition_password                |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-diego_brain-static_ips.yml                          | diego_brain/static_ips                                 |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-diego_cell-executor_disk_capacity.yml               | diego_cell/executor_disk_capacity                      |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-diego_cell-executor_memory_capacity.yml             | diego_cell/executor_memory_capacity                    |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-diego_cell-insecure_docker_registry_list.yml        | diego_cell/insecure_docker_registry_list               |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-ha_proxy-internal_only_domains.yml                  | ha_proxy/internal_only_domains                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-ha_proxy-static_ips.yml                             | ha_proxy/static_ips                                    |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-ha_proxy-trusted_domain_cidrs.yml                   | ha_proxy/trusted_domain_cidrs                          |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-logger_endpoint_port.yml                            | logger_endpoint_port                                   |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-mysql-cluster_probe_timeout.yml                     | mysql/cluster_probe_timeout                            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-mysql_proxy-service_hostname.yml                    | mysql_proxy/service_hostname                           |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-mysql_proxy-static_ips.yml                          | mysql_proxy/static_ips                                 |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_accent_color.yml                  | push_apps_manager_accent_color                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_company_name.yml                  | push_apps_manager_company_name                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_favicon.yml                       | push_apps_manager_favicon                              |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_footer_text.yml                   | push_apps_manager_footer_text                          |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_global_wrapper_bg_color.yml       | push_apps_manager_global_wrapper_bg_color              |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_global_wrapper_footer_content.yml | push_apps_manager_global_wrapper_footer_content        |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_global_wrapper_header_content.yml | push_apps_manager_global_wrapper_header_content        |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_global_wrapper_text_color.yml     | push_apps_manager_global_wrapper_text_color            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_logo.yml                          | push_apps_manager_logo                                 |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_marketplace_name.yml              | push_apps_manager_marketplace_name                     |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_product_name.yml                  | push_apps_manager_product_name                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-push_apps_manager_square_logo.yml                   | push_apps_manager_square_logo                          |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-router-extra_headers_to_log.yml                     | router/extra_headers_to_log                            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-router-static_ips.yml                               | router/static_ips                                      |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-routing_custom_ca_certificates.yml                  | routing_custom_ca_certificates                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-saml_entity_id_override.yml                         | saml_entity_id_override                                |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-smtp_address.yml                                    | smtp_address                                           |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-smtp_crammd5_secret.yml                             | smtp_crammd5_secret                                    |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-smtp_credentials.yml                                | smtp_credentials                                       |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-smtp_from.yml                                       | smtp_from                                              |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-smtp_port.yml                                       | smtp_port                                              |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-syslog_host.yml                                     | syslog_host                                            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-syslog_port.yml                                     | syslog_port                                            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-syslog_protocol.yml                                 | syslog_protocol                                        |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-syslog_rule.yml                                     | syslog_rule                                            |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-tcp_router-static_ips.yml                           | tcp_router/static_ips                                  |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-uaa-issuer_uri.yml                                  | uaa/issuer_uri                                         |
+------------------------------------------------------------------+--------------------------------------------------------+
| optional/add-uaa-service_provider_key_password.yml               | uaa/service_provider_key_password                      |
+------------------------------------------------------------------+--------------------------------------------------------+

*****  Resource Operations Files *******
+---------------------------------------------------------------------+--------------------------------------------------------+
|                                FILE                                 |                       PARAMETERS                       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/backup-prepare_additional_vm_extensions.yml                | backup-prepare_additional_vm_extensions                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/backup-prepare_elb_names.yml                               | backup-prepare_elb_names                               |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/backup-prepare_internet_connected.yml                      | backup-prepare_internet_connected                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/clock_global_additional_vm_extensions.yml                  | clock_global_additional_vm_extensions                  |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/clock_global_elb_names.yml                                 | clock_global_elb_names                                 |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/clock_global_internet_connected.yml                        | clock_global_internet_connected                        |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/cloud_controller_additional_vm_extensions.yml              | cloud_controller_additional_vm_extensions              |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/cloud_controller_elb_names.yml                             | cloud_controller_elb_names                             |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/cloud_controller_internet_connected.yml                    | cloud_controller_internet_connected                    |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/cloud_controller_worker_additional_vm_extensions.yml       | cloud_controller_worker_additional_vm_extensions       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/cloud_controller_worker_elb_names.yml                      | cloud_controller_worker_elb_names                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/cloud_controller_worker_internet_connected.yml             | cloud_controller_worker_internet_connected             |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/consul_server_additional_vm_extensions.yml                 | consul_server_additional_vm_extensions                 |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/consul_server_elb_names.yml                                | consul_server_elb_names                                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/consul_server_internet_connected.yml                       | consul_server_internet_connected                       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/credhub_additional_vm_extensions.yml                       | credhub_additional_vm_extensions                       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/credhub_elb_names.yml                                      | credhub_elb_names                                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/credhub_internet_connected.yml                             | credhub_internet_connected                             |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_brain_additional_vm_extensions.yml                   | diego_brain_additional_vm_extensions                   |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_brain_elb_names.yml                                  | diego_brain_elb_names                                  |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_brain_internet_connected.yml                         | diego_brain_internet_connected                         |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_cell_additional_vm_extensions.yml                    | diego_cell_additional_vm_extensions                    |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_cell_elb_names.yml                                   | diego_cell_elb_names                                   |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_cell_internet_connected.yml                          | diego_cell_internet_connected                          |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_database_additional_vm_extensions.yml                | diego_database_additional_vm_extensions                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_database_elb_names.yml                               | diego_database_elb_names                               |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/diego_database_internet_connected.yml                      | diego_database_internet_connected                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/doppler_additional_vm_extensions.yml                       | doppler_additional_vm_extensions                       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/doppler_elb_names.yml                                      | doppler_elb_names                                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/doppler_internet_connected.yml                             | doppler_internet_connected                             |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/ha_proxy_additional_vm_extensions.yml                      | ha_proxy_additional_vm_extensions                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/ha_proxy_elb_names.yml                                     | ha_proxy_elb_names                                     |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/ha_proxy_internet_connected.yml                            | ha_proxy_internet_connected                            |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/loggregator_trafficcontroller_additional_vm_extensions.yml | loggregator_trafficcontroller_additional_vm_extensions |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/loggregator_trafficcontroller_elb_names.yml                | loggregator_trafficcontroller_elb_names                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/loggregator_trafficcontroller_internet_connected.yml       | loggregator_trafficcontroller_internet_connected       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_additional_vm_extensions.yml                         | mysql_additional_vm_extensions                         |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_elb_names.yml                                        | mysql_elb_names                                        |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_internet_connected.yml                               | mysql_internet_connected                               |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_monitor_additional_vm_extensions.yml                 | mysql_monitor_additional_vm_extensions                 |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_monitor_elb_names.yml                                | mysql_monitor_elb_names                                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_monitor_internet_connected.yml                       | mysql_monitor_internet_connected                       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_proxy_additional_vm_extensions.yml                   | mysql_proxy_additional_vm_extensions                   |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_proxy_elb_names.yml                                  | mysql_proxy_elb_names                                  |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/mysql_proxy_internet_connected.yml                         | mysql_proxy_internet_connected                         |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/nats_additional_vm_extensions.yml                          | nats_additional_vm_extensions                          |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/nats_elb_names.yml                                         | nats_elb_names                                         |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/nats_internet_connected.yml                                | nats_internet_connected                                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/nfs_server_additional_vm_extensions.yml                    | nfs_server_additional_vm_extensions                    |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/nfs_server_elb_names.yml                                   | nfs_server_elb_names                                   |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/nfs_server_internet_connected.yml                          | nfs_server_internet_connected                          |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/router_additional_vm_extensions.yml                        | router_additional_vm_extensions                        |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/router_elb_names.yml                                       | router_elb_names                                       |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/router_internet_connected.yml                              | router_internet_connected                              |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/service-discovery-controller_additional_vm_extensions.yml  | service-discovery-controller_additional_vm_extensions  |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/service-discovery-controller_elb_names.yml                 | service-discovery-controller_elb_names                 |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/service-discovery-controller_internet_connected.yml        | service-discovery-controller_internet_connected        |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/syslog_adapter_additional_vm_extensions.yml                | syslog_adapter_additional_vm_extensions                |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/syslog_adapter_elb_names.yml                               | syslog_adapter_elb_names                               |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/syslog_adapter_internet_connected.yml                      | syslog_adapter_internet_connected                      |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/syslog_scheduler_additional_vm_extensions.yml              | syslog_scheduler_additional_vm_extensions              |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/syslog_scheduler_elb_names.yml                             | syslog_scheduler_elb_names                             |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/syslog_scheduler_internet_connected.yml                    | syslog_scheduler_internet_connected                    |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/tcp_router_additional_vm_extensions.yml                    | tcp_router_additional_vm_extensions                    |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/tcp_router_elb_names.yml                                   | tcp_router_elb_names                                   |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/tcp_router_internet_connected.yml                          | tcp_router_internet_connected                          |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/uaa_additional_vm_extensions.yml                           | uaa_additional_vm_extensions                           |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/uaa_elb_names.yml                                          | uaa_elb_names                                          |
+---------------------------------------------------------------------+--------------------------------------------------------+
| resource/uaa_internet_connected.yml                                 | uaa_internet_connected                                 |
+---------------------------------------------------------------------+--------------------------------------------------------+

*****  Network Operations Files *******
+--------------------------------+------------+
|              FILE              | PARAMETERS |
+--------------------------------+------------+
| network/2-az-configuration.yml | az2_name   |
+--------------------------------+------------+
| network/3-az-configuration.yml | az2_name   |
|                                | az3_name   |
+--------------------------------+------------+
```
