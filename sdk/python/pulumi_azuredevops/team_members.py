# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['TeamMembersArgs', 'TeamMembers']

@pulumi.input_type
class TeamMembersArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]],
                 project_id: pulumi.Input[_builtins.str],
                 team_id: pulumi.Input[_builtins.str],
                 mode: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a TeamMembers resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] members: List of subject descriptors to define members of the team.
               
               > **NOTE:** It's possible to define team members both within the
               `Team` resource via the `members` block and by using the
               `TeamMembers` resource. However, it's not possible to use
               both methods to manage team members, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] project_id: The Project ID.
        :param pulumi.Input[_builtins.str] team_id: The ID of the Team.
        :param pulumi.Input[_builtins.str] mode: The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
               
               > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
               <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "team_id", team_id)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @_builtins.property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]:
        """
        List of subject descriptors to define members of the team.

        > **NOTE:** It's possible to define team members both within the
        `Team` resource via the `members` block and by using the
        `TeamMembers` resource. However, it's not possible to use
        both methods to manage team members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]):
        pulumi.set(self, "members", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[_builtins.str]:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the Team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "team_id", value)

    @_builtins.property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.

        > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
        <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class _TeamMembersState:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 mode: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 team_id: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering TeamMembers resources.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] members: List of subject descriptors to define members of the team.
               
               > **NOTE:** It's possible to define team members both within the
               `Team` resource via the `members` block and by using the
               `TeamMembers` resource. However, it's not possible to use
               both methods to manage team members, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] mode: The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
               
               > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
               <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        :param pulumi.Input[_builtins.str] project_id: The Project ID.
        :param pulumi.Input[_builtins.str] team_id: The ID of the Team.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @_builtins.property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        List of subject descriptors to define members of the team.

        > **NOTE:** It's possible to define team members both within the
        `Team` resource via the `members` block and by using the
        `TeamMembers` resource. However, it's not possible to use
        both methods to manage team members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "members", value)

    @_builtins.property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.

        > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
        <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mode", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the Team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "team_id", value)


@pulumi.type_token("azuredevops:index/teamMembers:TeamMembers")
class TeamMembers(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 mode: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 team_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages members of a team within a project in a Azure DevOps organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Pulumi")
        example_project_readers = azuredevops.get_group_output(project_id=example.id,
            name="Readers")
        example_team = azuredevops.Team("example",
            project_id=example.id,
            name=example.name.apply(lambda name: f"{name} Team 2"))
        example_team_members = azuredevops.TeamMembers("example-team-members",
            project_id=example_team.project_id,
            team_id=example_team.id,
            mode="overwrite",
            members=[example_project_readers.descriptor])
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **vso.project_write**:	Grants the ability to read and update projects and teams.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] members: List of subject descriptors to define members of the team.
               
               > **NOTE:** It's possible to define team members both within the
               `Team` resource via the `members` block and by using the
               `TeamMembers` resource. However, it's not possible to use
               both methods to manage team members, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] mode: The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
               
               > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
               <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        :param pulumi.Input[_builtins.str] project_id: The Project ID.
        :param pulumi.Input[_builtins.str] team_id: The ID of the Team.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeamMembersArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages members of a team within a project in a Azure DevOps organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Pulumi")
        example_project_readers = azuredevops.get_group_output(project_id=example.id,
            name="Readers")
        example_team = azuredevops.Team("example",
            project_id=example.id,
            name=example.name.apply(lambda name: f"{name} Team 2"))
        example_team_members = azuredevops.TeamMembers("example-team-members",
            project_id=example_team.project_id,
            team_id=example_team.id,
            mode="overwrite",
            members=[example_project_readers.descriptor])
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **vso.project_write**:	Grants the ability to read and update projects and teams.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param TeamMembersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamMembersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 mode: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 team_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamMembersArgs.__new__(TeamMembersArgs)

            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["mode"] = mode
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if team_id is None and not opts.urn:
                raise TypeError("Missing required property 'team_id'")
            __props__.__dict__["team_id"] = team_id
        super(TeamMembers, __self__).__init__(
            'azuredevops:index/teamMembers:TeamMembers',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            mode: Optional[pulumi.Input[_builtins.str]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None,
            team_id: Optional[pulumi.Input[_builtins.str]] = None) -> 'TeamMembers':
        """
        Get an existing TeamMembers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] members: List of subject descriptors to define members of the team.
               
               > **NOTE:** It's possible to define team members both within the
               `Team` resource via the `members` block and by using the
               `TeamMembers` resource. However, it's not possible to use
               both methods to manage team members, since there'll be conflicts.
        :param pulumi.Input[_builtins.str] mode: The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.
               
               > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
               <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        :param pulumi.Input[_builtins.str] project_id: The Project ID.
        :param pulumi.Input[_builtins.str] team_id: The ID of the Team.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamMembersState.__new__(_TeamMembersState)

        __props__.__dict__["members"] = members
        __props__.__dict__["mode"] = mode
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["team_id"] = team_id
        return TeamMembers(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        List of subject descriptors to define members of the team.

        > **NOTE:** It's possible to define team members both within the
        `Team` resource via the `members` block and by using the
        `TeamMembers` resource. However, it's not possible to use
        both methods to manage team members, since there'll be conflicts.
        """
        return pulumi.get(self, "members")

    @_builtins.property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The mode how the resource manages team members. Possible values: `add`, `overwrite`. Defaults to `add`.

        > **NOTE:** 1. `mode = add`: the resource will ensure that all specified members will be part of the referenced team
        <br>2. `mode = overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        return pulumi.get(self, "mode")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[_builtins.str]:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the Team.
        """
        return pulumi.get(self, "team_id")

