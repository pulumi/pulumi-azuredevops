// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a AWS service endpoint within Azure DevOps. Using this service endpoint requires you to first install [AWS Toolkit for Azure DevOps](https://marketplace.visualstudio.com/items?itemName=AmazonWebServices.aws-vsts-tools).
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
 * const serviceendpoint = new azuredevops.ServiceEndpointAws("serviceendpoint", {
 *     projectId: project.id,
 *     serviceEndpointName: "aws_serviceendpoint",
 *     description: "Managed by AzureDevOps",
 *     accessKeyId: "xxxx",
 *     secretAccessKey: "xxxx",
 * });
 * ```
 * ## Relevant Links
 *
 * * [aws-toolkit-azure-devops](https://github.com/aws/aws-toolkit-azure-devops)
 * * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
 */
export class ServiceEndpointAws extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointAws resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceEndpointAwsState, opts?: pulumi.CustomResourceOptions): ServiceEndpointAws {
        return new ServiceEndpointAws(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceEndpointAws:ServiceEndpointAws';

    /**
     * Returns true if the given object is an instance of ServiceEndpointAws.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointAws {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointAws.__pulumiType;
    }

    /**
     * The AWS access key ID for signing programmatic requests.
     */
    public readonly accessKeyId!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<{[key: string]: string}>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Optional identifier for the assumed role session.
     */
    public readonly roleSessionName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the role to assume.
     */
    public readonly roleToAssume!: pulumi.Output<string | undefined>;
    /**
     * The AWS secret access key for signing programmatic requests.
     */
    public readonly secretAccessKey!: pulumi.Output<string>;
    /**
     * A bcrypted hash of the attribute 'secret_access_key'
     */
    public /*out*/ readonly secretAccessKeyHash!: pulumi.Output<string>;
    /**
     * The Service Endpoint name.
     */
    public readonly serviceEndpointName!: pulumi.Output<string>;
    /**
     * The AWS session token for signing programmatic requests.
     */
    public readonly sessionToken!: pulumi.Output<string | undefined>;
    /**
     * A bcrypted hash of the attribute 'session_token'
     */
    public /*out*/ readonly sessionTokenHash!: pulumi.Output<string>;

    /**
     * Create a ServiceEndpointAws resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointAwsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceEndpointAwsArgs | ServiceEndpointAwsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceEndpointAwsState | undefined;
            inputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["externalId"] = state ? state.externalId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["roleSessionName"] = state ? state.roleSessionName : undefined;
            inputs["roleToAssume"] = state ? state.roleToAssume : undefined;
            inputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
            inputs["secretAccessKeyHash"] = state ? state.secretAccessKeyHash : undefined;
            inputs["serviceEndpointName"] = state ? state.serviceEndpointName : undefined;
            inputs["sessionToken"] = state ? state.sessionToken : undefined;
            inputs["sessionTokenHash"] = state ? state.sessionTokenHash : undefined;
        } else {
            const args = argsOrState as ServiceEndpointAwsArgs | undefined;
            if (!args || args.accessKeyId === undefined) {
                throw new Error("Missing required property 'accessKeyId'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.secretAccessKey === undefined) {
                throw new Error("Missing required property 'secretAccessKey'");
            }
            if (!args || args.serviceEndpointName === undefined) {
                throw new Error("Missing required property 'serviceEndpointName'");
            }
            inputs["accessKeyId"] = args ? args.accessKeyId : undefined;
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["externalId"] = args ? args.externalId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["roleSessionName"] = args ? args.roleSessionName : undefined;
            inputs["roleToAssume"] = args ? args.roleToAssume : undefined;
            inputs["secretAccessKey"] = args ? args.secretAccessKey : undefined;
            inputs["serviceEndpointName"] = args ? args.serviceEndpointName : undefined;
            inputs["sessionToken"] = args ? args.sessionToken : undefined;
            inputs["secretAccessKeyHash"] = undefined /*out*/;
            inputs["sessionTokenHash"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceEndpointAws.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceEndpointAws resources.
 */
export interface ServiceEndpointAwsState {
    /**
     * The AWS access key ID for signing programmatic requests.
     */
    readonly accessKeyId?: pulumi.Input<string>;
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly description?: pulumi.Input<string>;
    /**
     * A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Optional identifier for the assumed role session.
     */
    readonly roleSessionName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the role to assume.
     */
    readonly roleToAssume?: pulumi.Input<string>;
    /**
     * The AWS secret access key for signing programmatic requests.
     */
    readonly secretAccessKey?: pulumi.Input<string>;
    /**
     * A bcrypted hash of the attribute 'secret_access_key'
     */
    readonly secretAccessKeyHash?: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName?: pulumi.Input<string>;
    /**
     * The AWS session token for signing programmatic requests.
     */
    readonly sessionToken?: pulumi.Input<string>;
    /**
     * A bcrypted hash of the attribute 'session_token'
     */
    readonly sessionTokenHash?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceEndpointAws resource.
 */
export interface ServiceEndpointAwsArgs {
    /**
     * The AWS access key ID for signing programmatic requests.
     */
    readonly accessKeyId: pulumi.Input<string>;
    readonly authorization?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly description?: pulumi.Input<string>;
    /**
     * A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * Optional identifier for the assumed role session.
     */
    readonly roleSessionName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the role to assume.
     */
    readonly roleToAssume?: pulumi.Input<string>;
    /**
     * The AWS secret access key for signing programmatic requests.
     */
    readonly secretAccessKey: pulumi.Input<string>;
    /**
     * The Service Endpoint name.
     */
    readonly serviceEndpointName: pulumi.Input<string>;
    /**
     * The AWS session token for signing programmatic requests.
     */
    readonly sessionToken?: pulumi.Input<string>;
}
