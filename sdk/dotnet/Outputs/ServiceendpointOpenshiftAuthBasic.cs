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
    public sealed class ServiceendpointOpenshiftAuthBasic
    {
        /// <summary>
        /// The password of the user.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ServiceendpointOpenshiftAuthBasic(
            string password,

            string username)
        {
            Password = password;
            Username = username;
        }
    }
}
