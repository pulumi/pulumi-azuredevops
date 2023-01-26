// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServiceEndpointGithubResult {
    /**
     * @return Specifies the Authorization Scheme Map.
     * 
     */
    private Map<String,String> authorization;
    /**
     * @return Specifies the description of the Service Endpoint.
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

    private GetServiceEndpointGithubResult() {}
    /**
     * @return Specifies the Authorization Scheme Map.
     * 
     */
    public Map<String,String> authorization() {
        return this.authorization;
    }
    /**
     * @return Specifies the description of the Service Endpoint.
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

    public static Builder builder(GetServiceEndpointGithubResult defaults) {
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
        public Builder(GetServiceEndpointGithubResult defaults) {
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
            this.authorization = Objects.requireNonNull(authorization);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointId(String serviceEndpointId) {
            this.serviceEndpointId = Objects.requireNonNull(serviceEndpointId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceEndpointName(String serviceEndpointName) {
            this.serviceEndpointName = Objects.requireNonNull(serviceEndpointName);
            return this;
        }
        public GetServiceEndpointGithubResult build() {
            final var o = new GetServiceEndpointGithubResult();
            o.authorization = authorization;
            o.description = description;
            o.id = id;
            o.projectId = projectId;
            o.serviceEndpointId = serviceEndpointId;
            o.serviceEndpointName = serviceEndpointName;
            return o;
        }
    }
}
