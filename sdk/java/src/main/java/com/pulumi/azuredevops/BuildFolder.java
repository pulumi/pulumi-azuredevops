// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.BuildFolderArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.BuildFolderState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Build Folder.
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
 * import com.pulumi.azuredevops.BuildFolder;
 * import com.pulumi.azuredevops.BuildFolderArgs;
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
 *             .visibility("private")
 *             .versionControl("Git")
 *             .workItemTemplate("Agile")
 *             .build());
 * 
 *         var exampleBuildFolder = new BuildFolder("exampleBuildFolder", BuildFolderArgs.builder()
 *             .projectId(example.id())
 *             .path("\\ExampleFolder")
 *             .description("ExampleFolder description")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Build Folders can be imported using the `project name/path` or `project id/path`, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:index/buildFolder:BuildFolder example &#34;Example Project/\\ExampleFolder&#34;
 * ```
 * 
 * or
 * 
 * ```sh
 * $ pulumi import azuredevops:index/buildFolder:BuildFolder example 00000000-0000-0000-0000-000000000000/\\ExampleFolder
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/buildFolder:BuildFolder")
public class BuildFolder extends com.pulumi.resources.CustomResource {
    /**
     * Folder Description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Folder Description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The folder path.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The folder path.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * The ID of the project in which the folder will be created.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project in which the folder will be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BuildFolder(java.lang.String name) {
        this(name, BuildFolderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BuildFolder(java.lang.String name, BuildFolderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BuildFolder(java.lang.String name, BuildFolderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildFolder:BuildFolder", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BuildFolder(java.lang.String name, Output<java.lang.String> id, @Nullable BuildFolderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/buildFolder:BuildFolder", name, state, makeResourceOptions(options, id), false);
    }

    private static BuildFolderArgs makeArgs(BuildFolderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BuildFolderArgs.Empty : args;
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
    public static BuildFolder get(java.lang.String name, Output<java.lang.String> id, @Nullable BuildFolderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BuildFolder(name, id, state, options);
    }
}
