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
    /// Manages an Incoming WebHook service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to subscribe to a webhook event.
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
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    ///     var exampleServiceendpointIncomingwebhook = new AzureDevOps.ServiceendpointIncomingwebhook("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         WebhookName = "example_webhook",
    ///         Secret = "secret",
    ///         HttpHeader = "X-Hub-Signature",
    ///         ServiceEndpointName = "Example IncomingWebhook",
    ///         Description = "Managed by Terraform",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure DevOps Service Endpoint Incoming WebHook can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook")]
    public partial class ServiceendpointIncomingwebhook : global::Pulumi.CustomResource
    {
        [Output("authorization")]
        public Output<ImmutableDictionary<string, string>> Authorization { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Http header name on which checksum will be sent.
        /// </summary>
        [Output("httpHeader")]
        public Output<string?> HttpHeader { get; private set; } = null!;

        /// <summary>
        /// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        /// </summary>
        [Output("secret")]
        public Output<string?> Secret { get; private set; } = null!;

        /// <summary>
        /// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        /// </summary>
        [Output("serviceEndpointName")]
        public Output<string> ServiceEndpointName { get; private set; } = null!;

        /// <summary>
        /// The name of the WebHook.
        /// </summary>
        [Output("webhookName")]
        public Output<string> WebhookName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceendpointIncomingwebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceendpointIncomingwebhook(string name, ServiceendpointIncomingwebhookArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook", name, args ?? new ServiceendpointIncomingwebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceendpointIncomingwebhook(string name, Input<string> id, ServiceendpointIncomingwebhookState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/serviceendpointIncomingwebhook:ServiceendpointIncomingwebhook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceendpointIncomingwebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceendpointIncomingwebhook Get(string name, Input<string> id, ServiceendpointIncomingwebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceendpointIncomingwebhook(name, id, state, options);
        }
    }

    public sealed class ServiceendpointIncomingwebhookArgs : global::Pulumi.ResourceArgs
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
        /// Http header name on which checksum will be sent.
        /// </summary>
        [Input("httpHeader")]
        public Input<string>? HttpHeader { get; set; }

        /// <summary>
        /// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        /// </summary>
        [Input("serviceEndpointName", required: true)]
        public Input<string> ServiceEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the WebHook.
        /// </summary>
        [Input("webhookName", required: true)]
        public Input<string> WebhookName { get; set; } = null!;

        public ServiceendpointIncomingwebhookArgs()
        {
        }
        public static new ServiceendpointIncomingwebhookArgs Empty => new ServiceendpointIncomingwebhookArgs();
    }

    public sealed class ServiceendpointIncomingwebhookState : global::Pulumi.ResourceArgs
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
        /// Http header name on which checksum will be sent.
        /// </summary>
        [Input("httpHeader")]
        public Input<string>? HttpHeader { get; set; }

        /// <summary>
        /// The ID of the project. Changing this forces a new Service Connection Incoming WebHook to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// Secret for the WebHook. WebHook service will use this secret to calculate the payload checksum.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of the service endpoint. Changing this forces a new Service Connection Incoming WebHook to be created.
        /// </summary>
        [Input("serviceEndpointName")]
        public Input<string>? ServiceEndpointName { get; set; }

        /// <summary>
        /// The name of the WebHook.
        /// </summary>
        [Input("webhookName")]
        public Input<string>? WebhookName { get; set; }

        public ServiceendpointIncomingwebhookState()
        {
        }
        public static new ServiceendpointIncomingwebhookState Empty => new ServiceendpointIncomingwebhookState();
    }
}
