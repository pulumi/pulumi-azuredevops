// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceEndpointAzureRMCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service principal application Id
        /// </summary>
        [Input("serviceprincipalid", required: true)]
        public Input<string> Serviceprincipalid { get; set; } = null!;

        [Input("serviceprincipalkey", required: true)]
        private Input<string>? _serviceprincipalkey;

        /// <summary>
        /// The service principal secret.
        /// </summary>
        public Input<string>? Serviceprincipalkey
        {
            get => _serviceprincipalkey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _serviceprincipalkey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("serviceprincipalkeyHash")]
        private Input<string>? _serviceprincipalkeyHash;
        public Input<string>? ServiceprincipalkeyHash
        {
            get => _serviceprincipalkeyHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _serviceprincipalkeyHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ServiceEndpointAzureRMCredentialsArgs()
        {
        }
        public static new ServiceEndpointAzureRMCredentialsArgs Empty => new ServiceEndpointAzureRMCredentialsArgs();
    }
}
