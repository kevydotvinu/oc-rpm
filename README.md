## oc-rpm
[![licence](https://img.shields.io/github/license/kevydotvinu/oc-rpm)](https://github.com/kevydotvinu/oc-rpm/blob/main/LICENSE)
[![goversion](https://img.shields.io/github/go-mod/go-version/openshift/oc)](https://github.com/openshift/oc/blob/master/go.mod)
[![downloads](https://img.shields.io/github/downloads/kevydotvinu/oc-rpm/total)](https://github.com/kevydotvinu/oc-rpm/releases)
[![release](https://github.com/kevydotvinu/oc-rpm/actions/workflows/build-and-release.yaml/badge.svg)](https://github.com/kevydotvinu/oc-rpm/actions/workflows/build-and-release.yaml)
[![openshiftplugin](https://img.shields.io/badge/openshift%20cli-plug--in-orange)](https://docs.openshift.com/container-platform/latest/cli_reference/openshift_cli/extending-cli-plugins.html)

### OpenShift CLI plug-in to list release RPMs
The oc-rpm is a CLI plug-in that works with oc/kubectl and gives the list of RPM packages in an OpenShift release.

### Prerequisites
- [OpenShift command-line interface (oc)](https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-linux.tar.gz)
- [Pull secret](https://console.redhat.com/openshift/downloads#tool-pull-secret)

### Installation
```bash
curl -#Lo oc-rpm $(curl -s https://api.github.com/repos/kevydotvinu/oc-rpm/releases/latest | jq -r '.assets | .[] | select(.name | contains("linux")) | .browser_download_url')
sudo mv oc-rpm /usr/local/bin/ && sudo chmod +x /usr/local/bin/oc-rpm
```

### Usage
##### List RPMs in a local cluster
```bash
export REGISTRY_AUTH_FILE=<pullsecret path>
oc rpm
```
##### List RPMs in an OpenShift release
```bash
export REGISTRY_AUTH_FILE=<pullsecret path>
oc rpm 4.13.0
```
> Default path of the authentication file is `${XDG_RUNTIME_DIR}/containers/auth.json` on Linux, and `$HOME/.config/containers/auth.json` on Windows/macOS. The file is created by podman login. If the authorization state is not found there, `$HOME/.docker/config.json` is checked, which is set using docker login.

> There is also the option to override the default path of the authentication file by setting the `REGISTRY_AUTH_FILE` environment variable. This can be done with `export REGISTRY_AUTH_FILE=path`.
