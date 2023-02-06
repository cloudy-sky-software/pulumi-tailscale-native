// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKey(ctx *pulumi.Context, args *GetKeyArgs, opts ...pulumi.InvokeOption) (*GetKeyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetKeyResult
	err := ctx.Invoke("tailscale:tailnet:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetKeyArgs struct {
	Id      string `pulumi:"id"`
	Tailnet string `pulumi:"tailnet"`
}

type GetKeyResult struct {
	Items AuthKeyRead `pulumi:"items"`
}

func GetKeyOutput(ctx *pulumi.Context, args GetKeyOutputArgs, opts ...pulumi.InvokeOption) GetKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKeyResult, error) {
			args := v.(GetKeyArgs)
			r, err := GetKey(ctx, &args, opts...)
			var s GetKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKeyResultOutput)
}

type GetKeyOutputArgs struct {
	Id      pulumi.StringInput `pulumi:"id"`
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (GetKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyArgs)(nil)).Elem()
}

type GetKeyResultOutput struct{ *pulumi.OutputState }

func (GetKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyResult)(nil)).Elem()
}

func (o GetKeyResultOutput) ToGetKeyResultOutput() GetKeyResultOutput {
	return o
}

func (o GetKeyResultOutput) ToGetKeyResultOutputWithContext(ctx context.Context) GetKeyResultOutput {
	return o
}

func (o GetKeyResultOutput) Items() AuthKeyReadOutput {
	return o.ApplyT(func(v GetKeyResult) AuthKeyRead { return v.Items }).(AuthKeyReadOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKeyResultOutput{})
}
