// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPoolArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPoolArgs Empty = new GetPoolArgs();

    /**
     * Name of the Agent Pool.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the Agent Pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetPoolArgs() {}

    private GetPoolArgs(GetPoolArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPoolArgs $;

        public Builder() {
            $ = new GetPoolArgs();
        }

        public Builder(GetPoolArgs defaults) {
            $ = new GetPoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the Agent Pool.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the Agent Pool.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetPoolArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}