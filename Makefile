.PHONY: all
all: build

.PHONY: build
build: build-act build-goreleaser

release: build-goreleaser
	./bin/goreleaser release --rm-dist

.PHONY: build-act
build-act: ./bin/act
	./bin/act -j build
.PHONY: build-goreleaser
build-goreleaser: ./bin/goreleaser
	./bin/goreleaser release --snapshot --rm-dist

./bin/act:
	GOBIN=${PWD}/bin go install github.com/nektos/act@latest

./bin/goreleaser:
	GOBIN=${PWD}/bin go install github.com/goreleaser/goreleaser@latest

