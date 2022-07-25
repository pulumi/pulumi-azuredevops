// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BranchPolicyMinReviewersSettingsScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        /// </summary>
        [Input("repositoryRef")]
        public Input<string>? RepositoryRef { get; set; }

        public BranchPolicyMinReviewersSettingsScopeArgs()
        {
        }
    }
}
