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
    public sealed class BuildDefinitionPullRequestTriggerOverride
    {
        public readonly bool? AutoCancel;
        public readonly ImmutableArray<Outputs.BuildDefinitionPullRequestTriggerOverrideBranchFilter> BranchFilters;
        public readonly ImmutableArray<Outputs.BuildDefinitionPullRequestTriggerOverridePathFilter> PathFilters;

        [OutputConstructor]
        private BuildDefinitionPullRequestTriggerOverride(
            bool? autoCancel,

            ImmutableArray<Outputs.BuildDefinitionPullRequestTriggerOverrideBranchFilter> branchFilters,

            ImmutableArray<Outputs.BuildDefinitionPullRequestTriggerOverridePathFilter> pathFilters)
        {
            AutoCancel = autoCancel;
            BranchFilters = branchFilters;
            PathFilters = pathFilters;
        }
    }
}
