// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupEntitlementState extends com.pulumi.resources.ResourceArgs {

    public static final GroupEntitlementState Empty = new GroupEntitlementState();

    /**
     * Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    @Import(name="accountLicenseType")
    private @Nullable Output<String> accountLicenseType;

    /**
     * @return Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
     * 
     */
    public Optional<Output<String>> accountLicenseType() {
        return Optional.ofNullable(this.accountLicenseType);
    }

    /**
     * The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
     * 
     */
    @Import(name="descriptor")
    private @Nullable Output<String> descriptor;

    /**
     * @return The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
     * 
     */
    public Optional<Output<String>> descriptor() {
        return Optional.ofNullable(this.descriptor);
    }

    /**
     * The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The source of the licensing (e.g. Account. MSDN etc.). Possible values are: `account`, `auto`, `msdn`, `none`, `profile`, `trial`. Defaults to `account`.
     * 
     * &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
     * 
     */
    @Import(name="licensingSource")
    private @Nullable Output<String> licensingSource;

    /**
     * @return The source of the licensing (e.g. Account. MSDN etc.). Possible values are: `account`, `auto`, `msdn`, `none`, `profile`, `trial`. Defaults to `account`.
     * 
     * &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
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
     * The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     * 
     */
    @Import(name="originId")
    private @Nullable Output<String> originId;

    /**
     * @return The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
     * 
     */
    public Optional<Output<String>> originId() {
        return Optional.ofNullable(this.originId);
    }

    /**
     * The principal name of a graph member on Azure DevOps
     * 
     */
    @Import(name="principalName")
    private @Nullable Output<String> principalName;

    /**
     * @return The principal name of a graph member on Azure DevOps
     * 
     */
    public Optional<Output<String>> principalName() {
        return Optional.ofNullable(this.principalName);
    }

    private GroupEntitlementState() {}

    private GroupEntitlementState(GroupEntitlementState $) {
        this.accountLicenseType = $.accountLicenseType;
        this.descriptor = $.descriptor;
        this.displayName = $.displayName;
        this.licensingSource = $.licensingSource;
        this.origin = $.origin;
        this.originId = $.originId;
        this.principalName = $.principalName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupEntitlementState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupEntitlementState $;

        public Builder() {
            $ = new GroupEntitlementState();
        }

        public Builder(GroupEntitlementState defaults) {
            $ = new GroupEntitlementState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountLicenseType Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
         * 
         * @return builder
         * 
         */
        public Builder accountLicenseType(@Nullable Output<String> accountLicenseType) {
            $.accountLicenseType = accountLicenseType;
            return this;
        }

        /**
         * @param accountLicenseType Type of Account License. Possible values are: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition, the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
         * 
         * @return builder
         * 
         */
        public Builder accountLicenseType(String accountLicenseType) {
            return accountLicenseType(Output.of(accountLicenseType));
        }

        /**
         * @param descriptor The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(@Nullable Output<String> descriptor) {
            $.descriptor = descriptor;
            return this;
        }

        /**
         * @param descriptor The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the group graph subject.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(String descriptor) {
            return descriptor(Output.of(descriptor));
        }

        /**
         * @param displayName The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name is the name used in Azure DevOps UI. Cannot be set together with `origin_id` and `origin`.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param licensingSource The source of the licensing (e.g. Account. MSDN etc.). Possible values are: `account`, `auto`, `msdn`, `none`, `profile`, `trial`. Defaults to `account`.
         * 
         * &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
         * 
         * @return builder
         * 
         */
        public Builder licensingSource(@Nullable Output<String> licensingSource) {
            $.licensingSource = licensingSource;
            return this;
        }

        /**
         * @param licensingSource The source of the licensing (e.g. Account. MSDN etc.). Possible values are: `account`, `auto`, `msdn`, `none`, `profile`, `trial`. Defaults to `account`.
         * 
         * &gt; **NOTE:** A existing group in Azure AD can only be referenced by the combination of `origin_id` and `origin`.
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
         * @param originId The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder originId(@Nullable Output<String> originId) {
            $.originId = originId;
            return this;
        }

        /**
         * @param originId The unique identifier from the system of origin. Typically, a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder originId(String originId) {
            return originId(Output.of(originId));
        }

        /**
         * @param principalName The principal name of a graph member on Azure DevOps
         * 
         * @return builder
         * 
         */
        public Builder principalName(@Nullable Output<String> principalName) {
            $.principalName = principalName;
            return this;
        }

        /**
         * @param principalName The principal name of a graph member on Azure DevOps
         * 
         * @return builder
         * 
         */
        public Builder principalName(String principalName) {
            return principalName(Output.of(principalName));
        }

        public GroupEntitlementState build() {
            return $;
        }
    }

}
