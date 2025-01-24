// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a Service Fabric service endpoint within Azure DevOps.
 *
 * ## Example Usage
 *
 * ### Azure Active Directory Authentication
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     name: "Sample Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const test = new azuredevops.ServiceEndpointServiceFabric("test", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample Service Fabric",
 *     description: "Managed by Pulumi",
 *     clusterEndpoint: "tcp://test",
 *     azureActiveDirectory: {
 *         serverCertificateLookup: "Thumbprint",
 *         serverCertificateThumbprint: "0000000000000000000000000000000000000000",
 *         username: "username",
 *         password: "password",
 *     },
 * });
 * ```
 *
 * ### Windows Authentication
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     name: "Sample Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const test = new azuredevops.ServiceEndpointServiceFabric("test", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample Service Fabric",
 *     description: "Managed by Pulumi",
 *     clusterEndpoint: "tcp://test",
 *     none: {
 *         unsecured: false,
 *         clusterSpn: "HTTP/www.contoso.com",
 *     },
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Fabric Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceEndpointServiceFabric extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointServiceFabric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointServiceFabricState, opts?: pulumi.CustomResourceOptions): ServiceEndpointServiceFabric {
        return new ServiceEndpointServiceFabric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric';

    /**
     * Returns true if the given object is an instance of ServiceEndpointServiceFabric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointServiceFabric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointServiceFabric.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    /**
     * An `azureActiveDirectory` block as documented below.
     */
    public readonly azureActiveDirectory!: pulumi.Output<outputs.ServiceEndpointServiceFabricAzureActiveDirectory | undefined>;
    /**
     * A `certificate` block as documented below.
     */
    public readonly certificate!: pulumi.Output<outputs.ServiceEndpointServiceFabricCertificate | undefined>;
    /**
     * Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
     */
    public readonly clusterEndpoint!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A `none` block as documented below.
     */
    public readonly none!: pulumi.Output<outputs.ServiceEndpointServiceFabricNone | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointServiceFabric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointServiceFabricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointServiceFabricArgs | ServiceEndpointServiceFabricState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointServiceFabricState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["azureActiveDirectory"] = state ? state.azureActiveDirectory : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["clusterEndpoint"] = state ? state.clusterEndpoint : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["none"] = state ? state.none : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceEndpointServiceFabricArgs | undefined;
            if ((!args || args.clusterEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterEndpoint'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["azureActiveDirectory"] = args ? args.azureActiveDirectory : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["clusterEndpoint"] = args ? args.clusterEndpoint : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["none"] = args ? args.none : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceEndpointServiceFabric.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointServiceFabric resources.
 */
export interface ServiceEndpointServiceFabricState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An `azureActiveDirectory` block as documented below.
     */
    azureActiveDirectory?: pulumi.Input<inputs.ServiceEndpointServiceFabricAzureActiveDirectory>;
    /**
     * A `certificate` block as documented below.
     */
    certificate?: pulumi.Input<inputs.ServiceEndpointServiceFabricCertificate>;
    /**
     * Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
     */
    clusterEndpoint?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * A `none` block as documented below.
     */
    none?: pulumi.Input<inputs.ServiceEndpointServiceFabricNone>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointServiceFabric resource.
 */
export interface ServiceEndpointServiceFabricArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An `azureActiveDirectory` block as documented below.
     */
    azureActiveDirectory?: pulumi.Input<inputs.ServiceEndpointServiceFabricAzureActiveDirectory>;
    /**
     * A `certificate` block as documented below.
     */
    certificate?: pulumi.Input<inputs.ServiceEndpointServiceFabricCertificate>;
    /**
     * Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
     */
    clusterEndpoint: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * A `none` block as documented below.
     */
    none?: pulumi.Input<inputs.ServiceEndpointServiceFabricNone>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
