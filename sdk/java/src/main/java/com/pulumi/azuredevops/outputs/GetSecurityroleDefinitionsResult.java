// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetSecurityroleDefinitionsDefinition;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetSecurityroleDefinitionsResult {
    /**
     * @return A list of existing Security Role Definitions in a Scope in your Azure DevOps Organization with details about every definition which includes. A `definitions` block as defined below.
     * 
     */
    private List<GetSecurityroleDefinitionsDefinition> definitions;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The scope of the Security Role Definition.
     * 
     */
    private String scope;

    private GetSecurityroleDefinitionsResult() {}
    /**
     * @return A list of existing Security Role Definitions in a Scope in your Azure DevOps Organization with details about every definition which includes. A `definitions` block as defined below.
     * 
     */
    public List<GetSecurityroleDefinitionsDefinition> definitions() {
        return this.definitions;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The scope of the Security Role Definition.
     * 
     */
    public String scope() {
        return this.scope;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecurityroleDefinitionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetSecurityroleDefinitionsDefinition> definitions;
        private String id;
        private String scope;
        public Builder() {}
        public Builder(GetSecurityroleDefinitionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.definitions = defaults.definitions;
    	      this.id = defaults.id;
    	      this.scope = defaults.scope;
        }

        @CustomType.Setter
        public Builder definitions(List<GetSecurityroleDefinitionsDefinition> definitions) {
            if (definitions == null) {
              throw new MissingRequiredPropertyException("GetSecurityroleDefinitionsResult", "definitions");
            }
            this.definitions = definitions;
            return this;
        }
        public Builder definitions(GetSecurityroleDefinitionsDefinition... definitions) {
            return definitions(List.of(definitions));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSecurityroleDefinitionsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder scope(String scope) {
            if (scope == null) {
              throw new MissingRequiredPropertyException("GetSecurityroleDefinitionsResult", "scope");
            }
            this.scope = scope;
            return this;
        }
        public GetSecurityroleDefinitionsResult build() {
            final var _resultValue = new GetSecurityroleDefinitionsResult();
            _resultValue.definitions = definitions;
            _resultValue.id = id;
            _resultValue.scope = scope;
            return _resultValue;
        }
    }
}
