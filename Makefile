appVersion ?= 0.0.0
appBuildTime ?= $(TZ=Asia/Shanghai date "+%F %T GMT%:z")
appGitCommitHash ?= $(git rev-parse HEAD)
appModuleName = github.com/liuguangw/forumx
buildLdFlags = "-X $(appModuleName)/cmd.appVersion=$(appVersion)"
buildLdFlags += " -X '$(appModuleName)/cmd.appBuildTime=$(appBuildTime)'"
buildLdFlags += " -X $(appModuleName)/cmd.appGitCommitHash=$(appGitCommitHash)"
EXTRA_FILES = LICENSE README.md

define build_app

export GOOS=$(1)
export GOARCH=$(2)
echo build for ($(GOOS) , $(GOARCH))
projectName=forumx
outputFileName=$(projectName)$(3)
zipFileName=release-$(GOOS)-$(GOARCH).zip
echo go build -v -ldflags "$(buildLdFlags)" -o $(outputFileName)
go build -v -ldflags "$(buildLdFlags)" -o $(outputFileName)
zip -r $(zipFileName) $(outputFileName) $(EXTRA_FILES)
rm $(outputFileName)
ls -al
endef

build:
	outputFileName=forumx
	echo go build -v -ldflags "$(buildLdFlags)" -o $(outputFileName)
	go build -v -ldflags "$(buildLdFlags)" -o $(outputFileName)
	ls -al

build_all:
	$(call build_app,linux,amd64)
	$(call build_app,linux,arm64)
	$(call build_app,darwin,amd64)
	$(call build_app,darwin,arm64)
	$(call build_app,windows,amd64,.exe)

clean:
	rm *.zip

.PHONY: build build_all clean
