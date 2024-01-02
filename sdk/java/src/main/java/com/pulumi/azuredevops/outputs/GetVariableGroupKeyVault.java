// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetVariableGroupKeyVault {
    /**
     * @return The name of the Variable Group to retrieve.
     * 
     */
    private String name;
    /**
     * @return The id of the Azure subscription endpoint to access the key vault.
     * 
     */
    private String serviceEndpointId;

    private GetVariableGroupKeyVault() {}
    /**
     * @return The name of the Variable Group to retrieve.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The id of the Azure subscription endpoint to access the key vault.
     * 
     */
    public String serviceEndpointId() {
        return this.serviceEndpointId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVariableGroupKeyVault defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String serviceEndpointId;
        public Builder() {}
        public Builder(GetVariableGroupKeyVault defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.serviceEndpointId = defaults.serviceEndpointId;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetVariableGroupKeyVault", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointId(String serviceEndpointId) {
            if (serviceEndpointId == null) {
              throw new MissingRequiredPropertyException("GetVariableGroupKeyVault", "serviceEndpointId");
            }
            this.serviceEndpointId = serviceEndpointId;
            return this;
        }
        public GetVariableGroupKeyVault build() {
            final var _resultValue = new GetVariableGroupKeyVault();
            _resultValue.name = name;
            _resultValue.serviceEndpointId = serviceEndpointId;
            return _resultValue;
        }
    }
}
