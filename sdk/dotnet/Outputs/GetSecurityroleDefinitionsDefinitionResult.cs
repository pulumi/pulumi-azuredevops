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
    public sealed class GetSecurityroleDefinitionsDefinitionResult
    {
        /// <summary>
        /// The mask of allowed permissions of the Security Role Definition.
        /// </summary>
        public readonly int AllowPermissions;
        /// <summary>
        /// The mask of the denied permissions of the Security Role Definition.
        /// </summary>
        public readonly int? DenyPermissions;
        /// <summary>
        /// The description of the Security Role Definition.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the Security Role Definition.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The identifier of the Security Role Definition.
        /// </summary>
        public readonly string Identifier;
        /// <summary>
        /// The name of the Security Role Definition.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Name of the Scope for which Security Role Definitions will be returned.
        /// 
        /// DataSource without specifying any arguments will return all projects.
        /// </summary>
        public readonly string Scope;

        [OutputConstructor]
        private GetSecurityroleDefinitionsDefinitionResult(
            int allowPermissions,

            int? denyPermissions,

            string description,

            string displayName,

            string identifier,

            string name,

            string scope)
        {
            AllowPermissions = allowPermissions;
            DenyPermissions = denyPermissions;
            Description = description;
            DisplayName = displayName;
            Identifier = identifier;
            Name = name;
            Scope = scope;
        }
    }
}
