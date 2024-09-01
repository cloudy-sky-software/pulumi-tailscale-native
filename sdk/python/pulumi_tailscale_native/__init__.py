# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_tailscale_native.config as __config
    config = __config
    import pulumi_tailscale_native.device as __device
    device = __device
    import pulumi_tailscale_native.tailnet as __tailnet
    tailnet = __tailnet
else:
    config = _utilities.lazy_import('pulumi_tailscale_native.config')
    device = _utilities.lazy_import('pulumi_tailscale_native.device')
    tailnet = _utilities.lazy_import('pulumi_tailscale_native.tailnet')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "tailscale-native",
  "mod": "device",
  "fqn": "pulumi_tailscale_native.device",
  "classes": {
   "tailscale-native:device:AuthorizeDevice": "AuthorizeDevice",
   "tailscale-native:device:KeyExpiry": "KeyExpiry",
   "tailscale-native:device:RoutesConfig": "RoutesConfig",
   "tailscale-native:device:Tags": "Tags"
  }
 },
 {
  "pkg": "tailscale-native",
  "mod": "tailnet",
  "fqn": "pulumi_tailscale_native.tailnet",
  "classes": {
   "tailscale-native:tailnet:Acl": "Acl",
   "tailscale-native:tailnet:DNSPreferencesConfig": "DNSPreferencesConfig",
   "tailscale-native:tailnet:Key": "Key",
   "tailscale-native:tailnet:NameServersConfig": "NameServersConfig",
   "tailscale-native:tailnet:ReplaceSearchPaths": "ReplaceSearchPaths"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "tailscale-native",
  "token": "pulumi:providers:tailscale-native",
  "fqn": "pulumi_tailscale_native",
  "class": "Provider"
 }
]
"""
)
