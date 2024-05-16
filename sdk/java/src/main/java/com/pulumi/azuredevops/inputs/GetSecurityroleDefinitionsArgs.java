// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetSecurityroleDefinitionsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecurityroleDefinitionsArgs Empty = new GetSecurityroleDefinitionsArgs();

    /**
     * Name of the Scope for which Security Role Definitions will be returned.
     * 
     * DataSource without specifying any arguments will return all projects.
     * 
     */
    @Import(name="scope", required=true)
    private Output<String> scope;

    /**
     * @return Name of the Scope for which Security Role Definitions will be returned.
     * 
     * DataSource without specifying any arguments will return all projects.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    private GetSecurityroleDefinitionsArgs() {}

    private GetSecurityroleDefinitionsArgs(GetSecurityroleDefinitionsArgs $) {
        this.scope = $.scope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecurityroleDefinitionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecurityroleDefinitionsArgs $;

        public Builder() {
            $ = new GetSecurityroleDefinitionsArgs();
        }

        public Builder(GetSecurityroleDefinitionsArgs defaults) {
            $ = new GetSecurityroleDefinitionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param scope Name of the Scope for which Security Role Definitions will be returned.
         * 
         * DataSource without specifying any arguments will return all projects.
         * 
         * @return builder
         * 
         */
        public Builder scope(Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope Name of the Scope for which Security Role Definitions will be returned.
         * 
         * DataSource without specifying any arguments will return all projects.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public GetSecurityroleDefinitionsArgs build() {
            if ($.scope == null) {
                throw new MissingRequiredPropertyException("GetSecurityroleDefinitionsArgs", "scope");
            }
            return $;
        }
    }

}
