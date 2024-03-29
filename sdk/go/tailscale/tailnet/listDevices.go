// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDevices(ctx *pulumi.Context, args *ListDevicesArgs, opts ...pulumi.InvokeOption) (*ListDevicesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDevicesResult
	err := ctx.Invoke("tailscale-native:tailnet:listDevices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDevicesArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet string `pulumi:"tailnet"`
}

type ListDevicesResult struct {
	Items ListDevicesProperties `pulumi:"items"`
}

func ListDevicesOutput(ctx *pulumi.Context, args ListDevicesOutputArgs, opts ...pulumi.InvokeOption) ListDevicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDevicesResult, error) {
			args := v.(ListDevicesArgs)
			r, err := ListDevices(ctx, &args, opts...)
			var s ListDevicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDevicesResultOutput)
}

type ListDevicesOutputArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (ListDevicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDevicesArgs)(nil)).Elem()
}

type ListDevicesResultOutput struct{ *pulumi.OutputState }

func (ListDevicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDevicesResult)(nil)).Elem()
}

func (o ListDevicesResultOutput) ToListDevicesResultOutput() ListDevicesResultOutput {
	return o
}

func (o ListDevicesResultOutput) ToListDevicesResultOutputWithContext(ctx context.Context) ListDevicesResultOutput {
	return o
}

func (o ListDevicesResultOutput) Items() ListDevicesPropertiesOutput {
	return o.ApplyT(func(v ListDevicesResult) ListDevicesProperties { return v.Items }).(ListDevicesPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDevicesResultOutput{})
}
