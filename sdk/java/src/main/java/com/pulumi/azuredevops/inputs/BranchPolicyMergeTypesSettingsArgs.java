// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.BranchPolicyMergeTypesSettingsScopeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchPolicyMergeTypesSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyMergeTypesSettingsArgs Empty = new BranchPolicyMergeTypesSettingsArgs();

    /**
     * Allow basic merge with no fast forward. Defaults to `false`.
     * 
     */
    @Import(name="allowBasicNoFastForward")
    private @Nullable Output<Boolean> allowBasicNoFastForward;

    /**
     * @return Allow basic merge with no fast forward. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> allowBasicNoFastForward() {
        return Optional.ofNullable(this.allowBasicNoFastForward);
    }

    /**
     * Allow rebase with fast forward. Defaults to `false`.
     * 
     */
    @Import(name="allowRebaseAndFastForward")
    private @Nullable Output<Boolean> allowRebaseAndFastForward;

    /**
     * @return Allow rebase with fast forward. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> allowRebaseAndFastForward() {
        return Optional.ofNullable(this.allowRebaseAndFastForward);
    }

    /**
     * Allow rebase with merge commit. Defaults to `false`.
     * 
     */
    @Import(name="allowRebaseWithMerge")
    private @Nullable Output<Boolean> allowRebaseWithMerge;

    /**
     * @return Allow rebase with merge commit. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> allowRebaseWithMerge() {
        return Optional.ofNullable(this.allowRebaseWithMerge);
    }

    /**
     * Allow squash merge. Defaults to `false`
     * 
     */
    @Import(name="allowSquash")
    private @Nullable Output<Boolean> allowSquash;

    /**
     * @return Allow squash merge. Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> allowSquash() {
        return Optional.ofNullable(this.allowSquash);
    }

    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<BranchPolicyMergeTypesSettingsScopeArgs>> scopes;

    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public Output<List<BranchPolicyMergeTypesSettingsScopeArgs>> scopes() {
        return this.scopes;
    }

    private BranchPolicyMergeTypesSettingsArgs() {}

    private BranchPolicyMergeTypesSettingsArgs(BranchPolicyMergeTypesSettingsArgs $) {
        this.allowBasicNoFastForward = $.allowBasicNoFastForward;
        this.allowRebaseAndFastForward = $.allowRebaseAndFastForward;
        this.allowRebaseWithMerge = $.allowRebaseWithMerge;
        this.allowSquash = $.allowSquash;
        this.scopes = $.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyMergeTypesSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyMergeTypesSettingsArgs $;

        public Builder() {
            $ = new BranchPolicyMergeTypesSettingsArgs();
        }

        public Builder(BranchPolicyMergeTypesSettingsArgs defaults) {
            $ = new BranchPolicyMergeTypesSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowBasicNoFastForward Allow basic merge with no fast forward. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder allowBasicNoFastForward(@Nullable Output<Boolean> allowBasicNoFastForward) {
            $.allowBasicNoFastForward = allowBasicNoFastForward;
            return this;
        }

        /**
         * @param allowBasicNoFastForward Allow basic merge with no fast forward. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder allowBasicNoFastForward(Boolean allowBasicNoFastForward) {
            return allowBasicNoFastForward(Output.of(allowBasicNoFastForward));
        }

        /**
         * @param allowRebaseAndFastForward Allow rebase with fast forward. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder allowRebaseAndFastForward(@Nullable Output<Boolean> allowRebaseAndFastForward) {
            $.allowRebaseAndFastForward = allowRebaseAndFastForward;
            return this;
        }

        /**
         * @param allowRebaseAndFastForward Allow rebase with fast forward. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder allowRebaseAndFastForward(Boolean allowRebaseAndFastForward) {
            return allowRebaseAndFastForward(Output.of(allowRebaseAndFastForward));
        }

        /**
         * @param allowRebaseWithMerge Allow rebase with merge commit. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder allowRebaseWithMerge(@Nullable Output<Boolean> allowRebaseWithMerge) {
            $.allowRebaseWithMerge = allowRebaseWithMerge;
            return this;
        }

        /**
         * @param allowRebaseWithMerge Allow rebase with merge commit. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder allowRebaseWithMerge(Boolean allowRebaseWithMerge) {
            return allowRebaseWithMerge(Output.of(allowRebaseWithMerge));
        }

        /**
         * @param allowSquash Allow squash merge. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder allowSquash(@Nullable Output<Boolean> allowSquash) {
            $.allowSquash = allowSquash;
            return this;
        }

        /**
         * @param allowSquash Allow squash merge. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder allowSquash(Boolean allowSquash) {
            return allowSquash(Output.of(allowSquash));
        }

        /**
         * @param scopes Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<BranchPolicyMergeTypesSettingsScopeArgs>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<BranchPolicyMergeTypesSettingsScopeArgs> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(BranchPolicyMergeTypesSettingsScopeArgs... scopes) {
            return scopes(List.of(scopes));
        }

        public BranchPolicyMergeTypesSettingsArgs build() {
            $.scopes = Objects.requireNonNull($.scopes, "expected parameter 'scopes' to be non-null");
            return $;
        }
    }

}