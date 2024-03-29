// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBuildDefinitionArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBuildDefinitionArgs Empty = new GetBuildDefinitionArgs();

    /**
     * The name of this Build Definition.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of this Build Definition.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The path of the build definition. Default to `\`.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path of the build definition. Default to `\`.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
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

    private GetBuildDefinitionArgs() {}

    private GetBuildDefinitionArgs(GetBuildDefinitionArgs $) {
        this.name = $.name;
        this.path = $.path;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBuildDefinitionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBuildDefinitionArgs $;

        public Builder() {
            $ = new GetBuildDefinitionArgs();
        }

        public Builder(GetBuildDefinitionArgs defaults) {
            $ = new GetBuildDefinitionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of this Build Definition.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of this Build Definition.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param path The path of the build definition. Default to `\`.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path of the build definition. Default to `\`.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
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

        public GetBuildDefinitionArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetBuildDefinitionArgs", "name");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetBuildDefinitionArgs", "projectId");
            }
            return $;
        }
    }

}
