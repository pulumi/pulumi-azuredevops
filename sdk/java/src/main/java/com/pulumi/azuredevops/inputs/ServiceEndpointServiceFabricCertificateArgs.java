// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointServiceFabricCertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointServiceFabricCertificateArgs Empty = new ServiceEndpointServiceFabricCertificateArgs();

    /**
     * Base64 encoding of the cluster&#39;s client certificate file.
     * 
     */
    @Import(name="clientCertificate", required=true)
    private Output<String> clientCertificate;

    /**
     * @return Base64 encoding of the cluster&#39;s client certificate file.
     * 
     */
    public Output<String> clientCertificate() {
        return this.clientCertificate;
    }

    /**
     * Password for the certificate.
     * 
     */
    @Import(name="clientCertificatePassword")
    private @Nullable Output<String> clientCertificatePassword;

    /**
     * @return Password for the certificate.
     * 
     */
    public Optional<Output<String>> clientCertificatePassword() {
        return Optional.ofNullable(this.clientCertificatePassword);
    }

    /**
     * The common name(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (&#39;,&#39;)
     * 
     */
    @Import(name="serverCertificateCommonName")
    private @Nullable Output<String> serverCertificateCommonName;

    /**
     * @return The common name(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (&#39;,&#39;)
     * 
     */
    public Optional<Output<String>> serverCertificateCommonName() {
        return Optional.ofNullable(this.serverCertificateCommonName);
    }

    /**
     * Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
     * 
     */
    @Import(name="serverCertificateLookup", required=true)
    private Output<String> serverCertificateLookup;

    /**
     * @return Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
     * 
     */
    public Output<String> serverCertificateLookup() {
        return this.serverCertificateLookup;
    }

    /**
     * The thumbprint(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (&#39;,&#39;)
     * 
     */
    @Import(name="serverCertificateThumbprint")
    private @Nullable Output<String> serverCertificateThumbprint;

    /**
     * @return The thumbprint(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (&#39;,&#39;)
     * 
     */
    public Optional<Output<String>> serverCertificateThumbprint() {
        return Optional.ofNullable(this.serverCertificateThumbprint);
    }

    private ServiceEndpointServiceFabricCertificateArgs() {}

    private ServiceEndpointServiceFabricCertificateArgs(ServiceEndpointServiceFabricCertificateArgs $) {
        this.clientCertificate = $.clientCertificate;
        this.clientCertificatePassword = $.clientCertificatePassword;
        this.serverCertificateCommonName = $.serverCertificateCommonName;
        this.serverCertificateLookup = $.serverCertificateLookup;
        this.serverCertificateThumbprint = $.serverCertificateThumbprint;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointServiceFabricCertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointServiceFabricCertificateArgs $;

        public Builder() {
            $ = new ServiceEndpointServiceFabricCertificateArgs();
        }

        public Builder(ServiceEndpointServiceFabricCertificateArgs defaults) {
            $ = new ServiceEndpointServiceFabricCertificateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientCertificate Base64 encoding of the cluster&#39;s client certificate file.
         * 
         * @return builder
         * 
         */
        public Builder clientCertificate(Output<String> clientCertificate) {
            $.clientCertificate = clientCertificate;
            return this;
        }

        /**
         * @param clientCertificate Base64 encoding of the cluster&#39;s client certificate file.
         * 
         * @return builder
         * 
         */
        public Builder clientCertificate(String clientCertificate) {
            return clientCertificate(Output.of(clientCertificate));
        }

        /**
         * @param clientCertificatePassword Password for the certificate.
         * 
         * @return builder
         * 
         */
        public Builder clientCertificatePassword(@Nullable Output<String> clientCertificatePassword) {
            $.clientCertificatePassword = clientCertificatePassword;
            return this;
        }

        /**
         * @param clientCertificatePassword Password for the certificate.
         * 
         * @return builder
         * 
         */
        public Builder clientCertificatePassword(String clientCertificatePassword) {
            return clientCertificatePassword(Output.of(clientCertificatePassword));
        }

        /**
         * @param serverCertificateCommonName The common name(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (&#39;,&#39;)
         * 
         * @return builder
         * 
         */
        public Builder serverCertificateCommonName(@Nullable Output<String> serverCertificateCommonName) {
            $.serverCertificateCommonName = serverCertificateCommonName;
            return this;
        }

        /**
         * @param serverCertificateCommonName The common name(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (&#39;,&#39;)
         * 
         * @return builder
         * 
         */
        public Builder serverCertificateCommonName(String serverCertificateCommonName) {
            return serverCertificateCommonName(Output.of(serverCertificateCommonName));
        }

        /**
         * @param serverCertificateLookup Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
         * 
         * @return builder
         * 
         */
        public Builder serverCertificateLookup(Output<String> serverCertificateLookup) {
            $.serverCertificateLookup = serverCertificateLookup;
            return this;
        }

        /**
         * @param serverCertificateLookup Verification mode for the cluster. Possible values include `Thumbprint` or `CommonName`.
         * 
         * @return builder
         * 
         */
        public Builder serverCertificateLookup(String serverCertificateLookup) {
            return serverCertificateLookup(Output.of(serverCertificateLookup));
        }

        /**
         * @param serverCertificateThumbprint The thumbprint(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (&#39;,&#39;)
         * 
         * @return builder
         * 
         */
        public Builder serverCertificateThumbprint(@Nullable Output<String> serverCertificateThumbprint) {
            $.serverCertificateThumbprint = serverCertificateThumbprint;
            return this;
        }

        /**
         * @param serverCertificateThumbprint The thumbprint(s) of the cluster&#39;s certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (&#39;,&#39;)
         * 
         * @return builder
         * 
         */
        public Builder serverCertificateThumbprint(String serverCertificateThumbprint) {
            return serverCertificateThumbprint(Output.of(serverCertificateThumbprint));
        }

        public ServiceEndpointServiceFabricCertificateArgs build() {
            $.clientCertificate = Objects.requireNonNull($.clientCertificate, "expected parameter 'clientCertificate' to be non-null");
            $.serverCertificateLookup = Objects.requireNonNull($.serverCertificateLookup, "expected parameter 'serverCertificateLookup' to be non-null");
            return $;
        }
    }

}
