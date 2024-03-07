// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Approval Check.
 *
 * ## Example Usage
 *
 * ### Protect an environment
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const exampleProject = new azuredevops.Project("exampleProject", {});
 * const exampleEnvironment = new azuredevops.Environment("exampleEnvironment", {projectId: exampleProject.id});
 * const exampleGroup = new azuredevops.Group("exampleGroup", {displayName: "some-azdo-group"});
 * const exampleCheckApproval = new azuredevops.CheckApproval("exampleCheckApproval", {
 *     projectId: exampleProject.id,
 *     targetResourceId: exampleEnvironment.id,
 *     targetResourceType: "environment",
 *     requesterCanApprove: true,
 *     approvers: [exampleGroup.originId],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Importing this resource is not supported.
 */
export class CheckApproval extends pulumi.CustomResource {
    /**
     * Get an existing CheckApproval resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CheckApprovalState, opts?: pulumi.CustomResourceOptions): CheckApproval {
        return new CheckApproval(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/checkApproval:CheckApproval';

    /**
     * Returns true if the given object is an instance of CheckApproval.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CheckApproval {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CheckApproval.__pulumiType;
    }

    /**
     * Specifies a list of approver IDs.
     */
    public readonly approvers!: pulumi.Output<string[]>;
    /**
     * The instructions for the approvers.
     */
    public readonly instructions!: pulumi.Output<string | undefined>;
    /**
     * The minimum number of approvers. This property is applicable when there is more than 1 approver.
     */
    public readonly minimumRequiredApprovers!: pulumi.Output<number | undefined>;
    /**
     * The project ID. Changing this forces a new Approval Check to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Can the requestor approve? Defaults to `false`.
     */
    public readonly requesterCanApprove!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
     */
    public readonly targetResourceId!: pulumi.Output<string>;
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
     */
    public readonly targetResourceType!: pulumi.Output<string>;
    /**
     * The timeout in minutes for the approval.  Defaults to `43200`.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a CheckApproval resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CheckApprovalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CheckApprovalArgs | CheckApprovalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CheckApprovalState | undefined;
            resourceInputs["approvers"] = state ? state.approvers : undefined;
            resourceInputs["instructions"] = state ? state.instructions : undefined;
            resourceInputs["minimumRequiredApprovers"] = state ? state.minimumRequiredApprovers : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["requesterCanApprove"] = state ? state.requesterCanApprove : undefined;
            resourceInputs["targetResourceId"] = state ? state.targetResourceId : undefined;
            resourceInputs["targetResourceType"] = state ? state.targetResourceType : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as CheckApprovalArgs | undefined;
            if ((!args || args.approvers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'approvers'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.targetResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceId'");
            }
            if ((!args || args.targetResourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceType'");
            }
            resourceInputs["approvers"] = args ? args.approvers : undefined;
            resourceInputs["instructions"] = args ? args.instructions : undefined;
            resourceInputs["minimumRequiredApprovers"] = args ? args.minimumRequiredApprovers : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["requesterCanApprove"] = args ? args.requesterCanApprove : undefined;
            resourceInputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            resourceInputs["targetResourceType"] = args ? args.targetResourceType : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CheckApproval.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CheckApproval resources.
 */
export interface CheckApprovalState {
    /**
     * Specifies a list of approver IDs.
     */
    approvers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instructions for the approvers.
     */
    instructions?: pulumi.Input<string>;
    /**
     * The minimum number of approvers. This property is applicable when there is more than 1 approver.
     */
    minimumRequiredApprovers?: pulumi.Input<number>;
    /**
     * The project ID. Changing this forces a new Approval Check to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Can the requestor approve? Defaults to `false`.
     */
    requesterCanApprove?: pulumi.Input<boolean>;
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
     */
    targetResourceId?: pulumi.Input<string>;
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
     */
    targetResourceType?: pulumi.Input<string>;
    /**
     * The timeout in minutes for the approval.  Defaults to `43200`.
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CheckApproval resource.
 */
export interface CheckApprovalArgs {
    /**
     * Specifies a list of approver IDs.
     */
    approvers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instructions for the approvers.
     */
    instructions?: pulumi.Input<string>;
    /**
     * The minimum number of approvers. This property is applicable when there is more than 1 approver.
     */
    minimumRequiredApprovers?: pulumi.Input<number>;
    /**
     * The project ID. Changing this forces a new Approval Check to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * Can the requestor approve? Defaults to `false`.
     */
    requesterCanApprove?: pulumi.Input<boolean>;
    /**
     * The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
     */
    targetResourceId: pulumi.Input<string>;
    /**
     * The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
     */
    targetResourceType: pulumi.Input<string>;
    /**
     * The timeout in minutes for the approval.  Defaults to `43200`.
     */
    timeout?: pulumi.Input<number>;
}
