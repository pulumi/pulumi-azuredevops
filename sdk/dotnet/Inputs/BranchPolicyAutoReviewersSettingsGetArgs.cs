// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BranchPolicyAutoReviewersSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoReviewerIds", required: true)]
        private InputList<string>? _autoReviewerIds;

        /// <summary>
        /// Required reviewers ids. Supports multiples user Ids.
        /// </summary>
        public InputList<string> AutoReviewerIds
        {
            get => _autoReviewerIds ?? (_autoReviewerIds = new InputList<string>());
            set => _autoReviewerIds = value;
        }

        /// <summary>
        /// Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Minimum number of required reviewers. Defaults to `1`.
        /// 
        /// &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
        /// </summary>
        [Input("minimumNumberOfReviewers")]
        public Input<int>? MinimumNumberOfReviewers { get; set; }

        [Input("pathFilters")]
        private InputList<string>? _pathFilters;

        /// <summary>
        /// Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
        /// </summary>
        public InputList<string> PathFilters
        {
            get => _pathFilters ?? (_pathFilters = new InputList<string>());
            set => _pathFilters = value;
        }

        [Input("scopes", required: true)]
        private InputList<Inputs.BranchPolicyAutoReviewersSettingsScopeGetArgs>? _scopes;

        /// <summary>
        /// A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        /// </summary>
        public InputList<Inputs.BranchPolicyAutoReviewersSettingsScopeGetArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.BranchPolicyAutoReviewersSettingsScopeGetArgs>());
            set => _scopes = value;
        }

        /// <summary>
        /// Controls whether or not the submitter's vote counts. Defaults to `false`.
        /// </summary>
        [Input("submitterCanVote")]
        public Input<bool>? SubmitterCanVote { get; set; }

        public BranchPolicyAutoReviewersSettingsGetArgs()
        {
        }
        public static new BranchPolicyAutoReviewersSettingsGetArgs Empty => new BranchPolicyAutoReviewersSettingsGetArgs();
    }
}
