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

__all__ = [
    'GetDeviceRoutesResult',
    'AwaitableGetDeviceRoutesResult',
    'get_device_routes',
    'get_device_routes_output',
]

@pulumi.output_type
class GetDeviceRoutesResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.DeviceRoutes':
        return pulumi.get(self, "items")


class AwaitableGetDeviceRoutesResult(GetDeviceRoutesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceRoutesResult(
            items=self.items)


def get_device_routes(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceRoutesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale:device:getDeviceRoutes', __args__, opts=opts, typ=GetDeviceRoutesResult).value

    return AwaitableGetDeviceRoutesResult(
        items=__ret__.items)


@_utilities.lift_output_func(get_device_routes)
def get_device_routes_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeviceRoutesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
