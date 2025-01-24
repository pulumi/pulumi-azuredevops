// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Checkmarx SAST service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Checkmarx SAST](https://marketplace.visualstudio.com/items?itemName=checkmarx.cxsast)
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
 * const exampleServiceendpointCheckmarxSast = new azuredevops.ServiceendpointCheckmarxSast("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Checkmarx SAST",
 *     serverUrl: "https://server.com",
 *     username: "username",
 *     password: "password",
 *     team: "team",
 *     preset: "preset",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
 *
 * ## Import
 *
 * Azure DevOps Service Endpoint Check Marx SAST can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointCheckmarxSast:ServiceendpointCheckmarxSast example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointCheckmarxSast extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointCheckmarxSast resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointCheckmarxSastState, opts?: pulumi.CustomResourceOptions): ServiceendpointCheckmarxSast {
        return new ServiceendpointCheckmarxSast(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointCheckmarxSast:ServiceendpointCheckmarxSast';

    /**
     * Returns true if the given object is an instance of ServiceendpointCheckmarxSast.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointCheckmarxSast {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointCheckmarxSast.__pulumiType;
    }

    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The password of the Checkmarx SAST.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
     */
    public readonly preset!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Server URL of the Checkmarx SAST.
     */
    public readonly serverUrl!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The full team name of the Checkmarx.
     */
    public readonly team!: pulumi.Output<string | undefined>;
    /**
     * The username of the Checkmarx SAST.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointCheckmarxSast resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointCheckmarxSastArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointCheckmarxSastArgs | ServiceendpointCheckmarxSastState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointCheckmarxSastState | undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["preset"] = state ? state.preset : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serverUrl"] = state ? state.serverUrl : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["team"] = state ? state.team : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceendpointCheckmarxSastArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
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
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["preset"] = args ? args.preset : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serverUrl"] = args ? args.serverUrl : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["team"] = args ? args.team : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceendpointCheckmarxSast.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointCheckmarxSast resources.
 */
export interface ServiceendpointCheckmarxSastState {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The password of the Checkmarx SAST.
     */
    password?: pulumi.Input<string>;
    /**
     * Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
     */
    preset?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Server URL of the Checkmarx SAST.
     */
    serverUrl?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The full team name of the Checkmarx.
     */
    team?: pulumi.Input<string>;
    /**
     * The username of the Checkmarx SAST.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointCheckmarxSast resource.
 */
export interface ServiceendpointCheckmarxSastArgs {
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The password of the Checkmarx SAST.
     */
    password: pulumi.Input<string>;
    /**
     * Predefined sets of queries that you can select when Creating, Configuring and Branching Projects. Predefined presets are provided by Checkmarx and you can configure your own. You can also import and export presets (on the server).In Service Connection if preset(optional) value is added, then it will igonres Preset available in pipeline and uses preset available in service connection only.If Preset is blank in service connection then it will use pipelines preset.
     */
    preset?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The Server URL of the Checkmarx SAST.
     */
    serverUrl: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The full team name of the Checkmarx.
     */
    team?: pulumi.Input<string>;
    /**
     * The username of the Checkmarx SAST.
     */
    username: pulumi.Input<string>;
}
