// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Build Folder.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     name: "Example Project",
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 * });
 * const exampleBuildFolder = new azuredevops.BuildFolder("example", {
 *     projectId: example.id,
 *     path: "\\ExampleFolder",
 *     description: "ExampleFolder description",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Build Folders can be imported using the `project name/path` or `project id/path`, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:index/buildFolder:BuildFolder example "Example Project/\\ExampleFolder"
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import azuredevops:index/buildFolder:BuildFolder example 00000000-0000-0000-0000-000000000000/\\ExampleFolder
 * ```
 */
export class BuildFolder extends pulumi.CustomResource {
    /**
     * Get an existing BuildFolder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BuildFolderState, opts?: pulumi.CustomResourceOptions): BuildFolder {
        return new BuildFolder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/buildFolder:BuildFolder';

    /**
     * Returns true if the given object is an instance of BuildFolder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BuildFolder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BuildFolder.__pulumiType;
    }

    /**
     * Folder Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The folder path.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * The ID of the project in which the folder will be created.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a BuildFolder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BuildFolderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BuildFolderArgs | BuildFolderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BuildFolderState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as BuildFolderArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BuildFolder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BuildFolder resources.
 */
export interface BuildFolderState {
    /**
     * Folder Description.
     */
    description?: pulumi.Input<string>;
    /**
     * The folder path.
     */
    path?: pulumi.Input<string>;
    /**
     * The ID of the project in which the folder will be created.
     */
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BuildFolder resource.
 */
export interface BuildFolderArgs {
    /**
     * Folder Description.
     */
    description?: pulumi.Input<string>;
    /**
     * The folder path.
     */
    path: pulumi.Input<string>;
    /**
     * The ID of the project in which the folder will be created.
     */
    projectId: pulumi.Input<string>;
}
