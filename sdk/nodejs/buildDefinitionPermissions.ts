// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages permissions for a Build Definition
 *
 * > **Note** Permissions can be assigned to group principals and not to single user principals.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     workItemTemplate: "Agile",
 *     versionControl: "Git",
 *     visibility: "private",
 *     description: "Managed by Terraform",
 * });
 * const project-readers = azuredevops.getGroupOutput({
 *     projectId: project.id,
 *     name: "Readers",
 * });
 * const repository = new azuredevops.Git("repository", {
 *     projectId: project.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const build = new azuredevops.BuildDefinition("build", {
 *     projectId: project.id,
 *     path: "\\ExampleFolder",
 *     ciTrigger: {
 *         useYaml: true,
 *     },
 *     repository: {
 *         repoType: "TfsGit",
 *         repoId: repository.id,
 *         branchName: repository.defaultBranch,
 *         ymlPath: "azure-pipelines.yml",
 *     },
 * });
 * const permissions = new azuredevops.BuildDefinitionPermissions("permissions", {
 *     projectId: project.id,
 *     principal: project_readers.apply(project_readers => project_readers.id),
 *     buildDefinitionId: build.id,
 *     permissions: {
 *         ViewBuilds: "Allow",
 *         EditBuildQuality: "Deny",
 *         DeleteBuilds: "Deny",
 *         StopBuilds: "Allow",
 *     },
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-5.1)
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
 *
 * ## Import
 *
 * The resource does not support import.
 */
export class BuildDefinitionPermissions extends pulumi.CustomResource {
    /**
     * Get an existing BuildDefinitionPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BuildDefinitionPermissionsState, opts?: pulumi.CustomResourceOptions): BuildDefinitionPermissions {
        return new BuildDefinitionPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/buildDefinitionPermissions:BuildDefinitionPermissions';

    /**
     * Returns true if the given object is an instance of BuildDefinitionPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BuildDefinitionPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BuildDefinitionPermissions.__pulumiType;
    }

    /**
     * The id of the build definition to assign the permissions.
     */
    public readonly buildDefinitionId!: pulumi.Output<string>;
    /**
     * the permissions to assign. The following permissions are available.
     */
    public readonly permissions!: pulumi.Output<{[key: string]: string}>;
    /**
     * The **group** principal to assign the permissions.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     */
    public readonly replace!: pulumi.Output<boolean | undefined>;

    /**
     * Create a BuildDefinitionPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BuildDefinitionPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BuildDefinitionPermissionsArgs | BuildDefinitionPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BuildDefinitionPermissionsState | undefined;
            resourceInputs["buildDefinitionId"] = state ? state.buildDefinitionId : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["replace"] = state ? state.replace : undefined;
        } else {
            const args = argsOrState as BuildDefinitionPermissionsArgs | undefined;
            if ((!args || args.buildDefinitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buildDefinitionId'");
            }
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["buildDefinitionId"] = args ? args.buildDefinitionId : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["replace"] = args ? args.replace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BuildDefinitionPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BuildDefinitionPermissions resources.
 */
export interface BuildDefinitionPermissionsState {
    /**
     * The id of the build definition to assign the permissions.
     */
    buildDefinitionId?: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal?: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     */
    replace?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a BuildDefinitionPermissions resource.
 */
export interface BuildDefinitionPermissionsArgs {
    /**
     * The id of the build definition to assign the permissions.
     */
    buildDefinitionId: pulumi.Input<string>;
    /**
     * the permissions to assign. The following permissions are available.
     */
    permissions: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The **group** principal to assign the permissions.
     */
    principal: pulumi.Input<string>;
    /**
     * The ID of the project to assign the permissions.
     */
    projectId: pulumi.Input<string>;
    /**
     * Replace (`true`) or merge (`false`) the permissions. Default: `true`.
     */
    replace?: pulumi.Input<boolean>;
}
