// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages extension within Azure DevOps organization.
 *
 * ## Example Usage
 *
 * ### Install Extension
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Extension("example", {
 *     extensionId: "extension ID",
 *     publisherId: "publisher ID",
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Azure DevOps Service REST API 7.1 - Extension Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/extensionmanagement/installed-extensions?view=azure-devops-rest-7.1)
 *
 * ## Import
 *
 * Azure DevOps Extension can be imported using the publisher ID and extension ID:
 *
 * ```sh
 * $ pulumi import azuredevops:index/extension:Extension example publisherId/extensionId
 * ```
 */
export class Extension extends pulumi.CustomResource {
    /**
     * Get an existing Extension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionState, opts?: pulumi.CustomResourceOptions): Extension {
        return new Extension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/extension:Extension';

    /**
     * Returns true if the given object is an instance of Extension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extension.__pulumiType;
    }

    /**
     * Whether to disable the extension.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * The publisher ID of the extension.
     */
    public readonly extensionId!: pulumi.Output<string>;
    /**
     * The name of the extension.
     */
    public /*out*/ readonly extensionName!: pulumi.Output<string>;
    /**
     * The extension ID of the extension.
     */
    public readonly publisherId!: pulumi.Output<string>;
    /**
     * The name of the publisher.
     */
    public /*out*/ readonly publisherName!: pulumi.Output<string>;
    /**
     * List of all oauth scopes required by this extension.
     */
    public /*out*/ readonly scopes!: pulumi.Output<string[]>;
    /**
     * The version of the extension.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Extension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionArgs | ExtensionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtensionState | undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["extensionId"] = state ? state.extensionId : undefined;
            resourceInputs["extensionName"] = state ? state.extensionName : undefined;
            resourceInputs["publisherId"] = state ? state.publisherId : undefined;
            resourceInputs["publisherName"] = state ? state.publisherName : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ExtensionArgs | undefined;
            if ((!args || args.extensionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extensionId'");
            }
            if ((!args || args.publisherId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publisherId'");
            }
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["extensionId"] = args ? args.extensionId : undefined;
            resourceInputs["publisherId"] = args ? args.publisherId : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["extensionName"] = undefined /*out*/;
            resourceInputs["publisherName"] = undefined /*out*/;
            resourceInputs["scopes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Extension.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Extension resources.
 */
export interface ExtensionState {
    /**
     * Whether to disable the extension.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The publisher ID of the extension.
     */
    extensionId?: pulumi.Input<string>;
    /**
     * The name of the extension.
     */
    extensionName?: pulumi.Input<string>;
    /**
     * The extension ID of the extension.
     */
    publisherId?: pulumi.Input<string>;
    /**
     * The name of the publisher.
     */
    publisherName?: pulumi.Input<string>;
    /**
     * List of all oauth scopes required by this extension.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The version of the extension.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Extension resource.
 */
export interface ExtensionArgs {
    /**
     * Whether to disable the extension.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The publisher ID of the extension.
     */
    extensionId: pulumi.Input<string>;
    /**
     * The extension ID of the extension.
     */
    publisherId: pulumi.Input<string>;
    /**
     * The version of the extension.
     */
    version?: pulumi.Input<string>;
}
