// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointGitHubEnterpriseAuthPersonalGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Personal Access Token for GitHub.
        /// </summary>
        [Input("personalAccessToken", required: true)]
        public Input<string> PersonalAccessToken { get; set; } = null!;

        [Input("personalAccessTokenHash")]
        public Input<string>? PersonalAccessTokenHash { get; set; }

        public ServiceEndpointGitHubEnterpriseAuthPersonalGetArgs()
        {
        }
    }
}
