// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetEnvironmentPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnvironmentPlainArgs Empty = new GetEnvironmentPlainArgs();

    /**
     * The ID of the Environment.
     * 
     */
    @Import(name="environmentId", required=true)
    private Integer environmentId;

    /**
     * @return The ID of the Environment.
     * 
     */
    public Integer environmentId() {
        return this.environmentId;
    }

    /**
     * The ID of the project.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetEnvironmentPlainArgs() {}

    private GetEnvironmentPlainArgs(GetEnvironmentPlainArgs $) {
        this.environmentId = $.environmentId;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnvironmentPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnvironmentPlainArgs $;

        public Builder() {
            $ = new GetEnvironmentPlainArgs();
        }

        public Builder(GetEnvironmentPlainArgs defaults) {
            $ = new GetEnvironmentPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param environmentId The ID of the Environment.
         * 
         * @return builder
         * 
         */
        public Builder environmentId(Integer environmentId) {
            $.environmentId = environmentId;
            return this;
        }

        /**
         * @param projectId The ID of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetEnvironmentPlainArgs build() {
            if ($.environmentId == null) {
                throw new MissingRequiredPropertyException("GetEnvironmentPlainArgs", "environmentId");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetEnvironmentPlainArgs", "projectId");
            }
            return $;
        }
    }

}
