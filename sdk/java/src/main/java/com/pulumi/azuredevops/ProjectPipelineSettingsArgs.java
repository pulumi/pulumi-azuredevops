// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectPipelineSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectPipelineSettingsArgs Empty = new ProjectPipelineSettingsArgs();

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
     * The `id` of the project for which the project pipeline settings will be managed.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The `id` of the project for which the project pipeline settings will be managed.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
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

    private ProjectPipelineSettingsArgs() {}

    private ProjectPipelineSettingsArgs(ProjectPipelineSettingsArgs $) {
        this.enforceJobScope = $.enforceJobScope;
        this.enforceReferencedRepoScopedToken = $.enforceReferencedRepoScopedToken;
        this.enforceSettableVar = $.enforceSettableVar;
        this.projectId = $.projectId;
        this.publishPipelineMetadata = $.publishPipelineMetadata;
        this.statusBadgesArePrivate = $.statusBadgesArePrivate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectPipelineSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectPipelineSettingsArgs $;

        public Builder() {
            $ = new ProjectPipelineSettingsArgs();
        }

        public Builder(ProjectPipelineSettingsArgs defaults) {
            $ = new ProjectPipelineSettingsArgs(Objects.requireNonNull(defaults));
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
         * @param projectId The `id` of the project for which the project pipeline settings will be managed.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The `id` of the project for which the project pipeline settings will be managed.
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

        public ProjectPipelineSettingsArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
