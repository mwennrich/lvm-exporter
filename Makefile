GO111MODULE := on
DOCKER_TAG := $(or ${GIT_TAG_NAME}, latest)

all: lvm-exporter

.PHONY: lvm-exporter
lvm-exporter:
	go build -tags netgo -o bin/lvm-exporter *.go
	strip bin/lvm-exporter

.PHONY: fix
fix:
	go fix ./...

.PHONY: dockerimages
dockerimages:
	docker build -t mwennrich/lvm-exporter:${DOCKER_TAG} .

.PHONY: dockerpush
dockerpush:
	docker push mwennrich/lvm-exporter:${DOCKER_TAG}

.PHONY: clean
clean:
	rm -f bin/*
