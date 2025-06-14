// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps
{
    /// <summary>
    /// Manages an Other Git service endpoint within Azure DevOps, which can be used to authenticate to any external git service
    /// using basic authentication via a username and password. This is mostly useful for importing private git repositories.
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
    ///     var example = new AzureDevOps.Project("example", new()
    ///     {
    ///         Name = "Example Project",
    ///         Visibility = "private",
    ///         VersionControl = "Git",
    ///         WorkItemTemplate = "Agile",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleServiceEndpointGenericGit = new AzureDevOps.ServiceEndpointGenericGit("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         RepositoryUrl = "https://dev.azure.com/org/project/_git/repository",
    ///         Username = "username",
    ///         Password = "password",
    ///         ServiceEndpointName = "Example Generic Git",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Other Git Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit")]
    public partial class ServiceEndpointGenericGit : global::Pulumi.CustomResource
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
        /// 
        /// &gt; **Note** For AzureDevOps Git, PAT should be used as the password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
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
                AdditionalSecretOutputs =
                {
                    "password",
                },
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

    public sealed class ServiceEndpointGenericGitArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        /// </summary>
        [Input("enablePipelinesAccess")]
        public Input<bool>? EnablePipelinesAccess { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The PAT or password used to authenticate to the git repository.
        /// 
        /// &gt; **Note** For AzureDevOps Git, PAT should be used as the password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the project.
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
        public static new ServiceEndpointGenericGitArgs Empty => new ServiceEndpointGenericGitArgs();
    }

    public sealed class ServiceEndpointGenericGitState : global::Pulumi.ResourceArgs
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

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The PAT or password used to authenticate to the git repository.
        /// 
        /// &gt; **Note** For AzureDevOps Git, PAT should be used as the password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the project.
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
        public static new ServiceEndpointGenericGitState Empty => new ServiceEndpointGenericGitState();
    }
}
