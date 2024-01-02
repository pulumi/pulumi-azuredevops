// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.BranchPolicyCommentResolutionSettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchPolicyCommentResolutionArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyCommentResolutionArgs Empty = new BranchPolicyCommentResolutionArgs();

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
     * Configuration for the policy. This block must be defined exactly once.
     * 
     */
    @Import(name="settings", required=true)
    private Output<BranchPolicyCommentResolutionSettingsArgs> settings;

    /**
     * @return Configuration for the policy. This block must be defined exactly once.
     * 
     */
    public Output<BranchPolicyCommentResolutionSettingsArgs> settings() {
        return this.settings;
    }

    private BranchPolicyCommentResolutionArgs() {}

    private BranchPolicyCommentResolutionArgs(BranchPolicyCommentResolutionArgs $) {
        this.blocking = $.blocking;
        this.enabled = $.enabled;
        this.projectId = $.projectId;
        this.settings = $.settings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyCommentResolutionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyCommentResolutionArgs $;

        public Builder() {
            $ = new BranchPolicyCommentResolutionArgs();
        }

        public Builder(BranchPolicyCommentResolutionArgs defaults) {
            $ = new BranchPolicyCommentResolutionArgs(Objects.requireNonNull(defaults));
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
         * @param settings Configuration for the policy. This block must be defined exactly once.
         * 
         * @return builder
         * 
         */
        public Builder settings(Output<BranchPolicyCommentResolutionSettingsArgs> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings Configuration for the policy. This block must be defined exactly once.
         * 
         * @return builder
         * 
         */
        public Builder settings(BranchPolicyCommentResolutionSettingsArgs settings) {
            return settings(Output.of(settings));
        }

        public BranchPolicyCommentResolutionArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("BranchPolicyCommentResolutionArgs", "projectId");
            }
            if ($.settings == null) {
                throw new MissingRequiredPropertyException("BranchPolicyCommentResolutionArgs", "settings");
            }
            return $;
        }
    }

}
