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
    /**
     * @return The Descriptor of the user.
     * 
     */
    private String descriptor;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private @Nullable String searchFilter;
    /**
     * @return The Subject Descriptor of the user.
     * 
     */
    private String subjectDescriptor;

    private GetIdentityUsersResult() {}
    /**
     * @return The Descriptor of the user.
     * 
     */
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
    public String name() {
        return this.name;
    }
    public Optional<String> searchFilter() {
        return Optional.ofNullable(this.searchFilter);
    }
    /**
     * @return The Subject Descriptor of the user.
     * 
     */
    public String subjectDescriptor() {
        return this.subjectDescriptor;
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
        private String subjectDescriptor;
        public Builder() {}
        public Builder(GetIdentityUsersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptor = defaults.descriptor;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.searchFilter = defaults.searchFilter;
    	      this.subjectDescriptor = defaults.subjectDescriptor;
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
        @CustomType.Setter
        public Builder subjectDescriptor(String subjectDescriptor) {
            if (subjectDescriptor == null) {
              throw new MissingRequiredPropertyException("GetIdentityUsersResult", "subjectDescriptor");
            }
            this.subjectDescriptor = subjectDescriptor;
            return this;
        }
        public GetIdentityUsersResult build() {
            final var _resultValue = new GetIdentityUsersResult();
            _resultValue.descriptor = descriptor;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.searchFilter = searchFilter;
            _resultValue.subjectDescriptor = subjectDescriptor;
            return _resultValue;
        }
    }
}
