// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class FeedFeatureGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if Feed should be Permanently removed, Defaults to `false`
        /// </summary>
        [Input("permanentDelete")]
        public Input<bool>? PermanentDelete { get; set; }

        /// <summary>
        /// Determines if Feed should be Restored during creation (if possible), Defaults to `false`
        /// </summary>
        [Input("restore")]
        public Input<bool>? Restore { get; set; }

        public FeedFeatureGetArgs()
        {
        }
        public static new FeedFeatureGetArgs Empty => new FeedFeatureGetArgs();
    }
}
