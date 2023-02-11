// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DNSPreferences struct {
	pulumi.CustomResourceState

	MagicDNS pulumi.BoolOutput `pulumi:"magicDNS"`
}

// NewDNSPreferences registers a new resource with the given unique name, arguments, and options.
func NewDNSPreferences(ctx *pulumi.Context,
	name string, args *DNSPreferencesArgs, opts ...pulumi.ResourceOption) (*DNSPreferences, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MagicDNS == nil {
		return nil, errors.New("invalid value for required argument 'MagicDNS'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DNSPreferences
	err := ctx.RegisterResource("tailscale:tailnet:DNSPreferences", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDNSPreferences gets an existing DNSPreferences resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDNSPreferences(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DNSPreferencesState, opts ...pulumi.ResourceOption) (*DNSPreferences, error) {
	var resource DNSPreferences
	err := ctx.ReadResource("tailscale:tailnet:DNSPreferences", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DNSPreferences resources.
type dnspreferencesState struct {
}

type DNSPreferencesState struct {
}

func (DNSPreferencesState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnspreferencesState)(nil)).Elem()
}

type dnspreferencesArgs struct {
	MagicDNS bool    `pulumi:"magicDNS"`
	Tailnet  *string `pulumi:"tailnet"`
}

// The set of arguments for constructing a DNSPreferences resource.
type DNSPreferencesArgs struct {
	MagicDNS pulumi.BoolInput
	Tailnet  pulumi.StringPtrInput
}

func (DNSPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnspreferencesArgs)(nil)).Elem()
}

type DNSPreferencesInput interface {
	pulumi.Input

	ToDNSPreferencesOutput() DNSPreferencesOutput
	ToDNSPreferencesOutputWithContext(ctx context.Context) DNSPreferencesOutput
}

func (*DNSPreferences) ElementType() reflect.Type {
	return reflect.TypeOf((**DNSPreferences)(nil)).Elem()
}

func (i *DNSPreferences) ToDNSPreferencesOutput() DNSPreferencesOutput {
	return i.ToDNSPreferencesOutputWithContext(context.Background())
}

func (i *DNSPreferences) ToDNSPreferencesOutputWithContext(ctx context.Context) DNSPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DNSPreferencesOutput)
}

type DNSPreferencesOutput struct{ *pulumi.OutputState }

func (DNSPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DNSPreferences)(nil)).Elem()
}

func (o DNSPreferencesOutput) ToDNSPreferencesOutput() DNSPreferencesOutput {
	return o
}

func (o DNSPreferencesOutput) ToDNSPreferencesOutputWithContext(ctx context.Context) DNSPreferencesOutput {
	return o
}

func (o DNSPreferencesOutput) MagicDNS() pulumi.BoolOutput {
	return o.ApplyT(func(v *DNSPreferences) pulumi.BoolOutput { return v.MagicDNS }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DNSPreferencesInput)(nil)).Elem(), &DNSPreferences{})
	pulumi.RegisterOutputType(DNSPreferencesOutput{})
}
