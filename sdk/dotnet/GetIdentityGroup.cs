// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    public static class GetIdentityGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Group within Azure DevOps On-Premise(Azure DevOps Server).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // load existing group with specific name
        ///     var example_project_group = AzureDevOps.GetIdentityGroup.Invoke(new()
        ///     {
        ///         ProjectId = example.Id,
        ///         Name = "[Project-Name]\\Group-Name",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.1 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
        /// </summary>
        public static Task<GetIdentityGroupResult> InvokeAsync(GetIdentityGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityGroupResult>("azuredevops:index/getIdentityGroup:getIdentityGroup", args ?? new GetIdentityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Group within Azure DevOps On-Premise(Azure DevOps Server).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // load existing group with specific name
        ///     var example_project_group = AzureDevOps.GetIdentityGroup.Invoke(new()
        ///     {
        ///         ProjectId = example.Id,
        ///         Name = "[Project-Name]\\Group-Name",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.1 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
        /// </summary>
        public static Output<GetIdentityGroupResult> Invoke(GetIdentityGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupResult>("azuredevops:index/getIdentityGroup:getIdentityGroup", args ?? new GetIdentityGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Group within Azure DevOps On-Premise(Azure DevOps Server).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureDevOps = Pulumi.AzureDevOps;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // load existing group with specific name
        ///     var example_project_group = AzureDevOps.GetIdentityGroup.Invoke(new()
        ///     {
        ///         ProjectId = example.Id,
        ///         Name = "[Project-Name]\\Group-Name",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Relevant Links
        /// 
        /// - [Azure DevOps Service REST API 7.1 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)
        /// </summary>
        public static Output<GetIdentityGroupResult> Invoke(GetIdentityGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupResult>("azuredevops:index/getIdentityGroup:getIdentityGroup", args ?? new GetIdentityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetIdentityGroupArgs()
        {
        }
        public static new GetIdentityGroupArgs Empty => new GetIdentityGroupArgs();
    }

    public sealed class GetIdentityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetIdentityGroupInvokeArgs()
        {
        }
        public static new GetIdentityGroupInvokeArgs Empty => new GetIdentityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentityGroupResult
    {
        /// <summary>
        /// The descriptor of the identity group.
        /// </summary>
        public readonly string Descriptor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
        /// </summary>
        public readonly string Name;
        public readonly string ProjectId;
        /// <summary>
        /// The subject descriptor of the identity group.
        /// </summary>
        public readonly string SubjectDescriptor;

        [OutputConstructor]
        private GetIdentityGroupResult(
            string descriptor,

            string id,

            string name,

            string projectId,

            string subjectDescriptor)
        {
            Descriptor = descriptor;
            Id = id;
            Name = name;
            ProjectId = projectId;
            SubjectDescriptor = subjectDescriptor;
        }
    }
}
