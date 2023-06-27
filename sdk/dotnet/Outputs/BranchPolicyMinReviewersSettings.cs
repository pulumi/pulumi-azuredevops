// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class BranchPolicyMinReviewersSettings
    {
        /// <summary>
        /// Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
        /// </summary>
        public readonly bool? AllowCompletionWithRejectsOrWaits;
        /// <summary>
        /// Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
        /// </summary>
        public readonly bool? LastPusherCannotApprove;
        /// <summary>
        /// On last iteration require vote. Defaults to `false`.
        /// </summary>
        public readonly bool? OnLastIterationRequireVote;
        /// <summary>
        /// When new changes are pushed reset all code reviewer votes. Defaults to `false`.
        /// 
        /// &gt; **Note:** If `on_push_reset_all_votes` is `true` then `on_push_reset_approved_votes` will be set to `true`. To enable `on_push_reset_approved_votes`, you need explicitly set `on_push_reset_all_votes` `false` or not configure.
        /// </summary>
        public readonly bool? OnPushResetAllVotes;
        /// <summary>
        /// When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
        /// </summary>
        public readonly bool? OnPushResetApprovedVotes;
        /// <summary>
        /// The number of reviewers needed to approve.
        /// </summary>
        public readonly int? ReviewerCount;
        /// <summary>
        /// A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        /// </summary>
        public readonly ImmutableArray<Outputs.BranchPolicyMinReviewersSettingsScope> Scopes;
        /// <summary>
        /// Allow requesters to approve their own changes. Defaults to `false`.
        /// </summary>
        public readonly bool? SubmitterCanVote;

        [OutputConstructor]
        private BranchPolicyMinReviewersSettings(
            bool? allowCompletionWithRejectsOrWaits,

            bool? lastPusherCannotApprove,

            bool? onLastIterationRequireVote,

            bool? onPushResetAllVotes,

            bool? onPushResetApprovedVotes,

            int? reviewerCount,

            ImmutableArray<Outputs.BranchPolicyMinReviewersSettingsScope> scopes,

            bool? submitterCanVote)
        {
            AllowCompletionWithRejectsOrWaits = allowCompletionWithRejectsOrWaits;
            LastPusherCannotApprove = lastPusherCannotApprove;
            OnLastIterationRequireVote = onLastIterationRequireVote;
            OnPushResetAllVotes = onPushResetAllVotes;
            OnPushResetApprovedVotes = onPushResetApprovedVotes;
            ReviewerCount = reviewerCount;
            Scopes = scopes;
            SubmitterCanVote = submitterCanVote;
        }
    }
}
