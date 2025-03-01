// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.BuildDefinitionBuildCompletionTriggerBranchFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;


public final class BuildDefinitionBuildCompletionTriggerArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionBuildCompletionTriggerArgs Empty = new BuildDefinitionBuildCompletionTriggerArgs();

    /**
     * The branches to include and exclude from the trigger. A `branch_filter` block as documented below.
     * 
     */
    @Import(name="branchFilters", required=true)
    private Output<List<BuildDefinitionBuildCompletionTriggerBranchFilterArgs>> branchFilters;

    /**
     * @return The branches to include and exclude from the trigger. A `branch_filter` block as documented below.
     * 
     */
    public Output<List<BuildDefinitionBuildCompletionTriggerBranchFilterArgs>> branchFilters() {
        return this.branchFilters;
    }

    /**
     * The ID of the build pipeline will be triggered.
     * 
     */
    @Import(name="buildDefinitionId", required=true)
    private Output<Integer> buildDefinitionId;

    /**
     * @return The ID of the build pipeline will be triggered.
     * 
     */
    public Output<Integer> buildDefinitionId() {
        return this.buildDefinitionId;
    }

    private BuildDefinitionBuildCompletionTriggerArgs() {}

    private BuildDefinitionBuildCompletionTriggerArgs(BuildDefinitionBuildCompletionTriggerArgs $) {
        this.branchFilters = $.branchFilters;
        this.buildDefinitionId = $.buildDefinitionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionBuildCompletionTriggerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionBuildCompletionTriggerArgs $;

        public Builder() {
            $ = new BuildDefinitionBuildCompletionTriggerArgs();
        }

        public Builder(BuildDefinitionBuildCompletionTriggerArgs defaults) {
            $ = new BuildDefinitionBuildCompletionTriggerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger. A `branch_filter` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(Output<List<BuildDefinitionBuildCompletionTriggerBranchFilterArgs>> branchFilters) {
            $.branchFilters = branchFilters;
            return this;
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger. A `branch_filter` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(List<BuildDefinitionBuildCompletionTriggerBranchFilterArgs> branchFilters) {
            return branchFilters(Output.of(branchFilters));
        }

        /**
         * @param branchFilters The branches to include and exclude from the trigger. A `branch_filter` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder branchFilters(BuildDefinitionBuildCompletionTriggerBranchFilterArgs... branchFilters) {
            return branchFilters(List.of(branchFilters));
        }

        /**
         * @param buildDefinitionId The ID of the build pipeline will be triggered.
         * 
         * @return builder
         * 
         */
        public Builder buildDefinitionId(Output<Integer> buildDefinitionId) {
            $.buildDefinitionId = buildDefinitionId;
            return this;
        }

        /**
         * @param buildDefinitionId The ID of the build pipeline will be triggered.
         * 
         * @return builder
         * 
         */
        public Builder buildDefinitionId(Integer buildDefinitionId) {
            return buildDefinitionId(Output.of(buildDefinitionId));
        }

        public BuildDefinitionBuildCompletionTriggerArgs build() {
            if ($.branchFilters == null) {
                throw new MissingRequiredPropertyException("BuildDefinitionBuildCompletionTriggerArgs", "branchFilters");
            }
            if ($.buildDefinitionId == null) {
                throw new MissingRequiredPropertyException("BuildDefinitionBuildCompletionTriggerArgs", "buildDefinitionId");
            }
            return $;
        }
    }

}
