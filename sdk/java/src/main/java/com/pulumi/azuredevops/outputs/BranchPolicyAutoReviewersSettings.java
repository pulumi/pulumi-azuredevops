// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.BranchPolicyAutoReviewersSettingsScope;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchPolicyAutoReviewersSettings {
    /**
     * @return Required reviewers ids. Supports multiples user Ids.
     * 
     */
    private List<String> autoReviewerIds;
    /**
     * @return Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
     * 
     */
    private @Nullable String message;
    /**
     * @return Minimum number of required reviewers. Defaults to `1`.
     * 
     * &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
     * 
     */
    private @Nullable Integer minimumNumberOfReviewers;
    /**
     * @return Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
     * 
     */
    private @Nullable List<String> pathFilters;
    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    private List<BranchPolicyAutoReviewersSettingsScope> scopes;
    /**
     * @return Controls whether or not the submitter&#39;s vote counts. Defaults to `false`.
     * 
     */
    private @Nullable Boolean submitterCanVote;

    private BranchPolicyAutoReviewersSettings() {}
    /**
     * @return Required reviewers ids. Supports multiples user Ids.
     * 
     */
    public List<String> autoReviewerIds() {
        return this.autoReviewerIds;
    }
    /**
     * @return Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
     * 
     */
    public Optional<String> message() {
        return Optional.ofNullable(this.message);
    }
    /**
     * @return Minimum number of required reviewers. Defaults to `1`.
     * 
     * &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
     * 
     */
    public Optional<Integer> minimumNumberOfReviewers() {
        return Optional.ofNullable(this.minimumNumberOfReviewers);
    }
    /**
     * @return Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
     * 
     */
    public List<String> pathFilters() {
        return this.pathFilters == null ? List.of() : this.pathFilters;
    }
    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public List<BranchPolicyAutoReviewersSettingsScope> scopes() {
        return this.scopes;
    }
    /**
     * @return Controls whether or not the submitter&#39;s vote counts. Defaults to `false`.
     * 
     */
    public Optional<Boolean> submitterCanVote() {
        return Optional.ofNullable(this.submitterCanVote);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchPolicyAutoReviewersSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> autoReviewerIds;
        private @Nullable String message;
        private @Nullable Integer minimumNumberOfReviewers;
        private @Nullable List<String> pathFilters;
        private List<BranchPolicyAutoReviewersSettingsScope> scopes;
        private @Nullable Boolean submitterCanVote;
        public Builder() {}
        public Builder(BranchPolicyAutoReviewersSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoReviewerIds = defaults.autoReviewerIds;
    	      this.message = defaults.message;
    	      this.minimumNumberOfReviewers = defaults.minimumNumberOfReviewers;
    	      this.pathFilters = defaults.pathFilters;
    	      this.scopes = defaults.scopes;
    	      this.submitterCanVote = defaults.submitterCanVote;
        }

        @CustomType.Setter
        public Builder autoReviewerIds(List<String> autoReviewerIds) {
            this.autoReviewerIds = Objects.requireNonNull(autoReviewerIds);
            return this;
        }
        public Builder autoReviewerIds(String... autoReviewerIds) {
            return autoReviewerIds(List.of(autoReviewerIds));
        }
        @CustomType.Setter
        public Builder message(@Nullable String message) {
            this.message = message;
            return this;
        }
        @CustomType.Setter
        public Builder minimumNumberOfReviewers(@Nullable Integer minimumNumberOfReviewers) {
            this.minimumNumberOfReviewers = minimumNumberOfReviewers;
            return this;
        }
        @CustomType.Setter
        public Builder pathFilters(@Nullable List<String> pathFilters) {
            this.pathFilters = pathFilters;
            return this;
        }
        public Builder pathFilters(String... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }
        @CustomType.Setter
        public Builder scopes(List<BranchPolicyAutoReviewersSettingsScope> scopes) {
            this.scopes = Objects.requireNonNull(scopes);
            return this;
        }
        public Builder scopes(BranchPolicyAutoReviewersSettingsScope... scopes) {
            return scopes(List.of(scopes));
        }
        @CustomType.Setter
        public Builder submitterCanVote(@Nullable Boolean submitterCanVote) {
            this.submitterCanVote = submitterCanVote;
            return this;
        }
        public BranchPolicyAutoReviewersSettings build() {
            final var o = new BranchPolicyAutoReviewersSettings();
            o.autoReviewerIds = autoReviewerIds;
            o.message = message;
            o.minimumNumberOfReviewers = minimumNumberOfReviewers;
            o.pathFilters = pathFilters;
            o.scopes = scopes;
            o.submitterCanVote = submitterCanVote;
            return o;
        }
    }
}
