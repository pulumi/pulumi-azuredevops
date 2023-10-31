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
    /// Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.
    /// 
    /// ## Requirements (Manual AzureRM Service Endpoint)
    /// 
    /// Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.
    /// 
    /// For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)
    /// 
    /// ## Example Usage
    /// ### Service Principal Manual AzureRM Service Endpoint (Subscription Scoped)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example AzureRM",
    ///         Description = "Managed by Terraform",
    ///         ServiceEndpointAuthenticationScheme = "ServicePrincipal",
    ///         Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///         {
    ///             Serviceprincipalid = "00000000-0000-0000-0000-000000000000",
    ///             Serviceprincipalkey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///         },
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Example Subscription Name",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Service Principal Manual AzureRM Service Endpoint (ManagementGroup Scoped)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example AzureRM",
    ///         Description = "Managed by Terraform",
    ///         ServiceEndpointAuthenticationScheme = "ServicePrincipal",
    ///         Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///         {
    ///             Serviceprincipalid = "00000000-0000-0000-0000-000000000000",
    ///             Serviceprincipalkey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///         },
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermManagementGroupId = "managementGroup",
    ///         AzurermManagementGroupName = "managementGroup",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Service Principal Automatic AzureRM Service Endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example AzureRM",
    ///         ServiceEndpointAuthenticationScheme = "ServicePrincipal",
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Example Subscription Name",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Workload Identity Federation Manual AzureRM Service Endpoint (Subscription Scoped)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// using Azurerm = Pulumi.Azurerm;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var serviceConnectionName = "example-federated-sc";
    /// 
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var identity = new Azurerm.Index.Azurerm_resource_group("identity", new()
    ///     {
    ///         Name = "identity",
    ///         Location = "UK South",
    ///     });
    /// 
    ///     var exampleazurerm_user_assigned_identity = new Azurerm.Index.Azurerm_user_assigned_identity("exampleazurerm_user_assigned_identity", new()
    ///     {
    ///         Location = identity.Location,
    ///         Name = "example-identity",
    ///         ResourceGroupName = "azurerm_resource_group.identity.name",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = serviceConnectionName,
    ///         Description = "Managed by Terraform",
    ///         ServiceEndpointAuthenticationScheme = "WorkloadIdentityFederation",
    ///         Credentials = new AzureDevOps.Inputs.ServiceEndpointAzureRMCredentialsArgs
    ///         {
    ///             Serviceprincipalid = exampleazurerm_user_assigned_identity.ClientId,
    ///         },
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Example Subscription Name",
    ///     });
    /// 
    ///     var exampleazurerm_federated_identity_credential = new Azurerm.Index.Azurerm_federated_identity_credential("exampleazurerm_federated_identity_credential", new()
    ///     {
    ///         Name = "example-federated-credential",
    ///         ResourceGroupName = identity.Name,
    ///         ParentId = exampleazurerm_user_assigned_identity.Id,
    ///         Audience = new[]
    ///         {
    ///             "api://AzureADTokenExchange",
    ///         },
    ///         Issuer = exampleServiceEndpointAzureRM.WorkloadIdentityFederationIssuer,
    ///         Subject = exampleServiceEndpointAzureRM.WorkloadIdentityFederationSubject,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Workload Identity Federation Automatic AzureRM Service Endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example AzureRM",
    ///         ServiceEndpointAuthenticationScheme = "WorkloadIdentityFederation",
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Example Subscription Name",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Managed Identity AzureRM Service Endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleProject = new AzureDevOps.Project("exampleProject", new()
    ///     {
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///     });
    /// 
    ///     var exampleServiceEndpointAzureRM = new AzureDevOps.ServiceEndpointAzureRM("exampleServiceEndpointAzureRM", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         ServiceEndpointName = "Example AzureRM",
    ///         ServiceEndpointAuthenticationScheme = "ManagedServiceIdentity",
    ///         AzurermSpnTenantid = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         AzurermSubscriptionName = "Example Subscription Name",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM")]
    public partial class ServiceEndpointAzureRM : global::Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        /// <summary>
        /// The Management group ID of the Azure targets.
        /// </summary>
        [Output("azurermManagementGroupId")]
        public Output<string?> AzurermManagementGroupId { get; private set; } = null!;

        /// <summary>
        /// The Management group Name of the targets.
        /// </summary>
        [Output("azurermManagementGroupName")]
        public Output<string?> AzurermManagementGroupName { get; private set; } = null!;

        /// <summary>
        /// The Tenant ID if the service principal.
        /// </summary>
        [Output("azurermSpnTenantid")]
        public Output<string> AzurermSpnTenantid { get; private set; } = null!;

        /// <summary>
        /// The Subscription ID of the Azure targets.
        /// </summary>
        [Output("azurermSubscriptionId")]
        public Output<string?> AzurermSubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The Subscription Name of the targets.
        /// </summary>
        [Output("azurermSubscriptionName")]
        public Output<string?> AzurermSubscriptionName { get; private set; } = null!;

        /// <summary>
        /// A `credentials` block.
        /// </summary>
        [Output("credentials")]
        public Output<Outputs.ServiceEndpointAzureRMCredentials?> Credentials { get; private set; } = null!;

        /// <summary>
        /// Service connection description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
        /// 
        /// &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
        /// </summary>
        [Output("environment")]
        public Output<string?> Environment { get; private set; } = null!;

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
        /// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
        /// 
        /// &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
        /// </summary>
        [Output("serviceEndpointAuthenticationScheme")]
        public Output<string?> ServiceEndpointAuthenticationScheme { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint Name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The Application(Client) ID of the Service Principal.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;

        /// <summary>
        /// The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
        /// </summary>
        [Output("workloadIdentityFederationIssuer")]
        public Output<string> WorkloadIdentityFederationIssuer { get; private set; } = null!;

        /// <summary>
        /// The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
        /// </summary>
        [Output("workloadIdentityFederationSubject")]
        public Output<string> WorkloadIdentityFederationSubject { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointAzureRM resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointAzureRM(string name, ServiceEndpointAzureRMArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM", name, args ?? new ServiceEndpointAzureRMArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointAzureRM(string name, Input<string> id, ServiceEndpointAzureRMState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAzureRM:ServiceEndpointAzureRM", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azuredevops:ServiceEndpoint/azureRM:AzureRM"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointAzureRM resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointAzureRM Get(string name, Input<string> id, ServiceEndpointAzureRMState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointAzureRM(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointAzureRMArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The Management group ID of the Azure targets.
        /// </summary>
        [Input("azurermManagementGroupId")]
        public Input<string>? AzurermManagementGroupId { get; set; }

        /// <summary>
        /// The Management group Name of the targets.
        /// </summary>
        [Input("azurermManagementGroupName")]
        public Input<string>? AzurermManagementGroupName { get; set; }

        /// <summary>
        /// The Tenant ID if the service principal.
        /// </summary>
        [Input("azurermSpnTenantid", required: true)]
        public Input<string> AzurermSpnTenantid { get; set; } = null!;

        /// <summary>
        /// The Subscription ID of the Azure targets.
        /// </summary>
        [Input("azurermSubscriptionId")]
        public Input<string>? AzurermSubscriptionId { get; set; }

        /// <summary>
        /// The Subscription Name of the targets.
        /// </summary>
        [Input("azurermSubscriptionName")]
        public Input<string>? AzurermSubscriptionName { get; set; }

        /// <summary>
        /// A `credentials` block.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.ServiceEndpointAzureRMCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// Service connection description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
        /// 
        /// &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

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
        /// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
        /// 
        /// &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
        /// </summary>
        [Input("serviceEndpointAuthenticationScheme")]
        public Input<string>? ServiceEndpointAuthenticationScheme { get; set; }

        /// <summary>
        /// The Service Endpoint Name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointAzureRMArgs()
        {
        }
        public static new ServiceEndpointAzureRMArgs Empty => new ServiceEndpointAzureRMArgs();
    }

    public sealed class ServiceEndpointAzureRMState : global::Pulumi.ResourceArgs
    {
        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The Management group ID of the Azure targets.
        /// </summary>
        [Input("azurermManagementGroupId")]
        public Input<string>? AzurermManagementGroupId { get; set; }

        /// <summary>
        /// The Management group Name of the targets.
        /// </summary>
        [Input("azurermManagementGroupName")]
        public Input<string>? AzurermManagementGroupName { get; set; }

        /// <summary>
        /// The Tenant ID if the service principal.
        /// </summary>
        [Input("azurermSpnTenantid")]
        public Input<string>? AzurermSpnTenantid { get; set; }

        /// <summary>
        /// The Subscription ID of the Azure targets.
        /// </summary>
        [Input("azurermSubscriptionId")]
        public Input<string>? AzurermSubscriptionId { get; set; }

        /// <summary>
        /// The Subscription Name of the targets.
        /// </summary>
        [Input("azurermSubscriptionName")]
        public Input<string>? AzurermSubscriptionName { get; set; }

        /// <summary>
        /// A `credentials` block.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.ServiceEndpointAzureRMCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// Service connection description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Cloud Environment to use. Defaults to `AzureCloud`. Possible values are `AzureCloud`, `AzureChinaCloud`. Changing this forces a new resource to be created.
        /// 
        /// &gt; **NOTE:** One of either `Subscription` scoped i.e. `azurerm_subscription_id`, `azurerm_subscription_name` or `ManagementGroup` scoped i.e. `azurerm_management_group_id`, `azurerm_management_group_name` values must be specified.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

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
        /// Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility.
        /// 
        /// &gt; **NOTE:** The `WorkloadIdentityFederation` authentication scheme is currently in private preview. Your organisation must be part of the preview and the feature toggle must be turned on to use it. More details can be found [here](https://aka.ms/azdo-rm-workload-identity).
        /// </summary>
        [Input("serviceEndpointAuthenticationScheme")]
        public Input<string>? ServiceEndpointAuthenticationScheme { get; set; }

        /// <summary>
        /// The Service Endpoint Name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The Application(Client) ID of the Service Principal.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
        /// </summary>
        [Input("workloadIdentityFederationIssuer")]
        public Input<string>? WorkloadIdentityFederationIssuer { get; set; }

        /// <summary>
        /// The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://&lt;organisation&gt;/&lt;project&gt;/&lt;service-connection-name&gt;`.
        /// </summary>
        [Input("workloadIdentityFederationSubject")]
        public Input<string>? WorkloadIdentityFederationSubject { get; set; }

        public ServiceEndpointAzureRMState()
        {
        }
        public static new ServiceEndpointAzureRMState Empty => new ServiceEndpointAzureRMState();
    }
}
