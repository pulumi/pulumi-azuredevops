// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ElasticPoolArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ElasticPoolState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Elastic pool within Azure DevOps.
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
 * import com.pulumi.azuredevops.ServiceEndpointAzureRM;
 * import com.pulumi.azuredevops.ServiceEndpointAzureRMArgs;
 * import com.pulumi.azuredevops.inputs.ServiceEndpointAzureRMCredentialsArgs;
 * import com.pulumi.azuredevops.ElasticPool;
 * import com.pulumi.azuredevops.ElasticPoolArgs;
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
 *             .visibility(&#34;private&#34;)
 *             .versionControl(&#34;Git&#34;)
 *             .workItemTemplate(&#34;Agile&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .build());
 * 
 *         var exampleServiceEndpointAzureRM = new ServiceEndpointAzureRM(&#34;exampleServiceEndpointAzureRM&#34;, ServiceEndpointAzureRMArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .serviceEndpointName(&#34;Example Azure Connection&#34;)
 *             .description(&#34;Managed by Terraform&#34;)
 *             .serviceEndpointAuthenticationScheme(&#34;ServicePrincipal&#34;)
 *             .credentials(ServiceEndpointAzureRMCredentialsArgs.builder()
 *                 .serviceprincipalid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .serviceprincipalkey(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *                 .build())
 *             .azurermSpnTenantid(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .azurermSubscriptionName(&#34;Subscription Name&#34;)
 *             .build());
 * 
 *         var exampleElasticPool = new ElasticPool(&#34;exampleElasticPool&#34;, ElasticPoolArgs.builder()        
 *             .serviceEndpointId(exampleServiceEndpointAzureRM.id())
 *             .serviceEndpointScope(exampleProject.id())
 *             .desiredIdle(2)
 *             .maxCapacity(3)
 *             .azureResourceId(&#34;/subscriptions/&lt;Subscription Id&gt;/resourceGroups/&lt;Resource Name&gt;/providers/Microsoft.Compute/virtualMachineScaleSets/&lt;VMSS Name&gt;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 7.0 - Elastic Pools](https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/elasticpools/create?view=azure-devops-rest-7.0)
 * 
 * ## Import
 * 
 * Azure DevOps Agent Pools can be imported using the Elastic pool ID, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/elasticPool:ElasticPool example 0
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/elasticPool:ElasticPool")
public class ElasticPool extends com.pulumi.resources.CustomResource {
    /**
     * Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     * 
     */
    @Export(name="agentInteractiveUi", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> agentInteractiveUi;

    /**
     * @return Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> agentInteractiveUi() {
        return Codegen.optional(this.agentInteractiveUi);
    }
    /**
     * Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     * 
     */
    @Export(name="autoProvision", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoProvision;

    /**
     * @return Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> autoProvision() {
        return Codegen.optional(this.autoProvision);
    }
    /**
     * Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     * 
     */
    @Export(name="autoUpdate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoUpdate;

    /**
     * @return Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> autoUpdate() {
        return Codegen.optional(this.autoUpdate);
    }
    /**
     * The ID of the Azure resource.
     * 
     */
    @Export(name="azureResourceId", refs={String.class}, tree="[0]")
    private Output<String> azureResourceId;

    /**
     * @return The ID of the Azure resource.
     * 
     */
    public Output<String> azureResourceId() {
        return this.azureResourceId;
    }
    /**
     * Number of agents to keep on standby.
     * 
     */
    @Export(name="desiredIdle", refs={Integer.class}, tree="[0]")
    private Output<Integer> desiredIdle;

    /**
     * @return Number of agents to keep on standby.
     * 
     */
    public Output<Integer> desiredIdle() {
        return this.desiredIdle;
    }
    /**
     * Maximum number of virtual machines in the scale set.
     * 
     */
    @Export(name="maxCapacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxCapacity;

    /**
     * @return Maximum number of virtual machines in the scale set.
     * 
     */
    public Output<Integer> maxCapacity() {
        return this.maxCapacity;
    }
    /**
     * The name of the Elastic pool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Elastic pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Tear down virtual machines after every use. Defaults to `false`.
     * 
     */
    @Export(name="recycleAfterEachUse", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> recycleAfterEachUse;

    /**
     * @return Tear down virtual machines after every use. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> recycleAfterEachUse() {
        return Codegen.optional(this.recycleAfterEachUse);
    }
    /**
     * The ID of Service Endpoint used to connect to Azure.
     * 
     */
    @Export(name="serviceEndpointId", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointId;

    /**
     * @return The ID of Service Endpoint used to connect to Azure.
     * 
     */
    public Output<String> serviceEndpointId() {
        return this.serviceEndpointId;
    }
    /**
     * The Project ID of Service Endpoint belongs to.
     * 
     */
    @Export(name="serviceEndpointScope", refs={String.class}, tree="[0]")
    private Output<String> serviceEndpointScope;

    /**
     * @return The Project ID of Service Endpoint belongs to.
     * 
     */
    public Output<String> serviceEndpointScope() {
        return this.serviceEndpointScope;
    }
    /**
     * Delay in minutes before deleting excess idle agents. Defaults to `30`.
     * 
     */
    @Export(name="timeToLiveMinutes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeToLiveMinutes;

    /**
     * @return Delay in minutes before deleting excess idle agents. Defaults to `30`.
     * 
     */
    public Output<Optional<Integer>> timeToLiveMinutes() {
        return Codegen.optional(this.timeToLiveMinutes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ElasticPool(String name) {
        this(name, ElasticPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ElasticPool(String name, ElasticPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ElasticPool(String name, ElasticPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/elasticPool:ElasticPool", name, args == null ? ElasticPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ElasticPool(String name, Output<String> id, @Nullable ElasticPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/elasticPool:ElasticPool", name, state, makeResourceOptions(options, id));
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
    public static ElasticPool get(String name, Output<String> id, @Nullable ElasticPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ElasticPool(name, id, state, options);
    }
}
