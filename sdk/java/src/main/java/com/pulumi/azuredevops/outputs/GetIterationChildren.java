// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIterationChildren {
    /**
     * @return Indicator if the child Iteration node has child nodes
     * 
     */
    private Boolean hasChildren;
    /**
     * @return The ID of the child Iteration node
     * 
     */
    private String id;
    /**
     * @return The name of the child Iteration node
     * 
     */
    private String name;
    /**
     * @return The path to the Iteration, _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Iteration will be returned
     * 
     */
    private String path;
    /**
     * @return The project ID.
     * 
     */
    private String projectId;

    private GetIterationChildren() {}
    /**
     * @return Indicator if the child Iteration node has child nodes
     * 
     */
    public Boolean hasChildren() {
        return this.hasChildren;
    }
    /**
     * @return The ID of the child Iteration node
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the child Iteration node
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The path to the Iteration, _Format_: URL relative; if omitted, or value `&#34;/&#34;` is used, the root Iteration will be returned
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return The project ID.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIterationChildren defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean hasChildren;
        private String id;
        private String name;
        private String path;
        private String projectId;
        public Builder() {}
        public Builder(GetIterationChildren defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hasChildren = defaults.hasChildren;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.path = defaults.path;
    	      this.projectId = defaults.projectId;
        }

        @CustomType.Setter
        public Builder hasChildren(Boolean hasChildren) {
            if (hasChildren == null) {
              throw new MissingRequiredPropertyException("GetIterationChildren", "hasChildren");
            }
            this.hasChildren = hasChildren;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIterationChildren", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIterationChildren", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            if (path == null) {
              throw new MissingRequiredPropertyException("GetIterationChildren", "path");
            }
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetIterationChildren", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        public GetIterationChildren build() {
            final var _resultValue = new GetIterationChildren();
            _resultValue.hasChildren = hasChildren;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.path = path;
            _resultValue.projectId = projectId;
            return _resultValue;
        }
    }
}
