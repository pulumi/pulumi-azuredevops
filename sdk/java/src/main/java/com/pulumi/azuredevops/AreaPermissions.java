// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.AreaPermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.AreaPermissionsState;
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
 * Manages permissions for an Area (Component)
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
 * 
 * ## Permission levels
 * 
 * Permission for Areas within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
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
 * import com.pulumi.azuredevops.AreaPermissions;
 * import com.pulumi.azuredevops.AreaPermissionsArgs;
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
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *         final var example-project-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(example.id())
 *             .name("Readers")
 *             .build());
 * 
 *         var example_root_permissions = new AreaPermissions("example-root-permissions", AreaPermissionsArgs.builder()
 *             .projectId(example.id())
 *             .principal(example_project_readers.applyValue(_example_project_readers -> _example_project_readers.id()))
 *             .path("/")
 *             .permissions(Map.ofEntries(
 *                 Map.entry("CREATE_CHILDREN", "Deny"),
 *                 Map.entry("GENERIC_READ", "Allow"),
 *                 Map.entry("DELETE", "Deny"),
 *                 Map.entry("WORK_ITEM_READ", "Allow")
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
@ResourceType(type="azuredevops:index/areaPermissions:AreaPermissions")
public class AreaPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The name of the branch to assign the permissions.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The name of the branch to assign the permissions.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * the permissions to assign. The following permissions are available.
     * 
     * | Permission             | Description                          |
     * |------------------------|--------------------------------------|
     * | GENERIC_READ           | View permissions for this node       |
     * | GENERIC_WRITE          | Edit this node                       |
     * | CREATE_CHILDREN        | Create child nodes                   |
     * | DELETE                 | Delete this node                     |
     * | WORK_ITEM_READ         | View work items in this node         |
     * | WORK_ITEM_WRITE        | Edit work items in this node         |
     * | MANAGE_TEST_PLANS      | Manage test plans                    |
     * | MANAGE_TEST_SUITES     | Manage test suites                   |
     * | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
     * 
     */
    @Export(name="permissions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     * | Permission             | Description                          |
     * |------------------------|--------------------------------------|
     * | GENERIC_READ           | View permissions for this node       |
     * | GENERIC_WRITE          | Edit this node                       |
     * | CREATE_CHILDREN        | Create child nodes                   |
     * | DELETE                 | Delete this node                     |
     * | WORK_ITEM_READ         | View work items in this node         |
     * | WORK_ITEM_WRITE        | Edit work items in this node         |
     * | MANAGE_TEST_PLANS      | Manage test plans                    |
     * | MANAGE_TEST_SUITES     | Manage test suites                   |
     * | WORK_ITEM_SAVE_COMMENT | Edit work item comments in this node |
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
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * 
     */
    @Export(name="replace", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AreaPermissions(java.lang.String name) {
        this(name, AreaPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AreaPermissions(java.lang.String name, AreaPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AreaPermissions(java.lang.String name, AreaPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/areaPermissions:AreaPermissions", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AreaPermissions(java.lang.String name, Output<java.lang.String> id, @Nullable AreaPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/areaPermissions:AreaPermissions", name, state, makeResourceOptions(options, id), false);
    }

    private static AreaPermissionsArgs makeArgs(AreaPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AreaPermissionsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static AreaPermissions get(java.lang.String name, Output<java.lang.String> id, @Nullable AreaPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AreaPermissions(name, id, state, options);
    }
}
