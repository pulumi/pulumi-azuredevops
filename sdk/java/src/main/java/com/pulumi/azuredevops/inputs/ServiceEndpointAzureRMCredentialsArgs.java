// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointAzureRMCredentialsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointAzureRMCredentialsArgs Empty = new ServiceEndpointAzureRMCredentialsArgs();

    /**
     * The service principal application Id
     * 
     */
    @Import(name="serviceprincipalid", required=true)
    private Output<String> serviceprincipalid;

    /**
     * @return The service principal application Id
     * 
     */
    public Output<String> serviceprincipalid() {
        return this.serviceprincipalid;
    }

    /**
     * The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
     * 
     */
    @Import(name="serviceprincipalkey")
    private @Nullable Output<String> serviceprincipalkey;

    /**
     * @return The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
     * 
     */
    public Optional<Output<String>> serviceprincipalkey() {
        return Optional.ofNullable(this.serviceprincipalkey);
    }

    private ServiceEndpointAzureRMCredentialsArgs() {}

    private ServiceEndpointAzureRMCredentialsArgs(ServiceEndpointAzureRMCredentialsArgs $) {
        this.serviceprincipalid = $.serviceprincipalid;
        this.serviceprincipalkey = $.serviceprincipalkey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointAzureRMCredentialsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointAzureRMCredentialsArgs $;

        public Builder() {
            $ = new ServiceEndpointAzureRMCredentialsArgs();
        }

        public Builder(ServiceEndpointAzureRMCredentialsArgs defaults) {
            $ = new ServiceEndpointAzureRMCredentialsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceprincipalid The service principal application Id
         * 
         * @return builder
         * 
         */
        public Builder serviceprincipalid(Output<String> serviceprincipalid) {
            $.serviceprincipalid = serviceprincipalid;
            return this;
        }

        /**
         * @param serviceprincipalid The service principal application Id
         * 
         * @return builder
         * 
         */
        public Builder serviceprincipalid(String serviceprincipalid) {
            return serviceprincipalid(Output.of(serviceprincipalid));
        }

        /**
         * @param serviceprincipalkey The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
         * 
         * @return builder
         * 
         */
        public Builder serviceprincipalkey(@Nullable Output<String> serviceprincipalkey) {
            $.serviceprincipalkey = serviceprincipalkey;
            return this;
        }

        /**
         * @param serviceprincipalkey The service principal secret. This not required if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`.
         * 
         * @return builder
         * 
         */
        public Builder serviceprincipalkey(String serviceprincipalkey) {
            return serviceprincipalkey(Output.of(serviceprincipalkey));
        }

        public ServiceEndpointAzureRMCredentialsArgs build() {
            if ($.serviceprincipalid == null) {
                throw new MissingRequiredPropertyException("ServiceEndpointAzureRMCredentialsArgs", "serviceprincipalid");
            }
            return $;
        }
    }

}
