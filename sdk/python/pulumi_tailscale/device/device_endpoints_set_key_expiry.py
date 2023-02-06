# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeviceEndpoints_setKeyExpiryArgs', 'DeviceEndpoints_setKeyExpiry']

@pulumi.input_type
class DeviceEndpoints_setKeyExpiryArgs:
    def __init__(__self__, *,
                 key_expiry_disabled: pulumi.Input[bool],
                 id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DeviceEndpoints_setKeyExpiry resource.
        """
        pulumi.set(__self__, "key_expiry_disabled", key_expiry_disabled)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "key_expiry_disabled")

    @key_expiry_disabled.setter
    def key_expiry_disabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "key_expiry_disabled", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


class DeviceEndpoints_setKeyExpiry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 key_expiry_disabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a DeviceEndpoints_setKeyExpiry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceEndpoints_setKeyExpiryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DeviceEndpoints_setKeyExpiry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DeviceEndpoints_setKeyExpiryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceEndpoints_setKeyExpiryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 key_expiry_disabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceEndpoints_setKeyExpiryArgs.__new__(DeviceEndpoints_setKeyExpiryArgs)

            __props__.__dict__["id"] = id
            if key_expiry_disabled is None and not opts.urn:
                raise TypeError("Missing required property 'key_expiry_disabled'")
            __props__.__dict__["key_expiry_disabled"] = key_expiry_disabled
        super(DeviceEndpoints_setKeyExpiry, __self__).__init__(
            'tailscale:device:DeviceEndpoints_setKeyExpiry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DeviceEndpoints_setKeyExpiry':
        """
        Get an existing DeviceEndpoints_setKeyExpiry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeviceEndpoints_setKeyExpiryArgs.__new__(DeviceEndpoints_setKeyExpiryArgs)

        __props__.__dict__["key_expiry_disabled"] = None
        return DeviceEndpoints_setKeyExpiry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "key_expiry_disabled")

