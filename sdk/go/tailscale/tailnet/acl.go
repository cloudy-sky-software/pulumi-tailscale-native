// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Acl struct {
	pulumi.CustomResourceState

	Acls          AclRuleArrayOutput   `pulumi:"acls"`
	AutoApprovers pulumi.AnyOutput     `pulumi:"autoApprovers"`
	Groups        pulumi.AnyOutput     `pulumi:"groups"`
	Hosts         pulumi.AnyOutput     `pulumi:"hosts"`
	NodeAttrs     NodeAttrsArrayOutput `pulumi:"nodeAttrs"`
	Ssh           SshRuleArrayOutput   `pulumi:"ssh"`
	TagOwners     pulumi.AnyOutput     `pulumi:"tagOwners"`
	Tests         pulumi.StringOutput  `pulumi:"tests"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Acls == nil {
		return nil, errors.New("invalid value for required argument 'Acls'")
	}
	if args.AutoApprovers == nil {
		return nil, errors.New("invalid value for required argument 'AutoApprovers'")
	}
	if args.Groups == nil {
		return nil, errors.New("invalid value for required argument 'Groups'")
	}
	if args.Hosts == nil {
		return nil, errors.New("invalid value for required argument 'Hosts'")
	}
	if args.NodeAttrs == nil {
		return nil, errors.New("invalid value for required argument 'NodeAttrs'")
	}
	if args.Ssh == nil {
		return nil, errors.New("invalid value for required argument 'Ssh'")
	}
	if args.TagOwners == nil {
		return nil, errors.New("invalid value for required argument 'TagOwners'")
	}
	if args.Tests == nil {
		return nil, errors.New("invalid value for required argument 'Tests'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Acl
	err := ctx.RegisterResource("tailscale-native:tailnet:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("tailscale-native:tailnet:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
}

type AclState struct {
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	Acls          []AclRule   `pulumi:"acls"`
	AutoApprovers interface{} `pulumi:"autoApprovers"`
	Groups        interface{} `pulumi:"groups"`
	Hosts         interface{} `pulumi:"hosts"`
	NodeAttrs     []NodeAttrs `pulumi:"nodeAttrs"`
	Ssh           []SshRule   `pulumi:"ssh"`
	TagOwners     interface{} `pulumi:"tagOwners"`
	Tailnet       *string     `pulumi:"tailnet"`
	Tests         string      `pulumi:"tests"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	Acls          AclRuleArrayInput
	AutoApprovers pulumi.Input
	Groups        pulumi.Input
	Hosts         pulumi.Input
	NodeAttrs     NodeAttrsArrayInput
	Ssh           SshRuleArrayInput
	TagOwners     pulumi.Input
	Tailnet       pulumi.StringPtrInput
	Tests         pulumi.StringInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

func (o AclOutput) Acls() AclRuleArrayOutput {
	return o.ApplyT(func(v *Acl) AclRuleArrayOutput { return v.Acls }).(AclRuleArrayOutput)
}

func (o AclOutput) AutoApprovers() pulumi.AnyOutput {
	return o.ApplyT(func(v *Acl) pulumi.AnyOutput { return v.AutoApprovers }).(pulumi.AnyOutput)
}

func (o AclOutput) Groups() pulumi.AnyOutput {
	return o.ApplyT(func(v *Acl) pulumi.AnyOutput { return v.Groups }).(pulumi.AnyOutput)
}

func (o AclOutput) Hosts() pulumi.AnyOutput {
	return o.ApplyT(func(v *Acl) pulumi.AnyOutput { return v.Hosts }).(pulumi.AnyOutput)
}

func (o AclOutput) NodeAttrs() NodeAttrsArrayOutput {
	return o.ApplyT(func(v *Acl) NodeAttrsArrayOutput { return v.NodeAttrs }).(NodeAttrsArrayOutput)
}

func (o AclOutput) Ssh() SshRuleArrayOutput {
	return o.ApplyT(func(v *Acl) SshRuleArrayOutput { return v.Ssh }).(SshRuleArrayOutput)
}

func (o AclOutput) TagOwners() pulumi.AnyOutput {
	return o.ApplyT(func(v *Acl) pulumi.AnyOutput { return v.TagOwners }).(pulumi.AnyOutput)
}

func (o AclOutput) Tests() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.Tests }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclInput)(nil)).Elem(), &Acl{})
	pulumi.RegisterOutputType(AclOutput{})
}
