// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointArtifactoryAuthenticationTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// Authentication Token generated through Artifactory.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tokenHash")]
        private Input<string>? _tokenHash;
        public Input<string>? TokenHash
        {
            get => _tokenHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tokenHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ServiceEndpointArtifactoryAuthenticationTokenArgs()
        {
        }
        public static new ServiceEndpointArtifactoryAuthenticationTokenArgs Empty => new ServiceEndpointArtifactoryAuthenticationTokenArgs();
    }
}
