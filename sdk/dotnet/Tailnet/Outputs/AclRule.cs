// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet.Outputs
{

    [OutputType]
    public sealed class AclRule
    {
        /// <summary>
        /// Tailscale ACL rules are "default deny".
        /// So the only possible value for an ACL
        /// rule is `accept`.
        /// </summary>
        public readonly Pulumi.TailscaleNative.Tailnet.AclRuleAction Action;
        public readonly ImmutableArray<string> Ports;
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private AclRule(
            Pulumi.TailscaleNative.Tailnet.AclRuleAction action,

            ImmutableArray<string> ports,

            ImmutableArray<string> users)
        {
            Action = action;
            Ports = ports;
            Users = users;
        }
    }
}
