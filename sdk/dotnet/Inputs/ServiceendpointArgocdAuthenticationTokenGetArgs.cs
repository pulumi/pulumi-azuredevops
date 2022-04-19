// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class ServiceendpointArgocdAuthenticationTokenGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication Token generated through ArgoCD.
        /// </summary>
        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        [Input("tokenHash")]
        public Input<string>? TokenHash { get; set; }

        public ServiceendpointArgocdAuthenticationTokenGetArgs()
        {
        }
    }
}
