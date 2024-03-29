// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BranchPolicyStatusCheckSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy applicability. If policy `applicability` is `default`, apply unless "Not Applicable" 
        /// status is posted to the pull request. If policy `applicability` is `conditional`, policy is applied only after a status
        /// is posted to the pull request.
        /// </summary>
        [Input("applicability")]
        public Input<string>? Applicability { get; set; }

        /// <summary>
        /// The authorized user can post the status.
        /// </summary>
        [Input("authorId")]
        public Input<string>? AuthorId { get; set; }

        /// <summary>
        /// The display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("filenamePatterns")]
        private InputList<string>? _filenamePatterns;

        /// <summary>
        /// If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
        /// </summary>
        public InputList<string> FilenamePatterns
        {
            get => _filenamePatterns ?? (_filenamePatterns = new InputList<string>());
            set => _filenamePatterns = value;
        }

        /// <summary>
        /// The genre of the status to check (see [Microsoft Documentation](https://docs.microsoft.com/en-us/azure/devops/repos/git/pull-request-status?view=azure-devops#status-policy))
        /// </summary>
        [Input("genre")]
        public Input<string>? Genre { get; set; }

        /// <summary>
        /// Reset status whenever there are new changes.
        /// </summary>
        [Input("invalidateOnUpdate")]
        public Input<bool>? InvalidateOnUpdate { get; set; }

        /// <summary>
        /// The status name to check.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("scopes", required: true)]
        private InputList<Inputs.BranchPolicyStatusCheckSettingsScopeGetArgs>? _scopes;

        /// <summary>
        /// Controls which repositories and branches the policy will be enabled for. This block must be defined
        /// at least once.
        /// </summary>
        public InputList<Inputs.BranchPolicyStatusCheckSettingsScopeGetArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.BranchPolicyStatusCheckSettingsScopeGetArgs>());
            set => _scopes = value;
        }

        public BranchPolicyStatusCheckSettingsGetArgs()
        {
        }
        public static new BranchPolicyStatusCheckSettingsGetArgs Empty => new BranchPolicyStatusCheckSettingsGetArgs();
    }
}
