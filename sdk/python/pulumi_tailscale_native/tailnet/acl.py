# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AclArgs', 'Acl']

@pulumi.input_type
class AclArgs:
    def __init__(__self__, *,
                 acls: pulumi.Input[Sequence[pulumi.Input['AclRuleArgs']]],
                 auto_approvers: Any,
                 groups: Any,
                 hosts: Any,
                 node_attrs: pulumi.Input[Sequence[pulumi.Input['NodeAttrsArgs']]],
                 ssh: pulumi.Input[Sequence[pulumi.Input['SshRuleArgs']]],
                 tag_owners: Any,
                 tests: pulumi.Input[str],
                 tailnet: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Acl resource.
        """
        pulumi.set(__self__, "acls", acls)
        pulumi.set(__self__, "auto_approvers", auto_approvers)
        pulumi.set(__self__, "groups", groups)
        pulumi.set(__self__, "hosts", hosts)
        pulumi.set(__self__, "node_attrs", node_attrs)
        pulumi.set(__self__, "ssh", ssh)
        pulumi.set(__self__, "tag_owners", tag_owners)
        pulumi.set(__self__, "tests", tests)
        if tailnet is not None:
            pulumi.set(__self__, "tailnet", tailnet)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Input[Sequence[pulumi.Input['AclRuleArgs']]]:
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: pulumi.Input[Sequence[pulumi.Input['AclRuleArgs']]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="autoApprovers")
    def auto_approvers(self) -> Any:
        return pulumi.get(self, "auto_approvers")

    @auto_approvers.setter
    def auto_approvers(self, value: Any):
        pulumi.set(self, "auto_approvers", value)

    @property
    @pulumi.getter
    def groups(self) -> Any:
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Any):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def hosts(self) -> Any:
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Any):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter(name="nodeAttrs")
    def node_attrs(self) -> pulumi.Input[Sequence[pulumi.Input['NodeAttrsArgs']]]:
        return pulumi.get(self, "node_attrs")

    @node_attrs.setter
    def node_attrs(self, value: pulumi.Input[Sequence[pulumi.Input['NodeAttrsArgs']]]):
        pulumi.set(self, "node_attrs", value)

    @property
    @pulumi.getter
    def ssh(self) -> pulumi.Input[Sequence[pulumi.Input['SshRuleArgs']]]:
        return pulumi.get(self, "ssh")

    @ssh.setter
    def ssh(self, value: pulumi.Input[Sequence[pulumi.Input['SshRuleArgs']]]):
        pulumi.set(self, "ssh", value)

    @property
    @pulumi.getter(name="tagOwners")
    def tag_owners(self) -> Any:
        return pulumi.get(self, "tag_owners")

    @tag_owners.setter
    def tag_owners(self, value: Any):
        pulumi.set(self, "tag_owners", value)

    @property
    @pulumi.getter
    def tests(self) -> pulumi.Input[str]:
        return pulumi.get(self, "tests")

    @tests.setter
    def tests(self, value: pulumi.Input[str]):
        pulumi.set(self, "tests", value)

    @property
    @pulumi.getter
    def tailnet(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tailnet")

    @tailnet.setter
    def tailnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tailnet", value)


class Acl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclRuleArgs']]]]] = None,
                 auto_approvers: Optional[Any] = None,
                 groups: Optional[Any] = None,
                 hosts: Optional[Any] = None,
                 node_attrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NodeAttrsArgs']]]]] = None,
                 ssh: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshRuleArgs']]]]] = None,
                 tag_owners: Optional[Any] = None,
                 tailnet: Optional[pulumi.Input[str]] = None,
                 tests: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Acl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Acl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclRuleArgs']]]]] = None,
                 auto_approvers: Optional[Any] = None,
                 groups: Optional[Any] = None,
                 hosts: Optional[Any] = None,
                 node_attrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NodeAttrsArgs']]]]] = None,
                 ssh: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshRuleArgs']]]]] = None,
                 tag_owners: Optional[Any] = None,
                 tailnet: Optional[pulumi.Input[str]] = None,
                 tests: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AclArgs.__new__(AclArgs)

            if acls is None and not opts.urn:
                raise TypeError("Missing required property 'acls'")
            __props__.__dict__["acls"] = acls
            if auto_approvers is None and not opts.urn:
                raise TypeError("Missing required property 'auto_approvers'")
            __props__.__dict__["auto_approvers"] = auto_approvers
            if groups is None and not opts.urn:
                raise TypeError("Missing required property 'groups'")
            __props__.__dict__["groups"] = groups
            if hosts is None and not opts.urn:
                raise TypeError("Missing required property 'hosts'")
            __props__.__dict__["hosts"] = hosts
            if node_attrs is None and not opts.urn:
                raise TypeError("Missing required property 'node_attrs'")
            __props__.__dict__["node_attrs"] = node_attrs
            if ssh is None and not opts.urn:
                raise TypeError("Missing required property 'ssh'")
            __props__.__dict__["ssh"] = ssh
            if tag_owners is None and not opts.urn:
                raise TypeError("Missing required property 'tag_owners'")
            __props__.__dict__["tag_owners"] = tag_owners
            __props__.__dict__["tailnet"] = tailnet
            if tests is None and not opts.urn:
                raise TypeError("Missing required property 'tests'")
            __props__.__dict__["tests"] = tests
        super(Acl, __self__).__init__(
            'tailscale-native:tailnet:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Acl':
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AclArgs.__new__(AclArgs)

        __props__.__dict__["acls"] = None
        __props__.__dict__["auto_approvers"] = None
        __props__.__dict__["groups"] = None
        __props__.__dict__["hosts"] = None
        __props__.__dict__["node_attrs"] = None
        __props__.__dict__["ssh"] = None
        __props__.__dict__["tag_owners"] = None
        __props__.__dict__["tests"] = None
        return Acl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Output[Sequence['outputs.AclRule']]:
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="autoApprovers")
    def auto_approvers(self) -> pulumi.Output[Any]:
        return pulumi.get(self, "auto_approvers")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Any]:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Output[Any]:
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter(name="nodeAttrs")
    def node_attrs(self) -> pulumi.Output[Sequence['outputs.NodeAttrs']]:
        return pulumi.get(self, "node_attrs")

    @property
    @pulumi.getter
    def ssh(self) -> pulumi.Output[Sequence['outputs.SshRule']]:
        return pulumi.get(self, "ssh")

    @property
    @pulumi.getter(name="tagOwners")
    def tag_owners(self) -> pulumi.Output[Any]:
        return pulumi.get(self, "tag_owners")

    @property
    @pulumi.getter
    def tests(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tests")

