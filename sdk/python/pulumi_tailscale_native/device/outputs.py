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
from . import outputs

__all__ = [
    'ClientConnectivity',
    'ClientSupports',
]

@pulumi.output_type
class ClientConnectivity(dict):
    def __init__(__self__, *,
                 client_supports: 'outputs.ClientSupports',
                 derp: builtins.str,
                 endpoints: builtins.str,
                 latency: Any,
                 mapping_varies_by_dest_ip: builtins.bool):
        pulumi.set(__self__, "client_supports", client_supports)
        pulumi.set(__self__, "derp", derp)
        pulumi.set(__self__, "endpoints", endpoints)
        pulumi.set(__self__, "latency", latency)
        pulumi.set(__self__, "mapping_varies_by_dest_ip", mapping_varies_by_dest_ip)

    @property
    @pulumi.getter(name="clientSupports")
    def client_supports(self) -> 'outputs.ClientSupports':
        return pulumi.get(self, "client_supports")

    @property
    @pulumi.getter
    def derp(self) -> builtins.str:
        return pulumi.get(self, "derp")

    @property
    @pulumi.getter
    def endpoints(self) -> builtins.str:
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def latency(self) -> Any:
        return pulumi.get(self, "latency")

    @property
    @pulumi.getter(name="mappingVariesByDestIP")
    def mapping_varies_by_dest_ip(self) -> builtins.bool:
        return pulumi.get(self, "mapping_varies_by_dest_ip")


@pulumi.output_type
class ClientSupports(dict):
    def __init__(__self__, *,
                 hair_pinning: builtins.bool,
                 ipv6: builtins.bool,
                 pcp: builtins.bool,
                 pmp: builtins.bool,
                 udp: builtins.bool,
                 upnp: builtins.bool):
        pulumi.set(__self__, "hair_pinning", hair_pinning)
        pulumi.set(__self__, "ipv6", ipv6)
        pulumi.set(__self__, "pcp", pcp)
        pulumi.set(__self__, "pmp", pmp)
        pulumi.set(__self__, "udp", udp)
        pulumi.set(__self__, "upnp", upnp)

    @property
    @pulumi.getter(name="hairPinning")
    def hair_pinning(self) -> builtins.bool:
        return pulumi.get(self, "hair_pinning")

    @property
    @pulumi.getter
    def ipv6(self) -> builtins.bool:
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def pcp(self) -> builtins.bool:
        return pulumi.get(self, "pcp")

    @property
    @pulumi.getter
    def pmp(self) -> builtins.bool:
        return pulumi.get(self, "pmp")

    @property
    @pulumi.getter
    def udp(self) -> builtins.bool:
        return pulumi.get(self, "udp")

    @property
    @pulumi.getter
    def upnp(self) -> builtins.bool:
        return pulumi.get(self, "upnp")


