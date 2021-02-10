// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Policy.Inputs
{

    public sealed class BranchPolicyMinReviewersSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
        /// </summary>
        [Input("allowCompletionWithRejectsOrWaits")]
        public Input<bool>? AllowCompletionWithRejectsOrWaits { get; set; }

        /// <summary>
        /// Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
        /// </summary>
        [Input("lastPusherCannotApprove")]
        public Input<bool>? LastPusherCannotApprove { get; set; }

        /// <summary>
        /// On last iteration require vote. Defaults to `false`.
        /// </summary>
        [Input("onLastIterationRequireVote")]
        public Input<bool>? OnLastIterationRequireVote { get; set; }

        /// <summary>
        /// When new changes are pushed reset all code reviewer votes. Defaults to `false`.
        /// </summary>
        [Input("onPushResetAllVotes")]
        public Input<bool>? OnPushResetAllVotes { get; set; }

        /// <summary>
        /// When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
        /// </summary>
        [Input("onPushResetApprovedVotes")]
        public Input<bool>? OnPushResetApprovedVotes { get; set; }

        /// <summary>
        /// The number of reviewers needed to approve.
        /// </summary>
        [Input("reviewerCount")]
        public Input<int>? ReviewerCount { get; set; }

        [Input("scopes", required: true)]
        private InputList<Inputs.BranchPolicyMinReviewersSettingsScopeArgs>? _scopes;

        /// <summary>
        /// Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        /// </summary>
        public InputList<Inputs.BranchPolicyMinReviewersSettingsScopeArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.BranchPolicyMinReviewersSettingsScopeArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// Allow requesters to approve their own changes. Defaults to `false`.
        /// </summary>
        [Input("submitterCanVote")]
        public Input<bool>? SubmitterCanVote { get; set; }

        public BranchPolicyMinReviewersSettingsArgs()
        {
        }
    }
}
