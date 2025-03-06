// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages an JFrog Artifactory server endpoint within an Azure DevOps organization. Using this service endpoint requires you to first install [JFrog Artifactory Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-artifactory-vsts-extension).
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
 * const exampleServiceEndpointArtifactory = new azuredevops.ServiceEndpointArtifactory("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Artifactory",
 *     description: "Managed by Pulumi",
 *     url: "https://artifactory.my.com",
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
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Pulumi",
 * });
 * const exampleServiceEndpointArtifactory = new azuredevops.ServiceEndpointArtifactory("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Artifactory",
 *     description: "Managed by Pulumi",
 *     url: "https://artifactory.my.com",
 *     authenticationBasic: {
 *         username: "username",
 *         password: "password",
 *     },
 * });
 * ```
 *
 * ## Relevant Links
 *
 * * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
 * * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
 *
 * ## Import
 *
 * Azure DevOps JFrog Artifactory Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointArtifactory extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointArtifactory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointArtifactoryState, opts?: pulumi.CustomResourceOptions): ServiceEndpointArtifactory {
        return new ServiceEndpointArtifactory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointArtifactory:ServiceEndpointArtifactory';

    /**
     * Returns true if the given object is an instance of ServiceEndpointArtifactory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointArtifactory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointArtifactory.__pulumiType;
    }

    public readonly authenticationBasic!: pulumi.Output<outputs.ServiceEndpointArtifactoryAuthenticationBasic | undefined>;
    /**
     * A `authenticationBasic` block as defined below.
     */
    public readonly authenticationToken!: pulumi.Output<outputs.ServiceEndpointArtifactoryAuthenticationToken | undefined>;
    public /*out*/ readonly authorization!: pulumi.Output<{[key: string]: string}>;
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
     * URL of the Artifactory server to connect with.
     *
     * _**Note: URL should not end in a slash character.**_
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointArtifactory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointArtifactoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointArtifactoryArgs | ServiceEndpointArtifactoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointArtifactoryState | undefined;
            resourceInputs["authenticationBasic"] = state ? state.authenticationBasic : undefined;
            resourceInputs["authenticationToken"] = state ? state.authenticationToken : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceEndpointArtifactoryArgs | undefined;
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
        super(ServiceEndpointArtifactory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointArtifactory resources.
 */
export interface ServiceEndpointArtifactoryState {
    authenticationBasic?: pulumi.Input<inputs.ServiceEndpointArtifactoryAuthenticationBasic>;
    /**
     * A `authenticationBasic` block as defined below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceEndpointArtifactoryAuthenticationToken>;
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
     * URL of the Artifactory server to connect with.
     *
     * _**Note: URL should not end in a slash character.**_
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointArtifactory resource.
 */
export interface ServiceEndpointArtifactoryArgs {
    authenticationBasic?: pulumi.Input<inputs.ServiceEndpointArtifactoryAuthenticationBasic>;
    /**
     * A `authenticationBasic` block as defined below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceEndpointArtifactoryAuthenticationToken>;
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
     * URL of the Artifactory server to connect with.
     *
     * _**Note: URL should not end in a slash character.**_
     */
    url: pulumi.Input<string>;
}
