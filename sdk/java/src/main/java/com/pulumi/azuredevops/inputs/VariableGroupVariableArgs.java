// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VariableGroupVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final VariableGroupVariableArgs Empty = new VariableGroupVariableArgs();

    @Import(name="contentType")
    private @Nullable Output<String> contentType;

    public Optional<Output<String>> contentType() {
        return Optional.ofNullable(this.contentType);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="expires")
    private @Nullable Output<String> expires;

    public Optional<Output<String>> expires() {
        return Optional.ofNullable(this.expires);
    }

    /**
     * A boolean flag describing if the variable value is sensitive. Defaults to `false`.
     * 
     */
    @Import(name="isSecret")
    private @Nullable Output<Boolean> isSecret;

    /**
     * @return A boolean flag describing if the variable value is sensitive. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> isSecret() {
        return Optional.ofNullable(this.isSecret);
    }

    /**
     * The key value used for the variable. Must be unique within the Variable Group.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The key value used for the variable. Must be unique within the Variable Group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
     * 
     */
    @Import(name="secretValue")
    private @Nullable Output<String> secretValue;

    /**
     * @return The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
     * 
     */
    public Optional<Output<String>> secretValue() {
        return Optional.ofNullable(this.secretValue);
    }

    /**
     * The value of the variable. If omitted, it will default to empty string.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value of the variable. If omitted, it will default to empty string.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private VariableGroupVariableArgs() {}

    private VariableGroupVariableArgs(VariableGroupVariableArgs $) {
        this.contentType = $.contentType;
        this.enabled = $.enabled;
        this.expires = $.expires;
        this.isSecret = $.isSecret;
        this.name = $.name;
        this.secretValue = $.secretValue;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VariableGroupVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VariableGroupVariableArgs $;

        public Builder() {
            $ = new VariableGroupVariableArgs();
        }

        public Builder(VariableGroupVariableArgs defaults) {
            $ = new VariableGroupVariableArgs(Objects.requireNonNull(defaults));
        }

        public Builder contentType(@Nullable Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        public Builder contentType(String contentType) {
            return contentType(Output.of(contentType));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder expires(@Nullable Output<String> expires) {
            $.expires = expires;
            return this;
        }

        public Builder expires(String expires) {
            return expires(Output.of(expires));
        }

        /**
         * @param isSecret A boolean flag describing if the variable value is sensitive. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isSecret(@Nullable Output<Boolean> isSecret) {
            $.isSecret = isSecret;
            return this;
        }

        /**
         * @param isSecret A boolean flag describing if the variable value is sensitive. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isSecret(Boolean isSecret) {
            return isSecret(Output.of(isSecret));
        }

        /**
         * @param name The key value used for the variable. Must be unique within the Variable Group.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The key value used for the variable. Must be unique within the Variable Group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param secretValue The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder secretValue(@Nullable Output<String> secretValue) {
            $.secretValue = secretValue;
            return this;
        }

        /**
         * @param secretValue The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder secretValue(String secretValue) {
            return secretValue(Output.of(secretValue));
        }

        /**
         * @param value The value of the variable. If omitted, it will default to empty string.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of the variable. If omitted, it will default to empty string.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public VariableGroupVariableArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
