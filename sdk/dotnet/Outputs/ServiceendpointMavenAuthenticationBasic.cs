// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class ServiceendpointMavenAuthenticationBasic
    {
        /// <summary>
        /// The password Maven Repository.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The Username of the Maven Repository.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ServiceendpointMavenAuthenticationBasic(
            string password,

            string username)
        {
            Password = password;
            Username = username;
        }
    }
}