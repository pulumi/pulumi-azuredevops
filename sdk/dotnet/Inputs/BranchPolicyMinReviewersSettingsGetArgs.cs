// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BranchPolicyMinReviewersSettingsGetArgs : global::Pulumi.ResourceArgs
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
        /// 
        /// &gt; **Note:** If `on_push_reset_all_votes` is `true` then `on_push_reset_approved_votes` will be set to `true`. To enable `on_push_reset_approved_votes`, you need explicitly set `on_push_reset_all_votes` `false` or not configure.
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
        private InputList<Inputs.BranchPolicyMinReviewersSettingsScopeGetArgs>? _scopes;

        /// <summary>
        /// A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        /// </summary>
        public InputList<Inputs.BranchPolicyMinReviewersSettingsScopeGetArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.BranchPolicyMinReviewersSettingsScopeGetArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// Allow requesters to approve their own changes. Defaults to `false`.
        /// </summary>
        [Input("submitterCanVote")]
        public Input<bool>? SubmitterCanVote { get; set; }

        public BranchPolicyMinReviewersSettingsGetArgs()
        {
        }
        public static new BranchPolicyMinReviewersSettingsGetArgs Empty => new BranchPolicyMinReviewersSettingsGetArgs();
    }
}
