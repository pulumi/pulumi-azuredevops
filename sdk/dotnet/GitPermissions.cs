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
    /// Manages permissions for Git repositories.
    /// 
    /// &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
    /// 
    /// ## Permission levels
    /// 
    /// Permission for Git Repositories within Azure DevOps can be applied on three different levels.
    /// Those levels are reflected by specifying (or omitting) values for the arguments `project_id`, `repository_id` and `branch_name`.
    /// 
    /// ### Project level
    /// 
    /// Permissions for all Git Repositories inside a project (existing or newly created ones) are specified, if only the argument `project_id` has a value.
    /// 
    /// #### Example usage
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
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var example_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.GitPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Principal = example_readers.Apply(example_readers =&gt; example_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "CreateRepository", "Deny" },
    ///             { "DeleteRepository", "Deny" },
    ///             { "RenameRepository", "NotSet" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Repository level
    /// 
    /// Permissions for a specific Git Repository and all existing or newly created branches are specified if the arguments `project_id` and `repository_id` are set.
    /// 
    /// #### Example usage
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
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var example_group = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         Name = "Project Collection Administrators",
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Empty Git Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.GitPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = exampleGit.ProjectId,
    ///         RepositoryId = exampleGit.Id,
    ///         Principal = example_group.Apply(example_group =&gt; example_group.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "RemoveOthersLocks", "Allow" },
    ///             { "ManagePermissions", "Deny" },
    ///             { "CreateTag", "Deny" },
    ///             { "CreateBranch", "NotSet" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Branch level
    /// 
    /// Permissions for a specific branch inside a Git Repository are specified if all above mentioned the arguments are set.
    /// 
    /// #### Example usage
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
    ///         WorkItemTemplate = "Agile",
    ///         VersionControl = "Git",
    ///         Visibility = "private",
    ///         Description = "Managed by Pulumi",
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Example Empty Git Repository",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var example_group = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         Name = "Project Collection Administrators",
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.GitPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = exampleGit.ProjectId,
    ///         RepositoryId = exampleGit.Id,
    ///         BranchName = "refs/heads/master",
    ///         Principal = example_group.Apply(example_group =&gt; example_group.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "RemoveOthersLocks", "Allow" },
    ///             { "ForcePush", "Deny" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
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
    ///     var example_project_readers = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Readers",
    ///     });
    /// 
    ///     var example_project_contributors = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Contributors",
    ///     });
    /// 
    ///     var example_project_administrators = AzureDevOps.GetGroup.Invoke(new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "Project administrators",
    ///     });
    /// 
    ///     var example_permissions = new AzureDevOps.GitPermissions("example-permissions", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Principal = example_project_readers.Apply(example_project_readers =&gt; example_project_readers.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "CreateRepository", "Deny" },
    ///             { "DeleteRepository", "Deny" },
    ///             { "RenameRepository", "NotSet" },
    ///         },
    ///     });
    /// 
    ///     var exampleGit = new AzureDevOps.Git("example", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         Name = "TestRepo",
    ///         DefaultBranch = "refs/heads/master",
    ///         Initialization = new AzureDevOps.Inputs.GitInitializationArgs
    ///         {
    ///             InitType = "Clean",
    ///         },
    ///     });
    /// 
    ///     var example_repo_permissions = new AzureDevOps.GitPermissions("example-repo-permissions", new()
    ///     {
    ///         ProjectId = exampleGit.ProjectId,
    ///         RepositoryId = exampleGit.Id,
    ///         Principal = example_project_administrators.Apply(example_project_administrators =&gt; example_project_administrators.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "RemoveOthersLocks", "Allow" },
    ///             { "ManagePermissions", "Deny" },
    ///             { "CreateTag", "Deny" },
    ///             { "CreateBranch", "NotSet" },
    ///         },
    ///     });
    /// 
    ///     var example_branch_permissions = new AzureDevOps.GitPermissions("example-branch-permissions", new()
    ///     {
    ///         ProjectId = exampleGit.ProjectId,
    ///         RepositoryId = exampleGit.Id,
    ///         BranchName = "master",
    ///         Principal = example_project_contributors.Apply(example_project_contributors =&gt; example_project_contributors.Apply(getGroupResult =&gt; getGroupResult.Id)),
    ///         Permissions = 
    ///         {
    ///             { "RemoveOthersLocks", "Allow" },
    ///             { "ForcePush", "Deny" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Relevant Links
    /// 
    /// * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
    /// 
    /// ## PAT Permissions Required
    /// 
    /// - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
    /// 
    /// ## Import
    /// 
    /// The resource does not support import.
    /// </summary>
    [AzureDevOpsResourceType("azuredevops:index/gitPermissions:GitPermissions")]
    public partial class GitPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the branch to assign the permissions. 
        /// 
        /// &gt; **Note** to assign permissions to a branch, the `repository_id` must be set as well.
        /// </summary>
        [Output("branchName")]
        public Output<string?> BranchName { get; private set; } = null!;

        /// <summary>
        /// the permissions to assign. The follwing permissions are available
        /// 
        /// 
        /// | Permissions             | Description                                            |
        /// |-------------------------|--------------------------------------------------------|
        /// | Administer              | Administer                                             |
        /// | GenericRead             | Read                                                   |
        /// | GenericContribute       | Contribute                                             |
        /// | ForcePush               | Force push (rewrite history, delete branches and tags) |
        /// | CreateBranch            | Create branch                                          |
        /// | CreateTag               | Create tag                                             |
        /// | ManageNote              | Manage notes                                           |
        /// | PolicyExempt            | Bypass policies when pushing                           |
        /// | CreateRepository        | Create repository                                      |
        /// | DeleteRepository        | Delete repository                                      |
        /// | RenameRepository        | Rename repository                                      |
        /// | EditPolicies            | Edit policies                                          |
        /// | RemoveOthersLocks       | Remove others' locks                                   |
        /// | ManagePermissions       | Manage permissions                                     |
        /// | PullRequestContribute   | Contribute to pull requests                            |
        /// | PullRequestBypassPolicy | Bypass policies when completing pull requests          |
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableDictionary<string, string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Output("replace")]
        public Output<bool?> Replace { get; private set; } = null!;

        /// <summary>
        /// The ID of the GIT repository to assign the permissions
        /// </summary>
        [Output("repositoryId")]
        public Output<string?> RepositoryId { get; private set; } = null!;


        /// <summary>
        /// Create a GitPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GitPermissions(string name, GitPermissionsArgs args, CustomResourceOptions? options = null)
            : base("azuredevops:index/gitPermissions:GitPermissions", name, args ?? new GitPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GitPermissions(string name, Input<string> id, GitPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("azuredevops:index/gitPermissions:GitPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GitPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GitPermissions Get(string name, Input<string> id, GitPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new GitPermissions(name, id, state, options);
        }
    }

    public sealed class GitPermissionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to assign the permissions. 
        /// 
        /// &gt; **Note** to assign permissions to a branch, the `repository_id` must be set as well.
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        [Input("permissions", required: true)]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The follwing permissions are available
        /// 
        /// 
        /// | Permissions             | Description                                            |
        /// |-------------------------|--------------------------------------------------------|
        /// | Administer              | Administer                                             |
        /// | GenericRead             | Read                                                   |
        /// | GenericContribute       | Contribute                                             |
        /// | ForcePush               | Force push (rewrite history, delete branches and tags) |
        /// | CreateBranch            | Create branch                                          |
        /// | CreateTag               | Create tag                                             |
        /// | ManageNote              | Manage notes                                           |
        /// | PolicyExempt            | Bypass policies when pushing                           |
        /// | CreateRepository        | Create repository                                      |
        /// | DeleteRepository        | Delete repository                                      |
        /// | RenameRepository        | Rename repository                                      |
        /// | EditPolicies            | Edit policies                                          |
        /// | RemoveOthersLocks       | Remove others' locks                                   |
        /// | ManagePermissions       | Manage permissions                                     |
        /// | PullRequestContribute   | Contribute to pull requests                            |
        /// | PullRequestBypassPolicy | Bypass policies when completing pull requests          |
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        /// <summary>
        /// The ID of the GIT repository to assign the permissions
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        public GitPermissionsArgs()
        {
        }
        public static new GitPermissionsArgs Empty => new GitPermissionsArgs();
    }

    public sealed class GitPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to assign the permissions. 
        /// 
        /// &gt; **Note** to assign permissions to a branch, the `repository_id` must be set as well.
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        [Input("permissions")]
        private InputMap<string>? _permissions;

        /// <summary>
        /// the permissions to assign. The follwing permissions are available
        /// 
        /// 
        /// | Permissions             | Description                                            |
        /// |-------------------------|--------------------------------------------------------|
        /// | Administer              | Administer                                             |
        /// | GenericRead             | Read                                                   |
        /// | GenericContribute       | Contribute                                             |
        /// | ForcePush               | Force push (rewrite history, delete branches and tags) |
        /// | CreateBranch            | Create branch                                          |
        /// | CreateTag               | Create tag                                             |
        /// | ManageNote              | Manage notes                                           |
        /// | PolicyExempt            | Bypass policies when pushing                           |
        /// | CreateRepository        | Create repository                                      |
        /// | DeleteRepository        | Delete repository                                      |
        /// | RenameRepository        | Rename repository                                      |
        /// | EditPolicies            | Edit policies                                          |
        /// | RemoveOthersLocks       | Remove others' locks                                   |
        /// | ManagePermissions       | Manage permissions                                     |
        /// | PullRequestContribute   | Contribute to pull requests                            |
        /// | PullRequestBypassPolicy | Bypass policies when completing pull requests          |
        /// </summary>
        public InputMap<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputMap<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The **group** principal to assign the permissions.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The ID of the project to assign the permissions.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Replace (`true`) or merge (`false`) the permissions. Default: `true`
        /// </summary>
        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        /// <summary>
        /// The ID of the GIT repository to assign the permissions
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        public GitPermissionsState()
        {
        }
        public static new GitPermissionsState Empty => new GitPermissionsState();
    }
}
