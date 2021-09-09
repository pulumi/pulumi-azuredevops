// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Docker Registry service endpoint within Azure DevOps.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * // dockerhub registry service connection
 * const dockerhubregistry = new azuredevops.ServiceEndpointDockerRegistry("dockerhubregistry", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample Docker Hub",
 *     dockerUsername: "sample",
 *     dockerEmail: "email@example.com",
 *     dockerPassword: "12345",
 *     registryType: "DockerHub",
 * });
 * // other docker registry service connection
 * const otherregistry = new azuredevops.ServiceEndpointDockerRegistry("otherregistry", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample Docker Registry",
 *     dockerRegistry: "https://sample.azurecr.io/v1",
 *     dockerUsername: "sample",
 *     dockerPassword: "12345",
 *     registryType: "Others",
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 * - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Docker Registry can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 *
 * @deprecated azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry
 */
export class DockerRegistry extends pulumi.CustomResource {
    /**
     * Get an existing DockerRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DockerRegistryState, opts?: pulumi.CustomResourceOptions): DockerRegistry {
        pulumi.log.warn("DockerRegistry is deprecated: azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry")
        return new DockerRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:ServiceEndpoint/dockerRegistry:DockerRegistry';

    /**
     * Returns true if the given object is an instance of DockerRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DockerRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DockerRegistry.__pulumiType;
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
     * A bcrypted hash of the attribute 'docker_password'
     */
    public /*out*/ readonly dockerPasswordHash!: pulumi.Output<string>;
    /**
     * The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
     */
    public readonly dockerRegistry!: pulumi.Output<string>;
    /**
     * The identifier of the Docker account user.
     */
    public readonly dockerUsername!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
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
     * Create a DockerRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry */
    constructor(name: string, args: DockerRegistryArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry */
    constructor(name: string, argsOrState?: DockerRegistryArgs | DockerRegistryState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DockerRegistry is deprecated: azuredevops.serviceendpoint.DockerRegistry has been deprecated in favor of azuredevops.ServiceEndpointDockerRegistry")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DockerRegistryState | undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dockerEmail"] = state ? state.dockerEmail : undefined;
            inputs["dockerPassword"] = state ? state.dockerPassword : undefined;
            inputs["dockerPasswordHash"] = state ? state.dockerPasswordHash : undefined;
            inputs["dockerRegistry"] = state ? state.dockerRegistry : undefined;
            inputs["dockerUsername"] = state ? state.dockerUsername : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["registryType"] = state ? state.registryType : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as DockerRegistryArgs | undefined;
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
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dockerEmail"] = args ? args.dockerEmail : undefined;
            inputs["dockerPassword"] = args ? args.dockerPassword : undefined;
            inputs["dockerRegistry"] = args ? args.dockerRegistry : undefined;
            inputs["dockerUsername"] = args ? args.dockerUsername : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["registryType"] = args ? args.registryType : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            inputs["dockerPasswordHash"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DockerRegistry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DockerRegistry resources.
 */
export interface DockerRegistryState {
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
     * A bcrypted hash of the attribute 'docker_password'
     */
    dockerPasswordHash?: pulumi.Input<string>;
    /**
     * The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
     */
    dockerRegistry?: pulumi.Input<string>;
    /**
     * The identifier of the Docker account user.
     */
    dockerUsername?: pulumi.Input<string>;
    /**
     * The project ID or project name.
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
 * The set of arguments for constructing a DockerRegistry resource.
 */
export interface DockerRegistryArgs {
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
     * The project ID or project name.
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
