// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages authorization of resources, e.g. for access in build pipelines.
 *
 * Currently supported resources: service endpoint (aka service connection, endpoint).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const exampleServiceEndpointBitBucket = new azuredevops.ServiceEndpointBitBucket("exampleServiceEndpointBitBucket", {
 *     projectId: exampleProject.id,
 *     username: "username",
 *     password: "password",
 *     serviceEndpointName: "example-bitbucket",
 *     description: "Managed by Terraform",
 * });
 * const exampleResourceAuthorization = new azuredevops.ResourceAuthorization("exampleResourceAuthorization", {
 *     projectId: exampleProject.id,
 *     resourceId: exampleServiceEndpointBitBucket.id,
 *     authorized: true,
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 6.0 - Authorize Definition Resource](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/resources/authorize%20definition%20resources?view=azure-devops-rest-6.0)
 */
export class ResourceAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing ResourceAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceAuthorizationState, opts?: pulumi.CustomResourceOptions): ResourceAuthorization {
        return new ResourceAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/resourceAuthorization:ResourceAuthorization';

    /**
     * Returns true if the given object is an instance of ResourceAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceAuthorization.__pulumiType;
    }

    /**
     * Set to true to allow public access in the project. Type: boolean.
     */
    public readonly authorized!: pulumi.Output<boolean>;
    /**
     * The ID of the build definition to authorize. Type: string.
     */
    public readonly definitionId!: pulumi.Output<number | undefined>;
    /**
     * The project ID or project name. Type: string.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The ID of the resource to authorize. Type: string.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ResourceAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceAuthorizationArgs | ResourceAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceAuthorizationState | undefined;
            resourceInputs["authorized"] = state ? state.authorized : undefined;
            resourceInputs["definitionId"] = state ? state.definitionId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ResourceAuthorizationArgs | undefined;
            if ((!args || args.authorized === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorized'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            resourceInputs["authorized"] = args ? args.authorized : undefined;
            resourceInputs["definitionId"] = args ? args.definitionId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azuredevops:Security/resourceAuthorization:ResourceAuthorization" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ResourceAuthorization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceAuthorization resources.
 */
export interface ResourceAuthorizationState {
    /**
     * Set to true to allow public access in the project. Type: boolean.
     */
    authorized?: pulumi.Input<boolean>;
    /**
     * The ID of the build definition to authorize. Type: string.
     */
    definitionId?: pulumi.Input<number>;
    /**
     * The project ID or project name. Type: string.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the resource to authorize. Type: string.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceAuthorization resource.
 */
export interface ResourceAuthorizationArgs {
    /**
     * Set to true to allow public access in the project. Type: boolean.
     */
    authorized: pulumi.Input<boolean>;
    /**
     * The ID of the build definition to authorize. Type: string.
     */
    definitionId?: pulumi.Input<number>;
    /**
     * The project ID or project name. Type: string.
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the resource to authorize. Type: string.
     */
    resourceId: pulumi.Input<string>;
    /**
     * The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
     */
    type?: pulumi.Input<string>;
}
