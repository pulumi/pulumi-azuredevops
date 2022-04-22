// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manages a GitHub Enterprise Server service endpoint within Azure DevOps.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleProject = new AzureDevOps.Project("exampleProject", new AzureDevOps.ProjectArgs
    ///         {
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///             Description = "Managed by Terraform",
    ///         });
    ///         var exampleServiceEndpointGitHubEnterprise = new AzureDevOps.ServiceEndpointGitHubEnterprise("exampleServiceEndpointGitHubEnterprise", new AzureDevOps.ServiceEndpointGitHubEnterpriseArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             ServiceEndpointName = "Example GitHub Enterprise",
    ///             Url = "https://github.contoso.com",
    ///             Description = "Managed by Terraform",
    ///             AuthPersonal = new AzureDevOps.Inputs.ServiceEndpointGitHubEnterpriseAuthPersonalArgs
    ///             {
    ///                 PersonalAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint GitHub Enterprise Server can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise")]
    public partial class ServiceEndpointGitHubEnterprise : Pulumi.CustomResource
    {
        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Output("authPersonal")]
        public Output<Outputs.ServiceEndpointGitHubEnterpriseAuthPersonal> AuthPersonal { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// GitHub Enterprise Server Url.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointGitHubEnterprise resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointGitHubEnterprise(string name, ServiceEndpointGitHubEnterpriseArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, args ?? new ServiceEndpointGitHubEnterpriseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointGitHubEnterprise(string name, Input<string> id, ServiceEndpointGitHubEnterpriseState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointGitHubEnterprise resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointGitHubEnterprise Get(string name, Input<string> id, ServiceEndpointGitHubEnterpriseState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointGitHubEnterprise(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointGitHubEnterpriseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal", required: true)]
        public Input<Inputs.ServiceEndpointGitHubEnterpriseAuthPersonalArgs> AuthPersonal { get; set; } = null!;

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// GitHub Enterprise Server Url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ServiceEndpointGitHubEnterpriseArgs()
        {
        }
    }

    public sealed class ServiceEndpointGitHubEnterpriseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal")]
        public Input<Inputs.ServiceEndpointGitHubEnterpriseAuthPersonalGetArgs>? AuthPersonal { get; set; }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// GitHub Enterprise Server Url.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ServiceEndpointGitHubEnterpriseState()
        {
        }
    }
}
