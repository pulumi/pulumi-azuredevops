// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Identities.Outputs
{

    [OutputType]
    public sealed class GetUsersUserResult
    {
        /// <summary>
        /// The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
        /// </summary>
        public readonly string Descriptor;
        /// <summary>
        /// This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The user ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The email address of record for a given graph member. This may be different than the principal name.
        /// </summary>
        public readonly string MailAddress;
        /// <summary>
        /// The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
        /// </summary>
        public readonly string Origin;
        /// <summary>
        /// The unique identifier from the system of origin.
        /// </summary>
        public readonly string? OriginId;
        /// <summary>
        /// The PrincipalName of this graph member from the source provider.
        /// </summary>
        public readonly string PrincipalName;

        [OutputConstructor]
        private GetUsersUserResult(
            string descriptor,

            string displayName,

            string id,

            string mailAddress,

            string origin,

            string? originId,

            string principalName)
        {
            Descriptor = descriptor;
            DisplayName = displayName;
            Id = id;
            MailAddress = mailAddress;
            Origin = origin;
            OriginId = originId;
            PrincipalName = principalName;
        }
    }
}
