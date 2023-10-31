// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.Core.outputs;

import com.pulumi.azuredevops.Core.outputs.GetProjectsProject;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the Project.
     * 
     */
    private @Nullable String name;
    /**
     * @return A list of existing projects in your Azure DevOps Organization with details about every project which includes:
     * 
     */
    private List<GetProjectsProject> projects;
    /**
     * @return Project state.
     * 
     */
    private @Nullable String state;

    private GetProjectsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the Project.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return A list of existing projects in your Azure DevOps Organization with details about every project which includes:
     * 
     */
    public List<GetProjectsProject> projects() {
        return this.projects;
    }
    /**
     * @return Project state.
     * 
     */
    public Optional<String> state() {
        return Optional.ofNullable(this.state);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String name;
        private List<GetProjectsProject> projects;
        private @Nullable String state;
        public Builder() {}
        public Builder(GetProjectsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projects = defaults.projects;
    	      this.state = defaults.state;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projects(List<GetProjectsProject> projects) {
            this.projects = Objects.requireNonNull(projects);
            return this;
        }
        public Builder projects(GetProjectsProject... projects) {
            return projects(List.of(projects));
        }
        @CustomType.Setter
        public Builder state(@Nullable String state) {
            this.state = state;
            return this;
        }
        public GetProjectsResult build() {
            final var _resultValue = new GetProjectsResult();
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.projects = projects;
            _resultValue.state = state;
            return _resultValue;
        }
    }
}
