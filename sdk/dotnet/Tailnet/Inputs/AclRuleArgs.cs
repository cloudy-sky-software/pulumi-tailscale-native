// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Tailnet.Inputs
{

    public sealed class AclRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tailscale ACL rules are "default deny".
        /// So the only possible value for an ACL
        /// rule is `accept`.
        /// </summary>
        [Input("action", required: true)]
        public Input<Pulumi.TailscaleNative.Tailnet.AclRuleAction> Action { get; set; } = null!;

        [Input("ports", required: true)]
        private InputList<string>? _ports;
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        [Input("users", required: true)]
        private InputList<string>? _users;
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public AclRuleArgs()
        {
        }
        public static new AclRuleArgs Empty => new AclRuleArgs();
    }
}
