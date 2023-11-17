// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Build.Outputs
{

    [OutputType]
    public sealed class BuildDefinitionScheduleBranchFilter
    {
        /// <summary>
        /// List of branch patterns to exclude.
        /// 
        /// 
        /// `exclude` - (Optional) List of branch patterns to exclude.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// List of branch patterns to include.`include` - (Optional) List of branch patterns to include.
        /// </summary>
        public readonly ImmutableArray<string> Includes;

        [OutputConstructor]
        private BuildDefinitionScheduleBranchFilter(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes)
        {
            Excludes = excludes;
            Includes = includes;
        }
    }
}
