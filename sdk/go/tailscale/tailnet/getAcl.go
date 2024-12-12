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
	Acls          []AclRule   `pulumi:"acls"`
	AutoApprovers interface{} `pulumi:"autoApprovers"`
	Groups        interface{} `pulumi:"groups"`
	Hosts         interface{} `pulumi:"hosts"`
	NodeAttrs     []NodeAttrs `pulumi:"nodeAttrs"`
	Ssh           []SshRule   `pulumi:"ssh"`
	TagOwners     interface{} `pulumi:"tagOwners"`
	Tests         string      `pulumi:"tests"`
}

func LookupAclOutput(ctx *pulumi.Context, args LookupAclOutputArgs, opts ...pulumi.InvokeOption) LookupAclResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAclResultOutput, error) {
			args := v.(LookupAclArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("tailscale-native:tailnet:getAcl", args, LookupAclResultOutput{}, options).(LookupAclResultOutput), nil
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

func (o LookupAclResultOutput) Acls() AclRuleArrayOutput {
	return o.ApplyT(func(v LookupAclResult) []AclRule { return v.Acls }).(AclRuleArrayOutput)
}

func (o LookupAclResultOutput) AutoApprovers() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAclResult) interface{} { return v.AutoApprovers }).(pulumi.AnyOutput)
}

func (o LookupAclResultOutput) Groups() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAclResult) interface{} { return v.Groups }).(pulumi.AnyOutput)
}

func (o LookupAclResultOutput) Hosts() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAclResult) interface{} { return v.Hosts }).(pulumi.AnyOutput)
}

func (o LookupAclResultOutput) NodeAttrs() NodeAttrsArrayOutput {
	return o.ApplyT(func(v LookupAclResult) []NodeAttrs { return v.NodeAttrs }).(NodeAttrsArrayOutput)
}

func (o LookupAclResultOutput) Ssh() SshRuleArrayOutput {
	return o.ApplyT(func(v LookupAclResult) []SshRule { return v.Ssh }).(SshRuleArrayOutput)
}

func (o LookupAclResultOutput) TagOwners() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAclResult) interface{} { return v.TagOwners }).(pulumi.AnyOutput)
}

func (o LookupAclResultOutput) Tests() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAclResult) string { return v.Tests }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAclResultOutput{})
}
