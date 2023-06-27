// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.BranchPolicyMinReviewersSettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchPolicyMinReviewersArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyMinReviewersArgs Empty = new BranchPolicyMinReviewersArgs();

    /**
     * A flag indicating if the policy should be blocking. Defaults to `true`.
     * 
     */
    @Import(name="blocking")
    private @Nullable Output<Boolean> blocking;

    /**
     * @return A flag indicating if the policy should be blocking. Defaults to `true`.
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
     * A `settings` block as defined below.. This block must be defined exactly once.
     * 
     */
    @Import(name="settings", required=true)
    private Output<BranchPolicyMinReviewersSettingsArgs> settings;

    /**
     * @return A `settings` block as defined below.. This block must be defined exactly once.
     * 
     */
    public Output<BranchPolicyMinReviewersSettingsArgs> settings() {
        return this.settings;
    }

    private BranchPolicyMinReviewersArgs() {}

    private BranchPolicyMinReviewersArgs(BranchPolicyMinReviewersArgs $) {
        this.blocking = $.blocking;
        this.enabled = $.enabled;
        this.projectId = $.projectId;
        this.settings = $.settings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyMinReviewersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyMinReviewersArgs $;

        public Builder() {
            $ = new BranchPolicyMinReviewersArgs();
        }

        public Builder(BranchPolicyMinReviewersArgs defaults) {
            $ = new BranchPolicyMinReviewersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blocking A flag indicating if the policy should be blocking. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder blocking(@Nullable Output<Boolean> blocking) {
            $.blocking = blocking;
            return this;
        }

        /**
         * @param blocking A flag indicating if the policy should be blocking. Defaults to `true`.
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
         * @param settings A `settings` block as defined below.. This block must be defined exactly once.
         * 
         * @return builder
         * 
         */
        public Builder settings(Output<BranchPolicyMinReviewersSettingsArgs> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings A `settings` block as defined below.. This block must be defined exactly once.
         * 
         * @return builder
         * 
         */
        public Builder settings(BranchPolicyMinReviewersSettingsArgs settings) {
            return settings(Output.of(settings));
        }

        public BranchPolicyMinReviewersArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.settings = Objects.requireNonNull($.settings, "expected parameter 'settings' to be non-null");
            return $;
        }
    }

}
