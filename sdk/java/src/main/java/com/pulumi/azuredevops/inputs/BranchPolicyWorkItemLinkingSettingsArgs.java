// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.BranchPolicyWorkItemLinkingSettingsScopeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.List;
import java.util.Objects;


public final class BranchPolicyWorkItemLinkingSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyWorkItemLinkingSettingsArgs Empty = new BranchPolicyWorkItemLinkingSettingsArgs();

    /**
     * Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<BranchPolicyWorkItemLinkingSettingsScopeArgs>> scopes;

    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public Output<List<BranchPolicyWorkItemLinkingSettingsScopeArgs>> scopes() {
        return this.scopes;
    }

    private BranchPolicyWorkItemLinkingSettingsArgs() {}

    private BranchPolicyWorkItemLinkingSettingsArgs(BranchPolicyWorkItemLinkingSettingsArgs $) {
        this.scopes = $.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyWorkItemLinkingSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyWorkItemLinkingSettingsArgs $;

        public Builder() {
            $ = new BranchPolicyWorkItemLinkingSettingsArgs();
        }

        public Builder(BranchPolicyWorkItemLinkingSettingsArgs defaults) {
            $ = new BranchPolicyWorkItemLinkingSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param scopes Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<BranchPolicyWorkItemLinkingSettingsScopeArgs>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<BranchPolicyWorkItemLinkingSettingsScopeArgs> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(BranchPolicyWorkItemLinkingSettingsScopeArgs... scopes) {
            return scopes(List.of(scopes));
        }

        public BranchPolicyWorkItemLinkingSettingsArgs build() {
            if ($.scopes == null) {
                throw new MissingRequiredPropertyException("BranchPolicyWorkItemLinkingSettingsArgs", "scopes");
            }
            return $;
        }
    }

}
