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


public final class GetServiceendpointSonarcloudArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceendpointSonarcloudArgs Empty = new GetServiceendpointSonarcloudArgs();

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * the ID of the Service Endpoint.
     * 
     */
    @Import(name="serviceEndpointId")
    private @Nullable Output<String> serviceEndpointId;

    /**
     * @return the ID of the Service Endpoint.
     * 
     */
    public Optional<Output<String>> serviceEndpointId() {
        return Optional.ofNullable(this.serviceEndpointId);
    }

    /**
     * the Name of the Service Endpoint.
     * 
     * &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
     * 
     */
    @Import(name="serviceEndpointName")
    private @Nullable Output<String> serviceEndpointName;

    /**
     * @return the Name of the Service Endpoint.
     * 
     * &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
     * 
     */
    public Optional<Output<String>> serviceEndpointName() {
        return Optional.ofNullable(this.serviceEndpointName);
    }

    private GetServiceendpointSonarcloudArgs() {}

    private GetServiceendpointSonarcloudArgs(GetServiceendpointSonarcloudArgs $) {
        this.projectId = $.projectId;
        this.serviceEndpointId = $.serviceEndpointId;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceendpointSonarcloudArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceendpointSonarcloudArgs $;

        public Builder() {
            $ = new GetServiceendpointSonarcloudArgs();
        }

        public Builder(GetServiceendpointSonarcloudArgs defaults) {
            $ = new GetServiceendpointSonarcloudArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param serviceEndpointId the ID of the Service Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointId(@Nullable Output<String> serviceEndpointId) {
            $.serviceEndpointId = serviceEndpointId;
            return this;
        }

        /**
         * @param serviceEndpointId the ID of the Service Endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointId(String serviceEndpointId) {
            return serviceEndpointId(Output.of(serviceEndpointId));
        }

        /**
         * @param serviceEndpointName the Name of the Service Endpoint.
         * 
         * &gt; **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(@Nullable Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
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
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        public GetServiceendpointSonarcloudArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetServiceendpointSonarcloudArgs", "projectId");
            }
            return $;
        }
    }

}
