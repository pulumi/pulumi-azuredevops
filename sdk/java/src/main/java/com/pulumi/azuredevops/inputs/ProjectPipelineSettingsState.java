// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectPipelineSettingsState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectPipelineSettingsState Empty = new ProjectPipelineSettingsState();

    /**
     * Limit job authorization scope to current project for non-release pipelines.
     * 
     */
    @Import(name="enforceJobScope")
    private @Nullable Output<Boolean> enforceJobScope;

    /**
     * @return Limit job authorization scope to current project for non-release pipelines.
     * 
     */
    public Optional<Output<Boolean>> enforceJobScope() {
        return Optional.ofNullable(this.enforceJobScope);
    }

    /**
     * Limit job authorization scope to current project for release pipelines.
     * 
     * &gt; **NOTE:** The settings at the organization will override settings specified on the project.
     * For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
     * In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
     * 
     */
    @Import(name="enforceJobScopeForRelease")
    private @Nullable Output<Boolean> enforceJobScopeForRelease;

    /**
     * @return Limit job authorization scope to current project for release pipelines.
     * 
     * &gt; **NOTE:** The settings at the organization will override settings specified on the project.
     * For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
     * In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
     * 
     */
    public Optional<Output<Boolean>> enforceJobScopeForRelease() {
        return Optional.ofNullable(this.enforceJobScopeForRelease);
    }

    /**
     * Protect access to repositories in YAML pipelines.
     * 
     */
    @Import(name="enforceReferencedRepoScopedToken")
    private @Nullable Output<Boolean> enforceReferencedRepoScopedToken;

    /**
     * @return Protect access to repositories in YAML pipelines.
     * 
     */
    public Optional<Output<Boolean>> enforceReferencedRepoScopedToken() {
        return Optional.ofNullable(this.enforceReferencedRepoScopedToken);
    }

    /**
     * Limit variables that can be set at queue time.
     * 
     */
    @Import(name="enforceSettableVar")
    private @Nullable Output<Boolean> enforceSettableVar;

    /**
     * @return Limit variables that can be set at queue time.
     * 
     */
    public Optional<Output<Boolean>> enforceSettableVar() {
        return Optional.ofNullable(this.enforceSettableVar);
    }

    /**
     * The ID of the project for which the project pipeline settings will be managed.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project for which the project pipeline settings will be managed.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Publish metadata from pipelines.
     * 
     */
    @Import(name="publishPipelineMetadata")
    private @Nullable Output<Boolean> publishPipelineMetadata;

    /**
     * @return Publish metadata from pipelines.
     * 
     */
    public Optional<Output<Boolean>> publishPipelineMetadata() {
        return Optional.ofNullable(this.publishPipelineMetadata);
    }

    /**
     * Disable anonymous access to badges.
     * 
     */
    @Import(name="statusBadgesArePrivate")
    private @Nullable Output<Boolean> statusBadgesArePrivate;

    /**
     * @return Disable anonymous access to badges.
     * 
     */
    public Optional<Output<Boolean>> statusBadgesArePrivate() {
        return Optional.ofNullable(this.statusBadgesArePrivate);
    }

    private ProjectPipelineSettingsState() {}

    private ProjectPipelineSettingsState(ProjectPipelineSettingsState $) {
        this.enforceJobScope = $.enforceJobScope;
        this.enforceJobScopeForRelease = $.enforceJobScopeForRelease;
        this.enforceReferencedRepoScopedToken = $.enforceReferencedRepoScopedToken;
        this.enforceSettableVar = $.enforceSettableVar;
        this.projectId = $.projectId;
        this.publishPipelineMetadata = $.publishPipelineMetadata;
        this.statusBadgesArePrivate = $.statusBadgesArePrivate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectPipelineSettingsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectPipelineSettingsState $;

        public Builder() {
            $ = new ProjectPipelineSettingsState();
        }

        public Builder(ProjectPipelineSettingsState defaults) {
            $ = new ProjectPipelineSettingsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param enforceJobScope Limit job authorization scope to current project for non-release pipelines.
         * 
         * @return builder
         * 
         */
        public Builder enforceJobScope(@Nullable Output<Boolean> enforceJobScope) {
            $.enforceJobScope = enforceJobScope;
            return this;
        }

        /**
         * @param enforceJobScope Limit job authorization scope to current project for non-release pipelines.
         * 
         * @return builder
         * 
         */
        public Builder enforceJobScope(Boolean enforceJobScope) {
            return enforceJobScope(Output.of(enforceJobScope));
        }

        /**
         * @param enforceJobScopeForRelease Limit job authorization scope to current project for release pipelines.
         * 
         * &gt; **NOTE:** The settings at the organization will override settings specified on the project.
         * For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
         * In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
         * 
         * @return builder
         * 
         */
        public Builder enforceJobScopeForRelease(@Nullable Output<Boolean> enforceJobScopeForRelease) {
            $.enforceJobScopeForRelease = enforceJobScopeForRelease;
            return this;
        }

        /**
         * @param enforceJobScopeForRelease Limit job authorization scope to current project for release pipelines.
         * 
         * &gt; **NOTE:** The settings at the organization will override settings specified on the project.
         * For example, if `enforce_job_scope` is true at the organization, the `azuredevops.ProjectPipelineSettings` resource cannot set it to false.
         * In this scenario, the plan will always show that the resource is trying to change `enforce_job_scope` from `true` to `false`.
         * 
         * @return builder
         * 
         */
        public Builder enforceJobScopeForRelease(Boolean enforceJobScopeForRelease) {
            return enforceJobScopeForRelease(Output.of(enforceJobScopeForRelease));
        }

        /**
         * @param enforceReferencedRepoScopedToken Protect access to repositories in YAML pipelines.
         * 
         * @return builder
         * 
         */
        public Builder enforceReferencedRepoScopedToken(@Nullable Output<Boolean> enforceReferencedRepoScopedToken) {
            $.enforceReferencedRepoScopedToken = enforceReferencedRepoScopedToken;
            return this;
        }

        /**
         * @param enforceReferencedRepoScopedToken Protect access to repositories in YAML pipelines.
         * 
         * @return builder
         * 
         */
        public Builder enforceReferencedRepoScopedToken(Boolean enforceReferencedRepoScopedToken) {
            return enforceReferencedRepoScopedToken(Output.of(enforceReferencedRepoScopedToken));
        }

        /**
         * @param enforceSettableVar Limit variables that can be set at queue time.
         * 
         * @return builder
         * 
         */
        public Builder enforceSettableVar(@Nullable Output<Boolean> enforceSettableVar) {
            $.enforceSettableVar = enforceSettableVar;
            return this;
        }

        /**
         * @param enforceSettableVar Limit variables that can be set at queue time.
         * 
         * @return builder
         * 
         */
        public Builder enforceSettableVar(Boolean enforceSettableVar) {
            return enforceSettableVar(Output.of(enforceSettableVar));
        }

        /**
         * @param projectId The ID of the project for which the project pipeline settings will be managed.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project for which the project pipeline settings will be managed.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param publishPipelineMetadata Publish metadata from pipelines.
         * 
         * @return builder
         * 
         */
        public Builder publishPipelineMetadata(@Nullable Output<Boolean> publishPipelineMetadata) {
            $.publishPipelineMetadata = publishPipelineMetadata;
            return this;
        }

        /**
         * @param publishPipelineMetadata Publish metadata from pipelines.
         * 
         * @return builder
         * 
         */
        public Builder publishPipelineMetadata(Boolean publishPipelineMetadata) {
            return publishPipelineMetadata(Output.of(publishPipelineMetadata));
        }

        /**
         * @param statusBadgesArePrivate Disable anonymous access to badges.
         * 
         * @return builder
         * 
         */
        public Builder statusBadgesArePrivate(@Nullable Output<Boolean> statusBadgesArePrivate) {
            $.statusBadgesArePrivate = statusBadgesArePrivate;
            return this;
        }

        /**
         * @param statusBadgesArePrivate Disable anonymous access to badges.
         * 
         * @return builder
         * 
         */
        public Builder statusBadgesArePrivate(Boolean statusBadgesArePrivate) {
            return statusBadgesArePrivate(Output.of(statusBadgesArePrivate));
        }

        public ProjectPipelineSettingsState build() {
            return $;
        }
    }

}
