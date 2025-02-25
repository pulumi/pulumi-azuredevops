// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIdentityGroupResult {
    /**
     * @return The descriptor of the identity group.
     * 
     */
    private String descriptor;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
     * 
     */
    private String name;
    private String projectId;

    private GetIdentityGroupResult() {}
    /**
     * @return The descriptor of the identity group.
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
    /**
     * @return This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
     * 
     */
    public String name() {
        return this.name;
    }
    public String projectId() {
        return this.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String descriptor;
        private String id;
        private String name;
        private String projectId;
        public Builder() {}
        public Builder(GetIdentityGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptor = defaults.descriptor;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
        }

        @CustomType.Setter
        public Builder descriptor(String descriptor) {
            if (descriptor == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "descriptor");
            }
            this.descriptor = descriptor;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        public GetIdentityGroupResult build() {
            final var _resultValue = new GetIdentityGroupResult();
            _resultValue.descriptor = descriptor;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.projectId = projectId;
            return _resultValue;
        }
    }
}
