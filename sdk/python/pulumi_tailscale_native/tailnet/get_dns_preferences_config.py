# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'NameServersPreference',
    'AwaitableNameServersPreference',
    'get_dns_preferences_config',
    'get_dns_preferences_config_output',
]

@pulumi.output_type
class NameServersPreference:
    def __init__(__self__, magic_dns=None):
        if magic_dns and not isinstance(magic_dns, bool):
            raise TypeError("Expected argument 'magic_dns' to be a bool")
        pulumi.set(__self__, "magic_dns", magic_dns)

    @property
    @pulumi.getter(name="magicDNS")
    def magic_dns(self) -> builtins.bool:
        return pulumi.get(self, "magic_dns")


class AwaitableNameServersPreference(NameServersPreference):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return NameServersPreference(
            magic_dns=self.magic_dns)


def get_dns_preferences_config(tailnet: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableNameServersPreference:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    __args__ = dict()
    __args__['tailnet'] = tailnet
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale-native:tailnet:getDNSPreferencesConfig', __args__, opts=opts, typ=NameServersPreference).value

    return AwaitableNameServersPreference(
        magic_dns=pulumi.get(__ret__, 'magic_dns'))
def get_dns_preferences_config_output(tailnet: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[NameServersPreference]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str tailnet: For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value "-" to refer to the default tailnet of the authenticated user making the API call.
    """
    __args__ = dict()
    __args__['tailnet'] = tailnet
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('tailscale-native:tailnet:getDNSPreferencesConfig', __args__, opts=opts, typ=NameServersPreference)
    return __ret__.apply(lambda __response__: NameServersPreference(
        magic_dns=pulumi.get(__response__, 'magic_dns')))
