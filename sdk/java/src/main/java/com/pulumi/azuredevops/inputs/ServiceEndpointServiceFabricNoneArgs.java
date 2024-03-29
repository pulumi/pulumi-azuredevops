// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceEndpointServiceFabricNoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceEndpointServiceFabricNoneArgs Empty = new ServiceEndpointServiceFabricNoneArgs();

    /**
     * Fully qualified domain SPN for gMSA account. This is applicable only if `unsecured` option is disabled.
     * 
     */
    @Import(name="clusterSpn")
    private @Nullable Output<String> clusterSpn;

    /**
     * @return Fully qualified domain SPN for gMSA account. This is applicable only if `unsecured` option is disabled.
     * 
     */
    public Optional<Output<String>> clusterSpn() {
        return Optional.ofNullable(this.clusterSpn);
    }

    /**
     * Skip using windows security for authentication.
     * 
     */
    @Import(name="unsecured")
    private @Nullable Output<Boolean> unsecured;

    /**
     * @return Skip using windows security for authentication.
     * 
     */
    public Optional<Output<Boolean>> unsecured() {
        return Optional.ofNullable(this.unsecured);
    }

    private ServiceEndpointServiceFabricNoneArgs() {}

    private ServiceEndpointServiceFabricNoneArgs(ServiceEndpointServiceFabricNoneArgs $) {
        this.clusterSpn = $.clusterSpn;
        this.unsecured = $.unsecured;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceEndpointServiceFabricNoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceEndpointServiceFabricNoneArgs $;

        public Builder() {
            $ = new ServiceEndpointServiceFabricNoneArgs();
        }

        public Builder(ServiceEndpointServiceFabricNoneArgs defaults) {
            $ = new ServiceEndpointServiceFabricNoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterSpn Fully qualified domain SPN for gMSA account. This is applicable only if `unsecured` option is disabled.
         * 
         * @return builder
         * 
         */
        public Builder clusterSpn(@Nullable Output<String> clusterSpn) {
            $.clusterSpn = clusterSpn;
            return this;
        }

        /**
         * @param clusterSpn Fully qualified domain SPN for gMSA account. This is applicable only if `unsecured` option is disabled.
         * 
         * @return builder
         * 
         */
        public Builder clusterSpn(String clusterSpn) {
            return clusterSpn(Output.of(clusterSpn));
        }

        /**
         * @param unsecured Skip using windows security for authentication.
         * 
         * @return builder
         * 
         */
        public Builder unsecured(@Nullable Output<Boolean> unsecured) {
            $.unsecured = unsecured;
            return this;
        }

        /**
         * @param unsecured Skip using windows security for authentication.
         * 
         * @return builder
         * 
         */
        public Builder unsecured(Boolean unsecured) {
            return unsecured(Output.of(unsecured));
        }

        public ServiceEndpointServiceFabricNoneArgs build() {
            return $;
        }
    }

}
