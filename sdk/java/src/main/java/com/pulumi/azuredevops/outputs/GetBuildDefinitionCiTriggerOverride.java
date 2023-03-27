// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetBuildDefinitionCiTriggerOverrideBranchFilter;
import com.pulumi.azuredevops.outputs.GetBuildDefinitionCiTriggerOverridePathFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionCiTriggerOverride {
    /**
     * @return If batch is true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built.
     * 
     */
    private Boolean batch;
    /**
     * @return A `branch_filter` block as defined above.
     * 
     */
    private List<GetBuildDefinitionCiTriggerOverrideBranchFilter> branchFilters;
    /**
     * @return The number of max builds per branch.
     * 
     */
    private Integer maxConcurrentBuildsPerBranch;
    /**
     * @return block supports the following:
     * 
     */
    private List<GetBuildDefinitionCiTriggerOverridePathFilter> pathFilters;
    /**
     * @return How often the external repository is polled.
     * 
     */
    private Integer pollingInterval;
    /**
     * @return This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     * 
     */
    private String pollingJobId;

    private GetBuildDefinitionCiTriggerOverride() {}
    /**
     * @return If batch is true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built.
     * 
     */
    public Boolean batch() {
        return this.batch;
    }
    /**
     * @return A `branch_filter` block as defined above.
     * 
     */
    public List<GetBuildDefinitionCiTriggerOverrideBranchFilter> branchFilters() {
        return this.branchFilters;
    }
    /**
     * @return The number of max builds per branch.
     * 
     */
    public Integer maxConcurrentBuildsPerBranch() {
        return this.maxConcurrentBuildsPerBranch;
    }
    /**
     * @return block supports the following:
     * 
     */
    public List<GetBuildDefinitionCiTriggerOverridePathFilter> pathFilters() {
        return this.pathFilters;
    }
    /**
     * @return How often the external repository is polled.
     * 
     */
    public Integer pollingInterval() {
        return this.pollingInterval;
    }
    /**
     * @return This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     * 
     */
    public String pollingJobId() {
        return this.pollingJobId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionCiTriggerOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean batch;
        private List<GetBuildDefinitionCiTriggerOverrideBranchFilter> branchFilters;
        private Integer maxConcurrentBuildsPerBranch;
        private List<GetBuildDefinitionCiTriggerOverridePathFilter> pathFilters;
        private Integer pollingInterval;
        private String pollingJobId;
        public Builder() {}
        public Builder(GetBuildDefinitionCiTriggerOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.batch = defaults.batch;
    	      this.branchFilters = defaults.branchFilters;
    	      this.maxConcurrentBuildsPerBranch = defaults.maxConcurrentBuildsPerBranch;
    	      this.pathFilters = defaults.pathFilters;
    	      this.pollingInterval = defaults.pollingInterval;
    	      this.pollingJobId = defaults.pollingJobId;
        }

        @CustomType.Setter
        public Builder batch(Boolean batch) {
            this.batch = Objects.requireNonNull(batch);
            return this;
        }
        @CustomType.Setter
        public Builder branchFilters(List<GetBuildDefinitionCiTriggerOverrideBranchFilter> branchFilters) {
            this.branchFilters = Objects.requireNonNull(branchFilters);
            return this;
        }
        public Builder branchFilters(GetBuildDefinitionCiTriggerOverrideBranchFilter... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }
        @CustomType.Setter
        public Builder maxConcurrentBuildsPerBranch(Integer maxConcurrentBuildsPerBranch) {
            this.maxConcurrentBuildsPerBranch = Objects.requireNonNull(maxConcurrentBuildsPerBranch);
            return this;
        }
        @CustomType.Setter
        public Builder pathFilters(List<GetBuildDefinitionCiTriggerOverridePathFilter> pathFilters) {
            this.pathFilters = Objects.requireNonNull(pathFilters);
            return this;
        }
        public Builder pathFilters(GetBuildDefinitionCiTriggerOverridePathFilter... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }
        @CustomType.Setter
        public Builder pollingInterval(Integer pollingInterval) {
            this.pollingInterval = Objects.requireNonNull(pollingInterval);
            return this;
        }
        @CustomType.Setter
        public Builder pollingJobId(String pollingJobId) {
            this.pollingJobId = Objects.requireNonNull(pollingJobId);
            return this;
        }
        public GetBuildDefinitionCiTriggerOverride build() {
            final var o = new GetBuildDefinitionCiTriggerOverride();
            o.batch = batch;
            o.branchFilters = branchFilters;
            o.maxConcurrentBuildsPerBranch = maxConcurrentBuildsPerBranch;
            o.pathFilters = pathFilters;
            o.pollingInterval = pollingInterval;
            o.pollingJobId = pollingJobId;
            return o;
        }
    }
}
