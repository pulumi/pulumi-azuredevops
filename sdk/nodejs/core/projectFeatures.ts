// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages features for Azure DevOps projects
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {
 *     visibility: "private",
 *     versionControl: "Git",
 *     workItemTemplate: "Agile",
 *     description: "Managed by Terraform",
 * });
 * const example_features = new azuredevops.ProjectFeatures("example-features", {
 *     projectId: example.id,
 *     features: {
 *         testplans: "disabled",
 *         artifacts: "enabled",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Relevant Links
 *
 * No official documentation available
 *
 * ## PAT Permissions Required
 *
 * - **Project & Team**: Read, Write, & Manage
 *
 * ## Import
 *
 * Azure DevOps feature settings can be imported using the project id, e.g.
 *
 * ```sh
 * $ pulumi import azuredevops:Core/projectFeatures:ProjectFeatures example 00000000-0000-0000-0000-000000000000
 * ```
 *
 * @deprecated azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures
 */
export class ProjectFeatures extends pulumi.CustomResource {
    /**
     * Get an existing ProjectFeatures resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectFeaturesState, opts?: pulumi.CustomResourceOptions): ProjectFeatures {
        pulumi.log.warn("ProjectFeatures is deprecated: azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures")
        return new ProjectFeatures(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:Core/projectFeatures:ProjectFeatures';

    /**
     * Returns true if the given object is an instance of ProjectFeatures.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectFeatures {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectFeatures.__pulumiType;
    }

    /**
     * Defines the status (`enabled`, `disabled`) of the project features.  
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     *
     * > **NOTE:**
     * > It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * > via the `features` block by using the `azuredevops.Project` resource.
     * > However it's not possible to use both methods to manage features, since there'll be conflicts.
     */
    public readonly features!: pulumi.Output<{[key: string]: string}>;
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a ProjectFeatures resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures */
    constructor(name: string, args: ProjectFeaturesArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures */
    constructor(name: string, argsOrState?: ProjectFeaturesArgs | ProjectFeaturesState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ProjectFeatures is deprecated: azuredevops.core.ProjectFeatures has been deprecated in favor of azuredevops.ProjectFeatures")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectFeaturesState | undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as ProjectFeaturesArgs | undefined;
            if ((!args || args.features === undefined) && !opts.urn) {
                throw new Error("Missing required property 'features'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectFeatures.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectFeatures resources.
 */
export interface ProjectFeaturesState {
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.  
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     *
     * > **NOTE:**
     * > It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * > via the `features` block by using the `azuredevops.Project` resource.
     * > However it's not possible to use both methods to manage features, since there'll be conflicts.
     */
    features?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectFeatures resource.
 */
export interface ProjectFeaturesArgs {
    /**
     * Defines the status (`enabled`, `disabled`) of the project features.  
     * Valid features `boards`, `repositories`, `pipelines`, `testplans`, `artifacts`
     *
     * > **NOTE:**
     * > It's possible to define project features both within the `azuredevops.ProjectFeatures` resource and
     * > via the `features` block by using the `azuredevops.Project` resource.
     * > However it's not possible to use both methods to manage features, since there'll be conflicts.
     */
    features: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    projectId: pulumi.Input<string>;
}
