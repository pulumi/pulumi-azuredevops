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
    /// Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external git service
    /// using basic authentication via a username and password. This is mostly useful for importing private git repositories.
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
    ///         var project = new AzureDevOps.Project("project", new AzureDevOps.ProjectArgs
    ///         {
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var serviceendpoint = new AzureDevOps.ServiceEndpointGenericGit("serviceendpoint", new AzureDevOps.ServiceEndpointGenericGitArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             RepositoryUrl = "https://dev.azure.com/org/project/_git/repository",
    ///             Username = "username",
    ///             Password = "password",
    ///             ServiceEndpointName = "Sample Generic Git",
    ///             Description = "Managed by Terraform",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Generic Git can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit")]
    public partial class ServiceEndpointGenericGit : Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        /// </summary>
        [Output("enablePipelinesAccess")]
        public Output<bool?> EnablePipelinesAccess { get; private set; } = null!;

        /// <summary>
        /// The PAT or password used to authenticate to the git repository.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// A bcrypted hash of the attribute 'password'
        /// </summary>
        [Output("passwordHash")]
        public Output<string> PasswordHash { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name to associate with the service endpoint.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The URL of the repository associated with the service endpoint.
        /// </summary>
        [Output("repositoryUrl")]
        public Output<string> RepositoryUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the service endpoint.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The username used to authenticate to the git repository.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointGenericGit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointGenericGit(string name, ServiceEndpointGenericGitArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit", name, args ?? new ServiceEndpointGenericGitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointGenericGit(string name, Input<string> id, ServiceEndpointGenericGitState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointGenericGit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointGenericGit Get(string name, Input<string> id, ServiceEndpointGenericGitState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointGenericGit(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointGenericGitArgs : Pulumi.ResourceArgs
    {
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
        /// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        /// </summary>
        [Input("enablePipelinesAccess")]
        public Input<bool>? EnablePipelinesAccess { get; set; }

        /// <summary>
        /// The PAT or password used to authenticate to the git repository.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The project ID or project name to associate with the service endpoint.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The URL of the repository associated with the service endpoint.
        /// </summary>
        [Input("repositoryUrl", required: true)]
        public Input<string> RepositoryUrl { get; set; } = null!;

        /// <summary>
        /// The name of the service endpoint.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// The username used to authenticate to the git repository.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceEndpointGenericGitArgs()
        {
        }
    }

    public sealed class ServiceEndpointGenericGitState : Pulumi.ResourceArgs
    {
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
        /// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        /// </summary>
        [Input("enablePipelinesAccess")]
        public Input<bool>? EnablePipelinesAccess { get; set; }

        /// <summary>
        /// The PAT or password used to authenticate to the git repository.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// A bcrypted hash of the attribute 'password'
        /// </summary>
        [Input("passwordHash")]
        public Input<string>? PasswordHash { get; set; }

        /// <summary>
        /// The project ID or project name to associate with the service endpoint.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The URL of the repository associated with the service endpoint.
        /// </summary>
        [Input("repositoryUrl")]
        public Input<string>? RepositoryUrl { get; set; }

        /// <summary>
        /// The name of the service endpoint.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The username used to authenticate to the git repository.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceEndpointGenericGitState()
        {
        }
    }
}