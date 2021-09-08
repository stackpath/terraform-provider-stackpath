TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=stackpath
BINARY=terraform-provider-${PKG_NAME}
OS_ARCH=darwin_amd64

default: build

build:
	@echo "==> Building ${BINARY}..."
	go build -o ${BINARY}
	@echo

install: build
	@echo "==> Installing ${BINARY}..."
	mkdir -p ~/.terraform.d/plugins/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${OS_ARCH}
	@echo

install-13: build # https://debruyn.dev/2020/setting-up-your-machine-for-local-terraform-provider-development/
	@echo "==> Installing ${BINARY}..."
	mkdir -p ~/terraform-providers/local/providers/stackpath/1.0.0/${OS_ARCH}
	-rm ~/terraform-providers/local/providers/stackpath/1.0.0/${OS_ARCH}/${BINARY}
	mv ${BINARY} ~/terraform-providers/local/providers/stackpath/1.0.0/${OS_ARCH}
	@echo

test:
	@echo "==> Running tests..."
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
	@echo

testacc:
	@echo "==> Running acceptance tests..."
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
	@echo

vet:
	@echo "==> Checking for suspicious Go constructs..."
	@echo "go vet \$$(go list ./... | grep -v vendor/)"
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -ne 0 ]; then \
		echo; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		echo; \
		exit 1; \
	fi
	@echo

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w $(GOFMT_FILES)
	@echo

fmtcheck:
	@echo "==> Checking that code complies with gofmt requirements..."
	gofmt_files=$$(gofmt -l $$(find . -name '*.go' | grep -v vendor))
	@if [[ -n "$(gofmt_files)" ]]; then \
  		echo; \
		echo "gofmt needs running on the following files:"; \
		echo "$(gofmt_files)"; \
		echo "You can use the command: \`make fmt\` to reformat code."; \
		echo; \
		exit 1; \
	fi;
	@echo

test-compile:
	@echo "==> Compiling test binary..."
	@if [ "$(TEST)" = "./..." ]; then \
  		echo; \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		echo; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)
	@echo

generate:
	@which swagger ; if [ $$? -ne 0 ] ; then \
		echo "Please install go-swagger to generate StackPath API client code"; \
		echo "See: https://goswagger.io/install.html"; \
		echo; \
		exit 1; \
	fi

	@echo "==> Generating code from StackPath API swagger specs..."
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

	@echo "==> Patching generated code..."
	sed -i '' -e 's/^type V1ProtocolAh interface{}$$/type V1ProtocolAh struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_ah.go
	sed -i '' -e 's/^type V1ProtocolEsp interface{}$$/type V1ProtocolEsp struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_esp.go
	sed -i '' -e 's/^type V1ProtocolGre interface{}$$/type V1ProtocolGre struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_gre.go
	sed -i '' -e 's/^type V1ProtocolIcmp interface{}$$/type V1ProtocolIcmp struct{}/' ./stackpath/api/ipam/ipam_models/v1_protocol_icmp.go

	sed -i '' -e 's/^	Ah V1ProtocolAh `json:"ah,omitempty"`$$/	Ah *V1ProtocolAh `json:"ah,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
	sed -i '' -e 's/^	Esp V1ProtocolEsp `json:"esp,omitempty"`$$/	Esp *V1ProtocolEsp `json:"esp,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
	sed -i '' -e 's/^	Gre V1ProtocolGre `json:"gre,omitempty"`$$/	Gre *V1ProtocolGre `json:"gre,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
	sed -i '' -e 's/^	Icmp V1ProtocolIcmp `json:"icmp,omitempty"`$$/	Icmp *V1ProtocolIcmp `json:"icmp,omitempty"`/' ./stackpath/api/ipam/ipam_models/v1_protocols.go
	@echo

.PHONY: build install test testacc vet fmt fmtcheck errcheck test-compile
