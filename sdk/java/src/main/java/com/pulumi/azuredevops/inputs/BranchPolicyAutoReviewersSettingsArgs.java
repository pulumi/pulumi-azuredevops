// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.BranchPolicyAutoReviewersSettingsScopeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchPolicyAutoReviewersSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchPolicyAutoReviewersSettingsArgs Empty = new BranchPolicyAutoReviewersSettingsArgs();

    /**
     * Required reviewers ids. Supports multiples user Ids.
     * 
     */
    @Import(name="autoReviewerIds", required=true)
    private Output<List<String>> autoReviewerIds;

    /**
     * @return Required reviewers ids. Supports multiples user Ids.
     * 
     */
    public Output<List<String>> autoReviewerIds() {
        return this.autoReviewerIds;
    }

    /**
     * Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
     * 
     */
    @Import(name="message")
    private @Nullable Output<String> message;

    /**
     * @return Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
     * 
     */
    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    /**
     * Minimum number of required reviewers. Defaults to `1`.
     * 
     * &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
     * 
     */
    @Import(name="minimumNumberOfReviewers")
    private @Nullable Output<Integer> minimumNumberOfReviewers;

    /**
     * @return Minimum number of required reviewers. Defaults to `1`.
     * 
     * &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
     * 
     */
    public Optional<Output<Integer>> minimumNumberOfReviewers() {
        return Optional.ofNullable(this.minimumNumberOfReviewers);
    }

    /**
     * Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
     * 
     */
    @Import(name="pathFilters")
    private @Nullable Output<List<String>> pathFilters;

    /**
     * @return Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
     * 
     */
    public Optional<Output<List<String>>> pathFilters() {
        return Optional.ofNullable(this.pathFilters);
    }

    /**
     * A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    @Import(name="scopes", required=true)
    private Output<List<BranchPolicyAutoReviewersSettingsScopeArgs>> scopes;

    /**
     * @return A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public Output<List<BranchPolicyAutoReviewersSettingsScopeArgs>> scopes() {
        return this.scopes;
    }

    /**
     * Controls whether or not the submitter&#39;s vote counts. Defaults to `false`.
     * 
     */
    @Import(name="submitterCanVote")
    private @Nullable Output<Boolean> submitterCanVote;

    /**
     * @return Controls whether or not the submitter&#39;s vote counts. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> submitterCanVote() {
        return Optional.ofNullable(this.submitterCanVote);
    }

    private BranchPolicyAutoReviewersSettingsArgs() {}

    private BranchPolicyAutoReviewersSettingsArgs(BranchPolicyAutoReviewersSettingsArgs $) {
        this.autoReviewerIds = $.autoReviewerIds;
        this.message = $.message;
        this.minimumNumberOfReviewers = $.minimumNumberOfReviewers;
        this.pathFilters = $.pathFilters;
        this.scopes = $.scopes;
        this.submitterCanVote = $.submitterCanVote;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchPolicyAutoReviewersSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchPolicyAutoReviewersSettingsArgs $;

        public Builder() {
            $ = new BranchPolicyAutoReviewersSettingsArgs();
        }

        public Builder(BranchPolicyAutoReviewersSettingsArgs defaults) {
            $ = new BranchPolicyAutoReviewersSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoReviewerIds Required reviewers ids. Supports multiples user Ids.
         * 
         * @return builder
         * 
         */
        public Builder autoReviewerIds(Output<List<String>> autoReviewerIds) {
            $.autoReviewerIds = autoReviewerIds;
            return this;
        }

        /**
         * @param autoReviewerIds Required reviewers ids. Supports multiples user Ids.
         * 
         * @return builder
         * 
         */
        public Builder autoReviewerIds(List<String> autoReviewerIds) {
            return autoReviewerIds(Output.of(autoReviewerIds));
        }

        /**
         * @param autoReviewerIds Required reviewers ids. Supports multiples user Ids.
         * 
         * @return builder
         * 
         */
        public Builder autoReviewerIds(String... autoReviewerIds) {
            return autoReviewerIds(List.of(autoReviewerIds));
        }

        /**
         * @param message Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
         * 
         * @return builder
         * 
         */
        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        /**
         * @param message Activity feed message, Message will appear in the activity feed of pull requests with automatically added reviewers.
         * 
         * @return builder
         * 
         */
        public Builder message(String message) {
            return message(Output.of(message));
        }

        /**
         * @param minimumNumberOfReviewers Minimum number of required reviewers. Defaults to `1`.
         * 
         * &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder minimumNumberOfReviewers(@Nullable Output<Integer> minimumNumberOfReviewers) {
            $.minimumNumberOfReviewers = minimumNumberOfReviewers;
            return this;
        }

        /**
         * @param minimumNumberOfReviewers Minimum number of required reviewers. Defaults to `1`.
         * 
         * &gt; **Note** Has to be greater than `0`. Can only be greater than `1` when attribute `auto_reviewer_ids` contains exactly one group! Only has an effect when attribute `blocking` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder minimumNumberOfReviewers(Integer minimumNumberOfReviewers) {
            return minimumNumberOfReviewers(Output.of(minimumNumberOfReviewers));
        }

        /**
         * @param pathFilters Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(@Nullable Output<List<String>> pathFilters) {
            $.pathFilters = pathFilters;
            return this;
        }

        /**
         * @param pathFilters Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(List<String> pathFilters) {
            return pathFilters(Output.of(pathFilters));
        }

        /**
         * @param pathFilters Filter path(s) on which the policy is applied. Supports absolute paths, wildcards and multiple paths. Example: /WebApp/Models/Data.cs, /WebApp/* or *.cs,/WebApp/Models/Data.cs;ClientApp/Models/Data.cs.
         * 
         * @return builder
         * 
         */
        public Builder pathFilters(String... pathFilters) {
            return pathFilters(List.of(pathFilters));
        }

        /**
         * @param scopes A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(Output<List<BranchPolicyAutoReviewersSettingsScopeArgs>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<BranchPolicyAutoReviewersSettingsScopeArgs> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
         * 
         * @return builder
         * 
         */
        public Builder scopes(BranchPolicyAutoReviewersSettingsScopeArgs... scopes) {
            return scopes(List.of(scopes));
        }

        /**
         * @param submitterCanVote Controls whether or not the submitter&#39;s vote counts. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder submitterCanVote(@Nullable Output<Boolean> submitterCanVote) {
            $.submitterCanVote = submitterCanVote;
            return this;
        }

        /**
         * @param submitterCanVote Controls whether or not the submitter&#39;s vote counts. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder submitterCanVote(Boolean submitterCanVote) {
            return submitterCanVote(Output.of(submitterCanVote));
        }

        public BranchPolicyAutoReviewersSettingsArgs build() {
            if ($.autoReviewerIds == null) {
                throw new MissingRequiredPropertyException("BranchPolicyAutoReviewersSettingsArgs", "autoReviewerIds");
            }
            if ($.scopes == null) {
                throw new MissingRequiredPropertyException("BranchPolicyAutoReviewersSettingsArgs", "scopes");
            }
            return $;
        }
    }

}
