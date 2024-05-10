// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.WorkItemQueryPermissionsArgs;
import com.pulumi.azuredevops.inputs.WorkItemQueryPermissionsState;
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
 * Manages permissions for Work Item Queries.
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
 * 
 * ## Permission levels
 * 
 * Permission for Work Item Queries within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
 * 
 * ### Project level
 * 
 * Permissions for all Work Item Queries inside a project (existing or newly created ones) are specified, if only the argument `project_id` has a value.
 * 
 * #### Example usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.WorkItemQueryPermissions;
 * import com.pulumi.azuredevops.WorkItemQueryPermissionsArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()        
 *             .name("Example Project")
 *             .workItemTemplate("Agile")
 *             .versionControl("Git")
 *             .visibility("private")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Readers")
 *             .build());
 * 
 *         var project_wiq_root_permissions = new WorkItemQueryPermissions("project-wiq-root-permissions", WorkItemQueryPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .principal(example_readers.applyValue(example_readers -> example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry("CreateRepository", "Deny"),
 *                 Map.entry("DeleteRepository", "Deny"),
 *                 Map.entry("RenameRepository", "NotSet")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Shared Queries folder level
 * 
 * Permissions for a specific folder inside Shared Queries are specified if the arguments `project_id` and `path` are set.
 * 
 * &gt; **Note** To set permissions for the Shared Queries folder itself use `/` as path value
 * 
 * #### Example usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.WorkItemQueryPermissions;
 * import com.pulumi.azuredevops.WorkItemQueryPermissionsArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()        
 *             .name("Example Project")
 *             .workItemTemplate("Agile")
 *             .versionControl("Git")
 *             .visibility("private")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Readers")
 *             .build());
 * 
 *         var example_permissions = new WorkItemQueryPermissions("example-permissions", WorkItemQueryPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .path("/Team")
 *             .principal(example_readers.applyValue(example_readers -> example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry("Contribute", "Allow"),
 *                 Map.entry("Delete", "Deny"),
 *                 Map.entry("Read", "NotSet")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.ProjectArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.WorkItemQueryPermissions;
 * import com.pulumi.azuredevops.WorkItemQueryPermissionsArgs;
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
 *         var example = new Project("example", ProjectArgs.builder()        
 *             .name("Example Project")
 *             .workItemTemplate("Agile")
 *             .versionControl("Git")
 *             .visibility("private")
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         final var example-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Readers")
 *             .build());
 * 
 *         final var example-contributors = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Contributors")
 *             .build());
 * 
 *         var example_project_permissions = new WorkItemQueryPermissions("example-project-permissions", WorkItemQueryPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .principal(example_readers.applyValue(example_readers -> example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry("Read", "Allow"),
 *                 Map.entry("Delete", "Deny"),
 *                 Map.entry("Contribute", "Deny"),
 *                 Map.entry("ManagePermissions", "Deny")
 *             ))
 *             .build());
 * 
 *         var example_sharedqueries_permissions = new WorkItemQueryPermissions("example-sharedqueries-permissions", WorkItemQueryPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .path("/")
 *             .principal(example_contributors.applyValue(example_contributors -> example_contributors.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry("Read", "Allow"),
 *                 Map.entry("Delete", "Deny")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
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
@ResourceType(type="azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions")
public class WorkItemQueryPermissions extends com.pulumi.resources.CustomResource {
    /**
     * Path to a query or folder beneath `Shared Queries`
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Path to a query or folder beneath `Shared Queries`
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * the permissions to assign. The following permissions are available
     * 
     * | Permissions              | Description                        |
     * |--------------------------|------------------------------------|
     * | Read                     | Read                               |
     * | Contribute               | Contribute                         |
     * | Delete                   | Delete                             |
     * | ManagePermissions        | Manage Permissions                 |
     * 
     */
    @Export(name="permissions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available
     * 
     * | Permissions              | Description                        |
     * |--------------------------|------------------------------------|
     * | Read                     | Read                               |
     * | Contribute               | Contribute                         |
     * | Delete                   | Delete                             |
     * | ManagePermissions        | Manage Permissions                 |
     * 
     */
    public Output<Map<String,String>> permissions() {
        return this.permissions;
    }
    /**
     * The **group** principal to assign the permissions.
     * 
     */
    @Export(name="principal", refs={String.class}, tree="[0]")
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
    @Export(name="projectId", refs={String.class}, tree="[0]")
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
    @Export(name="replace", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WorkItemQueryPermissions(String name) {
        this(name, WorkItemQueryPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WorkItemQueryPermissions(String name, WorkItemQueryPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WorkItemQueryPermissions(String name, WorkItemQueryPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions", name, args == null ? WorkItemQueryPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private WorkItemQueryPermissions(String name, Output<String> id, @Nullable WorkItemQueryPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/workItemQueryPermissions:WorkItemQueryPermissions", name, state, makeResourceOptions(options, id));
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
    public static WorkItemQueryPermissions get(String name, Output<String> id, @Nullable WorkItemQueryPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WorkItemQueryPermissions(name, id, state, options);
    }
}
