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
public final class GetPoolsAgentPool {
    /**
     * @return Specifies whether or not a queue should be automatically provisioned for each project collection.
     * 
     */
    private Boolean autoProvision;
    /**
     * @return Specifies whether or not agents within the pool should be automatically updated.
     * 
     */
    private Boolean autoUpdate;
    /**
     * @return The ID of the agent pool.
     * 
     */
    private Integer id;
    /**
     * @return The name of the agent pool.
     * 
     */
    private String name;
    /**
     * @return Specifies whether the agent pool type is Automation or Deployment.
     * 
     */
    private String poolType;

    private GetPoolsAgentPool() {}
    /**
     * @return Specifies whether or not a queue should be automatically provisioned for each project collection.
     * 
     */
    public Boolean autoProvision() {
        return this.autoProvision;
    }
    /**
     * @return Specifies whether or not agents within the pool should be automatically updated.
     * 
     */
    public Boolean autoUpdate() {
        return this.autoUpdate;
    }
    /**
     * @return The ID of the agent pool.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The name of the agent pool.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Specifies whether the agent pool type is Automation or Deployment.
     * 
     */
    public String poolType() {
        return this.poolType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPoolsAgentPool defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean autoProvision;
        private Boolean autoUpdate;
        private Integer id;
        private String name;
        private String poolType;
        public Builder() {}
        public Builder(GetPoolsAgentPool defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoProvision = defaults.autoProvision;
    	      this.autoUpdate = defaults.autoUpdate;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.poolType = defaults.poolType;
        }

        @CustomType.Setter
        public Builder autoProvision(Boolean autoProvision) {
            if (autoProvision == null) {
              throw new MissingRequiredPropertyException("GetPoolsAgentPool", "autoProvision");
            }
            this.autoProvision = autoProvision;
            return this;
        }
        @CustomType.Setter
        public Builder autoUpdate(Boolean autoUpdate) {
            if (autoUpdate == null) {
              throw new MissingRequiredPropertyException("GetPoolsAgentPool", "autoUpdate");
            }
            this.autoUpdate = autoUpdate;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPoolsAgentPool", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetPoolsAgentPool", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder poolType(String poolType) {
            if (poolType == null) {
              throw new MissingRequiredPropertyException("GetPoolsAgentPool", "poolType");
            }
            this.poolType = poolType;
            return this;
        }
        public GetPoolsAgentPool build() {
            final var _resultValue = new GetPoolsAgentPool();
            _resultValue.autoProvision = autoProvision;
            _resultValue.autoUpdate = autoUpdate;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.poolType = poolType;
            return _resultValue;
        }
    }
}
