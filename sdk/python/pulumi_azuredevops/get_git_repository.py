# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'GetGitRepositoryResult',
    'AwaitableGetGitRepositoryResult',
    'get_git_repository',
]

@pulumi.output_type
class GetGitRepositoryResult:
    """
    A collection of values returned by getGitRepository.
    """
    def __init__(__self__, default_branch=None, id=None, is_fork=None, name=None, project_id=None, remote_url=None, size=None, ssh_url=None, url=None, web_url=None):
        if default_branch and not isinstance(default_branch, str):
            raise TypeError("Expected argument 'default_branch' to be a str")
        pulumi.set(__self__, "default_branch", default_branch)
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
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
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
        return pulumi.get(self, "default_branch")

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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> str:
        return pulumi.get(self, "remote_url")

    @property
    @pulumi.getter
    def size(self) -> float:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sshUrl")
    def ssh_url(self) -> str:
        return pulumi.get(self, "ssh_url")

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> str:
        return pulumi.get(self, "web_url")


class AwaitableGetGitRepositoryResult(GetGitRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGitRepositoryResult(
            default_branch=self.default_branch,
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
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getGitRepository:getGitRepository', __args__, opts=opts, typ=GetGitRepositoryResult).value

    return AwaitableGetGitRepositoryResult(
        default_branch=__ret__.default_branch,
        id=__ret__.id,
        is_fork=__ret__.is_fork,
        name=__ret__.name,
        project_id=__ret__.project_id,
        remote_url=__ret__.remote_url,
        size=__ret__.size,
        ssh_url=__ret__.ssh_url,
        url=__ret__.url,
        web_url=__ret__.web_url)
