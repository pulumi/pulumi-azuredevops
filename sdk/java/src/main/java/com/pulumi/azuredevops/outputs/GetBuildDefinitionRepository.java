// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBuildDefinitionRepository {
    /**
     * @return The branch name for which builds are triggered.
     * 
     */
    private String branchName;
    /**
     * @return The Github Enterprise URL.
     * 
     */
    private String githubEnterpriseUrl;
    /**
     * @return The id of the repository.
     * 
     */
    private String repoId;
    /**
     * @return The repository type.
     * 
     */
    private String repoType;
    /**
     * @return Report build status.
     * 
     */
    private Boolean reportBuildStatus;
    /**
     * @return The service connection ID.
     * 
     */
    private String serviceConnectionId;
    /**
     * @return The path of the Yaml file describing the build definition.
     * 
     */
    private String ymlPath;

    private GetBuildDefinitionRepository() {}
    /**
     * @return The branch name for which builds are triggered.
     * 
     */
    public String branchName() {
        return this.branchName;
    }
    /**
     * @return The Github Enterprise URL.
     * 
     */
    public String githubEnterpriseUrl() {
        return this.githubEnterpriseUrl;
    }
    /**
     * @return The id of the repository.
     * 
     */
    public String repoId() {
        return this.repoId;
    }
    /**
     * @return The repository type.
     * 
     */
    public String repoType() {
        return this.repoType;
    }
    /**
     * @return Report build status.
     * 
     */
    public Boolean reportBuildStatus() {
        return this.reportBuildStatus;
    }
    /**
     * @return The service connection ID.
     * 
     */
    public String serviceConnectionId() {
        return this.serviceConnectionId;
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

    public static Builder builder(GetBuildDefinitionRepository defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String branchName;
        private String githubEnterpriseUrl;
        private String repoId;
        private String repoType;
        private Boolean reportBuildStatus;
        private String serviceConnectionId;
        private String ymlPath;
        public Builder() {}
        public Builder(GetBuildDefinitionRepository defaults) {
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
        public Builder branchName(String branchName) {
            this.branchName = Objects.requireNonNull(branchName);
            return this;
        }
        @CustomType.Setter
        public Builder githubEnterpriseUrl(String githubEnterpriseUrl) {
            this.githubEnterpriseUrl = Objects.requireNonNull(githubEnterpriseUrl);
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
        public Builder reportBuildStatus(Boolean reportBuildStatus) {
            this.reportBuildStatus = Objects.requireNonNull(reportBuildStatus);
            return this;
        }
        @CustomType.Setter
        public Builder serviceConnectionId(String serviceConnectionId) {
            this.serviceConnectionId = Objects.requireNonNull(serviceConnectionId);
            return this;
        }
        @CustomType.Setter
        public Builder ymlPath(String ymlPath) {
            this.ymlPath = Objects.requireNonNull(ymlPath);
            return this;
        }
        public GetBuildDefinitionRepository build() {
            final var o = new GetBuildDefinitionRepository();
            o.branchName = branchName;
            o.githubEnterpriseUrl = githubEnterpriseUrl;
            o.repoId = repoId;
            o.repoType = repoType;
            o.reportBuildStatus = reportBuildStatus;
            o.serviceConnectionId = serviceConnectionId;
            o.ymlPath = ymlPath;
            return o;
        }
    }
}