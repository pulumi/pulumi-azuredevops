// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.LibraryPermissionsArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.LibraryPermissionsState;
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
 * Manages permissions for a Library
 * 
 * ## Example Usage
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
 * import com.pulumi.azuredevops.LibraryPermissions;
 * import com.pulumi.azuredevops.LibraryPermissionsArgs;
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
 *         var project = new Project(&#34;project&#34;, ProjectArgs.builder()        
 *             .name(&#34;Testing&#34;)
 *             .description(&#34;Testing-description&#34;)
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *         final var tf-project-readers = AzuredevopsFunctions.getGroup(GetGroupArgs.builder()
 *             .projectId(project.id())
 *             .name(&#34;Readers&#34;)
 *             .build());
 * 
 *         var permissions = new LibraryPermissions(&#34;permissions&#34;, LibraryPermissionsArgs.builder()        
 *             .projectId(project.id())
 *             .principal(tf_project_readers.applyValue(tf_project_readers -&gt; tf_project_readers.id()))
 *             .permissions(Map.ofEntries(
 *                 Map.entry(&#34;View&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;Administer&#34;, &#34;allow&#34;),
 *                 Map.entry(&#34;Use&#34;, &#34;allow&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Roles
 * 
 * The Azure DevOps UI uses roles to assign permissions for the Library.
 * 
 * | Role          | Allowed Permissions    |
 * | ------------- | ---------------------- |
 * | Reader        | View                   |
 * | Creator       | View, Create           |
 * | User          | View, Use              |
 * | Administrator | View, Use, Administer  |
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
@ResourceType(type="azuredevops:index/libraryPermissions:LibraryPermissions")
public class LibraryPermissions extends com.pulumi.resources.CustomResource {
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
     * 
     * | Permission        | Description                         |
     * | ----------------- | ----------------------------------- |
     * | View              | View library item                   |
     * | Administer        | Administer library item             |
     * | Create            | Create library item                 |
     * | ViewSecrets       | View library item secrets           |
     * | Use               | Use library item                    |
     * | Owner             | Owner library item                  |
     * 
     */
    @Export(name="replace", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> replace;

    /**
     * @return Replace (`true`) or merge (`false`) the permissions. Default: `true`
     * 
     * | Permission        | Description                         |
     * | ----------------- | ----------------------------------- |
     * | View              | View library item                   |
     * | Administer        | Administer library item             |
     * | Create            | Create library item                 |
     * | ViewSecrets       | View library item secrets           |
     * | Use               | Use library item                    |
     * | Owner             | Owner library item                  |
     * 
     */
    public Output<Optional<Boolean>> replace() {
        return Codegen.optional(this.replace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LibraryPermissions(String name) {
        this(name, LibraryPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LibraryPermissions(String name, LibraryPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LibraryPermissions(String name, LibraryPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/libraryPermissions:LibraryPermissions", name, args == null ? LibraryPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LibraryPermissions(String name, Output<String> id, @Nullable LibraryPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/libraryPermissions:LibraryPermissions", name, state, makeResourceOptions(options, id));
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
    public static LibraryPermissions get(String name, Output<String> id, @Nullable LibraryPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LibraryPermissions(name, id, state, options);
    }
}
