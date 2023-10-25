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
    /// Manages authorization of resources, e.g. for access in build pipelines.
    /// 
    /// Currently supported resources: service endpoint (aka service connection, endpoint).
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Authorize Definition Resource](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/resources/authorize%20definition%20resources?view=azure-devops-rest-7.0)
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/resourceAuthorization:ResourceAuthorization")]
    public partial class ResourceAuthorization : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set to true to allow public access in the project. Type: boolean.
        /// </summary>
        [Output("authorized")]
        public Output<bool> Authorized { get; private set; } = null!;

        /// <summary>
        /// The ID of the build definition to authorize. Type: string.
        /// </summary>
        [Output("definitionId")]
        public Output<int?> DefinitionId { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name. Type: string.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource to authorize. Type: string.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceAuthorization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceAuthorization(string name, ResourceAuthorizationArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/resourceAuthorization:ResourceAuthorization", name, args ?? new ResourceAuthorizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceAuthorization(string name, Input<string> id, ResourceAuthorizationState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/resourceAuthorization:ResourceAuthorization", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azuredevops:Security/resourceAuthorization:ResourceAuthorization"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourceAuthorization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceAuthorization Get(string name, Input<string> id, ResourceAuthorizationState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceAuthorization(name, id, state, options);
        }
    }

    public sealed class ResourceAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true to allow public access in the project. Type: boolean.
        /// </summary>
        [Input("authorized", required: true)]
        public Input<bool> Authorized { get; set; } = null!;

        /// <summary>
        /// The ID of the build definition to authorize. Type: string.
        /// </summary>
        [Input("definitionId")]
        public Input<int>? DefinitionId { get; set; }

        /// <summary>
        /// The project ID or project name. Type: string.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the resource to authorize. Type: string.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResourceAuthorizationArgs()
        {
        }
        public static new ResourceAuthorizationArgs Empty => new ResourceAuthorizationArgs();
    }

    public sealed class ResourceAuthorizationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true to allow public access in the project. Type: boolean.
        /// </summary>
        [Input("authorized")]
        public Input<bool>? Authorized { get; set; }

        /// <summary>
        /// The ID of the build definition to authorize. Type: string.
        /// </summary>
        [Input("definitionId")]
        public Input<int>? DefinitionId { get; set; }

        /// <summary>
        /// The project ID or project name. Type: string.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the resource to authorize. Type: string.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResourceAuthorizationState()
        {
        }
        public static new ResourceAuthorizationState Empty => new ResourceAuthorizationState();
    }
}
