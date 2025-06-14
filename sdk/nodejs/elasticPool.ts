// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Elastic pool within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Pulumi",
 * });
 * const exampleServiceEndpointAzureRM = new azuredevops.ServiceEndpointAzureRM("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Azure Connection",
 *     description: "Managed by Pulumi",
 *     serviceEndpointAuthenticationScheme: "ServicePrincipal",
 *     credentials: {
 *         serviceprincipalid: "00000000-0000-0000-0000-000000000000",
 *         serviceprincipalkey: "00000000-0000-0000-0000-000000000000",
 *     },
 *     azurermSpnTenantid: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionId: "00000000-0000-0000-0000-000000000000",
 *     azurermSubscriptionName: "Subscription Name",
 * });
 * const exampleElasticPool = new azuredevops.ElasticPool("example", {
 *     name: "Example Elastic Pool",
 *     serviceEndpointId: exampleServiceEndpointAzureRM.id,
 *     serviceEndpointScope: example.id,
 *     desiredIdle: 2,
 *     maxCapacity: 3,
 *     azureResourceId: "/subscriptions/<Subscription Id>/resourceGroups/<Resource Name>/providers/Microsoft.Compute/virtualMachineScaleSets/<VMSS Name>",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Elastic Pools](https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/elasticpools/create?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Agent Pools can be imported using the Elastic pool ID, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/elasticPool:ElasticPool example 0
 * ```
 */
export class ElasticPool extends pulumi.CustomResource {
    /**
     * Get an existing ElasticPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ElasticPoolState, opts?: pulumi.CustomResourceOptions): ElasticPool {
        return new ElasticPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/elasticPool:ElasticPool';

    /**
     * Returns true if the given object is an instance of ElasticPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElasticPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElasticPool.__pulumiType;
    }

    /**
     * Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     */
    public readonly agentInteractiveUi!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     */
    public readonly autoProvision!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     */
    public readonly autoUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Azure resource.
     */
    public readonly azureResourceId!: pulumi.Output<string>;
    /**
     * Number of agents to keep on standby.
     */
    public readonly desiredIdle!: pulumi.Output<number>;
    /**
     * Maximum number of virtual machines in the scale set.
     */
    public readonly maxCapacity!: pulumi.Output<number>;
    /**
     * The name of the Elastic pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project where a new Elastic Pool will be created.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * Tear down virtual machines after every use. Defaults to `false`.
     */
    public readonly recycleAfterEachUse!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of Service Endpoint used to connect to Azure.
     */
    public readonly serviceEndpointId!: pulumi.Output<string>;
    /**
     * The Project ID of Service Endpoint belongs to.
     */
    public readonly serviceEndpointScope!: pulumi.Output<string>;
    /**
     * Delay in minutes before deleting excess idle agents. Defaults to `30`.
     */
    public readonly timeToLiveMinutes!: pulumi.Output<number | undefined>;

    /**
     * Create a ElasticPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElasticPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ElasticPoolArgs | ElasticPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ElasticPoolState | undefined;
            resourceInputs["agentInteractiveUi"] = state ? state.agentInteractiveUi : undefined;
            resourceInputs["autoProvision"] = state ? state.autoProvision : undefined;
            resourceInputs["autoUpdate"] = state ? state.autoUpdate : undefined;
            resourceInputs["azureResourceId"] = state ? state.azureResourceId : undefined;
            resourceInputs["desiredIdle"] = state ? state.desiredIdle : undefined;
            resourceInputs["maxCapacity"] = state ? state.maxCapacity : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["recycleAfterEachUse"] = state ? state.recycleAfterEachUse : undefined;
            resourceInputs["serviceEndpointId"] = state ? state.serviceEndpointId : undefined;
            resourceInputs["serviceEndpointScope"] = state ? state.serviceEndpointScope : undefined;
            resourceInputs["timeToLiveMinutes"] = state ? state.timeToLiveMinutes : undefined;
        } else {
            const args = argsOrState as ElasticPoolArgs | undefined;
            if ((!args || args.azureResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azureResourceId'");
            }
            if ((!args || args.desiredIdle === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desiredIdle'");
            }
            if ((!args || args.maxCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxCapacity'");
            }
            if ((!args || args.serviceEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointId'");
            }
            if ((!args || args.serviceEndpointScope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointScope'");
            }
            resourceInputs["agentInteractiveUi"] = args ? args.agentInteractiveUi : undefined;
            resourceInputs["autoProvision"] = args ? args.autoProvision : undefined;
            resourceInputs["autoUpdate"] = args ? args.autoUpdate : undefined;
            resourceInputs["azureResourceId"] = args ? args.azureResourceId : undefined;
            resourceInputs["desiredIdle"] = args ? args.desiredIdle : undefined;
            resourceInputs["maxCapacity"] = args ? args.maxCapacity : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["recycleAfterEachUse"] = args ? args.recycleAfterEachUse : undefined;
            resourceInputs["serviceEndpointId"] = args ? args.serviceEndpointId : undefined;
            resourceInputs["serviceEndpointScope"] = args ? args.serviceEndpointScope : undefined;
            resourceInputs["timeToLiveMinutes"] = args ? args.timeToLiveMinutes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ElasticPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ElasticPool resources.
 */
export interface ElasticPoolState {
    /**
     * Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     */
    agentInteractiveUi?: pulumi.Input<boolean>;
    /**
     * Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     */
    autoProvision?: pulumi.Input<boolean>;
    /**
     * Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     */
    autoUpdate?: pulumi.Input<boolean>;
    /**
     * The ID of the Azure resource.
     */
    azureResourceId?: pulumi.Input<string>;
    /**
     * Number of agents to keep on standby.
     */
    desiredIdle?: pulumi.Input<number>;
    /**
     * Maximum number of virtual machines in the scale set.
     */
    maxCapacity?: pulumi.Input<number>;
    /**
     * The name of the Elastic pool.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project where a new Elastic Pool will be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Tear down virtual machines after every use. Defaults to `false`.
     */
    recycleAfterEachUse?: pulumi.Input<boolean>;
    /**
     * The ID of Service Endpoint used to connect to Azure.
     */
    serviceEndpointId?: pulumi.Input<string>;
    /**
     * The Project ID of Service Endpoint belongs to.
     */
    serviceEndpointScope?: pulumi.Input<string>;
    /**
     * Delay in minutes before deleting excess idle agents. Defaults to `30`.
     */
    timeToLiveMinutes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ElasticPool resource.
 */
export interface ElasticPoolArgs {
    /**
     * Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     */
    agentInteractiveUi?: pulumi.Input<boolean>;
    /**
     * Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     */
    autoProvision?: pulumi.Input<boolean>;
    /**
     * Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     */
    autoUpdate?: pulumi.Input<boolean>;
    /**
     * The ID of the Azure resource.
     */
    azureResourceId: pulumi.Input<string>;
    /**
     * Number of agents to keep on standby.
     */
    desiredIdle: pulumi.Input<number>;
    /**
     * Maximum number of virtual machines in the scale set.
     */
    maxCapacity: pulumi.Input<number>;
    /**
     * The name of the Elastic pool.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project where a new Elastic Pool will be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Tear down virtual machines after every use. Defaults to `false`.
     */
    recycleAfterEachUse?: pulumi.Input<boolean>;
    /**
     * The ID of Service Endpoint used to connect to Azure.
     */
    serviceEndpointId: pulumi.Input<string>;
    /**
     * The Project ID of Service Endpoint belongs to.
     */
    serviceEndpointScope: pulumi.Input<string>;
    /**
     * Delay in minutes before deleting excess idle agents. Defaults to `30`.
     */
    timeToLiveMinutes?: pulumi.Input<number>;
}
