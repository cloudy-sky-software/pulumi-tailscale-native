// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Key struct {
	pulumi.CustomResourceState

	Capabilities  KeyCapabilitiesOutput  `pulumi:"capabilities"`
	Created       pulumi.StringPtrOutput `pulumi:"created"`
	Expires       pulumi.StringOutput    `pulumi:"expires"`
	ExpirySeconds pulumi.IntOutput       `pulumi:"expirySeconds"`
	Key           pulumi.StringOutput    `pulumi:"key"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capabilities == nil {
		return nil, errors.New("invalid value for required argument 'Capabilities'")
	}
	if args.ExpirySeconds == nil {
		return nil, errors.New("invalid value for required argument 'ExpirySeconds'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Key
	err := ctx.RegisterResource("tailscale:tailnet:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("tailscale:tailnet:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	Capabilities  KeyCapabilities `pulumi:"capabilities"`
	ExpirySeconds int             `pulumi:"expirySeconds"`
	Tailnet       *string         `pulumi:"tailnet"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	Capabilities  KeyCapabilitiesInput
	ExpirySeconds pulumi.IntInput
	Tailnet       pulumi.StringPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func (o KeyOutput) Capabilities() KeyCapabilitiesOutput {
	return o.ApplyT(func(v *Key) KeyCapabilitiesOutput { return v.Capabilities }).(KeyCapabilitiesOutput)
}

func (o KeyOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.Created }).(pulumi.StringPtrOutput)
}

func (o KeyOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Expires }).(pulumi.StringOutput)
}

func (o KeyOutput) ExpirySeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *Key) pulumi.IntOutput { return v.ExpirySeconds }).(pulumi.IntOutput)
}

func (o KeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterOutputType(KeyOutput{})
}
