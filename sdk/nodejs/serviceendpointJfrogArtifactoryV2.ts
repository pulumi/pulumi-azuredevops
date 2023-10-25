// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a JFrog Artifactory V2 server endpoint within an Azure DevOps organization.
 *
 * > **Note:** Using this service endpoint requires you to first install [JFrog Extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-azure-devops-extension).
 *
 * ## Relevant Links
 *
 * * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
 * * [Artifactory User Token](https://docs.artifactory.org/latest/user-guide/user-token/)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint JFrog Artifactory V2 can be imported using the **projectID/serviceEndpointID**, e.g.
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2 example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointJfrogArtifactoryV2 extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointJfrogArtifactoryV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointJfrogArtifactoryV2State, opts?: pulumi.CustomResourceOptions): ServiceendpointJfrogArtifactoryV2 {
        return new ServiceendpointJfrogArtifactoryV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointJfrogArtifactoryV2:ServiceendpointJfrogArtifactoryV2';

    /**
     * Returns true if the given object is an instance of ServiceendpointJfrogArtifactoryV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointJfrogArtifactoryV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointJfrogArtifactoryV2.__pulumiType;
    }

    /**
     * A `authenticationBasic` block as documented below.
     */
    public readonly authenticationBasic!: pulumi.Output<outputs.ServiceendpointJfrogArtifactoryV2AuthenticationBasic | undefined>;
    /**
     * A `authenticationToken` block as documented below.
     */
    public readonly authenticationToken!: pulumi.Output<outputs.ServiceendpointJfrogArtifactoryV2AuthenticationToken | undefined>;
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
     * URL of the Artifactory server to connect with.
     *
     * > **NOTE:** URL should not end in a slash character.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointJfrogArtifactoryV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointJfrogArtifactoryV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointJfrogArtifactoryV2Args | ServiceendpointJfrogArtifactoryV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointJfrogArtifactoryV2State | undefined;
            resourceInputs["authenticationBasic"] = state ? state.authenticationBasic : undefined;
            resourceInputs["authenticationToken"] = state ? state.authenticationToken : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServiceendpointJfrogArtifactoryV2Args | undefined;
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
        super(ServiceendpointJfrogArtifactoryV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointJfrogArtifactoryV2 resources.
 */
export interface ServiceendpointJfrogArtifactoryV2State {
    /**
     * A `authenticationBasic` block as documented below.
     */
    authenticationBasic?: pulumi.Input<inputs.ServiceendpointJfrogArtifactoryV2AuthenticationBasic>;
    /**
     * A `authenticationToken` block as documented below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceendpointJfrogArtifactoryV2AuthenticationToken>;
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
     * > **NOTE:** URL should not end in a slash character.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointJfrogArtifactoryV2 resource.
 */
export interface ServiceendpointJfrogArtifactoryV2Args {
    /**
     * A `authenticationBasic` block as documented below.
     */
    authenticationBasic?: pulumi.Input<inputs.ServiceendpointJfrogArtifactoryV2AuthenticationBasic>;
    /**
     * A `authenticationToken` block as documented below.
     */
    authenticationToken?: pulumi.Input<inputs.ServiceendpointJfrogArtifactoryV2AuthenticationToken>;
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
     * URL of the Artifactory server to connect with.
     *
     * > **NOTE:** URL should not end in a slash character.
     */
    url: pulumi.Input<string>;
}
