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

type ReplaceSearchPaths struct {
	pulumi.CustomResourceState

	SearchPaths pulumi.StringArrayOutput `pulumi:"searchPaths"`
}

// NewReplaceSearchPaths registers a new resource with the given unique name, arguments, and options.
func NewReplaceSearchPaths(ctx *pulumi.Context,
	name string, args *ReplaceSearchPathsArgs, opts ...pulumi.ResourceOption) (*ReplaceSearchPaths, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SearchPaths == nil {
		return nil, errors.New("invalid value for required argument 'SearchPaths'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplaceSearchPaths
	err := ctx.RegisterResource("tailscale-native:tailnet:ReplaceSearchPaths", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplaceSearchPaths gets an existing ReplaceSearchPaths resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplaceSearchPaths(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplaceSearchPathsState, opts ...pulumi.ResourceOption) (*ReplaceSearchPaths, error) {
	var resource ReplaceSearchPaths
	err := ctx.ReadResource("tailscale-native:tailnet:ReplaceSearchPaths", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplaceSearchPaths resources.
type replaceSearchPathsState struct {
}

type ReplaceSearchPathsState struct {
}

func (ReplaceSearchPathsState) ElementType() reflect.Type {
	return reflect.TypeOf((*replaceSearchPathsState)(nil)).Elem()
}

type replaceSearchPathsArgs struct {
	SearchPaths []string `pulumi:"searchPaths"`
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet *string `pulumi:"tailnet"`
}

// The set of arguments for constructing a ReplaceSearchPaths resource.
type ReplaceSearchPathsArgs struct {
	SearchPaths pulumi.StringArrayInput
	// For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
	Tailnet pulumi.StringPtrInput
}

func (ReplaceSearchPathsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replaceSearchPathsArgs)(nil)).Elem()
}

type ReplaceSearchPathsInput interface {
	pulumi.Input

	ToReplaceSearchPathsOutput() ReplaceSearchPathsOutput
	ToReplaceSearchPathsOutputWithContext(ctx context.Context) ReplaceSearchPathsOutput
}

func (*ReplaceSearchPaths) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplaceSearchPaths)(nil)).Elem()
}

func (i *ReplaceSearchPaths) ToReplaceSearchPathsOutput() ReplaceSearchPathsOutput {
	return i.ToReplaceSearchPathsOutputWithContext(context.Background())
}

func (i *ReplaceSearchPaths) ToReplaceSearchPathsOutputWithContext(ctx context.Context) ReplaceSearchPathsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplaceSearchPathsOutput)
}

type ReplaceSearchPathsOutput struct{ *pulumi.OutputState }

func (ReplaceSearchPathsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplaceSearchPaths)(nil)).Elem()
}

func (o ReplaceSearchPathsOutput) ToReplaceSearchPathsOutput() ReplaceSearchPathsOutput {
	return o
}

func (o ReplaceSearchPathsOutput) ToReplaceSearchPathsOutputWithContext(ctx context.Context) ReplaceSearchPathsOutput {
	return o
}

func (o ReplaceSearchPathsOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReplaceSearchPaths) pulumi.StringArrayOutput { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplaceSearchPathsInput)(nil)).Elem(), &ReplaceSearchPaths{})
	pulumi.RegisterOutputType(ReplaceSearchPathsOutput{})
}
