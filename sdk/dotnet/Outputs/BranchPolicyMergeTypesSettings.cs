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
    public sealed class BranchPolicyMergeTypesSettings
    {
        /// <summary>
        /// Allow basic merge with no fast forward. Defaults to `false`.
        /// </summary>
        public readonly bool? AllowBasicNoFastForward;
        /// <summary>
        /// Allow rebase with fast forward. Defaults to `false`.
        /// </summary>
        public readonly bool? AllowRebaseAndFastForward;
        /// <summary>
        /// Allow rebase with merge commit. Defaults to `false`.
        /// </summary>
        public readonly bool? AllowRebaseWithMerge;
        /// <summary>
        /// Allow squash merge. Defaults to `false`
        /// </summary>
        public readonly bool? AllowSquash;
        /// <summary>
        /// Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        /// </summary>
        public readonly ImmutableArray<Outputs.BranchPolicyMergeTypesSettingsScope> Scopes;

        [OutputConstructor]
        private BranchPolicyMergeTypesSettings(
            bool? allowBasicNoFastForward,

            bool? allowRebaseAndFastForward,

            bool? allowRebaseWithMerge,

            bool? allowSquash,

            ImmutableArray<Outputs.BranchPolicyMergeTypesSettingsScope> scopes)
        {
            AllowBasicNoFastForward = allowBasicNoFastForward;
            AllowRebaseAndFastForward = allowRebaseAndFastForward;
            AllowRebaseWithMerge = allowRebaseWithMerge;
            AllowSquash = allowSquash;
            Scopes = scopes;
        }
    }
}
