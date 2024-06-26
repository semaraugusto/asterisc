GITCOMMIT ?= $(shell git rev-parse HEAD)
GITDATE ?= $(shell git show -s --format='%ct')
VERSION := v0.0.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X github.com/ethereum-optimism/asterisc/rvgo/version.Version=$(VERSION)
LDFLAGSSTRING +=-X github.com/ethereum-optimism/asterisc/rvgo/version.Meta=$(VERSION_META)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

build:
	env GO111MODULE=on GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -v $(LDFLAGS) -o ./bin/asterisc .
.PHONY: build

build-ffi:
	env GO111MODULE=on GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -v $(LDFLAGS) -o ./scripts/go-ffi/go-ffi ./scripts/go-ffi
.PHONY: build-ffi

clean:
	rm -rf ./bin
	rm ./scripts/go-ffi/go-ffi
.PHONY: clean

test:
	go test ./...
.PHONY: test

lint-check:
	golangci-lint run
.PHONY: lint-check

lint-fix:
	golangci-lint run --fix
.PHONY: lint-fix

build-bindgen:
	cd ../rvsol/lib/optimism/op-bindings && \
	go build -o ./bin/op-bindings ./cmd/. && \
	cp ./bin/op-bindings ../../../../rvgo/
.PHONY: build-bindgen

bindgen-generate-local:
	./op-bindings \
		generate \
		--metadata-out ./bindings \
		--bindings-package bindings \
		--contracts-list ../rvsol/artifacts.json \
		--log.level info \
		local \
		--forge-artifacts ../rvsol/out
.PHONY: bindgen-generate-local
