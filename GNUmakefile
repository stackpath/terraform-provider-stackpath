TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=stackpath
BINARY=terraform-provider-${PKG_NAME}
OS_ARCH=darwin_amd64

default: build

build: fmtcheck vet errcheck
	@echo "==> Building ${BINARY}..."
	go build -o ${BINARY}

install: build
	@echo "==> Installing ${BINARY}..."
	mkdir -p ~/.terraform.d/plugins/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${OS_ARCH}

test: fmtcheck vet errcheck
	@echo "==> Running tests..."
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck vet errcheck
	@echo "==> Running acceptance tests..."
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "==> Checking for suspicious Go constructs..."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: build install test testacc vet fmt fmtcheck errcheck test-compile
