// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Build.Outputs
{

    [OutputType]
    public sealed class BuildDefinitionVariable
    {
        /// <summary>
        /// True if the variable can be overridden. Defaults to `true`.
        /// </summary>
        public readonly bool? AllowOverride;
        /// <summary>
        /// True if the variable is a secret. Defaults to `false`.
        /// </summary>
        public readonly bool? IsSecret;
        /// <summary>
        /// The name of the variable.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The secret value of the variable. Used when `is_secret` set to `true`.
        /// </summary>
        public readonly string? SecretValue;
        /// <summary>
        /// The value of the variable.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private BuildDefinitionVariable(
            bool? allowOverride,

            bool? isSecret,

            string name,

            string? secretValue,

            string? value)
        {
            AllowOverride = allowOverride;
            IsSecret = isSecret;
            Name = name;
            SecretValue = secretValue;
            Value = value;
        }
    }
}
