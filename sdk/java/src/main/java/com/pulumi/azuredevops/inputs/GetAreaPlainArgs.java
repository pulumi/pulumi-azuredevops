// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAreaPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAreaPlainArgs Empty = new GetAreaPlainArgs();

    /**
     * Read children nodes, _Depth_: 1, _Default_: `true`
     * 
     */
    @Import(name="fetchChildren")
    private @Nullable Boolean fetchChildren;

    /**
     * @return Read children nodes, _Depth_: 1, _Default_: `true`
     * 
     */
    public Optional<Boolean> fetchChildren() {
        return Optional.ofNullable(this.fetchChildren);
    }

    /**
     * The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
     * 
     */
    @Import(name="path")
    private @Nullable String path;

    /**
     * @return The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The project ID.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The project ID.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetAreaPlainArgs() {}

    private GetAreaPlainArgs(GetAreaPlainArgs $) {
        this.fetchChildren = $.fetchChildren;
        this.path = $.path;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAreaPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAreaPlainArgs $;

        public Builder() {
            $ = new GetAreaPlainArgs();
        }

        public Builder(GetAreaPlainArgs defaults) {
            $ = new GetAreaPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param fetchChildren Read children nodes, _Depth_: 1, _Default_: `true`
         * 
         * @return builder
         * 
         */
        public Builder fetchChildren(@Nullable Boolean fetchChildren) {
            $.fetchChildren = fetchChildren;
            return this;
        }

        /**
         * @param path The path to the Area; _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Area will be returned
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable String path) {
            $.path = path;
            return this;
        }

        /**
         * @param projectId The project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetAreaPlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetAreaPlainArgs", "projectId");
            }
            return $;
        }
    }

}
