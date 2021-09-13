# Dynamic NFS Volume Provisioner
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fopenebs%2Fdynamic-nfs-provisioner.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fopenebs%2Fdynamic-nfs-provisioner?ref=badge_shield)


[![Build Status](https://github.com/openebs/dynamic-nfs-provisioner/actions/workflows/build.yml/badge.svg)](https://github.com/openebs/dynamic-nfs-provisioner/actions/workflows/build.yml)
[![Go Report](https://goreportcard.com/badge/github.com/openebs/dynamic-nfs-provisioner)](https://goreportcard.com/report/github.com/openebs/dynamic-nfs-provisioner)
[![codecov](https://codecov.io/gh/openebs/dynamic-nfs-provisioner/branch/develop/graph/badge.svg)](https://app.codecov.io/gh/openebs/dynamic-nfs-provisioner)
[![Slack](https://img.shields.io/badge/chat!!!-slack-ff1493.svg?style=flat-square)](https://kubernetes.slack.com/messages/openebs)
[![BCH compliance](https://bettercodehub.com/edge/badge/openebs/dynamic-nfs-provisioner?branch=develop)](https://bettercodehub.com/results/openebs/dynamic-nfs-provisioner)

<img width="300" align="right" alt="OpenEBS Logo" src="https://raw.githubusercontent.com/cncf/artwork/HEAD/projects/openebs/stacked/color/openebs-stacked-color.png" xmlns="http://www.w3.org/1999/html">

<p align="justify">
<strong>OpenEBS Dynamic NFS PV provisioner</strong> can be used to dynamically provision 
NFS Volumes using different kinds of block storage available on the Kubernetes nodes. 
<br>
<br>
</p>

This project is under active development. 

## Prerequisites ##


Please ensure that an NFS client is functioning on all nodes that will run a pod that mounts an `openebs-rwx` volume. 

Here's how to prepare an NFS client on some common Operating Systems:

- **Ubuntu, Debian:** Install the `nfs-common` package if not already installed.

- **MacOS:** Should work out of the box.

- **Windows:**
Ensure that the default NFS client is operating. To do this start PowerShell as Administrator, and run `Install-WindowsFeature NFS-Client` if it's a Windows server or `Enable-WindowsOptionalFeature -FeatureName ServicesForNFS-ClientOnly, ClientForNFS-Infrastructure -Online -NoRestart` if it's a Windows host with a Desktop environment.

- **Fedora, CentOS, RedHat**: Install the `nfs-utils` package if not already installed.

- **FreeBSD**: 
   1) Edit the /etc/rc.conf file by setting or appending `nfs_client_enable="YES"`
   2) Run `service nfsclient start`

## Install
### Install NFS Provisioner through kubectl
```
kubectl apply -f https://openebs.github.io/charts/nfs-operator.yaml
```

Create a StorageClass with required backing storage class. Example:
```
#Sample storage classes for OpenEBS Local PV
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: openebs-rwx
  annotations:
    openebs.io/cas-type: nfsrwx
    cas.openebs.io/config: |
      - name: NFSServerType
        value: "kernel"
      - name: BackendStorageClass
        value: "openebs-hostpath"
      #  LeaseTime defines the renewal period(in seconds) for client state
      #- name: LeaseTime
      #  value: 30
      #  GraceTime defines the recovery period(in seconds) to reclaim locks
      #- name: GraceTime
      #  value: 30
provisioner: openebs.io/nfsrwx
reclaimPolicy: Delete
```

You can now use `openebs-rwx` storage class to create RWX volumes.

### Install NFS Provisioner through Helm
```
helm repo add openebs-nfs https://openebs.github.io/dynamic-nfs-provisioner
helm install [RELEASE_NAME] openebs-nfs/nfs-provisioner --namespace openebs --create-namespace
```

Refer https://github.com/openebs/dynamic-nfs-provisioner/tree/develop/deploy/helm/charts for the list of configuration parameter of the dynamic-nfs-provisioner chart.

## Contributing

OpenEBS welcomes your feedback and contributions in any form possible.

- [Join OpenEBS community on Kubernetes Slack](https://kubernetes.slack.com)
  - Already signed up? Head to our discussions at [#openebs](https://kubernetes.slack.com/messages/openebs/)
- Want to raise an issue or help with fixes and features?
  - See [open issues](https://github.com/openebs/openebs/issues)
  - See [contributing guide](./CONTRIBUTING.md)
  - See [Project Roadmap](https://github.com/orgs/openebs/projects/12)
- Join our OpenEBS CNCF Mailing lists
  - For OpenEBS project updates, subscribe to [OpenEBS Announcements](https://lists.cncf.io/g/cncf-openebs-announcements)
  - For interacting with other OpenEBS users, subscribe to [OpenEBS Users](https://lists.cncf.io/g/cncf-openebs-users)

## Community, discussion, and support

Learn how to engage with the OpenEBS community on the [community page](https://github.com/openebs/openebs/tree/HEAD/community).

You can reach the maintainers of this project at:

- [Kubernetes Slack](http://slack.k8s.io/) channels: 
      * [#openebs](https://kubernetes.slack.com/messages/openebs/)
      * [#openebs-dev](https://kubernetes.slack.com/messages/openebs-dev/)
- [Mailing List](https://lists.cncf.io/g/cncf-openebs-users)

### Code of conduct

Participation in the OpenEBS community is governed by the [CNCF Code of Conduct](CODE-OF-CONDUCT.md).

## Inspiration/Credit
- https://github.com/sjiveson/nfs-server-alpine


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fopenebs%2Fdynamic-nfs-provisioner.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fopenebs%2Fdynamic-nfs-provisioner?ref=badge_large)