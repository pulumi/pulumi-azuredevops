// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Policy.Inputs
{

    public sealed class BranchPolicyBuildValidationSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the build to monitor for the policy.
        /// </summary>
        [Input("buildDefinitionId", required: true)]
        public Input<int> BuildDefinitionId { get; set; } = null!;

        /// <summary>
        /// The display name for the policy.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// If set to true, the build will need to be manually queued. Defaults to `false`
        /// </summary>
        [Input("manualQueueOnly")]
        public Input<bool>? ManualQueueOnly { get; set; }

        /// <summary>
        /// True if the build should queue on source updates only. Defaults to `true`.
        /// </summary>
        [Input("queueOnSourceUpdateOnly")]
        public Input<bool>? QueueOnSourceUpdateOnly { get; set; }

        [Input("scopes", required: true)]
        private InputList<Inputs.BranchPolicyBuildValidationSettingsScopeArgs>? _scopes;

        /// <summary>
        /// Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        /// </summary>
        public InputList<Inputs.BranchPolicyBuildValidationSettingsScopeArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.BranchPolicyBuildValidationSettingsScopeArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
        /// </summary>
        [Input("validDuration")]
        public Input<int>? ValidDuration { get; set; }

        public BranchPolicyBuildValidationSettingsArgs()
        {
        }
    }
}
