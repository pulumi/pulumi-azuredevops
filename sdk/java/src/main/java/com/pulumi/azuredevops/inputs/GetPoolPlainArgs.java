// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPoolPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPoolPlainArgs Empty = new GetPoolPlainArgs();

    /**
     * Name of the Agent Pool.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the Agent Pool.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetPoolPlainArgs() {}

    private GetPoolPlainArgs(GetPoolPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPoolPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPoolPlainArgs $;

        public Builder() {
            $ = new GetPoolPlainArgs();
        }

        public Builder(GetPoolPlainArgs defaults) {
            $ = new GetPoolPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the Agent Pool.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetPoolPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
