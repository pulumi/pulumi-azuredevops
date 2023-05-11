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


public final class PoolState extends com.pulumi.resources.ResourceArgs {

    public static final PoolState Empty = new PoolState();

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
     * The name of the agent pool.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the agent pool.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
     * 
     */
    @Import(name="poolType")
    private @Nullable Output<String> poolType;

    /**
     * @return Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
     * 
     */
    public Optional<Output<String>> poolType() {
        return Optional.ofNullable(this.poolType);
    }

    private PoolState() {}

    private PoolState(PoolState $) {
        this.autoProvision = $.autoProvision;
        this.autoUpdate = $.autoUpdate;
        this.name = $.name;
        this.poolType = $.poolType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PoolState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PoolState $;

        public Builder() {
            $ = new PoolState();
        }

        public Builder(PoolState defaults) {
            $ = new PoolState(Objects.requireNonNull(defaults));
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
         * @param name The name of the agent pool.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the agent pool.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param poolType Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
         * 
         * @return builder
         * 
         */
        public Builder poolType(@Nullable Output<String> poolType) {
            $.poolType = poolType;
            return this;
        }

        /**
         * @param poolType Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
         * 
         * @return builder
         * 
         */
        public Builder poolType(String poolType) {
            return poolType(Output.of(poolType));
        }

        public PoolState build() {
            return $;
        }
    }

}