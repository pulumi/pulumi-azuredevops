// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops;

import com.pulumi.azuredevops.inputs.FeedFeatureArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeedArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeedArgs Empty = new FeedArgs();

    /**
     * A `features` blocks as documented below.
     * 
     * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
     * 
     */
    @Import(name="features")
    private @Nullable Output<List<FeedFeatureArgs>> features;

    /**
     * @return A `features` blocks as documented below.
     * 
     * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
     * 
     */
    public Optional<Output<List<FeedFeatureArgs>>> features() {
        return Optional.ofNullable(this.features);
    }

    /**
     * The name of the Feed.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Feed.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private FeedArgs() {}

    private FeedArgs(FeedArgs $) {
        this.features = $.features;
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeedArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeedArgs $;

        public Builder() {
            $ = new FeedArgs();
        }

        public Builder(FeedArgs defaults) {
            $ = new FeedArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param features A `features` blocks as documented below.
         * 
         * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
         * 
         * @return builder
         * 
         */
        public Builder features(@Nullable Output<List<FeedFeatureArgs>> features) {
            $.features = features;
            return this;
        }

        /**
         * @param features A `features` blocks as documented below.
         * 
         * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
         * 
         * @return builder
         * 
         */
        public Builder features(List<FeedFeatureArgs> features) {
            return features(Output.of(features));
        }

        /**
         * @param features A `features` blocks as documented below.
         * 
         * &gt; **Note** *Because of ADO limitations feed name can be **reserved** for up to 15 minutes after permanent delete of the feed*
         * 
         * @return builder
         * 
         */
        public Builder features(FeedFeatureArgs... features) {
            return features(List.of(features));
        }

        /**
         * @param name The name of the Feed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Feed.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public FeedArgs build() {
            return $;
        }
    }

}
