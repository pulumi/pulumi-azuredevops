// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.outputs;

import com.pulumi.azuredevops.Build.outputs.BuildDefinitionCiTriggerOverride;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BuildDefinitionCiTrigger {
    /**
     * @return Override the azure-pipeline file and use a this configuration for all builds.
     * 
     */
    private @Nullable BuildDefinitionCiTriggerOverride override;
    /**
     * @return Use the azure-pipeline file for the build configuration. Defaults to `false`.
     * 
     */
    private @Nullable Boolean useYaml;

    private BuildDefinitionCiTrigger() {}
    /**
     * @return Override the azure-pipeline file and use a this configuration for all builds.
     * 
     */
    public Optional<BuildDefinitionCiTriggerOverride> override() {
        return Optional.ofNullable(this.override);
    }
    /**
     * @return Use the azure-pipeline file for the build configuration. Defaults to `false`.
     * 
     */
    public Optional<Boolean> useYaml() {
        return Optional.ofNullable(this.useYaml);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BuildDefinitionCiTrigger defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable BuildDefinitionCiTriggerOverride override;
        private @Nullable Boolean useYaml;
        public Builder() {}
        public Builder(BuildDefinitionCiTrigger defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.override = defaults.override;
    	      this.useYaml = defaults.useYaml;
        }

        @CustomType.Setter
        public Builder override(@Nullable BuildDefinitionCiTriggerOverride override) {
            this.override = override;
            return this;
        }
        @CustomType.Setter
        public Builder useYaml(@Nullable Boolean useYaml) {
            this.useYaml = useYaml;
            return this;
        }
        public BuildDefinitionCiTrigger build() {
            final var _resultValue = new BuildDefinitionCiTrigger();
            _resultValue.override = override;
            _resultValue.useYaml = useYaml;
            return _resultValue;
        }
    }
}
