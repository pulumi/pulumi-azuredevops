// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetUsers
    {
        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         PrincipalName = "contoso-user@contoso.onmicrosoft.com",
        ///     });
        /// 
        ///     var example_all_users = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         Features = new AzureDevOps.Inputs.GetUsersFeaturesInputArgs
        ///         {
        ///             ConcurrentWorkers = 10,
        ///         },
        ///     });
        /// 
        ///     var example_all_from_origin = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         Origin = "aad",
        ///     });
        /// 
        ///     var example_all_from_subjectTypes = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         SubjectTypes = new[]
        ///         {
        ///             "aad",
        ///             "msa",
        ///         },
        ///     });
        /// 
        ///     var example_all_from_origin_id = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         Origin = "aad",
        ///         OriginId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-7.0)
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("azuredevops:index/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing users within Azure DevOps.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         PrincipalName = "contoso-user@contoso.onmicrosoft.com",
        ///     });
        /// 
        ///     var example_all_users = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         Features = new AzureDevOps.Inputs.GetUsersFeaturesInputArgs
        ///         {
        ///             ConcurrentWorkers = 10,
        ///         },
        ///     });
        /// 
        ///     var example_all_from_origin = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         Origin = "aad",
        ///     });
        /// 
        ///     var example_all_from_subjectTypes = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         SubjectTypes = new[]
        ///         {
        ///             "aad",
        ///             "msa",
        ///         },
        ///     });
        /// 
        ///     var example_all_from_origin_id = AzureDevOps.GetUsers.Invoke(new()
        ///     {
        ///         Origin = "aad",
        ///         OriginId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.0 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-7.0)
        /// </summary>
        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsersResult>("azuredevops:index/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A `features` block as defined below.
        /// 
        /// DataSource without specifying any arguments will return all users inside an organization.
        /// 
        /// List of possible subject types
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// List of possible origins
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        [Input("features")]
        public Inputs.GetUsersFeaturesArgs? Features { get; set; }

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
        public static new GetUsersArgs Empty => new GetUsersArgs();
    }

    public sealed class GetUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A `features` block as defined below.
        /// 
        /// DataSource without specifying any arguments will return all users inside an organization.
        /// 
        /// List of possible subject types
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// List of possible origins
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        [Input("features")]
        public Input<Inputs.GetUsersFeaturesInputArgs>? Features { get; set; }

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
        public static new GetUsersInvokeArgs Empty => new GetUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        public readonly Outputs.GetUsersFeaturesResult? Features;
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
            Outputs.GetUsersFeaturesResult? features,

            string id,

            string? origin,

            string? originId,

            string? principalName,

            ImmutableArray<string> subjectTypes,

            ImmutableArray<Outputs.GetUsersUserResult> users)
        {
            Features = features;
            Id = id;
            Origin = origin;
            OriginId = originId;
            PrincipalName = principalName;
            SubjectTypes = subjectTypes;
            Users = users;
        }
    }
}
