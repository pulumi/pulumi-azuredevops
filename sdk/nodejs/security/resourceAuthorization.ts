// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
 * const project = new azuredevops.Project("project", {});
 * const bitbucketAccount = new azuredevops.ServiceEndpointBitBucket("bitbucketAccount", {
 *     projectId: project.id,
 *     username: "xxxx",
 *     password: "xxxx",
 *     serviceEndpointName: "test-bitbucket",
 *     description: "test",
 * });
 * const auth = new azuredevops.ResourceAuthorization("auth", {
 *     projectId: project.id,
 *     resourceId: bitbucketAccount.id,
 *     authorized: true,
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Authorize Definition Resource](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/resources/authorize%20definition%20resources?view=azure-devops-rest-5.1)
 *
 * @deprecated azuredevops.security.ResourceAuthorization has been deprecated in favor of azuredevops.ResourceAuthorization
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
        pulumi.log.warn("ResourceAuthorization is deprecated: azuredevops.security.ResourceAuthorization has been deprecated in favor of azuredevops.ResourceAuthorization")
        return new ResourceAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Security/resourceAuthorization:ResourceAuthorization';

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
    /** @deprecated azuredevops.security.ResourceAuthorization has been deprecated in favor of azuredevops.ResourceAuthorization */
    constructor(name: string, args: ResourceAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.security.ResourceAuthorization has been deprecated in favor of azuredevops.ResourceAuthorization */
    constructor(name: string, argsOrState?: ResourceAuthorizationArgs | ResourceAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ResourceAuthorization is deprecated: azuredevops.security.ResourceAuthorization has been deprecated in favor of azuredevops.ResourceAuthorization")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceAuthorizationState | undefined;
            inputs["authorized"] = state ? state.authorized : undefined;
            inputs["definitionId"] = state ? state.definitionId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ResourceAuthorizationArgs | undefined;
            if ((!args || args.authorized === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'authorized'");
            }
            if ((!args || args.projectId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.resourceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceId'");
            }
            inputs["authorized"] = args ? args.authorized : undefined;
            inputs["definitionId"] = args ? args.definitionId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourceAuthorization.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceAuthorization resources.
 */
export interface ResourceAuthorizationState {
    /**
     * Set to true to allow public access in the project. Type: boolean.
     */
    readonly authorized?: pulumi.Input<boolean>;
    /**
     * The ID of the build definition to authorize. Type: string.
     */
    readonly definitionId?: pulumi.Input<number>;
    /**
     * The project ID or project name. Type: string.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The ID of the resource to authorize. Type: string.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceAuthorization resource.
 */
export interface ResourceAuthorizationArgs {
    /**
     * Set to true to allow public access in the project. Type: boolean.
     */
    readonly authorized: pulumi.Input<boolean>;
    /**
     * The ID of the build definition to authorize. Type: string.
     */
    readonly definitionId?: pulumi.Input<number>;
    /**
     * The project ID or project name. Type: string.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The ID of the resource to authorize. Type: string.
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
     */
    readonly type?: pulumi.Input<string>;
}
