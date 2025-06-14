// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Azure DevOps Repositories can be imported using the repo name or by the repo Guid e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/git:Git example projectName/repoName
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import azuredevops:index/git:Git example projectName/00000000-0000-0000-0000-000000000000
 * ```
 */
export class Git extends pulumi.CustomResource {
    /**
     * Get an existing Git resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GitState, opts?: pulumi.CustomResourceOptions): Git {
        return new Git(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/git:Git';

    /**
     * Returns true if the given object is an instance of Git.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Git {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Git.__pulumiType;
    }

    /**
     * The ref of the default branch. Will be used as the branch name for initialized repositories.
     */
    public readonly defaultBranch!: pulumi.Output<string>;
    /**
     * The ability to disable or enable the repository. Defaults to `false`.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * A `initialization` block as documented below.
     */
    public readonly initialization!: pulumi.Output<outputs.GitInitialization>;
    /**
     * True if the repository was created as a fork.
     */
    public /*out*/ readonly isFork!: pulumi.Output<boolean>;
    /**
     * The name of the git repository.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of a Git project from which a fork is to be created.
     */
    public readonly parentRepositoryId!: pulumi.Output<string | undefined>;
    /**
     * The project ID or project name.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Git HTTPS URL of the repository
     */
    public /*out*/ readonly remoteUrl!: pulumi.Output<string>;
    /**
     * Size in bytes.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Git SSH URL of the repository.
     */
    public /*out*/ readonly sshUrl!: pulumi.Output<string>;
    /**
     * REST API URL of the repository.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Web link to the repository.
     */
    public /*out*/ readonly webUrl!: pulumi.Output<string>;

    /**
     * Create a Git resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GitArgs | GitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GitState | undefined;
            resourceInputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["initialization"] = state ? state.initialization : undefined;
            resourceInputs["isFork"] = state ? state.isFork : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentRepositoryId"] = state ? state.parentRepositoryId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["remoteUrl"] = state ? state.remoteUrl : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["sshUrl"] = state ? state.sshUrl : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["webUrl"] = state ? state.webUrl : undefined;
        } else {
            const args = argsOrState as GitArgs | undefined;
            if ((!args || args.initialization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'initialization'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["initialization"] = args ? args.initialization : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentRepositoryId"] = args ? args.parentRepositoryId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["isFork"] = undefined /*out*/;
            resourceInputs["remoteUrl"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["sshUrl"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["webUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Git.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Git resources.
 */
export interface GitState {
    /**
     * The ref of the default branch. Will be used as the branch name for initialized repositories.
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * The ability to disable or enable the repository. Defaults to `false`.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * A `initialization` block as documented below.
     */
    initialization?: pulumi.Input<inputs.GitInitialization>;
    /**
     * True if the repository was created as a fork.
     */
    isFork?: pulumi.Input<boolean>;
    /**
     * The name of the git repository.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of a Git project from which a fork is to be created.
     */
    parentRepositoryId?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Git HTTPS URL of the repository
     */
    remoteUrl?: pulumi.Input<string>;
    /**
     * Size in bytes.
     */
    size?: pulumi.Input<number>;
    /**
     * Git SSH URL of the repository.
     */
    sshUrl?: pulumi.Input<string>;
    /**
     * REST API URL of the repository.
     */
    url?: pulumi.Input<string>;
    /**
     * Web link to the repository.
     */
    webUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Git resource.
 */
export interface GitArgs {
    /**
     * The ref of the default branch. Will be used as the branch name for initialized repositories.
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * The ability to disable or enable the repository. Defaults to `false`.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * A `initialization` block as documented below.
     */
    initialization: pulumi.Input<inputs.GitInitialization>;
    /**
     * The name of the git repository.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of a Git project from which a fork is to be created.
     */
    parentRepositoryId?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    projectId: pulumi.Input<string>;
}
