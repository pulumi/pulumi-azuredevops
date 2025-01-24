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
    public sealed class BranchPolicyStatusCheckSettingsScope
    {
        /// <summary>
        /// The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
        /// </summary>
        public readonly string? MatchType;
        /// <summary>
        /// The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type=DefaultBranch`, this should not be defined.
        /// </summary>
        public readonly string? RepositoryId;
        /// <summary>
        /// The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type=Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type=Prefix`, this should be a ref path such as `refs/heads/releases`.
        /// </summary>
        public readonly string? RepositoryRef;

        [OutputConstructor]
        private BranchPolicyStatusCheckSettingsScope(
            string? matchType,

            string? repositoryId,

            string? repositoryRef)
        {
            MatchType = matchType;
            RepositoryId = repositoryId;
            RepositoryRef = repositoryRef;
        }
    }
}
