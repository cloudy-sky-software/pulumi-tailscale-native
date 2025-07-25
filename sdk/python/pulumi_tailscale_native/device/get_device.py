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
    'GetDeviceResult',
    'AwaitableGetDeviceResult',
    'get_device',
    'get_device_output',
]

@pulumi.output_type
class GetDeviceResult:
    def __init__(__self__, addresses=None, advertised_routes=None, authorized=None, blocks_incoming_connections=None, client_connectivity=None, client_version=None, created=None, enabled_routes=None, expires=None, hostname=None, id=None, is_external=None, key_expiry_disabled=None, last_seen=None, machine_key=None, name=None, node_key=None, os=None, update_available=None, user=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if advertised_routes and not isinstance(advertised_routes, list):
            raise TypeError("Expected argument 'advertised_routes' to be a list")
        pulumi.set(__self__, "advertised_routes", advertised_routes)
        if authorized and not isinstance(authorized, bool):
            raise TypeError("Expected argument 'authorized' to be a bool")
        pulumi.set(__self__, "authorized", authorized)
        if blocks_incoming_connections and not isinstance(blocks_incoming_connections, bool):
            raise TypeError("Expected argument 'blocks_incoming_connections' to be a bool")
        pulumi.set(__self__, "blocks_incoming_connections", blocks_incoming_connections)
        if client_connectivity and not isinstance(client_connectivity, dict):
            raise TypeError("Expected argument 'client_connectivity' to be a dict")
        pulumi.set(__self__, "client_connectivity", client_connectivity)
        if client_version and not isinstance(client_version, str):
            raise TypeError("Expected argument 'client_version' to be a str")
        pulumi.set(__self__, "client_version", client_version)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if enabled_routes and not isinstance(enabled_routes, list):
            raise TypeError("Expected argument 'enabled_routes' to be a list")
        pulumi.set(__self__, "enabled_routes", enabled_routes)
        if expires and not isinstance(expires, str):
            raise TypeError("Expected argument 'expires' to be a str")
        pulumi.set(__self__, "expires", expires)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_external and not isinstance(is_external, bool):
            raise TypeError("Expected argument 'is_external' to be a bool")
        pulumi.set(__self__, "is_external", is_external)
        if key_expiry_disabled and not isinstance(key_expiry_disabled, bool):
            raise TypeError("Expected argument 'key_expiry_disabled' to be a bool")
        pulumi.set(__self__, "key_expiry_disabled", key_expiry_disabled)
        if last_seen and not isinstance(last_seen, str):
            raise TypeError("Expected argument 'last_seen' to be a str")
        pulumi.set(__self__, "last_seen", last_seen)
        if machine_key and not isinstance(machine_key, str):
            raise TypeError("Expected argument 'machine_key' to be a str")
        pulumi.set(__self__, "machine_key", machine_key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_key and not isinstance(node_key, str):
            raise TypeError("Expected argument 'node_key' to be a str")
        pulumi.set(__self__, "node_key", node_key)
        if os and not isinstance(os, str):
            raise TypeError("Expected argument 'os' to be a str")
        pulumi.set(__self__, "os", os)
        if update_available and not isinstance(update_available, bool):
            raise TypeError("Expected argument 'update_available' to be a bool")
        pulumi.set(__self__, "update_available", update_available)
        if user and not isinstance(user, str):
            raise TypeError("Expected argument 'user' to be a str")
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="advertisedRoutes")
    def advertised_routes(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "advertised_routes")

    @property
    @pulumi.getter
    def authorized(self) -> builtins.bool:
        return pulumi.get(self, "authorized")

    @property
    @pulumi.getter(name="blocksIncomingConnections")
    def blocks_incoming_connections(self) -> builtins.bool:
        return pulumi.get(self, "blocks_incoming_connections")

    @property
    @pulumi.getter(name="clientConnectivity")
    def client_connectivity(self) -> 'outputs.ClientConnectivity':
        return pulumi.get(self, "client_connectivity")

    @property
    @pulumi.getter(name="clientVersion")
    def client_version(self) -> builtins.str:
        return pulumi.get(self, "client_version")

    @property
    @pulumi.getter
    def created(self) -> builtins.str:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="enabledRoutes")
    def enabled_routes(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "enabled_routes")

    @property
    @pulumi.getter
    def expires(self) -> builtins.str:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def hostname(self) -> builtins.str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isExternal")
    def is_external(self) -> builtins.bool:
        return pulumi.get(self, "is_external")

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> builtins.bool:
        return pulumi.get(self, "key_expiry_disabled")

    @property
    @pulumi.getter(name="lastSeen")
    def last_seen(self) -> builtins.str:
        return pulumi.get(self, "last_seen")

    @property
    @pulumi.getter(name="machineKey")
    def machine_key(self) -> builtins.str:
        return pulumi.get(self, "machine_key")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeKey")
    def node_key(self) -> builtins.str:
        return pulumi.get(self, "node_key")

    @property
    @pulumi.getter
    def os(self) -> builtins.str:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="updateAvailable")
    def update_available(self) -> builtins.bool:
        return pulumi.get(self, "update_available")

    @property
    @pulumi.getter
    def user(self) -> builtins.str:
        return pulumi.get(self, "user")


class AwaitableGetDeviceResult(GetDeviceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceResult(
            addresses=self.addresses,
            advertised_routes=self.advertised_routes,
            authorized=self.authorized,
            blocks_incoming_connections=self.blocks_incoming_connections,
            client_connectivity=self.client_connectivity,
            client_version=self.client_version,
            created=self.created,
            enabled_routes=self.enabled_routes,
            expires=self.expires,
            hostname=self.hostname,
            id=self.id,
            is_external=self.is_external,
            key_expiry_disabled=self.key_expiry_disabled,
            last_seen=self.last_seen,
            machine_key=self.machine_key,
            name=self.name,
            node_key=self.node_key,
            os=self.os,
            update_available=self.update_available,
            user=self.user)


def get_device(id: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale-native:device:getDevice', __args__, opts=opts, typ=GetDeviceResult).value

    return AwaitableGetDeviceResult(
        addresses=pulumi.get(__ret__, 'addresses'),
        advertised_routes=pulumi.get(__ret__, 'advertised_routes'),
        authorized=pulumi.get(__ret__, 'authorized'),
        blocks_incoming_connections=pulumi.get(__ret__, 'blocks_incoming_connections'),
        client_connectivity=pulumi.get(__ret__, 'client_connectivity'),
        client_version=pulumi.get(__ret__, 'client_version'),
        created=pulumi.get(__ret__, 'created'),
        enabled_routes=pulumi.get(__ret__, 'enabled_routes'),
        expires=pulumi.get(__ret__, 'expires'),
        hostname=pulumi.get(__ret__, 'hostname'),
        id=pulumi.get(__ret__, 'id'),
        is_external=pulumi.get(__ret__, 'is_external'),
        key_expiry_disabled=pulumi.get(__ret__, 'key_expiry_disabled'),
        last_seen=pulumi.get(__ret__, 'last_seen'),
        machine_key=pulumi.get(__ret__, 'machine_key'),
        name=pulumi.get(__ret__, 'name'),
        node_key=pulumi.get(__ret__, 'node_key'),
        os=pulumi.get(__ret__, 'os'),
        update_available=pulumi.get(__ret__, 'update_available'),
        user=pulumi.get(__ret__, 'user'))
def get_device_output(id: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDeviceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('tailscale-native:device:getDevice', __args__, opts=opts, typ=GetDeviceResult)
    return __ret__.apply(lambda __response__: GetDeviceResult(
        addresses=pulumi.get(__response__, 'addresses'),
        advertised_routes=pulumi.get(__response__, 'advertised_routes'),
        authorized=pulumi.get(__response__, 'authorized'),
        blocks_incoming_connections=pulumi.get(__response__, 'blocks_incoming_connections'),
        client_connectivity=pulumi.get(__response__, 'client_connectivity'),
        client_version=pulumi.get(__response__, 'client_version'),
        created=pulumi.get(__response__, 'created'),
        enabled_routes=pulumi.get(__response__, 'enabled_routes'),
        expires=pulumi.get(__response__, 'expires'),
        hostname=pulumi.get(__response__, 'hostname'),
        id=pulumi.get(__response__, 'id'),
        is_external=pulumi.get(__response__, 'is_external'),
        key_expiry_disabled=pulumi.get(__response__, 'key_expiry_disabled'),
        last_seen=pulumi.get(__response__, 'last_seen'),
        machine_key=pulumi.get(__response__, 'machine_key'),
        name=pulumi.get(__response__, 'name'),
        node_key=pulumi.get(__response__, 'node_key'),
        os=pulumi.get(__response__, 'os'),
        update_available=pulumi.get(__response__, 'update_available'),
        user=pulumi.get(__response__, 'user')))
