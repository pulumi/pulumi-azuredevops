// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceEndpointAzureRMCredentials {
    /**
     * @return The service principal application Id
     * 
     */
    private String serviceprincipalid;
    /**
     * @return The service principal secret.
     * 
     */
    private String serviceprincipalkey;
    private @Nullable String serviceprincipalkeyHash;

    private ServiceEndpointAzureRMCredentials() {}
    /**
     * @return The service principal application Id
     * 
     */
    public String serviceprincipalid() {
        return this.serviceprincipalid;
    }
    /**
     * @return The service principal secret.
     * 
     */
    public String serviceprincipalkey() {
        return this.serviceprincipalkey;
    }
    public Optional<String> serviceprincipalkeyHash() {
        return Optional.ofNullable(this.serviceprincipalkeyHash);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointAzureRMCredentials defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String serviceprincipalid;
        private String serviceprincipalkey;
        private @Nullable String serviceprincipalkeyHash;
        public Builder() {}
        public Builder(ServiceEndpointAzureRMCredentials defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serviceprincipalid = defaults.serviceprincipalid;
    	      this.serviceprincipalkey = defaults.serviceprincipalkey;
    	      this.serviceprincipalkeyHash = defaults.serviceprincipalkeyHash;
        }

        @CustomType.Setter
        public Builder serviceprincipalid(String serviceprincipalid) {
            this.serviceprincipalid = Objects.requireNonNull(serviceprincipalid);
            return this;
        }
        @CustomType.Setter
        public Builder serviceprincipalkey(String serviceprincipalkey) {
            this.serviceprincipalkey = Objects.requireNonNull(serviceprincipalkey);
            return this;
        }
        @CustomType.Setter
        public Builder serviceprincipalkeyHash(@Nullable String serviceprincipalkeyHash) {
            this.serviceprincipalkeyHash = serviceprincipalkeyHash;
            return this;
        }
        public ServiceEndpointAzureRMCredentials build() {
            final var o = new ServiceEndpointAzureRMCredentials();
            o.serviceprincipalid = serviceprincipalid;
            o.serviceprincipalkey = serviceprincipalkey;
            o.serviceprincipalkeyHash = serviceprincipalkeyHash;
            return o;
        }
    }
}