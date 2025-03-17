// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class BuildDefinitionJobDependencyArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildDefinitionJobDependencyArgs Empty = new BuildDefinitionJobDependencyArgs();

    /**
     * The job reference name that depends on. Reference to `jobs.ref_name`
     * 
     */
    @Import(name="scope", required=true)
    private Output<String> scope;

    /**
     * @return The job reference name that depends on. Reference to `jobs.ref_name`
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    private BuildDefinitionJobDependencyArgs() {}

    private BuildDefinitionJobDependencyArgs(BuildDefinitionJobDependencyArgs $) {
        this.scope = $.scope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildDefinitionJobDependencyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildDefinitionJobDependencyArgs $;

        public Builder() {
            $ = new BuildDefinitionJobDependencyArgs();
        }

        public Builder(BuildDefinitionJobDependencyArgs defaults) {
            $ = new BuildDefinitionJobDependencyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param scope The job reference name that depends on. Reference to `jobs.ref_name`
         * 
         * @return builder
         * 
         */
        public Builder scope(Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope The job reference name that depends on. Reference to `jobs.ref_name`
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public BuildDefinitionJobDependencyArgs build() {
            if ($.scope == null) {
                throw new MissingRequiredPropertyException("BuildDefinitionJobDependencyArgs", "scope");
            }
            return $;
        }
    }

}
