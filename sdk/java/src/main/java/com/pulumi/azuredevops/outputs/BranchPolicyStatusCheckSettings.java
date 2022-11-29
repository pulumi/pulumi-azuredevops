// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.BranchPolicyStatusCheckSettingsScope;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchPolicyStatusCheckSettings {
    /**
     * @return Policy applicability. If policy `applicability` is `default`, apply unless &#34;Not Applicable&#34;
     * status is posted to the pull request. If policy `applicability` is `conditional`, policy is applied only after a status
     * is posted to the pull request.
     * 
     */
    private @Nullable String applicability;
    /**
     * @return The authorized user can post the status.
     * 
     */
    private @Nullable String authorId;
    /**
     * @return The display name.
     * 
     */
    private @Nullable String displayName;
    /**
     * @return If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `[&#34;/WebApp/Models/Data.cs&#34;, &#34;/WebApp/*&#34;, &#34;*.cs&#34;]`. Paths prefixed with &#34;!&#34; are excluded. Example: `[&#34;/WebApp/*&#34;, &#34;!/WebApp/Tests/*&#34;]`. Order is significant.
     * 
     */
    private @Nullable List<String> filenamePatterns;
    /**
     * @return The genre of the status to check (see [Microsoft Documentation](https://docs.microsoft.com/en-us/azure/devops/repos/git/pull-request-status?view=azure-devops#status-policy))
     * 
     */
    private @Nullable String genre;
    /**
     * @return Reset status whenever there are new changes.
     * 
     */
    private @Nullable Boolean invalidateOnUpdate;
    /**
     * @return The status name to check.
     * 
     */
    private String name;
    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined
     * at least once.
     * 
     */
    private List<BranchPolicyStatusCheckSettingsScope> scopes;

    private BranchPolicyStatusCheckSettings() {}
    /**
     * @return Policy applicability. If policy `applicability` is `default`, apply unless &#34;Not Applicable&#34;
     * status is posted to the pull request. If policy `applicability` is `conditional`, policy is applied only after a status
     * is posted to the pull request.
     * 
     */
    public Optional<String> applicability() {
        return Optional.ofNullable(this.applicability);
    }
    /**
     * @return The authorized user can post the status.
     * 
     */
    public Optional<String> authorId() {
        return Optional.ofNullable(this.authorId);
    }
    /**
     * @return The display name.
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    /**
     * @return If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `[&#34;/WebApp/Models/Data.cs&#34;, &#34;/WebApp/*&#34;, &#34;*.cs&#34;]`. Paths prefixed with &#34;!&#34; are excluded. Example: `[&#34;/WebApp/*&#34;, &#34;!/WebApp/Tests/*&#34;]`. Order is significant.
     * 
     */
    public List<String> filenamePatterns() {
        return this.filenamePatterns == null ? List.of() : this.filenamePatterns;
    }
    /**
     * @return The genre of the status to check (see [Microsoft Documentation](https://docs.microsoft.com/en-us/azure/devops/repos/git/pull-request-status?view=azure-devops#status-policy))
     * 
     */
    public Optional<String> genre() {
        return Optional.ofNullable(this.genre);
    }
    /**
     * @return Reset status whenever there are new changes.
     * 
     */
    public Optional<Boolean> invalidateOnUpdate() {
        return Optional.ofNullable(this.invalidateOnUpdate);
    }
    /**
     * @return The status name to check.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined
     * at least once.
     * 
     */
    public List<BranchPolicyStatusCheckSettingsScope> scopes() {
        return this.scopes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchPolicyStatusCheckSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String applicability;
        private @Nullable String authorId;
        private @Nullable String displayName;
        private @Nullable List<String> filenamePatterns;
        private @Nullable String genre;
        private @Nullable Boolean invalidateOnUpdate;
        private String name;
        private List<BranchPolicyStatusCheckSettingsScope> scopes;
        public Builder() {}
        public Builder(BranchPolicyStatusCheckSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.applicability = defaults.applicability;
    	      this.authorId = defaults.authorId;
    	      this.displayName = defaults.displayName;
    	      this.filenamePatterns = defaults.filenamePatterns;
    	      this.genre = defaults.genre;
    	      this.invalidateOnUpdate = defaults.invalidateOnUpdate;
    	      this.name = defaults.name;
    	      this.scopes = defaults.scopes;
        }

        @CustomType.Setter
        public Builder applicability(@Nullable String applicability) {
            this.applicability = applicability;
            return this;
        }
        @CustomType.Setter
        public Builder authorId(@Nullable String authorId) {
            this.authorId = authorId;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(@Nullable String displayName) {
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder filenamePatterns(@Nullable List<String> filenamePatterns) {
            this.filenamePatterns = filenamePatterns;
            return this;
        }
        public Builder filenamePatterns(String... filenamePatterns) {
            return filenamePatterns(List.of(filenamePatterns));
        }
        @CustomType.Setter
        public Builder genre(@Nullable String genre) {
            this.genre = genre;
            return this;
        }
        @CustomType.Setter
        public Builder invalidateOnUpdate(@Nullable Boolean invalidateOnUpdate) {
            this.invalidateOnUpdate = invalidateOnUpdate;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder scopes(List<BranchPolicyStatusCheckSettingsScope> scopes) {
            this.scopes = Objects.requireNonNull(scopes);
            return this;
        }
        public Builder scopes(BranchPolicyStatusCheckSettingsScope... scopes) {
            return scopes(List.of(scopes));
        }
        public BranchPolicyStatusCheckSettings build() {
            final var o = new BranchPolicyStatusCheckSettings();
            o.applicability = applicability;
            o.authorId = authorId;
            o.displayName = displayName;
            o.filenamePatterns = filenamePatterns;
            o.genre = genre;
            o.invalidateOnUpdate = invalidateOnUpdate;
            o.name = name;
            o.scopes = scopes;
            return o;
        }
    }
}
