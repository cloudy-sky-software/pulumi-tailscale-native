// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDNSPreferences(ctx *pulumi.Context, args *LookupDNSPreferencesArgs, opts ...pulumi.InvokeOption) (*LookupDNSPreferencesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDNSPreferencesResult
	err := ctx.Invoke("tailscale:tailnet:getDNSPreferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDNSPreferencesArgs struct {
	Tailnet string `pulumi:"tailnet"`
}

type LookupDNSPreferencesResult struct {
	Items NameServersPreference `pulumi:"items"`
}

func LookupDNSPreferencesOutput(ctx *pulumi.Context, args LookupDNSPreferencesOutputArgs, opts ...pulumi.InvokeOption) LookupDNSPreferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDNSPreferencesResult, error) {
			args := v.(LookupDNSPreferencesArgs)
			r, err := LookupDNSPreferences(ctx, &args, opts...)
			var s LookupDNSPreferencesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDNSPreferencesResultOutput)
}

type LookupDNSPreferencesOutputArgs struct {
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (LookupDNSPreferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDNSPreferencesArgs)(nil)).Elem()
}

type LookupDNSPreferencesResultOutput struct{ *pulumi.OutputState }

func (LookupDNSPreferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDNSPreferencesResult)(nil)).Elem()
}

func (o LookupDNSPreferencesResultOutput) ToLookupDNSPreferencesResultOutput() LookupDNSPreferencesResultOutput {
	return o
}

func (o LookupDNSPreferencesResultOutput) ToLookupDNSPreferencesResultOutputWithContext(ctx context.Context) LookupDNSPreferencesResultOutput {
	return o
}

func (o LookupDNSPreferencesResultOutput) Items() NameServersPreferenceOutput {
	return o.ApplyT(func(v LookupDNSPreferencesResult) NameServersPreference { return v.Items }).(NameServersPreferenceOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDNSPreferencesResultOutput{})
}