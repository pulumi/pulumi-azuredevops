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
    /// Manages a AWS service endpoint within Azure DevOps. Using this service endpoint requires you to first install [AWS Toolkit for Azure DevOps](https://marketplace.visualstudio.com/items?itemName=AmazonWebServices.aws-vsts-tools).
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
    ///         var serviceendpoint = new AzureDevOps.ServiceEndpointAws("serviceendpoint", new AzureDevOps.ServiceEndpointAwsArgs
    ///         {
    ///             ProjectId = project.Id,
    ///             ServiceEndpointName = "aws_serviceendpoint",
    ///             Description = "Managed by AzureDevOps",
    ///             AccessKeyId = "xxxx",
    ///             SecretAccessKey = "xxxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Relevant Links
    /// 
    /// * [aws-toolkit-azure-devops](https://github.com/aws/aws-toolkit-azure-devops)
    /// * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
    /// </summary>
    public partial class ServiceEndpointAws : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS access key ID for signing programmatic requests.
        /// </summary>
        [Output("accessKeyId")]
        public Output<string> AccessKeyId { get; private set; } = null!;

        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
        /// </summary>
        [Output("externalId")]
        public Output<string?> ExternalId { get; private set; } = null!;

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Optional identifier for the assumed role session.
        /// </summary>
        [Output("roleSessionName")]
        public Output<string?> RoleSessionName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the role to assume.
        /// </summary>
        [Output("roleToAssume")]
        public Output<string?> RoleToAssume { get; private set; } = null!;

        /// <summary>
        /// The AWS secret access key for signing programmatic requests.
        /// </summary>
        [Output("secretAccessKey")]
        public Output<string> SecretAccessKey { get; private set; } = null!;

        /// <summary>
        /// A bcrypted hash of the attribute 'secret_access_key'
        /// </summary>
        [Output("secretAccessKeyHash")]
        public Output<string> SecretAccessKeyHash { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The AWS session token for signing programmatic requests.
        /// </summary>
        [Output("sessionToken")]
        public Output<string?> SessionToken { get; private set; } = null!;

        /// <summary>
        /// A bcrypted hash of the attribute 'session_token'
        /// </summary>
        [Output("sessionTokenHash")]
        public Output<string> SessionTokenHash { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEndpointAws resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEndpointAws(string name, ServiceEndpointAwsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAws:ServiceEndpointAws", name, args ?? new ServiceEndpointAwsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEndpointAws(string name, Input<string> id, ServiceEndpointAwsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceEndpointAws:ServiceEndpointAws", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEndpointAws resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEndpointAws Get(string name, Input<string> id, ServiceEndpointAwsState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEndpointAws(name, id, state, options);
        }
    }

    public sealed class ServiceEndpointAwsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS access key ID for signing programmatic requests.
        /// </summary>
        [Input("accessKeyId", required: true)]
        public Input<string> AccessKeyId { get; set; } = null!;

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
        /// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Optional identifier for the assumed role session.
        /// </summary>
        [Input("roleSessionName")]
        public Input<string>? RoleSessionName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the role to assume.
        /// </summary>
        [Input("roleToAssume")]
        public Input<string>? RoleToAssume { get; set; }

        /// <summary>
        /// The AWS secret access key for signing programmatic requests.
        /// </summary>
        [Input("secretAccessKey", required: true)]
        public Input<string> SecretAccessKey { get; set; } = null!;

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// The AWS session token for signing programmatic requests.
        /// </summary>
        [Input("sessionToken")]
        public Input<string>? SessionToken { get; set; }

        public ServiceEndpointAwsArgs()
        {
        }
    }

    public sealed class ServiceEndpointAwsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS access key ID for signing programmatic requests.
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

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
        /// A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The project ID or project name.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Optional identifier for the assumed role session.
        /// </summary>
        [Input("roleSessionName")]
        public Input<string>? RoleSessionName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the role to assume.
        /// </summary>
        [Input("roleToAssume")]
        public Input<string>? RoleToAssume { get; set; }

        /// <summary>
        /// The AWS secret access key for signing programmatic requests.
        /// </summary>
        [Input("secretAccessKey")]
        public Input<string>? SecretAccessKey { get; set; }

        /// <summary>
        /// A bcrypted hash of the attribute 'secret_access_key'
        /// </summary>
        [Input("secretAccessKeyHash")]
        public Input<string>? SecretAccessKeyHash { get; set; }

        /// <summary>
        /// The Service Endpoint name.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The AWS session token for signing programmatic requests.
        /// </summary>
        [Input("sessionToken")]
        public Input<string>? SessionToken { get; set; }

        /// <summary>
        /// A bcrypted hash of the attribute 'session_token'
        /// </summary>
        [Input("sessionTokenHash")]
        public Input<string>? SessionTokenHash { get; set; }

        public ServiceEndpointAwsState()
        {
        }
    }
}
