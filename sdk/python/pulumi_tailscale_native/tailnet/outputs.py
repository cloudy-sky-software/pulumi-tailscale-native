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
from ._enums import *

__all__ = [
    'Acl',
    'AclRule',
    'AuthKey',
    'ClientConnectivity',
    'ClientSupports',
    'CreateKey',
    'Device',
    'DeviceKeyCapabilities',
    'DnsSearchPaths',
    'KeyCapabilities',
    'ListDevicesProperties',
    'NameServers',
    'NameServersPreference',
    'NodeAttrs',
    'SshRule',
]

@pulumi.output_type
class Acl(dict):
    def __init__(__self__, *,
                 acls: Sequence['outputs.AclRule'],
                 auto_approvers: Any,
                 groups: Any,
                 hosts: Any,
                 node_attrs: Sequence['outputs.NodeAttrs'],
                 ssh: Sequence['outputs.SshRule'],
                 tag_owners: Any,
                 tests: str):
        pulumi.set(__self__, "acls", acls)
        pulumi.set(__self__, "auto_approvers", auto_approvers)
        pulumi.set(__self__, "groups", groups)
        pulumi.set(__self__, "hosts", hosts)
        pulumi.set(__self__, "node_attrs", node_attrs)
        pulumi.set(__self__, "ssh", ssh)
        pulumi.set(__self__, "tag_owners", tag_owners)
        pulumi.set(__self__, "tests", tests)

    @property
    @pulumi.getter
    def acls(self) -> Sequence['outputs.AclRule']:
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="autoApprovers")
    def auto_approvers(self) -> Any:
        return pulumi.get(self, "auto_approvers")

    @property
    @pulumi.getter
    def groups(self) -> Any:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def hosts(self) -> Any:
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter(name="nodeAttrs")
    def node_attrs(self) -> Sequence['outputs.NodeAttrs']:
        return pulumi.get(self, "node_attrs")

    @property
    @pulumi.getter
    def ssh(self) -> Sequence['outputs.SshRule']:
        return pulumi.get(self, "ssh")

    @property
    @pulumi.getter(name="tagOwners")
    def tag_owners(self) -> Any:
        return pulumi.get(self, "tag_owners")

    @property
    @pulumi.getter
    def tests(self) -> str:
        return pulumi.get(self, "tests")


@pulumi.output_type
class AclRule(dict):
    def __init__(__self__, *,
                 action: 'AclRuleAction',
                 ports: Sequence[str],
                 users: Sequence[str]):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def action(self) -> 'AclRuleAction':
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def ports(self) -> Sequence[str]:
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def users(self) -> Sequence[str]:
        return pulumi.get(self, "users")


@pulumi.output_type
class AuthKey(dict):
    def __init__(__self__, *,
                 expires: str,
                 key: str,
                 created: Optional[str] = None):
        pulumi.set(__self__, "expires", expires)
        pulumi.set(__self__, "key", key)
        if created is not None:
            pulumi.set(__self__, "created", created)

    @property
    @pulumi.getter
    def expires(self) -> str:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def created(self) -> Optional[str]:
        return pulumi.get(self, "created")


@pulumi.output_type
class ClientConnectivity(dict):
    def __init__(__self__, *,
                 client_supports: 'outputs.ClientSupports',
                 derp: str,
                 endpoints: str,
                 latency: Any,
                 mapping_varies_by_dest_ip: bool):
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
    def derp(self) -> str:
        return pulumi.get(self, "derp")

    @property
    @pulumi.getter
    def endpoints(self) -> str:
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def latency(self) -> Any:
        return pulumi.get(self, "latency")

    @property
    @pulumi.getter(name="mappingVariesByDestIP")
    def mapping_varies_by_dest_ip(self) -> bool:
        return pulumi.get(self, "mapping_varies_by_dest_ip")


@pulumi.output_type
class ClientSupports(dict):
    def __init__(__self__, *,
                 hair_pinning: bool,
                 ipv6: bool,
                 pcp: bool,
                 pmp: bool,
                 udp: bool,
                 upnp: bool):
        pulumi.set(__self__, "hair_pinning", hair_pinning)
        pulumi.set(__self__, "ipv6", ipv6)
        pulumi.set(__self__, "pcp", pcp)
        pulumi.set(__self__, "pmp", pmp)
        pulumi.set(__self__, "udp", udp)
        pulumi.set(__self__, "upnp", upnp)

    @property
    @pulumi.getter(name="hairPinning")
    def hair_pinning(self) -> bool:
        return pulumi.get(self, "hair_pinning")

    @property
    @pulumi.getter
    def ipv6(self) -> bool:
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def pcp(self) -> bool:
        return pulumi.get(self, "pcp")

    @property
    @pulumi.getter
    def pmp(self) -> bool:
        return pulumi.get(self, "pmp")

    @property
    @pulumi.getter
    def udp(self) -> bool:
        return pulumi.get(self, "udp")

    @property
    @pulumi.getter
    def upnp(self) -> bool:
        return pulumi.get(self, "upnp")


@pulumi.output_type
class CreateKey(dict):
    def __init__(__self__, *,
                 ephemeral: bool,
                 preauthorized: bool,
                 reusable: bool,
                 tags: Sequence[str]):
        pulumi.set(__self__, "ephemeral", ephemeral)
        pulumi.set(__self__, "preauthorized", preauthorized)
        pulumi.set(__self__, "reusable", reusable)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def ephemeral(self) -> bool:
        return pulumi.get(self, "ephemeral")

    @property
    @pulumi.getter
    def preauthorized(self) -> bool:
        return pulumi.get(self, "preauthorized")

    @property
    @pulumi.getter
    def reusable(self) -> bool:
        return pulumi.get(self, "reusable")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")


@pulumi.output_type
class Device(dict):
    def __init__(__self__, *,
                 addresses: Sequence[str],
                 advertised_routes: Sequence[str],
                 authorized: bool,
                 blocks_incoming_connections: bool,
                 client_connectivity: 'outputs.ClientConnectivity',
                 client_version: str,
                 created: str,
                 enabled_routes: Sequence[str],
                 expires: str,
                 hostname: str,
                 id: str,
                 is_external: bool,
                 key_expiry_disabled: bool,
                 last_seen: str,
                 machine_key: str,
                 name: str,
                 node_key: str,
                 os: str,
                 update_available: bool,
                 user: str):
        pulumi.set(__self__, "addresses", addresses)
        pulumi.set(__self__, "advertised_routes", advertised_routes)
        pulumi.set(__self__, "authorized", authorized)
        pulumi.set(__self__, "blocks_incoming_connections", blocks_incoming_connections)
        pulumi.set(__self__, "client_connectivity", client_connectivity)
        pulumi.set(__self__, "client_version", client_version)
        pulumi.set(__self__, "created", created)
        pulumi.set(__self__, "enabled_routes", enabled_routes)
        pulumi.set(__self__, "expires", expires)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_external", is_external)
        pulumi.set(__self__, "key_expiry_disabled", key_expiry_disabled)
        pulumi.set(__self__, "last_seen", last_seen)
        pulumi.set(__self__, "machine_key", machine_key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "node_key", node_key)
        pulumi.set(__self__, "os", os)
        pulumi.set(__self__, "update_available", update_available)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[str]:
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="advertisedRoutes")
    def advertised_routes(self) -> Sequence[str]:
        return pulumi.get(self, "advertised_routes")

    @property
    @pulumi.getter
    def authorized(self) -> bool:
        return pulumi.get(self, "authorized")

    @property
    @pulumi.getter(name="blocksIncomingConnections")
    def blocks_incoming_connections(self) -> bool:
        return pulumi.get(self, "blocks_incoming_connections")

    @property
    @pulumi.getter(name="clientConnectivity")
    def client_connectivity(self) -> 'outputs.ClientConnectivity':
        return pulumi.get(self, "client_connectivity")

    @property
    @pulumi.getter(name="clientVersion")
    def client_version(self) -> str:
        return pulumi.get(self, "client_version")

    @property
    @pulumi.getter
    def created(self) -> str:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="enabledRoutes")
    def enabled_routes(self) -> Sequence[str]:
        return pulumi.get(self, "enabled_routes")

    @property
    @pulumi.getter
    def expires(self) -> str:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isExternal")
    def is_external(self) -> bool:
        return pulumi.get(self, "is_external")

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> bool:
        return pulumi.get(self, "key_expiry_disabled")

    @property
    @pulumi.getter(name="lastSeen")
    def last_seen(self) -> str:
        return pulumi.get(self, "last_seen")

    @property
    @pulumi.getter(name="machineKey")
    def machine_key(self) -> str:
        return pulumi.get(self, "machine_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeKey")
    def node_key(self) -> str:
        return pulumi.get(self, "node_key")

    @property
    @pulumi.getter
    def os(self) -> str:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="updateAvailable")
    def update_available(self) -> bool:
        return pulumi.get(self, "update_available")

    @property
    @pulumi.getter
    def user(self) -> str:
        return pulumi.get(self, "user")


@pulumi.output_type
class DeviceKeyCapabilities(dict):
    def __init__(__self__, *,
                 create: 'outputs.CreateKey'):
        pulumi.set(__self__, "create", create)

    @property
    @pulumi.getter
    def create(self) -> 'outputs.CreateKey':
        return pulumi.get(self, "create")


@pulumi.output_type
class DnsSearchPaths(dict):
    def __init__(__self__, *,
                 search_paths: Sequence[str]):
        pulumi.set(__self__, "search_paths", search_paths)

    @property
    @pulumi.getter(name="searchPaths")
    def search_paths(self) -> Sequence[str]:
        return pulumi.get(self, "search_paths")


@pulumi.output_type
class KeyCapabilities(dict):
    def __init__(__self__, *,
                 devices: 'outputs.DeviceKeyCapabilities'):
        pulumi.set(__self__, "devices", devices)

    @property
    @pulumi.getter
    def devices(self) -> 'outputs.DeviceKeyCapabilities':
        return pulumi.get(self, "devices")


@pulumi.output_type
class ListDevicesProperties(dict):
    def __init__(__self__, *,
                 devices: Sequence['outputs.Device']):
        pulumi.set(__self__, "devices", devices)

    @property
    @pulumi.getter
    def devices(self) -> Sequence['outputs.Device']:
        return pulumi.get(self, "devices")


@pulumi.output_type
class NameServers(dict):
    def __init__(__self__, *,
                 dns: Sequence[str],
                 magic_dns: bool):
        pulumi.set(__self__, "dns", dns)
        pulumi.set(__self__, "magic_dns", magic_dns)

    @property
    @pulumi.getter
    def dns(self) -> Sequence[str]:
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter(name="magicDNS")
    def magic_dns(self) -> bool:
        return pulumi.get(self, "magic_dns")


@pulumi.output_type
class NameServersPreference(dict):
    def __init__(__self__, *,
                 magic_dns: bool):
        pulumi.set(__self__, "magic_dns", magic_dns)

    @property
    @pulumi.getter(name="magicDNS")
    def magic_dns(self) -> bool:
        return pulumi.get(self, "magic_dns")


@pulumi.output_type
class NodeAttrs(dict):
    def __init__(__self__, *,
                 attr: Sequence[str],
                 target: Sequence[str]):
        pulumi.set(__self__, "attr", attr)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def attr(self) -> Sequence[str]:
        return pulumi.get(self, "attr")

    @property
    @pulumi.getter
    def target(self) -> Sequence[str]:
        return pulumi.get(self, "target")


@pulumi.output_type
class SshRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "checkPeriod":
            suggest = "check_period"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SshRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SshRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SshRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: 'SshRuleAction',
                 check_period: str,
                 dst: Sequence[str],
                 src: Sequence[str],
                 users: Sequence[str]):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "check_period", check_period)
        pulumi.set(__self__, "dst", dst)
        pulumi.set(__self__, "src", src)
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def action(self) -> 'SshRuleAction':
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="checkPeriod")
    def check_period(self) -> str:
        return pulumi.get(self, "check_period")

    @property
    @pulumi.getter
    def dst(self) -> Sequence[str]:
        return pulumi.get(self, "dst")

    @property
    @pulumi.getter
    def src(self) -> Sequence[str]:
        return pulumi.get(self, "src")

    @property
    @pulumi.getter
    def users(self) -> Sequence[str]:
        return pulumi.get(self, "users")


