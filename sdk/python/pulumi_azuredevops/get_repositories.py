# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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
from . import outputs

__all__ = [
    'GetRepositoriesResult',
    'AwaitableGetRepositoriesResult',
    'get_repositories',
    'get_repositories_output',
]

@pulumi.output_type
class GetRepositoriesResult:
    """
    A collection of values returned by getRepositories.
    """
    def __init__(__self__, id=None, include_hidden=None, name=None, project_id=None, repositories=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_hidden and not isinstance(include_hidden, bool):
            raise TypeError("Expected argument 'include_hidden' to be a bool")
        pulumi.set(__self__, "include_hidden", include_hidden)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if repositories and not isinstance(repositories, list):
            raise TypeError("Expected argument 'repositories' to be a list")
        pulumi.set(__self__, "repositories", repositories)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeHidden")
    def include_hidden(self) -> Optional[bool]:
        return pulumi.get(self, "include_hidden")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Git repository name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        Project identifier to which the Git repository belongs.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def repositories(self) -> Sequence['outputs.GetRepositoriesRepositoryResult']:
        """
        A list of existing projects in your Azure DevOps Organization with details about every project which includes:
        """
        return pulumi.get(self, "repositories")


class AwaitableGetRepositoriesResult(GetRepositoriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoriesResult(
            id=self.id,
            include_hidden=self.include_hidden,
            name=self.name,
            project_id=self.project_id,
            repositories=self.repositories)


def get_repositories(include_hidden: Optional[bool] = None,
                     name: Optional[str] = None,
                     project_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoriesResult:
    """
    Use this data source to access information about **multiple** existing Git Repositories within Azure DevOps.
    To read informations about a **single** Git Repository use the data source `Git`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # Load all Git repositories of a project, which are accessible for the current user
    example_all_repos = azuredevops.get_repositories(project_id=example.id,
        include_hidden=True)
    # Load a specific Git repository by name
    example_single_repo = azuredevops.get_repositories(project_id=example.id,
        name="Example Repository")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)


    :param bool include_hidden: DataSource without specifying any arguments will return all Git repositories of an organization.
    :param str name: Name of the Git repository to retrieve; requires `project_id` to be specified as well
    :param str project_id: ID of project to list Git repositories
    """
    __args__ = dict()
    __args__['includeHidden'] = include_hidden
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getRepositories:getRepositories', __args__, opts=opts, typ=GetRepositoriesResult).value

    return AwaitableGetRepositoriesResult(
        id=pulumi.get(__ret__, 'id'),
        include_hidden=pulumi.get(__ret__, 'include_hidden'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        repositories=pulumi.get(__ret__, 'repositories'))
def get_repositories_output(include_hidden: Optional[pulumi.Input[Optional[bool]]] = None,
                            name: Optional[pulumi.Input[Optional[str]]] = None,
                            project_id: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoriesResult]:
    """
    Use this data source to access information about **multiple** existing Git Repositories within Azure DevOps.
    To read informations about a **single** Git Repository use the data source `Git`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # Load all Git repositories of a project, which are accessible for the current user
    example_all_repos = azuredevops.get_repositories(project_id=example.id,
        include_hidden=True)
    # Load a specific Git repository by name
    example_single_repo = azuredevops.get_repositories(project_id=example.id,
        name="Example Repository")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)


    :param bool include_hidden: DataSource without specifying any arguments will return all Git repositories of an organization.
    :param str name: Name of the Git repository to retrieve; requires `project_id` to be specified as well
    :param str project_id: ID of project to list Git repositories
    """
    __args__ = dict()
    __args__['includeHidden'] = include_hidden
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getRepositories:getRepositories', __args__, opts=opts, typ=GetRepositoriesResult)
    return __ret__.apply(lambda __response__: GetRepositoriesResult(
        id=pulumi.get(__response__, 'id'),
        include_hidden=pulumi.get(__response__, 'include_hidden'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        repositories=pulumi.get(__response__, 'repositories')))
