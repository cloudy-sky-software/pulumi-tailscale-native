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
    public sealed class AuthKey
    {
        public readonly string? Created;
        public readonly string Expires;
        public readonly string Key;

        [OutputConstructor]
        private AuthKey(
            string? created,

            string expires,

            string key)
        {
            Created = created;
            Expires = expires;
            Key = key;
        }
    }
}
