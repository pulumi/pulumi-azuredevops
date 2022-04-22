// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Build.Inputs
{

    public sealed class BuildDefinitionPullRequestTriggerOverrideGetArgs : Pulumi.ResourceArgs
    {
        [Input("autoCancel")]
        public Input<bool>? AutoCancel { get; set; }

        [Input("branchFilters")]
        private InputList<Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterGetArgs>? _branchFilters;
        public InputList<Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterGetArgs> BranchFilters
        {
            get => _branchFilters ?? (_branchFilters = new InputList<Inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterGetArgs>());
            set => _branchFilters = value;
        }

        [Input("pathFilters")]
        private InputList<Inputs.BuildDefinitionPullRequestTriggerOverridePathFilterGetArgs>? _pathFilters;
        public InputList<Inputs.BuildDefinitionPullRequestTriggerOverridePathFilterGetArgs> PathFilters
        {
            get => _pathFilters ?? (_pathFilters = new InputList<Inputs.BuildDefinitionPullRequestTriggerOverridePathFilterGetArgs>());
            set => _pathFilters = value;
        }

        public BuildDefinitionPullRequestTriggerOverrideGetArgs()
        {
        }
    }
}
