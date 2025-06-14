# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = [
    'GetGitRepositoryFileResult',
    'AwaitableGetGitRepositoryFileResult',
    'get_git_repository_file',
    'get_git_repository_file_output',
]

@pulumi.output_type
class GetGitRepositoryFileResult:
    """
    A collection of values returned by getGitRepositoryFile.
    """
    def __init__(__self__, branch=None, content=None, file=None, id=None, last_commit_message=None, repository_id=None, tag=None):
        if branch and not isinstance(branch, str):
            raise TypeError("Expected argument 'branch' to be a str")
        pulumi.set(__self__, "branch", branch)
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        pulumi.set(__self__, "content", content)
        if file and not isinstance(file, str):
            raise TypeError("Expected argument 'file' to be a str")
        pulumi.set(__self__, "file", file)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_commit_message and not isinstance(last_commit_message, str):
            raise TypeError("Expected argument 'last_commit_message' to be a str")
        pulumi.set(__self__, "last_commit_message", last_commit_message)
        if repository_id and not isinstance(repository_id, str):
            raise TypeError("Expected argument 'repository_id' to be a str")
        pulumi.set(__self__, "repository_id", repository_id)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def branch(self) -> Optional[builtins.str]:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def content(self) -> builtins.str:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def file(self) -> builtins.str:
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastCommitMessage")
    def last_commit_message(self) -> builtins.str:
        """
        The commit message for the file.
        """
        return pulumi.get(self, "last_commit_message")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> builtins.str:
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter
    def tag(self) -> Optional[builtins.str]:
        return pulumi.get(self, "tag")


class AwaitableGetGitRepositoryFileResult(GetGitRepositoryFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGitRepositoryFileResult(
            branch=self.branch,
            content=self.content,
            file=self.file,
            id=self.id,
            last_commit_message=self.last_commit_message,
            repository_id=self.repository_id,
            tag=self.tag)


def get_git_repository_file(branch: Optional[builtins.str] = None,
                            file: Optional[builtins.str] = None,
                            repository_id: Optional[builtins.str] = None,
                            tag: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGitRepositoryFileResult:
    """
    Use this data source to get an existing Git Repository File.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # Load a specific Git repository by name
    example_get_git_repository = azuredevops.get_git_repository(project_id=example.id,
        name="Example Repository")
    example_get_git_repository_file = azuredevops.get_git_repository_file(repository_id=example_get_git_repository.id,
        branch="refs/heads/main",
        file="MyFile.txt")
    ```


    :param builtins.str branch: The git branch to use. Conflicts with `tag`; one or the other must be specified.
    :param builtins.str file: The path of the file to get.
    :param builtins.str repository_id: The ID of the Git repository.
    :param builtins.str tag: The tag to use.Conflicts with `branch`; one or the other must be specified.
    """
    __args__ = dict()
    __args__['branch'] = branch
    __args__['file'] = file
    __args__['repositoryId'] = repository_id
    __args__['tag'] = tag
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getGitRepositoryFile:getGitRepositoryFile', __args__, opts=opts, typ=GetGitRepositoryFileResult).value

    return AwaitableGetGitRepositoryFileResult(
        branch=pulumi.get(__ret__, 'branch'),
        content=pulumi.get(__ret__, 'content'),
        file=pulumi.get(__ret__, 'file'),
        id=pulumi.get(__ret__, 'id'),
        last_commit_message=pulumi.get(__ret__, 'last_commit_message'),
        repository_id=pulumi.get(__ret__, 'repository_id'),
        tag=pulumi.get(__ret__, 'tag'))
def get_git_repository_file_output(branch: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   file: Optional[pulumi.Input[builtins.str]] = None,
                                   repository_id: Optional[pulumi.Input[builtins.str]] = None,
                                   tag: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGitRepositoryFileResult]:
    """
    Use this data source to get an existing Git Repository File.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    # Load a specific Git repository by name
    example_get_git_repository = azuredevops.get_git_repository(project_id=example.id,
        name="Example Repository")
    example_get_git_repository_file = azuredevops.get_git_repository_file(repository_id=example_get_git_repository.id,
        branch="refs/heads/main",
        file="MyFile.txt")
    ```


    :param builtins.str branch: The git branch to use. Conflicts with `tag`; one or the other must be specified.
    :param builtins.str file: The path of the file to get.
    :param builtins.str repository_id: The ID of the Git repository.
    :param builtins.str tag: The tag to use.Conflicts with `branch`; one or the other must be specified.
    """
    __args__ = dict()
    __args__['branch'] = branch
    __args__['file'] = file
    __args__['repositoryId'] = repository_id
    __args__['tag'] = tag
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getGitRepositoryFile:getGitRepositoryFile', __args__, opts=opts, typ=GetGitRepositoryFileResult)
    return __ret__.apply(lambda __response__: GetGitRepositoryFileResult(
        branch=pulumi.get(__response__, 'branch'),
        content=pulumi.get(__response__, 'content'),
        file=pulumi.get(__response__, 'file'),
        id=pulumi.get(__response__, 'id'),
        last_commit_message=pulumi.get(__response__, 'last_commit_message'),
        repository_id=pulumi.get(__response__, 'repository_id'),
        tag=pulumi.get(__response__, 'tag')))
