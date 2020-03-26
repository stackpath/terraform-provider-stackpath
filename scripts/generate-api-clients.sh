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
  --spec=swagger/stackpath_storage.oas2.json \
  --target=stackpath/api/storage \
  --model-package=storage_models \
  --client-package=storage_client

# Patch the generated files
#
# The network policy AH, ESP, GRE, and ICMP prootocols don't have any
# parameters. If they're present in the network policy then they're on. As such
# these should be empty structs instead of interfaces and pointers in the
# upstream protocols struct. If the pointer is not nil then the protocol is
# bound to the policy.
sed -i '' -e 's/^type V1ProtocolAh interface{}$/type V1ProtocolAh struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_ah.go
sed -i '' -e 's/^type V1ProtocolEsp interface{}$/type V1ProtocolEsp struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_esp.go
sed -i '' -e 's/^type V1ProtocolGre interface{}$/type V1ProtocolGre struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_gre.go
sed -i '' -e 's/^type V1ProtocolIcmp interface{}$/type V1ProtocolIcmp struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_icmp.go

sed -i '' -e 's/^	Ah V1ProtocolAh `json:"ah,omitempty"`$/	Ah *V1ProtocolAh `json:"ah,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
sed -i '' -e 's/^	Esp V1ProtocolEsp `json:"esp,omitempty"`$/	Esp *V1ProtocolEsp `json:"esp,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
sed -i '' -e 's/^	Gre V1ProtocolGre `json:"gre,omitempty"`$/	Gre *V1ProtocolGre `json:"gre,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
sed -i '' -e 's/^	Icmp V1ProtocolIcmp `json:"icmp,omitempty"`$/	Icmp *V1ProtocolIcmp `json:"icmp,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go

# Format the project
make fmt
