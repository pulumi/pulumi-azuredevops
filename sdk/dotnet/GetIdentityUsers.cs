// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetIdentityUsers
    {
        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).
        /// </summary>
        public static Task<GetIdentityUsersResult> InvokeAsync(GetIdentityUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityUsersResult>("azuredevops:index/getIdentityUsers:getIdentityUsers", args ?? new GetIdentityUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).
        /// </summary>
        public static Output<GetIdentityUsersResult> Invoke(GetIdentityUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityUsersResult>("azuredevops:index/getIdentityUsers:getIdentityUsers", args ?? new GetIdentityUsersInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).
        /// </summary>
        public static Output<GetIdentityUsersResult> Invoke(GetIdentityUsersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityUsersResult>("azuredevops:index/getIdentityUsers:getIdentityUsers", args ?? new GetIdentityUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PrincipalName of this identity member from the source provider.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The type of search to perform. Default is `General`. Possible values are `AccountName`, `DisplayName`, and `MailAddress`.
        /// </summary>
        [Input("searchFilter")]
        public string? SearchFilter { get; set; }

        public GetIdentityUsersArgs()
        {
        }
        public static new GetIdentityUsersArgs Empty => new GetIdentityUsersArgs();
    }

    public sealed class GetIdentityUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PrincipalName of this identity member from the source provider.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of search to perform. Default is `General`. Possible values are `AccountName`, `DisplayName`, and `MailAddress`.
        /// </summary>
        [Input("searchFilter")]
        public Input<string>? SearchFilter { get; set; }

        public GetIdentityUsersInvokeArgs()
        {
        }
        public static new GetIdentityUsersInvokeArgs Empty => new GetIdentityUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentityUsersResult
    {
        public readonly string Descriptor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This is the PrincipalName of this identity member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the identity member.
        /// </summary>
        public readonly string Name;
        public readonly string? SearchFilter;

        [OutputConstructor]
        private GetIdentityUsersResult(
            string descriptor,

            string id,

            string name,

            string? searchFilter)
        {
            Descriptor = descriptor;
            Id = id;
            Name = name;
            SearchFilter = searchFilter;
        }
    }
}
