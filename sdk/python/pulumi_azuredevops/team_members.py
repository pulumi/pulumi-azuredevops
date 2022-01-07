# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TeamMembersArgs', 'TeamMembers']

@pulumi.input_type
class TeamMembersArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 project_id: pulumi.Input[str],
                 team_id: pulumi.Input[str],
                 mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TeamMembers resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of subject descriptors to define members of the team.
        :param pulumi.Input[str] project_id: The Project ID.
        :param pulumi.Input[str] team_id: The ID of the Team.
        :param pulumi.Input[str] mode: The mode how the resource manages team members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "team_id", team_id)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of subject descriptors to define members of the team.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Input[str]:
        """
        The ID of the Team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode how the resource manages team members.
        - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
        - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class _TeamMembersState:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TeamMembers resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of subject descriptors to define members of the team.
        :param pulumi.Input[str] mode: The mode how the resource manages team members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        :param pulumi.Input[str] project_id: The Project ID.
        :param pulumi.Input[str] team_id: The ID of the Team.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of subject descriptors to define members of the team.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode how the resource manages team members.
        - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
        - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


class TeamMembers(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages members of a team within a project in a Azure DevOps organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="My first project")
        builtin_project_readers = project.id.apply(lambda id: azuredevops.get_group(project_id=id,
            name="Readers"))
        team = azuredevops.Team("team", project_id=project.id)
        team_members = azuredevops.TeamMembers("teamMembers",
            project_id=team.project_id,
            team_id=team.id,
            mode="overwrite",
            members=[builtin_project_readers.descriptor])
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-5.1)

        ## PAT Permissions Required

        - **vso.project_write**:	Grants the ability to read and update projects and teams.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of subject descriptors to define members of the team.
        :param pulumi.Input[str] mode: The mode how the resource manages team members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        :param pulumi.Input[str] project_id: The Project ID.
        :param pulumi.Input[str] team_id: The ID of the Team.
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

        project = azuredevops.Project("project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="My first project")
        builtin_project_readers = project.id.apply(lambda id: azuredevops.get_group(project_id=id,
            name="Readers"))
        team = azuredevops.Team("team", project_id=project.id)
        team_members = azuredevops.TeamMembers("teamMembers",
            project_id=team.project_id,
            team_id=team.id,
            mode="overwrite",
            members=[builtin_project_readers.descriptor])
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-5.1)

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
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
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
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None) -> 'TeamMembers':
        """
        Get an existing TeamMembers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of subject descriptors to define members of the team.
        :param pulumi.Input[str] mode: The mode how the resource manages team members.
               - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
               - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        :param pulumi.Input[str] project_id: The Project ID.
        :param pulumi.Input[str] team_id: The ID of the Team.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamMembersState.__new__(_TeamMembersState)

        __props__.__dict__["members"] = members
        __props__.__dict__["mode"] = mode
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["team_id"] = team_id
        return TeamMembers(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        List of subject descriptors to define members of the team.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        The mode how the resource manages team members.
        - `mode == add`: the resource will ensure that all specified members will be part of the referenced team
        - `mode == overwrite`: the resource will replace all existing members with the members specified within the `members` block
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The ID of the Team.
        """
        return pulumi.get(self, "team_id")
