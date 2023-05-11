// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GitRepositoryFileArgs extends com.pulumi.resources.ResourceArgs {

    public static final GitRepositoryFileArgs Empty = new GitRepositoryFileArgs();

    /**
     * Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
     * does not already exist.
     * 
     */
    @Import(name="branch")
    private @Nullable Output<String> branch;

    /**
     * @return Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
     * does not already exist.
     * 
     */
    public Optional<Output<String>> branch() {
        return Optional.ofNullable(this.branch);
    }

    /**
     * Commit message when adding or updating the managed file.
     * 
     */
    @Import(name="commitMessage")
    private @Nullable Output<String> commitMessage;

    /**
     * @return Commit message when adding or updating the managed file.
     * 
     */
    public Optional<Output<String>> commitMessage() {
        return Optional.ofNullable(this.commitMessage);
    }

    /**
     * The file content.
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The file content.
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The path of the file to manage.
     * 
     */
    @Import(name="file", required=true)
    private Output<String> file;

    /**
     * @return The path of the file to manage.
     * 
     */
    public Output<String> file() {
        return this.file;
    }

    /**
     * Enable overwriting existing files (defaults to `false`).
     * 
     */
    @Import(name="overwriteOnCreate")
    private @Nullable Output<Boolean> overwriteOnCreate;

    /**
     * @return Enable overwriting existing files (defaults to `false`).
     * 
     */
    public Optional<Output<Boolean>> overwriteOnCreate() {
        return Optional.ofNullable(this.overwriteOnCreate);
    }

    /**
     * The ID of the Git repository.
     * 
     */
    @Import(name="repositoryId", required=true)
    private Output<String> repositoryId;

    /**
     * @return The ID of the Git repository.
     * 
     */
    public Output<String> repositoryId() {
        return this.repositoryId;
    }

    private GitRepositoryFileArgs() {}

    private GitRepositoryFileArgs(GitRepositoryFileArgs $) {
        this.branch = $.branch;
        this.commitMessage = $.commitMessage;
        this.content = $.content;
        this.file = $.file;
        this.overwriteOnCreate = $.overwriteOnCreate;
        this.repositoryId = $.repositoryId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GitRepositoryFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GitRepositoryFileArgs $;

        public Builder() {
            $ = new GitRepositoryFileArgs();
        }

        public Builder(GitRepositoryFileArgs defaults) {
            $ = new GitRepositoryFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branch Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
         * does not already exist.
         * 
         * @return builder
         * 
         */
        public Builder branch(@Nullable Output<String> branch) {
            $.branch = branch;
            return this;
        }

        /**
         * @param branch Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
         * does not already exist.
         * 
         * @return builder
         * 
         */
        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        /**
         * @param commitMessage Commit message when adding or updating the managed file.
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(@Nullable Output<String> commitMessage) {
            $.commitMessage = commitMessage;
            return this;
        }

        /**
         * @param commitMessage Commit message when adding or updating the managed file.
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(String commitMessage) {
            return commitMessage(Output.of(commitMessage));
        }

        /**
         * @param content The file content.
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The file content.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param file The path of the file to manage.
         * 
         * @return builder
         * 
         */
        public Builder file(Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file The path of the file to manage.
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files (defaults to `false`).
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(@Nullable Output<Boolean> overwriteOnCreate) {
            $.overwriteOnCreate = overwriteOnCreate;
            return this;
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files (defaults to `false`).
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(Boolean overwriteOnCreate) {
            return overwriteOnCreate(Output.of(overwriteOnCreate));
        }

        /**
         * @param repositoryId The ID of the Git repository.
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(Output<String> repositoryId) {
            $.repositoryId = repositoryId;
            return this;
        }

        /**
         * @param repositoryId The ID of the Git repository.
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(String repositoryId) {
            return repositoryId(Output.of(repositoryId));
        }

        public GitRepositoryFileArgs build() {
            $.content = Objects.requireNonNull($.content, "expected parameter 'content' to be non-null");
            $.file = Objects.requireNonNull($.file, "expected parameter 'file' to be non-null");
            $.repositoryId = Objects.requireNonNull($.repositoryId, "expected parameter 'repositoryId' to be non-null");
            return $;
        }
    }

}