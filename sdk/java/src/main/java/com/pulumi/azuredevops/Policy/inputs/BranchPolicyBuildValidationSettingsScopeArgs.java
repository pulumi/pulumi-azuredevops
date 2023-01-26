// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Policy.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchPolicyBuildValidationSettingsScopeArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyBuildValidationSettingsScopeArgs Empty = new BranchPolicyBuildValidationSettingsScopeArgs();

    /**
     * The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     * 
     */
    @Import(name="matchType")
    private @Nullable Output<String> matchType;

    /**
     * @return The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     * 
     */
    public Optional<Output<String>> matchType() {
        return Optional.ofNullable(this.matchType);
    }

    /**
     * The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
     * 
     */
    @Import(name="repositoryId")
    private @Nullable Output<String> repositoryId;

    /**
     * @return The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
     * 
     */
    public Optional<Output<String>> repositoryId() {
        return Optional.ofNullable(this.repositoryId);
    }

    /**
     * The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     * 
     */
    @Import(name="repositoryRef")
    private @Nullable Output<String> repositoryRef;

    /**
     * @return The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     * 
     */
    public Optional<Output<String>> repositoryRef() {
        return Optional.ofNullable(this.repositoryRef);
    }

    private BranchPolicyBuildValidationSettingsScopeArgs() {}

    private BranchPolicyBuildValidationSettingsScopeArgs(BranchPolicyBuildValidationSettingsScopeArgs $) {
        this.matchType = $.matchType;
        this.repositoryId = $.repositoryId;
        this.repositoryRef = $.repositoryRef;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyBuildValidationSettingsScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyBuildValidationSettingsScopeArgs $;

        public Builder() {
            $ = new BranchPolicyBuildValidationSettingsScopeArgs();
        }

        public Builder(BranchPolicyBuildValidationSettingsScopeArgs defaults) {
            $ = new BranchPolicyBuildValidationSettingsScopeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param matchType The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
         * 
         * @return builder
         * 
         */
        public Builder matchType(@Nullable Output<String> matchType) {
            $.matchType = matchType;
            return this;
        }

        /**
         * @param matchType The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
         * 
         * @return builder
         * 
         */
        public Builder matchType(String matchType) {
            return matchType(Output.of(matchType));
        }

        /**
         * @param repositoryId The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(@Nullable Output<String> repositoryId) {
            $.repositoryId = repositoryId;
            return this;
        }

        /**
         * @param repositoryId The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(String repositoryId) {
            return repositoryId(Output.of(repositoryId));
        }

        /**
         * @param repositoryRef The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
         * 
         * @return builder
         * 
         */
        public Builder repositoryRef(@Nullable Output<String> repositoryRef) {
            $.repositoryRef = repositoryRef;
            return this;
        }

        /**
         * @param repositoryRef The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
         * 
         * @return builder
         * 
         */
        public Builder repositoryRef(String repositoryRef) {
            return repositoryRef(Output.of(repositoryRef));
        }

        public BranchPolicyBuildValidationSettingsScopeArgs build() {
            return $;
        }
    }

}
