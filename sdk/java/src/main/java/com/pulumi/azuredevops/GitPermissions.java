// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.GitPermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.GitPermissionsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages permissions for Git repositories.
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
 * 
 * ## Permission levels
 * 
 * Permission for Git Repositories within Azure DevOps can be applied on three different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `project_id`, `repository_id` and `branch_name`.
 * 
 * ### Project level
 * 
 * Permissions for all Git Repositories inside a project (existing or newly created ones) are specified, if only the argument `project_id` has a value.
 * 
 * #### Example usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.GitPermissions;
 * import com.pulumi.azuredevops.GitPermissionsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Project(&#34;example&#34;, ProjectArgs.builder()        
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         var example_permissions = new GitPermissions(&#34;example-permissions&#34;, GitPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;CreateRepository&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;DeleteRepository&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;RenameRepository&#34;, &#34;NotSet&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### Repository level
 * 
 * Permissions for a specific Git Repository and all existing or newly created branches are specified if the arguments `project_id` and `repository_id` are set.
 * 
 * #### Example usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.GitPermissions;
 * import com.pulumi.azuredevops.GitPermissionsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         final var example-group = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .name(&#34;Project Collection Administrators&#34;)
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         var example_permissions = new GitPermissions(&#34;example-permissions&#34;, GitPermissionsArgs.builder()        
 *             .projectId(exampleGit.projectId())
 *             .repositoryId(exampleGit.id())
 *             .principal(example_group.id())
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;RemoveOthersLocks&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;ManagePermissions&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;CreateTag&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;CreateBranch&#34;, &#34;NotSet&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### Branch level
 * 
 * Permissions for a specific branch inside a Git Repository are specified if all above mentioned the arguments are set.
 * 
 * #### Example usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.GitPermissions;
 * import com.pulumi.azuredevops.GitPermissionsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         final var example-group = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .name(&#34;Project Collection Administrators&#34;)
 *             .build());
 * 
 *         var example_permissions = new GitPermissions(&#34;example-permissions&#34;, GitPermissionsArgs.builder()        
 *             .projectId(exampleGit.projectId())
 *             .repositoryId(exampleGit.id())
 *             .branchName(&#34;refs/heads/master&#34;)
 *             .principal(example_group.id())
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;RemoveOthersLocks&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;ForcePush&#34;, &#34;Deny&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.GitPermissions;
 * import com.pulumi.azuredevops.GitPermissionsArgs;
 * import com.pulumi.azuredevops.Git;
 * import com.pulumi.azuredevops.GitArgs;
 * import com.pulumi.azuredevops.inputs.GitInitializationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         final var example-project-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         final var example-project-contributors = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Contributors&#34;)
 *             .build());
 * 
 *         final var example-project-administrators = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(exampleProject.id())
 *             .name(&#34;Project administrators&#34;)
 *             .build());
 * 
 *         var example_permissions = new GitPermissions(&#34;example-permissions&#34;, GitPermissionsArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .principal(example_project_readers.applyValue(example_project_readers -&gt; example_project_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;CreateRepository&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;DeleteRepository&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;RenameRepository&#34;, &#34;NotSet&#34;)
 *             ))
 *             .build());
 * 
 *         var exampleGit = new Git(&#34;exampleGit&#34;, GitArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .defaultBranch(&#34;refs/heads/master&#34;)
 *             .initialization(GitInitializationArgs.builder()
 *                 .initType(&#34;Clean&#34;)
 *                 .build())
 *             .build());
 * 
 *         var example_repo_permissions = new GitPermissions(&#34;example-repo-permissions&#34;, GitPermissionsArgs.builder()        
 *             .projectId(exampleGit.projectId())
 *             .repositoryId(exampleGit.id())
 *             .principal(example_project_administrators.applyValue(example_project_administrators -&gt; example_project_administrators.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;RemoveOthersLocks&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;ManagePermissions&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;CreateTag&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;CreateBranch&#34;, &#34;NotSet&#34;)
 *             ))
 *             .build());
 * 
 *         var example_branch_permissions = new GitPermissions(&#34;example-branch-permissions&#34;, GitPermissionsArgs.builder()        
 *             .projectId(exampleGit.projectId())
 *             .repositoryId(exampleGit.id())
 *             .branchName(&#34;master&#34;)
 *             .principal(example_project_contributors.applyValue(example_project_contributors -&gt; example_project_contributors.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;RemoveOthersLocks&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;ForcePush&#34;, &#34;Deny&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Project &amp; Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 * 
 * ## Import
 * 
 * The resource does not support import.
 * 
 */
@ResourceType(type="azuredevops:index/gitPermissions:GitPermissions")
public class GitPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The name of the branch to assign the permissions.
     * 
     */
    @Export(name="branchName", type=String.class, parameters={})
    private Output</* @Nullable */ String> branchName;

    /**
     * @return The name of the branch to assign the permissions.
     * 
     */
    public Output<Optional<String>> branchName() {
        return Codegen.optional(this.branchName);
    }
    /**
     * the permissions to assign. The follwing permissions are available
     * 
     */
    @Export(name="permissions", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The follwing permissions are available
     * 
     */
    public Output<Map<String,String>> permissions() {
        return this.permissions;
    }
    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Export(name="principal", type=String.class, parameters={})
    private Output<String> principal;

    /**
     * @return The **group** principal to assign the permissions.
     * 
     */
    public Output<String> principal() {
        return this.principal;
    }
    /**
     * The ID of the project to assign the permissions.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The ID of the project to assign the permissions.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     */
    @Export(name="replace", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }
    /**
     * The ID of the GIT repository to assign the permissions
     * 
     */
    @Export(name="repositoryId", type=String.class, parameters={})
    private Output</* @Nullable */ String> repositoryId;

    /**
     * @return The ID of the GIT repository to assign the permissions
     * 
     */
    public Output<Optional<String>> repositoryId() {
        return Codegen.optional(this.repositoryId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GitPermissions(String name) {
        this(name, GitPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GitPermissions(String name, GitPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GitPermissions(String name, GitPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/gitPermissions:GitPermissions", name, args == null ? GitPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GitPermissions(String name, Output<String> id, @Nullable GitPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/gitPermissions:GitPermissions", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static GitPermissions get(String name, Output<String> id, @Nullable GitPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GitPermissions(name, id, state, options);
    }
}