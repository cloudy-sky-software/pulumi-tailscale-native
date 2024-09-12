// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKey(ctx *pulumi.Context, args *LookupKeyArgs, opts ...pulumi.InvokeOption) (*LookupKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKeyResult
	err := ctx.Invoke("tailscale-native:tailnet:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyArgs struct {
	Id string `pulumi:"id"`
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet string `pulumi:"tailnet"`
}

type LookupKeyResult struct {
	Created *string `pulumi:"created"`
	Expires string  `pulumi:"expires"`
	Key     string  `pulumi:"key"`
}

func LookupKeyOutput(ctx *pulumi.Context, args LookupKeyOutputArgs, opts ...pulumi.InvokeOption) LookupKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeyResultOutput, error) {
			args := v.(LookupKeyArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupKeyResult
			secret, err := ctx.InvokePackageRaw("tailscale-native:tailnet:getKey", args, &rv, "", opts...)
			if err != nil {
				return LookupKeyResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupKeyResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupKeyResultOutput), nil
			}
			return output, nil
		}).(LookupKeyResultOutput)
}

type LookupKeyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (LookupKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyArgs)(nil)).Elem()
}

type LookupKeyResultOutput struct{ *pulumi.OutputState }

func (LookupKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyResult)(nil)).Elem()
}

func (o LookupKeyResultOutput) ToLookupKeyResultOutput() LookupKeyResultOutput {
	return o
}

func (o LookupKeyResultOutput) ToLookupKeyResultOutputWithContext(ctx context.Context) LookupKeyResultOutput {
	return o
}

func (o LookupKeyResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o LookupKeyResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Expires }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

type LookupKeyResultArrayOutput struct{ *pulumi.OutputState }

func (LookupKeyResultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupKeyResult)(nil)).Elem()
}

func (o LookupKeyResultArrayOutput) ToLookupKeyResultArrayOutput() LookupKeyResultArrayOutput {
	return o
}

func (o LookupKeyResultArrayOutput) ToLookupKeyResultArrayOutputWithContext(ctx context.Context) LookupKeyResultArrayOutput {
	return o
}

func (o LookupKeyResultArrayOutput) Index(i pulumi.IntInput) LookupKeyResultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupKeyResult {
		return vs[0].([]LookupKeyResult)[vs[1].(int)]
	}).(LookupKeyResultOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyResultOutput{})
}
