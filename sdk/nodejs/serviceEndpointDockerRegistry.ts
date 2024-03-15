// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Docker Registry service endpoint within Azure DevOps.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
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
 * // dockerhub registry service connection
 * const exampleServiceEndpointDockerRegistry = new azuredevops.ServiceEndpointDockerRegistry("exampleServiceEndpointDockerRegistry", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example Docker Hub",
 *     dockerUsername: "example",
 *     dockerEmail: "email@example.com",
 *     dockerPassword: "12345",
 *     registryType: "DockerHub",
 * });
 * // other docker registry service connection
 * const example_other = new azuredevops.ServiceEndpointDockerRegistry("example-other", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "Example Docker Registry",
 *     dockerRegistry: "https://sample.azurecr.io/v1",
 *     dockerUsername: "sample",
 *     dockerPassword: "12345",
 *     registryType: "Others",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 * - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Docker Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointDockerRegistry extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointDockerRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointDockerRegistryState, opts?: pulumi.CustomResourceOptions): ServiceEndpointDockerRegistry {
        return new ServiceEndpointDockerRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry';

    /**
     * Returns true if the given object is an instance of ServiceEndpointDockerRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointDockerRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointDockerRegistry.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The email for Docker account user.
     */
    public readonly dockerEmail!: pulumi.Output<string | undefined>;
    /**
     * The password for the account user identified above.
     */
    public readonly dockerPassword!: pulumi.Output<string | undefined>;
    /**
     * The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
     */
    public readonly dockerRegistry!: pulumi.Output<string>;
    /**
     * The identifier of the Docker account user.
     */
    public readonly dockerUsername!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Can be "DockerHub" or "Others" (Default "DockerHub")
     */
    public readonly registryType!: pulumi.Output<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointDockerRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointDockerRegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointDockerRegistryArgs | ServiceEndpointDockerRegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointDockerRegistryState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dockerEmail"] = state ? state.dockerEmail : undefined;
            resourceInputs["dockerPassword"] = state ? state.dockerPassword : undefined;
            resourceInputs["dockerRegistry"] = state ? state.dockerRegistry : undefined;
            resourceInputs["dockerUsername"] = state ? state.dockerUsername : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["registryType"] = state ? state.registryType : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceEndpointDockerRegistryArgs | undefined;
            if ((!args || args.dockerRegistry === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dockerRegistry'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.registryType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryType'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dockerEmail"] = args ? args.dockerEmail : undefined;
            resourceInputs["dockerPassword"] = args?.dockerPassword ? pulumi.secret(args.dockerPassword) : undefined;
            resourceInputs["dockerRegistry"] = args ? args.dockerRegistry : undefined;
            resourceInputs["dockerUsername"] = args ? args.dockerUsername : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["registryType"] = args ? args.registryType : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["dockerPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceEndpointDockerRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointDockerRegistry resources.
 */
export interface ServiceEndpointDockerRegistryState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The email for Docker account user.
     */
    dockerEmail?: pulumi.Input<string>;
    /**
     * The password for the account user identified above.
     */
    dockerPassword?: pulumi.Input<string>;
    /**
     * The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
     */
    dockerRegistry?: pulumi.Input<string>;
    /**
     * The identifier of the Docker account user.
     */
    dockerUsername?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Can be "DockerHub" or "Others" (Default "DockerHub")
     */
    registryType?: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointDockerRegistry resource.
 */
export interface ServiceEndpointDockerRegistryArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The email for Docker account user.
     */
    dockerEmail?: pulumi.Input<string>;
    /**
     * The password for the account user identified above.
     */
    dockerPassword?: pulumi.Input<string>;
    /**
     * The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
     */
    dockerRegistry: pulumi.Input<string>;
    /**
     * The identifier of the Docker account user.
     */
    dockerUsername?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * Can be "DockerHub" or "Others" (Default "DockerHub")
     */
    registryType: pulumi.Input<string>;
    /**
     * The name you will use to refer to this service connection in task inputs.
     */
    serviceEndpointName: pulumi.Input<string>;
}
