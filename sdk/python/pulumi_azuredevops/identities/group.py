# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mail: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[str] description: The Description of the Project.
        :param pulumi.Input[str] display_name: The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        :param pulumi.Input[str] mail: The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] origin_id: The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        :param pulumi.Input[str] scope: The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if mail is not None:
            pulumi.set(__self__, "mail", mail)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if origin_id is not None:
            pulumi.set(__self__, "origin_id", origin_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the Project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def mail(self) -> Optional[pulumi.Input[str]]:
        """
        The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        """
        return pulumi.get(self, "mail")

    @mail.setter
    def mail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mail", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        """
        return pulumi.get(self, "origin_id")

    @origin_id.setter
    def origin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 descriptor: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 mail: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 principal_name: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 subject_kind: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[str] description: The Description of the Project.
        :param pulumi.Input[str] descriptor: The identity (subject) descriptor of the Group.
        :param pulumi.Input[str] display_name: The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        :param pulumi.Input[str] domain: This represents the name of the container of origin for a graph member.
        :param pulumi.Input[str] mail: The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] origin: The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        :param pulumi.Input[str] origin_id: The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        :param pulumi.Input[str] principal_name: This is the PrincipalName of this graph member from the source provider.
        :param pulumi.Input[str] scope: The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        :param pulumi.Input[str] subject_kind: This field identifies the type of the graph subject (ex: Group, Scope, User).
        :param pulumi.Input[str] url: This url is the full route to the source resource of this graph subject.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if descriptor is not None:
            pulumi.set(__self__, "descriptor", descriptor)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if mail is not None:
            pulumi.set(__self__, "mail", mail)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if origin is not None:
            pulumi.set(__self__, "origin", origin)
        if origin_id is not None:
            pulumi.set(__self__, "origin_id", origin_id)
        if principal_name is not None:
            pulumi.set(__self__, "principal_name", principal_name)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if subject_kind is not None:
            pulumi.set(__self__, "subject_kind", subject_kind)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the Project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def descriptor(self) -> Optional[pulumi.Input[str]]:
        """
        The identity (subject) descriptor of the Group.
        """
        return pulumi.get(self, "descriptor")

    @descriptor.setter
    def descriptor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "descriptor", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        This represents the name of the container of origin for a graph member.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def mail(self) -> Optional[pulumi.Input[str]]:
        """
        The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        """
        return pulumi.get(self, "mail")

    @mail.setter
    def mail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mail", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def origin(self) -> Optional[pulumi.Input[str]]:
        """
        The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        """
        return pulumi.get(self, "origin")

    @origin.setter
    def origin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin", value)

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        """
        return pulumi.get(self, "origin_id")

    @origin_id.setter
    def origin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_id", value)

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the PrincipalName of this graph member from the source provider.
        """
        return pulumi.get(self, "principal_name")

    @principal_name.setter
    def principal_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_name", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="subjectKind")
    def subject_kind(self) -> Optional[pulumi.Input[str]]:
        """
        This field identifies the type of the graph subject (ex: Group, Scope, User).
        """
        return pulumi.get(self, "subject_kind")

    @subject_kind.setter
    def subject_kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject_kind", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        This url is the full route to the source resource of this graph subject.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


warnings.warn("""azuredevops.identities.Group has been deprecated in favor of azuredevops.Group""", DeprecationWarning)


class Group(pulumi.CustomResource):
    warnings.warn("""azuredevops.identities.Group has been deprecated in favor of azuredevops.Group""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mail: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a group within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project")
        tf_project_readers = azuredevops.get_group_output(project_id=project.id,
            name="Readers")
        tf_project_contributors = azuredevops.get_group_output(project_id=project.id,
            name="Contributors")
        group = azuredevops.Group("group",
            scope=project.id,
            display_name="Test group",
            description="Test description",
            members=[
                tf_project_readers.descriptor,
                tf_project_contributors.descriptor,
            ])
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-5.1)

        ## PAT Permissions Required

        - **Project & Team**: Read, Write, & Manage

        ## Import

        Azure DevOps groups can be imported using the group identity descriptor, e.g.

        ```sh
         $ pulumi import azuredevops:Identities/group:Group id aadgp.Uy0xLTktMTU1MTM3NDI0NS0xMjA0NDAwOTY5LTI0MDI5ODY0MTMtMjE3OTQwODYxNi0zLTIxNjc2NjQyNTMtMzI1Nzg0NDI4OS0yMjU4MjcwOTc0LTI2MDYxODY2NDU
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The Description of the Project.
        :param pulumi.Input[str] display_name: The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        :param pulumi.Input[str] mail: The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] origin_id: The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        :param pulumi.Input[str] scope: The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a group within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project")
        tf_project_readers = azuredevops.get_group_output(project_id=project.id,
            name="Readers")
        tf_project_contributors = azuredevops.get_group_output(project_id=project.id,
            name="Contributors")
        group = azuredevops.Group("group",
            scope=project.id,
            display_name="Test group",
            description="Test description",
            members=[
                tf_project_readers.descriptor,
                tf_project_contributors.descriptor,
            ])
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups?view=azure-devops-rest-5.1)

        ## PAT Permissions Required

        - **Project & Team**: Read, Write, & Manage

        ## Import

        Azure DevOps groups can be imported using the group identity descriptor, e.g.

        ```sh
         $ pulumi import azuredevops:Identities/group:Group id aadgp.Uy0xLTktMTU1MTM3NDI0NS0xMjA0NDAwOTY5LTI0MDI5ODY0MTMtMjE3OTQwODYxNi0zLTIxNjc2NjQyNTMtMzI1Nzg0NDI4OS0yMjU4MjcwOTc0LTI2MDYxODY2NDU
        ```

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mail: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Group is deprecated: azuredevops.identities.Group has been deprecated in favor of azuredevops.Group""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["mail"] = mail
            __props__.__dict__["members"] = members
            __props__.__dict__["origin_id"] = origin_id
            __props__.__dict__["scope"] = scope
            __props__.__dict__["descriptor"] = None
            __props__.__dict__["domain"] = None
            __props__.__dict__["origin"] = None
            __props__.__dict__["principal_name"] = None
            __props__.__dict__["subject_kind"] = None
            __props__.__dict__["url"] = None
        super(Group, __self__).__init__(
            'azuredevops:Identities/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            descriptor: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            mail: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            origin: Optional[pulumi.Input[str]] = None,
            origin_id: Optional[pulumi.Input[str]] = None,
            principal_name: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            subject_kind: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The Description of the Project.
        :param pulumi.Input[str] descriptor: The identity (subject) descriptor of the Group.
        :param pulumi.Input[str] display_name: The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        :param pulumi.Input[str] domain: This represents the name of the container of origin for a graph member.
        :param pulumi.Input[str] mail: The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] origin: The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        :param pulumi.Input[str] origin_id: The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        :param pulumi.Input[str] principal_name: This is the PrincipalName of this graph member from the source provider.
        :param pulumi.Input[str] scope: The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        :param pulumi.Input[str] subject_kind: This field identifies the type of the graph subject (ex: Group, Scope, User).
        :param pulumi.Input[str] url: This url is the full route to the source resource of this graph subject.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["descriptor"] = descriptor
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["domain"] = domain
        __props__.__dict__["mail"] = mail
        __props__.__dict__["members"] = members
        __props__.__dict__["origin"] = origin
        __props__.__dict__["origin_id"] = origin_id
        __props__.__dict__["principal_name"] = principal_name
        __props__.__dict__["scope"] = scope
        __props__.__dict__["subject_kind"] = subject_kind
        __props__.__dict__["url"] = url
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Description of the Project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def descriptor(self) -> pulumi.Output[str]:
        """
        The identity (subject) descriptor of the Group.
        """
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The name of a new Azure DevOps group that is not backed by an external provider. The `origin_id` and `mail` arguments cannot be used simultaneously with `display_name`.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        This represents the name of the container of origin for a graph member.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def mail(self) -> pulumi.Output[str]:
        """
        The mail address as a reference to an existing group from an external AD or AAD backed provider. The `scope`, `origin_id` and `display_name` arguments cannot be used simultaneously with `mail`.
        """
        return pulumi.get(self, "mail")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        > NOTE: It's possible to define group members both within the `Group` resource via the members block and by using the `GroupMembership` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def origin(self) -> pulumi.Output[str]:
        """
        The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> pulumi.Output[str]:
        """
        The OriginID as a reference to a group from an external AD or AAD backed provider. The `scope`, `mail` and `display_name` arguments cannot be used simultaneously with `origin_id`.
        """
        return pulumi.get(self, "origin_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> pulumi.Output[str]:
        """
        This is the PrincipalName of this graph member from the source provider.
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="subjectKind")
    def subject_kind(self) -> pulumi.Output[str]:
        """
        This field identifies the type of the graph subject (ex: Group, Scope, User).
        """
        return pulumi.get(self, "subject_kind")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        This url is the full route to the source resource of this graph subject.
        """
        return pulumi.get(self, "url")

