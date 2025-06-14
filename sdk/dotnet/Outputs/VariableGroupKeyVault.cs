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
    public sealed class VariableGroupKeyVault
    {
        /// <summary>
        /// The name of the Azure key vault to link secrets from as variables.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set the Azure Key Vault Secret search depth. Defaults to `20`.
        /// </summary>
        public readonly int? SearchDepth;
        /// <summary>
        /// The id of the Azure subscription endpoint to access the key vault.
        /// </summary>
        public readonly string ServiceEndpointId;

        [OutputConstructor]
        private VariableGroupKeyVault(
            string name,

            int? searchDepth,

            string serviceEndpointId)
        {
            Name = name;
            SearchDepth = searchDepth;
            ServiceEndpointId = serviceEndpointId;
        }
    }
}
