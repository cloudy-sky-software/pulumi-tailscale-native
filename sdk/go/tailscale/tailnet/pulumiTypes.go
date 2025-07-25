// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type AclType struct {
	Acls          []AclRule   `pulumi:"acls"`
	AutoApprovers interface{} `pulumi:"autoApprovers"`
	Groups        interface{} `pulumi:"groups"`
	Hosts         interface{} `pulumi:"hosts"`
	NodeAttrs     []NodeAttrs `pulumi:"nodeAttrs"`
	Ssh           []SshRule   `pulumi:"ssh"`
	TagOwners     interface{} `pulumi:"tagOwners"`
	Tests         string      `pulumi:"tests"`
}

type AclRule struct {
	// Tailscale ACL rules are "default deny".
	// So the only possible value for an ACL
	// rule is `accept`.
	Action AclRuleAction `pulumi:"action"`
	Ports  []string      `pulumi:"ports"`
	Users  []string      `pulumi:"users"`
}

// AclRuleInput is an input type that accepts AclRuleArgs and AclRuleOutput values.
// You can construct a concrete instance of `AclRuleInput` via:
//
//	AclRuleArgs{...}
type AclRuleInput interface {
	pulumi.Input

	ToAclRuleOutput() AclRuleOutput
	ToAclRuleOutputWithContext(context.Context) AclRuleOutput
}

type AclRuleArgs struct {
	// Tailscale ACL rules are "default deny".
	// So the only possible value for an ACL
	// rule is `accept`.
	Action AclRuleActionInput      `pulumi:"action"`
	Ports  pulumi.StringArrayInput `pulumi:"ports"`
	Users  pulumi.StringArrayInput `pulumi:"users"`
}

func (AclRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AclRule)(nil)).Elem()
}

func (i AclRuleArgs) ToAclRuleOutput() AclRuleOutput {
	return i.ToAclRuleOutputWithContext(context.Background())
}

func (i AclRuleArgs) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleOutput)
}

// AclRuleArrayInput is an input type that accepts AclRuleArray and AclRuleArrayOutput values.
// You can construct a concrete instance of `AclRuleArrayInput` via:
//
//	AclRuleArray{ AclRuleArgs{...} }
type AclRuleArrayInput interface {
	pulumi.Input

	ToAclRuleArrayOutput() AclRuleArrayOutput
	ToAclRuleArrayOutputWithContext(context.Context) AclRuleArrayOutput
}

type AclRuleArray []AclRuleInput

func (AclRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AclRule)(nil)).Elem()
}

func (i AclRuleArray) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return i.ToAclRuleArrayOutputWithContext(context.Background())
}

func (i AclRuleArray) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleArrayOutput)
}

type AclRuleOutput struct{ *pulumi.OutputState }

func (AclRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AclRule)(nil)).Elem()
}

func (o AclRuleOutput) ToAclRuleOutput() AclRuleOutput {
	return o
}

func (o AclRuleOutput) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return o
}

// Tailscale ACL rules are "default deny".
// So the only possible value for an ACL
// rule is `accept`.
func (o AclRuleOutput) Action() AclRuleActionOutput {
	return o.ApplyT(func(v AclRule) AclRuleAction { return v.Action }).(AclRuleActionOutput)
}

func (o AclRuleOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AclRule) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

func (o AclRuleOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AclRule) []string { return v.Users }).(pulumi.StringArrayOutput)
}

type AclRuleArrayOutput struct{ *pulumi.OutputState }

func (AclRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AclRule)(nil)).Elem()
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) Index(i pulumi.IntInput) AclRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AclRule {
		return vs[0].([]AclRule)[vs[1].(int)]
	}).(AclRuleOutput)
}

type AuthKey struct {
	Created *string `pulumi:"created"`
	Expires string  `pulumi:"expires"`
	Key     string  `pulumi:"key"`
}

type AuthKeyOutput struct{ *pulumi.OutputState }

func (AuthKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthKey)(nil)).Elem()
}

func (o AuthKeyOutput) ToAuthKeyOutput() AuthKeyOutput {
	return o
}

func (o AuthKeyOutput) ToAuthKeyOutputWithContext(ctx context.Context) AuthKeyOutput {
	return o
}

func (o AuthKeyOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthKey) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o AuthKeyOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v AuthKey) string { return v.Expires }).(pulumi.StringOutput)
}

func (o AuthKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v AuthKey) string { return v.Key }).(pulumi.StringOutput)
}

type AuthKeyArrayOutput struct{ *pulumi.OutputState }

func (AuthKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthKey)(nil)).Elem()
}

func (o AuthKeyArrayOutput) ToAuthKeyArrayOutput() AuthKeyArrayOutput {
	return o
}

func (o AuthKeyArrayOutput) ToAuthKeyArrayOutputWithContext(ctx context.Context) AuthKeyArrayOutput {
	return o
}

func (o AuthKeyArrayOutput) Index(i pulumi.IntInput) AuthKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthKey {
		return vs[0].([]AuthKey)[vs[1].(int)]
	}).(AuthKeyOutput)
}

type ClientConnectivity struct {
	ClientSupports        ClientSupports `pulumi:"clientSupports"`
	Derp                  string         `pulumi:"derp"`
	Endpoints             string         `pulumi:"endpoints"`
	Latency               interface{}    `pulumi:"latency"`
	MappingVariesByDestIP bool           `pulumi:"mappingVariesByDestIP"`
}

type ClientConnectivityOutput struct{ *pulumi.OutputState }

func (ClientConnectivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientConnectivity)(nil)).Elem()
}

func (o ClientConnectivityOutput) ToClientConnectivityOutput() ClientConnectivityOutput {
	return o
}

func (o ClientConnectivityOutput) ToClientConnectivityOutputWithContext(ctx context.Context) ClientConnectivityOutput {
	return o
}

func (o ClientConnectivityOutput) ClientSupports() ClientSupportsOutput {
	return o.ApplyT(func(v ClientConnectivity) ClientSupports { return v.ClientSupports }).(ClientSupportsOutput)
}

func (o ClientConnectivityOutput) Derp() pulumi.StringOutput {
	return o.ApplyT(func(v ClientConnectivity) string { return v.Derp }).(pulumi.StringOutput)
}

func (o ClientConnectivityOutput) Endpoints() pulumi.StringOutput {
	return o.ApplyT(func(v ClientConnectivity) string { return v.Endpoints }).(pulumi.StringOutput)
}

func (o ClientConnectivityOutput) Latency() pulumi.AnyOutput {
	return o.ApplyT(func(v ClientConnectivity) interface{} { return v.Latency }).(pulumi.AnyOutput)
}

func (o ClientConnectivityOutput) MappingVariesByDestIP() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientConnectivity) bool { return v.MappingVariesByDestIP }).(pulumi.BoolOutput)
}

type ClientSupports struct {
	HairPinning bool `pulumi:"hairPinning"`
	Ipv6        bool `pulumi:"ipv6"`
	Pcp         bool `pulumi:"pcp"`
	Pmp         bool `pulumi:"pmp"`
	Udp         bool `pulumi:"udp"`
	Upnp        bool `pulumi:"upnp"`
}

type ClientSupportsOutput struct{ *pulumi.OutputState }

func (ClientSupportsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientSupports)(nil)).Elem()
}

func (o ClientSupportsOutput) ToClientSupportsOutput() ClientSupportsOutput {
	return o
}

func (o ClientSupportsOutput) ToClientSupportsOutputWithContext(ctx context.Context) ClientSupportsOutput {
	return o
}

func (o ClientSupportsOutput) HairPinning() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientSupports) bool { return v.HairPinning }).(pulumi.BoolOutput)
}

func (o ClientSupportsOutput) Ipv6() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientSupports) bool { return v.Ipv6 }).(pulumi.BoolOutput)
}

func (o ClientSupportsOutput) Pcp() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientSupports) bool { return v.Pcp }).(pulumi.BoolOutput)
}

func (o ClientSupportsOutput) Pmp() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientSupports) bool { return v.Pmp }).(pulumi.BoolOutput)
}

func (o ClientSupportsOutput) Udp() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientSupports) bool { return v.Udp }).(pulumi.BoolOutput)
}

func (o ClientSupportsOutput) Upnp() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientSupports) bool { return v.Upnp }).(pulumi.BoolOutput)
}

type CreateKey struct {
	Ephemeral     bool     `pulumi:"ephemeral"`
	Preauthorized bool     `pulumi:"preauthorized"`
	Reusable      bool     `pulumi:"reusable"`
	Tags          []string `pulumi:"tags"`
}

// CreateKeyInput is an input type that accepts CreateKeyArgs and CreateKeyOutput values.
// You can construct a concrete instance of `CreateKeyInput` via:
//
//	CreateKeyArgs{...}
type CreateKeyInput interface {
	pulumi.Input

	ToCreateKeyOutput() CreateKeyOutput
	ToCreateKeyOutputWithContext(context.Context) CreateKeyOutput
}

type CreateKeyArgs struct {
	Ephemeral     pulumi.BoolInput        `pulumi:"ephemeral"`
	Preauthorized pulumi.BoolInput        `pulumi:"preauthorized"`
	Reusable      pulumi.BoolInput        `pulumi:"reusable"`
	Tags          pulumi.StringArrayInput `pulumi:"tags"`
}

func (CreateKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateKey)(nil)).Elem()
}

func (i CreateKeyArgs) ToCreateKeyOutput() CreateKeyOutput {
	return i.ToCreateKeyOutputWithContext(context.Background())
}

func (i CreateKeyArgs) ToCreateKeyOutputWithContext(ctx context.Context) CreateKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateKeyOutput)
}

type CreateKeyOutput struct{ *pulumi.OutputState }

func (CreateKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateKey)(nil)).Elem()
}

func (o CreateKeyOutput) ToCreateKeyOutput() CreateKeyOutput {
	return o
}

func (o CreateKeyOutput) ToCreateKeyOutputWithContext(ctx context.Context) CreateKeyOutput {
	return o
}

func (o CreateKeyOutput) Ephemeral() pulumi.BoolOutput {
	return o.ApplyT(func(v CreateKey) bool { return v.Ephemeral }).(pulumi.BoolOutput)
}

func (o CreateKeyOutput) Preauthorized() pulumi.BoolOutput {
	return o.ApplyT(func(v CreateKey) bool { return v.Preauthorized }).(pulumi.BoolOutput)
}

func (o CreateKeyOutput) Reusable() pulumi.BoolOutput {
	return o.ApplyT(func(v CreateKey) bool { return v.Reusable }).(pulumi.BoolOutput)
}

func (o CreateKeyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CreateKey) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

type Device struct {
	Addresses                 []string           `pulumi:"addresses"`
	AdvertisedRoutes          []string           `pulumi:"advertisedRoutes"`
	Authorized                bool               `pulumi:"authorized"`
	BlocksIncomingConnections bool               `pulumi:"blocksIncomingConnections"`
	ClientConnectivity        ClientConnectivity `pulumi:"clientConnectivity"`
	ClientVersion             string             `pulumi:"clientVersion"`
	Created                   string             `pulumi:"created"`
	EnabledRoutes             []string           `pulumi:"enabledRoutes"`
	Expires                   string             `pulumi:"expires"`
	Hostname                  string             `pulumi:"hostname"`
	Id                        string             `pulumi:"id"`
	IsExternal                bool               `pulumi:"isExternal"`
	KeyExpiryDisabled         bool               `pulumi:"keyExpiryDisabled"`
	LastSeen                  string             `pulumi:"lastSeen"`
	MachineKey                string             `pulumi:"machineKey"`
	Name                      string             `pulumi:"name"`
	NodeKey                   string             `pulumi:"nodeKey"`
	Os                        string             `pulumi:"os"`
	UpdateAvailable           bool               `pulumi:"updateAvailable"`
	User                      string             `pulumi:"user"`
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func (o DeviceOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Device) []string { return v.Addresses }).(pulumi.StringArrayOutput)
}

func (o DeviceOutput) AdvertisedRoutes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Device) []string { return v.AdvertisedRoutes }).(pulumi.StringArrayOutput)
}

func (o DeviceOutput) Authorized() pulumi.BoolOutput {
	return o.ApplyT(func(v Device) bool { return v.Authorized }).(pulumi.BoolOutput)
}

func (o DeviceOutput) BlocksIncomingConnections() pulumi.BoolOutput {
	return o.ApplyT(func(v Device) bool { return v.BlocksIncomingConnections }).(pulumi.BoolOutput)
}

func (o DeviceOutput) ClientConnectivity() ClientConnectivityOutput {
	return o.ApplyT(func(v Device) ClientConnectivity { return v.ClientConnectivity }).(ClientConnectivityOutput)
}

func (o DeviceOutput) ClientVersion() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.ClientVersion }).(pulumi.StringOutput)
}

func (o DeviceOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.Created }).(pulumi.StringOutput)
}

func (o DeviceOutput) EnabledRoutes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Device) []string { return v.EnabledRoutes }).(pulumi.StringArrayOutput)
}

func (o DeviceOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.Expires }).(pulumi.StringOutput)
}

func (o DeviceOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o DeviceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.Id }).(pulumi.StringOutput)
}

func (o DeviceOutput) IsExternal() pulumi.BoolOutput {
	return o.ApplyT(func(v Device) bool { return v.IsExternal }).(pulumi.BoolOutput)
}

func (o DeviceOutput) KeyExpiryDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v Device) bool { return v.KeyExpiryDisabled }).(pulumi.BoolOutput)
}

func (o DeviceOutput) LastSeen() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.LastSeen }).(pulumi.StringOutput)
}

func (o DeviceOutput) MachineKey() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.MachineKey }).(pulumi.StringOutput)
}

func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceOutput) NodeKey() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.NodeKey }).(pulumi.StringOutput)
}

func (o DeviceOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.Os }).(pulumi.StringOutput)
}

func (o DeviceOutput) UpdateAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v Device) bool { return v.UpdateAvailable }).(pulumi.BoolOutput)
}

func (o DeviceOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v Device) string { return v.User }).(pulumi.StringOutput)
}

type DeviceArrayOutput struct{ *pulumi.OutputState }

func (DeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Device)(nil)).Elem()
}

func (o DeviceArrayOutput) ToDeviceArrayOutput() DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) Index(i pulumi.IntInput) DeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Device {
		return vs[0].([]Device)[vs[1].(int)]
	}).(DeviceOutput)
}

type DeviceKeyCapabilities struct {
	Create CreateKey `pulumi:"create"`
}

// DeviceKeyCapabilitiesInput is an input type that accepts DeviceKeyCapabilitiesArgs and DeviceKeyCapabilitiesOutput values.
// You can construct a concrete instance of `DeviceKeyCapabilitiesInput` via:
//
//	DeviceKeyCapabilitiesArgs{...}
type DeviceKeyCapabilitiesInput interface {
	pulumi.Input

	ToDeviceKeyCapabilitiesOutput() DeviceKeyCapabilitiesOutput
	ToDeviceKeyCapabilitiesOutputWithContext(context.Context) DeviceKeyCapabilitiesOutput
}

type DeviceKeyCapabilitiesArgs struct {
	Create CreateKeyInput `pulumi:"create"`
}

func (DeviceKeyCapabilitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceKeyCapabilities)(nil)).Elem()
}

func (i DeviceKeyCapabilitiesArgs) ToDeviceKeyCapabilitiesOutput() DeviceKeyCapabilitiesOutput {
	return i.ToDeviceKeyCapabilitiesOutputWithContext(context.Background())
}

func (i DeviceKeyCapabilitiesArgs) ToDeviceKeyCapabilitiesOutputWithContext(ctx context.Context) DeviceKeyCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceKeyCapabilitiesOutput)
}

type DeviceKeyCapabilitiesOutput struct{ *pulumi.OutputState }

func (DeviceKeyCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceKeyCapabilities)(nil)).Elem()
}

func (o DeviceKeyCapabilitiesOutput) ToDeviceKeyCapabilitiesOutput() DeviceKeyCapabilitiesOutput {
	return o
}

func (o DeviceKeyCapabilitiesOutput) ToDeviceKeyCapabilitiesOutputWithContext(ctx context.Context) DeviceKeyCapabilitiesOutput {
	return o
}

func (o DeviceKeyCapabilitiesOutput) Create() CreateKeyOutput {
	return o.ApplyT(func(v DeviceKeyCapabilities) CreateKey { return v.Create }).(CreateKeyOutput)
}

type DnsSearchPaths struct {
	SearchPaths []string `pulumi:"searchPaths"`
}

type KeyCapabilities struct {
	Devices DeviceKeyCapabilities `pulumi:"devices"`
}

// KeyCapabilitiesInput is an input type that accepts KeyCapabilitiesArgs and KeyCapabilitiesOutput values.
// You can construct a concrete instance of `KeyCapabilitiesInput` via:
//
//	KeyCapabilitiesArgs{...}
type KeyCapabilitiesInput interface {
	pulumi.Input

	ToKeyCapabilitiesOutput() KeyCapabilitiesOutput
	ToKeyCapabilitiesOutputWithContext(context.Context) KeyCapabilitiesOutput
}

type KeyCapabilitiesArgs struct {
	Devices DeviceKeyCapabilitiesInput `pulumi:"devices"`
}

func (KeyCapabilitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyCapabilities)(nil)).Elem()
}

func (i KeyCapabilitiesArgs) ToKeyCapabilitiesOutput() KeyCapabilitiesOutput {
	return i.ToKeyCapabilitiesOutputWithContext(context.Background())
}

func (i KeyCapabilitiesArgs) ToKeyCapabilitiesOutputWithContext(ctx context.Context) KeyCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyCapabilitiesOutput)
}

type KeyCapabilitiesOutput struct{ *pulumi.OutputState }

func (KeyCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyCapabilities)(nil)).Elem()
}

func (o KeyCapabilitiesOutput) ToKeyCapabilitiesOutput() KeyCapabilitiesOutput {
	return o
}

func (o KeyCapabilitiesOutput) ToKeyCapabilitiesOutputWithContext(ctx context.Context) KeyCapabilitiesOutput {
	return o
}

func (o KeyCapabilitiesOutput) Devices() DeviceKeyCapabilitiesOutput {
	return o.ApplyT(func(v KeyCapabilities) DeviceKeyCapabilities { return v.Devices }).(DeviceKeyCapabilitiesOutput)
}

type ListDevicesProperties struct {
	Devices []Device `pulumi:"devices"`
}

type NameServers struct {
	Dns      []string `pulumi:"dns"`
	MagicDNS bool     `pulumi:"magicDNS"`
}

type NameServersPreference struct {
	MagicDNS bool `pulumi:"magicDNS"`
}

type NodeAttrs struct {
	Attr   []string `pulumi:"attr"`
	Target []string `pulumi:"target"`
}

// NodeAttrsInput is an input type that accepts NodeAttrsArgs and NodeAttrsOutput values.
// You can construct a concrete instance of `NodeAttrsInput` via:
//
//	NodeAttrsArgs{...}
type NodeAttrsInput interface {
	pulumi.Input

	ToNodeAttrsOutput() NodeAttrsOutput
	ToNodeAttrsOutputWithContext(context.Context) NodeAttrsOutput
}

type NodeAttrsArgs struct {
	Attr   pulumi.StringArrayInput `pulumi:"attr"`
	Target pulumi.StringArrayInput `pulumi:"target"`
}

func (NodeAttrsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeAttrs)(nil)).Elem()
}

func (i NodeAttrsArgs) ToNodeAttrsOutput() NodeAttrsOutput {
	return i.ToNodeAttrsOutputWithContext(context.Background())
}

func (i NodeAttrsArgs) ToNodeAttrsOutputWithContext(ctx context.Context) NodeAttrsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeAttrsOutput)
}

// NodeAttrsArrayInput is an input type that accepts NodeAttrsArray and NodeAttrsArrayOutput values.
// You can construct a concrete instance of `NodeAttrsArrayInput` via:
//
//	NodeAttrsArray{ NodeAttrsArgs{...} }
type NodeAttrsArrayInput interface {
	pulumi.Input

	ToNodeAttrsArrayOutput() NodeAttrsArrayOutput
	ToNodeAttrsArrayOutputWithContext(context.Context) NodeAttrsArrayOutput
}

type NodeAttrsArray []NodeAttrsInput

func (NodeAttrsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeAttrs)(nil)).Elem()
}

func (i NodeAttrsArray) ToNodeAttrsArrayOutput() NodeAttrsArrayOutput {
	return i.ToNodeAttrsArrayOutputWithContext(context.Background())
}

func (i NodeAttrsArray) ToNodeAttrsArrayOutputWithContext(ctx context.Context) NodeAttrsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeAttrsArrayOutput)
}

type NodeAttrsOutput struct{ *pulumi.OutputState }

func (NodeAttrsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeAttrs)(nil)).Elem()
}

func (o NodeAttrsOutput) ToNodeAttrsOutput() NodeAttrsOutput {
	return o
}

func (o NodeAttrsOutput) ToNodeAttrsOutputWithContext(ctx context.Context) NodeAttrsOutput {
	return o
}

func (o NodeAttrsOutput) Attr() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeAttrs) []string { return v.Attr }).(pulumi.StringArrayOutput)
}

func (o NodeAttrsOutput) Target() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeAttrs) []string { return v.Target }).(pulumi.StringArrayOutput)
}

type NodeAttrsArrayOutput struct{ *pulumi.OutputState }

func (NodeAttrsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeAttrs)(nil)).Elem()
}

func (o NodeAttrsArrayOutput) ToNodeAttrsArrayOutput() NodeAttrsArrayOutput {
	return o
}

func (o NodeAttrsArrayOutput) ToNodeAttrsArrayOutputWithContext(ctx context.Context) NodeAttrsArrayOutput {
	return o
}

func (o NodeAttrsArrayOutput) Index(i pulumi.IntInput) NodeAttrsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeAttrs {
		return vs[0].([]NodeAttrs)[vs[1].(int)]
	}).(NodeAttrsOutput)
}

type SshRule struct {
	Action      SshRuleAction `pulumi:"action"`
	CheckPeriod string        `pulumi:"checkPeriod"`
	Dst         []string      `pulumi:"dst"`
	Src         []string      `pulumi:"src"`
	Users       []string      `pulumi:"users"`
}

// SshRuleInput is an input type that accepts SshRuleArgs and SshRuleOutput values.
// You can construct a concrete instance of `SshRuleInput` via:
//
//	SshRuleArgs{...}
type SshRuleInput interface {
	pulumi.Input

	ToSshRuleOutput() SshRuleOutput
	ToSshRuleOutputWithContext(context.Context) SshRuleOutput
}

type SshRuleArgs struct {
	Action      SshRuleActionInput      `pulumi:"action"`
	CheckPeriod pulumi.StringInput      `pulumi:"checkPeriod"`
	Dst         pulumi.StringArrayInput `pulumi:"dst"`
	Src         pulumi.StringArrayInput `pulumi:"src"`
	Users       pulumi.StringArrayInput `pulumi:"users"`
}

func (SshRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshRule)(nil)).Elem()
}

func (i SshRuleArgs) ToSshRuleOutput() SshRuleOutput {
	return i.ToSshRuleOutputWithContext(context.Background())
}

func (i SshRuleArgs) ToSshRuleOutputWithContext(ctx context.Context) SshRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshRuleOutput)
}

// SshRuleArrayInput is an input type that accepts SshRuleArray and SshRuleArrayOutput values.
// You can construct a concrete instance of `SshRuleArrayInput` via:
//
//	SshRuleArray{ SshRuleArgs{...} }
type SshRuleArrayInput interface {
	pulumi.Input

	ToSshRuleArrayOutput() SshRuleArrayOutput
	ToSshRuleArrayOutputWithContext(context.Context) SshRuleArrayOutput
}

type SshRuleArray []SshRuleInput

func (SshRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshRule)(nil)).Elem()
}

func (i SshRuleArray) ToSshRuleArrayOutput() SshRuleArrayOutput {
	return i.ToSshRuleArrayOutputWithContext(context.Background())
}

func (i SshRuleArray) ToSshRuleArrayOutputWithContext(ctx context.Context) SshRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshRuleArrayOutput)
}

type SshRuleOutput struct{ *pulumi.OutputState }

func (SshRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshRule)(nil)).Elem()
}

func (o SshRuleOutput) ToSshRuleOutput() SshRuleOutput {
	return o
}

func (o SshRuleOutput) ToSshRuleOutputWithContext(ctx context.Context) SshRuleOutput {
	return o
}

func (o SshRuleOutput) Action() SshRuleActionOutput {
	return o.ApplyT(func(v SshRule) SshRuleAction { return v.Action }).(SshRuleActionOutput)
}

func (o SshRuleOutput) CheckPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v SshRule) string { return v.CheckPeriod }).(pulumi.StringOutput)
}

func (o SshRuleOutput) Dst() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SshRule) []string { return v.Dst }).(pulumi.StringArrayOutput)
}

func (o SshRuleOutput) Src() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SshRule) []string { return v.Src }).(pulumi.StringArrayOutput)
}

func (o SshRuleOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SshRule) []string { return v.Users }).(pulumi.StringArrayOutput)
}

type SshRuleArrayOutput struct{ *pulumi.OutputState }

func (SshRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshRule)(nil)).Elem()
}

func (o SshRuleArrayOutput) ToSshRuleArrayOutput() SshRuleArrayOutput {
	return o
}

func (o SshRuleArrayOutput) ToSshRuleArrayOutputWithContext(ctx context.Context) SshRuleArrayOutput {
	return o
}

func (o SshRuleArrayOutput) Index(i pulumi.IntInput) SshRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshRule {
		return vs[0].([]SshRule)[vs[1].(int)]
	}).(SshRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleInput)(nil)).Elem(), AclRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleArrayInput)(nil)).Elem(), AclRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateKeyInput)(nil)).Elem(), CreateKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceKeyCapabilitiesInput)(nil)).Elem(), DeviceKeyCapabilitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyCapabilitiesInput)(nil)).Elem(), KeyCapabilitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeAttrsInput)(nil)).Elem(), NodeAttrsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeAttrsArrayInput)(nil)).Elem(), NodeAttrsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshRuleInput)(nil)).Elem(), SshRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshRuleArrayInput)(nil)).Elem(), SshRuleArray{})
	pulumi.RegisterOutputType(AclRuleOutput{})
	pulumi.RegisterOutputType(AclRuleArrayOutput{})
	pulumi.RegisterOutputType(AuthKeyOutput{})
	pulumi.RegisterOutputType(AuthKeyArrayOutput{})
	pulumi.RegisterOutputType(ClientConnectivityOutput{})
	pulumi.RegisterOutputType(ClientSupportsOutput{})
	pulumi.RegisterOutputType(CreateKeyOutput{})
	pulumi.RegisterOutputType(DeviceOutput{})
	pulumi.RegisterOutputType(DeviceArrayOutput{})
	pulumi.RegisterOutputType(DeviceKeyCapabilitiesOutput{})
	pulumi.RegisterOutputType(KeyCapabilitiesOutput{})
	pulumi.RegisterOutputType(NodeAttrsOutput{})
	pulumi.RegisterOutputType(NodeAttrsArrayOutput{})
	pulumi.RegisterOutputType(SshRuleOutput{})
	pulumi.RegisterOutputType(SshRuleArrayOutput{})
}
