// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.azuredevops.inputs.GitInitializationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GitState extends com.pulumi.resources.ResourceArgs {

    public static final GitState Empty = new GitState();

    /**
     * The ref of the default branch. Will be used as the branch name for initialized repositories.
     * 
     */
    @Import(name="defaultBranch")
    private @Nullable Output<String> defaultBranch;

    /**
     * @return The ref of the default branch. Will be used as the branch name for initialized repositories.
     * 
     */
    public Optional<Output<String>> defaultBranch() {
        return Optional.ofNullable(this.defaultBranch);
    }

    /**
     * Is the repository disabled?
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Is the repository disabled?
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * An `initialization` block as documented below.
     * 
     */
    @Import(name="initialization")
    private @Nullable Output<GitInitializationArgs> initialization;

    /**
     * @return An `initialization` block as documented below.
     * 
     */
    public Optional<Output<GitInitializationArgs>> initialization() {
        return Optional.ofNullable(this.initialization);
    }

    /**
     * True if the repository was created as a fork.
     * 
     */
    @Import(name="isFork")
    private @Nullable Output<Boolean> isFork;

    /**
     * @return True if the repository was created as a fork.
     * 
     */
    public Optional<Output<Boolean>> isFork() {
        return Optional.ofNullable(this.isFork);
    }

    /**
     * The name of the git repository.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the git repository.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of a Git project from which a fork is to be created.
     * 
     */
    @Import(name="parentRepositoryId")
    private @Nullable Output<String> parentRepositoryId;

    /**
     * @return The ID of a Git project from which a fork is to be created.
     * 
     */
    public Optional<Output<String>> parentRepositoryId() {
        return Optional.ofNullable(this.parentRepositoryId);
    }

    /**
     * The project ID or project name.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project ID or project name.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Git HTTPS URL of the repository
     * 
     */
    @Import(name="remoteUrl")
    private @Nullable Output<String> remoteUrl;

    /**
     * @return Git HTTPS URL of the repository
     * 
     */
    public Optional<Output<String>> remoteUrl() {
        return Optional.ofNullable(this.remoteUrl);
    }

    /**
     * Size in bytes.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return Size in bytes.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * Git SSH URL of the repository.
     * 
     */
    @Import(name="sshUrl")
    private @Nullable Output<String> sshUrl;

    /**
     * @return Git SSH URL of the repository.
     * 
     */
    public Optional<Output<String>> sshUrl() {
        return Optional.ofNullable(this.sshUrl);
    }

    /**
     * REST API URL of the repository.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return REST API URL of the repository.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * Web link to the repository.
     * 
     */
    @Import(name="webUrl")
    private @Nullable Output<String> webUrl;

    /**
     * @return Web link to the repository.
     * 
     */
    public Optional<Output<String>> webUrl() {
        return Optional.ofNullable(this.webUrl);
    }

    private GitState() {}

    private GitState(GitState $) {
        this.defaultBranch = $.defaultBranch;
        this.disabled = $.disabled;
        this.initialization = $.initialization;
        this.isFork = $.isFork;
        this.name = $.name;
        this.parentRepositoryId = $.parentRepositoryId;
        this.projectId = $.projectId;
        this.remoteUrl = $.remoteUrl;
        this.size = $.size;
        this.sshUrl = $.sshUrl;
        this.url = $.url;
        this.webUrl = $.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GitState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GitState $;

        public Builder() {
            $ = new GitState();
        }

        public Builder(GitState defaults) {
            $ = new GitState(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultBranch The ref of the default branch. Will be used as the branch name for initialized repositories.
         * 
         * @return builder
         * 
         */
        public Builder defaultBranch(@Nullable Output<String> defaultBranch) {
            $.defaultBranch = defaultBranch;
            return this;
        }

        /**
         * @param defaultBranch The ref of the default branch. Will be used as the branch name for initialized repositories.
         * 
         * @return builder
         * 
         */
        public Builder defaultBranch(String defaultBranch) {
            return defaultBranch(Output.of(defaultBranch));
        }

        /**
         * @param disabled Is the repository disabled?
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Is the repository disabled?
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param initialization An `initialization` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder initialization(@Nullable Output<GitInitializationArgs> initialization) {
            $.initialization = initialization;
            return this;
        }

        /**
         * @param initialization An `initialization` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder initialization(GitInitializationArgs initialization) {
            return initialization(Output.of(initialization));
        }

        /**
         * @param isFork True if the repository was created as a fork.
         * 
         * @return builder
         * 
         */
        public Builder isFork(@Nullable Output<Boolean> isFork) {
            $.isFork = isFork;
            return this;
        }

        /**
         * @param isFork True if the repository was created as a fork.
         * 
         * @return builder
         * 
         */
        public Builder isFork(Boolean isFork) {
            return isFork(Output.of(isFork));
        }

        /**
         * @param name The name of the git repository.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the git repository.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentRepositoryId The ID of a Git project from which a fork is to be created.
         * 
         * @return builder
         * 
         */
        public Builder parentRepositoryId(@Nullable Output<String> parentRepositoryId) {
            $.parentRepositoryId = parentRepositoryId;
            return this;
        }

        /**
         * @param parentRepositoryId The ID of a Git project from which a fork is to be created.
         * 
         * @return builder
         * 
         */
        public Builder parentRepositoryId(String parentRepositoryId) {
            return parentRepositoryId(Output.of(parentRepositoryId));
        }

        /**
         * @param projectId The project ID or project name.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project ID or project name.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param remoteUrl Git HTTPS URL of the repository
         * 
         * @return builder
         * 
         */
        public Builder remoteUrl(@Nullable Output<String> remoteUrl) {
            $.remoteUrl = remoteUrl;
            return this;
        }

        /**
         * @param remoteUrl Git HTTPS URL of the repository
         * 
         * @return builder
         * 
         */
        public Builder remoteUrl(String remoteUrl) {
            return remoteUrl(Output.of(remoteUrl));
        }

        /**
         * @param size Size in bytes.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size in bytes.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param sshUrl Git SSH URL of the repository.
         * 
         * @return builder
         * 
         */
        public Builder sshUrl(@Nullable Output<String> sshUrl) {
            $.sshUrl = sshUrl;
            return this;
        }

        /**
         * @param sshUrl Git SSH URL of the repository.
         * 
         * @return builder
         * 
         */
        public Builder sshUrl(String sshUrl) {
            return sshUrl(Output.of(sshUrl));
        }

        /**
         * @param url REST API URL of the repository.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url REST API URL of the repository.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param webUrl Web link to the repository.
         * 
         * @return builder
         * 
         */
        public Builder webUrl(@Nullable Output<String> webUrl) {
            $.webUrl = webUrl;
            return this;
        }

        /**
         * @param webUrl Web link to the repository.
         * 
         * @return builder
         * 
         */
        public Builder webUrl(String webUrl) {
            return webUrl(Output.of(webUrl));
        }

        public GitState build() {
            return $;
        }
    }

}
