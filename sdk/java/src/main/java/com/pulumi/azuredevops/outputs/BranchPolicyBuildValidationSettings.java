// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.BranchPolicyBuildValidationSettingsScope;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchPolicyBuildValidationSettings {
    /**
     * @return The ID of the build to monitor for the policy.
     * 
     */
    private Integer buildDefinitionId;
    /**
     * @return The display name for the policy.
     * 
     */
    private String displayName;
    /**
     * @return If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `[&#34;/WebApp/Models/Data.cs&#34;, &#34;/WebApp/*&#34;, &#34;*.cs&#34;]`. Paths prefixed with &#34;!&#34; are excluded. Example: `[&#34;/WebApp/*&#34;, &#34;!/WebApp/Tests/*&#34;]`. Order is significant.
     * 
     */
    private @Nullable List<String> filenamePatterns;
    /**
     * @return If set to true, the build will need to be manually queued. Defaults to `false`
     * 
     */
    private @Nullable Boolean manualQueueOnly;
    /**
     * @return True if the build should queue on source updates only. Defaults to `true`.
     * 
     */
    private @Nullable Boolean queueOnSourceUpdateOnly;
    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    private List<BranchPolicyBuildValidationSettingsScope> scopes;
    /**
     * @return The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
     * 
     */
    private @Nullable Integer validDuration;

    private BranchPolicyBuildValidationSettings() {}
    /**
     * @return The ID of the build to monitor for the policy.
     * 
     */
    public Integer buildDefinitionId() {
        return this.buildDefinitionId;
    }
    /**
     * @return The display name for the policy.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `[&#34;/WebApp/Models/Data.cs&#34;, &#34;/WebApp/*&#34;, &#34;*.cs&#34;]`. Paths prefixed with &#34;!&#34; are excluded. Example: `[&#34;/WebApp/*&#34;, &#34;!/WebApp/Tests/*&#34;]`. Order is significant.
     * 
     */
    public List<String> filenamePatterns() {
        return this.filenamePatterns == null ? List.of() : this.filenamePatterns;
    }
    /**
     * @return If set to true, the build will need to be manually queued. Defaults to `false`
     * 
     */
    public Optional<Boolean> manualQueueOnly() {
        return Optional.ofNullable(this.manualQueueOnly);
    }
    /**
     * @return True if the build should queue on source updates only. Defaults to `true`.
     * 
     */
    public Optional<Boolean> queueOnSourceUpdateOnly() {
        return Optional.ofNullable(this.queueOnSourceUpdateOnly);
    }
    /**
     * @return Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
     * 
     */
    public List<BranchPolicyBuildValidationSettingsScope> scopes() {
        return this.scopes;
    }
    /**
     * @return The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
     * 
     */
    public Optional<Integer> validDuration() {
        return Optional.ofNullable(this.validDuration);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchPolicyBuildValidationSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer buildDefinitionId;
        private String displayName;
        private @Nullable List<String> filenamePatterns;
        private @Nullable Boolean manualQueueOnly;
        private @Nullable Boolean queueOnSourceUpdateOnly;
        private List<BranchPolicyBuildValidationSettingsScope> scopes;
        private @Nullable Integer validDuration;
        public Builder() {}
        public Builder(BranchPolicyBuildValidationSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.buildDefinitionId = defaults.buildDefinitionId;
    	      this.displayName = defaults.displayName;
    	      this.filenamePatterns = defaults.filenamePatterns;
    	      this.manualQueueOnly = defaults.manualQueueOnly;
    	      this.queueOnSourceUpdateOnly = defaults.queueOnSourceUpdateOnly;
    	      this.scopes = defaults.scopes;
    	      this.validDuration = defaults.validDuration;
        }

        @CustomType.Setter
        public Builder buildDefinitionId(Integer buildDefinitionId) {
            this.buildDefinitionId = Objects.requireNonNull(buildDefinitionId);
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            this.displayName = Objects.requireNonNull(displayName);
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
        public Builder manualQueueOnly(@Nullable Boolean manualQueueOnly) {
            this.manualQueueOnly = manualQueueOnly;
            return this;
        }
        @CustomType.Setter
        public Builder queueOnSourceUpdateOnly(@Nullable Boolean queueOnSourceUpdateOnly) {
            this.queueOnSourceUpdateOnly = queueOnSourceUpdateOnly;
            return this;
        }
        @CustomType.Setter
        public Builder scopes(List<BranchPolicyBuildValidationSettingsScope> scopes) {
            this.scopes = Objects.requireNonNull(scopes);
            return this;
        }
        public Builder scopes(BranchPolicyBuildValidationSettingsScope... scopes) {
            return scopes(List.of(scopes));
        }
        @CustomType.Setter
        public Builder validDuration(@Nullable Integer validDuration) {
            this.validDuration = validDuration;
            return this;
        }
        public BranchPolicyBuildValidationSettings build() {
            final var o = new BranchPolicyBuildValidationSettings();
            o.buildDefinitionId = buildDefinitionId;
            o.displayName = displayName;
            o.filenamePatterns = filenamePatterns;
            o.manualQueueOnly = manualQueueOnly;
            o.queueOnSourceUpdateOnly = queueOnSourceUpdateOnly;
            o.scopes = scopes;
            o.validDuration = validDuration;
            return o;
        }
    }
}
