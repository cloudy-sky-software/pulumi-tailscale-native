# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetNameServersResult',
    'AwaitableGetNameServersResult',
    'get_name_servers',
    'get_name_servers_output',
]

@pulumi.output_type
class GetNameServersResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.NameServers':
        return pulumi.get(self, "items")


class AwaitableGetNameServersResult(GetNameServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNameServersResult(
            items=self.items)


def get_name_servers(tailnet: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNameServersResult:
    """
    Use this data source to access information about an existing resource.

    :param str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    __args__ = dict()
    __args__['tailnet'] = tailnet
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale-native:tailnet:getNameServers', __args__, opts=opts, typ=GetNameServersResult).value

    return AwaitableGetNameServersResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_name_servers)
def get_name_servers_output(tailnet: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNameServersResult]:
    """
    Use this data source to access information about an existing resource.

    :param str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    ...
