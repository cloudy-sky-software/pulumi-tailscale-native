# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['AuthorizeDeviceArgs', 'AuthorizeDevice']

@pulumi.input_type
class AuthorizeDeviceArgs:
    def __init__(__self__, *,
                 authorized: pulumi.Input[builtins.bool],
                 id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a AuthorizeDevice resource.
        """
        pulumi.set(__self__, "authorized", authorized)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Input[builtins.bool]:
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)


@pulumi.type_token("tailscale-native:device:AuthorizeDevice")
class AuthorizeDevice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[builtins.bool]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a AuthorizeDevice resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizeDeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AuthorizeDevice resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AuthorizeDeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizeDeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[builtins.bool]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizeDeviceArgs.__new__(AuthorizeDeviceArgs)

            if authorized is None and not opts.urn:
                raise TypeError("Missing required property 'authorized'")
            __props__.__dict__["authorized"] = authorized
            __props__.__dict__["id"] = id
        super(AuthorizeDevice, __self__).__init__(
            'tailscale-native:device:AuthorizeDevice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AuthorizeDevice':
        """
        Get an existing AuthorizeDevice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AuthorizeDeviceArgs.__new__(AuthorizeDeviceArgs)

        __props__.__dict__["authorized"] = None
        return AuthorizeDevice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Output[builtins.bool]:
        return pulumi.get(self, "authorized")

