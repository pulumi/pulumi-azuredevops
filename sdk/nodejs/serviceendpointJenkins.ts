// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Jenkins service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Jenkins instance.
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
 *     description: "Managed by Terraform",
 * });
 * const exampleServiceendpointJenkins = new azuredevops.ServiceendpointJenkins("exampleServiceendpointJenkins", {
 *     projectId: exampleProject.id,
 *     serviceEndpointName: "jenkins-example",
 *     description: "Service Endpoint for 'Jenkins' (Managed by Terraform)",
 *     url: "https://example.com",
 *     acceptUntrustedCerts: false,
 *     username: "username",
 *     password: "password",
 * });
 * ```
 *
 * ## Import
 *
 * Service Connection Jenkins can be imported using the `projectId/id` or or `projectName/id`, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins example projectName/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServiceendpointJenkins extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointJenkins resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointJenkinsState, opts?: pulumi.CustomResourceOptions): ServiceendpointJenkins {
        return new ServiceendpointJenkins(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointJenkins:ServiceendpointJenkins';

    /**
     * Returns true if the given object is an instance of ServiceendpointJenkins.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointJenkins {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointJenkins.__pulumiType;
    }

    /**
     * Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
     */
    public readonly acceptUntrustedCerts!: pulumi.Output<boolean | undefined>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Service Endpoint password to authenticate at the Jenkins Instance.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The Service Endpoint url.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The Service Endpoint username to authenticate at the Jenkins Instance.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a ServiceendpointJenkins resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointJenkinsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointJenkinsArgs | ServiceendpointJenkinsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointJenkinsState | undefined;
            resourceInputs["acceptUntrustedCerts"] = state ? state.acceptUntrustedCerts : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ServiceendpointJenkinsArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["acceptUntrustedCerts"] = args ? args.acceptUntrustedCerts : undefined;
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceendpointJenkins.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointJenkins resources.
 */
export interface ServiceendpointJenkinsState {
    /**
     * Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
     */
    acceptUntrustedCerts?: pulumi.Input<boolean>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The Service Endpoint password to authenticate at the Jenkins Instance.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
     */
    serviceEndpointName?: pulumi.Input<string>;
    /**
     * The Service Endpoint url.
     */
    url?: pulumi.Input<string>;
    /**
     * The Service Endpoint username to authenticate at the Jenkins Instance.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointJenkins resource.
 */
export interface ServiceendpointJenkinsArgs {
    /**
     * Allows the Jenkins clients to accept self-signed SSL server certificates. Defaults to `false.`
     */
    acceptUntrustedCerts?: pulumi.Input<boolean>;
    authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * The Service Endpoint password to authenticate at the Jenkins Instance.
     */
    password: pulumi.Input<string>;
    /**
     * The ID of the project. Changing this forces a new Service Connection Jenkins to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * The name of the service endpoint. Changing this forces a new Service Connection Jenkins to be created.
     */
    serviceEndpointName: pulumi.Input<string>;
    /**
     * The Service Endpoint url.
     */
    url: pulumi.Input<string>;
    /**
     * The Service Endpoint username to authenticate at the Jenkins Instance.
     */
    username: pulumi.Input<string>;
}
