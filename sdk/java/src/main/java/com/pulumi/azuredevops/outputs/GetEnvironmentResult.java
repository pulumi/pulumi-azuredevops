// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetEnvironmentResult {
    /**
     * @return A description for the Environment.
     * 
     */
    private String description;
    private Integer environmentId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the Environment.
     * 
     */
    private String name;
    private String projectId;

    private GetEnvironmentResult() {}
    /**
     * @return A description for the Environment.
     * 
     */
    public String description() {
        return this.description;
    }
    public Integer environmentId() {
        return this.environmentId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the Environment.
     * 
     */
    public String name() {
        return this.name;
    }
    public String projectId() {
        return this.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEnvironmentResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Integer environmentId;
        private String id;
        private String name;
        private String projectId;
        public Builder() {}
        public Builder(GetEnvironmentResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.environmentId = defaults.environmentId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder environmentId(Integer environmentId) {
            this.environmentId = Objects.requireNonNull(environmentId);
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
        public GetEnvironmentResult build() {
            final var o = new GetEnvironmentResult();
            o.description = description;
            o.environmentId = environmentId;
            o.id = id;
            o.name = name;
            o.projectId = projectId;
            return o;
        }
    }
}
