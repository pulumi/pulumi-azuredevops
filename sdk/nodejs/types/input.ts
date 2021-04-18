// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

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
     * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
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
     * If a path filter is set, the policy wil only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
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
     * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
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
     * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
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
     * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
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
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     */
    scopes: pulumi.Input<pulumi.Input<inputs.BranchPolicyMinReviewersSettingsScope>[]>;
    /**
     * Allow requesters to approve their own changes. Defaults to `false`.
     */
    submitterCanVote?: pulumi.Input<boolean>;
}

export interface BranchPolicyMinReviewersSettingsScope {
    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     */
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
     * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
     */
    matchType?: pulumi.Input<string>;
    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
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
     * List of branch patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of branch patterns to include.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Build pull requests form forms of this repository.
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
     * List of branch patterns to exclude.
     */
    excludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of branch patterns to include.
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

export interface GitInitialization {
    /**
     * The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`. Defaults to `Uninitialized`.
     */
    initType: pulumi.Input<string>;
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
     * Artifactory Password.
     */
    password: pulumi.Input<string>;
    passwordHash?: pulumi.Input<string>;
    /**
     * Artifactory Username.
     */
    username: pulumi.Input<string>;
    usernameHash?: pulumi.Input<string>;
}

export interface ServiceEndpointArtifactoryAuthenticationToken {
    /**
     * Authentication Token generated through Artifactory.
     * * `authenticationBasic`
     */
    token: pulumi.Input<string>;
    tokenHash?: pulumi.Input<string>;
}

export interface ServiceEndpointAzureRMCredentials {
    /**
     * The service principal application Id
     */
    serviceprincipalid: pulumi.Input<string>;
    /**
     * The service principal secret.
     */
    serviceprincipalkey: pulumi.Input<string>;
    serviceprincipalkeyHash?: pulumi.Input<string>;
}

export interface ServiceEndpointGitHubAuthOauth {
    oauthConfigurationId: pulumi.Input<string>;
}

export interface ServiceEndpointGitHubAuthPersonal {
    /**
     * The Personal Access Token for Github.
     */
    personalAccessToken: pulumi.Input<string>;
    personalAccessTokenHash?: pulumi.Input<string>;
}

export interface ServiceEndpointGitHubEnterpriseAuthPersonal {
    /**
     * The Personal Access Token for Github.
     */
    personalAccessToken: pulumi.Input<string>;
    personalAccessTokenHash?: pulumi.Input<string>;
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
    kubeConfigHash?: pulumi.Input<string>;
}

export interface ServiceEndpointKubernetesServiceAccount {
    /**
     * The certificate from a Kubernetes secret object.
     */
    caCert: pulumi.Input<string>;
    caCertHash?: pulumi.Input<string>;
    /**
     * The token from a Kubernetes secret object.
     */
    token: pulumi.Input<string>;
    tokenHash?: pulumi.Input<string>;
}

export interface ServiceEndpointPipelineAuthPersonal {
    /**
     * The Personal Access Token for Azure DevOps Pipeline. It also can be set with AZDO_PERSONAL_ACCESS_TOKEN environment variable.
     */
    personalAccessToken: pulumi.Input<string>;
    personalAccessTokenHash?: pulumi.Input<string>;
}

export interface VariableGroupKeyVault {
    /**
     * The name of the Variable Group.
     */
    name: pulumi.Input<string>;
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
export namespace Agent {
}

export namespace Build {
    export interface BuildDefinitionCiTrigger {
        /**
         * Override the azure-pipeline file and use a this configuration for all builds.
         */
        override?: pulumi.Input<inputs.Build.BuildDefinitionCiTriggerOverride>;
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
        branchFilters?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionCiTriggerOverrideBranchFilter>[]>;
        /**
         * The number of max builds per branch. Defaults to `1`.
         */
        maxConcurrentBuildsPerBranch?: pulumi.Input<number>;
        /**
         * Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         */
        pathFilters?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionCiTriggerOverridePathFilter>[]>;
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
         * List of branch patterns to exclude.
         */
        excludes?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * List of branch patterns to include.
         */
        includes?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface BuildDefinitionPullRequestTrigger {
        commentRequired?: pulumi.Input<string>;
        /**
         * Set permissions for Forked repositories.
         */
        forks: pulumi.Input<inputs.Build.BuildDefinitionPullRequestTriggerForks>;
        initialBranch?: pulumi.Input<string>;
        /**
         * Override the azure-pipeline file and use this configuration for all builds.
         */
        override?: pulumi.Input<inputs.Build.BuildDefinitionPullRequestTriggerOverride>;
        /**
         * Use the azure-pipeline file for the build configuration. Defaults to `false`.
         */
        useYaml?: pulumi.Input<boolean>;
    }

    export interface BuildDefinitionPullRequestTriggerForks {
        /**
         * Build pull requests form forms of this repository.
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
        branchFilters?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionPullRequestTriggerOverrideBranchFilter>[]>;
        /**
         * Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         */
        pathFilters?: pulumi.Input<pulumi.Input<inputs.Build.BuildDefinitionPullRequestTriggerOverridePathFilter>[]>;
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
         * List of branch patterns to exclude.
         */
        excludes?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * List of branch patterns to include.
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
}

export namespace Core {
}

export namespace Identities {
}

export namespace Pipeline {
    export interface VariableGroupKeyVault {
        /**
         * The name of the Variable Group.
         */
        name: pulumi.Input<string>;
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
}

export namespace Policy {
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
         * If a path filter is set, the policy wil only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
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
        scopes: pulumi.Input<pulumi.Input<inputs.Policy.BranchPolicyBuildValidationSettingsScope>[]>;
        /**
         * The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
         */
        validDuration?: pulumi.Input<number>;
    }

    export interface BranchPolicyBuildValidationSettingsScope {
        /**
         * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
         */
        matchType?: pulumi.Input<string>;
        /**
         * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
         */
        repositoryId?: pulumi.Input<string>;
        /**
         * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
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
         * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         */
        scopes: pulumi.Input<pulumi.Input<inputs.Policy.BranchPolicyMinReviewersSettingsScope>[]>;
        /**
         * Allow requesters to approve their own changes. Defaults to `false`.
         */
        submitterCanVote?: pulumi.Input<boolean>;
    }

    export interface BranchPolicyMinReviewersSettingsScope {
        /**
         * The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
         */
        matchType?: pulumi.Input<string>;
        /**
         * The repository ID. Needed only if the scope of the policy will be limited to a single repository.
         */
        repositoryId?: pulumi.Input<string>;
        /**
         * The ref pattern to use for the match. If `matchType` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `matchType` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
         */
        repositoryRef?: pulumi.Input<string>;
    }
}

export namespace Repository {
    export interface GitInitialization {
        /**
         * The type of repository to create. Valid values: `Uninitialized`, `Clean` or `Import`. Defaults to `Uninitialized`.
         */
        initType: pulumi.Input<string>;
        /**
         * Type of the source repository. Used if the `initType` is `Import`. Valid values: `Git`.
         */
        sourceType?: pulumi.Input<string>;
        /**
         * The URL of the source repository. Used if the `initType` is `Import`.
         */
        sourceUrl?: pulumi.Input<string>;
    }
}

export namespace ServiceEndpoint {
    export interface AzureRMCredentials {
        /**
         * The service principal application Id
         */
        serviceprincipalid: pulumi.Input<string>;
        /**
         * The service principal secret.
         */
        serviceprincipalkey: pulumi.Input<string>;
        serviceprincipalkeyHash?: pulumi.Input<string>;
    }

    export interface GitHubAuthOauth {
        oauthConfigurationId: pulumi.Input<string>;
    }

    export interface GitHubAuthPersonal {
        /**
         * The Personal Access Token for Github.
         */
        personalAccessToken: pulumi.Input<string>;
        personalAccessTokenHash?: pulumi.Input<string>;
    }

    export interface KubernetesAzureSubscription {
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

    export interface KubernetesKubeconfig {
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
        kubeConfigHash?: pulumi.Input<string>;
    }

    export interface KubernetesServiceAccount {
        /**
         * The certificate from a Kubernetes secret object.
         */
        caCert: pulumi.Input<string>;
        caCertHash?: pulumi.Input<string>;
        /**
         * The token from a Kubernetes secret object.
         */
        token: pulumi.Input<string>;
        tokenHash?: pulumi.Input<string>;
    }
}
