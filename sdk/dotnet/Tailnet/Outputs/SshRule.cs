// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Tailnet.Outputs
{

    [OutputType]
    public sealed class SshRule
    {
        public readonly Pulumi.Tailscale.Tailnet.SshRuleAction Action;
        public readonly string CheckPeriod;
        public readonly ImmutableArray<string> Dst;
        public readonly ImmutableArray<string> Src;
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private SshRule(
            Pulumi.Tailscale.Tailnet.SshRuleAction action,

            string checkPeriod,

            ImmutableArray<string> dst,

            ImmutableArray<string> src,

            ImmutableArray<string> users)
        {
            Action = action;
            CheckPeriod = checkPeriod;
            Dst = dst;
            Src = src;
            Users = users;
        }
    }
}
