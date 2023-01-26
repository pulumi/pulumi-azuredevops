// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetVariableGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVariableGroupArgs Empty = new GetVariableGroupArgs();

    /**
     * The name of the Variable Group to retrieve.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the Variable Group to retrieve.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The project ID.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetVariableGroupArgs() {}

    private GetVariableGroupArgs(GetVariableGroupArgs $) {
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVariableGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVariableGroupArgs $;

        public Builder() {
            $ = new GetVariableGroupArgs();
        }

        public Builder(GetVariableGroupArgs defaults) {
            $ = new GetVariableGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the Variable Group to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Variable Group to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetVariableGroupArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
