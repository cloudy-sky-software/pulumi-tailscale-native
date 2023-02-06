// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NameServers struct {
	pulumi.CustomResourceState

	Dns      pulumi.StringArrayOutput `pulumi:"dns"`
	MagicDNS pulumi.BoolPtrOutput     `pulumi:"magicDNS"`
}

// NewNameServers registers a new resource with the given unique name, arguments, and options.
func NewNameServers(ctx *pulumi.Context,
	name string, args *NameServersArgs, opts ...pulumi.ResourceOption) (*NameServers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dns == nil {
		return nil, errors.New("invalid value for required argument 'Dns'")
	}
	if args.MagicDNS == nil {
		return nil, errors.New("invalid value for required argument 'MagicDNS'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NameServers
	err := ctx.RegisterResource("tailscale:tailnet:NameServers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNameServers gets an existing NameServers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNameServers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NameServersState, opts ...pulumi.ResourceOption) (*NameServers, error) {
	var resource NameServers
	err := ctx.ReadResource("tailscale:tailnet:NameServers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NameServers resources.
type nameServersState struct {
}

type NameServersState struct {
}

func (NameServersState) ElementType() reflect.Type {
	return reflect.TypeOf((*nameServersState)(nil)).Elem()
}

type nameServersArgs struct {
	Dns      []string `pulumi:"dns"`
	MagicDNS bool     `pulumi:"magicDNS"`
	Tailnet  *string  `pulumi:"tailnet"`
}

// The set of arguments for constructing a NameServers resource.
type NameServersArgs struct {
	Dns      pulumi.StringArrayInput
	MagicDNS pulumi.BoolInput
	Tailnet  pulumi.StringPtrInput
}

func (NameServersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nameServersArgs)(nil)).Elem()
}

type NameServersInput interface {
	pulumi.Input

	ToNameServersOutput() NameServersOutput
	ToNameServersOutputWithContext(ctx context.Context) NameServersOutput
}

func (*NameServers) ElementType() reflect.Type {
	return reflect.TypeOf((**NameServers)(nil)).Elem()
}

func (i *NameServers) ToNameServersOutput() NameServersOutput {
	return i.ToNameServersOutputWithContext(context.Background())
}

func (i *NameServers) ToNameServersOutputWithContext(ctx context.Context) NameServersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameServersOutput)
}

type NameServersOutput struct{ *pulumi.OutputState }

func (NameServersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NameServers)(nil)).Elem()
}

func (o NameServersOutput) ToNameServersOutput() NameServersOutput {
	return o
}

func (o NameServersOutput) ToNameServersOutputWithContext(ctx context.Context) NameServersOutput {
	return o
}

func (o NameServersOutput) Dns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NameServers) pulumi.StringArrayOutput { return v.Dns }).(pulumi.StringArrayOutput)
}

func (o NameServersOutput) MagicDNS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NameServers) pulumi.BoolPtrOutput { return v.MagicDNS }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NameServersInput)(nil)).Elem(), &NameServers{})
	pulumi.RegisterOutputType(NameServersOutput{})
}
