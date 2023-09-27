// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tailscale-native:tailnet:Acl":
		r = &Acl{}
	case "tailscale-native:tailnet:DNSPreferences":
		r = &DNSPreferences{}
	case "tailscale-native:tailnet:Key":
		r = &Key{}
	case "tailscale-native:tailnet:NameServers":
		r = &NameServers{}
	case "tailscale-native:tailnet:ReplaceSearchPaths":
		r = &ReplaceSearchPaths{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tailscale-native",
		"tailnet",
		&module{version},
	)
}
