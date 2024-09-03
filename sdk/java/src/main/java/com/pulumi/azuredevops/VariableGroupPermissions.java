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
 * &lt;table&gt;
 * &lt;thead&gt;
 * &lt;tr&gt;
 * &lt;th&gt;Role&lt;/th&gt;
 * &lt;th&gt;Allow Permissions&lt;/th&gt;
 * &lt;/tr&gt;
 * &lt;/thead&gt;
 * &lt;tbody&gt;
 * &lt;tr&gt;
 * &lt;td&gt;Reader&lt;/td&gt;
 * &lt;td&gt;View&lt;/td&gt;
 * &lt;/tr&gt;
 * &lt;tr&gt;
 * &lt;td&gt;User&lt;/td&gt;
 * &lt;td&gt;View, Use&lt;/td&gt;
 * &lt;/tr&gt;
 * &lt;tr&gt;
 * &lt;td&gt;Administrator&lt;/td&gt;
 * &lt;td&gt;View, Use, Administer&lt;/td&gt;
 * &lt;/tr&gt;
 * &lt;/tbody&gt;
 * &lt;/table&gt;
 * 
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
@ResourceType(type="azuredevops:index/variableGroupPermissions:VariableGroupPermissions")
public class VariableGroupPermissions extends com.pulumi.resources.CustomResource {
    /**
     * the permissions to assign. The following permissions are available.
     * 
     */
    @Export(name="permissions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
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
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;View&lt;/td&gt;
     * &lt;td&gt;View library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Administer&lt;/td&gt;
     * &lt;td&gt;Administer library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Create&lt;/td&gt;
     * &lt;td&gt;Create library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewSecrets&lt;/td&gt;
     * &lt;td&gt;View library item secrets&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Use&lt;/td&gt;
     * &lt;td&gt;Use library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Owner&lt;/td&gt;
     * &lt;td&gt;Owner library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
     * 
     */
    @Export(name="replace", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * &lt;table&gt;
     * &lt;thead&gt;
     * &lt;tr&gt;
     * &lt;th&gt;Permission&lt;/th&gt;
     * &lt;th&gt;Description&lt;/th&gt;
     * &lt;/tr&gt;
     * &lt;/thead&gt;
     * &lt;tbody&gt;
     * &lt;tr&gt;
     * &lt;td&gt;View&lt;/td&gt;
     * &lt;td&gt;View library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Administer&lt;/td&gt;
     * &lt;td&gt;Administer library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Create&lt;/td&gt;
     * &lt;td&gt;Create library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;ViewSecrets&lt;/td&gt;
     * &lt;td&gt;View library item secrets&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Use&lt;/td&gt;
     * &lt;td&gt;Use library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;tr&gt;
     * &lt;td&gt;Owner&lt;/td&gt;
     * &lt;td&gt;Owner library item&lt;/td&gt;
     * &lt;/tr&gt;
     * &lt;/tbody&gt;
     * &lt;/table&gt;
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
