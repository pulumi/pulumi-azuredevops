// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Identities.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetUsersUser {
    /**
     * @return The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
     * 
     */
    private String descriptor;
    /**
     * @return This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
     * 
     */
    private String displayName;
    /**
     * @return The user ID.
     * 
     */
    private String id;
    /**
     * @return The email address of record for a given graph member. This may be different than the principal name.
     * 
     */
    private String mailAddress;
    /**
     * @return The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
     * 
     */
    private String origin;
    /**
     * @return The unique identifier from the system of origin.
     * 
     */
    private @Nullable String originId;
    /**
     * @return The PrincipalName of this graph member from the source provider.
     * 
     */
    private String principalName;

    private GetUsersUser() {}
    /**
     * @return The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
     * 
     */
    public String descriptor() {
        return this.descriptor;
    }
    /**
     * @return This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return The user ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The email address of record for a given graph member. This may be different than the principal name.
     * 
     */
    public String mailAddress() {
        return this.mailAddress;
    }
    /**
     * @return The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
     * 
     */
    public String origin() {
        return this.origin;
    }
    /**
     * @return The unique identifier from the system of origin.
     * 
     */
    public Optional<String> originId() {
        return Optional.ofNullable(this.originId);
    }
    /**
     * @return The PrincipalName of this graph member from the source provider.
     * 
     */
    public String principalName() {
        return this.principalName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUsersUser defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String descriptor;
        private String displayName;
        private String id;
        private String mailAddress;
        private String origin;
        private @Nullable String originId;
        private String principalName;
        public Builder() {}
        public Builder(GetUsersUser defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptor = defaults.descriptor;
    	      this.displayName = defaults.displayName;
    	      this.id = defaults.id;
    	      this.mailAddress = defaults.mailAddress;
    	      this.origin = defaults.origin;
    	      this.originId = defaults.originId;
    	      this.principalName = defaults.principalName;
        }

        @CustomType.Setter
        public Builder descriptor(String descriptor) {
            if (descriptor == null) {
              throw new MissingRequiredPropertyException("GetUsersUser", "descriptor");
            }
            this.descriptor = descriptor;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetUsersUser", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUsersUser", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mailAddress(String mailAddress) {
            if (mailAddress == null) {
              throw new MissingRequiredPropertyException("GetUsersUser", "mailAddress");
            }
            this.mailAddress = mailAddress;
            return this;
        }
        @CustomType.Setter
        public Builder origin(String origin) {
            if (origin == null) {
              throw new MissingRequiredPropertyException("GetUsersUser", "origin");
            }
            this.origin = origin;
            return this;
        }
        @CustomType.Setter
        public Builder originId(@Nullable String originId) {

            this.originId = originId;
            return this;
        }
        @CustomType.Setter
        public Builder principalName(String principalName) {
            if (principalName == null) {
              throw new MissingRequiredPropertyException("GetUsersUser", "principalName");
            }
            this.principalName = principalName;
            return this;
        }
        public GetUsersUser build() {
            final var _resultValue = new GetUsersUser();
            _resultValue.descriptor = descriptor;
            _resultValue.displayName = displayName;
            _resultValue.id = id;
            _resultValue.mailAddress = mailAddress;
            _resultValue.origin = origin;
            _resultValue.originId = originId;
            _resultValue.principalName = principalName;
            return _resultValue;
        }
    }
}
