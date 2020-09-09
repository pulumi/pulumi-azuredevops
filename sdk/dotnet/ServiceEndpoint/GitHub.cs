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
    /// Manages a GitHub service endpoint within Azure DevOps.
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
    ///             ProjectName = "Sample Project",
    ///             Visibility = "private",
    ///             VersionControl = "Git",
    ///             WorkItemTemplate = "Agile",
    ///         });
    ///         var serviceendpointGh1 = new AzureDevOps.GitHub("serviceendpointGh1", new AzureDevOps.GitHubArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "Sample GithHub Personal Access Token",
    ///             AuthPersonal = new AzureDevOps.Inputs.GitHubAuthPersonalArgs
    ///             {
    ///                 PersonalAccessToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceendpointGh2 = new AzureDevOps.GitHub("serviceendpointGh2", new AzureDevOps.GitHubArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             ServiceEndpointName = "Sample GithHub Grant",
    ///             AuthOauth = new AzureDevOps.Inputs.GitHubAuthOauthArgs
    ///             {
    ///                 OauthConfigurationId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureDevOps = Pulumi.AzureDevOps;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceendpointGh3 = new AzureDevOps.GitHub("serviceendpointGh3", new AzureDevOps.GitHubArgs
    ///         {
    ///             ProjectId = azuredevops_project.Project.Id,
    ///             ServiceEndpointName = "Sample GithHub Apps: Azure Pipelines",
    ///             Description = "",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
    /// </summary>
    [Obsolete(@"azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.GitHub")]
    public partial class GitHub : Pulumi.CustomResource
    {
        /// <summary>
        /// An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
        /// </summary>
        [Output("authOauth")]
        public Output<Outputs.GitHubAuthOauth?> AuthOauth { get; private set; } = null!;

        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Output("authPersonal")]
        public Output<Outputs.GitHubAuthPersonal?> AuthPersonal { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;


        /// <summary>
        /// Create a GitHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GitHub(string name, GitHubArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:ServiceEndpoint/gitHub:GitHub", name, args ?? new GitHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GitHub(string name, Input<string> id, GitHubState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:ServiceEndpoint/gitHub:GitHub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GitHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GitHub Get(string name, Input<string> id, GitHubState? state = null, CustomResourceOptions? options = null)
        {
            return new GitHub(name, id, state, options);
        }
    }

    public sealed class GitHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
        /// </summary>
        [Input("authOauth")]
        public Input<Inputs.GitHubAuthOauthArgs>? AuthOauth { get; set; }

        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal")]
        public Input<Inputs.GitHubAuthPersonalArgs>? AuthPersonal { get; set; }

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
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        public GitHubArgs()
        {
        }
    }

    public sealed class GitHubState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `auth_oauth` block as documented below. Allows connecting using an Oauth token.
        /// </summary>
        [Input("authOauth")]
        public Input<Inputs.GitHubAuthOauthGetArgs>? AuthOauth { get; set; }

        /// <summary>
        /// An `auth_personal` block as documented below. Allows connecting using a personal access token.
        /// </summary>
        [Input("authPersonal")]
        public Input<Inputs.GitHubAuthPersonalGetArgs>? AuthPersonal { get; set; }

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
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        public GitHubState()
        {
        }
    }
}
