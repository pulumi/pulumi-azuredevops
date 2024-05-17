// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface BranchPolicyAutoReviewersSettings {
    /**
     * Required reviewers ids. Supports multiples user Ids.
     */
    autoReviewerIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
     */
    message?: pulumi.Input<string>;
    /**
     * Minimum number of required reviewers. Defaults to `1`.
     *
     * > **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `autoReviewerIds` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
     */
    minimumNumberOfReviewers?: pulumi.Input<number>;
    /**
     * Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
     */
    pathFilters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyAutoReviewersSettingsScope>[]>;
    /**
     * Controls whether or not the submitter's vote counts. Defaults to `false`.
     */
    submitterCanVote?: pulumi.Input<boolean>;
}

export interface BranchPolicyAutoReviewersSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `matchType` is `DefaultBranch`, this should not be defined.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match when `matchType` other than `DefaultBranch`. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
    repositoryRef?: pulumi.Input<string>;
}

export interface BranchPolicyBuildValidationSettings {
    /**
     * The ID of the build to monitor for the policy.
     */
    buildDefinitionId: pulumi.Input<number>;
    /**
     * The display name for the policy.
     */
    displayName: pulumi.Input<string>;
    /**
     * If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
     */
    filenamePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to true, the build will need to be manually queued. Defaults to `false`
     */
    manualQueueOnly?: pulumi.Input<boolean>;
    /**
     * True if the build should queue on source updates only. Defaults to `true`.
     */
    queueOnSourceUpdateOnly?: pulumi.Input<boolean>;
    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyBuildValidationSettingsScope>[]>;
    /**
     * The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
     */
    validDuration?: pulumi.Input<number>;
}

export interface BranchPolicyBuildValidationSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `matchType` is `DefaultBranch`, this should not be defined.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match when `matchType` other than `DefaultBranch`. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
    repositoryRef?: pulumi.Input<string>;
}

export interface BranchPolicyCommentResolutionSettings {
    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyCommentResolutionSettingsScope>[]>;
}

export interface BranchPolicyCommentResolutionSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `matchType` is `DefaultBranch`, this should not be defined.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match when `matchType` other than `DefaultBranch`. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
    repositoryRef?: pulumi.Input<string>;
}

export interface BranchPolicyMergeTypesSettings {
    /**
     * Allow basic merge with no fast forward. Defaults to `false`.
     */
    allowBasicNoFastForward?: pulumi.Input<boolean>;
    /**
     * Allow rebase with fast forward. Defaults to `false`.
     */
    allowRebaseAndFastForward?: pulumi.Input<boolean>;
    /**
     * Allow rebase with merge commit. Defaults to `false`.
     */
    allowRebaseWithMerge?: pulumi.Input<boolean>;
    /**
     * Allow squash merge. Defaults to `false`
     */
    allowSquash?: pulumi.Input<boolean>;
    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyMergeTypesSettingsScope>[]>;
}

export interface BranchPolicyMergeTypesSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `matchType` is `DefaultBranch`, this should not be defined.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match when `matchType` other than `DefaultBranch`. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
    repositoryRef?: pulumi.Input<string>;
}

export interface BranchPolicyMinReviewersSettings {
    /**
     * Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
     */
    allowCompletionWithRejectsOrWaits?: pulumi.Input<boolean>;
    /**
     * Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
     */
    lastPusherCannotApprove?: pulumi.Input<boolean>;
    /**
     * On last iteration require vote. Defaults to `false`.
     */
    onLastIterationRequireVote?: pulumi.Input<boolean>;
    /**
     * When new changes are pushed reset all code reviewer votes. Defaults to `false`.
     *
     * > **Note:** If `onPushResetAllVotes` is `true` then `onPushResetApprovedVotes` will be set to `true`. To enable `onPushResetApprovedVotes`, you need explicitly set `onPushResetAllVotes` `false` or not configure.
     */
    onPushResetAllVotes?: pulumi.Input<boolean>;
    /**
     * When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
     */
    onPushResetApprovedVotes?: pulumi.Input<boolean>;
    /**
     * The number of reviewers needed to approve.
     */
    reviewerCount?: pulumi.Input<number>;
    /**
     * A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyMinReviewersSettingsScope>[]>;
    /**
     * Allow requesters to approve their own changes. Defaults to `false`.
     */
    submitterCanVote?: pulumi.Input<boolean>;
}

export interface BranchPolicyMinReviewersSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `matchType` is `DefaultBranch`, this should not be defined.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match when `matchType` other than `DefaultBranch`. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
    repositoryRef?: pulumi.Input<string>;
}

export interface BranchPolicyStatusCheckSettings {
    /**
     * Policy applicability. If policy `applicability` is `default`, apply unless "Not Applicable" 
     * status is posted to the pull request. If policy `applicability` is `conditional`, policy is applied only after a status
     * is posted to the pull request.
     */
    applicability?: pulumi.Input<string>;
    /**
     * The authorized user can post the status.
     */
    authorId?: pulumi.Input<string>;
    /**
     * The display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
     */
    filenamePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The genre of the status to check (see [Microsoft Documentation](https://docs.microsoft.com/en-us/azure/devops/repos/git/pull-request-status?view=azure-devops#status-policy))
     */
    genre?: pulumi.Input<string>;
    /**
     * Reset status whenever there are new changes.
     */
    invalidateOnUpdate?: pulumi.Input<boolean>;
    /**
     * The status name to check.
     */
    name: pulumi.Input<string>;
    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined
     * at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyStatusCheckSettingsScope>[]>;
}

export interface BranchPolicyStatusCheckSettingsScope {
    matchType?: pulumi.Input<string>;
    repositoryId?: pulumi.Input<string>;
    repositoryRef?: pulumi.Input<string>;
}

export interface BranchPolicyWorkItemLinkingSettings {
    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyWorkItemLinkingSettingsScope>[]>;
}

export interface BranchPolicyWorkItemLinkingSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `matchType` is `DefaultBranch`, this should not be defined.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match when `matchType` other than `DefaultBranch`. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
    repositoryRef?: pulumi.Input<string>;
}

export interface BuildDefinitionCiTrigger {
    /**
     * Override the azure-pipeline file and use a this configuration for all builds.
     */
    override?: pulumi.Input<inputs.BuildDefinitionCiTriggerOverride>;
    /**
     * Use the azure-pipeline file for the build configuration. Defaults to `false`.
     */
    useYaml?: pulumi.Input<boolean>;
}

export interface BuildDefinitionCiTriggerOverride {
    /**
     * If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
     */
    batch?: pulumi.Input<boolean>;
    /**
     * The branches to include and exclude from the trigger.
     */
    branchFilters?: pulumi.Input<pulumi.Input<inputs.BuildDefinitionCiTriggerOverrideBranchFilter>[]>;
    /**
     * The number of max builds per branch. Defaults to `1`.
     */
    maxConcurrentBuildsPerBranch?: pulumi.Input<number>;
    /**
     * Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     */
    pathFilters?: pulumi.Input<pulumi.Input<inputs.BuildDefinitionCiTriggerOverridePathFilter>[]>;
    /**
     * How often the external repository is polled. Defaults to `0`.
     */
    pollingInterval?: pulumi.Input<number>;
    /**
     * This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     */
    pollingJobId?: pulumi.Input<string>;
}

export interface BuildDefinitionCiTriggerOverrideBranchFilter {
    /**
     * List of branch patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of branch patterns to include.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BuildDefinitionCiTriggerOverridePathFilter {
    /**
     * List of path patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of path patterns to include.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BuildDefinitionFeature {
    skipFirstRun?: pulumi.Input<boolean>;
}

export interface BuildDefinitionPullRequestTrigger {
    commentRequired?: pulumi.Input<string>;
    /**
     * Set permissions for Forked repositories.
     */
    forks: pulumi.Input<inputs.BuildDefinitionPullRequestTriggerForks>;
    initialBranch?: pulumi.Input<string>;
    /**
     * Override the azure-pipeline file and use this configuration for all builds.
     */
    override?: pulumi.Input<inputs.BuildDefinitionPullRequestTriggerOverride>;
    /**
     * Use the azure-pipeline file for the build configuration. Defaults to `false`.
     */
    useYaml?: pulumi.Input<boolean>;
}

export interface BuildDefinitionPullRequestTriggerForks {
    /**
     * Build pull requests from forks of this repository.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Make secrets available to builds of forks.
     */
    shareSecrets: pulumi.Input<boolean>;
}

export interface BuildDefinitionPullRequestTriggerOverride {
    /**
     * . Defaults to `true`.
     */
    autoCancel?: pulumi.Input<boolean>;
    /**
     * The branches to include and exclude from the trigger.
     */
    branchFilters?: pulumi.Input<pulumi.Input<inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilter>[]>;
    /**
     * Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     */
    pathFilters?: pulumi.Input<pulumi.Input<inputs.BuildDefinitionPullRequestTriggerOverridePathFilter>[]>;
}

export interface BuildDefinitionPullRequestTriggerOverrideBranchFilter {
    /**
     * List of branch patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of branch patterns to include.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BuildDefinitionPullRequestTriggerOverridePathFilter {
    /**
     * List of path patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of path patterns to include.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BuildDefinitionRepository {
    /**
     * The branch name for which builds are triggered. Defaults to `master`.
     */
    branchName?: pulumi.Input<string>;
    /**
     * The Github Enterprise URL. Used if `repoType` is `GithubEnterprise`.
     */
    githubEnterpriseUrl?: pulumi.Input<string>;
    /**
     * The id of the repository. For `TfsGit` repos, this is simply the ID of the repository. For `Github` repos, this will take the form of `<GitHub Org>/<Repo Name>`. For `Bitbucket` repos, this will take the form of `<Workspace ID>/<Repo Name>`.
     */
    repoId: pulumi.Input<string>;
    /**
     * The repository type. Valid values: `GitHub` or `TfsGit` or `Bitbucket` or `GitHub Enterprise`. Defaults to `GitHub`. If `repoType` is `GitHubEnterprise`, must use existing project and GitHub Enterprise service connection.
     */
    repoType: pulumi.Input<string>;
    /**
     * Report build status. Default is true.
     */
    reportBuildStatus?: pulumi.Input<boolean>;
    /**
     * The service connection ID. Used if the `repoType` is `GitHub` or `GitHubEnterprise`.
     */
    serviceConnectionId?: pulumi.Input<string>;
    /**
     * The path of the Yaml file describing the build definition.
     */
    ymlPath: pulumi.Input<string>;
}

export interface BuildDefinitionSchedule {
    /**
     * block supports the following:
     */
    branchFilters?: pulumi.Input<pulumi.Input<inputs.BuildDefinitionScheduleBranchFilter>[]>;
    /**
     * When to build. Valid values: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`.
     */
    daysToBuilds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the schedule job
     */
    scheduleJobId?: pulumi.Input<string>;
    /**
     * Schedule builds if the source or pipeline has changed. Defaults to `true`.
     */
    scheduleOnlyWithChanges?: pulumi.Input<boolean>;
    /**
     * Build start hour. Defaults to `0`. Valid values: `0 ~ 23`.
     */
    startHours?: pulumi.Input<number>;
    /**
     * Build start minute. Defaults to `0`. Valid values: `0 ~ 59`.
     */
    startMinutes?: pulumi.Input<number>;
    /**
     * Build time zone. Defaults to `(UTC) Coordinated Universal Time`. Valid values: 
     * `(UTC-12:00) International Date Line West`,
     * `(UTC-11:00) Coordinated Universal Time-11`,
     * `(UTC-10:00) Aleutian Islands`,
     * `(UTC-10:00) Hawaii`,
     * `(UTC-09:30) Marquesas Islands`,
     * `(UTC-09:00) Alaska`,
     * `(UTC-09:00) Coordinated Universal Time-09`,
     * `(UTC-08:00) Baja California`,
     * `(UTC-08:00) Coordinated Universal Time-08`,
     * `(UTC-08:00) Pacific Time (US &Canada)`,
     * `(UTC-07:00) Arizona`,
     * `(UTC-07:00) Chihuahua, La Paz, Mazatlan`,
     * `(UTC-07:00) Mountain Time (US &Canada)`,
     * `(UTC-07:00) Yukon`,
     * `(UTC-06:00) Central America`,
     * `(UTC-06:00) Central Time (US &Canada)`,
     * `(UTC-06:00) Easter Island`,
     * `(UTC-06:00) Guadalajara, Mexico City, Monterrey`,
     * `(UTC-06:00) Saskatchewan`,
     * `(UTC-05:00) Bogota, Lima, Quito, Rio Branco`,
     * `(UTC-05:00) Chetumal`,
     * `(UTC-05:00) Eastern Time (US &Canada)`,
     * `(UTC-05:00) Haiti`,
     * `(UTC-05:00) Havana`,
     * `(UTC-05:00) Indiana (East)`,
     * `(UTC-05:00) Turks and Caicos`,
     * `(UTC-04:00) Asuncion`,
     * `(UTC-04:00) Atlantic Time (Canada)`,
     * `(UTC-04:00) Caracas`,
     * `(UTC-04:00) Cuiaba`,
     * `(UTC-04:00) Georgetown, La Paz, Manaus, San Juan`,
     * `(UTC-04:00) Santiago`,
     * `(UTC-03:30) Newfoundland`,
     * `(UTC-03:00) Araguaina`,
     * `(UTC-03:00) Brasilia`,
     * `(UTC-03:00) Cayenne, Fortaleza`,
     * `(UTC-03:00) City of Buenos Aires`,
     * `(UTC-03:00) Greenland`,
     * `(UTC-03:00) Montevideo`,
     * `(UTC-03:00) Punta Arenas`,
     * `(UTC-03:00) Saint Pierre and Miquelon`,
     * `(UTC-03:00) Salvador`,
     * `(UTC-02:00) Coordinated Universal Time-02`,
     * `(UTC-02:00) Mid-Atlantic - Old`,
     * `(UTC-01:00) Azores`,
     * `(UTC-01:00) Cabo Verde Is.`,
     * `(UTC) Coordinated Universal Time`,
     * `(UTC+00:00) Dublin, Edinburgh, Lisbon, London`,
     * `(UTC+00:00) Monrovia, Reykjavik`,
     * `(UTC+00:00) Sao Tome`,
     * `(UTC+01:00) Casablanca`,
     * `(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna`,
     * `(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague`,
     * `(UTC+01:00) Brussels, Copenhagen, Madrid, Paris`,
     * `(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb`,
     * `(UTC+01:00) West Central Africa`,
     * `(UTC+02:00) Amman`,
     * `(UTC+02:00) Athens, Bucharest`,
     * `(UTC+02:00) Beirut`,
     * `(UTC+02:00) Cairo`,
     * `(UTC+02:00) Chisinau`,
     * `(UTC+02:00) Damascus`,
     * `(UTC+02:00) Gaza, Hebron`,
     * `(UTC+02:00) Harare, Pretoria`,
     * `(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius`,
     * `(UTC+02:00) Jerusalem`,
     * `(UTC+02:00) Juba`,
     * `(UTC+02:00) Kaliningrad`,
     * `(UTC+02:00) Khartoum`,
     * `(UTC+02:00) Tripoli`,
     * `(UTC+02:00) Windhoek`,
     * `(UTC+03:00) Baghdad`,
     * `(UTC+03:00) Istanbul`,
     * `(UTC+03:00) Kuwait, Riyadh`,
     * `(UTC+03:00) Minsk`,
     * `(UTC+03:00) Moscow, St. Petersburg`,
     * `(UTC+03:00) Nairobi`,
     * `(UTC+03:00) Volgograd`,
     * `(UTC+03:30) Tehran`,
     * `(UTC+04:00) Abu Dhabi, Muscat`,
     * `(UTC+04:00) Astrakhan, Ulyanovsk`,
     * `(UTC+04:00) Baku`,
     * `(UTC+04:00) Izhevsk, Samara`,
     * `(UTC+04:00) Port Louis`,
     * `(UTC+04:00) Saratov`,
     * `(UTC+04:00) Tbilisi`,
     * `(UTC+04:00) Yerevan`,
     * `(UTC+04:30) Kabul`,
     * `(UTC+05:00) Ashgabat, Tashkent`,
     * `(UTC+05:00) Ekaterinburg`,
     * `(UTC+05:00) Islamabad, Karachi`,
     * `(UTC+05:00) Qyzylorda`,
     * `(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi`,
     * `(UTC+05:30) Sri Jayawardenepura`,
     * `(UTC+05:45) Kathmandu`,
     * `(UTC+06:00) Astana`,
     * `(UTC+06:00) Dhaka`,
     * `(UTC+06:00) Omsk`,
     * `(UTC+06:30) Yangon (Rangoon)`,
     * `(UTC+07:00) Bangkok, Hanoi, Jakarta`,
     * `(UTC+07:00) Barnaul, Gorno-Altaysk`,
     * `(UTC+07:00) Hovd`,
     * `(UTC+07:00) Krasnoyarsk`,
     * `(UTC+07:00) Novosibirsk`,
     * `(UTC+07:00) Tomsk`,
     * `(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi`,
     * `(UTC+08:00) Irkutsk`,
     * `(UTC+08:00) Kuala Lumpur, Singapore`,
     * `(UTC+08:00) Perth`,
     * `(UTC+08:00) Taipei`,
     * `(UTC+08:00) Ulaanbaatar`,
     * `(UTC+08:45) Eucla`,
     * `(UTC+09:00) Chita`,
     * `(UTC+09:00) Osaka, Sapporo, Tokyo`,
     * `(UTC+09:00) Pyongyang`,
     * `(UTC+09:00) Seoul`,
     * `(UTC+09:00) Yakutsk`,
     * `(UTC+09:30) Adelaide`,
     * `(UTC+09:30) Darwin`,
     * `(UTC+10:00) Brisbane`,
     * `(UTC+10:00) Canberra, Melbourne, Sydney`,
     * `(UTC+10:00) Guam, Port Moresby`,
     * `(UTC+10:00) Hobart`,
     * `(UTC+10:00) Vladivostok`,
     * `(UTC+10:30) Lord Howe Island`,
     * `(UTC+11:00) Bougainville Island`,
     * `(UTC+11:00) Chokurdakh`,
     * `(UTC+11:00) Magadan`,
     * `(UTC+11:00) Norfolk Island`,
     * `(UTC+11:00) Sakhalin`,
     * `(UTC+11:00) Solomon Is., New Caledonia`,
     * `(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky`,
     * `(UTC+12:00) Auckland, Wellington`,
     * `(UTC+12:00) Coordinated Universal Time+12`,
     * `(UTC+12:00) Fiji`,
     * `(UTC+12:00) Petropavlovsk-Kamchatsky - Old`,
     * `(UTC+12:45) Chatham Islands`,
     * `(UTC+13:00) Coordinated Universal Time+13`,
     * `(UTC+13:00) Nuku'alofa`,
     * `(UTC+13:00) Samoa`,
     * `(UTC+14:00) Kiritimati Island`.
     */
    timeZone?: pulumi.Input<string>;
}

export interface BuildDefinitionScheduleBranchFilter {
    /**
     * List of branch patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of branch patterns to include.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BuildDefinitionVariable {
    /**
     * True if the variable can be overridden. Defaults to `true`.
     */
    allowOverride?: pulumi.Input<boolean>;
    /**
     * True if the variable is a secret. Defaults to `false`.
     */
    isSecret?: pulumi.Input<boolean>;
    /**
     * The name of the variable.
     */
    name: pulumi.Input<string>;
    /**
     * The secret value of the variable. Used when `isSecret` set to `true`.
     */
    secretValue?: pulumi.Input<string>;
    /**
     * The value of the variable.
     */
    value?: pulumi.Input<string>;
}

export interface CheckRequiredTemplateRequiredTemplate {
    /**
     * The name of the repository storing the template.
     */
    repositoryName: pulumi.Input<string>;
    /**
     * The branch in which the template will be referenced.
     */
    repositoryRef: pulumi.Input<string>;
    /**
     * The type of the repository storing the template. Valid values: `azuregit`, `github`, `githubenterprise`, `bitbucket`. Defaults to `azuregit`.
     */
    repositoryType?: pulumi.Input<string>;
    /**
     * The path to the template yaml.
     */
    templatePath: pulumi.Input<string>;
}

export interface FeedFeature {
    /**
     * Determines if Feed should be Permanently removed, Defaults to `false`
     */
    permanentDelete?: pulumi.Input<boolean>;
    /**
     * Determines if Feed should be Restored during creation (if possible), Defaults to `false`
     */
    restore?: pulumi.Input<boolean>;
}

export interface GetUsersFeatures {
    /**
     * Number of workers to process user data concurrently.
     *
     * > **Note** Setting `concurrentWorkers` to a value greater than 1 can greatly decrease the time it takes to read the data source.
     */
    concurrentWorkers?: number;
}

export interface GetUsersFeaturesArgs {
    /**
     * Number of workers to process user data concurrently.
     *
     * > **Note** Setting `concurrentWorkers` to a value greater than 1 can greatly decrease the time it takes to read the data source.
     */
    concurrentWorkers?: pulumi.Input<number>;
}

export interface GitInitialization {
    /**
     * The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`.
     */
    initType: pulumi.Input<string>;
    /**
     * The id of service connection used to authenticate to a private repository for import initialization.
     */
    serviceConnectionId?: pulumi.Input<string>;
    /**
     * Type of the source repository. Used if the `initType` is `Import`. Valid values: `Git`.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * The URL of the source repository. Used if the `initType` is `Import`.
     */
    sourceUrl?: pulumi.Input<string>;
}

export interface ServiceEndpointArtifactoryAuthenticationBasic {
    /**
     * The Artifactory password.
     */
    password: pulumi.Input<string>;
    /**
     * The Artifactory user name.
     */
    username: pulumi.Input<string>;
}

export interface ServiceEndpointArtifactoryAuthenticationToken {
    /**
     * The Artifactory access token.
     */
    token: pulumi.Input<string>;
}

export interface ServiceEndpointAzureRMCredentials {
    /**
     * The service principal application Id
     */
    serviceprincipalid: pulumi.Input<string>;
    /**
     * The service principal secret. This not required if `serviceEndpointAuthenticationScheme` is set to `WorkloadIdentityFederation`.
     */
    serviceprincipalkey?: pulumi.Input<string>;
}

export interface ServiceEndpointAzureRMFeatures {
    /**
     * Whether or not to validate connection with Azure after create or update operations. Defaults to `false`
     */
    validate?: pulumi.Input<boolean>;
}

export interface ServiceEndpointGitHubAuthOauth {
    oauthConfigurationId: pulumi.Input<string>;
}

export interface ServiceEndpointGitHubAuthPersonal {
    /**
     * The Personal Access Token for GitHub.
     */
    personalAccessToken: pulumi.Input<string>;
}

export interface ServiceEndpointGitHubEnterpriseAuthPersonal {
    /**
     * The Personal Access Token for GitHub.
     */
    personalAccessToken: pulumi.Input<string>;
}

export interface ServiceEndpointKubernetesAzureSubscription {
    /**
     * Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.
     */
    azureEnvironment?: pulumi.Input<string>;
    /**
     * Set this option to allow use cluster admin credentials.
     */
    clusterAdmin?: pulumi.Input<boolean>;
    /**
     * The name of the Kubernetes cluster.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The Kubernetes namespace. Default value is "default".
     */
    namespace?: pulumi.Input<string>;
    /**
     * The resource group name, to which the Kubernetes cluster is deployed.
     */
    resourcegroupId: pulumi.Input<string>;
    /**
     * The id of the Azure subscription.
     */
    subscriptionId: pulumi.Input<string>;
    /**
     * The name of the Azure subscription.
     */
    subscriptionName: pulumi.Input<string>;
    /**
     * The id of the tenant used by the subscription.
     */
    tenantId: pulumi.Input<string>;
}

export interface ServiceEndpointKubernetesKubeconfig {
    /**
     * Set this option to allow clients to accept a self-signed certificate.
     */
    acceptUntrustedCerts?: pulumi.Input<boolean>;
    /**
     * Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.
     */
    clusterContext?: pulumi.Input<string>;
    /**
     * The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.
     */
    kubeConfig: pulumi.Input<string>;
}

export interface ServiceEndpointKubernetesServiceAccount {
    /**
     * The certificate from a Kubernetes secret object.
     */
    caCert: pulumi.Input<string>;
    /**
     * The token from a Kubernetes secret object.
     */
    token: pulumi.Input<string>;
}

export interface ServiceEndpointPipelineAuthPersonal {
    /**
     * The Personal Access Token for Azure DevOps Pipeline. It also can be set with AZDO_PERSONAL_ACCESS_TOKEN environment variable.
     */
    personalAccessToken: pulumi.Input<string>;
}

export interface ServiceEndpointServiceFabricAzureActiveDirectory {
    /**
     * Password for the Azure Active Directory account.
     */
    password: pulumi.Input<string>;
    /**
     * The common name(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (',')
     */
    serverCertificateCommonName?: pulumi.Input<string>;
    /**
     * Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
     */
    serverCertificateLookup: pulumi.Input<string>;
    /**
     * The thumbprint(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (',')
     */
    serverCertificateThumbprint?: pulumi.Input<string>;
    /**
     * Specify an Azure Active Directory account.
     */
    username: pulumi.Input<string>;
}

export interface ServiceEndpointServiceFabricCertificate {
    /**
     * Base64 encoding of the cluster's client certificate file.
     */
    clientCertificate: pulumi.Input<string>;
    /**
     * Password for the certificate.
     */
    clientCertificatePassword?: pulumi.Input<string>;
    /**
     * The common name(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (',')
     */
    serverCertificateCommonName?: pulumi.Input<string>;
    /**
     * Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
     */
    serverCertificateLookup: pulumi.Input<string>;
    /**
     * The thumbprint(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (',')
     */
    serverCertificateThumbprint?: pulumi.Input<string>;
}

export interface ServiceEndpointServiceFabricNone {
    /**
     * Fully qualified domain SPN for gMSA account. This is applicable only if `unsecured` option is disabled.
     */
    clusterSpn?: pulumi.Input<string>;
    /**
     * Skip using windows security for authentication.
     */
    unsecured?: pulumi.Input<boolean>;
}

export interface ServiceendpointArgocdAuthenticationBasic {
    /**
     * The ArgoCD password.
     */
    password: pulumi.Input<string>;
    /**
     * The ArgoCD user name.
     */
    username: pulumi.Input<string>;
}

export interface ServiceendpointArgocdAuthenticationToken {
    /**
     * The ArgoCD access token.
     */
    token: pulumi.Input<string>;
}

export interface ServiceendpointExternaltfsAuthPersonal {
    /**
     * The Personal Access Token for Azure DevOps Organization.
     */
    personalAccessToken: pulumi.Input<string>;
}

export interface ServiceendpointJfrogArtifactoryV2AuthenticationBasic {
    /**
     * Artifactory Password.
     */
    password: pulumi.Input<string>;
    /**
     * Artifactory Username.
     */
    username: pulumi.Input<string>;
}

export interface ServiceendpointJfrogArtifactoryV2AuthenticationToken {
    /**
     * Authentication Token generated through Artifactory.
     */
    token: pulumi.Input<string>;
}

export interface ServiceendpointJfrogDistributionV2AuthenticationBasic {
    /**
     * Artifactory Password.
     */
    password: pulumi.Input<string>;
    /**
     * Artifactory Username.
     */
    username: pulumi.Input<string>;
}

export interface ServiceendpointJfrogDistributionV2AuthenticationToken {
    /**
     * Authentication Token generated through Artifactory.
     */
    token: pulumi.Input<string>;
}

export interface ServiceendpointJfrogPlatformV2AuthenticationBasic {
    /**
     * Artifactory Password.
     */
    password: pulumi.Input<string>;
    /**
     * Artifactory Username.
     */
    username: pulumi.Input<string>;
}

export interface ServiceendpointJfrogPlatformV2AuthenticationToken {
    /**
     * Authentication Token generated through Artifactory.
     */
    token: pulumi.Input<string>;
}

export interface ServiceendpointJfrogXrayV2AuthenticationBasic {
    /**
     * Artifactory Password.
     */
    password: pulumi.Input<string>;
    /**
     * Artifactory Username.
     */
    username: pulumi.Input<string>;
}

export interface ServiceendpointJfrogXrayV2AuthenticationToken {
    /**
     * Authentication Token generated through Artifactory.
     */
    token: pulumi.Input<string>;
}

export interface ServiceendpointMavenAuthenticationBasic {
    /**
     * The password Maven Repository.
     */
    password: pulumi.Input<string>;
    /**
     * The Username of the Maven Repository.
     */
    username: pulumi.Input<string>;
}

export interface ServiceendpointMavenAuthenticationToken {
    /**
     * Authentication Token generated through maven repository.
     */
    token: pulumi.Input<string>;
}

export interface ServicehookStorageQueuePipelinesRunStateChangedEvent {
    /**
     * The pipeline ID that will generate an event. If not specified, all pipelines in the project will trigger the event.
     */
    pipelineId?: pulumi.Input<string>;
    /**
     * Which run result should generate an event. Only valid if publishedEvent is `RunStateChanged`. If not specified, all results will trigger the event.
     */
    runResultFilter?: pulumi.Input<string>;
    /**
     * Which run state should generate an event. Only valid if publishedEvent is `RunStateChanged`. If not specified, all states will trigger the event.
     */
    runStateFilter?: pulumi.Input<string>;
}

export interface ServicehookStorageQueuePipelinesStageStateChangedEvent {
    /**
     * The pipeline ID that will generate an event.
     */
    pipelineId?: pulumi.Input<string>;
    /**
     * Which stage should generate an event. Only valid if publishedEvent is `StageStateChanged`. If not specified, all stages will trigger the event.
     */
    stageName?: pulumi.Input<string>;
    /**
     * Which stage result should generate an event. Only valid if publishedEvent is `StageStateChanged`. If not specified, all results will trigger the event.
     */
    stageResultFilter?: pulumi.Input<string>;
    /**
     * Which stage state should generate an event. Only valid if publishedEvent is `StageStateChanged`. If not specified, all states will trigger the event.
     */
    stageStateFilter?: pulumi.Input<string>;
}

export interface VariableGroupKeyVault {
    /**
     * The name of the Azure key vault to link secrets from as variables.
     */
    name: pulumi.Input<string>;
    /**
     * Set the Azure Key Vault Secret search depth. Defaults to `20`.
     */
    searchDepth?: pulumi.Input<number>;
    /**
     * The id of the Azure subscription endpoint to access the key vault.
     */
    serviceEndpointId: pulumi.Input<string>;
}

export interface VariableGroupVariable {
    contentType?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    expires?: pulumi.Input<string>;
    /**
     * A boolean flag describing if the variable value is sensitive. Defaults to `false`.
     */
    isSecret?: pulumi.Input<boolean>;
    /**
     * The key value used for the variable. Must be unique within the Variable Group.
     */
    name: pulumi.Input<string>;
    /**
     * The secret value of the variable. If omitted, it will default to empty string. Used when `isSecret` set to `true`.
     */
    secretValue?: pulumi.Input<string>;
    /**
     * The value of the variable. If omitted, it will default to empty string.
     */
    value?: pulumi.Input<string>;
}
