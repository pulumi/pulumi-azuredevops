// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIdentityUsersResult {
    private String descriptor;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return This is the PrincipalName of this identity member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the identity member.
     * 
     */
    private String name;
    private @Nullable String searchFilter;

    private GetIdentityUsersResult() {}
    public String descriptor() {
        return this.descriptor;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return This is the PrincipalName of this identity member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the identity member.
     * 
     */
    public String name() {
        return this.name;
    }
    public Optional<String> searchFilter() {
        return Optional.ofNullable(this.searchFilter);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityUsersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String descriptor;
        private String id;
        private String name;
        private @Nullable String searchFilter;
        public Builder() {}
        public Builder(GetIdentityUsersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptor = defaults.descriptor;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.searchFilter = defaults.searchFilter;
        }

        @CustomType.Setter
        public Builder descriptor(String descriptor) {
            if (descriptor == null) {
              throw new MissingRequiredPropertyException("GetIdentityUsersResult", "descriptor");
            }
            this.descriptor = descriptor;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentityUsersResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIdentityUsersResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder searchFilter(@Nullable String searchFilter) {

            this.searchFilter = searchFilter;
            return this;
        }
        public GetIdentityUsersResult build() {
            final var _resultValue = new GetIdentityUsersResult();
            _resultValue.descriptor = descriptor;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.searchFilter = searchFilter;
            return _resultValue;
        }
    }
}
