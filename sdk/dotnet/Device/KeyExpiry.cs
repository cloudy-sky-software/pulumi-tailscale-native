// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.TailscaleNative.Device
{
    [TailscaleNativeResourceType("tailscale-native:device:KeyExpiry")]
    public partial class KeyExpiry : global::Pulumi.CustomResource
    {
        [Output("keyExpiryDisabled")]
        public Output<bool> KeyExpiryDisabled { get; private set; } = null!;


        /// <summary>
        /// Create a KeyExpiry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyExpiry(string name, KeyExpiryArgs args, CustomResourceOptions? options = null)
            : base("tailscale-native:device:KeyExpiry", name, args ?? new KeyExpiryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyExpiry(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("tailscale-native:device:KeyExpiry", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-tailscale-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KeyExpiry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyExpiry Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KeyExpiry(name, id, options);
        }
    }

    public sealed class KeyExpiryArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("keyExpiryDisabled", required: true)]
        public Input<bool> KeyExpiryDisabled { get; set; } = null!;

        public KeyExpiryArgs()
        {
        }
        public static new KeyExpiryArgs Empty => new KeyExpiryArgs();
    }
}
