// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceEndpointAzureRMCredentials {
    /**
     * @return The service principal certificate. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
     * 
     */
    private @Nullable String serviceprincipalcertificate;
    /**
     * @return The service principal application ID
     * 
     */
    private String serviceprincipalid;
    /**
     * @return The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
     * 
     */
    private @Nullable String serviceprincipalkey;

    private ServiceEndpointAzureRMCredentials() {}
    /**
     * @return The service principal certificate. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
     * 
     */
    public Optional<String> serviceprincipalcertificate() {
        return Optional.ofNullable(this.serviceprincipalcertificate);
    }
    /**
     * @return The service principal application ID
     * 
     */
    public String serviceprincipalid() {
        return this.serviceprincipalid;
    }
    /**
     * @return The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
     * 
     */
    public Optional<String> serviceprincipalkey() {
        return Optional.ofNullable(this.serviceprincipalkey);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointAzureRMCredentials defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String serviceprincipalcertificate;
        private String serviceprincipalid;
        private @Nullable String serviceprincipalkey;
        public Builder() {}
        public Builder(ServiceEndpointAzureRMCredentials defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serviceprincipalcertificate = defaults.serviceprincipalcertificate;
    	      this.serviceprincipalid = defaults.serviceprincipalid;
    	      this.serviceprincipalkey = defaults.serviceprincipalkey;
        }

        @CustomType.Setter
        public Builder serviceprincipalcertificate(@Nullable String serviceprincipalcertificate) {

            this.serviceprincipalcertificate = serviceprincipalcertificate;
            return this;
        }
        @CustomType.Setter
        public Builder serviceprincipalid(String serviceprincipalid) {
            if (serviceprincipalid == null) {
              throw new MissingRequiredPropertyException("ServiceEndpointAzureRMCredentials", "serviceprincipalid");
            }
            this.serviceprincipalid = serviceprincipalid;
            return this;
        }
        @CustomType.Setter
        public Builder serviceprincipalkey(@Nullable String serviceprincipalkey) {

            this.serviceprincipalkey = serviceprincipalkey;
            return this;
        }
        public ServiceEndpointAzureRMCredentials build() {
            final var _resultValue = new ServiceEndpointAzureRMCredentials();
            _resultValue.serviceprincipalcertificate = serviceprincipalcertificate;
            _resultValue.serviceprincipalid = serviceprincipalid;
            _resultValue.serviceprincipalkey = serviceprincipalkey;
            return _resultValue;
        }
    }
}
