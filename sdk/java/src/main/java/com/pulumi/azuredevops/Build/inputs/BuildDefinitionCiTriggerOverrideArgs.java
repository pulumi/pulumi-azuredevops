// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.inputs;

import com.pulumi.azuredevops.Build.inputs.BuildDefinitionCiTriggerOverrideBranchFilterArgs;
import com.pulumi.azuredevops.Build.inputs.BuildDefinitionCiTriggerOverridePathFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildDefinitionCiTriggerOverrideArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionCiTriggerOverrideArgs Empty = new BuildDefinitionCiTriggerOverrideArgs();

    /**
     * If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
     * 
     */
    @Import(name="batch")
    private @Nullable Output<Boolean> batch;

    /**
     * @return If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> batch() {
        return Optional.ofNullable(this.batch);
    }

    /**
     * The branches to include and exclude from the trigger.
     * 
     */
    @Import(name="branchFilters")
    private @Nullable Output<List<BuildDefinitionCiTriggerOverrideBranchFilterArgs>> branchFilters;

    /**
     * @return The branches to include and exclude from the trigger.
     * 
     */
    public Optional<Output<List<BuildDefinitionCiTriggerOverrideBranchFilterArgs>>> branchFilters() {
        return Optional.ofNullable(this.branchFilters);
    }

    /**
     * The number of max builds per branch. Defaults to `1`.
     * 
     */
    @Import(name="maxConcurrentBuildsPerBranch")
    private @Nullable Output<Integer> maxConcurrentBuildsPerBranch;

    /**
     * @return The number of max builds per branch. Defaults to `1`.
     * 
     */
    public Optional<Output<Integer>> maxConcurrentBuildsPerBranch() {
        return Optional.ofNullable(this.maxConcurrentBuildsPerBranch);
    }

    /**
     * Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     * 
     */
    @Import(name="pathFilters")
    private @Nullable Output<List<BuildDefinitionCiTriggerOverridePathFilterArgs>> pathFilters;

    /**
     * @return Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     * 
     */
    public Optional<Output<List<BuildDefinitionCiTriggerOverridePathFilterArgs>>> pathFilters() {
        return Optional.ofNullable(this.pathFilters);
    }

    /**
     * How often the external repository is polled. Defaults to `0`.
     * 
     */
    @Import(name="pollingInterval")
    private @Nullable Output<Integer> pollingInterval;

    /**
     * @return How often the external repository is polled. Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> pollingInterval() {
        return Optional.ofNullable(this.pollingInterval);
    }

    /**
     * This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     * 
     */
    @Import(name="pollingJobId")
    private @Nullable Output<String> pollingJobId;

    /**
     * @return This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
     * 
     */
    public Optional<Output<String>> pollingJobId() {
        return Optional.ofNullable(this.pollingJobId);
    }

    private BuildDefinitionCiTriggerOverrideArgs() {}

    private BuildDefinitionCiTriggerOverrideArgs(BuildDefinitionCiTriggerOverrideArgs $) {
        this.batch = $.batch;
        this.branchFilters = $.branchFilters;
        this.maxConcurrentBuildsPerBranch = $.maxConcurrentBuildsPerBranch;
        this.pathFilters = $.pathFilters;
        this.pollingInterval = $.pollingInterval;
        this.pollingJobId = $.pollingJobId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionCiTriggerOverrideArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionCiTriggerOverrideArgs $;

        public Builder() {
            $ = new BuildDefinitionCiTriggerOverrideArgs();
        }

        public Builder(BuildDefinitionCiTriggerOverrideArgs defaults) {
            $ = new BuildDefinitionCiTriggerOverrideArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param batch If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder batch(@Nullable Output<Boolean> batch) {
            $.batch = batch;
            return this;
        }

        /**
         * @param batch If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder batch(Boolean batch) {
            return batch(Output.of(batch));
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(@Nullable Output<List<BuildDefinitionCiTriggerOverrideBranchFilterArgs>> branchFilters) {
            $.branchFilters = branchFilters;
            return this;
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(List<BuildDefinitionCiTriggerOverrideBranchFilterArgs> branchFilters) {
            return branchFilters(Output.of(branchFilters));
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(BuildDefinitionCiTriggerOverrideBranchFilterArgs... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }

        /**
         * @param maxConcurrentBuildsPerBranch The number of max builds per branch. Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder maxConcurrentBuildsPerBranch(@Nullable Output<Integer> maxConcurrentBuildsPerBranch) {
            $.maxConcurrentBuildsPerBranch = maxConcurrentBuildsPerBranch;
            return this;
        }

        /**
         * @param maxConcurrentBuildsPerBranch The number of max builds per branch. Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder maxConcurrentBuildsPerBranch(Integer maxConcurrentBuildsPerBranch) {
            return maxConcurrentBuildsPerBranch(Output.of(maxConcurrentBuildsPerBranch));
        }

        /**
         * @param pathFilters Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(@Nullable Output<List<BuildDefinitionCiTriggerOverridePathFilterArgs>> pathFilters) {
            $.pathFilters = pathFilters;
            return this;
        }

        /**
         * @param pathFilters Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(List<BuildDefinitionCiTriggerOverridePathFilterArgs> pathFilters) {
            return pathFilters(Output.of(pathFilters));
        }

        /**
         * @param pathFilters Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(BuildDefinitionCiTriggerOverridePathFilterArgs... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }

        /**
         * @param pollingInterval How often the external repository is polled. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder pollingInterval(@Nullable Output<Integer> pollingInterval) {
            $.pollingInterval = pollingInterval;
            return this;
        }

        /**
         * @param pollingInterval How often the external repository is polled. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder pollingInterval(Integer pollingInterval) {
            return pollingInterval(Output.of(pollingInterval));
        }

        /**
         * @param pollingJobId This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
         * 
         * @return builder
         * 
         */
        public Builder pollingJobId(@Nullable Output<String> pollingJobId) {
            $.pollingJobId = pollingJobId;
            return this;
        }

        /**
         * @param pollingJobId This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.
         * 
         * @return builder
         * 
         */
        public Builder pollingJobId(String pollingJobId) {
            return pollingJobId(Output.of(pollingJobId));
        }

        public BuildDefinitionCiTriggerOverrideArgs build() {
            return $;
        }
    }

}
