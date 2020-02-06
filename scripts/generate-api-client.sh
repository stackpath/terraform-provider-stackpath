#!/bin/bash

# Generate Go API client using openapi-generator
docker run --rm -v ${PWD}:/local -w /local openapitools/openapi-generator-cli \
  generate \
  -i /local/scripts/generate-api-client-spec.json \
  -g go \
  --git-user-id terraform-providers \
  --git-repo-id terraform-provider-stackpath/stackpath \
  -c /local/scripts/generate-api-client.yaml \
  -o /local/stackpath/api_client

# Format the project
make fmt

# Remove mod/sum files from api client
rm ./stackpath/api_client/go.mod
rm ./stackpath/api_client/go.sum
