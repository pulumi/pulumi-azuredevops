// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a git repository within Azure DevOps.
 *
 * ## Example Usage
 * ### Create Git repository
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const project = new azuredevops.Project("project", {
 *     projectName: "Sample Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const repo = new azuredevops.Git("repo", {
 *     projectId: project.id,
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * ```
 * ### Create Fork of another Azure DevOps Git repository
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const repo = new azuredevops.Git("repo", {
 *     projectId: azuredevops_project.project.id,
 *     parentId: azuredevops_git_repository.parent.id,
 * });
 * ```
 * ## Relevant Links
 *
 * * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories?view=azure-devops-rest-5.1)
 *
 * @deprecated azuredevops.repository.Git has been deprecated in favor of azuredevops.Git
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
        pulumi.log.warn("Git is deprecated: azuredevops.repository.Git has been deprecated in favor of azuredevops.Git")
        return new Git(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Repository/git:Git';

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
     * The ref of the default branch.
     */
    public readonly defaultBranch!: pulumi.Output<string>;
    /**
     * An `initialization` block as documented below.
     */
    public readonly initialization!: pulumi.Output<outputs.Repository.GitInitialization | undefined>;
    /**
     * True if the repository was created as a fork.
     */
    public /*out*/ readonly isFork!: pulumi.Output<boolean>;
    /**
     * The name of the git repository.
     */
    public readonly name!: pulumi.Output<string>;
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
    /** @deprecated azuredevops.repository.Git has been deprecated in favor of azuredevops.Git */
    constructor(name: string, args: GitArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.repository.Git has been deprecated in favor of azuredevops.Git */
    constructor(name: string, argsOrState?: GitArgs | GitState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Git is deprecated: azuredevops.repository.Git has been deprecated in favor of azuredevops.Git")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GitState | undefined;
            inputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            inputs["initialization"] = state ? state.initialization : undefined;
            inputs["isFork"] = state ? state.isFork : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parentRepositoryId"] = state ? state.parentRepositoryId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["remoteUrl"] = state ? state.remoteUrl : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["sshUrl"] = state ? state.sshUrl : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["webUrl"] = state ? state.webUrl : undefined;
        } else {
            const args = argsOrState as GitArgs | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            inputs["initialization"] = args ? args.initialization : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parentRepositoryId"] = args ? args.parentRepositoryId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["isFork"] = undefined /*out*/;
            inputs["remoteUrl"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["sshUrl"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
            inputs["webUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Git.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Git resources.
 */
export interface GitState {
    /**
     * The ref of the default branch.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * An `initialization` block as documented below.
     */
    readonly initialization?: pulumi.Input<inputs.Repository.GitInitialization>;
    /**
     * True if the repository was created as a fork.
     */
    readonly isFork?: pulumi.Input<boolean>;
    /**
     * The name of the git repository.
     */
    readonly name?: pulumi.Input<string>;
    readonly parentRepositoryId?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Git HTTPS URL of the repository
     */
    readonly remoteUrl?: pulumi.Input<string>;
    /**
     * Size in bytes.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * Git SSH URL of the repository.
     */
    readonly sshUrl?: pulumi.Input<string>;
    /**
     * REST API URL of the repository.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Web link to the repository.
     */
    readonly webUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Git resource.
 */
export interface GitArgs {
    /**
     * The ref of the default branch.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * An `initialization` block as documented below.
     */
    readonly initialization?: pulumi.Input<inputs.Repository.GitInitialization>;
    /**
     * The name of the git repository.
     */
    readonly name?: pulumi.Input<string>;
    readonly parentRepositoryId?: pulumi.Input<string>;
    /**
     * The project ID or project name.
     */
    readonly projectId: pulumi.Input<string>;
}