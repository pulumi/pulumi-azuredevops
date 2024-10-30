// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WikiPageState extends com.pulumi.resources.ResourceArgs {

    public static final WikiPageState Empty = new WikiPageState();

    /**
     * The content of the wiki page.
     * 
     */
    @Import(name="content")
    private @Nullable Output<String> content;

    /**
     * @return The content of the wiki page.
     * 
     */
    public Optional<Output<String>> content() {
        return Optional.ofNullable(this.content);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The path of the wiki page.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path of the wiki page.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The ID of the Project.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the Project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The ID of the Wiki.
     * 
     */
    @Import(name="wikiId")
    private @Nullable Output<String> wikiId;

    /**
     * @return The ID of the Wiki.
     * 
     */
    public Optional<Output<String>> wikiId() {
        return Optional.ofNullable(this.wikiId);
    }

    private WikiPageState() {}

    private WikiPageState(WikiPageState $) {
        this.content = $.content;
        this.etag = $.etag;
        this.path = $.path;
        this.projectId = $.projectId;
        this.wikiId = $.wikiId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WikiPageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WikiPageState $;

        public Builder() {
            $ = new WikiPageState();
        }

        public Builder(WikiPageState defaults) {
            $ = new WikiPageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param content The content of the wiki page.
         * 
         * @return builder
         * 
         */
        public Builder content(@Nullable Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The content of the wiki page.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param path The path of the wiki page.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path of the wiki page.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param projectId The ID of the Project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the Project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param wikiId The ID of the Wiki.
         * 
         * @return builder
         * 
         */
        public Builder wikiId(@Nullable Output<String> wikiId) {
            $.wikiId = wikiId;
            return this;
        }

        /**
         * @param wikiId The ID of the Wiki.
         * 
         * @return builder
         * 
         */
        public Builder wikiId(String wikiId) {
            return wikiId(Output.of(wikiId));
        }

        public WikiPageState build() {
            return $;
        }
    }

}
