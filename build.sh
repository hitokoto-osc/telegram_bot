#!/usr/bin/env bash
set -e

#
# 来源： https://github.com/wolfeidau/somproject
#
OWNER=hitokoto
BIN_NAME=worker
PROJECT_NAME=notification_worker

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

GIT_COMMIT="$(git rev-parse HEAD)"
GIT_DIRTY="$(test -n "$(git status --porcelain)" && echo "+CHANGES" || true)"
VERSION=$(grep "const Version " src/build/info.go | sed -E 's/.*"(.+)"$/\1/' )

go get -d -v ./...

# building the master branch on ci
if [ "$BUILDBOX_BRANCH" = "master" ]; then
	go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -tags release -o ./bin/${BIN_NAME}_"${VERSION}"_linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -tags release -o ./bin/${BIN_NAME}_"${VERSION}"_linux_arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -tags release -o ./bin/${BIN_NAME}_"${VERSION}"_windows_amd64.exe
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -tags release -o ./bin/${BIN_NAME}_"${VERSION}"_darwin_amd64

else
	go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o ./bin/${BIN_NAME}_"${VERSION}"_linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o ./bin/${BIN_NAME}_"${VERSION}"_linux_arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o ./bin/${BIN_NAME}_"${VERSION}"_windows_amd64.exe
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X build.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o ./bin/${BIN_NAME}_"${VERSION}"_darwin_amd64
fi
