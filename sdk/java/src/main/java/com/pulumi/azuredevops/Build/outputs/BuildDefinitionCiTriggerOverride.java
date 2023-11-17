// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.outputs;

import com.pulumi.azuredevops.Build.outputs.BuildDefinitionCiTriggerOverrideBranchFilter;
import com.pulumi.azuredevops.Build.outputs.BuildDefinitionCiTriggerOverridePathFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BuildDefinitionCiTriggerOverride {
    /**
     * @return If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
     * 
     */
    private @Nullable Boolean batch;
    /**
     * @return The branches to include and exclude from the trigger.`branch_filter` - (Optional) The branches to include and exclude from the trigger.
     * 
     */
    private @Nullable List<BuildDefinitionCiTriggerOverrideBranchFilter> branchFilters;
    /**
     * @return The number of max builds per branch. Defaults to `1`.
     * 
     */
    private @Nullable Integer maxConcurrentBuildsPerBranch;
    /**
     * @return Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.`path_filter` - (Optional) Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     * 
     */
    private @Nullable List<BuildDefinitionCiTriggerOverridePathFilter> pathFilters;
    /**
     * @return How often the external repository is polled. Defaults to `0`.
     * 
     */
    private @Nullable Integer pollingInterval;
    /**
     * @return This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     * 
     */
    private @Nullable String pollingJobId;

    private BuildDefinitionCiTriggerOverride() {}
    /**
     * @return If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
     * 
     */
    public Optional<Boolean> batch() {
        return Optional.ofNullable(this.batch);
    }
    /**
     * @return The branches to include and exclude from the trigger.`branch_filter` - (Optional) The branches to include and exclude from the trigger.
     * 
     */
    public List<BuildDefinitionCiTriggerOverrideBranchFilter> branchFilters() {
        return this.branchFilters == null ? List.of() : this.branchFilters;
    }
    /**
     * @return The number of max builds per branch. Defaults to `1`.
     * 
     */
    public Optional<Integer> maxConcurrentBuildsPerBranch() {
        return Optional.ofNullable(this.maxConcurrentBuildsPerBranch);
    }
    /**
     * @return Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.`path_filter` - (Optional) Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     * 
     */
    public List<BuildDefinitionCiTriggerOverridePathFilter> pathFilters() {
        return this.pathFilters == null ? List.of() : this.pathFilters;
    }
    /**
     * @return How often the external repository is polled. Defaults to `0`.
     * 
     */
    public Optional<Integer> pollingInterval() {
        return Optional.ofNullable(this.pollingInterval);
    }
    /**
     * @return This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     * 
     */
    public Optional<String> pollingJobId() {
        return Optional.ofNullable(this.pollingJobId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BuildDefinitionCiTriggerOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean batch;
        private @Nullable List<BuildDefinitionCiTriggerOverrideBranchFilter> branchFilters;
        private @Nullable Integer maxConcurrentBuildsPerBranch;
        private @Nullable List<BuildDefinitionCiTriggerOverridePathFilter> pathFilters;
        private @Nullable Integer pollingInterval;
        private @Nullable String pollingJobId;
        public Builder() {}
        public Builder(BuildDefinitionCiTriggerOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.batch = defaults.batch;
    	      this.branchFilters = defaults.branchFilters;
    	      this.maxConcurrentBuildsPerBranch = defaults.maxConcurrentBuildsPerBranch;
    	      this.pathFilters = defaults.pathFilters;
    	      this.pollingInterval = defaults.pollingInterval;
    	      this.pollingJobId = defaults.pollingJobId;
        }

        @CustomType.Setter
        public Builder batch(@Nullable Boolean batch) {
            this.batch = batch;
            return this;
        }
        @CustomType.Setter
        public Builder branchFilters(@Nullable List<BuildDefinitionCiTriggerOverrideBranchFilter> branchFilters) {
            this.branchFilters = branchFilters;
            return this;
        }
        public Builder branchFilters(BuildDefinitionCiTriggerOverrideBranchFilter... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }
        @CustomType.Setter
        public Builder maxConcurrentBuildsPerBranch(@Nullable Integer maxConcurrentBuildsPerBranch) {
            this.maxConcurrentBuildsPerBranch = maxConcurrentBuildsPerBranch;
            return this;
        }
        @CustomType.Setter
        public Builder pathFilters(@Nullable List<BuildDefinitionCiTriggerOverridePathFilter> pathFilters) {
            this.pathFilters = pathFilters;
            return this;
        }
        public Builder pathFilters(BuildDefinitionCiTriggerOverridePathFilter... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }
        @CustomType.Setter
        public Builder pollingInterval(@Nullable Integer pollingInterval) {
            this.pollingInterval = pollingInterval;
            return this;
        }
        @CustomType.Setter
        public Builder pollingJobId(@Nullable String pollingJobId) {
            this.pollingJobId = pollingJobId;
            return this;
        }
        public BuildDefinitionCiTriggerOverride build() {
            final var _resultValue = new BuildDefinitionCiTriggerOverride();
            _resultValue.batch = batch;
            _resultValue.branchFilters = branchFilters;
            _resultValue.maxConcurrentBuildsPerBranch = maxConcurrentBuildsPerBranch;
            _resultValue.pathFilters = pathFilters;
            _resultValue.pollingInterval = pollingInterval;
            _resultValue.pollingJobId = pollingJobId;
            return _resultValue;
        }
    }
}
