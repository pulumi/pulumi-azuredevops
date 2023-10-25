// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Bitbucket service endpoint within Azure DevOps.
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
 *     serviceEndpointName: "Example Bitbucket",
 *     description: "Managed by Terraform",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Bitbucket can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointBitBucket extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointBitBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointBitBucketState, opts?: pulumi.CustomResourceOptions): ServiceEndpointBitBucket {
        return new ServiceEndpointBitBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointBitBucket:ServiceEndpointBitBucket';

    /**
     * Returns true if the given object is an instance of ServiceEndpointBitBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointBitBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointBitBucket.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Bitbucket account password.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * Bitbucket account username.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointBitBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointBitBucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointBitBucketArgs | ServiceEndpointBitBucketState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointBitBucketState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceEndpointBitBucketArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azuredevops:ServiceEndpoint/bitBucket:BitBucket" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceEndpointBitBucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointBitBucket resources.
 */
export interface ServiceEndpointBitBucketState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Bitbucket account password.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * Bitbucket account username.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointBitBucket resource.
 */
export interface ServiceEndpointBitBucketArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Bitbucket account password.
     */
    password: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * Bitbucket account username.
     */
    username: pulumi.Input<string>;
}
