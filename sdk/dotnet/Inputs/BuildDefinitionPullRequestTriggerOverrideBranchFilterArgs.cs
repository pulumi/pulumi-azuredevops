// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludes")]
        private InputList<string>? _excludes;

        /// <summary>
        /// List of branch patterns to exclude.
        /// </summary>
        public InputList<string> Excludes
        {
            get => _excludes ?? (_excludes = new InputList<string>());
            set => _excludes = value;
        }

        [Input("includes")]
        private InputList<string>? _includes;

        /// <summary>
        /// List of branch patterns to include.
        /// </summary>
        public InputList<string> Includes
        {
            get => _includes ?? (_includes = new InputList<string>());
            set => _includes = value;
        }

        public BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs()
        {
        }
        public static new BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs Empty => new BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs();
    }
}
