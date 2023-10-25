# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetTeamsResult',
    'AwaitableGetTeamsResult',
    'get_teams',
    'get_teams_output',
]

@pulumi.output_type
class GetTeamsResult:
    """
    A collection of values returned by getTeams.
    """
    def __init__(__self__, id=None, project_id=None, teams=None, top=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if teams and not isinstance(teams, list):
            raise TypeError("Expected argument 'teams' to be a list")
        pulumi.set(__self__, "teams", teams)
        if top and not isinstance(top, int):
            raise TypeError("Expected argument 'top' to be a int")
        pulumi.set(__self__, "top", top)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        Project identifier.
        - `id - Team identifier
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def teams(self) -> Sequence['outputs.GetTeamsTeamResult']:
        """
        A list of existing projects in your Azure DevOps Organization with details about every project which includes:
        """
        return pulumi.get(self, "teams")

    @property
    @pulumi.getter
    def top(self) -> Optional[int]:
        return pulumi.get(self, "top")


class AwaitableGetTeamsResult(GetTeamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTeamsResult(
            id=self.id,
            project_id=self.project_id,
            teams=self.teams,
            top=self.top)


def get_teams(project_id: Optional[str] = None,
              top: Optional[int] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTeamsResult:
    """
    Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_teams()
    pulumi.export("projectId", [__item.project_id for __item in example.teams])
    pulumi.export("name", [__item.name for __item in example.teams])
    pulumi.export("alladministrators", [__item.administrators for __item in example.teams])
    pulumi.export("administrators", [__item.members for __item in example.teams])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)

    ## PAT Permissions Required

    - **vso.project**:	Grants the ability to read projects and teams.


    :param str project_id: The Project ID. If no project ID all teams of the organization will be returned.
    :param int top: The maximum number of teams to return. Defaults to `100`.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['top'] = top
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getTeams:getTeams', __args__, opts=opts, typ=GetTeamsResult).value

    return AwaitableGetTeamsResult(
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        teams=pulumi.get(__ret__, 'teams'),
        top=pulumi.get(__ret__, 'top'))


@_utilities.lift_output_func(get_teams)
def get_teams_output(project_id: Optional[pulumi.Input[Optional[str]]] = None,
                     top: Optional[pulumi.Input[Optional[int]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTeamsResult]:
    """
    Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_teams()
    pulumi.export("projectId", [__item.project_id for __item in example.teams])
    pulumi.export("name", [__item.name for __item in example.teams])
    pulumi.export("alladministrators", [__item.administrators for __item in example.teams])
    pulumi.export("administrators", [__item.members for __item in example.teams])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)

    ## PAT Permissions Required

    - **vso.project**:	Grants the ability to read projects and teams.


    :param str project_id: The Project ID. If no project ID all teams of the organization will be returned.
    :param int top: The maximum number of teams to return. Defaults to `100`.
    """
    ...
