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
    public sealed class BuildDefinitionBuildCompletionTrigger
    {
        /// <summary>
        /// The branches to include and exclude from the trigger. A `branch_filter` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.BuildDefinitionBuildCompletionTriggerBranchFilter> BranchFilters;
        /// <summary>
        /// The ID of the build pipeline will be triggered.
        /// </summary>
        public readonly int BuildDefinitionId;

        [OutputConstructor]
        private BuildDefinitionBuildCompletionTrigger(
            ImmutableArray<Outputs.BuildDefinitionBuildCompletionTriggerBranchFilter> branchFilters,

            int buildDefinitionId)
        {
            BranchFilters = branchFilters;
            BuildDefinitionId = buildDefinitionId;
        }
    }
}
