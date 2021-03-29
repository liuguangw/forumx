appVersion ?= 0.0.0
appBuildTime ?= $(shell TZ=Asia/Shanghai date "+%F %T GMT%:z")
appGitCommitHash ?= $(shell git rev-parse HEAD)
projectName=forumx
appModuleName = github.com/liuguangw/$(projectName)
buildLdFlags =-X $(appModuleName)/cmd.appVersion=$(appVersion)
buildLdFlags += -X '$(appModuleName)/cmd.appBuildTime=$(appBuildTime)'
buildLdFlags += -X $(appModuleName)/cmd.appGitCommitHash=$(appGitCommitHash)
GO_BUILD=go build -v -ldflags "$(buildLdFlags)"
EXTRA_FILES = LICENSE README.md

define build_app

$(eval GOOS:=$(1))
$(eval GOARCH:=$(2))
@echo build for \($(GOOS) , $(GOARCH)\)
$(eval outputFileName:=$(projectName)$(3))
$(eval zipFileName:=release-$(GOOS)-$(GOARCH).zip)
@echo $(GO_BUILD) -o ${outputFileName}
$(GO_BUILD) -o ${outputFileName}
zip -r ${zipFileName} ${outputFileName} $(EXTRA_FILES)
rm ${outputFileName}
ls -al
endef

build:
	echo $(GO_BUILD) -o $(projectName)
	$(GO_BUILD) -o $(projectName)
	ls -al

all:
	$(call build_app,linux,amd64)
	$(call build_app,linux,arm64)
	$(call build_app,darwin,amd64)
	$(call build_app,darwin,arm64)
	$(call build_app,windows,amd64,.exe)

clean:
	rm -rf ./forumx
	rm -rf ./*.zip

.PHONY: build build_all clean
