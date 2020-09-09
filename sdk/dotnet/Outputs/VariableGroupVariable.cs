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
    public sealed class VariableGroupVariable
    {
        public readonly string? ContentType;
        public readonly bool? Enabled;
        public readonly string? Expires;
        /// <summary>
        /// A boolean flag describing if the variable value is sensitive. Defaults to `false`.
        /// </summary>
        public readonly bool? IsSecret;
        /// <summary>
        /// The key value used for the variable. Must be unique within the Variable Group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
        /// </summary>
        public readonly string? SecretValue;
        /// <summary>
        /// The value of the variable. If omitted, it will default to empty string.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private VariableGroupVariable(
            string? contentType,

            bool? enabled,

            string? expires,

            bool? isSecret,

            string name,

            string? secretValue,

            string? value)
        {
            ContentType = contentType;
            Enabled = enabled;
            Expires = expires;
            IsSecret = isSecret;
            Name = name;
            SecretValue = secretValue;
            Value = value;
        }
    }
}