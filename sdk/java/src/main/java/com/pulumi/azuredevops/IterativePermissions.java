// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.IterativePermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.IterativePermissionsState;
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
 * Manages permissions for an Iteration (Sprint)
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
 * 
 * ## Permission levels
 * 
 * Permission for Iterations within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.
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
 * import com.pulumi.azuredevops.IterativePermissions;
 * import com.pulumi.azuredevops.IterativePermissionsArgs;
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
 *         var example_root_permissions = new IterativePermissions(&#34;example-root-permissions&#34;, IterativePermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;CREATE_CHILDREN&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;GENERIC_READ&#34;, &#34;NotSet&#34;),
 *                 Map.entry(&#34;DELETE&#34;, &#34;Deny&#34;)
 *             ))
 *             .build());
 * 
 *         var example_iteration_permissions = new IterativePermissions(&#34;example-iteration-permissions&#34;, IterativePermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .path(&#34;Iteration 1&#34;)
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;CREATE_CHILDREN&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;GENERIC_READ&#34;, &#34;NotSet&#34;),
 *                 Map.entry(&#34;DELETE&#34;, &#34;Allow&#34;)
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
@ResourceType(type="azuredevops:index/iterativePermissions:IterativePermissions")
public class IterativePermissions extends com.pulumi.resources.CustomResource {
    /**
     * The name of the branch to assign the permissions.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
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
     */
    @Export(name="permissions", type=Map.class, parameters={String.class, String.class})
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
     * | Permission      | Description                    |
     * |-----------------|--------------------------------|
     * | GENERIC_READ    | View permissions for this node |
     * | GENERIC_WRITE   | Edit this node                 |
     * | CREATE_CHILDREN | Create child nodes             |
     * | DELETE          | Delete this node               |
     * 
     */
    @Export(name="replace", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     * | Permission      | Description                    |
     * |-----------------|--------------------------------|
     * | GENERIC_READ    | View permissions for this node |
     * | GENERIC_WRITE   | Edit this node                 |
     * | CREATE_CHILDREN | Create child nodes             |
     * | DELETE          | Delete this node               |
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IterativePermissions(String name) {
        this(name, IterativePermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IterativePermissions(String name, IterativePermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IterativePermissions(String name, IterativePermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/iterativePermissions:IterativePermissions", name, args == null ? IterativePermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IterativePermissions(String name, Output<String> id, @Nullable IterativePermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/iterativePermissions:IterativePermissions", name, state, makeResourceOptions(options, id));
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
    public static IterativePermissions get(String name, Output<String> id, @Nullable IterativePermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IterativePermissions(name, id, state, options);
    }
}
