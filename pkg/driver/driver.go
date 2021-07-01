/*
Copyright 2021 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/utils/mount"
)

type Driver struct {
	// driver name to be registered at CSI
	name string

	// pluginType defines the type of plugin
	// node plugin or controller plugin
	pluginType string

	// csi driver version
	version string

	// endPoint represent the address on which
	// requests are made by kubelet/external-provisioner
	endPoint string

	// nodeID helps in differentiating the nodes on
	// which node drivers are running. This is useful
	// in case of topologies and publishing or
	// unpublishing volumes on nodes
	nodeID string

	ids csi.IdentityServer
	ns  csi.NodeServer
	cs  csi.ControllerServer

	cap []*csi.VolumeCapability_AccessMode
}

type pluginType int

const (
	// Driver name
	DriverName = "nfs.csi.openebs.io"

	// TopologyNodenameKey is supported topology key for the nfs driver
	TopologyNodenameKey = "nfs.openebs.io/nodename"

	// volume paramteres for node server
	NodeParamServer = "server"
	NodeParamPath   = "path"

	// plugin type
	ControllerPlugin pluginType = iota
	NodePlugin
)

// New returns a new driver instance
func New(nodeid, endpoint string, plugin pluginType) *Driver {
	driver := &Driver{}

	switch plugin {
	case NodePlugin:
		driver.ns = NewNodeServer(driver, mount.New(""))
	}

	// Identity server is common to both node and
	// controller, it is required to register,
	// share capabilities and probe the corresponding
	// driver
	driver.ids = NewIdentityServer(driver)
	return driver
}
