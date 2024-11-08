.PHONY: all
all: clean build

# ==============================================================================
# Includes

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/docker.mk
include scripts/make-rules/test.mk


# ==============================================================================
# Golang

.PHONY: build
build:
	@echo "===========> Building Go bin for Linux amd64"
	@$(MAKE) go.build

.PHONY: push-pkg-tc
push-pkg-tc: build
	@echo "===========> Uploading Go bin and config files to Tencent Cloud COS"
#	@source ~/.tencentcloud/tmigrate.sh
	@./scripts/cos-upload.sh

.PHONY: run
run:
	@echo "===========> Running Go app"
	@$(MAKE) go.run

.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)


# ==============================================================================
# Docker & K8s

.PHONY: build-docker
build-docker:
	@echo "===========> Building Docker image"
	@$(MAKE) docker.build

.PHONY: push-docker
push-docker: build-docker
	@echo "===========> Pushing Docker image"
	@$(MAKE) docker.push

.PHONY: run-docker
run-docker: build-docker
	@echo "===========> Running Docker"
	@$(MAKE) docker.run

.PHONY: run-docker-without-build
run-docker-without-build:
	@echo "===========> Running Docker without build"
	@$(MAKE) docker.run

.PHONY: run-k8s
run-k8s: build-docker
	@echo "===========> Running Kubernetes"
	@$(MAKE) k8s.run

.PHONY: run-k8s-without-build
run-k8s-without-build:
	@echo "===========> Running Kubernetes without build"
	@$(MAKE) k8s.run

.PHONY: run-helm
run-helm: build-docker
	@echo "===========> Running Helm"
	@$(MAKE) helm.run

.PHONY: run-helm-without-build
run-helm-without-build:
	@echo "===========> Running Helm without build"
	@$(MAKE) helm.run

.PHONY: clean-k8s
clean-k8s:
	@echo "===========> Cleaning all local k8s"
	@$(MAKE) k8s.clean

.PHONY: clean-helm
clean-helm:
	@echo "===========> Cleaning Helm"
	@$(MAKE) helm.clean


# ==============================================================================
# Test

.PHONY: test-api
test-api:
	@$(MAKE) test.api

.PHONY: test
test:
	@scripts/cos-upload.sh


# ==============================================================================
# Swagger OpenAPI
.PHONY: swag
swag:
	@echo "===========> Generating Swagger OpenAPI"
	@swag init -g cmd/app1/main.go -o ./docs
