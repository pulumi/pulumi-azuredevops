// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ProjectArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ProjectState;
import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a project within Azure DevOps.
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
 *             .description(&#34;Managed by Terraform&#34;)
 *             .features(Map.ofEntries(
 *                 Map.entry(&#34;artifacts&#34;, &#34;disabled&#34;),
 *                 Map.entry(&#34;testplans&#34;, &#34;disabled&#34;)
 *             ))
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Projects](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects?view=azure-devops-rest-6.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Project &amp; Team**: Read, Write, &amp; Manage
 * 
 * ## Import
 * 
 * Azure DevOps Projects can be imported using the project name or by the project Guid, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/project:Project example &#34;Example Project&#34;
 * ```
 * 
 *  or
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/project:Project example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/project:Project")
public class Project extends com.pulumi.resources.CustomResource {
    /**
     * The Description of the Project.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The Description of the Project.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.
     * Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     * 
     */
    @Export(name="features", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> features;

    /**
     * @return Defines the status (`enabled`, `disabled`) of the project features.
     * Valid features are `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     * 
     */
    public Output<Optional<Map<String,String>>> features() {
        return Codegen.optional(this.features);
    }
    /**
     * The Project Name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The Project Name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Process Template ID used by the Project.
     * 
     */
    @Export(name="processTemplateId", type=String.class, parameters={})
    private Output<String> processTemplateId;

    /**
     * @return The Process Template ID used by the Project.
     * 
     */
    public Output<String> processTemplateId() {
        return this.processTemplateId;
    }
    /**
     * Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     * 
     */
    @Export(name="versionControl", type=String.class, parameters={})
    private Output</* @Nullable */ String> versionControl;

    /**
     * @return Specifies the version control system. Valid values: `Git` or `Tfvc`. Defaults to `Git`.
     * 
     */
    public Output<Optional<String>> versionControl() {
        return Codegen.optional(this.versionControl);
    }
    /**
     * Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     * 
     */
    @Export(name="visibility", type=String.class, parameters={})
    private Output</* @Nullable */ String> visibility;

    /**
     * @return Specifies the visibility of the Project. Valid values: `private` or `public`. Defaults to `private`.
     * 
     */
    public Output<Optional<String>> visibility() {
        return Codegen.optional(this.visibility);
    }
    /**
     * Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     * 
     */
    @Export(name="workItemTemplate", type=String.class, parameters={})
    private Output</* @Nullable */ String> workItemTemplate;

    /**
     * @return Specifies the work item template. Valid values: `Agile`, `Basic`, `CMMI`, `Scrum` or a custom, pre-existing one. Defaults to `Agile`. An empty string will use the parent organization default.
     * 
     */
    public Output<Optional<String>> workItemTemplate() {
        return Codegen.optional(this.workItemTemplate);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Project(String name) {
        this(name, ProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Project(String name, @Nullable ProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Project(String name, @Nullable ProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/project:Project", name, args == null ? ProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Project(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/project:Project", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("azuredevops:Core/project:Project").build())
            ))
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
    public static Project get(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Project(name, id, state, options);
    }
}
