// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.ServiceEndpoint.Inputs
{

    public sealed class AzureRMCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service principal application Id
        /// </summary>
        [Input("serviceprincipalid", required: true)]
        public Input<string> Serviceprincipalid { get; set; } = null!;

        [Input("serviceprincipalkey")]
        private Input<string>? _serviceprincipalkey;

        /// <summary>
        /// The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
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

        public AzureRMCredentialsGetArgs()
        {
        }
        public static new AzureRMCredentialsGetArgs Empty => new AzureRMCredentialsGetArgs();
    }
}
