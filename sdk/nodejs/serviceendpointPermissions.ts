// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for a Service Endpoint
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Permission levels
 *
 * Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
 * Those levels are reflected by specifying (or omitting) values for the arguments `projectId` and `serviceendpointId`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Pulumi",
 * });
 * const example_readers = azuredevops.getGroupOutput({
 *     projectId: example.id,
 *     name: "Readers",
 * });
 * const example_root_permissions = new azuredevops.ServiceendpointPermissions("example-root-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     permissions: {
 *         Use: "allow",
 *         Administer: "allow",
 *         Create: "allow",
 *         ViewAuthorization: "allow",
 *         ViewEndpoint: "allow",
 *     },
 * });
 * const exampleServiceEndpointDockerRegistry = new azuredevops.ServiceEndpointDockerRegistry("example", {
 *     projectId: example.id,
 *     serviceEndpointName: "Example Docker Hub",
 *     dockerUsername: "username",
 *     dockerEmail: "email@example.com",
 *     dockerPassword: "password",
 *     registryType: "DockerHub",
 * });
 * const example_permissions = new azuredevops.ServiceendpointPermissions("example-permissions", {
 *     projectId: example.id,
 *     principal: example_readers.apply(example_readers => example_readers.id),
 *     serviceendpointId: exampleServiceEndpointDockerRegistry.id,
 *     permissions: {
 *         Use: "allow",
 *         Administer: "deny",
 *         Create: "deny",
 *         ViewAuthorization: "allow",
 *         ViewEndpoint: "allow",
 *     },
 * });
 * ```
 *
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 *
 * ## Import
 *
 * The resource does not support import.
 */
export class ServiceendpointPermissions extends pulumi.CustomResource {
    /**
     * Get an existing ServiceendpointPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceendpointPermissionsState, opts?: pulumi.CustomResourceOptions): ServiceendpointPermissions {
        return new ServiceendpointPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions';

    /**
     * Returns true if the given object is an instance of ServiceendpointPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceendpointPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceendpointPermissions.__pulumiType;
    }

    /**
     * the permissions to assign. The following permissions are available.
     *
     * | Permission        | Description                         |
     * |-------------------|-------------------------------------|
     * | Use               | Use service endpoint                |
     * | Administer        | Full control over service endpoints |
     * | Create            | Create service endpoints            |
     * | ViewAuthorization | View authorizations                 |
     * | ViewEndpoint      | View service endpoint properties    |
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string}>;
    /**
     * The **group** principal to assign the permissions.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the service endpoint to assign the permissions.
     */
    public readonly serviceendpointId!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceendpointPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceendpointPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceendpointPermissionsArgs | ServiceendpointPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceendpointPermissionsState | undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
            resourceInputs["serviceendpointId"] = state ? state.serviceendpointId : undefined;
        } else {
            const args = argsOrState as ServiceendpointPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replace"] = args ? args.replace : undefined;
            resourceInputs["serviceendpointId"] = args ? args.serviceendpointId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceendpointPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceendpointPermissions resources.
 */
export interface ServiceendpointPermissionsState {
    /**
     * the permissions to assign. The following permissions are available.
     *
     * | Permission        | Description                         |
     * |-------------------|-------------------------------------|
     * | Use               | Use service endpoint                |
     * | Administer        | Full control over service endpoints |
     * | Create            | Create service endpoints            |
     * | ViewAuthorization | View authorizations                 |
     * | ViewEndpoint      | View service endpoint properties    |
     */
    permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     */
    replace?: pulumi.Input<boolean>;
    /**
     * The id of the service endpoint to assign the permissions.
     */
    serviceendpointId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceendpointPermissions resource.
 */
export interface ServiceendpointPermissionsArgs {
    /**
     * the permissions to assign. The following permissions are available.
     *
     * | Permission        | Description                         |
     * |-------------------|-------------------------------------|
     * | Use               | Use service endpoint                |
     * | Administer        | Full control over service endpoints |
     * | Create            | Create service endpoints            |
     * | ViewAuthorization | View authorizations                 |
     * | ViewEndpoint      | View service endpoint properties    |
     */
    permissions: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`
     */
    replace?: pulumi.Input<boolean>;
    /**
     * The id of the service endpoint to assign the permissions.
     */
    serviceendpointId?: pulumi.Input<string>;
}
