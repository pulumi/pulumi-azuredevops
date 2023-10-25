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
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
    'get_projects_output',
]

@pulumi.output_type
class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, id=None, name=None, projects=None, state=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the Project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def projects(self) -> Sequence['outputs.GetProjectsProjectResult']:
        """
        A list of existing projects in your Azure DevOps Organization with details about every project which includes:
        """
        return pulumi.get(self, "projects")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Project state.
        """
        return pulumi.get(self, "state")


class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            id=self.id,
            name=self.name,
            projects=self.projects,
            state=self.state)


def get_projects(name: Optional[str] = None,
                 state: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectsResult:
    """
    Use this data source to access information about existing Projects within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_projects(name="Example Project",
        state="wellFormed")
    pulumi.export("projectId", [__item.project_id for __item in example.projects])
    pulumi.export("name", [__item.name for __item in example.projects])
    pulumi.export("projectUrl", [__item.project_url for __item in example.projects])
    pulumi.export("state", [__item.state for __item in example.projects])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)


    :param str name: Name of the Project, if not specified all projects will be returned.
    :param str state: State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
           
           DataSource without specifying any arguments will return all projects.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['state'] = state
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        projects=pulumi.get(__ret__, 'projects'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_projects)
def get_projects_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        state: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectsResult]:
    """
    Use this data source to access information about existing Projects within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_projects(name="Example Project",
        state="wellFormed")
    pulumi.export("projectId", [__item.project_id for __item in example.projects])
    pulumi.export("name", [__item.name for __item in example.projects])
    pulumi.export("projectUrl", [__item.project_url for __item in example.projects])
    pulumi.export("state", [__item.state for __item in example.projects])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-7.0)


    :param str name: Name of the Project, if not specified all projects will be returned.
    :param str state: State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
           
           DataSource without specifying any arguments will return all projects.
    """
    ...
