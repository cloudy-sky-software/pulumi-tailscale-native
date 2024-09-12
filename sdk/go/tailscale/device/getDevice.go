// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package device

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDevice(ctx *pulumi.Context, args *GetDeviceArgs, opts ...pulumi.InvokeOption) (*GetDeviceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDeviceResult
	err := ctx.Invoke("tailscale-native:device:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeviceArgs struct {
	Id string `pulumi:"id"`
}

type GetDeviceResult struct {
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

func GetDeviceOutput(ctx *pulumi.Context, args GetDeviceOutputArgs, opts ...pulumi.InvokeOption) GetDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeviceResultOutput, error) {
			args := v.(GetDeviceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDeviceResult
			secret, err := ctx.InvokePackageRaw("tailscale-native:device:getDevice", args, &rv, "", opts...)
			if err != nil {
				return GetDeviceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDeviceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDeviceResultOutput), nil
			}
			return output, nil
		}).(GetDeviceResultOutput)
}

type GetDeviceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceArgs)(nil)).Elem()
}

type GetDeviceResultOutput struct{ *pulumi.OutputState }

func (GetDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceResult)(nil)).Elem()
}

func (o GetDeviceResultOutput) ToGetDeviceResultOutput() GetDeviceResultOutput {
	return o
}

func (o GetDeviceResultOutput) ToGetDeviceResultOutputWithContext(ctx context.Context) GetDeviceResultOutput {
	return o
}

func (o GetDeviceResultOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeviceResult) []string { return v.Addresses }).(pulumi.StringArrayOutput)
}

func (o GetDeviceResultOutput) AdvertisedRoutes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeviceResult) []string { return v.AdvertisedRoutes }).(pulumi.StringArrayOutput)
}

func (o GetDeviceResultOutput) Authorized() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDeviceResult) bool { return v.Authorized }).(pulumi.BoolOutput)
}

func (o GetDeviceResultOutput) BlocksIncomingConnections() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDeviceResult) bool { return v.BlocksIncomingConnections }).(pulumi.BoolOutput)
}

func (o GetDeviceResultOutput) ClientConnectivity() ClientConnectivityOutput {
	return o.ApplyT(func(v GetDeviceResult) ClientConnectivity { return v.ClientConnectivity }).(ClientConnectivityOutput)
}

func (o GetDeviceResultOutput) ClientVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.ClientVersion }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) EnabledRoutes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeviceResult) []string { return v.EnabledRoutes }).(pulumi.StringArrayOutput)
}

func (o GetDeviceResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Expires }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) IsExternal() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDeviceResult) bool { return v.IsExternal }).(pulumi.BoolOutput)
}

func (o GetDeviceResultOutput) KeyExpiryDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDeviceResult) bool { return v.KeyExpiryDisabled }).(pulumi.BoolOutput)
}

func (o GetDeviceResultOutput) LastSeen() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.LastSeen }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) MachineKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.MachineKey }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) NodeKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.NodeKey }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Os }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) UpdateAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDeviceResult) bool { return v.UpdateAvailable }).(pulumi.BoolOutput)
}

func (o GetDeviceResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.User }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeviceResultOutput{})
}
