PROJECT_NAME := "telegram_bot"
PROJECT_PATH := "github.com/hitokoto-osc/telegram_bot"
PKG := "$(PROJECT_PATH)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_COMMIT_TIME := $(shell git log --pretty=format:"%cd" ${GIT_COMMIT} -1)
GIT_DIRTY :=$(shell test -n "$(git status --porcelain)" && echo "+CHANGES" || true)

.PHONY: all dep get-tools lint vet test test-coverage build clean

all:
	build

get-tools:
	@echo Installing tools...
	go install github.com/mgechev/revive

dep: # get dependencies
	@echo Installing Dependencies...
	go mod download

lint: get-tools ## Lint Golang files
	@echo
	@echo Linting go codes with revive...
	@revive -config ./.revive.toml -formatter stylish ${PKG_LIST}

vet:
	@echo Linting go codes with go vet...
	go vet ./...

precommit: vet lint test
	go mod tidy

build: dep
	@echo;
	@echo Building...;
	@mkdir -p dist;
	@go build -ldflags "-X 'github.com/hitokoto-osc/telegram_bot/build.CommitTag=${GIT_COMMIT}' -X 'github.com/hitokoto-osc/telegram_bot/build.CommitTime=${GIT_COMMIT_TIME}'" -v -o dist/${PROJECT_NAME} .;

test:
	@echo Testing...
	@go test -short ${PKG_LIST}

test-coverage:
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

clean:
	@rm -f coverage.txt
	@rm -f cover.out

release:
	@echo Releasing by GoReleaser...
	@goreleaser release --rm-dist
