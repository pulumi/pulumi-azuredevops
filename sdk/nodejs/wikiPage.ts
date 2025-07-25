// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Wiki pages within Azure DevOps project.
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
 * const exampleWiki = new azuredevops.Wiki("example", {
 *     projectId: example.id,
 *     name: "Example project wiki ",
 *     type: "projectWiki",
 * });
 * const exampleWikiPage = new azuredevops.WikiPage("example", {
 *     projectId: example.id,
 *     wikiId: exampleWiki.id,
 *     path: "/page",
 *     content: "content",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.1 - Wiki Page](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/pages?view=azure-devops-rest-7.1)
 */
export class WikiPage extends pulumi.CustomResource {
    /**
     * Get an existing WikiPage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WikiPageState, opts?: pulumi.CustomResourceOptions): WikiPage {
        return new WikiPage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/wikiPage:WikiPage';

    /**
     * Returns true if the given object is an instance of WikiPage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WikiPage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WikiPage.__pulumiType;
    }

    /**
     * The content of the wiki page.
     */
    public readonly content!: pulumi.Output<string>;
    public readonly etag!: pulumi.Output<string>;
    /**
     * The path of the wiki page.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * The ID of the Project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The ID of the Wiki.
     */
    public readonly wikiId!: pulumi.Output<string>;

    /**
     * Create a WikiPage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WikiPageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WikiPageArgs | WikiPageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WikiPageState | undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["wikiId"] = state ? state.wikiId : undefined;
        } else {
            const args = argsOrState as WikiPageArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.wikiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'wikiId'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["wikiId"] = args ? args.wikiId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WikiPage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WikiPage resources.
 */
export interface WikiPageState {
    /**
     * The content of the wiki page.
     */
    content?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The path of the wiki page.
     */
    path?: pulumi.Input<string>;
    /**
     * The ID of the Project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the Wiki.
     */
    wikiId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WikiPage resource.
 */
export interface WikiPageArgs {
    /**
     * The content of the wiki page.
     */
    content: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The path of the wiki page.
     */
    path: pulumi.Input<string>;
    /**
     * The ID of the Project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the Wiki.
     */
    wikiId: pulumi.Input<string>;
}
