// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionVariable {
    /**
     * @return `true` if the variable can be overridden.
     * 
     */
    private Boolean allowOverride;
    /**
     * @return `true` if the variable is a secret.
     * 
     */
    private Boolean isSecret;
    /**
     * @return The name of this Build Definition.
     * 
     */
    private String name;
    /**
     * @return The secret value of the variable.
     * 
     */
    private String secretValue;
    /**
     * @return The value of the variable.
     * 
     */
    private String value;

    private GetBuildDefinitionVariable() {}
    /**
     * @return `true` if the variable can be overridden.
     * 
     */
    public Boolean allowOverride() {
        return this.allowOverride;
    }
    /**
     * @return `true` if the variable is a secret.
     * 
     */
    public Boolean isSecret() {
        return this.isSecret;
    }
    /**
     * @return The name of this Build Definition.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The secret value of the variable.
     * 
     */
    public String secretValue() {
        return this.secretValue;
    }
    /**
     * @return The value of the variable.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBuildDefinitionVariable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean allowOverride;
        private Boolean isSecret;
        private String name;
        private String secretValue;
        private String value;
        public Builder() {}
        public Builder(GetBuildDefinitionVariable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowOverride = defaults.allowOverride;
    	      this.isSecret = defaults.isSecret;
    	      this.name = defaults.name;
    	      this.secretValue = defaults.secretValue;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder allowOverride(Boolean allowOverride) {
            if (allowOverride == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionVariable", "allowOverride");
            }
            this.allowOverride = allowOverride;
            return this;
        }
        @CustomType.Setter
        public Builder isSecret(Boolean isSecret) {
            if (isSecret == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionVariable", "isSecret");
            }
            this.isSecret = isSecret;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionVariable", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder secretValue(String secretValue) {
            if (secretValue == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionVariable", "secretValue");
            }
            this.secretValue = secretValue;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("GetBuildDefinitionVariable", "value");
            }
            this.value = value;
            return this;
        }
        public GetBuildDefinitionVariable build() {
            final var _resultValue = new GetBuildDefinitionVariable();
            _resultValue.allowOverride = allowOverride;
            _resultValue.isSecret = isSecret;
            _resultValue.name = name;
            _resultValue.secretValue = secretValue;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
