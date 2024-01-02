// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetEnvironmentArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnvironmentArgs Empty = new GetEnvironmentArgs();

    /**
     * The ID of the Environment.
     * 
     */
    @Import(name="environmentId", required=true)
    private Output<Integer> environmentId;

    /**
     * @return The ID of the Environment.
     * 
     */
    public Output<Integer> environmentId() {
        return this.environmentId;
    }

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetEnvironmentArgs() {}

    private GetEnvironmentArgs(GetEnvironmentArgs $) {
        this.environmentId = $.environmentId;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnvironmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnvironmentArgs $;

        public Builder() {
            $ = new GetEnvironmentArgs();
        }

        public Builder(GetEnvironmentArgs defaults) {
            $ = new GetEnvironmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param environmentId The ID of the Environment.
         * 
         * @return builder
         * 
         */
        public Builder environmentId(Output<Integer> environmentId) {
            $.environmentId = environmentId;
            return this;
        }

        /**
         * @param environmentId The ID of the Environment.
         * 
         * @return builder
         * 
         */
        public Builder environmentId(Integer environmentId) {
            return environmentId(Output.of(environmentId));
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetEnvironmentArgs build() {
            if ($.environmentId == null) {
                throw new MissingRequiredPropertyException("GetEnvironmentArgs", "environmentId");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetEnvironmentArgs", "projectId");
            }
            return $;
        }
    }

}
