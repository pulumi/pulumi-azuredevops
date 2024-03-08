# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupMembershipArgs', 'GroupMembership']

@pulumi.input_type
class GroupMembershipArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupMembership resource.
        :param pulumi.Input[str] group: The descriptor of the group being managed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of user or group descriptors that will become members of the group.
               > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] mode: The mode how the resource manages group members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
               > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "members", members)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The descriptor of the group being managed.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of user or group descriptors that will become members of the group.
        > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode how the resource manages group members.
        - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
        - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class _GroupMembershipState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupMembership resources.
        :param pulumi.Input[str] group: The descriptor of the group being managed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of user or group descriptors that will become members of the group.
               > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] mode: The mode how the resource manages group members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
               > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The descriptor of the group being managed.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of user or group descriptors that will become members of the group.
        > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode how the resource manages group members.
        - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
        - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


warnings.warn("""azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership""", DeprecationWarning)


class GroupMembership(pulumi.CustomResource):
    warnings.warn("""azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages group membership within Azure DevOps.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_user = azuredevops.User("exampleUser", principal_name="foo@contoso.com")
        example_group = azuredevops.get_group_output(project_id=example_project.id,
            name="Build Administrators")
        example_group_membership = azuredevops.GroupMembership("exampleGroupMembership",
            group=example_group.descriptor,
            members=[example_user.descriptor])
        ```
        <!--End PulumiCodeChooser -->

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Memberships](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/memberships?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Deployment Groups**: Read & Manage

        ## Import

        Not supported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The descriptor of the group being managed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of user or group descriptors that will become members of the group.
               > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] mode: The mode how the resource manages group members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
               > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages group membership within Azure DevOps.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_user = azuredevops.User("exampleUser", principal_name="foo@contoso.com")
        example_group = azuredevops.get_group_output(project_id=example_project.id,
            name="Build Administrators")
        example_group_membership = azuredevops.GroupMembership("exampleGroupMembership",
            group=example_group.descriptor,
            members=[example_user.descriptor])
        ```
        <!--End PulumiCodeChooser -->

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Memberships](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/memberships?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Deployment Groups**: Read & Manage

        ## Import

        Not supported.

        :param str resource_name: The name of the resource.
        :param GroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""GroupMembership is deprecated: azuredevops.identities.GroupMembership has been deprecated in favor of azuredevops.GroupMembership""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupMembershipArgs.__new__(GroupMembershipArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["mode"] = mode
        super(GroupMembership, __self__).__init__(
            'azuredevops:Identities/groupMembership:GroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mode: Optional[pulumi.Input[str]] = None) -> 'GroupMembership':
        """
        Get an existing GroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The descriptor of the group being managed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of user or group descriptors that will become members of the group.
               > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        :param pulumi.Input[str] mode: The mode how the resource manages group members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
               > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupMembershipState.__new__(_GroupMembershipState)

        __props__.__dict__["group"] = group
        __props__.__dict__["members"] = members
        __props__.__dict__["mode"] = mode
        return GroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The descriptor of the group being managed.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of user or group descriptors that will become members of the group.
        > NOTE: It's possible to define group members both within the `GroupMembership resource` via the members block and by using the `Group` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        The mode how the resource manages group members.
        - `mode == add`: the resource will ensure that all specified members will be part of the referenced group
        - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        > NOTE: To clear all members from a group, specify an empty list of descriptors in the `members` attribute and set the `mode` member to `overwrite`.
        """
        return pulumi.get(self, "mode")

