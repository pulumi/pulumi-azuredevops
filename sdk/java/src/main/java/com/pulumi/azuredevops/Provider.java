// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.ProviderArgs;
import com.pulumi.azuredevops.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the azuredevops package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:azuredevops")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * The url of the Azure DevOps instance which should be used.
     * 
     */
    @Export(name="orgServiceUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> orgServiceUrl;

    /**
     * @return The url of the Azure DevOps instance which should be used.
     * 
     */
    public Output<Optional<String>> orgServiceUrl() {
        return Codegen.optional(this.orgServiceUrl);
    }
    /**
     * The personal access token which should be used.
     * 
     */
    @Export(name="personalAccessToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> personalAccessToken;

    /**
     * @return The personal access token which should be used.
     * 
     */
    public Output<Optional<String>> personalAccessToken() {
        return Codegen.optional(this.personalAccessToken);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuredevops", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "personalAccessToken"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}