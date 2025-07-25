// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class GetBuildDefinitionPullRequestTriggerOverrideBranchFilterResult
    {
        /// <summary>
        /// (Optional) List of path patterns to exclude.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// (Optional) List of path patterns to include.
        /// </summary>
        public readonly ImmutableArray<string> Includes;

        [OutputConstructor]
        private GetBuildDefinitionPullRequestTriggerOverrideBranchFilterResult(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes)
        {
            Excludes = excludes;
            Includes = includes;
        }
    }
}
