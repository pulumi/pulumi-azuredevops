// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a Visual Studio Marketplace service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Azure DevOps Extension Tasks](https://marketplace.visualstudio.com/items?itemName=ms-devlabs.vsts-developer-tools-build-tasks)
 *
 * ## Example Usage
 *
 * ### Authorize with token
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
 * const exampleServiceendpointVisualstudiomarketplace = new azuredevops.ServiceendpointVisualstudiomarketplace("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Marketplace",
 *     url: "https://markpetplace.com",
 *     authenticationToken: {
 *         token: "token",
 *     },
 *     description: "Managed by Pulumi",
 * });
 * ```
 *
 * ### Authorize with username and password
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
 * const exampleServiceendpointVisualstudiomarketplace = new azuredevops.ServiceendpointVisualstudiomarketplace("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Marketplace",
 *     url: "https://markpetplace.com",
 *     authenticationBasic: {
 *         username: "username",
 *         password: "password",
 *     },
 *     description: "Managed by Pulumi",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Visual Studio Marketplace Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointVisualstudiomarketplace:ServiceendpointVisualstudiomarketplace example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointVisualstudiomarketplace extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointVisualstudiomarketplace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointVisualstudiomarketplaceState, opts?: pulumi.CustomResourceOptions): ServiceendpointVisualstudiomarketplace {
        return new ServiceendpointVisualstudiomarketplace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointVisualstudiomarketplace:ServiceendpointVisualstudiomarketplace';

    /**
     * Returns true if the given object is an instance of ServiceendpointVisualstudiomarketplace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointVisualstudiomarketplace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointVisualstudiomarketplace.__pulumiType;
    }

    /**
     * An `authenticationBasic` block as documented below.
     *
     * > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
     */
    public readonly authenticationBasic!: pulumi.Output<outputs.ServiceendpointVisualstudiomarketplaceAuthenticationBasic | undefined>;
    /**
     * An `authenticationToken` block as documented below.
     */
    public readonly authenticationToken!: pulumi.Output<outputs.ServiceendpointVisualstudiomarketplaceAuthenticationToken | undefined>;
    public /*out*/ readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The server URL for Visual Studio Marketplace.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointVisualstudiomarketplace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointVisualstudiomarketplaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointVisualstudiomarketplaceArgs | ServiceendpointVisualstudiomarketplaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointVisualstudiomarketplaceState | undefined;
            resourceInputs["authenticationBasic"] = state ? state.authenticationBasic : undefined;
            resourceInputs["authenticationToken"] = state ? state.authenticationToken : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceendpointVisualstudiomarketplaceArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["authenticationBasic"] = args ? args.authenticationBasic : undefined;
            resourceInputs["authenticationToken"] = args ? args.authenticationToken : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["authorization"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceendpointVisualstudiomarketplace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointVisualstudiomarketplace resources.
 */
export interface ServiceendpointVisualstudiomarketplaceState {
    /**
     * An `authenticationBasic` block as documented below.
     *
     * > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
     */
    authenticationBasic?: pulumi.Input<inputs.ServiceendpointVisualstudiomarketplaceAuthenticationBasic>;
    /**
     * An `authenticationToken` block as documented below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceendpointVisualstudiomarketplaceAuthenticationToken>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The server URL for Visual Studio Marketplace.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointVisualstudiomarketplace resource.
 */
export interface ServiceendpointVisualstudiomarketplaceArgs {
    /**
     * An `authenticationBasic` block as documented below.
     *
     * > **NOTE:** `authenticationBasic` and `authenticationToken` conflict with each other, only one is required.
     */
    authenticationBasic?: pulumi.Input<inputs.ServiceendpointVisualstudiomarketplaceAuthenticationBasic>;
    /**
     * An `authenticationToken` block as documented below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceendpointVisualstudiomarketplaceAuthenticationToken>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The server URL for Visual Studio Marketplace.
     */
    url: pulumi.Input<string>;
}
