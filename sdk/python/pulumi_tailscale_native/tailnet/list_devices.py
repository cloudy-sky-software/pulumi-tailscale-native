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
    'ListDevicesResult',
    'AwaitableListDevicesResult',
    'list_devices',
    'list_devices_output',
]

@pulumi.output_type
class ListDevicesResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListDevicesProperties':
        return pulumi.get(self, "items")


class AwaitableListDevicesResult(ListDevicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDevicesResult(
            items=self.items)


def list_devices(tailnet: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDevicesResult:
    """
    Use this data source to access information about an existing resource.

    :param str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    __args__ = dict()
    __args__['tailnet'] = tailnet
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale-native:tailnet:listDevices', __args__, opts=opts, typ=ListDevicesResult).value

    return AwaitableListDevicesResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_devices)
def list_devices_output(tailnet: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDevicesResult]:
    """
    Use this data source to access information about an existing resource.

    :param str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    ...
