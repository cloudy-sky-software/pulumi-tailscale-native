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
from ._inputs import *

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 capabilities: pulumi.Input['KeyCapabilitiesArgs'],
                 expiry_seconds: pulumi.Input[int],
                 tailnet: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Key resource.
        """
        pulumi.set(__self__, "capabilities", capabilities)
        pulumi.set(__self__, "expiry_seconds", expiry_seconds)
        if tailnet is not None:
            pulumi.set(__self__, "tailnet", tailnet)

    @property
    @pulumi.getter
    def capabilities(self) -> pulumi.Input['KeyCapabilitiesArgs']:
        return pulumi.get(self, "capabilities")

    @capabilities.setter
    def capabilities(self, value: pulumi.Input['KeyCapabilitiesArgs']):
        pulumi.set(self, "capabilities", value)

    @property
    @pulumi.getter(name="expirySeconds")
    def expiry_seconds(self) -> pulumi.Input[int]:
        return pulumi.get(self, "expiry_seconds")

    @expiry_seconds.setter
    def expiry_seconds(self, value: pulumi.Input[int]):
        pulumi.set(self, "expiry_seconds", value)

    @property
    @pulumi.getter
    def tailnet(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tailnet")

    @tailnet.setter
    def tailnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tailnet", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capabilities: Optional[pulumi.Input[pulumi.InputType['KeyCapabilitiesArgs']]] = None,
                 expiry_seconds: Optional[pulumi.Input[int]] = None,
                 tailnet: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Key resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Key resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capabilities: Optional[pulumi.Input[pulumi.InputType['KeyCapabilitiesArgs']]] = None,
                 expiry_seconds: Optional[pulumi.Input[int]] = None,
                 tailnet: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            if capabilities is None and not opts.urn:
                raise TypeError("Missing required property 'capabilities'")
            __props__.__dict__["capabilities"] = capabilities
            if expiry_seconds is None and not opts.urn:
                raise TypeError("Missing required property 'expiry_seconds'")
            __props__.__dict__["expiry_seconds"] = expiry_seconds
            __props__.__dict__["tailnet"] = tailnet
            __props__.__dict__["created"] = None
            __props__.__dict__["expires"] = None
            __props__.__dict__["key"] = None
        super(Key, __self__).__init__(
            'tailscale:tailnet:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["capabilities"] = None
        __props__.__dict__["created"] = None
        __props__.__dict__["expires"] = None
        __props__.__dict__["expiry_seconds"] = None
        __props__.__dict__["key"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capabilities(self) -> pulumi.Output['outputs.KeyCapabilities']:
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def expires(self) -> pulumi.Output[str]:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter(name="expirySeconds")
    def expiry_seconds(self) -> pulumi.Output[int]:
        return pulumi.get(self, "expiry_seconds")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key")

