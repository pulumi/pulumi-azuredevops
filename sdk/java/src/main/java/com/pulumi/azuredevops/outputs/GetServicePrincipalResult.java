// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServicePrincipalResult {
    /**
     * @return The descriptor of the Service Principal.
     * 
     */
    private String descriptor;
    private String displayName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The origin of the Service Principal.
     * 
     */
    private String origin;
    /**
     * @return The origin ID of the Service Principal..
     * 
     */
    private String originId;

    private GetServicePrincipalResult() {}
    /**
     * @return The descriptor of the Service Principal.
     * 
     */
    public String descriptor() {
        return this.descriptor;
    }
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The origin of the Service Principal.
     * 
     */
    public String origin() {
        return this.origin;
    }
    /**
     * @return The origin ID of the Service Principal..
     * 
     */
    public String originId() {
        return this.originId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServicePrincipalResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String descriptor;
        private String displayName;
        private String id;
        private String origin;
        private String originId;
        public Builder() {}
        public Builder(GetServicePrincipalResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptor = defaults.descriptor;
    	      this.displayName = defaults.displayName;
    	      this.id = defaults.id;
    	      this.origin = defaults.origin;
    	      this.originId = defaults.originId;
        }

        @CustomType.Setter
        public Builder descriptor(String descriptor) {
            if (descriptor == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "descriptor");
            }
            this.descriptor = descriptor;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder origin(String origin) {
            if (origin == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "origin");
            }
            this.origin = origin;
            return this;
        }
        @CustomType.Setter
        public Builder originId(String originId) {
            if (originId == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "originId");
            }
            this.originId = originId;
            return this;
        }
        public GetServicePrincipalResult build() {
            final var _resultValue = new GetServicePrincipalResult();
            _resultValue.descriptor = descriptor;
            _resultValue.displayName = displayName;
            _resultValue.id = id;
            _resultValue.origin = origin;
            _resultValue.originId = originId;
            return _resultValue;
        }
    }
}
