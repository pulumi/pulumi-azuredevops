// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class GitInitializationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`. Defaults to `Uninitialized`.
        /// </summary>
        [Input("initType", required: true)]
        public Input<string> InitType { get; set; } = null!;

        /// <summary>
        /// Type of the source repository. Used if the `init_type` is `Import`. Valid values: `Git`.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// The URL of the source repository. Used if the `init_type` is `Import`.
        /// </summary>
        [Input("sourceUrl")]
        public Input<string>? SourceUrl { get; set; }

        public GitInitializationGetArgs()
        {
        }
    }
}
