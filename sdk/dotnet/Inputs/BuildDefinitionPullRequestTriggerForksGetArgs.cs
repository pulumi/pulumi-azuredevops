// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionPullRequestTriggerForksGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Build pull requests from forks of this repository.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Make secrets available to builds of forks.
        /// </summary>
        [Input("shareSecrets", required: true)]
        public Input<bool> ShareSecrets { get; set; } = null!;

        public BuildDefinitionPullRequestTriggerForksGetArgs()
        {
        }
        public static new BuildDefinitionPullRequestTriggerForksGetArgs Empty => new BuildDefinitionPullRequestTriggerForksGetArgs();
    }
}
