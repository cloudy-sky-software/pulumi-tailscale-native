// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAcl(ctx *pulumi.Context, args *LookupAclArgs, opts ...pulumi.InvokeOption) (*LookupAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAclResult
	err := ctx.Invoke("tailscale-native:tailnet:getAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAclArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet string `pulumi:"tailnet"`
}

type LookupAclResult struct {
	Items AclType `pulumi:"items"`
}

func LookupAclOutput(ctx *pulumi.Context, args LookupAclOutputArgs, opts ...pulumi.InvokeOption) LookupAclResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAclResult, error) {
			args := v.(LookupAclArgs)
			r, err := LookupAcl(ctx, &args, opts...)
			var s LookupAclResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAclResultOutput)
}

type LookupAclOutputArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (LookupAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAclArgs)(nil)).Elem()
}

type LookupAclResultOutput struct{ *pulumi.OutputState }

func (LookupAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAclResult)(nil)).Elem()
}

func (o LookupAclResultOutput) ToLookupAclResultOutput() LookupAclResultOutput {
	return o
}

func (o LookupAclResultOutput) ToLookupAclResultOutputWithContext(ctx context.Context) LookupAclResultOutput {
	return o
}

func (o LookupAclResultOutput) Items() AclTypeOutput {
	return o.ApplyT(func(v LookupAclResult) AclType { return v.Items }).(AclTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAclResultOutput{})
}
