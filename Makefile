# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Tools for deploy
KUBECTL?=kubectl
PWD=$(shell pwd)

IMAGE ?= quay.io/open-cluster-management/multicluster-mesh-addon:latest

all: build
.PHONY: all

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: golangci-lint ## Run golangci-lint against code.
	$(GOLANGCI_LINT) run

.PHONY: test
test: fmt vet ## Run tests.
	go test ./... -coverprofile cover.out

.PHONY: test-integration
test-integration:
	@echo "TODO: Run integration test"

.PHONY: test-e2e
test-e2e:
	@echo "TODO: Run e2e test"

##@ Build

.PHONY: build
build: fmt vet ## Build multicluster-mesh-addon binary.
	go build -o bin/multicluster-mesh-addon main.go

.PHONY: docker-build
docker-build: test ## Build docker image with the multicluster-mesh-addon.
	docker build ${IMAGE_BUILD_EXTRA_FLAGS} -t ${IMAGE} .

.PHONY: docker-push
docker-push: ## Push docker image with the multicluster-mesh-addon.
	docker push ${IMAGE}

##@ Deployment

.PHONY: manifests
manifests: controller-gen ## Generate CustomResourceDefinition objects.
	$(CONTROLLER_GEN) crd paths="./apis/..." output:crd:artifacts:config=deploy/crd

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./apis/..."

.PHONY: update-codegen
update-codegen: # update the client set tools.
	go mod vendor
	./hack/update-codegen.sh
	rm -rf vendor

.PHONY: verify-codegen
verify-codegen: # verify the client set tools.
	go mod vendor
	./hack/verify-codegen.sh
	rm -rf vendor

deploy: kustomize
	cp deploy/kustomization.yaml deploy/kustomization.yaml.tmp
	cd deploy && $(KUSTOMIZE) edit set image multicluster-mesh-addon=$(IMAGE) && $(KUSTOMIZE) edit add configmap image-config --from-literal=MULTICLUSTER_MESH_ADDON_IMAGE=$(IMAGE)
	$(KUSTOMIZE) build deploy | $(KUBECTL) apply -f -
	mv deploy/kustomization.yaml.tmp deploy/kustomization.yaml

undeploy: kustomize
	$(KUSTOMIZE) build deploy | $(KUBECTL) delete --ignore-not-found -f -

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

GOLANGCI_LINT = $(LOCALBIN)/golangci-lint
KUSTOMIZE ?= $(LOCALBIN)/kustomize
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen

## Tool Versions
GOLANGCI_LINT_VERSION ?= v1.54.2
KUSTOMIZE_VERSION ?= v5.2.1
CONTROLLER_TOOLS_VERSION ?= v0.13.0


.PHONY: golangci-lint
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary. If wrong version is installed, it will be removed before downloading.
$(KUSTOMIZE): $(LOCALBIN)
	@if test -x $(LOCALBIN)/kustomize && ! $(LOCALBIN)/kustomize version | grep -q $(KUSTOMIZE_VERSION); then \
		echo "$(LOCALBIN)/kustomize version is not expected $(KUSTOMIZE_VERSION). Removing it before installing."; \
		rm -rf $(LOCALBIN)/kustomize; \
	fi
	test -s $(LOCALBIN)/kustomize || GOBIN=$(LOCALBIN) GO111MODULE=on go install sigs.k8s.io/kustomize/kustomize/v5@$(KUSTOMIZE_VERSION)

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary. If wrong version is installed, it will be overwritten.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen && $(LOCALBIN)/controller-gen --version | grep -q $(CONTROLLER_TOOLS_VERSION) || \
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

