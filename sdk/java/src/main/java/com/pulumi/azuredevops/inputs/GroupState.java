// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuredevops.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupState extends com.pulumi.resources.ResourceArgs {

    public static final GroupState Empty = new GroupState();

    /**
     * The Description of the Project.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Description of the Project.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The identity (subject) descriptor of the Group.
     * 
     */
    @Import(name="descriptor")
    private @Nullable Output<String> descriptor;

    /**
     * @return The identity (subject) descriptor of the Group.
     * 
     */
    public Optional<Output<String>> descriptor() {
        return Optional.ofNullable(this.descriptor);
    }

    /**
     * The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * This represents the name of the container of origin for a graph member.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return This represents the name of the container of origin for a graph member.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * The ID of the Group.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return The ID of the Group.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
     * 
     */
    @Import(name="mail")
    private @Nullable Output<String> mail;

    /**
     * @return The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
     * 
     */
    public Optional<Output<String>> mail() {
        return Optional.ofNullable(this.mail);
    }

    /**
     * The member of the Group.
     * 
     * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return The member of the Group.
     * 
     * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     * 
     */
    @Import(name="origin")
    private @Nullable Output<String> origin;

    /**
     * @return The type of source provider for the origin identifier (ex:AD, AAD, MSA)
     * 
     */
    public Optional<Output<String>> origin() {
        return Optional.ofNullable(this.origin);
    }

    /**
     * The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
     * 
     */
    @Import(name="originId")
    private @Nullable Output<String> originId;

    /**
     * @return The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
     * 
     */
    public Optional<Output<String>> originId() {
        return Optional.ofNullable(this.originId);
    }

    /**
     * This is the PrincipalName of this graph member from the source provider.
     * 
     */
    @Import(name="principalName")
    private @Nullable Output<String> principalName;

    /**
     * @return This is the PrincipalName of this graph member from the source provider.
     * 
     */
    public Optional<Output<String>> principalName() {
        return Optional.ofNullable(this.principalName);
    }

    /**
     * The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * This field identifies the type of the graph subject (ex: Group, Scope, User).
     * 
     */
    @Import(name="subjectKind")
    private @Nullable Output<String> subjectKind;

    /**
     * @return This field identifies the type of the graph subject (ex: Group, Scope, User).
     * 
     */
    public Optional<Output<String>> subjectKind() {
        return Optional.ofNullable(this.subjectKind);
    }

    /**
     * This url is the full route to the source resource of this graph subject.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return This url is the full route to the source resource of this graph subject.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private GroupState() {}

    private GroupState(GroupState $) {
        this.description = $.description;
        this.descriptor = $.descriptor;
        this.displayName = $.displayName;
        this.domain = $.domain;
        this.groupId = $.groupId;
        this.mail = $.mail;
        this.members = $.members;
        this.origin = $.origin;
        this.originId = $.originId;
        this.principalName = $.principalName;
        this.scope = $.scope;
        this.subjectKind = $.subjectKind;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupState $;

        public Builder() {
            $ = new GroupState();
        }

        public Builder(GroupState defaults) {
            $ = new GroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The Description of the Project.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Description of the Project.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param descriptor The identity (subject) descriptor of the Group.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(@Nullable Output<String> descriptor) {
            $.descriptor = descriptor;
            return this;
        }

        /**
         * @param descriptor The identity (subject) descriptor of the Group.
         * 
         * @return builder
         * 
         */
        public Builder descriptor(String descriptor) {
            return descriptor(Output.of(descriptor));
        }

        /**
         * @param displayName The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param domain This represents the name of the container of origin for a graph member.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain This represents the name of the container of origin for a graph member.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param groupId The ID of the Group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the Group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param mail The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
         * 
         * @return builder
         * 
         */
        public Builder mail(@Nullable Output<String> mail) {
            $.mail = mail;
            return this;
        }

        /**
         * @param mail The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
         * 
         * @return builder
         * 
         */
        public Builder mail(String mail) {
            return mail(Output.of(mail));
        }

        /**
         * @param members The member of the Group.
         * 
         * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members The member of the Group.
         * 
         * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members The member of the Group.
         * 
         * &gt; **NOTE:** It&#39;s possible to define group members both within the `azuredevops.Group` resource via the members block and by using the `azuredevops.GroupMembership` resource. However it&#39;s not possible to use both methods to manage group members, since there&#39;ll be conflicts.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param origin The type of source provider for the origin identifier (ex:AD, AAD, MSA)
         * 
         * @return builder
         * 
         */
        public Builder origin(@Nullable Output<String> origin) {
            $.origin = origin;
            return this;
        }

        /**
         * @param origin The type of source provider for the origin identifier (ex:AD, AAD, MSA)
         * 
         * @return builder
         * 
         */
        public Builder origin(String origin) {
            return origin(Output.of(origin));
        }

        /**
         * @param originId The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
         * 
         * @return builder
         * 
         */
        public Builder originId(@Nullable Output<String> originId) {
            $.originId = originId;
            return this;
        }

        /**
         * @param originId The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
         * 
         * @return builder
         * 
         */
        public Builder originId(String originId) {
            return originId(Output.of(originId));
        }

        /**
         * @param principalName This is the PrincipalName of this graph member from the source provider.
         * 
         * @return builder
         * 
         */
        public Builder principalName(@Nullable Output<String> principalName) {
            $.principalName = principalName;
            return this;
        }

        /**
         * @param principalName This is the PrincipalName of this graph member from the source provider.
         * 
         * @return builder
         * 
         */
        public Builder principalName(String principalName) {
            return principalName(Output.of(principalName));
        }

        /**
         * @param scope The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        /**
         * @param subjectKind This field identifies the type of the graph subject (ex: Group, Scope, User).
         * 
         * @return builder
         * 
         */
        public Builder subjectKind(@Nullable Output<String> subjectKind) {
            $.subjectKind = subjectKind;
            return this;
        }

        /**
         * @param subjectKind This field identifies the type of the graph subject (ex: Group, Scope, User).
         * 
         * @return builder
         * 
         */
        public Builder subjectKind(String subjectKind) {
            return subjectKind(Output.of(subjectKind));
        }

        /**
         * @param url This url is the full route to the source resource of this graph subject.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url This url is the full route to the source resource of this graph subject.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public GroupState build() {
            return $;
        }
    }

}
