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
    /// Manages Wikis within Azure DevOps project.
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
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var exampleWiki = new AzureDevOps.Wiki("example", new()
    ///     {
    ///         Name = "Example project wiki ",
    ///         ProjectId = example.Id,
    ///         Type = "projectWiki",
    ///     });
    /// 
    ///     var example2 = new AzureDevOps.Wiki("example2", new()
    ///     {
    ///         Name = "Example wiki in repository",
    ///         ProjectId = example.Id,
    ///         RepositoryId = exampleGit.Id,
    ///         Version = "main",
    ///         Type = "codeWiki",
    ///         Mappedpath = "/",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.1 - Wiki ](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/wikis?view=azure-devops-rest-7.1)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Wiki can be imported using the `id`
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/wiki:Wiki wiki 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/wiki:Wiki")]
    public partial class Wiki : global::Pulumi.CustomResource
    {
        [Output("mappedPath")]
        public Output<string> MappedPath { get; private set; } = null!;

        /// <summary>
        /// The name of the Wiki.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Project.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The remote web url to the wiki.
        /// </summary>
        [Output("remoteUrl")]
        public Output<string> RemoteUrl { get; private set; } = null!;

        /// <summary>
        /// The ID of the repository.
        /// </summary>
        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;

        /// <summary>
        /// The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The REST url for this wiki.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Version of the wiki.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Wiki resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Wiki(string name, WikiArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/wiki:Wiki", name, args ?? new WikiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Wiki(string name, Input<string> id, WikiState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/wiki:Wiki", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Wiki resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Wiki Get(string name, Input<string> id, WikiState? state = null, CustomResourceOptions? options = null)
        {
            return new Wiki(name, id, state, options);
        }
    }

    public sealed class WikiArgs : global::Pulumi.ResourceArgs
    {
        [Input("mappedPath")]
        public Input<string>? MappedPath { get; set; }

        /// <summary>
        /// The name of the Wiki.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the repository.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Version of the wiki.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public WikiArgs()
        {
        }
        public static new WikiArgs Empty => new WikiArgs();
    }

    public sealed class WikiState : global::Pulumi.ResourceArgs
    {
        [Input("mappedPath")]
        public Input<string>? MappedPath { get; set; }

        /// <summary>
        /// The name of the Wiki.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The remote web url to the wiki.
        /// </summary>
        [Input("remoteUrl")]
        public Input<string>? RemoteUrl { get; set; }

        /// <summary>
        /// The ID of the repository.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The REST url for this wiki.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Version of the wiki.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public WikiState()
        {
        }
        public static new WikiState Empty => new WikiState();
    }
}
