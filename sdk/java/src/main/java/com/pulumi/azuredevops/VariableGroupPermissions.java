// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.VariableGroupPermissionsArgs;
import com.pulumi.azuredevops.inputs.VariableGroupPermissionsState;
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
 * Manages permissions for a Variable Group
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
 * import com.pulumi.azuredevops.VariableGroup;
 * import com.pulumi.azuredevops.VariableGroupArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupVariableArgs;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.inputs.GetGroupArgs;
 * import com.pulumi.azuredevops.VariableGroupPermissions;
 * import com.pulumi.azuredevops.VariableGroupPermissionsArgs;
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
 *         var project = new Project("project", ProjectArgs.builder()
 *             .name("Testing")
 *             .description("Testing-description")
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .build());
 * 
 *         var example = new VariableGroup("example", VariableGroupArgs.builder()
 *             .projectId(project.id())
 *             .name("test")
 *             .description("Test Description")
 *             .allowAccess(true)
 *             .variables(VariableGroupVariableArgs.builder()
 *                 .name("key1")
 *                 .value("val1")
 *                 .build())
 *             .build());
 * 
 *         final var tf-project-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(project.id())
 *             .name("Readers")
 *             .build());
 * 
 *         var permissions = new VariableGroupPermissions("permissions", VariableGroupPermissionsArgs.builder()
 *             .projectId(project.id())
 *             .variableGroupId(example.id())
 *             .principal(tf_project_readers.applyValue(tf_project_readers -> tf_project_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry("View", "allow"),
 *                 Map.entry("Administer", "allow"),
 *                 Map.entry("Use", "allow")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Roles
 * 
 * The Azure DevOps UI uses roles to assign permissions for variable groups.
 * 
 * | Role          | Allow Permissions     |
 * |---------------|-----------------------|
 * | Reader        | View                  |
 * | User          | View, Use             |
 * | Administrator | View, Use, Administer |
 * 
 * ## Relevant Links
 * 
 * * [Azure DevOps Service REST API 7.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.1)
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
@ResourceType(type="azuredevops:index/variableGroupPermissions:VariableGroupPermissions")
public class VariableGroupPermissions extends com.pulumi.resources.CustomResource {
    /**
     * the permissions to assign. The following permissions are available.
     * 
     * | Permission  | Description               |
     * |-------------|---------------------------|
     * | View        | View library item         |
     * | Administer  | Administer library item   |
     * | Create      | Create library item       |
     * | ViewSecrets | View library item secrets |
     * | Use         | Use library item          |
     * | Owner       | Owner library item        |
     * 
     */
    @Export(name="permissions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     * | Permission  | Description               |
     * |-------------|---------------------------|
     * | View        | View library item         |
     * | Administer  | Administer library item   |
     * | Create      | Create library item       |
     * | ViewSecrets | View library item secrets |
     * | Use         | Use library item          |
     * | Owner       | Owner library item        |
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
     * The ID of the project.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project.
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
     * The id of the variable group to assign the permissions.
     * 
     */
    @Export(name="variableGroupId", refs={String.class}, tree="[0]")
    private Output<String> variableGroupId;

    /**
     * @return The id of the variable group to assign the permissions.
     * 
     */
    public Output<String> variableGroupId() {
        return this.variableGroupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VariableGroupPermissions(java.lang.String name) {
        this(name, VariableGroupPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VariableGroupPermissions(java.lang.String name, VariableGroupPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VariableGroupPermissions(java.lang.String name, VariableGroupPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/variableGroupPermissions:VariableGroupPermissions", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VariableGroupPermissions(java.lang.String name, Output<java.lang.String> id, @Nullable VariableGroupPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/variableGroupPermissions:VariableGroupPermissions", name, state, makeResourceOptions(options, id), false);
    }

    private static VariableGroupPermissionsArgs makeArgs(VariableGroupPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VariableGroupPermissionsArgs.Empty : args;
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
    public static VariableGroupPermissions get(java.lang.String name, Output<java.lang.String> id, @Nullable VariableGroupPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VariableGroupPermissions(name, id, state, options);
    }
}
