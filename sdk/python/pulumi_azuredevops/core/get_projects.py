# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
]

warnings.warn("""azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects""", DeprecationWarning)

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
        Project name.
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

    test = azuredevops.get_projects(name="contoso",
        state="wellFormed")
    pulumi.export("projectId", [__item.project_id for __item in test.projects])
    pulumi.export("name", [__item.name for __item in test.projects])
    pulumi.export("projectUrl", [__item.project_url for __item in test.projects])
    pulumi.export("state", [__item.state for __item in test.projects])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.1)


    :param str name: Name of the Project, if not specified all projects will be returned.
    :param str state: State of the Project, if not specified all projects will be returned. Valid values are `all`, `deleting`, `new`, `wellFormed`, `createPending`, `unchanged`,`deleted`.
    """
    pulumi.log.warn("""get_projects is deprecated: azuredevops.core.getProjects has been deprecated in favor of azuredevops.getProjects""")
    __args__ = dict()
    __args__['name'] = name
    __args__['state'] = state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:Core/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        id=__ret__.id,
        name=__ret__.name,
        projects=__ret__.projects,
        state=__ret__.state)
