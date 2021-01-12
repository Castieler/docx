
ETCD_NAMESPACE="platform"
ETCD_NS_PROJ="docx"
ETCD_ENVS=ETCD_PROJ=$(ETCD_PROJ) ETCD_NAMESPACE=$(ETCD_NAMESPACE) ETCD_NS_PROJ=$(ETCD_NS_PROJ)

GOPROXY=GOPROXY=https://internal-gomod.milvzn.com GONOSUMDB='*'
GOCMD=go
GOBUILD=$(GOPROXY) $(GOCMD) build
GOLINT=golangci-lint run

## deps: install all the deps the project needs
deps:
	@# $(WIRE_CMD)
	@cp .hooks/* .git/hooks/
	@$(GOBUILD) -i ./...

## lint: run golint in all packages
lint:
	@$(GOLINT) ./...
