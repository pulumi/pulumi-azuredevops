// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Identities
{
    [Obsolete(@"azuredevops.identities.getUsers has been deprecated in favor of azuredevops.getUsers")]
    public static class GetUsers
    {
        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var user = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             PrincipalName = "contoso-user@contoso.onmicrosoft.com",
        ///         }));
        ///         var all_users = Output.Create(AzureDevOps.GetUsers.InvokeAsync());
        ///         var all_from_origin = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             Origin = "aad",
        ///         }));
        ///         var all_from_subjectTypes = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             SubjectTypes = 
        ///             {
        ///                 "aad",
        ///                 "msa",
        ///             },
        ///         }));
        ///         var all_from_origin_id = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             Origin = "aad",
        ///             OriginId = "a7ead982-8438-4cd2-b9e3-c3aa51a7b675",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 5.1 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-5.1)
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("azuredevops:Identities/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var user = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             PrincipalName = "contoso-user@contoso.onmicrosoft.com",
        ///         }));
        ///         var all_users = Output.Create(AzureDevOps.GetUsers.InvokeAsync());
        ///         var all_from_origin = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             Origin = "aad",
        ///         }));
        ///         var all_from_subjectTypes = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             SubjectTypes = 
        ///             {
        ///                 "aad",
        ///                 "msa",
        ///             },
        ///         }));
        ///         var all_from_origin_id = Output.Create(AzureDevOps.GetUsers.InvokeAsync(new AzureDevOps.GetUsersArgs
        ///         {
        ///             Origin = "aad",
        ///             OriginId = "a7ead982-8438-4cd2-b9e3-c3aa51a7b675",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 5.1 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-5.1)
        /// </summary>
        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUsersResult>("azuredevops:Identities/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
        /// </summary>
        [Input("origin")]
        public string? Origin { get; set; }

        /// <summary>
        /// The unique identifier from the system of origin.
        /// </summary>
        [Input("originId")]
        public string? OriginId { get; set; }

        /// <summary>
        /// The PrincipalName of this graph member from the source provider.
        /// </summary>
        [Input("principalName")]
        public string? PrincipalName { get; set; }

        [Input("subjectTypes")]
        private List<string>? _subjectTypes;

        /// <summary>
        /// A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
        /// </summary>
        public List<string> SubjectTypes
        {
            get => _subjectTypes ?? (_subjectTypes = new List<string>());
            set => _subjectTypes = value;
        }

        public GetUsersArgs()
        {
        }
    }

    public sealed class GetUsersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The unique identifier from the system of origin.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// The PrincipalName of this graph member from the source provider.
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        [Input("subjectTypes")]
        private InputList<string>? _subjectTypes;

        /// <summary>
        /// A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
        /// </summary>
        public InputList<string> SubjectTypes
        {
            get => _subjectTypes ?? (_subjectTypes = new InputList<string>());
            set => _subjectTypes = value;
        }

        public GetUsersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        /// </summary>
        public readonly string? Origin;
        /// <summary>
        /// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
        /// </summary>
        public readonly string? OriginId;
        /// <summary>
        /// This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
        /// </summary>
        public readonly string? PrincipalName;
        public readonly ImmutableArray<string> SubjectTypes;
        /// <summary>
        /// A set of existing users in your Azure DevOps Organization with details about every single user which includes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersUserResult> Users;

        [OutputConstructor]
        private GetUsersResult(
            string id,

            string? origin,

            string? originId,

            string? principalName,

            ImmutableArray<string> subjectTypes,

            ImmutableArray<Outputs.GetUsersUserResult> users)
        {
            Id = id;
            Origin = origin;
            OriginId = originId;
            PrincipalName = principalName;
            SubjectTypes = subjectTypes;
            Users = users;
        }
    }
}
