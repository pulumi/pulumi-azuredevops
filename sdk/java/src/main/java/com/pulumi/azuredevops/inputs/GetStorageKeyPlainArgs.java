// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetStorageKeyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetStorageKeyPlainArgs Empty = new GetStorageKeyPlainArgs();

    /**
     * The descriptor that will be resolved to a storage key.
     * 
     */
    @Import(name="descriptor", required=true)
    private String descriptor;

    /**
     * @return The descriptor that will be resolved to a storage key.
     * 
     */
    public String descriptor() {
        return this.descriptor;
    }

    private GetStorageKeyPlainArgs() {}

    private GetStorageKeyPlainArgs(GetStorageKeyPlainArgs $) {
        this.descriptor = $.descriptor;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetStorageKeyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetStorageKeyPlainArgs $;

        public Builder() {
            $ = new GetStorageKeyPlainArgs();
        }

        public Builder(GetStorageKeyPlainArgs defaults) {
            $ = new GetStorageKeyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param descriptor The descriptor that will be resolved to a storage key.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(String descriptor) {
            $.descriptor = descriptor;
            return this;
        }

        public GetStorageKeyPlainArgs build() {
            if ($.descriptor == null) {
                throw new MissingRequiredPropertyException("GetStorageKeyPlainArgs", "descriptor");
            }
            return $;
        }
    }

}
