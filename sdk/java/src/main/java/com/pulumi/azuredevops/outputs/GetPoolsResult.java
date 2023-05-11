// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.azuredevops.outputs.GetPoolsAgentPool;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetPoolsResult {
    /**
     * @return A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
     * 
     */
    private List<GetPoolsAgentPool> agentPools;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetPoolsResult() {}
    /**
     * @return A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
     * 
     */
    public List<GetPoolsAgentPool> agentPools() {
        return this.agentPools;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPoolsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetPoolsAgentPool> agentPools;
        private String id;
        public Builder() {}
        public Builder(GetPoolsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.agentPools = defaults.agentPools;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder agentPools(List<GetPoolsAgentPool> agentPools) {
            this.agentPools = Objects.requireNonNull(agentPools);
            return this;
        }
        public Builder agentPools(GetPoolsAgentPool... agentPools) {
            return agentPools(List.of(agentPools));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetPoolsResult build() {
            final var o = new GetPoolsResult();
            o.agentPools = agentPools;
            o.id = id;
            return o;
        }
    }
}