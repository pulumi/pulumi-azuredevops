// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointGitHubEnterpriseAuthPersonalArgs : global::Pulumi.ResourceArgs
    {
        [Input("personalAccessToken", required: true)]
        private Input<string>? _personalAccessToken;

        /// <summary>
        /// The Personal Access Token for GitHub.
        /// </summary>
        public Input<string>? PersonalAccessToken
        {
            get => _personalAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _personalAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ServiceEndpointGitHubEnterpriseAuthPersonalArgs()
        {
        }
        public static new ServiceEndpointGitHubEnterpriseAuthPersonalArgs Empty => new ServiceEndpointGitHubEnterpriseAuthPersonalArgs();
    }
}
