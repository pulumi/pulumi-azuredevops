// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Build.Inputs
{

    public sealed class BuildDefinitionPullRequestTriggerOverrideArgs : Pulumi.ResourceArgs
    {
        [Input("autoCancel")]
        public Input<bool>? AutoCancel { get; set; }

        [Input("branchFilters")]
        private InputList<Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs>? _branchFilters;
        public InputList<Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs> BranchFilters
        {
            get => _branchFilters ?? (_branchFilters = new InputList<Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs>());
            set => _branchFilters = value;
        }

        [Input("pathFilters")]
        private InputList<Inputs.BuildDefinitionPullRequestTriggerOverridePathFilterArgs>? _pathFilters;
        public InputList<Inputs.BuildDefinitionPullRequestTriggerOverridePathFilterArgs> PathFilters
        {
            get => _pathFilters ?? (_pathFilters = new InputList<Inputs.BuildDefinitionPullRequestTriggerOverridePathFilterArgs>());
            set => _pathFilters = value;
        }

        public BuildDefinitionPullRequestTriggerOverrideArgs()
        {
        }
    }
}
