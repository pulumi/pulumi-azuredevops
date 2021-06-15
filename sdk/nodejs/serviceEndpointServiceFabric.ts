// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a Service Fabric service endpoint within Azure DevOps.
 *
 * ## Example Usage
 * ### Azure Active Directory Authentication
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
 * const test = new azuredevops.ServiceEndpointServiceFabric("test", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample Service Fabric",
 *     description: "Managed by Terraform",
 *     clusterEndpoint: "tcp://test",
 *     azureActiveDirectory: {
 *         serverCertificateLookup: "Thumbprint",
 *         serverCertificateThumbprint: "0000000000000000000000000000000000000000",
 *         username: "username",
 *         password: "password",
 *     },
 * });
 * ```
 * ### Windows Authentication
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
 * const test = new azuredevops.ServiceEndpointServiceFabric("test", {
 *     projectId: project.id,
 *     serviceEndpointName: "Sample Service Fabric",
 *     description: "Managed by Terraform",
 *     clusterEndpoint: "tcp://test",
 *     none: {
 *         unsecured: false,
 *         clusterSpn: "HTTP/www.contoso.com",
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 5.1 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Service Fabric can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 *  $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
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
    public readonly azureActiveDirectory!: pulumi.Output<outputs.ServiceEndpointServiceFabricAzureActiveDirectory | undefined>;
    public readonly certificate!: pulumi.Output<outputs.ServiceEndpointServiceFabricCertificate | undefined>;
    /**
     * Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
     */
    public readonly clusterEndpoint!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly none!: pulumi.Output<outputs.ServiceEndpointServiceFabricNone | undefined>;
    /**
     * The project ID or project name.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceEndpointServiceFabricState | undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["azureActiveDirectory"] = state ? state.azureActiveDirectory : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["clusterEndpoint"] = state ? state.clusterEndpoint : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["none"] = state ? state.none : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
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
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["azureActiveDirectory"] = args ? args.azureActiveDirectory : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["clusterEndpoint"] = args ? args.clusterEndpoint : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["none"] = args ? args.none : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceEndpointServiceFabric.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointServiceFabric resources.
 */
export interface ServiceEndpointServiceFabricState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    azureActiveDirectory?: pulumi.Input<inputs.ServiceEndpointServiceFabricAzureActiveDirectory>;
    certificate?: pulumi.Input<inputs.ServiceEndpointServiceFabricCertificate>;
    /**
     * Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
     */
    clusterEndpoint?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    none?: pulumi.Input<inputs.ServiceEndpointServiceFabricNone>;
    /**
     * The project ID or project name.
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
    azureActiveDirectory?: pulumi.Input<inputs.ServiceEndpointServiceFabricAzureActiveDirectory>;
    certificate?: pulumi.Input<inputs.ServiceEndpointServiceFabricCertificate>;
    /**
     * Client connection endpoint for the cluster. Prefix the value with 'tcp://';. This value overrides the publish profile.
     */
    clusterEndpoint: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    none?: pulumi.Input<inputs.ServiceEndpointServiceFabricNone>;
    /**
     * The project ID or project name.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
