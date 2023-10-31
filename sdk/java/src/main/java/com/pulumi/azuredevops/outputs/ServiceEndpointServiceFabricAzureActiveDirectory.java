// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceEndpointServiceFabricAzureActiveDirectory {
    /**
     * @return Password for the Azure Active Directory account.
     * 
     */
    private String password;
    /**
     * @return The common name(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (&#39;,&#39;)
     * 
     */
    private @Nullable String serverCertificateCommonName;
    /**
     * @return Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
     * 
     */
    private String serverCertificateLookup;
    /**
     * @return The thumbprint(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (&#39;,&#39;)
     * 
     */
    private @Nullable String serverCertificateThumbprint;
    /**
     * @return Specify an Azure Active Directory account.
     * 
     */
    private String username;

    private ServiceEndpointServiceFabricAzureActiveDirectory() {}
    /**
     * @return Password for the Azure Active Directory account.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return The common name(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (&#39;,&#39;)
     * 
     */
    public Optional<String> serverCertificateCommonName() {
        return Optional.ofNullable(this.serverCertificateCommonName);
    }
    /**
     * @return Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
     * 
     */
    public String serverCertificateLookup() {
        return this.serverCertificateLookup;
    }
    /**
     * @return The thumbprint(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (&#39;,&#39;)
     * 
     */
    public Optional<String> serverCertificateThumbprint() {
        return Optional.ofNullable(this.serverCertificateThumbprint);
    }
    /**
     * @return Specify an Azure Active Directory account.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointServiceFabricAzureActiveDirectory defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private @Nullable String serverCertificateCommonName;
        private String serverCertificateLookup;
        private @Nullable String serverCertificateThumbprint;
        private String username;
        public Builder() {}
        public Builder(ServiceEndpointServiceFabricAzureActiveDirectory defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.serverCertificateCommonName = defaults.serverCertificateCommonName;
    	      this.serverCertificateLookup = defaults.serverCertificateLookup;
    	      this.serverCertificateThumbprint = defaults.serverCertificateThumbprint;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder serverCertificateCommonName(@Nullable String serverCertificateCommonName) {
            this.serverCertificateCommonName = serverCertificateCommonName;
            return this;
        }
        @CustomType.Setter
        public Builder serverCertificateLookup(String serverCertificateLookup) {
            this.serverCertificateLookup = Objects.requireNonNull(serverCertificateLookup);
            return this;
        }
        @CustomType.Setter
        public Builder serverCertificateThumbprint(@Nullable String serverCertificateThumbprint) {
            this.serverCertificateThumbprint = serverCertificateThumbprint;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public ServiceEndpointServiceFabricAzureActiveDirectory build() {
            final var _resultValue = new ServiceEndpointServiceFabricAzureActiveDirectory();
            _resultValue.password = password;
            _resultValue.serverCertificateCommonName = serverCertificateCommonName;
            _resultValue.serverCertificateLookup = serverCertificateLookup;
            _resultValue.serverCertificateThumbprint = serverCertificateThumbprint;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
