# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_tailscale.config as __config
    config = __config
    import pulumi_tailscale.device as __device
    device = __device
    import pulumi_tailscale.tailnet as __tailnet
    tailnet = __tailnet
else:
    config = _utilities.lazy_import('pulumi_tailscale.config')
    device = _utilities.lazy_import('pulumi_tailscale.device')
    tailnet = _utilities.lazy_import('pulumi_tailscale.tailnet')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "tailscale",
  "mod": "device",
  "fqn": "pulumi_tailscale.device",
  "classes": {
   "tailscale:device:DeviceEndpoints_authorizeDevice": "DeviceEndpoints_authorizeDevice",
   "tailscale:device:DeviceEndpoints_setKeyExpiry": "DeviceEndpoints_setKeyExpiry",
   "tailscale:device:DeviceEndpoints_setRoutes": "DeviceEndpoints_setRoutes",
   "tailscale:device:DeviceEndpoints_setTags": "DeviceEndpoints_setTags"
  }
 },
 {
  "pkg": "tailscale",
  "mod": "tailnet",
  "fqn": "pulumi_tailscale.tailnet",
  "classes": {
   "tailscale:tailnet:Acl": "Acl",
   "tailscale:tailnet:DNSPreferences": "DNSPreferences",
   "tailscale:tailnet:NameServers": "NameServers",
   "tailscale:tailnet:TailnetEndpoints_replaceSearchPaths": "TailnetEndpoints_replaceSearchPaths"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "tailscale",
  "token": "pulumi:providers:tailscale",
  "fqn": "pulumi_tailscale",
  "class": "Provider"
 }
]
"""
)
