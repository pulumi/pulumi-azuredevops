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
    'GetGitRepositoryResult',
    'AwaitableGetGitRepositoryResult',
    'get_git_repository',
    'get_git_repository_output',
]

@pulumi.output_type
class GetGitRepositoryResult:
    """
    A collection of values returned by getGitRepository.
    """
    def __init__(__self__, default_branch=None, disabled=None, id=None, is_fork=None, name=None, project_id=None, remote_url=None, size=None, ssh_url=None, url=None, web_url=None):
        if default_branch and not isinstance(default_branch, str):
            raise TypeError("Expected argument 'default_branch' to be a str")
        pulumi.set(__self__, "default_branch", default_branch)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_fork and not isinstance(is_fork, bool):
            raise TypeError("Expected argument 'is_fork' to be a bool")
        pulumi.set(__self__, "is_fork", is_fork)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if remote_url and not isinstance(remote_url, str):
            raise TypeError("Expected argument 'remote_url' to be a str")
        pulumi.set(__self__, "remote_url", remote_url)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if ssh_url and not isinstance(ssh_url, str):
            raise TypeError("Expected argument 'ssh_url' to be a str")
        pulumi.set(__self__, "ssh_url", ssh_url)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if web_url and not isinstance(web_url, str):
            raise TypeError("Expected argument 'web_url' to be a str")
        pulumi.set(__self__, "web_url", web_url)

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> str:
        """
        The ref of the default branch.
        """
        return pulumi.get(self, "default_branch")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Is the repository disabled?
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isFork")
    def is_fork(self) -> bool:
        return pulumi.get(self, "is_fork")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Git repository name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        Project identifier to which the Git repository belongs.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> str:
        """
        HTTPS Url to clone the Git repository
        """
        return pulumi.get(self, "remote_url")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        Compressed size (bytes) of the repository.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sshUrl")
    def ssh_url(self) -> str:
        """
        SSH Url to clone the Git repository
        """
        return pulumi.get(self, "ssh_url")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        Details REST API endpoint for the Git Repository.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> str:
        """
        Url of the Git repository web view
        """
        return pulumi.get(self, "web_url")


class AwaitableGetGitRepositoryResult(GetGitRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGitRepositoryResult(
            default_branch=self.default_branch,
            disabled=self.disabled,
            id=self.id,
            is_fork=self.is_fork,
            name=self.name,
            project_id=self.project_id,
            remote_url=self.remote_url,
            size=self.size,
            ssh_url=self.ssh_url,
            url=self.url,
            web_url=self.web_url)


def get_git_repository(name: Optional[str] = None,
                       project_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGitRepositoryResult:
    """
    Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
    To read information about **multiple** Git Repositories use the data source `get_repositories`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # Load a specific Git repository by name
    example_single_repo = azuredevops.get_git_repository(project_id=example.id,
        name="Example Repository")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)


    :param str name: Name of the Git repository to retrieve
    :param str project_id: ID of project to list Git repositories
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getGitRepository:getGitRepository', __args__, opts=opts, typ=GetGitRepositoryResult).value

    return AwaitableGetGitRepositoryResult(
        default_branch=pulumi.get(__ret__, 'default_branch'),
        disabled=pulumi.get(__ret__, 'disabled'),
        id=pulumi.get(__ret__, 'id'),
        is_fork=pulumi.get(__ret__, 'is_fork'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        remote_url=pulumi.get(__ret__, 'remote_url'),
        size=pulumi.get(__ret__, 'size'),
        ssh_url=pulumi.get(__ret__, 'ssh_url'),
        url=pulumi.get(__ret__, 'url'),
        web_url=pulumi.get(__ret__, 'web_url'))


@_utilities.lift_output_func(get_git_repository)
def get_git_repository_output(name: Optional[pulumi.Input[str]] = None,
                              project_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGitRepositoryResult]:
    """
    Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
    To read information about **multiple** Git Repositories use the data source `get_repositories`

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # Load a specific Git repository by name
    example_single_repo = azuredevops.get_git_repository(project_id=example.id,
        name="Example Repository")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)


    :param str name: Name of the Git repository to retrieve
    :param str project_id: ID of project to list Git repositories
    """
    ...
