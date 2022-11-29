// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.ServiceEndpoint.Inputs
{

    public sealed class KubernetesServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("caCert", required: true)]
        private Input<string>? _caCert;

        /// <summary>
        /// The certificate from a Kubernetes secret object.
        /// </summary>
        public Input<string>? CaCert
        {
            get => _caCert;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _caCert = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("caCertHash")]
        private Input<string>? _caCertHash;
        public Input<string>? CaCertHash
        {
            get => _caCertHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _caCertHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// The token from a Kubernetes secret object.
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

        public KubernetesServiceAccountArgs()
        {
        }
        public static new KubernetesServiceAccountArgs Empty => new KubernetesServiceAccountArgs();
    }
}
