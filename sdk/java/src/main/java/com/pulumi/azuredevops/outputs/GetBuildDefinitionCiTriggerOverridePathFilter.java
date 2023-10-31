// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionCiTriggerOverridePathFilter {
    /**
     * @return (Optional) List of path patterns to exclude.
     * 
     */
    private List<String> excludes;
    /**
     * @return (Optional) List of path patterns to include.
     * 
     */
    private List<String> includes;

    private GetBuildDefinitionCiTriggerOverridePathFilter() {}
    /**
     * @return (Optional) List of path patterns to exclude.
     * 
     */
    public List<String> excludes() {
        return this.excludes;
    }
    /**
     * @return (Optional) List of path patterns to include.
     * 
     */
    public List<String> includes() {
        return this.includes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionCiTriggerOverridePathFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> excludes;
        private List<String> includes;
        public Builder() {}
        public Builder(GetBuildDefinitionCiTriggerOverridePathFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.excludes = defaults.excludes;
    	      this.includes = defaults.includes;
        }

        @CustomType.Setter
        public Builder excludes(List<String> excludes) {
            this.excludes = Objects.requireNonNull(excludes);
            return this;
        }
        public Builder excludes(String... excludes) {
            return excludes(List.of(excludes));
        }
        @CustomType.Setter
        public Builder includes(List<String> includes) {
            this.includes = Objects.requireNonNull(includes);
            return this;
        }
        public Builder includes(String... includes) {
            return includes(List.of(includes));
        }
        public GetBuildDefinitionCiTriggerOverridePathFilter build() {
            final var _resultValue = new GetBuildDefinitionCiTriggerOverridePathFilter();
            _resultValue.excludes = excludes;
            _resultValue.includes = includes;
            return _resultValue;
        }
    }
}
