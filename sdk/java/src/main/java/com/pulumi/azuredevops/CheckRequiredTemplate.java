// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.CheckRequiredTemplateArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.CheckRequiredTemplateState;
import com.pulumi.azuredevops.outputs.CheckRequiredTemplateRequiredTemplate;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Manages a Required Template Check.
 * 
 * ## Example Usage
 * 
 * ### Protect a service connection
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
 * import com.pulumi.azuredevops.ServiceEndpointGeneric;
 * import com.pulumi.azuredevops.ServiceEndpointGenericArgs;
 * import com.pulumi.azuredevops.CheckRequiredTemplate;
 * import com.pulumi.azuredevops.CheckRequiredTemplateArgs;
 * import com.pulumi.azuredevops.inputs.CheckRequiredTemplateRequiredTemplateArgs;
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
 *             .build());
 * 
 *         var exampleServiceEndpointGeneric = new ServiceEndpointGeneric(&#34;exampleServiceEndpointGeneric&#34;, ServiceEndpointGenericArgs.builder()        
 *             .projectId(example.id())
 *             .serverUrl(&#34;https://some-server.example.com&#34;)
 *             .username(&#34;username&#34;)
 *             .password(&#34;password&#34;)
 *             .serviceEndpointName(&#34;Example Generic&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleCheckRequiredTemplate = new CheckRequiredTemplate(&#34;exampleCheckRequiredTemplate&#34;, CheckRequiredTemplateArgs.builder()        
 *             .projectId(example.id())
 *             .targetResourceId(exampleServiceEndpointGeneric.id())
 *             .targetResourceType(&#34;endpoint&#34;)
 *             .requiredTemplates(CheckRequiredTemplateRequiredTemplateArgs.builder()
 *                 .repositoryType(&#34;azuregit&#34;)
 *                 .repositoryName(&#34;project/repository&#34;)
 *                 .repositoryRef(&#34;refs/heads/main&#34;)
 *                 .templatePath(&#34;template/path.yml&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Protect an environment
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
 * import com.pulumi.azuredevops.Environment;
 * import com.pulumi.azuredevops.EnvironmentArgs;
 * import com.pulumi.azuredevops.CheckRequiredTemplate;
 * import com.pulumi.azuredevops.CheckRequiredTemplateArgs;
 * import com.pulumi.azuredevops.inputs.CheckRequiredTemplateRequiredTemplateArgs;
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
 *             .build());
 * 
 *         var exampleEnvironment = new Environment(&#34;exampleEnvironment&#34;, EnvironmentArgs.builder()        
 *             .projectId(example.id())
 *             .name(&#34;Example Environment&#34;)
 *             .build());
 * 
 *         var exampleCheckRequiredTemplate = new CheckRequiredTemplate(&#34;exampleCheckRequiredTemplate&#34;, CheckRequiredTemplateArgs.builder()        
 *             .projectId(example.id())
 *             .targetResourceId(exampleEnvironment.id())
 *             .targetResourceType(&#34;environment&#34;)
 *             .requiredTemplates(            
 *                 CheckRequiredTemplateRequiredTemplateArgs.builder()
 *                     .repositoryName(&#34;project/repository&#34;)
 *                     .repositoryRef(&#34;refs/heads/main&#34;)
 *                     .templatePath(&#34;template/path.yml&#34;)
 *                     .build(),
 *                 CheckRequiredTemplateRequiredTemplateArgs.builder()
 *                     .repositoryName(&#34;project/repository&#34;)
 *                     .repositoryRef(&#34;refs/heads/main&#34;)
 *                     .templatePath(&#34;template/alternate-path.yml&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Importing this resource is not supported.
 * 
 */
@ResourceType(type="azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate")
public class CheckRequiredTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The project ID. Changing this forces a new Required Template Check to be created.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project ID. Changing this forces a new Required Template Check to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * One or more `required_template` blocks documented below.
     * 
     */
    @Export(name="requiredTemplates", refs={List.class,CheckRequiredTemplateRequiredTemplate.class}, tree="[0,1]")
    private Output<List<CheckRequiredTemplateRequiredTemplate>> requiredTemplates;

    /**
     * @return One or more `required_template` blocks documented below.
     * 
     */
    public Output<List<CheckRequiredTemplateRequiredTemplate>> requiredTemplates() {
        return this.requiredTemplates;
    }
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
     * 
     */
    @Export(name="targetResourceId", refs={String.class}, tree="[0]")
    private Output<String> targetResourceId;

    /**
     * @return The ID of the resource being protected by the check. Changing this forces a new Required Template Check to be created.
     * 
     */
    public Output<String> targetResourceId() {
        return this.targetResourceId;
    }
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
     * 
     */
    @Export(name="targetResourceType", refs={String.class}, tree="[0]")
    private Output<String> targetResourceType;

    /**
     * @return The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Required Template Check to be created.
     * 
     */
    public Output<String> targetResourceType() {
        return this.targetResourceType;
    }
    /**
     * The version of the check.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    /**
     * @return The version of the check.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CheckRequiredTemplate(String name) {
        this(name, CheckRequiredTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CheckRequiredTemplate(String name, CheckRequiredTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CheckRequiredTemplate(String name, CheckRequiredTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate", name, args == null ? CheckRequiredTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CheckRequiredTemplate(String name, Output<String> id, @Nullable CheckRequiredTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/checkRequiredTemplate:CheckRequiredTemplate", name, state, makeResourceOptions(options, id));
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
    public static CheckRequiredTemplate get(String name, Output<String> id, @Nullable CheckRequiredTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CheckRequiredTemplate(name, id, state, options);
    }
}
