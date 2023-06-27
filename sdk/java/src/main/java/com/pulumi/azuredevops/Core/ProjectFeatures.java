// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Core;

import com.pulumi.azuredevops.Core.ProjectFeaturesArgs;
import com.pulumi.azuredevops.Core.inputs.ProjectFeaturesState;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * Manages features for Azure DevOps projects
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
 * import com.pulumi.azuredevops.ProjectFeatures;
 * import com.pulumi.azuredevops.ProjectFeaturesArgs;
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
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var example_features = new ProjectFeatures(&#34;example-features&#34;, ProjectFeaturesArgs.builder()        
 *             .projectId(example.id())
 *             .features(Map.ofEntries(
 *                 Map.entry(&#34;testplans&#34;, &#34;disabled&#34;),
 *                 Map.entry(&#34;artifacts&#34;, &#34;enabled&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * No official documentation available
 * 
 * ## PAT Permissions Required
 * 
 * - **Project &amp; Team**: Read, Write, &amp; Manage
 * 
 * ## Import
 * 
 * Azure DevOps feature settings can be imported using the project id, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:Core/projectFeatures:ProjectFeatures example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 * @deprecated
 * azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures
 * 
 */
@Deprecated /* azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures */
@ResourceType(type="azuredevops:Core/projectFeatures:ProjectFeatures")
public class ProjectFeatures extends com.pulumi.resources.CustomResource {
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.\
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     * 
     * &gt; **NOTE:**
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    @Export(name="features", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> features;

    /**
     * @return Defines the status (`enabled`, `disabled`) of the project features.\
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     * 
     * &gt; **NOTE:**
     * It&#39;s possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * via the `features` block by using the `azuredevops.Project` resource.
     * However it&#39;s not possible to use both methods to manage features, since there&#39;ll be conflicts.
     * 
     */
    public Output<Map<String,String>> features() {
        return this.features;
    }
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectFeatures(String name) {
        this(name, ProjectFeaturesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectFeatures(String name, ProjectFeaturesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectFeatures(String name, ProjectFeaturesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Core/projectFeatures:ProjectFeatures", name, args == null ? ProjectFeaturesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectFeatures(String name, Output<String> id, @Nullable ProjectFeaturesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Core/projectFeatures:ProjectFeatures", name, state, makeResourceOptions(options, id));
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
    public static ProjectFeatures get(String name, Output<String> id, @Nullable ProjectFeaturesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectFeatures(name, id, state, options);
    }
}
