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
    public sealed class BranchPolicyStatusCheckSettings
    {
        /// <summary>
        /// Policy applicability. If policy `applicability` is `default`, apply unless "Not Applicable" 
        /// status is posted to the pull request. If policy `applicability` is `conditional`, policy is applied only after a status
        /// is posted to the pull request.
        /// </summary>
        public readonly string? Applicability;
        /// <summary>
        /// The authorized user can post the status.
        /// </summary>
        public readonly string? AuthorId;
        /// <summary>
        /// The display name.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
        /// </summary>
        public readonly ImmutableArray<string> FilenamePatterns;
        /// <summary>
        /// The genre of the status to check (see [Microsoft Documentation](https://docs.microsoft.com/en-us/azure/devops/repos/git/pull-request-status?view=azure-devops#status-policy))
        /// </summary>
        public readonly string? Genre;
        /// <summary>
        /// Reset status whenever there are new changes.
        /// </summary>
        public readonly bool? InvalidateOnUpdate;
        /// <summary>
        /// The status name to check.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Controls which repositories and branches the policy will be enabled for. This block must be defined
        /// at least once.
        /// </summary>
        public readonly ImmutableArray<Outputs.BranchPolicyStatusCheckSettingsScope> Scopes;

        [OutputConstructor]
        private BranchPolicyStatusCheckSettings(
            string? applicability,

            string? authorId,

            string? displayName,

            ImmutableArray<string> filenamePatterns,

            string? genre,

            bool? invalidateOnUpdate,

            string name,

            ImmutableArray<Outputs.BranchPolicyStatusCheckSettingsScope> scopes)
        {
            Applicability = applicability;
            AuthorId = authorId;
            DisplayName = displayName;
            FilenamePatterns = filenamePatterns;
            Genre = genre;
            InvalidateOnUpdate = invalidateOnUpdate;
            Name = name;
            Scopes = scopes;
        }
    }
}
