// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ServicehookStorageQueuePipelinesArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.azuredevops.inputs.ServicehookStorageQueuePipelinesState;
import com.pulumi.azuredevops.outputs.ServicehookStorageQueuePipelinesRunStateChangedEvent;
import com.pulumi.azuredevops.outputs.ServicehookStorageQueuePipelinesStageStateChangedEvent;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Service Hook Storage Queue Pipelines.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.Project;
 * import com.pulumi.azurerm.azurerm_resource_group;
 * import com.pulumi.azurerm.Azurerm_resource_groupArgs;
 * import com.pulumi.azurerm.azurerm_storage_account;
 * import com.pulumi.azurerm.Azurerm_storage_accountArgs;
 * import com.pulumi.azurerm.azurerm_storage_queue;
 * import com.pulumi.azurerm.Azurerm_storage_queueArgs;
 * import com.pulumi.azuredevops.ServicehookStorageQueuePipelines;
 * import com.pulumi.azuredevops.ServicehookStorageQueuePipelinesArgs;
 * import com.pulumi.azuredevops.inputs.ServicehookStorageQueuePipelinesRunStateChangedEventArgs;
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
 *         var exampleazurerm_resource_group = new Azurerm_resource_group(&#34;exampleazurerm_resource_group&#34;, Azurerm_resource_groupArgs.builder()        
 *             .name(&#34;example-resources&#34;)
 *             .location(&#34;West Europe&#34;)
 *             .build());
 * 
 *         var exampleazurerm_storage_account = new Azurerm_storage_account(&#34;exampleazurerm_storage_account&#34;, Azurerm_storage_accountArgs.builder()        
 *             .name(&#34;servicehookexamplestacc&#34;)
 *             .resourceGroupName(exampleazurerm_resource_group.name())
 *             .location(exampleazurerm_resource_group.location())
 *             .accountTier(&#34;Standard&#34;)
 *             .accountReplicationType(&#34;LRS&#34;)
 *             .build());
 * 
 *         var exampleazurerm_storage_queue = new Azurerm_storage_queue(&#34;exampleazurerm_storage_queue&#34;, Azurerm_storage_queueArgs.builder()        
 *             .name(&#34;examplequeue&#34;)
 *             .storageAccountName(exampleazurerm_storage_account.name())
 *             .build());
 * 
 *         var exampleServicehookStorageQueuePipelines = new ServicehookStorageQueuePipelines(&#34;exampleServicehookStorageQueuePipelines&#34;, ServicehookStorageQueuePipelinesArgs.builder()        
 *             .projectId(exampleProject.id())
 *             .accountName(exampleazurerm_storage_account.name())
 *             .accountKey(exampleazurerm_storage_account.primaryAccessKey())
 *             .queueName(exampleazurerm_storage_queue.name())
 *             .visiTimeout(30)
 *             .runStateChangedEvent(ServicehookStorageQueuePipelinesRunStateChangedEventArgs.builder()
 *                 .runStateFilter(&#34;Completed&#34;)
 *                 .runResultFilter(&#34;Succeeded&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * An empty configuration block will occur in all events triggering the associated action.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuredevops.ServicehookStorageQueuePipelines;
 * import com.pulumi.azuredevops.ServicehookStorageQueuePipelinesArgs;
 * import com.pulumi.azuredevops.inputs.ServicehookStorageQueuePipelinesRunStateChangedEventArgs;
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
 *         var example = new ServicehookStorageQueuePipelines(&#34;example&#34;, ServicehookStorageQueuePipelinesArgs.builder()        
 *             .projectId(azuredevops_project.example().id())
 *             .accountName(azurerm_storage_account.example().name())
 *             .accountKey(azurerm_storage_account.example().primary_access_key())
 *             .queueName(azurerm_storage_queue.example().name())
 *             .visiTimeout(30)
 *             .runStateChangedEvent()
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Service Hook Storage Queue Pipeliness can be imported using the `resource id`, e.g.
 * 
 * ```sh
 *  $ pulumi import azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines")
public class ServicehookStorageQueuePipelines extends com.pulumi.resources.CustomResource {
    /**
     * A valid account key from the queue&#39;s storage account.
     * 
     */
    @Export(name="accountKey", refs={String.class}, tree="[0]")
    private Output<String> accountKey;

    /**
     * @return A valid account key from the queue&#39;s storage account.
     * 
     */
    public Output<String> accountKey() {
        return this.accountKey;
    }
    /**
     * The queue&#39;s storage account name.
     * 
     */
    @Export(name="accountName", refs={String.class}, tree="[0]")
    private Output<String> accountName;

    /**
     * @return The queue&#39;s storage account name.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The name of the queue that will store the events.
     * 
     */
    @Export(name="queueName", refs={String.class}, tree="[0]")
    private Output<String> queueName;

    /**
     * @return The name of the queue that will store the events.
     * 
     */
    public Output<String> queueName() {
        return this.queueName;
    }
    /**
     * A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
     * 
     */
    @Export(name="runStateChangedEvent", refs={ServicehookStorageQueuePipelinesRunStateChangedEvent.class}, tree="[0]")
    private Output</* @Nullable */ ServicehookStorageQueuePipelinesRunStateChangedEvent> runStateChangedEvent;

    /**
     * @return A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
     * 
     */
    public Output<Optional<ServicehookStorageQueuePipelinesRunStateChangedEvent>> runStateChangedEvent() {
        return Codegen.optional(this.runStateChangedEvent);
    }
    /**
     * A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`
     * 
     * &gt; **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
     * 
     */
    @Export(name="stageStateChangedEvent", refs={ServicehookStorageQueuePipelinesStageStateChangedEvent.class}, tree="[0]")
    private Output</* @Nullable */ ServicehookStorageQueuePipelinesStageStateChangedEvent> stageStateChangedEvent;

    /**
     * @return A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`
     * 
     * &gt; **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
     * 
     */
    public Output<Optional<ServicehookStorageQueuePipelinesStageStateChangedEvent>> stageStateChangedEvent() {
        return Codegen.optional(this.stageStateChangedEvent);
    }
    /**
     * event time-to-live - the duration a message can remain in the queue before it&#39;s automatically removed. Defaults to `604800`.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return event time-to-live - the duration a message can remain in the queue before it&#39;s automatically removed. Defaults to `604800`.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * event visibility timout - how long a message is invisible to other consumers after it&#39;s been dequeued. Defaults to `0`.
     * 
     */
    @Export(name="visiTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> visiTimeout;

    /**
     * @return event visibility timout - how long a message is invisible to other consumers after it&#39;s been dequeued. Defaults to `0`.
     * 
     */
    public Output<Optional<Integer>> visiTimeout() {
        return Codegen.optional(this.visiTimeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicehookStorageQueuePipelines(String name) {
        this(name, ServicehookStorageQueuePipelinesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicehookStorageQueuePipelines(String name, ServicehookStorageQueuePipelinesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicehookStorageQueuePipelines(String name, ServicehookStorageQueuePipelinesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines", name, args == null ? ServicehookStorageQueuePipelinesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServicehookStorageQueuePipelines(String name, Output<String> id, @Nullable ServicehookStorageQueuePipelinesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountKey"
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
    public static ServicehookStorageQueuePipelines get(String name, Output<String> id, @Nullable ServicehookStorageQueuePipelinesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicehookStorageQueuePipelines(name, id, state, options);
    }
}