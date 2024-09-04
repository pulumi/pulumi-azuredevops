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
    /// Manages a group within Azure DevOps.
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
    ///     });
    /// 
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_contributors = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Contributors",
    ///     });
    /// 
    ///     var exampleGroup = new AzureDevOps.Group("example", new()
    ///     {
    ///         Scope = example.Id,
    ///         DisplayName = "Example group",
    ///         Description = "Example description",
    ///         Members = new[]
    ///         {
    ///             example_readers.Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Descriptor)),
    ///             example_contributors.Apply(example_contributors =&gt; example_contributors.Apply(getGroupResult =&gt; getGroupResult.Descriptor)),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// - [Azure DevOps Service REST API 7.0 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-7.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: Read, Write, &amp; Manage
    /// 
    /// ## Import
    /// 
    /// Azure DevOps groups can be imported using the group identity descriptor, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuredevops:index/group:Group example aadgp.Uy0xLTktMTU1MTM3NDI0NS0xMjA0NDAwOTY5LTI0MDI5ODY0MTMtMjE3OTQwODYxNi0zLTIxNjc2NjQyNTMtMzI1Nzg0NDI4OS0yMjU4MjcwOTc0LTI2MDYxODY2NDU
    /// ```
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Description of the Project.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The identity (subject) descriptor of the Group.
        /// </summary>
        [Output("descriptor")]
        public Output<string> Descriptor { get; private set; } = null!;

        /// <summary>
        /// The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// This represents the name of the container of origin for a graph member.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The ID of the Group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        /// </summary>
        [Output("mail")]
        public Output<string> Mail { get; private set; } = null!;

        /// <summary>
        /// &gt; NOTE: It's possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        /// </summary>
        [Output("originId")]
        public Output<string> OriginId { get; private set; } = null!;

        /// <summary>
        /// This is the PrincipalName of this graph member from the source provider.
        /// </summary>
        [Output("principalName")]
        public Output<string> PrincipalName { get; private set; } = null!;

        /// <summary>
        /// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// This field identifies the type of the graph subject (ex: Group, Scope, User).
        /// </summary>
        [Output("subjectKind")]
        public Output<string> SubjectKind { get; private set; } = null!;

        /// <summary>
        /// This url is the full route to the source resource of this graph subject.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Description of the Project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        /// </summary>
        [Input("mail")]
        public Input<string>? Mail { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// &gt; NOTE: It's possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Description of the Project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identity (subject) descriptor of the Group.
        /// </summary>
        [Input("descriptor")]
        public Input<string>? Descriptor { get; set; }

        /// <summary>
        /// The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// This represents the name of the container of origin for a graph member.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The ID of the Group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        /// </summary>
        [Input("mail")]
        public Input<string>? Mail { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// &gt; NOTE: It's possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        /// </summary>
        [Input("originId")]
        public Input<string>? OriginId { get; set; }

        /// <summary>
        /// This is the PrincipalName of this graph member from the source provider.
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        /// <summary>
        /// The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// This field identifies the type of the graph subject (ex: Group, Scope, User).
        /// </summary>
        [Input("subjectKind")]
        public Input<string>? SubjectKind { get; set; }

        /// <summary>
        /// This url is the full route to the source resource of this graph subject.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
