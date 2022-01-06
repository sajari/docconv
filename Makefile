#!/usr/bin/make -f

IMAGE=docker.io/skpr/docconv

define buildimage
  docker build --build-arg alpine_version=$(1) -t $(IMAGE):alpine$(1) .
endef

define pushimage
	docker push $(IMAGE):alpine$(1)
endef

build:
	$(call buildimage,3.13)
	$(call buildimage,3.14)
	$(call buildimage,3.15)

push:
	$(call pushimage,3.13)
	$(call pushimage,3.14)
	$(call pushimage,3.15)
