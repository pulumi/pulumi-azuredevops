// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a business hours check on a resource within Azure DevOps.
 *
 * ## Example Usage
 *
 * ### Protect a service connection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const exampleServiceEndpointGeneric = new azuredevops.ServiceEndpointGeneric("example", {
 *     projectId: example.id,
 *     serverUrl: "https://some-server.example.com",
 *     username: "username",
 *     password: "password",
 *     serviceEndpointName: "Example Generic",
 *     description: "Managed by Pulumi",
 * });
 * const exampleCheckBusinessHours = new azuredevops.CheckBusinessHours("example", {
 *     projectId: example.id,
 *     displayName: "Managed by Pulumi",
 *     targetResourceId: exampleServiceEndpointGeneric.id,
 *     targetResourceType: "endpoint",
 *     startTime: "07:00",
 *     endTime: "15:30",
 *     timeZone: "UTC",
 *     monday: true,
 *     tuesday: true,
 *     timeout: 1440,
 * });
 * ```
 *
 * ### Protect an environment
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const exampleEnvironment = new azuredevops.Environment("example", {
 *     projectId: example.id,
 *     name: "Example Environment",
 * });
 * const exampleCheckBusinessHours = new azuredevops.CheckBusinessHours("example", {
 *     projectId: example.id,
 *     displayName: "Managed by Pulumi",
 *     targetResourceId: exampleEnvironment.id,
 *     targetResourceType: "environment",
 *     startTime: "07:00",
 *     endTime: "15:30",
 *     timeZone: "UTC",
 *     monday: true,
 *     tuesday: true,
 * });
 * ```
 *
 * ### Protect an agent queue
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const examplePool = new azuredevops.Pool("example", {name: "example-pool"});
 * const exampleQueue = new azuredevops.Queue("example", {
 *     projectId: example.id,
 *     agentPoolId: examplePool.id,
 * });
 * const exampleCheckBusinessHours = new azuredevops.CheckBusinessHours("example", {
 *     projectId: example.id,
 *     displayName: "Managed by Pulumi",
 *     targetResourceId: exampleQueue.id,
 *     targetResourceType: "queue",
 *     startTime: "07:00",
 *     endTime: "15:30",
 *     timeZone: "UTC",
 *     monday: true,
 *     tuesday: true,
 * });
 * ```
 *
 * ### Protect a repository
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const exampleGit = new azuredevops.Git("example", {
 *     projectId: example.id,
 *     name: "Example Empty Git Repository",
 *     initialization: {
 *         initType: "Clean",
 *     },
 * });
 * const exampleCheckBusinessHours = new azuredevops.CheckBusinessHours("example", {
 *     projectId: example.id,
 *     displayName: "Managed by Pulumi",
 *     targetResourceId: pulumi.interpolate`${example.id}.${exampleGit.id}`,
 *     targetResourceType: "repository",
 *     startTime: "07:00",
 *     endTime: "15:30",
 *     timeZone: "UTC",
 *     monday: true,
 *     tuesday: true,
 * });
 * ```
 *
 * ### Protect a variable group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuredevops from "@pulumi/azuredevops";
 *
 * const example = new azuredevops.Project("example", {name: "Example Project"});
 * const exampleVariableGroup = new azuredevops.VariableGroup("example", {
 *     projectId: example.id,
 *     name: "Example Variable Group",
 *     description: "Example Variable Group Description",
 *     allowAccess: true,
 *     variables: [
 *         {
 *             name: "key1",
 *             value: "val1",
 *         },
 *         {
 *             name: "key2",
 *             secretValue: "val2",
 *             isSecret: true,
 *         },
 *     ],
 * });
 * const exampleCheckBusinessHours = new azuredevops.CheckBusinessHours("example", {
 *     projectId: example.id,
 *     displayName: "Managed by Pulumi",
 *     targetResourceId: exampleVariableGroup.id,
 *     targetResourceType: "variablegroup",
 *     startTime: "07:00",
 *     endTime: "15:30",
 *     timeZone: "UTC",
 *     monday: true,
 *     tuesday: true,
 * });
 * ```
 *
 * ## Relevant Links
 *
 * - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&tabs=check-pass)
 *
 * ## Supported Time Zones
 *
 * - AUS Central Standard Time
 * - AUS Eastern Standard Time
 * - Afghanistan Standard Time
 * - Alaskan Standard Time
 * - Aleutian Standard Time
 * - Altai Standard Time
 * - Arab Standard Time
 * - Arabian Standard Time
 * - Arabic Standard Time
 * - Argentina Standard Time
 * - Astrakhan Standard Time
 * - Atlantic Standard Time
 * - Aus Central W. Standard Time
 * - Azerbaijan Standard Time
 * - Azores Standard Time
 * - Bahia Standard Time
 * - Bangladesh Standard Time
 * - Belarus Standard Time
 * - Bougainville Standard Time
 * - Canada Central Standard Time
 * - Cape Verde Standard Time
 * - Caucasus Standard Time
 * - Cen. Australia Standard Time
 * - Central America Standard Time
 * - Central Asia Standard Time
 * - Central Brazilian Standard Time
 * - Central Europe Standard Time
 * - Central European Standard Time
 * - Central Pacific Standard Time
 * - Central Standard Time (Mexico)
 * - Central Standard Time
 * - Chatham Islands Standard Time
 * - China Standard Time
 * - Cuba Standard Time
 * - Dateline Standard Time
 * - E. Africa Standard Time
 * - E. Australia Standard Time
 * - E. Europe Standard Time
 * - E. South America Standard Time
 * - Easter Island Standard Time
 * - Eastern Standard Time (Mexico)
 * - Eastern Standard Time
 * - Egypt Standard Time
 * - Ekaterinburg Standard Time
 * - FLE Standard Time
 * - Fiji Standard Time
 * - GMT Standard Time
 * - GTB Standard Time
 * - Georgian Standard Time
 * - Greenland Standard Time
 * - Greenwich Standard Time
 * - Haiti Standard Time
 * - Hawaiian Standard Time
 * - India Standard Time
 * - Iran Standard Time
 * - Israel Standard Time
 * - Jordan Standard Time
 * - Kaliningrad Standard Time
 * - Kamchatka Standard Time
 * - Korea Standard Time
 * - Libya Standard Time
 * - Line Islands Standard Time
 * - Lord Howe Standard Time
 * - Magadan Standard Time
 * - Magallanes Standard Time
 * - Marquesas Standard Time
 * - Mauritius Standard Time
 * - Mid-Atlantic Standard Time
 * - Middle East Standard Time
 * - Montevideo Standard Time
 * - Morocco Standard Time
 * - Mountain Standard Time (Mexico)
 * - Mountain Standard Time
 * - Myanmar Standard Time
 * - N. Central Asia Standard Time
 * - Namibia Standard Time
 * - Nepal Standard Time
 * - New Zealand Standard Time
 * - Newfoundland Standard Time
 * - Norfolk Standard Time
 * - North Asia East Standard Time
 * - North Asia Standard Time
 * - North Korea Standard Time
 * - Omsk Standard Time
 * - Pacific SA Standard Time
 * - Pacific Standard Time (Mexico)
 * - Pacific Standard Time
 * - Pakistan Standard Time
 * - Paraguay Standard Time
 * - Qyzylorda Standard Time
 * - Romance Standard Time
 * - Russia Time Zone 10
 * - Russia Time Zone 11
 * - Russia Time Zone 3
 * - Russian Standard Time
 * - SA Eastern Standard Time
 * - SA Pacific Standard Time
 * - SA Western Standard Time
 * - SE Asia Standard Time
 * - Saint Pierre Standard Time
 * - Sakhalin Standard Time
 * - Samoa Standard Time
 * - Sao Tome Standard Time
 * - Saratov Standard Time
 * - Singapore Standard Time
 * - South Africa Standard Time
 * - South Sudan Standard Time
 * - Sri Lanka Standard Time
 * - Sudan Standard Time
 * - Syria Standard Time
 * - Taipei Standard Time
 * - Tasmania Standard Time
 * - Tocantins Standard Time
 * - Tokyo Standard Time
 * - Tomsk Standard Time
 * - Tonga Standard Time
 * - Transbaikal Standard Time
 * - Turkey Standard Time
 * - Turks And Caicos Standard Time
 * - US Eastern Standard Time
 * - US Mountain Standard Time
 * - UTC
 * - UTC+12
 * - UTC+13
 * - UTC-02
 * - UTC-08
 * - UTC-09
 * - UTC-11
 * - Ulaanbaatar Standard Time
 * - Venezuela Standard Time
 * - Vladivostok Standard Time
 * - Volgograd Standard Time
 * - W. Australia Standard Time
 * - W. Central Africa Standard Time
 * - W. Europe Standard Time
 * - W. Mongolia Standard Time
 * - West Asia Standard Time
 * - West Bank Standard Time
 * - West Pacific Standard Time
 * - Yakutsk Standard Time
 * - Yukon Standard Time
 *
 * ## Import
 *
 * Importing this resource is not supported.
 */
export class CheckBusinessHours extends pulumi.CustomResource {
    /**
     * Get an existing CheckBusinessHours resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CheckBusinessHoursState, opts?: pulumi.CustomResourceOptions): CheckBusinessHours {
        return new CheckBusinessHours(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuredevops:index/checkBusinessHours:CheckBusinessHours';

    /**
     * Returns true if the given object is an instance of CheckBusinessHours.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CheckBusinessHours {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CheckBusinessHours.__pulumiType;
    }

    /**
     * The name of the business hours check displayed in the web UI.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * This check will pass on Fridays. Defaults to `false`.
     */
    public readonly friday!: pulumi.Output<boolean | undefined>;
    /**
     * This check will pass on Mondays. Defaults to `false`.
     */
    public readonly monday!: pulumi.Output<boolean | undefined>;
    /**
     * The project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * This check will pass on Saturdays. Defaults to `false`.
     */
    public readonly saturday!: pulumi.Output<boolean | undefined>;
    /**
     * The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * This check will pass on Sundays. Defaults to `false`.
     */
    public readonly sunday!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the resource being protected by the check.
     */
    public readonly targetResourceId!: pulumi.Output<string>;
    /**
     * The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     */
    public readonly targetResourceType!: pulumi.Output<string>;
    /**
     * This check will pass on Thursdays. Defaults to `false`.
     */
    public readonly thursday!: pulumi.Output<boolean | undefined>;
    /**
     * The time zone this check will be evaluated in. See below for supported values.
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * The timeout in minutes for the business hours check. Defaults to `1440`.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * This check will pass on Tuesday. Defaults to `false`.
     */
    public readonly tuesday!: pulumi.Output<boolean | undefined>;
    /**
     * The version of the check.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;
    /**
     * This check will pass on Wednesdays. Defaults to `false`.
     */
    public readonly wednesday!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CheckBusinessHours resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CheckBusinessHoursArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CheckBusinessHoursArgs | CheckBusinessHoursState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CheckBusinessHoursState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["friday"] = state ? state.friday : undefined;
            resourceInputs["monday"] = state ? state.monday : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["saturday"] = state ? state.saturday : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["sunday"] = state ? state.sunday : undefined;
            resourceInputs["targetResourceId"] = state ? state.targetResourceId : undefined;
            resourceInputs["targetResourceType"] = state ? state.targetResourceType : undefined;
            resourceInputs["thursday"] = state ? state.thursday : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["tuesday"] = state ? state.tuesday : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["wednesday"] = state ? state.wednesday : undefined;
        } else {
            const args = argsOrState as CheckBusinessHoursArgs | undefined;
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            if ((!args || args.targetResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceId'");
            }
            if ((!args || args.targetResourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceType'");
            }
            if ((!args || args.timeZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeZone'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["friday"] = args ? args.friday : undefined;
            resourceInputs["monday"] = args ? args.monday : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["saturday"] = args ? args.saturday : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["sunday"] = args ? args.sunday : undefined;
            resourceInputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            resourceInputs["targetResourceType"] = args ? args.targetResourceType : undefined;
            resourceInputs["thursday"] = args ? args.thursday : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["tuesday"] = args ? args.tuesday : undefined;
            resourceInputs["wednesday"] = args ? args.wednesday : undefined;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CheckBusinessHours.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CheckBusinessHours resources.
 */
export interface CheckBusinessHoursState {
    /**
     * The name of the business hours check displayed in the web UI.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     */
    endTime?: pulumi.Input<string>;
    /**
     * This check will pass on Fridays. Defaults to `false`.
     */
    friday?: pulumi.Input<boolean>;
    /**
     * This check will pass on Mondays. Defaults to `false`.
     */
    monday?: pulumi.Input<boolean>;
    /**
     * The project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * This check will pass on Saturdays. Defaults to `false`.
     */
    saturday?: pulumi.Input<boolean>;
    /**
     * The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     */
    startTime?: pulumi.Input<string>;
    /**
     * This check will pass on Sundays. Defaults to `false`.
     */
    sunday?: pulumi.Input<boolean>;
    /**
     * The ID of the resource being protected by the check.
     */
    targetResourceId?: pulumi.Input<string>;
    /**
     * The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     */
    targetResourceType?: pulumi.Input<string>;
    /**
     * This check will pass on Thursdays. Defaults to `false`.
     */
    thursday?: pulumi.Input<boolean>;
    /**
     * The time zone this check will be evaluated in. See below for supported values.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * The timeout in minutes for the business hours check. Defaults to `1440`.
     */
    timeout?: pulumi.Input<number>;
    /**
     * This check will pass on Tuesday. Defaults to `false`.
     */
    tuesday?: pulumi.Input<boolean>;
    /**
     * The version of the check.
     */
    version?: pulumi.Input<number>;
    /**
     * This check will pass on Wednesdays. Defaults to `false`.
     */
    wednesday?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CheckBusinessHours resource.
 */
export interface CheckBusinessHoursArgs {
    /**
     * The name of the business hours check displayed in the web UI.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     */
    endTime: pulumi.Input<string>;
    /**
     * This check will pass on Fridays. Defaults to `false`.
     */
    friday?: pulumi.Input<boolean>;
    /**
     * This check will pass on Mondays. Defaults to `false`.
     */
    monday?: pulumi.Input<boolean>;
    /**
     * The project ID.
     */
    projectId: pulumi.Input<string>;
    /**
     * This check will pass on Saturdays. Defaults to `false`.
     */
    saturday?: pulumi.Input<boolean>;
    /**
     * The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.
     */
    startTime: pulumi.Input<string>;
    /**
     * This check will pass on Sundays. Defaults to `false`.
     */
    sunday?: pulumi.Input<boolean>;
    /**
     * The ID of the resource being protected by the check.
     */
    targetResourceId: pulumi.Input<string>;
    /**
     * The type of resource being protected by the check. Possible values are: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`.
     */
    targetResourceType: pulumi.Input<string>;
    /**
     * This check will pass on Thursdays. Defaults to `false`.
     */
    thursday?: pulumi.Input<boolean>;
    /**
     * The time zone this check will be evaluated in. See below for supported values.
     */
    timeZone: pulumi.Input<string>;
    /**
     * The timeout in minutes for the business hours check. Defaults to `1440`.
     */
    timeout?: pulumi.Input<number>;
    /**
     * This check will pass on Tuesday. Defaults to `false`.
     */
    tuesday?: pulumi.Input<boolean>;
    /**
     * This check will pass on Wednesdays. Defaults to `false`.
     */
    wednesday?: pulumi.Input<boolean>;
}
