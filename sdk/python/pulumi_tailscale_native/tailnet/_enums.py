# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'AclRuleAction',
    'SshRuleAction',
]


class AclRuleAction(builtins.str, Enum):
    """
    Tailscale ACL rules are "default deny".
    So the only possible value for an ACL
    rule is `accept`.
    """
    ACCEPT = "accept"


class SshRuleAction(builtins.str, Enum):
    CHECK = "check"
    ACCEPT = "accept"
