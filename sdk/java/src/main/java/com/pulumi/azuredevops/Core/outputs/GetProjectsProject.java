// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Core.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProject {
    /**
     * @return Name of the Project, if not specified all projects will be returned.
     * 
     */
    private String name;
    /**
     * @return The ID of the Project.
     * 
     */
    private String projectId;
    /**
     * @return Url to the full version of the object.
     * 
     */
    private String projectUrl;
    /**
     * @return State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
     * 
     * DataSource without specifying any arguments will return all projects.
     * 
     */
    private String state;

    private GetProjectsProject() {}
    /**
     * @return Name of the Project, if not specified all projects will be returned.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The ID of the Project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Url to the full version of the object.
     * 
     */
    public String projectUrl() {
        return this.projectUrl;
    }
    /**
     * @return State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
     * 
     * DataSource without specifying any arguments will return all projects.
     * 
     */
    public String state() {
        return this.state;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String projectId;
        private String projectUrl;
        private String state;
        public Builder() {}
        public Builder(GetProjectsProject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.projectUrl = defaults.projectUrl;
    	      this.state = defaults.state;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder projectUrl(String projectUrl) {
            if (projectUrl == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "projectUrl");
            }
            this.projectUrl = projectUrl;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "state");
            }
            this.state = state;
            return this;
        }
        public GetProjectsProject build() {
            final var _resultValue = new GetProjectsProject();
            _resultValue.name = name;
            _resultValue.projectId = projectId;
            _resultValue.projectUrl = projectUrl;
            _resultValue.state = state;
            return _resultValue;
        }
    }
}
