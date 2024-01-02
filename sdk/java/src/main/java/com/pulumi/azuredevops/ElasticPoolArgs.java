// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ElasticPoolArgs extends com.pulumi.resources.ResourceArgs {

    public static final ElasticPoolArgs Empty = new ElasticPoolArgs();

    /**
     * Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     * 
     */
    @Import(name="agentInteractiveUi")
    private @Nullable Output<Boolean> agentInteractiveUi;

    /**
     * @return Set whether agents should be configured to run with interactive UI. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> agentInteractiveUi() {
        return Optional.ofNullable(this.agentInteractiveUi);
    }

    /**
     * Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     * 
     */
    @Import(name="autoProvision")
    private @Nullable Output<Boolean> autoProvision;

    /**
     * @return Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> autoProvision() {
        return Optional.ofNullable(this.autoProvision);
    }

    /**
     * Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     * 
     */
    @Import(name="autoUpdate")
    private @Nullable Output<Boolean> autoUpdate;

    /**
     * @return Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> autoUpdate() {
        return Optional.ofNullable(this.autoUpdate);
    }

    /**
     * The ID of the Azure resource.
     * 
     */
    @Import(name="azureResourceId", required=true)
    private Output<String> azureResourceId;

    /**
     * @return The ID of the Azure resource.
     * 
     */
    public Output<String> azureResourceId() {
        return this.azureResourceId;
    }

    /**
     * Number of agents to keep on standby.
     * 
     */
    @Import(name="desiredIdle", required=true)
    private Output<Integer> desiredIdle;

    /**
     * @return Number of agents to keep on standby.
     * 
     */
    public Output<Integer> desiredIdle() {
        return this.desiredIdle;
    }

    /**
     * Maximum number of virtual machines in the scale set.
     * 
     */
    @Import(name="maxCapacity", required=true)
    private Output<Integer> maxCapacity;

    /**
     * @return Maximum number of virtual machines in the scale set.
     * 
     */
    public Output<Integer> maxCapacity() {
        return this.maxCapacity;
    }

    /**
     * The name of the Elastic pool.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Elastic pool.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Tear down virtual machines after every use. Defaults to `false`.
     * 
     */
    @Import(name="recycleAfterEachUse")
    private @Nullable Output<Boolean> recycleAfterEachUse;

    /**
     * @return Tear down virtual machines after every use. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> recycleAfterEachUse() {
        return Optional.ofNullable(this.recycleAfterEachUse);
    }

    /**
     * The ID of Service Endpoint used to connect to Azure.
     * 
     */
    @Import(name="serviceEndpointId", required=true)
    private Output<String> serviceEndpointId;

    /**
     * @return The ID of Service Endpoint used to connect to Azure.
     * 
     */
    public Output<String> serviceEndpointId() {
        return this.serviceEndpointId;
    }

    /**
     * The Project ID of Service Endpoint belongs to.
     * 
     */
    @Import(name="serviceEndpointScope", required=true)
    private Output<String> serviceEndpointScope;

    /**
     * @return The Project ID of Service Endpoint belongs to.
     * 
     */
    public Output<String> serviceEndpointScope() {
        return this.serviceEndpointScope;
    }

    /**
     * Delay in minutes before deleting excess idle agents. Defaults to `30`.
     * 
     */
    @Import(name="timeToLiveMinutes")
    private @Nullable Output<Integer> timeToLiveMinutes;

    /**
     * @return Delay in minutes before deleting excess idle agents. Defaults to `30`.
     * 
     */
    public Optional<Output<Integer>> timeToLiveMinutes() {
        return Optional.ofNullable(this.timeToLiveMinutes);
    }

    private ElasticPoolArgs() {}

    private ElasticPoolArgs(ElasticPoolArgs $) {
        this.agentInteractiveUi = $.agentInteractiveUi;
        this.autoProvision = $.autoProvision;
        this.autoUpdate = $.autoUpdate;
        this.azureResourceId = $.azureResourceId;
        this.desiredIdle = $.desiredIdle;
        this.maxCapacity = $.maxCapacity;
        this.name = $.name;
        this.recycleAfterEachUse = $.recycleAfterEachUse;
        this.serviceEndpointId = $.serviceEndpointId;
        this.serviceEndpointScope = $.serviceEndpointScope;
        this.timeToLiveMinutes = $.timeToLiveMinutes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ElasticPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ElasticPoolArgs $;

        public Builder() {
            $ = new ElasticPoolArgs();
        }

        public Builder(ElasticPoolArgs defaults) {
            $ = new ElasticPoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param agentInteractiveUi Set whether agents should be configured to run with interactive UI. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder agentInteractiveUi(@Nullable Output<Boolean> agentInteractiveUi) {
            $.agentInteractiveUi = agentInteractiveUi;
            return this;
        }

        /**
         * @param agentInteractiveUi Set whether agents should be configured to run with interactive UI. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder agentInteractiveUi(Boolean agentInteractiveUi) {
            return agentInteractiveUi(Output.of(agentInteractiveUi));
        }

        /**
         * @param autoProvision Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder autoProvision(@Nullable Output<Boolean> autoProvision) {
            $.autoProvision = autoProvision;
            return this;
        }

        /**
         * @param autoProvision Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder autoProvision(Boolean autoProvision) {
            return autoProvision(Output.of(autoProvision));
        }

        /**
         * @param autoUpdate Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoUpdate(@Nullable Output<Boolean> autoUpdate) {
            $.autoUpdate = autoUpdate;
            return this;
        }

        /**
         * @param autoUpdate Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoUpdate(Boolean autoUpdate) {
            return autoUpdate(Output.of(autoUpdate));
        }

        /**
         * @param azureResourceId The ID of the Azure resource.
         * 
         * @return builder
         * 
         */
        public Builder azureResourceId(Output<String> azureResourceId) {
            $.azureResourceId = azureResourceId;
            return this;
        }

        /**
         * @param azureResourceId The ID of the Azure resource.
         * 
         * @return builder
         * 
         */
        public Builder azureResourceId(String azureResourceId) {
            return azureResourceId(Output.of(azureResourceId));
        }

        /**
         * @param desiredIdle Number of agents to keep on standby.
         * 
         * @return builder
         * 
         */
        public Builder desiredIdle(Output<Integer> desiredIdle) {
            $.desiredIdle = desiredIdle;
            return this;
        }

        /**
         * @param desiredIdle Number of agents to keep on standby.
         * 
         * @return builder
         * 
         */
        public Builder desiredIdle(Integer desiredIdle) {
            return desiredIdle(Output.of(desiredIdle));
        }

        /**
         * @param maxCapacity Maximum number of virtual machines in the scale set.
         * 
         * @return builder
         * 
         */
        public Builder maxCapacity(Output<Integer> maxCapacity) {
            $.maxCapacity = maxCapacity;
            return this;
        }

        /**
         * @param maxCapacity Maximum number of virtual machines in the scale set.
         * 
         * @return builder
         * 
         */
        public Builder maxCapacity(Integer maxCapacity) {
            return maxCapacity(Output.of(maxCapacity));
        }

        /**
         * @param name The name of the Elastic pool.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Elastic pool.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param recycleAfterEachUse Tear down virtual machines after every use. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder recycleAfterEachUse(@Nullable Output<Boolean> recycleAfterEachUse) {
            $.recycleAfterEachUse = recycleAfterEachUse;
            return this;
        }

        /**
         * @param recycleAfterEachUse Tear down virtual machines after every use. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder recycleAfterEachUse(Boolean recycleAfterEachUse) {
            return recycleAfterEachUse(Output.of(recycleAfterEachUse));
        }

        /**
         * @param serviceEndpointId The ID of Service Endpoint used to connect to Azure.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointId(Output<String> serviceEndpointId) {
            $.serviceEndpointId = serviceEndpointId;
            return this;
        }

        /**
         * @param serviceEndpointId The ID of Service Endpoint used to connect to Azure.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointId(String serviceEndpointId) {
            return serviceEndpointId(Output.of(serviceEndpointId));
        }

        /**
         * @param serviceEndpointScope The Project ID of Service Endpoint belongs to.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointScope(Output<String> serviceEndpointScope) {
            $.serviceEndpointScope = serviceEndpointScope;
            return this;
        }

        /**
         * @param serviceEndpointScope The Project ID of Service Endpoint belongs to.
         * 
         * @return builder
         * 
         */
        public Builder serviceEndpointScope(String serviceEndpointScope) {
            return serviceEndpointScope(Output.of(serviceEndpointScope));
        }

        /**
         * @param timeToLiveMinutes Delay in minutes before deleting excess idle agents. Defaults to `30`.
         * 
         * @return builder
         * 
         */
        public Builder timeToLiveMinutes(@Nullable Output<Integer> timeToLiveMinutes) {
            $.timeToLiveMinutes = timeToLiveMinutes;
            return this;
        }

        /**
         * @param timeToLiveMinutes Delay in minutes before deleting excess idle agents. Defaults to `30`.
         * 
         * @return builder
         * 
         */
        public Builder timeToLiveMinutes(Integer timeToLiveMinutes) {
            return timeToLiveMinutes(Output.of(timeToLiveMinutes));
        }

        public ElasticPoolArgs build() {
            if ($.azureResourceId == null) {
                throw new MissingRequiredPropertyException("ElasticPoolArgs", "azureResourceId");
            }
            if ($.desiredIdle == null) {
                throw new MissingRequiredPropertyException("ElasticPoolArgs", "desiredIdle");
            }
            if ($.maxCapacity == null) {
                throw new MissingRequiredPropertyException("ElasticPoolArgs", "maxCapacity");
            }
            if ($.serviceEndpointId == null) {
                throw new MissingRequiredPropertyException("ElasticPoolArgs", "serviceEndpointId");
            }
            if ($.serviceEndpointScope == null) {
                throw new MissingRequiredPropertyException("ElasticPoolArgs", "serviceEndpointScope");
            }
            return $;
        }
    }

}
