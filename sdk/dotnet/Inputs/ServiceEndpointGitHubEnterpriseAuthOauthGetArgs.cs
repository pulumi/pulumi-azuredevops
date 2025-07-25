// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointGitHubEnterpriseAuthOauthGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OAuth Configuration ID.
        /// </summary>
        [Input("oauthConfigurationId", required: true)]
        public Input<string> OauthConfigurationId { get; set; } = null!;

        public ServiceEndpointGitHubEnterpriseAuthOauthGetArgs()
        {
        }
        public static new ServiceEndpointGitHubEnterpriseAuthOauthGetArgs Empty => new ServiceEndpointGitHubEnterpriseAuthOauthGetArgs();
    }
}
