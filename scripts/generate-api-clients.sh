#!/bin/bash

# Generate Go API clients using go-swagger
#
# See https://github.com/go-swagger/go-swagger for more information and
# installation instructions.
swagger generate client \
  --spec=swagger/stackpath_workload.oas2.json \
  --target=stackpath/api/workload \
  --model-package=workload_models \
  --client-package=workload_client

swagger generate client \
  --spec=swagger/stackpath_ipam.oas2.json \
  --target=stackpath/api/ipam \
  --model-package=ipam_models \
  --client-package=ipam_client

swagger generate client \
  --spec=swagger/stackpath_object_storage.oas2.json \
  --target=stackpath/api/storage \
  --model-package=storage_models \
  --client-package=storage_client

# Format the project
make fmt
