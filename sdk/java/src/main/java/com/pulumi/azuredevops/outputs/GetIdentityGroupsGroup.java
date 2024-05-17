// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIdentityGroupsGroup {
    private String id;
    /**
     * @return This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
     * 
     */
    private String name;

    private GetIdentityGroupsGroup() {}
    public String id() {
        return this.id;
    }
    /**
     * @return This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityGroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetIdentityGroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupsGroup", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupsGroup", "name");
            }
            this.name = name;
            return this;
        }
        public GetIdentityGroupsGroup build() {
            final var _resultValue = new GetIdentityGroupsGroup();
            _resultValue.id = id;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
