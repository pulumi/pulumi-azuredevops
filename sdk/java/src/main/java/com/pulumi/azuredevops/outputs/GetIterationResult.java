// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetIterationChildren;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIterationResult {
    /**
     * @return A list of `children` blocks as defined below, empty if `has_children == false`
     * 
     */
    private List<GetIterationChildren> childrens;
    private @Nullable Boolean fetchChildren;
    /**
     * @return Indicator if the child Iteration node has child nodes
     * 
     */
    private Boolean hasChildren;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the child Iteration node
     * 
     */
    private String name;
    /**
     * @return The complete path (in relative URL format) of the child Iteration
     * 
     */
    private String path;
    /**
     * @return The project ID of the child Iteration node
     * 
     */
    private String projectId;

    private GetIterationResult() {}
    /**
     * @return A list of `children` blocks as defined below, empty if `has_children == false`
     * 
     */
    public List<GetIterationChildren> childrens() {
        return this.childrens;
    }
    public Optional<Boolean> fetchChildren() {
        return Optional.ofNullable(this.fetchChildren);
    }
    /**
     * @return Indicator if the child Iteration node has child nodes
     * 
     */
    public Boolean hasChildren() {
        return this.hasChildren;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
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
     * @return The complete path (in relative URL format) of the child Iteration
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return The project ID of the child Iteration node
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIterationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetIterationChildren> childrens;
        private @Nullable Boolean fetchChildren;
        private Boolean hasChildren;
        private String id;
        private String name;
        private String path;
        private String projectId;
        public Builder() {}
        public Builder(GetIterationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.childrens = defaults.childrens;
    	      this.fetchChildren = defaults.fetchChildren;
    	      this.hasChildren = defaults.hasChildren;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.path = defaults.path;
    	      this.projectId = defaults.projectId;
        }

        @CustomType.Setter
        public Builder childrens(List<GetIterationChildren> childrens) {
            this.childrens = Objects.requireNonNull(childrens);
            return this;
        }
        public Builder childrens(GetIterationChildren... childrens) {
            return childrens(List.of(childrens));
        }
        @CustomType.Setter
        public Builder fetchChildren(@Nullable Boolean fetchChildren) {
            this.fetchChildren = fetchChildren;
            return this;
        }
        @CustomType.Setter
        public Builder hasChildren(Boolean hasChildren) {
            this.hasChildren = Objects.requireNonNull(hasChildren);
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
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        public GetIterationResult build() {
            final var o = new GetIterationResult();
            o.childrens = childrens;
            o.fetchChildren = fetchChildren;
            o.hasChildren = hasChildren;
            o.id = id;
            o.name = name;
            o.path = path;
            o.projectId = projectId;
            return o;
        }
    }
}
