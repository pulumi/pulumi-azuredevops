# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetTeamResult',
    'AwaitableGetTeamResult',
    'get_team',
    'get_team_output',
]

@pulumi.output_type
class GetTeamResult:
    """
    A collection of values returned by getTeam.
    """
    def __init__(__self__, administrators=None, description=None, descriptor=None, id=None, members=None, name=None, project_id=None, top=None):
        if administrators and not isinstance(administrators, list):
            raise TypeError("Expected argument 'administrators' to be a list")
        pulumi.set(__self__, "administrators", administrators)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if descriptor and not isinstance(descriptor, str):
            raise TypeError("Expected argument 'descriptor' to be a str")
        pulumi.set(__self__, "descriptor", descriptor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if top and not isinstance(top, int):
            raise TypeError("Expected argument 'top' to be a int")
        pulumi.set(__self__, "top", top)

    @property
    @pulumi.getter
    def administrators(self) -> Sequence[str]:
        """
        List of subject descriptors for `administrators` of the team.
        """
        return pulumi.get(self, "administrators")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Team description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def descriptor(self) -> str:
        """
        The descriptor of the Team.
        """
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        List of subject descriptors for `members` of the team.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def top(self) -> Optional[int]:
        return pulumi.get(self, "top")


class AwaitableGetTeamResult(GetTeamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTeamResult(
            administrators=self.administrators,
            description=self.description,
            descriptor=self.descriptor,
            id=self.id,
            members=self.members,
            name=self.name,
            project_id=self.project_id,
            top=self.top)


def get_team(name: Optional[str] = None,
             project_id: Optional[str] = None,
             top: Optional[int] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTeamResult:
    """
    Use this data source to access information about an existing Team in a Project within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.Project("example",
        name="Example Project",
        work_item_template="Agile",
        version_control="Git",
        visibility="private",
        description="Managed by Terraform")
    example = azuredevops.get_team_output(project_id=example_project.id,
        name="Example Project Team")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)

    ## PAT Permissions Required

    - **vso.project**:	Grants the ability to read projects and teams.


    :param str name: The name of the Team.
    :param str project_id: The Project ID.
    :param int top: The maximum number of teams to return. Defaults to `100`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['top'] = top
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getTeam:getTeam', __args__, opts=opts, typ=GetTeamResult).value

    return AwaitableGetTeamResult(
        administrators=pulumi.get(__ret__, 'administrators'),
        description=pulumi.get(__ret__, 'description'),
        descriptor=pulumi.get(__ret__, 'descriptor'),
        id=pulumi.get(__ret__, 'id'),
        members=pulumi.get(__ret__, 'members'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        top=pulumi.get(__ret__, 'top'))


@_utilities.lift_output_func(get_team)
def get_team_output(name: Optional[pulumi.Input[str]] = None,
                    project_id: Optional[pulumi.Input[str]] = None,
                    top: Optional[pulumi.Input[Optional[int]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTeamResult]:
    """
    Use this data source to access information about an existing Team in a Project within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.Project("example",
        name="Example Project",
        work_item_template="Agile",
        version_control="Git",
        visibility="private",
        description="Managed by Terraform")
    example = azuredevops.get_team_output(project_id=example_project.id,
        name="Example Project Team")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)

    ## PAT Permissions Required

    - **vso.project**:	Grants the ability to read projects and teams.


    :param str name: The name of the Team.
    :param str project_id: The Project ID.
    :param int top: The maximum number of teams to return. Defaults to `100`.
    """
    ...
