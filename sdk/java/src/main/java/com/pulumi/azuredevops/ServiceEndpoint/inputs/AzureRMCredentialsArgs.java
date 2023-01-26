// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.ServiceEndpoint.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AzureRMCredentialsArgs extends com.pulumi.resources.ResourceArgs {

    public static final AzureRMCredentialsArgs Empty = new AzureRMCredentialsArgs();

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
     * The service principal secret.
     * 
     */
    @Import(name="serviceprincipalkey", required=true)
    private Output<String> serviceprincipalkey;

    /**
     * @return The service principal secret.
     * 
     */
    public Output<String> serviceprincipalkey() {
        return this.serviceprincipalkey;
    }

    @Import(name="serviceprincipalkeyHash")
    private @Nullable Output<String> serviceprincipalkeyHash;

    public Optional<Output<String>> serviceprincipalkeyHash() {
        return Optional.ofNullable(this.serviceprincipalkeyHash);
    }

    private AzureRMCredentialsArgs() {}

    private AzureRMCredentialsArgs(AzureRMCredentialsArgs $) {
        this.serviceprincipalid = $.serviceprincipalid;
        this.serviceprincipalkey = $.serviceprincipalkey;
        this.serviceprincipalkeyHash = $.serviceprincipalkeyHash;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AzureRMCredentialsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AzureRMCredentialsArgs $;

        public Builder() {
            $ = new AzureRMCredentialsArgs();
        }

        public Builder(AzureRMCredentialsArgs defaults) {
            $ = new AzureRMCredentialsArgs(Objects.requireNonNull(defaults));
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
         * @param serviceprincipalkey The service principal secret.
         * 
         * @return builder
         * 
         */
        public Builder serviceprincipalkey(Output<String> serviceprincipalkey) {
            $.serviceprincipalkey = serviceprincipalkey;
            return this;
        }

        /**
         * @param serviceprincipalkey The service principal secret.
         * 
         * @return builder
         * 
         */
        public Builder serviceprincipalkey(String serviceprincipalkey) {
            return serviceprincipalkey(Output.of(serviceprincipalkey));
        }

        public Builder serviceprincipalkeyHash(@Nullable Output<String> serviceprincipalkeyHash) {
            $.serviceprincipalkeyHash = serviceprincipalkeyHash;
            return this;
        }

        public Builder serviceprincipalkeyHash(String serviceprincipalkeyHash) {
            return serviceprincipalkeyHash(Output.of(serviceprincipalkeyHash));
        }

        public AzureRMCredentialsArgs build() {
            $.serviceprincipalid = Objects.requireNonNull($.serviceprincipalid, "expected parameter 'serviceprincipalid' to be non-null");
            $.serviceprincipalkey = Objects.requireNonNull($.serviceprincipalkey, "expected parameter 'serviceprincipalkey' to be non-null");
            return $;
        }
    }

}
