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
    'GetRoutesConfigResult',
    'AwaitableGetRoutesConfigResult',
    'get_routes_config',
    'get_routes_config_output',
]

@pulumi.output_type
class GetRoutesConfigResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.DeviceRoutes':
        return pulumi.get(self, "items")


class AwaitableGetRoutesConfigResult(GetRoutesConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoutesConfigResult(
            items=self.items)


def get_routes_config(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoutesConfigResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale-native:device:getRoutesConfig', __args__, opts=opts, typ=GetRoutesConfigResult).value

    return AwaitableGetRoutesConfigResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_routes_config)
def get_routes_config_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoutesConfigResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...