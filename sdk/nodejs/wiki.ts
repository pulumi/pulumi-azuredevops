// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Wikis within Azure DevOps project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     description: "Managed by Pulumi",
 * });
 * const exampleGit = new azuredevops.Git("example", {
 *     projectId: example.id,
 *     name: "Example Repository",
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleWiki = new azuredevops.Wiki("example", {
 *     name: "Example project wiki ",
 *     projectId: example.id,
 *     type: "projectWiki",
 * });
 * const example2 = new azuredevops.Wiki("example2", {
 *     name: "Example wiki in repository",
 *     projectId: example.id,
 *     repositoryId: exampleGit.id,
 *     version: "main",
 *     type: "codeWiki",
 *     mappedpath: "/",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.1 - Wiki ](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/wikis?view=azure-devops-rest-7.1)
 *
 * ## Import
 *
 * Azure DevOps Wiki can be imported using the `id`
 *
 * ```sh
 * $ pulumi import azuredevops:index/wiki:Wiki wiki 00000000-0000-0000-0000-000000000000
 * ```
 */
export class Wiki extends pulumi.CustomResource {
    /**
     * Get an existing Wiki resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WikiState, opts?: pulumi.CustomResourceOptions): Wiki {
        return new Wiki(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/wiki:Wiki';

    /**
     * Returns true if the given object is an instance of Wiki.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wiki {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wiki.__pulumiType;
    }

    public readonly mappedPath!: pulumi.Output<string>;
    /**
     * The name of the Wiki.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Project.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * The remote web url to the wiki.
     */
    public /*out*/ readonly remoteUrl!: pulumi.Output<string>;
    /**
     * The ID of the repository.
     */
    public readonly repositoryId!: pulumi.Output<string>;
    /**
     * The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The REST url for this wiki.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Version of the wiki.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Wiki resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WikiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WikiArgs | WikiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WikiState | undefined;
            resourceInputs["mappedPath"] = state ? state.mappedPath : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["remoteUrl"] = state ? state.remoteUrl : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as WikiArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["mappedPath"] = args ? args.mappedPath : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["remoteUrl"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Wiki.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wiki resources.
 */
export interface WikiState {
    mappedPath?: pulumi.Input<string>;
    /**
     * The name of the Wiki.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The remote web url to the wiki.
     */
    remoteUrl?: pulumi.Input<string>;
    /**
     * The ID of the repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
     */
    type?: pulumi.Input<string>;
    /**
     * The REST url for this wiki.
     */
    url?: pulumi.Input<string>;
    /**
     * Version of the wiki.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Wiki resource.
 */
export interface WikiArgs {
    mappedPath?: pulumi.Input<string>;
    /**
     * The name of the Wiki.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
     */
    type: pulumi.Input<string>;
    /**
     * Version of the wiki.
     */
    version?: pulumi.Input<string>;
}
