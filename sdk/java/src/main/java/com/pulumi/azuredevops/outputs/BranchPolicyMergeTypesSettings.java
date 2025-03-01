// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.BranchPolicyMergeTypesSettingsScope;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchPolicyMergeTypesSettings {
    /**
     * @return Allow basic merge with no fast forward. Defaults to `false`.
     * 
     */
    private @Nullable Boolean allowBasicNoFastForward;
    /**
     * @return Allow rebase with fast forward. Defaults to `false`.
     * 
     */
    private @Nullable Boolean allowRebaseAndFastForward;
    /**
     * @return Allow rebase with merge commit. Defaults to `false`.
     * 
     */
    private @Nullable Boolean allowRebaseWithMerge;
    /**
     * @return Allow squash merge. Defaults to `false`
     * 
     */
    private @Nullable Boolean allowSquash;
    /**
     * @return A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    private List<BranchPolicyMergeTypesSettingsScope> scopes;

    private BranchPolicyMergeTypesSettings() {}
    /**
     * @return Allow basic merge with no fast forward. Defaults to `false`.
     * 
     */
    public Optional<Boolean> allowBasicNoFastForward() {
        return Optional.ofNullable(this.allowBasicNoFastForward);
    }
    /**
     * @return Allow rebase with fast forward. Defaults to `false`.
     * 
     */
    public Optional<Boolean> allowRebaseAndFastForward() {
        return Optional.ofNullable(this.allowRebaseAndFastForward);
    }
    /**
     * @return Allow rebase with merge commit. Defaults to `false`.
     * 
     */
    public Optional<Boolean> allowRebaseWithMerge() {
        return Optional.ofNullable(this.allowRebaseWithMerge);
    }
    /**
     * @return Allow squash merge. Defaults to `false`
     * 
     */
    public Optional<Boolean> allowSquash() {
        return Optional.ofNullable(this.allowSquash);
    }
    /**
     * @return A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public List<BranchPolicyMergeTypesSettingsScope> scopes() {
        return this.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchPolicyMergeTypesSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowBasicNoFastForward;
        private @Nullable Boolean allowRebaseAndFastForward;
        private @Nullable Boolean allowRebaseWithMerge;
        private @Nullable Boolean allowSquash;
        private List<BranchPolicyMergeTypesSettingsScope> scopes;
        public Builder() {}
        public Builder(BranchPolicyMergeTypesSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowBasicNoFastForward = defaults.allowBasicNoFastForward;
    	      this.allowRebaseAndFastForward = defaults.allowRebaseAndFastForward;
    	      this.allowRebaseWithMerge = defaults.allowRebaseWithMerge;
    	      this.allowSquash = defaults.allowSquash;
    	      this.scopes = defaults.scopes;
        }

        @CustomType.Setter
        public Builder allowBasicNoFastForward(@Nullable Boolean allowBasicNoFastForward) {

            this.allowBasicNoFastForward = allowBasicNoFastForward;
            return this;
        }
        @CustomType.Setter
        public Builder allowRebaseAndFastForward(@Nullable Boolean allowRebaseAndFastForward) {

            this.allowRebaseAndFastForward = allowRebaseAndFastForward;
            return this;
        }
        @CustomType.Setter
        public Builder allowRebaseWithMerge(@Nullable Boolean allowRebaseWithMerge) {

            this.allowRebaseWithMerge = allowRebaseWithMerge;
            return this;
        }
        @CustomType.Setter
        public Builder allowSquash(@Nullable Boolean allowSquash) {

            this.allowSquash = allowSquash;
            return this;
        }
        @CustomType.Setter
        public Builder scopes(List<BranchPolicyMergeTypesSettingsScope> scopes) {
            if (scopes == null) {
              throw new MissingRequiredPropertyException("BranchPolicyMergeTypesSettings", "scopes");
            }
            this.scopes = scopes;
            return this;
        }
        public Builder scopes(BranchPolicyMergeTypesSettingsScope... scopes) {
            return scopes(List.of(scopes));
        }
        public BranchPolicyMergeTypesSettings build() {
            final var _resultValue = new BranchPolicyMergeTypesSettings();
            _resultValue.allowBasicNoFastForward = allowBasicNoFastForward;
            _resultValue.allowRebaseAndFastForward = allowRebaseAndFastForward;
            _resultValue.allowRebaseWithMerge = allowRebaseWithMerge;
            _resultValue.allowSquash = allowSquash;
            _resultValue.scopes = scopes;
            return _resultValue;
        }
    }
}
