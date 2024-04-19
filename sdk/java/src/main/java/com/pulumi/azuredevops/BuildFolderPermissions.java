// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.BuildFolderPermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.BuildFolderPermissionsState;
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
 * Manages permissions for a Build Folder
 * 
 * &gt; **Note** Permissions can be assigned to group principals and not to single user principals.
 * 
 * ## Example Usage
 * 
 * ### Set specific folder permissions
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
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
 * import com.pulumi.azuredevops.BuildFolder;
 * import com.pulumi.azuredevops.BuildFolderArgs;
 * import com.pulumi.azuredevops.BuildFolderPermissions;
 * import com.pulumi.azuredevops.BuildFolderPermissionsArgs;
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
 *             .name(&#34;Example Project&#34;)
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
 *         var exampleBuildFolder = new BuildFolder(&#34;exampleBuildFolder&#34;, BuildFolderArgs.builder()        
 *             .projectId(example.id())
 *             .path(&#34;\\ExampleFolder&#34;)
 *             .description(&#34;ExampleFolder description&#34;)
 *             .build());
 * 
 *         var exampleBuildFolderPermissions = new BuildFolderPermissions(&#34;exampleBuildFolderPermissions&#34;, BuildFolderPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .path(&#34;\\ExampleFolder&#34;)
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;ViewBuilds&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;EditBuildQuality&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;RetainIndefinitely&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;DeleteBuilds&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;ManageBuildQualities&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;DestroyBuilds&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;UpdateBuildInformation&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;QueueBuilds&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;ManageBuildQueue&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;StopBuilds&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;ViewBuildDefinition&#34;, &#34;Allow&#34;),
 *                 Map.entry(&#34;EditBuildDefinition&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;DeleteBuildDefinition&#34;, &#34;Deny&#34;),
 *                 Map.entry(&#34;AdministerBuildPermissions&#34;, &#34;NotSet&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Set root folder permissions
 * &lt;!--Start PulumiCodeChooser --&gt;
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
 * import com.pulumi.azuredevops.BuildFolderPermissions;
 * import com.pulumi.azuredevops.BuildFolderPermissionsArgs;
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
 *             .name(&#34;Example Project&#34;)
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
 *         var exampleBuildFolderPermissions = new BuildFolderPermissions(&#34;exampleBuildFolderPermissions&#34;, BuildFolderPermissionsArgs.builder()        
 *             .projectId(example.id())
 *             .path(&#34;\\&#34;)
 *             .principal(example_readers.applyValue(example_readers -&gt; example_readers.id()))
 *             .permissions(Map.of(&#34;RetainIndefinitely&#34;, &#34;Allow&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
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
@ResourceType(type="azuredevops:index/buildFolderPermissions:BuildFolderPermissions")
public class BuildFolderPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The folder path to assign the permissions.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The folder path to assign the permissions.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * the permissions to assign. The following permissions are available.
     * 
     * | Permission                     | Description                           |
     * |--------------------------------|---------------------------------------|
     * | ViewBuilds                     | View builds                           |
     * | EditBuildQuality               | Edit build quality                    |
     * | RetainIndefinitely             | Retain indefinitely                   |
     * | DeleteBuilds                   | Delete builds                         |
     * | ManageBuildQualities           | Manage build qualities                |
     * | DestroyBuilds                  | Destroy builds                        |
     * | UpdateBuildInformation         | Update build information              |
     * | QueueBuilds                    | Queue builds                          |
     * | ManageBuildQueue               | Manage build queue                    |
     * | StopBuilds                     | Stop builds                           |
     * | ViewBuildDefinition            | View build pipeline                   |
     * | EditBuildDefinition            | Edit build pipeline                   |
     * | DeleteBuildDefinition          | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation | Override check-in validation by build |
     * | AdministerBuildPermissions     | Administer build permissions          |
     * 
     */
    @Export(name="permissions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> permissions;

    /**
     * @return the permissions to assign. The following permissions are available.
     * 
     * | Permission                     | Description                           |
     * |--------------------------------|---------------------------------------|
     * | ViewBuilds                     | View builds                           |
     * | EditBuildQuality               | Edit build quality                    |
     * | RetainIndefinitely             | Retain indefinitely                   |
     * | DeleteBuilds                   | Delete builds                         |
     * | ManageBuildQualities           | Manage build qualities                |
     * | DestroyBuilds                  | Destroy builds                        |
     * | UpdateBuildInformation         | Update build information              |
     * | QueueBuilds                    | Queue builds                          |
     * | ManageBuildQueue               | Manage build queue                    |
     * | StopBuilds                     | Stop builds                           |
     * | ViewBuildDefinition            | View build pipeline                   |
     * | EditBuildDefinition            | Edit build pipeline                   |
     * | DeleteBuildDefinition          | Delete build pipeline                 |
     * | OverrideBuildCheckInValidation | Override check-in validation by build |
     * | AdministerBuildPermissions     | Administer build permissions          |
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
    public BuildFolderPermissions(String name) {
        this(name, BuildFolderPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BuildFolderPermissions(String name, BuildFolderPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BuildFolderPermissions(String name, BuildFolderPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildFolderPermissions:BuildFolderPermissions", name, args == null ? BuildFolderPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BuildFolderPermissions(String name, Output<String> id, @Nullable BuildFolderPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildFolderPermissions:BuildFolderPermissions", name, state, makeResourceOptions(options, id));
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
    public static BuildFolderPermissions get(String name, Output<String> id, @Nullable BuildFolderPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BuildFolderPermissions(name, id, state, options);
    }
}
