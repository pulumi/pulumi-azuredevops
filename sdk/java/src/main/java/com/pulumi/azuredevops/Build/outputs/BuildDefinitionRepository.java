// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Build.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BuildDefinitionRepository {
    /**
     * @return The branch name for which builds are triggered. Defaults to `master`.
     * 
     */
    private @Nullable String branchName;
    /**
     * @return The Github Enterprise URL. Used if `repo_type` is `GithubEnterprise`.
     * 
     */
    private @Nullable String githubEnterpriseUrl;
    /**
     * @return The id of the repository. For `TfsGit` repos, this is simply the ID of the repository. For `Github` repos, this will take the form of `&lt;GitHub Org&gt;/&lt;Repo Name&gt;`. For `Bitbucket` repos, this will take the form of `&lt;Workspace ID&gt;/&lt;Repo Name&gt;`.
     * 
     */
    private String repoId;
    /**
     * @return The repository type. Valid values: `GitHub` or `TfsGit` or `Bitbucket` or `GitHub Enterprise`. Defaults to `GitHub`. If `repo_type` is `GitHubEnterprise`, must use existing project and GitHub Enterprise service connection.
     * 
     */
    private String repoType;
    /**
     * @return Report build status. Default is true.
     * 
     */
    private @Nullable Boolean reportBuildStatus;
    /**
     * @return The service connection ID. Used if the `repo_type` is `GitHub` or `GitHubEnterprise`.
     * 
     */
    private @Nullable String serviceConnectionId;
    /**
     * @return The path of the Yaml file describing the build definition.
     * 
     */
    private String ymlPath;

    private BuildDefinitionRepository() {}
    /**
     * @return The branch name for which builds are triggered. Defaults to `master`.
     * 
     */
    public Optional<String> branchName() {
        return Optional.ofNullable(this.branchName);
    }
    /**
     * @return The Github Enterprise URL. Used if `repo_type` is `GithubEnterprise`.
     * 
     */
    public Optional<String> githubEnterpriseUrl() {
        return Optional.ofNullable(this.githubEnterpriseUrl);
    }
    /**
     * @return The id of the repository. For `TfsGit` repos, this is simply the ID of the repository. For `Github` repos, this will take the form of `&lt;GitHub Org&gt;/&lt;Repo Name&gt;`. For `Bitbucket` repos, this will take the form of `&lt;Workspace ID&gt;/&lt;Repo Name&gt;`.
     * 
     */
    public String repoId() {
        return this.repoId;
    }
    /**
     * @return The repository type. Valid values: `GitHub` or `TfsGit` or `Bitbucket` or `GitHub Enterprise`. Defaults to `GitHub`. If `repo_type` is `GitHubEnterprise`, must use existing project and GitHub Enterprise service connection.
     * 
     */
    public String repoType() {
        return this.repoType;
    }
    /**
     * @return Report build status. Default is true.
     * 
     */
    public Optional<Boolean> reportBuildStatus() {
        return Optional.ofNullable(this.reportBuildStatus);
    }
    /**
     * @return The service connection ID. Used if the `repo_type` is `GitHub` or `GitHubEnterprise`.
     * 
     */
    public Optional<String> serviceConnectionId() {
        return Optional.ofNullable(this.serviceConnectionId);
    }
    /**
     * @return The path of the Yaml file describing the build definition.
     * 
     */
    public String ymlPath() {
        return this.ymlPath;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BuildDefinitionRepository defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String branchName;
        private @Nullable String githubEnterpriseUrl;
        private String repoId;
        private String repoType;
        private @Nullable Boolean reportBuildStatus;
        private @Nullable String serviceConnectionId;
        private String ymlPath;
        public Builder() {}
        public Builder(BuildDefinitionRepository defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branchName = defaults.branchName;
    	      this.githubEnterpriseUrl = defaults.githubEnterpriseUrl;
    	      this.repoId = defaults.repoId;
    	      this.repoType = defaults.repoType;
    	      this.reportBuildStatus = defaults.reportBuildStatus;
    	      this.serviceConnectionId = defaults.serviceConnectionId;
    	      this.ymlPath = defaults.ymlPath;
        }

        @CustomType.Setter
        public Builder branchName(@Nullable String branchName) {
            this.branchName = branchName;
            return this;
        }
        @CustomType.Setter
        public Builder githubEnterpriseUrl(@Nullable String githubEnterpriseUrl) {
            this.githubEnterpriseUrl = githubEnterpriseUrl;
            return this;
        }
        @CustomType.Setter
        public Builder repoId(String repoId) {
            this.repoId = Objects.requireNonNull(repoId);
            return this;
        }
        @CustomType.Setter
        public Builder repoType(String repoType) {
            this.repoType = Objects.requireNonNull(repoType);
            return this;
        }
        @CustomType.Setter
        public Builder reportBuildStatus(@Nullable Boolean reportBuildStatus) {
            this.reportBuildStatus = reportBuildStatus;
            return this;
        }
        @CustomType.Setter
        public Builder serviceConnectionId(@Nullable String serviceConnectionId) {
            this.serviceConnectionId = serviceConnectionId;
            return this;
        }
        @CustomType.Setter
        public Builder ymlPath(String ymlPath) {
            this.ymlPath = Objects.requireNonNull(ymlPath);
            return this;
        }
        public BuildDefinitionRepository build() {
            final var _resultValue = new BuildDefinitionRepository();
            _resultValue.branchName = branchName;
            _resultValue.githubEnterpriseUrl = githubEnterpriseUrl;
            _resultValue.repoId = repoId;
            _resultValue.repoType = repoType;
            _resultValue.reportBuildStatus = reportBuildStatus;
            _resultValue.serviceConnectionId = serviceConnectionId;
            _resultValue.ymlPath = ymlPath;
            return _resultValue;
        }
    }
}
