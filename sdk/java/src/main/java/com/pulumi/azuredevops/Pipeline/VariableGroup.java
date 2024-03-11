// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Pipeline;

import com.pulumi.azuredevops.Pipeline.VariableGroupArgs;
import com.pulumi.azuredevops.Pipeline.inputs.VariableGroupState;
import com.pulumi.azuredevops.Pipeline.outputs.VariableGroupKeyVault;
import com.pulumi.azuredevops.Pipeline.outputs.VariableGroupVariable;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.azuredevops.VariableGroup;
 * import com.pulumi.azuredevops.VariableGroupArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupVariableArgs;
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleVariableGroup = new VariableGroup(&#34;exampleVariableGroup&#34;, VariableGroupArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .description(&#34;Example Variable Group Description&#34;)
 *             .allowAccess(true)
 *             .variables(            
 *                 VariableGroupVariableArgs.builder()
 *                     .name(&#34;key1&#34;)
 *                     .value(&#34;val1&#34;)
 *                     .build(),
 *                 VariableGroupVariableArgs.builder()
 *                     .name(&#34;key2&#34;)
 *                     .secretValue(&#34;val2&#34;)
 *                     .isSecret(true)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### With AzureRM Key Vault
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
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMCredentialsArgs;
 * import com.pulumi.azuredevops.VariableGroup;
 * import com.pulumi.azuredevops.VariableGroupArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupKeyVaultArgs;
 * import com.pulumi.azuredevops.inputs.VariableGroupVariableArgs;
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .visibility(&#34;private&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example AzureRM&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .serviceprincipalkey(&#34;xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&#34;)
 *                 .build())
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Example Subscription Name&#34;)
 *             .build());
 * 
 *         var exampleVariableGroup = new VariableGroup(&#34;exampleVariableGroup&#34;, VariableGroupArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .description(&#34;Example Variable Group Description&#34;)
 *             .allowAccess(true)
 *             .keyVault(VariableGroupKeyVaultArgs.builder()
 *                 .name(&#34;example-kv&#34;)
 *                 .serviceEndpointId(exampleServiceEndpointAzureRM.id())
 *                 .build())
 *             .variables(            
 *                 VariableGroupVariableArgs.builder()
 *                     .name(&#34;key1&#34;)
 *                     .build(),
 *                 VariableGroupVariableArgs.builder()
 *                     .name(&#34;key2&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
 * - [Azure DevOps Service REST API 7.0 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-7.0)
 * 
 * ## PAT Permissions Required
 * 
 * - **Variable Groups**: Read, Create, &amp; Manage
 * - **Build**: Read &amp; execute
 * - **Project and Team**: Read
 * - **Token Administration**: Read &amp; manage
 * - **Tokens**: Read &amp; manage
 * - **Work Items**: Read
 * 
 * ## Import
 * 
 * **Variable groups containing secret values cannot be imported.**
 * 
 * Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.
 * 
 * ```sh
 * $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example &#34;Example Project/10&#34;
 * ```
 * 
 * or
 * 
 * ```sh
 * $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 * _Note that for secret variables, the import command retrieve blank value in the tfstate._
 * 
 * @deprecated
 * azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup
 * 
 */
@Deprecated /* azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup */
@ResourceType(type="azuredevops:Pipeline/variableGroup:VariableGroup")
public class VariableGroup extends com.pulumi.resources.CustomResource {
    /**
     * Boolean that indicate if this variable group is shared by all pipelines of this project.
     * 
     */
    @Export(name="allowAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowAccess;

    /**
     * @return Boolean that indicate if this variable group is shared by all pipelines of this project.
     * 
     */
    public Output<Optional<Boolean>> allowAccess() {
        return Codegen.optional(this.allowAccess);
    }
    /**
     * The description of the Variable Group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Variable Group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A list of `key_vault` blocks as documented below.
     * 
     */
    @Export(name="keyVault", refs={VariableGroupKeyVault.class}, tree="[0]")
    private Output</* @Nullable */ VariableGroupKeyVault> keyVault;

    /**
     * @return A list of `key_vault` blocks as documented below.
     * 
     */
    public Output<Optional<VariableGroupKeyVault>> keyVault() {
        return Codegen.optional(this.keyVault);
    }
    /**
     * The name of the Variable Group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Variable Group.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * One or more `variable` blocks as documented below.
     * 
     */
    @Export(name="variables", refs={List.class,VariableGroupVariable.class}, tree="[0,1]")
    private Output<List<VariableGroupVariable>> variables;

    /**
     * @return One or more `variable` blocks as documented below.
     * 
     */
    public Output<List<VariableGroupVariable>> variables() {
        return this.variables;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VariableGroup(String name) {
        this(name, VariableGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VariableGroup(String name, VariableGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VariableGroup(String name, VariableGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Pipeline/variableGroup:VariableGroup", name, args == null ? VariableGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VariableGroup(String name, Output<String> id, @Nullable VariableGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Pipeline/variableGroup:VariableGroup", name, state, makeResourceOptions(options, id));
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
    public static VariableGroup get(String name, Output<String> id, @Nullable VariableGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VariableGroup(name, id, state, options);
    }
}
