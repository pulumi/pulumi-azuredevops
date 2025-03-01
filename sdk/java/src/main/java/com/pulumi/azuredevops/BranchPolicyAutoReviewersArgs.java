// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.BranchPolicyAutoReviewersSettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchPolicyAutoReviewersArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyAutoReviewersArgs Empty = new BranchPolicyAutoReviewersArgs();

    /**
     * A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms &#34;optional&#34; and &#34;required&#34; reviewers. Defaults to `true`.
     * 
     */
    @Import(name="blocking")
    private @Nullable Output<Boolean> blocking;

    /**
     * @return A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms &#34;optional&#34; and &#34;required&#34; reviewers. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> blocking() {
        return Optional.ofNullable(this.blocking);
    }

    /**
     * A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return A flag indicating if the policy should be enabled. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The ID of the project in which the policy will be created.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project in which the policy will be created.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     * 
     */
    @Import(name="settings", required=true)
    private Output<BranchPolicyAutoReviewersSettingsArgs> settings;

    /**
     * @return A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
     * 
     */
    public Output<BranchPolicyAutoReviewersSettingsArgs> settings() {
        return this.settings;
    }

    private BranchPolicyAutoReviewersArgs() {}

    private BranchPolicyAutoReviewersArgs(BranchPolicyAutoReviewersArgs $) {
        this.blocking = $.blocking;
        this.enabled = $.enabled;
        this.projectId = $.projectId;
        this.settings = $.settings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyAutoReviewersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyAutoReviewersArgs $;

        public Builder() {
            $ = new BranchPolicyAutoReviewersArgs();
        }

        public Builder(BranchPolicyAutoReviewersArgs defaults) {
            $ = new BranchPolicyAutoReviewersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blocking A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms &#34;optional&#34; and &#34;required&#34; reviewers. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder blocking(@Nullable Output<Boolean> blocking) {
            $.blocking = blocking;
            return this;
        }

        /**
         * @param blocking A flag indicating if the policy should be blocking. This relates to the Azure DevOps terms &#34;optional&#34; and &#34;required&#34; reviewers. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder blocking(Boolean blocking) {
            return blocking(Output.of(blocking));
        }

        /**
         * @param enabled A flag indicating if the policy should be enabled. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled A flag indicating if the policy should be enabled. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param projectId The ID of the project in which the policy will be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project in which the policy will be created.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param settings A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
         * 
         * @return builder
         * 
         */
        public Builder settings(Output<BranchPolicyAutoReviewersSettingsArgs> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
         * 
         * @return builder
         * 
         */
        public Builder settings(BranchPolicyAutoReviewersSettingsArgs settings) {
            return settings(Output.of(settings));
        }

        public BranchPolicyAutoReviewersArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("BranchPolicyAutoReviewersArgs", "projectId");
            }
            if ($.settings == null) {
                throw new MissingRequiredPropertyException("BranchPolicyAutoReviewersArgs", "settings");
            }
            return $;
        }
    }

}
