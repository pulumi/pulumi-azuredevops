// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.WikiPageArgs;
import com.pulumi.azuredevops.inputs.WikiPageState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages Wiki pages within Azure DevOps project.
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
 * import com.pulumi.azuredevops.Wiki;
 * import com.pulumi.azuredevops.WikiArgs;
 * import com.pulumi.azuredevops.WikiPage;
 * import com.pulumi.azuredevops.WikiPageArgs;
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
 *             .description("Managed by Terraform")
 *             .build());
 * 
 *         var exampleWiki = new Wiki("exampleWiki", WikiArgs.builder()
 *             .name("Example project wiki ")
 *             .projectId(example.id())
 *             .type("projectWiki")
 *             .build());
 * 
 *         var exampleWikiPage = new WikiPage("exampleWikiPage", WikiPageArgs.builder()
 *             .projectId(example.id())
 *             .wikiId(exampleWiki.id())
 *             .path("/page")
 *             .content("content")
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
 * - [Azure DevOps Service REST API 7.1 - Wiki Page](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/pages?view=azure-devops-rest-7.1)
 * 
 */
@ResourceType(type="azuredevops:index/wikiPage:WikiPage")
public class WikiPage extends com.pulumi.resources.CustomResource {
    /**
     * The content of the wiki page.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return The content of the wiki page.
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The path of the wiki page.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path of the wiki page.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * The ID of the Project.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the Project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The ID of the Wiki.
     * 
     */
    @Export(name="wikiId", refs={String.class}, tree="[0]")
    private Output<String> wikiId;

    /**
     * @return The ID of the Wiki.
     * 
     */
    public Output<String> wikiId() {
        return this.wikiId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WikiPage(java.lang.String name) {
        this(name, WikiPageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WikiPage(java.lang.String name, WikiPageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WikiPage(java.lang.String name, WikiPageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/wikiPage:WikiPage", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private WikiPage(java.lang.String name, Output<java.lang.String> id, @Nullable WikiPageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/wikiPage:WikiPage", name, state, makeResourceOptions(options, id), false);
    }

    private static WikiPageArgs makeArgs(WikiPageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? WikiPageArgs.Empty : args;
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
    public static WikiPage get(java.lang.String name, Output<java.lang.String> id, @Nullable WikiPageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WikiPage(name, id, state, options);
    }
}