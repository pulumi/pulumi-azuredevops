// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a ArgoCD service endpoint within Azure DevOps. Using this service endpoint requires you to first install [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd).
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
 * });
 * const exampleServiceendpointArgocd = new azuredevops.ServiceendpointArgocd("exampleServiceendpointArgocd", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example ArgoCD",
 *     description: "Managed by Terraform",
 *     url: "https://argocd.my.com",
 *     authenticationToken: {
 *         token: "0000000000000000000000000000000000000000",
 *     },
 * });
 * ```
 * Alternatively a username and password may be used.
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
 * const exampleServiceendpointArgocd = new azuredevops.ServiceendpointArgocd("exampleServiceendpointArgocd", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example ArgoCD",
 *     description: "Managed by Terraform",
 *     url: "https://argocd.my.com",
 *     authenticationBasic: {
 *         username: "username",
 *         password: "password",
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
 * - [ArgoCD Project/User Token](https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_account_generate-token/)
 * - [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint ArgoCD can be imported using the **projectID/serviceEndpointID**, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointArgocd extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointArgocd resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointArgocdState, opts?: pulumi.CustomResourceOptions): ServiceendpointArgocd {
        return new ServiceendpointArgocd(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointArgocd:ServiceendpointArgocd';

    /**
     * Returns true if the given object is an instance of ServiceendpointArgocd.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointArgocd {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointArgocd.__pulumiType;
    }

    /**
     * An `authenticationBasic` block for the ArgoCD as documented below.
     */
    public readonly authenticationBasic!: pulumi.Output<outputs.ServiceendpointArgocdAuthenticationBasic | undefined>;
    /**
     * An `authenticationToken` block for the ArgoCD as documented below.
     */
    public readonly authenticationToken!: pulumi.Output<outputs.ServiceendpointArgocdAuthenticationToken | undefined>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Service Endpoint description.
     */
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
     * URL of the ArgoCD server to connect with.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointArgocd resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointArgocdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointArgocdArgs | ServiceendpointArgocdState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointArgocdState | undefined;
            resourceInputs["authenticationBasic"] = state ? state.authenticationBasic : undefined;
            resourceInputs["authenticationToken"] = state ? state.authenticationToken : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceendpointArgocdArgs | undefined;
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
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceendpointArgocd.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointArgocd resources.
 */
export interface ServiceendpointArgocdState {
    /**
     * An `authenticationBasic` block for the ArgoCD as documented below.
     */
    authenticationBasic?: pulumi.Input<inputs.ServiceendpointArgocdAuthenticationBasic>;
    /**
     * An `authenticationToken` block for the ArgoCD as documented below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceendpointArgocdAuthenticationToken>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Service Endpoint description.
     */
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
     * URL of the ArgoCD server to connect with.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointArgocd resource.
 */
export interface ServiceendpointArgocdArgs {
    /**
     * An `authenticationBasic` block for the ArgoCD as documented below.
     */
    authenticationBasic?: pulumi.Input<inputs.ServiceendpointArgocdAuthenticationBasic>;
    /**
     * An `authenticationToken` block for the ArgoCD as documented below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceendpointArgocdAuthenticationToken>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Service Endpoint description.
     */
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
     * URL of the ArgoCD server to connect with.
     */
    url: pulumi.Input<string>;
}
