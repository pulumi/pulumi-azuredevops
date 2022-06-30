// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionPullRequestTriggerGetArgs : Pulumi.ResourceArgs
    {
        [Input("commentRequired")]
        public Input<string>? CommentRequired { get; set; }

        /// <summary>
        /// Set permissions for Forked repositories.
        /// </summary>
        [Input("forks", required: true)]
        public Input<Inputs.BuildDefinitionPullRequestTriggerForksGetArgs> Forks { get; set; } = null!;

        [Input("initialBranch")]
        public Input<string>? InitialBranch { get; set; }

        /// <summary>
        /// Override the azure-pipeline file and use this configuration for all builds.
        /// </summary>
        [Input("override")]
        public Input<Inputs.BuildDefinitionPullRequestTriggerOverrideGetArgs>? Override { get; set; }

        /// <summary>
        /// Use the azure-pipeline file for the build configuration. Defaults to `false`.
        /// </summary>
        [Input("useYaml")]
        public Input<bool>? UseYaml { get; set; }

        public BuildDefinitionPullRequestTriggerGetArgs()
        {
        }
    }
}
