// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointGitHubAuthOauthArgs : Pulumi.ResourceArgs
    {
        [Input("oauthConfigurationId", required: true)]
        public Input<string> OauthConfigurationId { get; set; } = null!;

        public ServiceEndpointGitHubAuthOauthArgs()
        {
        }
    }
}