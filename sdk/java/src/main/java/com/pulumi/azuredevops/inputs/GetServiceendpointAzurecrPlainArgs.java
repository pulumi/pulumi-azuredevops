// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceendpointAzurecrPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceendpointAzurecrPlainArgs Empty = new GetServiceendpointAzurecrPlainArgs();

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    /**
     * the ID of the Service Endpoint.
     * 
     */
    @Import(name="serviceEndpointId")
    private @Nullable String serviceEndpointId;

    /**
     * @return the ID of the Service Endpoint.
     * 
     */
    public Optional<String> serviceEndpointId() {
        return Optional.ofNullable(this.serviceEndpointId);
    }

    /**
     * the Name of the Service Endpoint.
     * 
     * &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
     * 
     */
    @Import(name="serviceEndpointName")
    private @Nullable String serviceEndpointName;

    /**
     * @return the Name of the Service Endpoint.
     * 
     * &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
     * 
     */
    public Optional<String> serviceEndpointName() {
        return Optional.ofNullable(this.serviceEndpointName);
    }

    private GetServiceendpointAzurecrPlainArgs() {}

    private GetServiceendpointAzurecrPlainArgs(GetServiceendpointAzurecrPlainArgs $) {
        this.projectId = $.projectId;
        this.serviceEndpointId = $.serviceEndpointId;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceendpointAzurecrPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceendpointAzurecrPlainArgs $;

        public Builder() {
            $ = new GetServiceendpointAzurecrPlainArgs();
        }

        public Builder(GetServiceendpointAzurecrPlainArgs defaults) {
            $ = new GetServiceendpointAzurecrPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param serviceEndpointId the ID of the Service Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointId(@Nullable String serviceEndpointId) {
            $.serviceEndpointId = serviceEndpointId;
            return this;
        }

        /**
         * @param serviceEndpointName the Name of the Service Endpoint.
         * 
         * &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(@Nullable String serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        public GetServiceendpointAzurecrPlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetServiceendpointAzurecrPlainArgs", "projectId");
            }
            return $;
        }
    }

}
