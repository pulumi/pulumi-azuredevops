// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTeamsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTeamsArgs Empty = new GetTeamsArgs();

    /**
     * The Project ID. If no project ID all teams of the organization will be returned.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The Project ID. If no project ID all teams of the organization will be returned.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private GetTeamsArgs() {}

    private GetTeamsArgs(GetTeamsArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTeamsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTeamsArgs $;

        public Builder() {
            $ = new GetTeamsArgs();
        }

        public Builder(GetTeamsArgs defaults) {
            $ = new GetTeamsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The Project ID. If no project ID all teams of the organization will be returned.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The Project ID. If no project ID all teams of the organization will be returned.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetTeamsArgs build() {
            return $;
        }
    }

}