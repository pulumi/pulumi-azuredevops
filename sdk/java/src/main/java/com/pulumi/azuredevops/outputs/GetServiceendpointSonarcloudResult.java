// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServiceendpointSonarcloudResult {
    /**
     * @return The Authorization scheme.
     * 
     */
    private Map<String,String> authorization;
    /**
     * @return The description of the Service Endpoint.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String projectId;
    private String serviceEndpointId;
    private String serviceEndpointName;

    private GetServiceendpointSonarcloudResult() {}
    /**
     * @return The Authorization scheme.
     * 
     */
    public Map<String,String> authorization() {
        return this.authorization;
    }
    /**
     * @return The description of the Service Endpoint.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String projectId() {
        return this.projectId;
    }
    public String serviceEndpointId() {
        return this.serviceEndpointId;
    }
    public String serviceEndpointName() {
        return this.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceendpointSonarcloudResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> authorization;
        private String description;
        private String id;
        private String projectId;
        private String serviceEndpointId;
        private String serviceEndpointName;
        public Builder() {}
        public Builder(GetServiceendpointSonarcloudResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorization = defaults.authorization;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.serviceEndpointId = defaults.serviceEndpointId;
    	      this.serviceEndpointName = defaults.serviceEndpointName;
        }

        @CustomType.Setter
        public Builder authorization(Map<String,String> authorization) {
            if (authorization == null) {
              throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudResult", "authorization");
            }
            this.authorization = authorization;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointId(String serviceEndpointId) {
            if (serviceEndpointId == null) {
              throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudResult", "serviceEndpointId");
            }
            this.serviceEndpointId = serviceEndpointId;
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointName(String serviceEndpointName) {
            if (serviceEndpointName == null) {
              throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudResult", "serviceEndpointName");
            }
            this.serviceEndpointName = serviceEndpointName;
            return this;
        }
        public GetServiceendpointSonarcloudResult build() {
            final var _resultValue = new GetServiceendpointSonarcloudResult();
            _resultValue.authorization = authorization;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.projectId = projectId;
            _resultValue.serviceEndpointId = serviceEndpointId;
            _resultValue.serviceEndpointName = serviceEndpointName;
            return _resultValue;
        }
    }
}
