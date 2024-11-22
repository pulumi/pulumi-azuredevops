// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.FeedArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.FeedState;
import com.pulumi.azuredevops.outputs.FeedFeature;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages creation of the Feed within Azure DevOps organization.
 * 
 * ## Example Usage
 * 
 * ### Create Feed in the scope of whole Organization
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Feed;
 * import com.pulumi.azuredevops.FeedArgs;
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
 *         var example = new Feed("example", FeedArgs.builder()
 *             .name("releases")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Create Feed in the scope of a Project
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
 * import com.pulumi.azuredevops.Feed;
 * import com.pulumi.azuredevops.FeedArgs;
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
 *             .description("Managed by Pulumi")
 *             .build());
 * 
 *         var exampleFeed = new Feed("exampleFeed", FeedArgs.builder()
 *             .name("releases")
 *             .projectId(example.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Create Feed with Soft Delete
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Feed;
 * import com.pulumi.azuredevops.FeedArgs;
 * import com.pulumi.azuredevops.inputs.FeedFeatureArgs;
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
 *         var example = new Feed("example", FeedArgs.builder()
 *             .name("releases")
 *             .features(FeedFeatureArgs.builder()
 *                 .permanentDelete(false)
 *                 .build())
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
 * - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Feed can be imported using the Project ID and Feed ID or Feed ID e.g.:
 * 
 * ```sh
 * $ pulumi import azuredevops:index/feed:Feed example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 * 
 * or
 * 
 * ```sh
 * $ pulumi import azuredevops:index/feed:Feed example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/feed:Feed")
public class Feed extends com.pulumi.resources.CustomResource {
    /**
     * A `features` blocks as documented below.
     * 
     * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
     * 
     */
    @Export(name="features", refs={List.class,FeedFeature.class}, tree="[0,1]")
    private Output</* @Nullable */ List<FeedFeature>> features;

    /**
     * @return A `features` blocks as documented below.
     * 
     * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
     * 
     */
    public Output<Optional<List<FeedFeature>>> features() {
        return Codegen.optional(this.features);
    }
    /**
     * The name of the Feed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Feed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    /**
     * @return The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Feed(java.lang.String name) {
        this(name, FeedArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Feed(java.lang.String name, @Nullable FeedArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Feed(java.lang.String name, @Nullable FeedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/feed:Feed", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Feed(java.lang.String name, Output<java.lang.String> id, @Nullable FeedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/feed:Feed", name, state, makeResourceOptions(options, id), false);
    }

    private static FeedArgs makeArgs(@Nullable FeedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FeedArgs.Empty : args;
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
    public static Feed get(java.lang.String name, Output<java.lang.String> id, @Nullable FeedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Feed(name, id, state, options);
    }
}
