// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetBuildDefinitionPullRequestTriggerFork;
import com.pulumi.azuredevops.outputs.GetBuildDefinitionPullRequestTriggerOverride;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionPullRequestTrigger {
    /**
     * @return Is a comment required on the PR?
     * 
     */
    private String commentRequired;
    /**
     * @return A `forks` block as defined above.
     * 
     */
    private List<GetBuildDefinitionPullRequestTriggerFork> forks;
    /**
     * @return When use_yaml is true set this to the name of the branch that the azure-pipelines.yml exists on.
     * 
     */
    private String initialBranch;
    /**
     * @return A `override` block as defined below.
     * 
     */
    private List<GetBuildDefinitionPullRequestTriggerOverride> overrides;
    /**
     * @return Use the azure-pipeline file for the build configuration.
     * 
     */
    private Boolean useYaml;

    private GetBuildDefinitionPullRequestTrigger() {}
    /**
     * @return Is a comment required on the PR?
     * 
     */
    public String commentRequired() {
        return this.commentRequired;
    }
    /**
     * @return A `forks` block as defined above.
     * 
     */
    public List<GetBuildDefinitionPullRequestTriggerFork> forks() {
        return this.forks;
    }
    /**
     * @return When use_yaml is true set this to the name of the branch that the azure-pipelines.yml exists on.
     * 
     */
    public String initialBranch() {
        return this.initialBranch;
    }
    /**
     * @return A `override` block as defined below.
     * 
     */
    public List<GetBuildDefinitionPullRequestTriggerOverride> overrides() {
        return this.overrides;
    }
    /**
     * @return Use the azure-pipeline file for the build configuration.
     * 
     */
    public Boolean useYaml() {
        return this.useYaml;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionPullRequestTrigger defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String commentRequired;
        private List<GetBuildDefinitionPullRequestTriggerFork> forks;
        private String initialBranch;
        private List<GetBuildDefinitionPullRequestTriggerOverride> overrides;
        private Boolean useYaml;
        public Builder() {}
        public Builder(GetBuildDefinitionPullRequestTrigger defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commentRequired = defaults.commentRequired;
    	      this.forks = defaults.forks;
    	      this.initialBranch = defaults.initialBranch;
    	      this.overrides = defaults.overrides;
    	      this.useYaml = defaults.useYaml;
        }

        @CustomType.Setter
        public Builder commentRequired(String commentRequired) {
            if (commentRequired == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTrigger", "commentRequired");
            }
            this.commentRequired = commentRequired;
            return this;
        }
        @CustomType.Setter
        public Builder forks(List<GetBuildDefinitionPullRequestTriggerFork> forks) {
            if (forks == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTrigger", "forks");
            }
            this.forks = forks;
            return this;
        }
        public Builder forks(GetBuildDefinitionPullRequestTriggerFork... forks) {
            return forks(List.of(forks));
        }
        @CustomType.Setter
        public Builder initialBranch(String initialBranch) {
            if (initialBranch == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTrigger", "initialBranch");
            }
            this.initialBranch = initialBranch;
            return this;
        }
        @CustomType.Setter
        public Builder overrides(List<GetBuildDefinitionPullRequestTriggerOverride> overrides) {
            if (overrides == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTrigger", "overrides");
            }
            this.overrides = overrides;
            return this;
        }
        public Builder overrides(GetBuildDefinitionPullRequestTriggerOverride... overrides) {
            return overrides(List.of(overrides));
        }
        @CustomType.Setter
        public Builder useYaml(Boolean useYaml) {
            if (useYaml == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionPullRequestTrigger", "useYaml");
            }
            this.useYaml = useYaml;
            return this;
        }
        public GetBuildDefinitionPullRequestTrigger build() {
            final var _resultValue = new GetBuildDefinitionPullRequestTrigger();
            _resultValue.commentRequired = commentRequired;
            _resultValue.forks = forks;
            _resultValue.initialBranch = initialBranch;
            _resultValue.overrides = overrides;
            _resultValue.useYaml = useYaml;
            return _resultValue;
        }
    }
}
