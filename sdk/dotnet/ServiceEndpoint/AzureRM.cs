// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.ServiceEndpoint
{
    /// <summary>
    /// Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.
    /// 
    /// ## Requirements (Manual AzureRM Service Endpoint)
    /// 
    /// Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.
    /// 
    /// For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)
    /// 
    /// ## Example Usage
    /// ### Manual AzureRM Service Endpoint (Subscription Scoped)
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
    ///         var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new AzureDevOps.ServiceEndpointAzureRMArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             ServiceEndpointName = "Example AzureRM",
    ///             Description = "Managed by Terraform",
    ///             Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///             {
    ///                 Serviceprincipalid = "00000000-0000-0000-0000-000000000000",
    ///                 Serviceprincipalkey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///             },
    ///             AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///             AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///             AzurermSubscriptionName = "Example Subscription Name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Manual AzureRM Service Endpoint (ManagementGroup Scoped)
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
    ///         var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new AzureDevOps.ServiceEndpointAzureRMArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             ServiceEndpointName = "Example AzureRM",
    ///             Description = "Managed by Terraform",
    ///             Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///             {
    ///                 Serviceprincipalid = "00000000-0000-0000-0000-000000000000",
    ///                 Serviceprincipalkey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///             },
    ///             AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///             AzurermManagementGroupId = "managementGroup",
    ///             AzurermManagementGroupName = "managementGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Automatic AzureRM Service Endpoint
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
    ///         });
    ///         var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new AzureDevOps.ServiceEndpointAzureRMArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             ServiceEndpointName = "Example AzureRM",
    ///             AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///             AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///             AzurermSubscriptionName = "Example Subscription Name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [Obsolete(@"azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM")]
    [AzureDevOpsResourceType("azuredevops:ServiceEndpoint/azureRM:AzureRM")]
    public partial class AzureRM : Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        /// <summary>
        /// The management group Id of the Azure targets.
        /// </summary>
        [Output("azurermManagementGroupId")]
        public Output<string?> AzurermManagementGroupId { get; private set; } = null!;

        /// <summary>
        /// The management group Name of the targets.
        /// </summary>
        [Output("azurermManagementGroupName")]
        public Output<string?> AzurermManagementGroupName { get; private set; } = null!;

        /// <summary>
        /// The tenant id if the service principal.
        /// </summary>
        [Output("azurermSpnTenantid")]
        public Output<string> AzurermSpnTenantid { get; private set; } = null!;

        /// <summary>
        /// The subscription Id of the Azure targets.
        /// </summary>
        [Output("azurermSubscriptionId")]
        public Output<string?> AzurermSubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The subscription Name of the targets.
        /// </summary>
        [Output("azurermSubscriptionName")]
        public Output<string?> AzurermSubscriptionName { get; private set; } = null!;

        /// <summary>
        /// A `credentials` block.
        /// </summary>
        [Output("credentials")]
        public Output<Outputs.AzureRMCredentials?> Credentials { get; private set; } = null!;

        /// <summary>
        /// Service connection description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The resource group used for scope of automatic service endpoint.
        /// </summary>
        [Output("resourceGroup")]
        public Output<string?> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a AzureRM resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AzureRM(string name, AzureRMArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:ServiceEndpoint/azureRM:AzureRM", name, args ?? new AzureRMArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AzureRM(string name, Input<string> id, AzureRMState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:ServiceEndpoint/azureRM:AzureRM", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AzureRM resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AzureRM Get(string name, Input<string> id, AzureRMState? state = null, CustomResourceOptions? options = null)
        {
            return new AzureRM(name, id, state, options);
        }
    }

    public sealed class AzureRMArgs : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The management group Id of the Azure targets.
        /// </summary>
        [Input("azurermManagementGroupId")]
        public Input<string>? AzurermManagementGroupId { get; set; }

        /// <summary>
        /// The management group Name of the targets.
        /// </summary>
        [Input("azurermManagementGroupName")]
        public Input<string>? AzurermManagementGroupName { get; set; }

        /// <summary>
        /// The tenant id if the service principal.
        /// </summary>
        [Input("azurermSpnTenantid", required: true)]
        public Input<string> AzurermSpnTenantid { get; set; } = null!;

        /// <summary>
        /// The subscription Id of the Azure targets.
        /// </summary>
        [Input("azurermSubscriptionId")]
        public Input<string>? AzurermSubscriptionId { get; set; }

        /// <summary>
        /// The subscription Name of the targets.
        /// </summary>
        [Input("azurermSubscriptionName")]
        public Input<string>? AzurermSubscriptionName { get; set; }

        /// <summary>
        /// A `credentials` block.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.AzureRMCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// Service connection description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The resource group used for scope of automatic service endpoint.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public AzureRMArgs()
        {
        }
    }

    public sealed class AzureRMState : Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The management group Id of the Azure targets.
        /// </summary>
        [Input("azurermManagementGroupId")]
        public Input<string>? AzurermManagementGroupId { get; set; }

        /// <summary>
        /// The management group Name of the targets.
        /// </summary>
        [Input("azurermManagementGroupName")]
        public Input<string>? AzurermManagementGroupName { get; set; }

        /// <summary>
        /// The tenant id if the service principal.
        /// </summary>
        [Input("azurermSpnTenantid")]
        public Input<string>? AzurermSpnTenantid { get; set; }

        /// <summary>
        /// The subscription Id of the Azure targets.
        /// </summary>
        [Input("azurermSubscriptionId")]
        public Input<string>? AzurermSubscriptionId { get; set; }

        /// <summary>
        /// The subscription Name of the targets.
        /// </summary>
        [Input("azurermSubscriptionName")]
        public Input<string>? AzurermSubscriptionName { get; set; }

        /// <summary>
        /// A `credentials` block.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.AzureRMCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// Service connection description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The resource group used for scope of automatic service endpoint.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public AzureRMState()
        {
        }
    }
}
