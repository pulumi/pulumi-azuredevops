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


public final class ServicePrincipalEntitlementArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicePrincipalEntitlementArgs Empty = new ServicePrincipalEntitlementArgs();

    /**
     * Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    @Import(name="accountLicenseType")
    private @Nullable Output<String> accountLicenseType;

    /**
     * @return Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    public Optional<Output<String>> accountLicenseType() {
        return Optional.ofNullable(this.accountLicenseType);
    }

    /**
     * The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     * 
     */
    @Import(name="licensingSource")
    private @Nullable Output<String> licensingSource;

    /**
     * @return The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
     * 
     */
    public Optional<Output<String>> licensingSource() {
        return Optional.ofNullable(this.licensingSource);
    }

    /**
     * The type of source provider for the origin identifier.
     * 
     */
    @Import(name="origin")
    private @Nullable Output<String> origin;

    /**
     * @return The type of source provider for the origin identifier.
     * 
     */
    public Optional<Output<String>> origin() {
        return Optional.ofNullable(this.origin);
    }

    /**
     * The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
     * 
     */
    @Import(name="originId", required=true)
    private Output<String> originId;

    /**
     * @return The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
     * 
     */
    public Output<String> originId() {
        return this.originId;
    }

    private ServicePrincipalEntitlementArgs() {}

    private ServicePrincipalEntitlementArgs(ServicePrincipalEntitlementArgs $) {
        this.accountLicenseType = $.accountLicenseType;
        this.licensingSource = $.licensingSource;
        this.origin = $.origin;
        this.originId = $.originId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePrincipalEntitlementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePrincipalEntitlementArgs $;

        public Builder() {
            $ = new ServicePrincipalEntitlementArgs();
        }

        public Builder(ServicePrincipalEntitlementArgs defaults) {
            $ = new ServicePrincipalEntitlementArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountLicenseType Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
         * 
         * @return builder
         * 
         */
        public Builder accountLicenseType(@Nullable Output<String> accountLicenseType) {
            $.accountLicenseType = accountLicenseType;
            return this;
        }

        /**
         * @param accountLicenseType Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
         * 
         * @return builder
         * 
         */
        public Builder accountLicenseType(String accountLicenseType) {
            return accountLicenseType(Output.of(accountLicenseType));
        }

        /**
         * @param licensingSource The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
         * 
         * @return builder
         * 
         */
        public Builder licensingSource(@Nullable Output<String> licensingSource) {
            $.licensingSource = licensingSource;
            return this;
        }

        /**
         * @param licensingSource The source of the licensing (e.g. Account. MSDN etc.) Possible values are: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`
         * 
         * @return builder
         * 
         */
        public Builder licensingSource(String licensingSource) {
            return licensingSource(Output.of(licensingSource));
        }

        /**
         * @param origin The type of source provider for the origin identifier.
         * 
         * @return builder
         * 
         */
        public Builder origin(@Nullable Output<String> origin) {
            $.origin = origin;
            return this;
        }

        /**
         * @param origin The type of source provider for the origin identifier.
         * 
         * @return builder
         * 
         */
        public Builder origin(String origin) {
            return origin(Output.of(origin));
        }

        /**
         * @param originId The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
         * 
         * @return builder
         * 
         */
        public Builder originId(Output<String> originId) {
            $.originId = originId;
            return this;
        }

        /**
         * @param originId The Object ID of the service principal in Entra ID. Changing this forces a new Service Principal Entitlement to be created.
         * 
         * @return builder
         * 
         */
        public Builder originId(String originId) {
            return originId(Output.of(originId));
        }

        public ServicePrincipalEntitlementArgs build() {
            if ($.originId == null) {
                throw new MissingRequiredPropertyException("ServicePrincipalEntitlementArgs", "originId");
            }
            return $;
        }
    }

}
