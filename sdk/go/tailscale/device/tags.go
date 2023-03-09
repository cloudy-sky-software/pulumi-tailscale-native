// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package device

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Tags struct {
	pulumi.CustomResourceState

	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewTags registers a new resource with the given unique name, arguments, and options.
func NewTags(ctx *pulumi.Context,
	name string, args *TagsArgs, opts ...pulumi.ResourceOption) (*Tags, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Tags == nil {
		return nil, errors.New("invalid value for required argument 'Tags'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Tags
	err := ctx.RegisterResource("tailscale-native:device:Tags", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTags gets an existing Tags resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTags(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagsState, opts ...pulumi.ResourceOption) (*Tags, error) {
	var resource Tags
	err := ctx.ReadResource("tailscale-native:device:Tags", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tags resources.
type tagsState struct {
}

type TagsState struct {
}

func (TagsState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsState)(nil)).Elem()
}

type tagsArgs struct {
	Id   *string  `pulumi:"id"`
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Tags resource.
type TagsArgs struct {
	Id   pulumi.StringPtrInput
	Tags pulumi.StringArrayInput
}

func (TagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagsArgs)(nil)).Elem()
}

type TagsInput interface {
	pulumi.Input

	ToTagsOutput() TagsOutput
	ToTagsOutputWithContext(ctx context.Context) TagsOutput
}

func (*Tags) ElementType() reflect.Type {
	return reflect.TypeOf((**Tags)(nil)).Elem()
}

func (i *Tags) ToTagsOutput() TagsOutput {
	return i.ToTagsOutputWithContext(context.Background())
}

func (i *Tags) ToTagsOutputWithContext(ctx context.Context) TagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsOutput)
}

type TagsOutput struct{ *pulumi.OutputState }

func (TagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tags)(nil)).Elem()
}

func (o TagsOutput) ToTagsOutput() TagsOutput {
	return o
}

func (o TagsOutput) ToTagsOutputWithContext(ctx context.Context) TagsOutput {
	return o
}

func (o TagsOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Tags) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagsInput)(nil)).Elem(), &Tags{})
	pulumi.RegisterOutputType(TagsOutput{})
}
