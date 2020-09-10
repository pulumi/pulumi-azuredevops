// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class GitInitialization
    {
        /// <summary>
        /// The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`. Defaults to `Uninitialized`.
        /// </summary>
        public readonly string InitType;
        /// <summary>
        /// Type of the source repository. Used if the `init_type` is `Import`. Valid values: `Git`. Defaults to `Git`.
        /// </summary>
        public readonly string? SourceType;
        /// <summary>
        /// The URL of the source repository. Used if the `init_type` is `Import`.
        /// </summary>
        public readonly string? SourceUrl;

        [OutputConstructor]
        private GitInitialization(
            string initType,

            string? sourceType,

            string? sourceUrl)
        {
            InitType = initType;
            SourceType = sourceType;
            SourceUrl = sourceUrl;
        }
    }
}
