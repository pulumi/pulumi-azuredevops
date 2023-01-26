// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Identities
{
    [Obsolete(@"azuredevops.identities.getGroup has been deprecated in favor of azuredevops.getGroup")]
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Group within Azure DevOps
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGroup = AzureDevOps.GetGroup.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Name = "Example Group",
        ///     });
        /// 
        ///     var example_collection_group = AzureDevOps.GetGroup.Invoke(new()
        ///     {
        ///         Name = "Project Collection Administrators",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["groupId"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Id),
        ///         ["groupDescriptor"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Descriptor),
        ///         ["collectionGroupId"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Id),
        ///         ["collectionGroupDescriptor"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Descriptor),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-6.0)
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("azuredevops:Identities/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Group within Azure DevOps
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = AzureDevOps.GetProject.Invoke(new()
        ///     {
        ///         Name = "Example Project",
        ///     });
        /// 
        ///     var exampleGroup = AzureDevOps.GetGroup.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Name = "Example Group",
        ///     });
        /// 
        ///     var example_collection_group = AzureDevOps.GetGroup.Invoke(new()
        ///     {
        ///         Name = "Project Collection Administrators",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["groupId"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Id),
        ///         ["groupDescriptor"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Descriptor),
        ///         ["collectionGroupId"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Id),
        ///         ["collectionGroupDescriptor"] = exampleGroup.Apply(getGroupResult =&gt; getGroupResult.Descriptor),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 6.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-6.0)
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("azuredevops:Identities/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Group Name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Project ID. If no project ID is specified the project collection groups will be searched.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Group Name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Project ID. If no project ID is specified the project collection groups will be searched.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations.
        /// </summary>
        public readonly string Descriptor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        /// </summary>
        public readonly string Origin;
        /// <summary>
        /// The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
        /// </summary>
        public readonly string OriginId;
        public readonly string? ProjectId;

        [OutputConstructor]
        private GetGroupResult(
            string descriptor,

            string id,

            string name,

            string origin,

            string originId,

            string? projectId)
        {
            Descriptor = descriptor;
            Id = id;
            Name = name;
            Origin = origin;
            OriginId = originId;
            ProjectId = projectId;
        }
    }
}
