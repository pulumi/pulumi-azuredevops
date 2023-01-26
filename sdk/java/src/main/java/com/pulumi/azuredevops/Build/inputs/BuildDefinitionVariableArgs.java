// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildDefinitionVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionVariableArgs Empty = new BuildDefinitionVariableArgs();

    /**
     * True if the variable can be overridden. Defaults to `true`.
     * 
     */
    @Import(name="allowOverride")
    private @Nullable Output<Boolean> allowOverride;

    /**
     * @return True if the variable can be overridden. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> allowOverride() {
        return Optional.ofNullable(this.allowOverride);
    }

    /**
     * True if the variable is a secret. Defaults to `false`.
     * 
     */
    @Import(name="isSecret")
    private @Nullable Output<Boolean> isSecret;

    /**
     * @return True if the variable is a secret. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> isSecret() {
        return Optional.ofNullable(this.isSecret);
    }

    /**
     * The name of the variable.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the variable.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The secret value of the variable. Used when `is_secret` set to `true`.
     * 
     */
    @Import(name="secretValue")
    private @Nullable Output<String> secretValue;

    /**
     * @return The secret value of the variable. Used when `is_secret` set to `true`.
     * 
     */
    public Optional<Output<String>> secretValue() {
        return Optional.ofNullable(this.secretValue);
    }

    /**
     * The value of the variable.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value of the variable.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private BuildDefinitionVariableArgs() {}

    private BuildDefinitionVariableArgs(BuildDefinitionVariableArgs $) {
        this.allowOverride = $.allowOverride;
        this.isSecret = $.isSecret;
        this.name = $.name;
        this.secretValue = $.secretValue;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionVariableArgs $;

        public Builder() {
            $ = new BuildDefinitionVariableArgs();
        }

        public Builder(BuildDefinitionVariableArgs defaults) {
            $ = new BuildDefinitionVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowOverride True if the variable can be overridden. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder allowOverride(@Nullable Output<Boolean> allowOverride) {
            $.allowOverride = allowOverride;
            return this;
        }

        /**
         * @param allowOverride True if the variable can be overridden. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder allowOverride(Boolean allowOverride) {
            return allowOverride(Output.of(allowOverride));
        }

        /**
         * @param isSecret True if the variable is a secret. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isSecret(@Nullable Output<Boolean> isSecret) {
            $.isSecret = isSecret;
            return this;
        }

        /**
         * @param isSecret True if the variable is a secret. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isSecret(Boolean isSecret) {
            return isSecret(Output.of(isSecret));
        }

        /**
         * @param name The name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param secretValue The secret value of the variable. Used when `is_secret` set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder secretValue(@Nullable Output<String> secretValue) {
            $.secretValue = secretValue;
            return this;
        }

        /**
         * @param secretValue The secret value of the variable. Used when `is_secret` set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder secretValue(String secretValue) {
            return secretValue(Output.of(secretValue));
        }

        /**
         * @param value The value of the variable.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of the variable.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public BuildDefinitionVariableArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
