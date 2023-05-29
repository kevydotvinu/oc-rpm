## oc-quicklab
[![licence](https://img.shields.io/github/license/kevydotvinu/oc-rpm)](https://github.com/kevydotvinu/oc-rpm/blob/main/LICENSE)
[![goversion](https://img.shields.io/github/go-mod/go-version/openshift/oc)](https://github.com/openshift/oc/blob/master/go.mod)
[![downloads](https://img.shields.io/github/downloads/kevydotvinu/oc-rpm/total)](https://github.com/kevydotvinu/oc-rpm/releases)
[![release](https://github.com/kevydotvinu/oc-rpm/actions/workflows/build-and-release.yaml/badge.svg)](https://github.com/kevydotvinu/oc-rpm/actions/workflows/build-and-release.yaml)
[![openshiftplugin](https://img.shields.io/badge/openshift%20cli-plug--in-orange)](https://docs.openshift.com/container-platform/latest/cli_reference/openshift_cli/extending-cli-plugins.html)

### OpenShift CLI plug-in to list RPMs
The oc-rpm is a CLI plug-in that works with oc/kubectl and gives the list of RPM packages in an OpenShift release.

### Prerequisites
- [OpenShift CLI](https://mirror.openshift.com/pub/openshift-v4/clients/ocp/)

### Installation
```bash
$ curl -#Lo oc-rpm $(curl -s https://api.github.com/repos/kevydotvinu/oc-rpm/releases/latest | jq -r '.assets | .[] | select(.name | contains("linux")) | .browser_download_url')
$ sudo mv oc-rpm /usr/local/bin/ && sudo chmod +x /usr/local/bin/oc-rpm
$ oc rpm 4.13.0
```

### Usage
##### List RPMs in a local cluster
```bash
$ oc rpm
```
##### List RPMs in an OpenShift release
```bash
$ oc rpm 4.13.0
```
