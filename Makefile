GO_FLAGS = GO111MODULE=on CGO_ENABLED=0

CURRENT_VERSION = $(shell grep 'const VERSION string' cmd/duration/duration.go | sed 's/.*"\(.*\)"/\1/')
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
GIT_BRANCH_UP_TO_DATE = $(shell git remote show origin | tail -n1 | sed 's/.*(\(.*\))/\1/')

default: build
	
build:
	 $(GO_FLAGS) go build -o ./bin/duration ./...

run: build
	./bin/duration ./test-script/script.sh

test: 
	go test -cover ./...

clean:
	rm -f ./bin/duration

release: _is_master_branch _git_branch_is_up_to_date current_version _tag_version _do_release
	@echo Release done. Go to Github and create a release.
	@echo Remember to push a newer docker image

_is_master_branch:
ifneq ($(GIT_BRANCH),master)
	@echo You are not on the master branch.
	@echo Please check out the master and try to release again
	@false
endif

_git_branch_is_up_to_date:
ifneq ($(GIT_BRANCH_UP_TO_DATE),up to date)
	@echo Your master branch is not up to date.
	@echo Please push your changes or pull changes from the remote.
	@false
endif

current_version:
	@echo the current version is: $(CURRENT_VERSION)

_do_release: clean test build run _build-all-binaries _compress-all-binaries

_tag_version:
	@read -p "Enter version to release: " version && \
	sed -i "s/const VERSION string = \".*\"/const VERSION string = \"$${version}\"/" ./cmd/duration/duration.go && \
	git add . && git commit -m "chore(release): $${version}" && git tag "$${version}" && \
	git push origin master && git push origin master --tags

_build-all-binaries:
	# doesn't work on my machine and not in travis, see: https://github.com/golang/go/wiki/GoArm
	# GOOS=android GOARCH=arm  $(COMPILE_COMMAND) && mv ./bin/ec ./bin/ec-android-arm
	# GOOS=darwin  GOARCH=arm $(COMPILE_COMMAND) && mv ./bin/ec ./bin/ec-darwin-arm
	# GOOS=darwin  GOARCH=arm64 $(COMPILE_COMMAND) && mv ./bin/ec ./bin/ec-darwin-arm64
	$(GO_FLAGS) GOOS=darwin    GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-darwin-386
	$(GO_FLAGS) GOOS=darwin    GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-darwin-amd64
	$(GO_FLAGS) GOOS=dragonfly GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-dragonfly-amd64
	$(GO_FLAGS) GOOS=freebsd   GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-freebsd-386
	$(GO_FLAGS) GOOS=freebsd   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-freebsd-amd64
	$(GO_FLAGS) GOOS=freebsd   GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-freebsd-arm
	$(GO_FLAGS) GOOS=linux     GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-386
	$(GO_FLAGS) GOOS=linux     GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-amd64
	$(GO_FLAGS) GOOS=linux     GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-arm
	$(GO_FLAGS) GOOS=linux     GOARCH=arm64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-arm64
	$(GO_FLAGS) GOOS=linux     GOARCH=ppc64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-ppc64
	$(GO_FLAGS) GOOS=linux     GOARCH=ppc64le  $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-ppc64le
	$(GO_FLAGS) GOOS=linux     GOARCH=mips     $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-mips
	$(GO_FLAGS) GOOS=linux     GOARCH=mipsle   $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-mipsle
	$(GO_FLAGS) GOOS=linux     GOARCH=mips64   $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-mips64
	$(GO_FLAGS) GOOS=linux     GOARCH=mips64le $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-linux-mips64le
	$(GO_FLAGS) GOOS=netbsd    GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-netbsd-386
	$(GO_FLAGS) GOOS=netbsd    GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-netbsd-amd64
	$(GO_FLAGS) GOOS=netbsd    GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-netbsd-arm
	$(GO_FLAGS) GOOS=openbsd   GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-openbsd-386
	$(GO_FLAGS) GOOS=openbsd   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-openbsd-amd64
	$(GO_FLAGS) GOOS=openbsd   GOARCH=arm      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-openbsd-arm
	$(GO_FLAGS) GOOS=plan9     GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-plan9-386
	$(GO_FLAGS) GOOS=plan9     GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-plan9-amd64
	$(GO_FLAGS) GOOS=solaris   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-solaris-amd64
	$(GO_FLAGS) GOOS=windows   GOARCH=386      $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-windows-386.exe
	$(GO_FLAGS) GOOS=windows   GOARCH=amd64    $(COMPILE_COMMAND) && mv ./bin/duration ./bin/duration-windows-amd64.exe

_compress-all-binaries:
	for f in $(BINARIES); do      \
		tar czf $$f.tar.gz $$f;    \
	done
	@rm $(BINARIES)
