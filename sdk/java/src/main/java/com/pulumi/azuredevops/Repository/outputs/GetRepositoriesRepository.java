// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Repository.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoriesRepository {
    /**
     * @return The ref of the default branch.
     * 
     */
    private String defaultBranch;
    /**
     * @return Git repository identifier.
     * 
     */
    private String id;
    /**
     * @return Name of the Git repository to retrieve; requires `project_id` to be specified as well
     * 
     */
    private String name;
    /**
     * @return ID of project to list Git repositories
     * 
     */
    private String projectId;
    /**
     * @return HTTPS Url to clone the Git repository
     * 
     */
    private String remoteUrl;
    /**
     * @return Compressed size (bytes) of the repository.
     * 
     */
    private Integer size;
    /**
     * @return SSH Url to clone the Git repository
     * 
     */
    private String sshUrl;
    /**
     * @return Details REST API endpoint for the Git Repository.
     * 
     */
    private String url;
    /**
     * @return Url of the Git repository web view
     * 
     */
    private String webUrl;

    private GetRepositoriesRepository() {}
    /**
     * @return The ref of the default branch.
     * 
     */
    public String defaultBranch() {
        return this.defaultBranch;
    }
    /**
     * @return Git repository identifier.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the Git repository to retrieve; requires `project_id` to be specified as well
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return ID of project to list Git repositories
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return HTTPS Url to clone the Git repository
     * 
     */
    public String remoteUrl() {
        return this.remoteUrl;
    }
    /**
     * @return Compressed size (bytes) of the repository.
     * 
     */
    public Integer size() {
        return this.size;
    }
    /**
     * @return SSH Url to clone the Git repository
     * 
     */
    public String sshUrl() {
        return this.sshUrl;
    }
    /**
     * @return Details REST API endpoint for the Git Repository.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Url of the Git repository web view
     * 
     */
    public String webUrl() {
        return this.webUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoriesRepository defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String defaultBranch;
        private String id;
        private String name;
        private String projectId;
        private String remoteUrl;
        private Integer size;
        private String sshUrl;
        private String url;
        private String webUrl;
        public Builder() {}
        public Builder(GetRepositoriesRepository defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultBranch = defaults.defaultBranch;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.remoteUrl = defaults.remoteUrl;
    	      this.size = defaults.size;
    	      this.sshUrl = defaults.sshUrl;
    	      this.url = defaults.url;
    	      this.webUrl = defaults.webUrl;
        }

        @CustomType.Setter
        public Builder defaultBranch(String defaultBranch) {
            this.defaultBranch = Objects.requireNonNull(defaultBranch);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder remoteUrl(String remoteUrl) {
            this.remoteUrl = Objects.requireNonNull(remoteUrl);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder sshUrl(String sshUrl) {
            this.sshUrl = Objects.requireNonNull(sshUrl);
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            this.webUrl = Objects.requireNonNull(webUrl);
            return this;
        }
        public GetRepositoriesRepository build() {
            final var o = new GetRepositoriesRepository();
            o.defaultBranch = defaultBranch;
            o.id = id;
            o.name = name;
            o.projectId = projectId;
            o.remoteUrl = remoteUrl;
            o.size = size;
            o.sshUrl = sshUrl;
            o.url = url;
            o.webUrl = webUrl;
            return o;
        }
    }
}
