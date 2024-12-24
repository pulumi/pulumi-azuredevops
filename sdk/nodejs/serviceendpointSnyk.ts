// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Snyk Security Scan service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Snyk Security Scan](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan)
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
 * const exampleServiceendpointSnyk = new azuredevops.ServiceendpointSnyk("example", {
 *     projectId: example.id,
 *     serverUrl: "https://snyk.io/",
 *     apiToken: "00000000-0000-0000-0000-000000000000",
 *     serviceEndpointName: "Example Snyk",
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
 * Azure DevOps Service Endpoint Snyk can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointSnyk extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointSnyk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointSnykState, opts?: pulumi.CustomResourceOptions): ServiceendpointSnyk {
        return new ServiceendpointSnyk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk';

    /**
     * Returns true if the given object is an instance of ServiceendpointSnyk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointSnyk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointSnyk.__pulumiType;
    }

    /**
     * The API token of the Snyk Security Scan.
     */
    public readonly apiToken!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The server URL of the Snyk Security Scan.
     */
    public readonly serverUrl!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointSnyk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointSnykArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointSnykArgs | ServiceendpointSnykState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointSnykState | undefined;
            resourceInputs["apiToken"] = state ? state.apiToken : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serverUrl"] = state ? state.serverUrl : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
        } else {
            const args = argsOrState as ServiceendpointSnykArgs | undefined;
            if ((!args || args.apiToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiToken'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serverUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverUrl'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            resourceInputs["apiToken"] = args?.apiToken ? pulumi.secret(args.apiToken) : undefined;
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serverUrl"] = args ? args.serverUrl : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceendpointSnyk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointSnyk resources.
 */
export interface ServiceendpointSnykState {
    /**
     * The API token of the Snyk Security Scan.
     */
    apiToken?: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The server URL of the Snyk Security Scan.
     */
    serverUrl?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointSnyk resource.
 */
export interface ServiceendpointSnykArgs {
    /**
     * The API token of the Snyk Security Scan.
     */
    apiToken: pulumi.Input<string>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The server URL of the Snyk Security Scan.
     */
    serverUrl: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
}
