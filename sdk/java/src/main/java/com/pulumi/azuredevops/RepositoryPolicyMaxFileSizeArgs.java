// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryPolicyMaxFileSizeArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryPolicyMaxFileSizeArgs Empty = new RepositoryPolicyMaxFileSizeArgs();

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
     * Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
     * 
     */
    @Import(name="maxFileSize", required=true)
    private Output<Integer> maxFileSize;

    /**
     * @return Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
     * 
     */
    public Output<Integer> maxFileSize() {
        return this.maxFileSize;
    }

    /**
     * The ID of the project in which the policy will be created.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project in which the policy will be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
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

    private RepositoryPolicyMaxFileSizeArgs() {}

    private RepositoryPolicyMaxFileSizeArgs(RepositoryPolicyMaxFileSizeArgs $) {
        this.blocking = $.blocking;
        this.enabled = $.enabled;
        this.maxFileSize = $.maxFileSize;
        this.projectId = $.projectId;
        this.repositoryIds = $.repositoryIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryPolicyMaxFileSizeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryPolicyMaxFileSizeArgs $;

        public Builder() {
            $ = new RepositoryPolicyMaxFileSizeArgs();
        }

        public Builder(RepositoryPolicyMaxFileSizeArgs defaults) {
            $ = new RepositoryPolicyMaxFileSizeArgs(Objects.requireNonNull(defaults));
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
         * @param maxFileSize Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
         * 
         * @return builder
         * 
         */
        public Builder maxFileSize(Output<Integer> maxFileSize) {
            $.maxFileSize = maxFileSize;
            return this;
        }

        /**
         * @param maxFileSize Block pushes that contain new or updated files larger than this limit. Possible values are: `1, 2, 5, 10, 50, 100, 200` (MB).
         * 
         * @return builder
         * 
         */
        public Builder maxFileSize(Integer maxFileSize) {
            return maxFileSize(Output.of(maxFileSize));
        }

        /**
         * @param projectId The ID of the project in which the policy will be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
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

        public RepositoryPolicyMaxFileSizeArgs build() {
            if ($.maxFileSize == null) {
                throw new MissingRequiredPropertyException("RepositoryPolicyMaxFileSizeArgs", "maxFileSize");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("RepositoryPolicyMaxFileSizeArgs", "projectId");
            }
            return $;
        }
    }

}
