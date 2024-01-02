// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Core.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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
            if (id == null) {
              throw new MissingRequiredPropertyException("GetClientConfigResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder organizationUrl(String organizationUrl) {
            if (organizationUrl == null) {
              throw new MissingRequiredPropertyException("GetClientConfigResult", "organizationUrl");
            }
            this.organizationUrl = organizationUrl;
            return this;
        }
        public GetClientConfigResult build() {
            final var _resultValue = new GetClientConfigResult();
            _resultValue.id = id;
            _resultValue.organizationUrl = organizationUrl;
            return _resultValue;
        }
    }
}
