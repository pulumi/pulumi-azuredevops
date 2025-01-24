// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a Storage Queue Pipelines Service Hook .
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "example-project"});
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     name: "example-resources",
 *     location: "West Europe",
 * });
 * const exampleAccount = new azure.storage.Account("example", {
 *     name: "servicehookexamplestacc",
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleQueue = new azure.storage.Queue("example", {
 *     name: "examplequeue",
 *     storageAccountName: exampleAccount.name,
 * });
 * const exampleServicehookStorageQueuePipelines = new azuredevops.ServicehookStorageQueuePipelines("example", {
 *     projectId: example.id,
 *     accountName: exampleAccount.name,
 *     accountKey: exampleAccount.primaryAccessKey,
 *     queueName: exampleQueue.name,
 *     visiTimeout: 30,
 *     runStateChangedEvent: {
 *         runStateFilter: "Completed",
 *         runResultFilter: "Succeeded",
 *     },
 * });
 * ```
 *
 * An empty configuration block will occur in all events triggering the associated action.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.ServicehookStorageQueuePipelines("example", {
 *     projectId: exampleAzuredevopsProject.id,
 *     accountName: exampleAzurermStorageAccount.name,
 *     accountKey: exampleAzurermStorageAccount.primaryAccessKey,
 *     queueName: exampleAzurermStorageQueue.name,
 *     visiTimeout: 30,
 *     runStateChangedEvent: {},
 * });
 * ```
 *
 * ## Import
 *
 * Storage Queue Pipelines Service Hook can be imported using the `resource id`, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines example 00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServicehookStorageQueuePipelines extends pulumi.CustomResource {
    /**
     * Get an existing ServicehookStorageQueuePipelines resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicehookStorageQueuePipelinesState, opts?: pulumi.CustomResourceOptions): ServicehookStorageQueuePipelines {
        return new ServicehookStorageQueuePipelines(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines';

    /**
     * Returns true if the given object is an instance of ServicehookStorageQueuePipelines.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicehookStorageQueuePipelines {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicehookStorageQueuePipelines.__pulumiType;
    }

    /**
     * A valid account key from the queue's storage account.
     */
    public readonly accountKey!: pulumi.Output<string>;
    /**
     * The queue's storage account name.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The name of the queue that will store the events.
     */
    public readonly queueName!: pulumi.Output<string>;
    /**
     * A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
     */
    public readonly runStateChangedEvent!: pulumi.Output<outputs.ServicehookStorageQueuePipelinesRunStateChangedEvent | undefined>;
    /**
     * A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
     *
     * > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
     */
    public readonly stageStateChangedEvent!: pulumi.Output<outputs.ServicehookStorageQueuePipelinesStageStateChangedEvent | undefined>;
    /**
     * event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
     */
    public readonly visiTimeout!: pulumi.Output<number | undefined>;

    /**
     * Create a ServicehookStorageQueuePipelines resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicehookStorageQueuePipelinesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicehookStorageQueuePipelinesArgs | ServicehookStorageQueuePipelinesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicehookStorageQueuePipelinesState | undefined;
            resourceInputs["accountKey"] = state ? state.accountKey : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["queueName"] = state ? state.queueName : undefined;
            resourceInputs["runStateChangedEvent"] = state ? state.runStateChangedEvent : undefined;
            resourceInputs["stageStateChangedEvent"] = state ? state.stageStateChangedEvent : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["visiTimeout"] = state ? state.visiTimeout : undefined;
        } else {
            const args = argsOrState as ServicehookStorageQueuePipelinesArgs | undefined;
            if ((!args || args.accountKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountKey'");
            }
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.queueName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queueName'");
            }
            resourceInputs["accountKey"] = args?.accountKey ? pulumi.secret(args.accountKey) : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["queueName"] = args ? args.queueName : undefined;
            resourceInputs["runStateChangedEvent"] = args ? args.runStateChangedEvent : undefined;
            resourceInputs["stageStateChangedEvent"] = args ? args.stageStateChangedEvent : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["visiTimeout"] = args ? args.visiTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServicehookStorageQueuePipelines.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicehookStorageQueuePipelines resources.
 */
export interface ServicehookStorageQueuePipelinesState {
    /**
     * A valid account key from the queue's storage account.
     */
    accountKey?: pulumi.Input<string>;
    /**
     * The queue's storage account name.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The name of the queue that will store the events.
     */
    queueName?: pulumi.Input<string>;
    /**
     * A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
     */
    runStateChangedEvent?: pulumi.Input<inputs.ServicehookStorageQueuePipelinesRunStateChangedEvent>;
    /**
     * A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
     *
     * > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
     */
    stageStateChangedEvent?: pulumi.Input<inputs.ServicehookStorageQueuePipelinesStageStateChangedEvent>;
    /**
     * event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
     */
    ttl?: pulumi.Input<number>;
    /**
     * event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
     */
    visiTimeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ServicehookStorageQueuePipelines resource.
 */
export interface ServicehookStorageQueuePipelinesArgs {
    /**
     * A valid account key from the queue's storage account.
     */
    accountKey: pulumi.Input<string>;
    /**
     * The queue's storage account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * The name of the queue that will store the events.
     */
    queueName: pulumi.Input<string>;
    /**
     * A `runStateChangedEvent` block as defined below. Conflicts with `stageStateChangedEvent`
     */
    runStateChangedEvent?: pulumi.Input<inputs.ServicehookStorageQueuePipelinesRunStateChangedEvent>;
    /**
     * A `stageStateChangedEvent` block as defined below. Conflicts with `runStateChangedEvent`
     *
     * > **Note** At least one of `runStateChangedEvent` and `stageStateChangedEvent` has to be set.
     */
    stageStateChangedEvent?: pulumi.Input<inputs.ServicehookStorageQueuePipelinesStageStateChangedEvent>;
    /**
     * event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
     */
    ttl?: pulumi.Input<number>;
    /**
     * event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
     */
    visiTimeout?: pulumi.Input<number>;
}
