// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDNSPreferencesConfig(ctx *pulumi.Context, args *LookupDNSPreferencesConfigArgs, opts ...pulumi.InvokeOption) (*LookupDNSPreferencesConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDNSPreferencesConfigResult
	err := ctx.Invoke("tailscale-native:tailnet:getDNSPreferencesConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDNSPreferencesConfigArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet string `pulumi:"tailnet"`
}

type LookupDNSPreferencesConfigResult struct {
	Items NameServersPreference `pulumi:"items"`
}

func LookupDNSPreferencesConfigOutput(ctx *pulumi.Context, args LookupDNSPreferencesConfigOutputArgs, opts ...pulumi.InvokeOption) LookupDNSPreferencesConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDNSPreferencesConfigResult, error) {
			args := v.(LookupDNSPreferencesConfigArgs)
			r, err := LookupDNSPreferencesConfig(ctx, &args, opts...)
			var s LookupDNSPreferencesConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDNSPreferencesConfigResultOutput)
}

type LookupDNSPreferencesConfigOutputArgs struct {
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet pulumi.StringInput `pulumi:"tailnet"`
}

func (LookupDNSPreferencesConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDNSPreferencesConfigArgs)(nil)).Elem()
}

type LookupDNSPreferencesConfigResultOutput struct{ *pulumi.OutputState }

func (LookupDNSPreferencesConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDNSPreferencesConfigResult)(nil)).Elem()
}

func (o LookupDNSPreferencesConfigResultOutput) ToLookupDNSPreferencesConfigResultOutput() LookupDNSPreferencesConfigResultOutput {
	return o
}

func (o LookupDNSPreferencesConfigResultOutput) ToLookupDNSPreferencesConfigResultOutputWithContext(ctx context.Context) LookupDNSPreferencesConfigResultOutput {
	return o
}

func (o LookupDNSPreferencesConfigResultOutput) Items() NameServersPreferenceOutput {
	return o.ApplyT(func(v LookupDNSPreferencesConfigResult) NameServersPreference { return v.Items }).(NameServersPreferenceOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDNSPreferencesConfigResultOutput{})
}
