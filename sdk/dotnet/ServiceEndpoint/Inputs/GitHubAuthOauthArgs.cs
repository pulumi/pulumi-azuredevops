// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.ServiceEndpoint.Inputs
{

    public sealed class GitHubAuthOauthArgs : global::Pulumi.ResourceArgs
    {
        [Input("oauthConfigurationId", required: true)]
        public Input<string> OauthConfigurationId { get; set; } = null!;

        public GitHubAuthOauthArgs()
        {
        }
        public static new GitHubAuthOauthArgs Empty => new GitHubAuthOauthArgs();
    }
}
