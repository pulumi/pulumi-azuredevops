// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchPolicyMinReviewersSettingsScope {
    /**
     * @return The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     * 
     */
    private @Nullable String matchType;
    /**
     * @return The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
     * 
     */
    private @Nullable String repositoryId;
    /**
     * @return The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     * 
     */
    private @Nullable String repositoryRef;

    private BranchPolicyMinReviewersSettingsScope() {}
    /**
     * @return The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
     * 
     */
    public Optional<String> matchType() {
        return Optional.ofNullable(this.matchType);
    }
    /**
     * @return The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
     * 
     */
    public Optional<String> repositoryId() {
        return Optional.ofNullable(this.repositoryId);
    }
    /**
     * @return The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
     * 
     */
    public Optional<String> repositoryRef() {
        return Optional.ofNullable(this.repositoryRef);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchPolicyMinReviewersSettingsScope defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String matchType;
        private @Nullable String repositoryId;
        private @Nullable String repositoryRef;
        public Builder() {}
        public Builder(BranchPolicyMinReviewersSettingsScope defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.matchType = defaults.matchType;
    	      this.repositoryId = defaults.repositoryId;
    	      this.repositoryRef = defaults.repositoryRef;
        }

        @CustomType.Setter
        public Builder matchType(@Nullable String matchType) {
            this.matchType = matchType;
            return this;
        }
        @CustomType.Setter
        public Builder repositoryId(@Nullable String repositoryId) {
            this.repositoryId = repositoryId;
            return this;
        }
        @CustomType.Setter
        public Builder repositoryRef(@Nullable String repositoryRef) {
            this.repositoryRef = repositoryRef;
            return this;
        }
        public BranchPolicyMinReviewersSettingsScope build() {
            final var _resultValue = new BranchPolicyMinReviewersSettingsScope();
            _resultValue.matchType = matchType;
            _resultValue.repositoryId = repositoryId;
            _resultValue.repositoryRef = repositoryRef;
            return _resultValue;
        }
    }
}
