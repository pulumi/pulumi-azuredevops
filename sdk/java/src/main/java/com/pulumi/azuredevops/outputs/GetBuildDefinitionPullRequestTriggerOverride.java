// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetBuildDefinitionPullRequestTriggerOverrideBranchFilter;
import com.pulumi.azuredevops.outputs.GetBuildDefinitionPullRequestTriggerOverridePathFilter;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionPullRequestTriggerOverride {
    /**
     * @return Should further updates to a PR cancel an in progress validation?
     * 
     */
    private Boolean autoCancel;
    /**
     * @return A `branch_filter` block as defined above.
     * 
     */
    private List<GetBuildDefinitionPullRequestTriggerOverrideBranchFilter> branchFilters;
    /**
     * @return The file paths to include or exclude. A `path_filter` block as defined above.
     * 
     */
    private List<GetBuildDefinitionPullRequestTriggerOverridePathFilter> pathFilters;

    private GetBuildDefinitionPullRequestTriggerOverride() {}
    /**
     * @return Should further updates to a PR cancel an in progress validation?
     * 
     */
    public Boolean autoCancel() {
        return this.autoCancel;
    }
    /**
     * @return A `branch_filter` block as defined above.
     * 
     */
    public List<GetBuildDefinitionPullRequestTriggerOverrideBranchFilter> branchFilters() {
        return this.branchFilters;
    }
    /**
     * @return The file paths to include or exclude. A `path_filter` block as defined above.
     * 
     */
    public List<GetBuildDefinitionPullRequestTriggerOverridePathFilter> pathFilters() {
        return this.pathFilters;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionPullRequestTriggerOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean autoCancel;
        private List<GetBuildDefinitionPullRequestTriggerOverrideBranchFilter> branchFilters;
        private List<GetBuildDefinitionPullRequestTriggerOverridePathFilter> pathFilters;
        public Builder() {}
        public Builder(GetBuildDefinitionPullRequestTriggerOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoCancel = defaults.autoCancel;
    	      this.branchFilters = defaults.branchFilters;
    	      this.pathFilters = defaults.pathFilters;
        }

        @CustomType.Setter
        public Builder autoCancel(Boolean autoCancel) {
            if (autoCancel == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTriggerOverride", "autoCancel");
            }
            this.autoCancel = autoCancel;
            return this;
        }
        @CustomType.Setter
        public Builder branchFilters(List<GetBuildDefinitionPullRequestTriggerOverrideBranchFilter> branchFilters) {
            if (branchFilters == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTriggerOverride", "branchFilters");
            }
            this.branchFilters = branchFilters;
            return this;
        }
        public Builder branchFilters(GetBuildDefinitionPullRequestTriggerOverrideBranchFilter... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }
        @CustomType.Setter
        public Builder pathFilters(List<GetBuildDefinitionPullRequestTriggerOverridePathFilter> pathFilters) {
            if (pathFilters == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTriggerOverride", "pathFilters");
            }
            this.pathFilters = pathFilters;
            return this;
        }
        public Builder pathFilters(GetBuildDefinitionPullRequestTriggerOverridePathFilter... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }
        public GetBuildDefinitionPullRequestTriggerOverride build() {
            final var _resultValue = new GetBuildDefinitionPullRequestTriggerOverride();
            _resultValue.autoCancel = autoCancel;
            _resultValue.branchFilters = branchFilters;
            _resultValue.pathFilters = pathFilters;
            return _resultValue;
        }
    }
}
