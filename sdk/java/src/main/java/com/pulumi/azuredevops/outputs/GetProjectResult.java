// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectResult {
    /**
     * @return The description of the project.
     * 
     */
    private String description;
    private Map<String,String> features;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the project.
     * 
     */
    private @Nullable String name;
    /**
     * @return The process template ID for the project.
     * 
     */
    private String processTemplateId;
    private @Nullable String projectId;
    /**
     * @return The version control of the project.
     * 
     */
    private String versionControl;
    /**
     * @return The visibility of the project.
     * 
     */
    private String visibility;
    /**
     * @return The work item template for the project.
     * 
     */
    private String workItemTemplate;

    private GetProjectResult() {}
    /**
     * @return The description of the project.
     * 
     */
    public String description() {
        return this.description;
    }
    public Map<String,String> features() {
        return this.features;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the project.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The process template ID for the project.
     * 
     */
    public String processTemplateId() {
        return this.processTemplateId;
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    /**
     * @return The version control of the project.
     * 
     */
    public String versionControl() {
        return this.versionControl;
    }
    /**
     * @return The visibility of the project.
     * 
     */
    public String visibility() {
        return this.visibility;
    }
    /**
     * @return The work item template for the project.
     * 
     */
    public String workItemTemplate() {
        return this.workItemTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Map<String,String> features;
        private String id;
        private @Nullable String name;
        private String processTemplateId;
        private @Nullable String projectId;
        private String versionControl;
        private String visibility;
        private String workItemTemplate;
        public Builder() {}
        public Builder(GetProjectResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.features = defaults.features;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.processTemplateId = defaults.processTemplateId;
    	      this.projectId = defaults.projectId;
    	      this.versionControl = defaults.versionControl;
    	      this.visibility = defaults.visibility;
    	      this.workItemTemplate = defaults.workItemTemplate;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder features(Map<String,String> features) {
            if (features == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "features");
            }
            this.features = features;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder processTemplateId(String processTemplateId) {
            if (processTemplateId == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "processTemplateId");
            }
            this.processTemplateId = processTemplateId;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {

            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder versionControl(String versionControl) {
            if (versionControl == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "versionControl");
            }
            this.versionControl = versionControl;
            return this;
        }
        @CustomType.Setter
        public Builder visibility(String visibility) {
            if (visibility == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "visibility");
            }
            this.visibility = visibility;
            return this;
        }
        @CustomType.Setter
        public Builder workItemTemplate(String workItemTemplate) {
            if (workItemTemplate == null) {
              throw new MissingRequiredPropertyException("GetProjectResult", "workItemTemplate");
            }
            this.workItemTemplate = workItemTemplate;
            return this;
        }
        public GetProjectResult build() {
            final var _resultValue = new GetProjectResult();
            _resultValue.description = description;
            _resultValue.features = features;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.processTemplateId = processTemplateId;
            _resultValue.projectId = projectId;
            _resultValue.versionControl = versionControl;
            _resultValue.visibility = visibility;
            _resultValue.workItemTemplate = workItemTemplate;
            return _resultValue;
        }
    }
}
