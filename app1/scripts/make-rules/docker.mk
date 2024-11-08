# ==============================================================================
# Docker & K8s options

#REGISTRY_PREFIX ?= ruan-nj.tencentcloudcr.com/tamlab/repo
#REGISTRY_PREFIX ?= registry.rebirthmonkey.com/ops/ops
REGISTRY_PREFIX ?= wukongsun
IMAGE := app1
IMAGE_VERSION := v1.0.0

NAMESPACE ?= $(IMAGE)

# ==============================================================================
# Targets

.PHONY: docker.build
docker.build:
	@echo "===========> Building Docker image $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)"
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@-docker image rm $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)
	$(eval BUILD_SUFFIX := $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION))
	@docker build --platform $(OS)/$(ARCH) -f ${ROOT_DIR}/build/Dockerfile -t $(BUILD_SUFFIX) .


.PHONY: docker.push
docker.push:
	@echo "===========> Pushing Docker image $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)"
	@docker push $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)

.PHONY: docker.run
docker.run:
	@echo "===========> Running Local Docker $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)"
	@docker run --rm -p 8889:8888 -v $(ROOT_DIR)/configs/config.yaml:/etc/app/config.yaml $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)


.PHONY: k8s.run
k8s.run:
	@echo "===========> Running Local k8s $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)"
	@kubectl -n $(NAMESPACE) create configmap vol-config --from-file=$(ROOT_DIR)/configs/config.yaml --dry-run=client -o yaml | kubectl apply -f -
	@kubectl -n $(NAMESPACE) apply -f $(ROOT_DIR)/manifests/k8s/local/deployment.yaml
	@kubectl -n $(NAMESPACE) apply -f $(ROOT_DIR)/manifests/k8s/local/svc.yaml

.PHONY: helm.run
helm.run:
	@echo "===========> Running Local Helm $(REGISTRY_PREFIX)/$(IMAGE):$(IMAGE_VERSION)"
	@helm -n $(NAMESPACE) install $(IMAGE) ./manifests/helm -f ./manifests/helm/values-gke.yaml

.PHONY: k8s.clean
k8s.clean:
	@echo "===========> Cleaning Local k8s $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION)"
	@kubectl -n $(NAMESPACE) delete -f $(ROOT_DIR)/manifests/k8s/local/svc.yaml
	@kubectl -n $(NAMESPACE) delete -f $(ROOT_DIR)/manifests/k8s/local/deployment.yaml
	@kubectl -n $(NAMESPACE) delete configmap vol-config

.PHONY: helm.clean
helm.clean:
	@echo "===========> Cleaning Helm $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION)"
	@helm -n $(NAMESPACE) uninstall $(IMAGE)