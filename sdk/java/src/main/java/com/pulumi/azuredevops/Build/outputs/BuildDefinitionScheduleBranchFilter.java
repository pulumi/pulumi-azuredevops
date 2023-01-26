// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class BuildDefinitionScheduleBranchFilter {
    /**
     * @return List of branch patterns to exclude.
     * 
     */
    private @Nullable List<String> excludes;
    /**
     * @return List of branch patterns to include.
     * 
     */
    private @Nullable List<String> includes;

    private BuildDefinitionScheduleBranchFilter() {}
    /**
     * @return List of branch patterns to exclude.
     * 
     */
    public List<String> excludes() {
        return this.excludes == null ? List.of() : this.excludes;
    }
    /**
     * @return List of branch patterns to include.
     * 
     */
    public List<String> includes() {
        return this.includes == null ? List.of() : this.includes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BuildDefinitionScheduleBranchFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> excludes;
        private @Nullable List<String> includes;
        public Builder() {}
        public Builder(BuildDefinitionScheduleBranchFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.excludes = defaults.excludes;
    	      this.includes = defaults.includes;
        }

        @CustomType.Setter
        public Builder excludes(@Nullable List<String> excludes) {
            this.excludes = excludes;
            return this;
        }
        public Builder excludes(String... excludes) {
            return excludes(List.of(excludes));
        }
        @CustomType.Setter
        public Builder includes(@Nullable List<String> includes) {
            this.includes = includes;
            return this;
        }
        public Builder includes(String... includes) {
            return includes(List.of(includes));
        }
        public BuildDefinitionScheduleBranchFilter build() {
            final var o = new BuildDefinitionScheduleBranchFilter();
            o.excludes = excludes;
            o.includes = includes;
            return o;
        }
    }
}
