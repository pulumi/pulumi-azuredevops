// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an agent queue within Azure DevOps. In the UI, this is equivelant to adding an
 * Organization defined pool to a project.
 *
 * The created queue is not authorized for use by all pipeliens in the project. However,
 * the `azuredevops.ResourceAuthorization` resource can be used to grant authorization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {projectName: "Sample Project"});
 * const pool = azuredevops.getPool({
 *     name: "contoso-pool",
 * });
 * const queue = new azuredevops.Queue("queue", {
 *     projectId: project.id,
 *     agentPoolId: pool.then(pool => pool.id),
 * });
 * // Grant acccess to queue to all pipelines in the project
 * const auth = new azuredevops.ResourceAuthorization("auth", {
 *     projectId: project.id,
 *     resourceId: queue.id,
 *     type: "queue",
 *     authorized: true,
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-5.1)
 *
 * @deprecated azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        pulumi.log.warn("Queue is deprecated: azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue")
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Agent/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * The ID of the organization agent pool.
     */
    public readonly agentPoolId!: pulumi.Output<number>;
    /**
     * The ID of the project in which to create the resource.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue */
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Queue is deprecated: azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QueueState | undefined;
            inputs["agentPoolId"] = state ? state.agentPoolId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            if (!args || args.agentPoolId === undefined) {
                throw new Error("Missing required property 'agentPoolId'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["agentPoolId"] = args ? args.agentPoolId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Queue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * The ID of the organization agent pool.
     */
    readonly agentPoolId?: pulumi.Input<number>;
    /**
     * The ID of the project in which to create the resource.
     */
    readonly projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * The ID of the organization agent pool.
     */
    readonly agentPoolId: pulumi.Input<number>;
    /**
     * The ID of the project in which to create the resource.
     */
    readonly projectId: pulumi.Input<string>;
}
