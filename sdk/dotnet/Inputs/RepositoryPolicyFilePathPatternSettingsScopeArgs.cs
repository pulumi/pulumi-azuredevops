// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class RepositoryPolicyFilePathPatternSettingsScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository ID.
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        public RepositoryPolicyFilePathPatternSettingsScopeArgs()
        {
        }
    }
}
