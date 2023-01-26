// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildDefinitionPullRequestTriggerOverridePathFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionPullRequestTriggerOverridePathFilterArgs Empty = new BuildDefinitionPullRequestTriggerOverridePathFilterArgs();

    /**
     * List of branch patterns to exclude.
     * 
     */
    @Import(name="excludes")
    private @Nullable Output<List<String>> excludes;

    /**
     * @return List of branch patterns to exclude.
     * 
     */
    public Optional<Output<List<String>>> excludes() {
        return Optional.ofNullable(this.excludes);
    }

    /**
     * List of branch patterns to include.
     * 
     */
    @Import(name="includes")
    private @Nullable Output<List<String>> includes;

    /**
     * @return List of branch patterns to include.
     * 
     */
    public Optional<Output<List<String>>> includes() {
        return Optional.ofNullable(this.includes);
    }

    private BuildDefinitionPullRequestTriggerOverridePathFilterArgs() {}

    private BuildDefinitionPullRequestTriggerOverridePathFilterArgs(BuildDefinitionPullRequestTriggerOverridePathFilterArgs $) {
        this.excludes = $.excludes;
        this.includes = $.includes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionPullRequestTriggerOverridePathFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionPullRequestTriggerOverridePathFilterArgs $;

        public Builder() {
            $ = new BuildDefinitionPullRequestTriggerOverridePathFilterArgs();
        }

        public Builder(BuildDefinitionPullRequestTriggerOverridePathFilterArgs defaults) {
            $ = new BuildDefinitionPullRequestTriggerOverridePathFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param excludes List of branch patterns to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludes(@Nullable Output<List<String>> excludes) {
            $.excludes = excludes;
            return this;
        }

        /**
         * @param excludes List of branch patterns to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludes(List<String> excludes) {
            return excludes(Output.of(excludes));
        }

        /**
         * @param excludes List of branch patterns to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludes(String... excludes) {
            return excludes(List.of(excludes));
        }

        /**
         * @param includes List of branch patterns to include.
         * 
         * @return builder
         * 
         */
        public Builder includes(@Nullable Output<List<String>> includes) {
            $.includes = includes;
            return this;
        }

        /**
         * @param includes List of branch patterns to include.
         * 
         * @return builder
         * 
         */
        public Builder includes(List<String> includes) {
            return includes(Output.of(includes));
        }

        /**
         * @param includes List of branch patterns to include.
         * 
         * @return builder
         * 
         */
        public Builder includes(String... includes) {
            return includes(List.of(includes));
        }

        public BuildDefinitionPullRequestTriggerOverridePathFilterArgs build() {
            return $;
        }
    }

}
