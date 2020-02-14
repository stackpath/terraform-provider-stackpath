#!/bin/bash

# Generate Go API client using openapi-generator
swagger generate client \
  --spec=swagger/stackpath_object_storage.oas2.json \
  --target=stackpath/api/object_storage

# Format the project
make fmt
