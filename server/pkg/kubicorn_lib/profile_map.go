package kubicornlib

import (
	"github.com/kris-nova/kubicorn/apis/cluster"
	"github.com/kris-nova/kubicorn/profiles/amazon"
	"github.com/kris-nova/kubicorn/profiles/azure"
	"github.com/kris-nova/kubicorn/profiles/digitalocean"
	"github.com/kris-nova/kubicorn/profiles/googlecompute"
	"github.com/kris-nova/kubicorn/profiles/packet"
)

type profileFunc func(name string) *cluster.Cluster

type profileMap struct {
	ProfileFunc profileFunc
	Description string
}

var ProfileMapIndexed = map[string]profileMap{
	"azure": {
		ProfileFunc: azure.NewUbuntuCluster,
		Description: "Ubuntu on Azure",
	},
	"azure-ubuntu": {
		ProfileFunc: azure.NewUbuntuCluster,
		Description: "Ubuntu on Azure",
	},
	"amazon": {
		ProfileFunc: amazon.NewUbuntuCluster,
		Description: "Ubuntu on Amazon",
	},
	"aws": {
		ProfileFunc: amazon.NewUbuntuCluster,
		Description: "Ubuntu on Amazon",
	},
	"do": {
		ProfileFunc: digitalocean.NewUbuntuCluster,
		Description: "Ubuntu on DigitalOcean",
	},
	"google": {
		ProfileFunc: googlecompute.NewUbuntuCluster,
		Description: "Ubuntu on Google Compute",
	},
	"digitalocean": {
		ProfileFunc: digitalocean.NewUbuntuCluster,
		Description: "Ubuntu on DigitalOcean",
	},
	"do-ubuntu": {
		ProfileFunc: digitalocean.NewUbuntuCluster,
		Description: "Ubuntu on DigitalOcean",
	},
	"aws-ubuntu": {
		ProfileFunc: amazon.NewUbuntuCluster,
		Description: "Ubuntu on Amazon",
	},
	"do-centos": {
		ProfileFunc: digitalocean.NewCentosCluster,
		Description: "CentOS on DigitalOcean",
	},
	"aws-centos": {
		ProfileFunc: amazon.NewCentosCluster,
		Description: "CentOS on Amazon",
	},
	"aws-debian": {
		ProfileFunc: amazon.NewDebianCluster,
		Description: "Debian on Amazon",
	},
	"packet": {
		ProfileFunc: packet.NewUbuntuCluster,
		Description: "Ubuntu on Packet x86",
	},
	"packet-ubuntu": {
		ProfileFunc: packet.NewUbuntuCluster,
		Description: "Ubuntu on Packet x86",
	},
}
