// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetBuildDefinitionCiTriggerOverride;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionCiTrigger {
    /**
     * @return A `override` block as defined below.
     * 
     */
    private List<GetBuildDefinitionCiTriggerOverride> overrides;
    /**
     * @return Use the azure-pipeline file for the build configuration.
     * 
     */
    private Boolean useYaml;

    private GetBuildDefinitionCiTrigger() {}
    /**
     * @return A `override` block as defined below.
     * 
     */
    public List<GetBuildDefinitionCiTriggerOverride> overrides() {
        return this.overrides;
    }
    /**
     * @return Use the azure-pipeline file for the build configuration.
     * 
     */
    public Boolean useYaml() {
        return this.useYaml;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionCiTrigger defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetBuildDefinitionCiTriggerOverride> overrides;
        private Boolean useYaml;
        public Builder() {}
        public Builder(GetBuildDefinitionCiTrigger defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.overrides = defaults.overrides;
    	      this.useYaml = defaults.useYaml;
        }

        @CustomType.Setter
        public Builder overrides(List<GetBuildDefinitionCiTriggerOverride> overrides) {
            if (overrides == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionCiTrigger", "overrides");
            }
            this.overrides = overrides;
            return this;
        }
        public Builder overrides(GetBuildDefinitionCiTriggerOverride... overrides) {
            return overrides(List.of(overrides));
        }
        @CustomType.Setter
        public Builder useYaml(Boolean useYaml) {
            if (useYaml == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionCiTrigger", "useYaml");
            }
            this.useYaml = useYaml;
            return this;
        }
        public GetBuildDefinitionCiTrigger build() {
            final var _resultValue = new GetBuildDefinitionCiTrigger();
            _resultValue.overrides = overrides;
            _resultValue.useYaml = useYaml;
            return _resultValue;
        }
    }
}
