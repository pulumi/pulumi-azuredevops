// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryPolicyCaseEnforcementState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryPolicyCaseEnforcementState Empty = new RepositoryPolicyCaseEnforcementState();

    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     * 
     */
    @Import(name="blocking")
    private @Nullable Output<Boolean> blocking;

    /**
     * @return A flag indicating if the policy should be blocking. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> blocking() {
        return Optional.ofNullable(this.blocking);
    }

    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
     * 
     */
    @Import(name="enforceConsistentCase")
    private @Nullable Output<Boolean> enforceConsistentCase;

    /**
     * @return Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
     * 
     */
    public Optional<Output<Boolean>> enforceConsistentCase() {
        return Optional.ofNullable(this.enforceConsistentCase);
    }

    /**
     * The ID of the project in which the policy will be created.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project in which the policy will be created.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
     * 
     */
    @Import(name="repositoryIds")
    private @Nullable Output<List<String>> repositoryIds;

    /**
     * @return Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
     * 
     */
    public Optional<Output<List<String>>> repositoryIds() {
        return Optional.ofNullable(this.repositoryIds);
    }

    private RepositoryPolicyCaseEnforcementState() {}

    private RepositoryPolicyCaseEnforcementState(RepositoryPolicyCaseEnforcementState $) {
        this.blocking = $.blocking;
        this.enabled = $.enabled;
        this.enforceConsistentCase = $.enforceConsistentCase;
        this.projectId = $.projectId;
        this.repositoryIds = $.repositoryIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryPolicyCaseEnforcementState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryPolicyCaseEnforcementState $;

        public Builder() {
            $ = new RepositoryPolicyCaseEnforcementState();
        }

        public Builder(RepositoryPolicyCaseEnforcementState defaults) {
            $ = new RepositoryPolicyCaseEnforcementState(Objects.requireNonNull(defaults));
        }

        /**
         * @param blocking A flag indicating if the policy should be blocking. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder blocking(@Nullable Output<Boolean> blocking) {
            $.blocking = blocking;
            return this;
        }

        /**
         * @param blocking A flag indicating if the policy should be blocking. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder blocking(Boolean blocking) {
            return blocking(Output.of(blocking));
        }

        /**
         * @param enabled A flag indicating if the policy should be enabled. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled A flag indicating if the policy should be enabled. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param enforceConsistentCase Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
         * 
         * @return builder
         * 
         */
        public Builder enforceConsistentCase(@Nullable Output<Boolean> enforceConsistentCase) {
            $.enforceConsistentCase = enforceConsistentCase;
            return this;
        }

        /**
         * @param enforceConsistentCase Avoid case-sensitivity conflicts by blocking pushes that change name casing on files, folders, branches, and tags.
         * 
         * @return builder
         * 
         */
        public Builder enforceConsistentCase(Boolean enforceConsistentCase) {
            return enforceConsistentCase(Output.of(enforceConsistentCase));
        }

        /**
         * @param projectId The ID of the project in which the policy will be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project in which the policy will be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param repositoryIds Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
         * 
         * @return builder
         * 
         */
        public Builder repositoryIds(@Nullable Output<List<String>> repositoryIds) {
            $.repositoryIds = repositoryIds;
            return this;
        }

        /**
         * @param repositoryIds Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
         * 
         * @return builder
         * 
         */
        public Builder repositoryIds(List<String> repositoryIds) {
            return repositoryIds(Output.of(repositoryIds));
        }

        /**
         * @param repositoryIds Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
         * 
         * @return builder
         * 
         */
        public Builder repositoryIds(String... repositoryIds) {
            return repositoryIds(List.of(repositoryIds));
        }

        public RepositoryPolicyCaseEnforcementState build() {
            return $;
        }
    }

}
