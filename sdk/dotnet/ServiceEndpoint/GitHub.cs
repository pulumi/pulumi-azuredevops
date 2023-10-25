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
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint GitHub can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:ServiceEndpoint/gitHub:GitHub example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [Obsolete(@"azuredevops.serviceendpoint.GitHub has been deprecated in favor of azuredevops.ServiceEndpointGitHub")]
    [AzureDevOpsResourceType("azuredevops:ServiceEndpoint/gitHub:GitHub")]
    public partial class GitHub : global::Pulumi.CustomResource
    {
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

    public sealed class GitHubArgs : global::Pulumi.ResourceArgs
    {
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
        /// The ID of the project.
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
        public static new GitHubArgs Empty => new GitHubArgs();
    }

    public sealed class GitHubState : global::Pulumi.ResourceArgs
    {
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
        /// The ID of the project.
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
        public static new GitHubState Empty => new GitHubState();
    }
}
