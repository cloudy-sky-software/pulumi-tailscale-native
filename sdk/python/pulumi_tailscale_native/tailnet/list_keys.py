# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from . import outputs

__all__ = [
    'ListKeysResult',
    'AwaitableListKeysResult',
    'list_keys',
    'list_keys_output',
]

@pulumi.output_type
class ListKeysResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.AuthKey']:
        return pulumi.get(self, "items")


class AwaitableListKeysResult(ListKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListKeysResult(
            items=self.items)


def list_keys(tailnet: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    __args__ = dict()
    __args__['tailnet'] = tailnet
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale-native:tailnet:listKeys', __args__, opts=opts, typ=ListKeysResult).value

    return AwaitableListKeysResult(
        items=pulumi.get(__ret__, 'items'))
def list_keys_output(tailnet: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListKeysResult]:
    """
    Use this data source to access information about an existing resource.

    :param str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    __args__ = dict()
    __args__['tailnet'] = tailnet
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('tailscale-native:tailnet:listKeys', __args__, opts=opts, typ=ListKeysResult)
    return __ret__.apply(lambda __response__: ListKeysResult(
        items=pulumi.get(__response__, 'items')))
