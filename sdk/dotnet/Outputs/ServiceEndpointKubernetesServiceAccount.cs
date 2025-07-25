// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class ServiceEndpointKubernetesServiceAccount
    {
        /// <summary>
        /// Set this option to allow clients to accept a self-signed certificate. Defaults to `false`.
        /// </summary>
        public readonly bool? AcceptUntrustedCerts;
        /// <summary>
        /// The certificate from a Kubernetes secret object.
        /// </summary>
        public readonly string CaCert;
        /// <summary>
        /// The token from a Kubernetes secret object.
        /// </summary>
        public readonly string Token;

        [OutputConstructor]
        private ServiceEndpointKubernetesServiceAccount(
            bool? acceptUntrustedCerts,

            string caCert,

            string token)
        {
            AcceptUntrustedCerts = acceptUntrustedCerts;
            CaCert = caCert;
            Token = token;
        }
    }
}
