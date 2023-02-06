// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailnet

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AclRuleAction string

const (
	AclRuleActionAccept = AclRuleAction("accept")
)

func (AclRuleAction) ElementType() reflect.Type {
	return reflect.TypeOf((*AclRuleAction)(nil)).Elem()
}

func (e AclRuleAction) ToAclRuleActionOutput() AclRuleActionOutput {
	return pulumi.ToOutput(e).(AclRuleActionOutput)
}

func (e AclRuleAction) ToAclRuleActionOutputWithContext(ctx context.Context) AclRuleActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AclRuleActionOutput)
}

func (e AclRuleAction) ToAclRuleActionPtrOutput() AclRuleActionPtrOutput {
	return e.ToAclRuleActionPtrOutputWithContext(context.Background())
}

func (e AclRuleAction) ToAclRuleActionPtrOutputWithContext(ctx context.Context) AclRuleActionPtrOutput {
	return AclRuleAction(e).ToAclRuleActionOutputWithContext(ctx).ToAclRuleActionPtrOutputWithContext(ctx)
}

func (e AclRuleAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AclRuleAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AclRuleAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AclRuleAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AclRuleActionOutput struct{ *pulumi.OutputState }

func (AclRuleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AclRuleAction)(nil)).Elem()
}

func (o AclRuleActionOutput) ToAclRuleActionOutput() AclRuleActionOutput {
	return o
}

func (o AclRuleActionOutput) ToAclRuleActionOutputWithContext(ctx context.Context) AclRuleActionOutput {
	return o
}

func (o AclRuleActionOutput) ToAclRuleActionPtrOutput() AclRuleActionPtrOutput {
	return o.ToAclRuleActionPtrOutputWithContext(context.Background())
}

func (o AclRuleActionOutput) ToAclRuleActionPtrOutputWithContext(ctx context.Context) AclRuleActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AclRuleAction) *AclRuleAction {
		return &v
	}).(AclRuleActionPtrOutput)
}

func (o AclRuleActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AclRuleActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AclRuleAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AclRuleActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AclRuleActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AclRuleAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AclRuleActionPtrOutput struct{ *pulumi.OutputState }

func (AclRuleActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRuleAction)(nil)).Elem()
}

func (o AclRuleActionPtrOutput) ToAclRuleActionPtrOutput() AclRuleActionPtrOutput {
	return o
}

func (o AclRuleActionPtrOutput) ToAclRuleActionPtrOutputWithContext(ctx context.Context) AclRuleActionPtrOutput {
	return o
}

func (o AclRuleActionPtrOutput) Elem() AclRuleActionOutput {
	return o.ApplyT(func(v *AclRuleAction) AclRuleAction {
		if v != nil {
			return *v
		}
		var ret AclRuleAction
		return ret
	}).(AclRuleActionOutput)
}

func (o AclRuleActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AclRuleActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AclRuleAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AclRuleActionInput is an input type that accepts AclRuleActionArgs and AclRuleActionOutput values.
// You can construct a concrete instance of `AclRuleActionInput` via:
//
//	AclRuleActionArgs{...}
type AclRuleActionInput interface {
	pulumi.Input

	ToAclRuleActionOutput() AclRuleActionOutput
	ToAclRuleActionOutputWithContext(context.Context) AclRuleActionOutput
}

var aclRuleActionPtrType = reflect.TypeOf((**AclRuleAction)(nil)).Elem()

type AclRuleActionPtrInput interface {
	pulumi.Input

	ToAclRuleActionPtrOutput() AclRuleActionPtrOutput
	ToAclRuleActionPtrOutputWithContext(context.Context) AclRuleActionPtrOutput
}

type aclRuleActionPtr string

func AclRuleActionPtr(v string) AclRuleActionPtrInput {
	return (*aclRuleActionPtr)(&v)
}

func (*aclRuleActionPtr) ElementType() reflect.Type {
	return aclRuleActionPtrType
}

func (in *aclRuleActionPtr) ToAclRuleActionPtrOutput() AclRuleActionPtrOutput {
	return pulumi.ToOutput(in).(AclRuleActionPtrOutput)
}

func (in *aclRuleActionPtr) ToAclRuleActionPtrOutputWithContext(ctx context.Context) AclRuleActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AclRuleActionPtrOutput)
}

type SshRuleAction string

const (
	SshRuleActionCheck  = SshRuleAction("check")
	SshRuleActionAccept = SshRuleAction("accept")
)

func (SshRuleAction) ElementType() reflect.Type {
	return reflect.TypeOf((*SshRuleAction)(nil)).Elem()
}

func (e SshRuleAction) ToSshRuleActionOutput() SshRuleActionOutput {
	return pulumi.ToOutput(e).(SshRuleActionOutput)
}

func (e SshRuleAction) ToSshRuleActionOutputWithContext(ctx context.Context) SshRuleActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SshRuleActionOutput)
}

func (e SshRuleAction) ToSshRuleActionPtrOutput() SshRuleActionPtrOutput {
	return e.ToSshRuleActionPtrOutputWithContext(context.Background())
}

func (e SshRuleAction) ToSshRuleActionPtrOutputWithContext(ctx context.Context) SshRuleActionPtrOutput {
	return SshRuleAction(e).ToSshRuleActionOutputWithContext(ctx).ToSshRuleActionPtrOutputWithContext(ctx)
}

func (e SshRuleAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SshRuleAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SshRuleAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SshRuleAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SshRuleActionOutput struct{ *pulumi.OutputState }

func (SshRuleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshRuleAction)(nil)).Elem()
}

func (o SshRuleActionOutput) ToSshRuleActionOutput() SshRuleActionOutput {
	return o
}

func (o SshRuleActionOutput) ToSshRuleActionOutputWithContext(ctx context.Context) SshRuleActionOutput {
	return o
}

func (o SshRuleActionOutput) ToSshRuleActionPtrOutput() SshRuleActionPtrOutput {
	return o.ToSshRuleActionPtrOutputWithContext(context.Background())
}

func (o SshRuleActionOutput) ToSshRuleActionPtrOutputWithContext(ctx context.Context) SshRuleActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshRuleAction) *SshRuleAction {
		return &v
	}).(SshRuleActionPtrOutput)
}

func (o SshRuleActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SshRuleActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SshRuleAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SshRuleActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SshRuleActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SshRuleAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SshRuleActionPtrOutput struct{ *pulumi.OutputState }

func (SshRuleActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshRuleAction)(nil)).Elem()
}

func (o SshRuleActionPtrOutput) ToSshRuleActionPtrOutput() SshRuleActionPtrOutput {
	return o
}

func (o SshRuleActionPtrOutput) ToSshRuleActionPtrOutputWithContext(ctx context.Context) SshRuleActionPtrOutput {
	return o
}

func (o SshRuleActionPtrOutput) Elem() SshRuleActionOutput {
	return o.ApplyT(func(v *SshRuleAction) SshRuleAction {
		if v != nil {
			return *v
		}
		var ret SshRuleAction
		return ret
	}).(SshRuleActionOutput)
}

func (o SshRuleActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SshRuleActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SshRuleAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SshRuleActionInput is an input type that accepts SshRuleActionArgs and SshRuleActionOutput values.
// You can construct a concrete instance of `SshRuleActionInput` via:
//
//	SshRuleActionArgs{...}
type SshRuleActionInput interface {
	pulumi.Input

	ToSshRuleActionOutput() SshRuleActionOutput
	ToSshRuleActionOutputWithContext(context.Context) SshRuleActionOutput
}

var sshRuleActionPtrType = reflect.TypeOf((**SshRuleAction)(nil)).Elem()

type SshRuleActionPtrInput interface {
	pulumi.Input

	ToSshRuleActionPtrOutput() SshRuleActionPtrOutput
	ToSshRuleActionPtrOutputWithContext(context.Context) SshRuleActionPtrOutput
}

type sshRuleActionPtr string

func SshRuleActionPtr(v string) SshRuleActionPtrInput {
	return (*sshRuleActionPtr)(&v)
}

func (*sshRuleActionPtr) ElementType() reflect.Type {
	return sshRuleActionPtrType
}

func (in *sshRuleActionPtr) ToSshRuleActionPtrOutput() SshRuleActionPtrOutput {
	return pulumi.ToOutput(in).(SshRuleActionPtrOutput)
}

func (in *sshRuleActionPtr) ToSshRuleActionPtrOutputWithContext(ctx context.Context) SshRuleActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SshRuleActionPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleActionInput)(nil)).Elem(), AclRuleAction("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleActionPtrInput)(nil)).Elem(), AclRuleAction("accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*SshRuleActionInput)(nil)).Elem(), SshRuleAction("check"))
	pulumi.RegisterInputType(reflect.TypeOf((*SshRuleActionPtrInput)(nil)).Elem(), SshRuleAction("check"))
	pulumi.RegisterOutputType(AclRuleActionOutput{})
	pulumi.RegisterOutputType(AclRuleActionPtrOutput{})
	pulumi.RegisterOutputType(SshRuleActionOutput{})
	pulumi.RegisterOutputType(SshRuleActionPtrOutput{})
}
