// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.BranchPolicyCommentResolutionSettingsScope;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.List;
import java.util.Objects;

@CustomType
public final class BranchPolicyCommentResolutionSettings {
    /**
     * @return A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    private List<BranchPolicyCommentResolutionSettingsScope> scopes;

    private BranchPolicyCommentResolutionSettings() {}
    /**
     * @return A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public List<BranchPolicyCommentResolutionSettingsScope> scopes() {
        return this.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchPolicyCommentResolutionSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<BranchPolicyCommentResolutionSettingsScope> scopes;
        public Builder() {}
        public Builder(BranchPolicyCommentResolutionSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.scopes = defaults.scopes;
        }

        @CustomType.Setter
        public Builder scopes(List<BranchPolicyCommentResolutionSettingsScope> scopes) {
            if (scopes == null) {
              throw new MissingRequiredPropertyException("BranchPolicyCommentResolutionSettings", "scopes");
            }
            this.scopes = scopes;
            return this;
        }
        public Builder scopes(BranchPolicyCommentResolutionSettingsScope... scopes) {
            return scopes(List.of(scopes));
        }
        public BranchPolicyCommentResolutionSettings build() {
            final var _resultValue = new BranchPolicyCommentResolutionSettings();
            _resultValue.scopes = scopes;
            return _resultValue;
        }
    }
}
