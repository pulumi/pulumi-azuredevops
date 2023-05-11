// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClientConfigResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String organizationUrl;

    private GetClientConfigResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String organizationUrl() {
        return this.organizationUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientConfigResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String organizationUrl;
        public Builder() {}
        public Builder(GetClientConfigResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.organizationUrl = defaults.organizationUrl;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder organizationUrl(String organizationUrl) {
            this.organizationUrl = Objects.requireNonNull(organizationUrl);
            return this;
        }
        public GetClientConfigResult build() {
            final var o = new GetClientConfigResult();
            o.id = id;
            o.organizationUrl = organizationUrl;
            return o;
        }
    }
}