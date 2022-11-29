// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs;
import com.pulumi.azuredevops.inputs.BuildDefinitionPullRequestTriggerOverridePathFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildDefinitionPullRequestTriggerOverrideArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionPullRequestTriggerOverrideArgs Empty = new BuildDefinitionPullRequestTriggerOverrideArgs();

    /**
     * . Defaults to `true`.
     * 
     */
    @Import(name="autoCancel")
    private @Nullable Output<Boolean> autoCancel;

    /**
     * @return . Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> autoCancel() {
        return Optional.ofNullable(this.autoCancel);
    }

    /**
     * The branches to include and exclude from the trigger.
     * 
     */
    @Import(name="branchFilters")
    private @Nullable Output<List<BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs>> branchFilters;

    /**
     * @return The branches to include and exclude from the trigger.
     * 
     */
    public Optional<Output<List<BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs>>> branchFilters() {
        return Optional.ofNullable(this.branchFilters);
    }

    /**
     * Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     * 
     */
    @Import(name="pathFilters")
    private @Nullable Output<List<BuildDefinitionPullRequestTriggerOverridePathFilterArgs>> pathFilters;

    /**
     * @return Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
     * 
     */
    public Optional<Output<List<BuildDefinitionPullRequestTriggerOverridePathFilterArgs>>> pathFilters() {
        return Optional.ofNullable(this.pathFilters);
    }

    private BuildDefinitionPullRequestTriggerOverrideArgs() {}

    private BuildDefinitionPullRequestTriggerOverrideArgs(BuildDefinitionPullRequestTriggerOverrideArgs $) {
        this.autoCancel = $.autoCancel;
        this.branchFilters = $.branchFilters;
        this.pathFilters = $.pathFilters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionPullRequestTriggerOverrideArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionPullRequestTriggerOverrideArgs $;

        public Builder() {
            $ = new BuildDefinitionPullRequestTriggerOverrideArgs();
        }

        public Builder(BuildDefinitionPullRequestTriggerOverrideArgs defaults) {
            $ = new BuildDefinitionPullRequestTriggerOverrideArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoCancel . Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoCancel(@Nullable Output<Boolean> autoCancel) {
            $.autoCancel = autoCancel;
            return this;
        }

        /**
         * @param autoCancel . Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoCancel(Boolean autoCancel) {
            return autoCancel(Output.of(autoCancel));
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(@Nullable Output<List<BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs>> branchFilters) {
            $.branchFilters = branchFilters;
            return this;
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(List<BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs> branchFilters) {
            return branchFilters(Output.of(branchFilters));
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }

        /**
         * @param pathFilters Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(@Nullable Output<List<BuildDefinitionPullRequestTriggerOverridePathFilterArgs>> pathFilters) {
            $.pathFilters = pathFilters;
            return this;
        }

        /**
         * @param pathFilters Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(List<BuildDefinitionPullRequestTriggerOverridePathFilterArgs> pathFilters) {
            return pathFilters(Output.of(pathFilters));
        }

        /**
         * @param pathFilters Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(BuildDefinitionPullRequestTriggerOverridePathFilterArgs... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }

        public BuildDefinitionPullRequestTriggerOverrideArgs build() {
            return $;
        }
    }

}
