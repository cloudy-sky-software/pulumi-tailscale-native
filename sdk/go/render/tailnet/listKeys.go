// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKeys(ctx *pulumi.Context, args *ListKeysArgs, opts ...pulumi.InvokeOption) (*ListKeysResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ListKeysResult
	err := ctx.Invoke("tailscale:tailnet:listKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKeysArgs struct {
	Tailnet string `pulumi:"tailnet"`
}

type ListKeysResult struct {
	Items AuthKeyRead `pulumi:"items"`
}

func ListKeysOutput(ctx *pulumi.Context, args ListKeysOutputArgs, opts ...pulumi.InvokeOption) ListKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListKeysResult, error) {
			args := v.(ListKeysArgs)
			r, err := ListKeys(ctx, &args, opts...)
			var s ListKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListKeysResultOutput)
}

type ListKeysOutputArgs struct {
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (ListKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKeysArgs)(nil)).Elem()
}

type ListKeysResultOutput struct{ *pulumi.OutputState }

func (ListKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKeysResult)(nil)).Elem()
}

func (o ListKeysResultOutput) ToListKeysResultOutput() ListKeysResultOutput {
	return o
}

func (o ListKeysResultOutput) ToListKeysResultOutputWithContext(ctx context.Context) ListKeysResultOutput {
	return o
}

func (o ListKeysResultOutput) Items() AuthKeyReadOutput {
	return o.ApplyT(func(v ListKeysResult) AuthKeyRead { return v.Items }).(AuthKeyReadOutput)
}

func init() {
	pulumi.RegisterOutputType(ListKeysResultOutput{})
}
