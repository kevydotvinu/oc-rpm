SHELL=/bin/bash
DIR=$(shell pwd)
ORANGE=\e[0;33m
NOCOLOR=\e[0m

.PHONY: update-repo

update-repo:
	@echo -e "${ORANGE}Updating submodule ...${NOCOLOR}"
	git submodule update --remote
	@echo -e "${ORANGE}Removing linked & copied files ...${NOCOLOR}"
	rm -rvf cmd \
		go.mod \
		go.sum \
		pkg \
		vendor
	@echo -e "${ORANGE}Linking and copying files from oc submodule ...${NOCOLOR}"
	ln -svf ./oc/cmd . && \
		ln -svf ./oc/go.mod . && \
		ln -svf ./oc/go.sum . && \
		ln -svf ./oc/vendor .
	mkdir pkg && \
		pushd pkg && \
		ln -svf ../oc/pkg/* . && \
		rm -rvf cli && \
		mkdir cli && \
		pushd cli && \
		ln -svf ../../oc/pkg/cli/* . && \
		rm -rvf admin && \
		mkdir admin && \
		pushd admin && \
		ln -svf ../../../oc/pkg/cli/admin/* . && \
		rm -rvf release && \
		mkdir release && \
		pushd release && \
		ln -svf ../../../../oc/pkg/cli/admin/release/* . && \
		rm -rvf release.go && \
		cp -rvf ../../../../oc/pkg/cli/admin/release/release.go . && \
		cp -rvf ../../../../package.go . && \
		popd && \
		popd && \
		popd && \
		popd
	@echo -e "${ORANGE}Patching package.go with cmd.AddCommand(NewPackage(f, streams)) ...${NOCOLOR}"
	sed -i '/return cmd/i cmd.AddCommand(NewPackage(f, streams))' pkg/cli/admin/release/release.go
	go fmt pkg/cli/admin/release/release.go

.PHONY: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o oc-rpm_linux_amd64 -ldflags="-s -w" main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o oc-rpm_darwin_amd64 -ldflags="-s -w" main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o oc-rpm_windows_amd64 -ldflags="-s -w" main.go
