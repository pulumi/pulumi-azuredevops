// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTeamArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTeamArgs Empty = new GetTeamArgs();

    /**
     * The name of the Team.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the Team.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The Project ID.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The Project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The maximum number of teams to return. Defaults to `100`.
     * 
     */
    @Import(name="top")
    private @Nullable Output<Integer> top;

    /**
     * @return The maximum number of teams to return. Defaults to `100`.
     * 
     */
    public Optional<Output<Integer>> top() {
        return Optional.ofNullable(this.top);
    }

    private GetTeamArgs() {}

    private GetTeamArgs(GetTeamArgs $) {
        this.name = $.name;
        this.projectId = $.projectId;
        this.top = $.top;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTeamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTeamArgs $;

        public Builder() {
            $ = new GetTeamArgs();
        }

        public Builder(GetTeamArgs defaults) {
            $ = new GetTeamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the Team.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Team.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The Project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The Project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param top The maximum number of teams to return. Defaults to `100`.
         * 
         * @return builder
         * 
         */
        public Builder top(@Nullable Output<Integer> top) {
            $.top = top;
            return this;
        }

        /**
         * @param top The maximum number of teams to return. Defaults to `100`.
         * 
         * @return builder
         * 
         */
        public Builder top(Integer top) {
            return top(Output.of(top));
        }

        public GetTeamArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetTeamArgs", "name");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetTeamArgs", "projectId");
            }
            return $;
        }
    }

}
