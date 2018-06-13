#!/bin/bash -e
pivotal_file=$(find product -name '*.pivotal')
mkdir -p config
tile-config-generator generate --pivotal-file-path ${pivotal_file} --base-directory config --do-not-include-product-version

product_name=$(bosh int config/product.yml --path /product_name)
product_version=$(bosh int config/product.yml --path /product_version)
opsfiles=""
for op in ${OPS_FILES}
do
  opsfiles="${opsfiles} -o config/${op}"
done

varsfiles="-l config/product-default-vars.yml -l config/resource-vars.yml"
for var in ${VARS_FILES}
do
  varsfiles="${varsfiles} -l ${var}"
done

tmpproduct=$(mktemp)
bosh int config/product.yml ${opsfiles} ${varsfiles} --var-errs > ${tmpproduct}

om -t ${OM_HOST} -k upload-product --product ${pivotal_file}
om -t ${OM_HOST} -k stage-product --product-name ${product_name} --product-version ${product_version}
om -t ${OM_HOST} -k configure-product --product-name ${product_name} --config ${tmpproduct}
