// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionPullRequestTriggerFork {
    /**
     * @return Build pull requests from forks of this repository.
     * 
     */
    private Boolean enabled;
    /**
     * @return Make secrets available to builds of forks.
     * 
     */
    private Boolean shareSecrets;

    private GetBuildDefinitionPullRequestTriggerFork() {}
    /**
     * @return Build pull requests from forks of this repository.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return Make secrets available to builds of forks.
     * 
     */
    public Boolean shareSecrets() {
        return this.shareSecrets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionPullRequestTriggerFork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enabled;
        private Boolean shareSecrets;
        public Builder() {}
        public Builder(GetBuildDefinitionPullRequestTriggerFork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.shareSecrets = defaults.shareSecrets;
        }

        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder shareSecrets(Boolean shareSecrets) {
            this.shareSecrets = Objects.requireNonNull(shareSecrets);
            return this;
        }
        public GetBuildDefinitionPullRequestTriggerFork build() {
            final var _resultValue = new GetBuildDefinitionPullRequestTriggerFork();
            _resultValue.enabled = enabled;
            _resultValue.shareSecrets = shareSecrets;
            return _resultValue;
        }
    }
}
