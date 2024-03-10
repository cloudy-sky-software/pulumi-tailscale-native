// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSearchPaths(ctx *pulumi.Context, args *ListSearchPathsArgs, opts ...pulumi.InvokeOption) (*ListSearchPathsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListSearchPathsResult
	err := ctx.Invoke("tailscale-native:tailnet:listSearchPaths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSearchPathsArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet string `pulumi:"tailnet"`
}

type ListSearchPathsResult struct {
	Items DnsSearchPaths `pulumi:"items"`
}

func ListSearchPathsOutput(ctx *pulumi.Context, args ListSearchPathsOutputArgs, opts ...pulumi.InvokeOption) ListSearchPathsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSearchPathsResult, error) {
			args := v.(ListSearchPathsArgs)
			r, err := ListSearchPaths(ctx, &args, opts...)
			var s ListSearchPathsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSearchPathsResultOutput)
}

type ListSearchPathsOutputArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (ListSearchPathsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSearchPathsArgs)(nil)).Elem()
}

type ListSearchPathsResultOutput struct{ *pulumi.OutputState }

func (ListSearchPathsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSearchPathsResult)(nil)).Elem()
}

func (o ListSearchPathsResultOutput) ToListSearchPathsResultOutput() ListSearchPathsResultOutput {
	return o
}

func (o ListSearchPathsResultOutput) ToListSearchPathsResultOutputWithContext(ctx context.Context) ListSearchPathsResultOutput {
	return o
}

func (o ListSearchPathsResultOutput) Items() DnsSearchPathsOutput {
	return o.ApplyT(func(v ListSearchPathsResult) DnsSearchPaths { return v.Items }).(DnsSearchPathsOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSearchPathsResultOutput{})
}
