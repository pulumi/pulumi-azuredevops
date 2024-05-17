// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGitRepositoryResult {
    /**
     * @return The ref of the default branch.
     * 
     */
    private String defaultBranch;
    /**
     * @return Is the repository disabled?
     * 
     */
    private Boolean disabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Boolean isFork;
    /**
     * @return Git repository name.
     * 
     */
    private String name;
    /**
     * @return Project identifier to which the Git repository belongs.
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

    private GetGitRepositoryResult() {}
    /**
     * @return The ref of the default branch.
     * 
     */
    public String defaultBranch() {
        return this.defaultBranch;
    }
    /**
     * @return Is the repository disabled?
     * 
     */
    public Boolean disabled() {
        return this.disabled;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean isFork() {
        return this.isFork;
    }
    /**
     * @return Git repository name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Project identifier to which the Git repository belongs.
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

    public static Builder builder(GetGitRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String defaultBranch;
        private Boolean disabled;
        private String id;
        private Boolean isFork;
        private String name;
        private String projectId;
        private String remoteUrl;
        private Integer size;
        private String sshUrl;
        private String url;
        private String webUrl;
        public Builder() {}
        public Builder(GetGitRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultBranch = defaults.defaultBranch;
    	      this.disabled = defaults.disabled;
    	      this.id = defaults.id;
    	      this.isFork = defaults.isFork;
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
            if (defaultBranch == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "defaultBranch");
            }
            this.defaultBranch = defaultBranch;
            return this;
        }
        @CustomType.Setter
        public Builder disabled(Boolean disabled) {
            if (disabled == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "disabled");
            }
            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isFork(Boolean isFork) {
            if (isFork == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "isFork");
            }
            this.isFork = isFork;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder remoteUrl(String remoteUrl) {
            if (remoteUrl == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "remoteUrl");
            }
            this.remoteUrl = remoteUrl;
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            if (size == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "size");
            }
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder sshUrl(String sshUrl) {
            if (sshUrl == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "sshUrl");
            }
            this.sshUrl = sshUrl;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder webUrl(String webUrl) {
            if (webUrl == null) {
              throw new MissingRequiredPropertyException("GetGitRepositoryResult", "webUrl");
            }
            this.webUrl = webUrl;
            return this;
        }
        public GetGitRepositoryResult build() {
            final var _resultValue = new GetGitRepositoryResult();
            _resultValue.defaultBranch = defaultBranch;
            _resultValue.disabled = disabled;
            _resultValue.id = id;
            _resultValue.isFork = isFork;
            _resultValue.name = name;
            _resultValue.projectId = projectId;
            _resultValue.remoteUrl = remoteUrl;
            _resultValue.size = size;
            _resultValue.sshUrl = sshUrl;
            _resultValue.url = url;
            _resultValue.webUrl = webUrl;
            return _resultValue;
        }
    }
}
