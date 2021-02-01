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
    /// Manages a Bitbucket service endpoint within Azure DevOps.
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
    ///         var serviceendpoint = new AzureDevOps.ServiceEndpointBitBucket("serviceendpoint", new AzureDevOps.ServiceEndpointBitBucketArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             Username = "username",
    ///             Password = "password",
    ///             ServiceEndpointName = "Sample Bitbucket",
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
    /// Azure DevOps Service Endpoint Bitbucket can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket")]
    public partial class ServiceEndpointBitBucket : Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Bitbucket account password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// A bcrypted hash of the attribute 'password'
        /// </summary>
        [Output("passwordHash")]
        public Output<string> PasswordHash { get; private set; } = null!;

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
        /// Bitbucket account username.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointBitBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointBitBucket(string name, ServiceEndpointBitBucketArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket", name, args ?? new ServiceEndpointBitBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointBitBucket(string name, Input<string> id, ServiceEndpointBitBucketState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azuredevops:ServiceEndpoint/bitBucket:BitBucket"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceEndpointBitBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointBitBucket Get(string name, Input<string> id, ServiceEndpointBitBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointBitBucket(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointBitBucketArgs : Pulumi.ResourceArgs
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
        /// Bitbucket account password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

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

        /// <summary>
        /// Bitbucket account username.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ServiceEndpointBitBucketArgs()
        {
        }
    }

    public sealed class ServiceEndpointBitBucketState : Pulumi.ResourceArgs
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
        /// Bitbucket account password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// A bcrypted hash of the attribute 'password'
        /// </summary>
        [Input("passwordHash")]
        public Input<string>? PasswordHash { get; set; }

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

        /// <summary>
        /// Bitbucket account username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceEndpointBitBucketState()
        {
        }
    }
}
