// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceEndpointArtifactoryAuthenticationBasic {
    /**
     * @return Artifactory Password.
     * 
     */
    private String password;
    /**
     * @return Artifactory Username.
     * 
     */
    private String username;

    private ServiceEndpointArtifactoryAuthenticationBasic() {}
    /**
     * @return Artifactory Password.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Artifactory Username.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointArtifactoryAuthenticationBasic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private String username;
        public Builder() {}
        public Builder(ServiceEndpointArtifactoryAuthenticationBasic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public ServiceEndpointArtifactoryAuthenticationBasic build() {
            final var o = new ServiceEndpointArtifactoryAuthenticationBasic();
            o.password = password;
            o.username = username;
            return o;
        }
    }
}
