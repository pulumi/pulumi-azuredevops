// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Agent;

import com.pulumi.azuredevops.Agent.QueueArgs;
import com.pulumi.azuredevops.Agent.inputs.QueueState;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
 * Organization defined pool to a project.
 * 
 * The created queue is not authorized for use by all pipelines in the project. However,
 * the `azuredevops.ResourceAuthorization` resource can be used to grant authorization.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azuredevops.AzuredevopsFunctions;
 * import com.pulumi.azuredevops.Agent.inputs.GetPoolArgs;
 * import com.pulumi.azuredevops.Queue;
 * import com.pulumi.azuredevops.QueueArgs;
 * import com.pulumi.azuredevops.ResourceAuthorization;
 * import com.pulumi.azuredevops.ResourceAuthorizationArgs;
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
 *         var exampleProject = new Project(&#34;exampleProject&#34;);
 * 
 *         final var examplePool = AzuredevopsFunctions.getPool(GetPoolArgs.builder()
 *             .name(&#34;example-pool&#34;)
 *             .build());
 * 
 *         var exampleQueue = new Queue(&#34;exampleQueue&#34;, QueueArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .agentPoolId(examplePool.applyValue(getPoolResult -&gt; getPoolResult.id()))
 *             .build());
 * 
 *         var exampleResourceAuthorization = new ResourceAuthorization(&#34;exampleResourceAuthorization&#34;, ResourceAuthorizationArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .resourceId(exampleQueue.id())
 *             .type(&#34;queue&#34;)
 *             .authorized(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Relevant Links
 * 
 * - [Azure DevOps Service REST API 6.0 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-6.0)
 * 
 * ## Import
 * 
 * Azure DevOps Agent Pools can be imported using the project ID and agent queue ID, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:Agent/queue:Queue example 00000000-0000-0000-0000-000000000000/0
 * ```
 * 
 * @deprecated
 * azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue
 * 
 */
@Deprecated /* azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue */
@ResourceType(type="azuredevops:Agent/queue:Queue")
public class Queue extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the organization agent pool.
     * 
     */
    @Export(name="agentPoolId", type=Integer.class, parameters={})
    private Output<Integer> agentPoolId;

    /**
     * @return The ID of the organization agent pool.
     * 
     */
    public Output<Integer> agentPoolId() {
        return this.agentPoolId;
    }
    /**
     * The ID of the project in which to create the resource.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The ID of the project in which to create the resource.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Queue(String name) {
        this(name, QueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Queue(String name, QueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Queue(String name, QueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Agent/queue:Queue", name, args == null ? QueueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Queue(String name, Output<String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:Agent/queue:Queue", name, state, makeResourceOptions(options, id));
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
    public static Queue get(String name, Output<String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Queue(name, id, state, options);
    }
}
