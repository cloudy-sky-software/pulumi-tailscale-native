// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package device

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoutes(ctx *pulumi.Context, args *LookupRoutesArgs, opts ...pulumi.InvokeOption) (*LookupRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoutesResult
	err := ctx.Invoke("tailscale-native:device:getRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutesArgs struct {
	Id string `pulumi:"id"`
}

type LookupRoutesResult struct {
	Items DeviceRoutes `pulumi:"items"`
}

func LookupRoutesOutput(ctx *pulumi.Context, args LookupRoutesOutputArgs, opts ...pulumi.InvokeOption) LookupRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutesResult, error) {
			args := v.(LookupRoutesArgs)
			r, err := LookupRoutes(ctx, &args, opts...)
			var s LookupRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutesResultOutput)
}

type LookupRoutesOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutesArgs)(nil)).Elem()
}

type LookupRoutesResultOutput struct{ *pulumi.OutputState }

func (LookupRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutesResult)(nil)).Elem()
}

func (o LookupRoutesResultOutput) ToLookupRoutesResultOutput() LookupRoutesResultOutput {
	return o
}

func (o LookupRoutesResultOutput) ToLookupRoutesResultOutputWithContext(ctx context.Context) LookupRoutesResultOutput {
	return o
}

func (o LookupRoutesResultOutput) Items() DeviceRoutesOutput {
	return o.ApplyT(func(v LookupRoutesResult) DeviceRoutes { return v.Items }).(DeviceRoutesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutesResultOutput{})
}
