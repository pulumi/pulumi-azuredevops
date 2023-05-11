// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAreaArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAreaArgs Empty = new GetAreaArgs();

    /**
     * Read children nodes, _Depth_: 1, _Default_: `true`
     * 
     */
    @Import(name="fetchChildren")
    private @Nullable Output<Boolean> fetchChildren;

    /**
     * @return Read children nodes, _Depth_: 1, _Default_: `true`
     * 
     */
    public Optional<Output<Boolean>> fetchChildren() {
        return Optional.ofNullable(this.fetchChildren);
    }

    /**
     * The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
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

    private GetAreaArgs() {}

    private GetAreaArgs(GetAreaArgs $) {
        this.fetchChildren = $.fetchChildren;
        this.path = $.path;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAreaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAreaArgs $;

        public Builder() {
            $ = new GetAreaArgs();
        }

        public Builder(GetAreaArgs defaults) {
            $ = new GetAreaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param fetchChildren Read children nodes, _Depth_: 1, _Default_: `true`
         * 
         * @return builder
         * 
         */
        public Builder fetchChildren(@Nullable Output<Boolean> fetchChildren) {
            $.fetchChildren = fetchChildren;
            return this;
        }

        /**
         * @param fetchChildren Read children nodes, _Depth_: 1, _Default_: `true`
         * 
         * @return builder
         * 
         */
        public Builder fetchChildren(Boolean fetchChildren) {
            return fetchChildren(Output.of(fetchChildren));
        }

        /**
         * @param path The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
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

        public GetAreaArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}