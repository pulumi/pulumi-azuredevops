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
    /// using System.Collections.Generic;
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
    ///     var exampleServiceEndpointBitBucket = new AzureDevOps.ServiceEndpointBitBucket("exampleServiceEndpointBitBucket", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Username = "username",
    ///         Password = "password",
    ///         ServiceEndpointName = "Example Bitbucket",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Bitbucket can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    ///  $ pulumi import azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket")]
    public partial class ServiceEndpointBitBucket : global::Pulumi.CustomResource
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
                    new global::Pulumi.Alias { Type = "azuredevops:ServiceEndpoint/bitBucket:BitBucket"},
                },
                AdditionalSecretOutputs =
                {
                    "password",
                    "passwordHash",
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

    public sealed class ServiceEndpointBitBucketArgs : global::Pulumi.ResourceArgs
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

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Bitbucket account password.
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
        public static new ServiceEndpointBitBucketArgs Empty => new ServiceEndpointBitBucketArgs();
    }

    public sealed class ServiceEndpointBitBucketState : global::Pulumi.ResourceArgs
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

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Bitbucket account password.
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

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// A bcrypted hash of the attribute 'password'
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        /// <summary>
        /// Bitbucket account username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ServiceEndpointBitBucketState()
        {
        }
        public static new ServiceEndpointBitBucketState Empty => new ServiceEndpointBitBucketState();
    }
}
