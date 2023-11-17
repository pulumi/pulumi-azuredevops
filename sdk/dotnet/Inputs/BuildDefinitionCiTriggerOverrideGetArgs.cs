// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionCiTriggerOverrideGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
        /// </summary>
        [Input("batch")]
        public Input<bool>? Batch { get; set; }

        [Input("branchFilters")]
        private InputList<Inputs.BuildDefinitionCiTriggerOverrideBranchFilterGetArgs>? _branchFilters;

        /// <summary>
        /// The branches to include and exclude from the trigger.`branch_filter` - (Optional) The branches to include and exclude from the trigger.
        /// </summary>
        public InputList<Inputs.BuildDefinitionCiTriggerOverrideBranchFilterGetArgs> BranchFilters
        {
            get => _branchFilters ?? (_branchFilters = new InputList<Inputs.BuildDefinitionCiTriggerOverrideBranchFilterGetArgs>());
            set => _branchFilters = value;
        }

        /// <summary>
        /// The number of max builds per branch. Defaults to `1`.
        /// </summary>
        [Input("maxConcurrentBuildsPerBranch")]
        public Input<int>? MaxConcurrentBuildsPerBranch { get; set; }

        [Input("pathFilters")]
        private InputList<Inputs.BuildDefinitionCiTriggerOverridePathFilterGetArgs>? _pathFilters;

        /// <summary>
        /// Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.`path_filter` - (Optional) Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
        /// </summary>
        public InputList<Inputs.BuildDefinitionCiTriggerOverridePathFilterGetArgs> PathFilters
        {
            get => _pathFilters ?? (_pathFilters = new InputList<Inputs.BuildDefinitionCiTriggerOverridePathFilterGetArgs>());
            set => _pathFilters = value;
        }

        /// <summary>
        /// How often the external repository is polled. Defaults to `0`.
        /// </summary>
        [Input("pollingInterval")]
        public Input<int>? PollingInterval { get; set; }

        /// <summary>
        /// This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
        /// </summary>
        [Input("pollingJobId")]
        public Input<string>? PollingJobId { get; set; }

        public BuildDefinitionCiTriggerOverrideGetArgs()
        {
        }
        public static new BuildDefinitionCiTriggerOverrideGetArgs Empty => new BuildDefinitionCiTriggerOverrideGetArgs();
    }
}
