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
    /// Manages a Kubernetes service endpoint within Azure DevOps.
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Kubernetes can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes azuredevops_serviceendpoint_kubernetes.serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    public partial class ServiceEndpointKubernetes : Pulumi.CustomResource
    {
        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        [Output("apiserverUrl")]
        public Output<string> ApiserverUrl { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        /// <summary>
        /// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        /// </summary>
        [Output("authorizationType")]
        public Output<string> AuthorizationType { get; private set; } = null!;

        /// <summary>
        /// The configuration for authorization_type="AzureSubscription".
        /// </summary>
        [Output("azureSubscriptions")]
        public Output<ImmutableArray<Outputs.ServiceEndpointKubernetesAzureSubscription>> AzureSubscriptions { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The configuration for authorization_type="Kubeconfig".
        /// </summary>
        [Output("kubeconfigs")]
        public Output<ImmutableArray<Outputs.ServiceEndpointKubernetesKubeconfig>> Kubeconfigs { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        /// </summary>
        [Output("serviceAccounts")]
        public Output<ImmutableArray<Outputs.ServiceEndpointKubernetesServiceAccount>> ServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointKubernetes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointKubernetes(string name, ServiceEndpointKubernetesArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes", name, args ?? new ServiceEndpointKubernetesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointKubernetes(string name, Input<string> id, ServiceEndpointKubernetesState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointKubernetes:ServiceEndpointKubernetes", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azuredevops:ServiceEndpoint/kubernetes:Kubernetes"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointKubernetes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointKubernetes Get(string name, Input<string> id, ServiceEndpointKubernetesState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointKubernetes(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointKubernetesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        [Input("apiserverUrl", required: true)]
        public Input<string> ApiserverUrl { get; set; } = null!;

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        /// </summary>
        [Input("authorizationType", required: true)]
        public Input<string> AuthorizationType { get; set; } = null!;

        [Input("azureSubscriptions")]
        private InputList<Inputs.ServiceEndpointKubernetesAzureSubscriptionArgs>? _azureSubscriptions;

        /// <summary>
        /// The configuration for authorization_type="AzureSubscription".
        /// </summary>
        public InputList<Inputs.ServiceEndpointKubernetesAzureSubscriptionArgs> AzureSubscriptions
        {
            get => _azureSubscriptions ?? (_azureSubscriptions = new InputList<Inputs.ServiceEndpointKubernetesAzureSubscriptionArgs>());
            set => _azureSubscriptions = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("kubeconfigs")]
        private InputList<Inputs.ServiceEndpointKubernetesKubeconfigArgs>? _kubeconfigs;

        /// <summary>
        /// The configuration for authorization_type="Kubeconfig".
        /// </summary>
        public InputList<Inputs.ServiceEndpointKubernetesKubeconfigArgs> Kubeconfigs
        {
            get => _kubeconfigs ?? (_kubeconfigs = new InputList<Inputs.ServiceEndpointKubernetesKubeconfigArgs>());
            set => _kubeconfigs = value;
        }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("serviceAccounts")]
        private InputList<Inputs.ServiceEndpointKubernetesServiceAccountArgs>? _serviceAccounts;

        /// <summary>
        /// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        /// </summary>
        public InputList<Inputs.ServiceEndpointKubernetesServiceAccountArgs> ServiceAccounts
        {
            get => _serviceAccounts ?? (_serviceAccounts = new InputList<Inputs.ServiceEndpointKubernetesServiceAccountArgs>());
            set => _serviceAccounts = value;
        }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public ServiceEndpointKubernetesArgs()
        {
        }
    }

    public sealed class ServiceEndpointKubernetesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Service Endpoint description.
        /// </summary>
        [Input("apiserverUrl")]
        public Input<string>? ApiserverUrl { get; set; }

        [Input("authorization")]
        private InputMap<string>? _authorization;
        public InputMap<string> Authorization
        {
            get => _authorization ?? (_authorization = new InputMap<string>());
            set => _authorization = value;
        }

        /// <summary>
        /// The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.
        /// </summary>
        [Input("authorizationType")]
        public Input<string>? AuthorizationType { get; set; }

        [Input("azureSubscriptions")]
        private InputList<Inputs.ServiceEndpointKubernetesAzureSubscriptionGetArgs>? _azureSubscriptions;

        /// <summary>
        /// The configuration for authorization_type="AzureSubscription".
        /// </summary>
        public InputList<Inputs.ServiceEndpointKubernetesAzureSubscriptionGetArgs> AzureSubscriptions
        {
            get => _azureSubscriptions ?? (_azureSubscriptions = new InputList<Inputs.ServiceEndpointKubernetesAzureSubscriptionGetArgs>());
            set => _azureSubscriptions = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("kubeconfigs")]
        private InputList<Inputs.ServiceEndpointKubernetesKubeconfigGetArgs>? _kubeconfigs;

        /// <summary>
        /// The configuration for authorization_type="Kubeconfig".
        /// </summary>
        public InputList<Inputs.ServiceEndpointKubernetesKubeconfigGetArgs> Kubeconfigs
        {
            get => _kubeconfigs ?? (_kubeconfigs = new InputList<Inputs.ServiceEndpointKubernetesKubeconfigGetArgs>());
            set => _kubeconfigs = value;
        }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("serviceAccounts")]
        private InputList<Inputs.ServiceEndpointKubernetesServiceAccountGetArgs>? _serviceAccounts;

        /// <summary>
        /// The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.
        /// </summary>
        public InputList<Inputs.ServiceEndpointKubernetesServiceAccountGetArgs> ServiceAccounts
        {
            get => _serviceAccounts ?? (_serviceAccounts = new InputList<Inputs.ServiceEndpointKubernetesServiceAccountGetArgs>());
            set => _serviceAccounts = value;
        }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public ServiceEndpointKubernetesState()
        {
        }
    }
}
