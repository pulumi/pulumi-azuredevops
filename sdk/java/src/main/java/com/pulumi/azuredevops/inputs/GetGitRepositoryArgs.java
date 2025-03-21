// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGitRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGitRepositoryArgs Empty = new GetGitRepositoryArgs();

    /**
     * The Name of the Git repository to retrieve
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The Name of the Git repository to retrieve
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The ID of project to list Git repositories
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of project to list Git repositories
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetGitRepositoryArgs() {}

    private GetGitRepositoryArgs(GetGitRepositoryArgs $) {
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGitRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGitRepositoryArgs $;

        public Builder() {
            $ = new GetGitRepositoryArgs();
        }

        public Builder(GetGitRepositoryArgs defaults) {
            $ = new GetGitRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The Name of the Git repository to retrieve
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The Name of the Git repository to retrieve
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The ID of project to list Git repositories
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of project to list Git repositories
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetGitRepositoryArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetGitRepositoryArgs", "name");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetGitRepositoryArgs", "projectId");
            }
            return $;
        }
    }

}
