// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceendpointAzureServiceBusArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceendpointAzureServiceBusArgs Empty = new ServiceendpointAzureServiceBusArgs();

    /**
     * The  Azure Service Bus Connection string.
     * 
     */
    @Import(name="connectionString", required=true)
    private Output<String> connectionString;

    /**
     * @return The  Azure Service Bus Connection string.
     * 
     */
    public Output<String> connectionString() {
        return this.connectionString;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

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
     * The Azure Service Bus Queue Name.
     * 
     */
    @Import(name="queueName", required=true)
    private Output<String> queueName;

    /**
     * @return The Azure Service Bus Queue Name.
     * 
     */
    public Output<String> queueName() {
        return this.queueName;
    }

    /**
     * The Service Endpoint name.
     * 
     */
    @Import(name="serviceEndpointName", required=true)
    private Output<String> serviceEndpointName;

    /**
     * @return The Service Endpoint name.
     * 
     */
    public Output<String> serviceEndpointName() {
        return this.serviceEndpointName;
    }

    private ServiceendpointAzureServiceBusArgs() {}

    private ServiceendpointAzureServiceBusArgs(ServiceendpointAzureServiceBusArgs $) {
        this.connectionString = $.connectionString;
        this.description = $.description;
        this.projectId = $.projectId;
        this.queueName = $.queueName;
        this.serviceEndpointName = $.serviceEndpointName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceendpointAzureServiceBusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceendpointAzureServiceBusArgs $;

        public Builder() {
            $ = new ServiceendpointAzureServiceBusArgs();
        }

        public Builder(ServiceendpointAzureServiceBusArgs defaults) {
            $ = new ServiceendpointAzureServiceBusArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionString The  Azure Service Bus Connection string.
         * 
         * @return builder
         * 
         */
        public Builder connectionString(Output<String> connectionString) {
            $.connectionString = connectionString;
            return this;
        }

        /**
         * @param connectionString The  Azure Service Bus Connection string.
         * 
         * @return builder
         * 
         */
        public Builder connectionString(String connectionString) {
            return connectionString(Output.of(connectionString));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param queueName The Azure Service Bus Queue Name.
         * 
         * @return builder
         * 
         */
        public Builder queueName(Output<String> queueName) {
            $.queueName = queueName;
            return this;
        }

        /**
         * @param queueName The Azure Service Bus Queue Name.
         * 
         * @return builder
         * 
         */
        public Builder queueName(String queueName) {
            return queueName(Output.of(queueName));
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(Output<String> serviceEndpointName) {
            $.serviceEndpointName = serviceEndpointName;
            return this;
        }

        /**
         * @param serviceEndpointName The Service Endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointName(String serviceEndpointName) {
            return serviceEndpointName(Output.of(serviceEndpointName));
        }

        public ServiceendpointAzureServiceBusArgs build() {
            if ($.connectionString == null) {
                throw new MissingRequiredPropertyException("ServiceendpointAzureServiceBusArgs", "connectionString");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("ServiceendpointAzureServiceBusArgs", "projectId");
            }
            if ($.queueName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointAzureServiceBusArgs", "queueName");
            }
            if ($.serviceEndpointName == null) {
                throw new MissingRequiredPropertyException("ServiceendpointAzureServiceBusArgs", "serviceEndpointName");
            }
            return $;
        }
    }

}
