# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GitRepositoryFileArgs', 'GitRepositoryFile']

@pulumi.input_type
class GitRepositoryFileArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 file: pulumi.Input[str],
                 repository_id: pulumi.Input[str],
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GitRepositoryFile resource.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[str] repository_id: The ID of the Git repository.
        :param pulumi.Input[str] branch: Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
               does not already exist.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files (defaults to `false`).
        """
        GitRepositoryFileArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            content=content,
            file=file,
            repository_id=repository_id,
            branch=branch,
            commit_message=commit_message,
            overwrite_on_create=overwrite_on_create,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             content: Optional[pulumi.Input[str]] = None,
             file: Optional[pulumi.Input[str]] = None,
             repository_id: Optional[pulumi.Input[str]] = None,
             branch: Optional[pulumi.Input[str]] = None,
             commit_message: Optional[pulumi.Input[str]] = None,
             overwrite_on_create: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if content is None:
            raise TypeError("Missing 'content' argument")
        if file is None:
            raise TypeError("Missing 'file' argument")
        if repository_id is None and 'repositoryId' in kwargs:
            repository_id = kwargs['repositoryId']
        if repository_id is None:
            raise TypeError("Missing 'repository_id' argument")
        if commit_message is None and 'commitMessage' in kwargs:
            commit_message = kwargs['commitMessage']
        if overwrite_on_create is None and 'overwriteOnCreate' in kwargs:
            overwrite_on_create = kwargs['overwriteOnCreate']

        _setter("content", content)
        _setter("file", file)
        _setter("repository_id", repository_id)
        if branch is not None:
            _setter("branch", branch)
        if commit_message is not None:
            _setter("commit_message", commit_message)
        if overwrite_on_create is not None:
            _setter("overwrite_on_create", overwrite_on_create)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def file(self) -> pulumi.Input[str]:
        """
        The path of the file to manage.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: pulumi.Input[str]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Input[str]:
        """
        The ID of the Git repository.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
        does not already exist.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> Optional[pulumi.Input[str]]:
        """
        Commit message when adding or updating the managed file.
        """
        return pulumi.get(self, "commit_message")

    @commit_message.setter
    def commit_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message", value)

    @property
    @pulumi.getter(name="overwriteOnCreate")
    def overwrite_on_create(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable overwriting existing files (defaults to `false`).
        """
        return pulumi.get(self, "overwrite_on_create")

    @overwrite_on_create.setter
    def overwrite_on_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overwrite_on_create", value)


@pulumi.input_type
class _GitRepositoryFileState:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GitRepositoryFile resources.
        :param pulumi.Input[str] branch: Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
               does not already exist.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files (defaults to `false`).
        :param pulumi.Input[str] repository_id: The ID of the Git repository.
        """
        _GitRepositoryFileState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            branch=branch,
            commit_message=commit_message,
            content=content,
            file=file,
            overwrite_on_create=overwrite_on_create,
            repository_id=repository_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             branch: Optional[pulumi.Input[str]] = None,
             commit_message: Optional[pulumi.Input[str]] = None,
             content: Optional[pulumi.Input[str]] = None,
             file: Optional[pulumi.Input[str]] = None,
             overwrite_on_create: Optional[pulumi.Input[bool]] = None,
             repository_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if commit_message is None and 'commitMessage' in kwargs:
            commit_message = kwargs['commitMessage']
        if overwrite_on_create is None and 'overwriteOnCreate' in kwargs:
            overwrite_on_create = kwargs['overwriteOnCreate']
        if repository_id is None and 'repositoryId' in kwargs:
            repository_id = kwargs['repositoryId']

        if branch is not None:
            _setter("branch", branch)
        if commit_message is not None:
            _setter("commit_message", commit_message)
        if content is not None:
            _setter("content", content)
        if file is not None:
            _setter("file", file)
        if overwrite_on_create is not None:
            _setter("overwrite_on_create", overwrite_on_create)
        if repository_id is not None:
            _setter("repository_id", repository_id)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
        does not already exist.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> Optional[pulumi.Input[str]]:
        """
        Commit message when adding or updating the managed file.
        """
        return pulumi.get(self, "commit_message")

    @commit_message.setter
    def commit_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the file to manage.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter(name="overwriteOnCreate")
    def overwrite_on_create(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable overwriting existing files (defaults to `false`).
        """
        return pulumi.get(self, "overwrite_on_create")

    @overwrite_on_create.setter
    def overwrite_on_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overwrite_on_create", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Git repository.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_id", value)


class GitRepositoryFile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage files within an Azure DevOps Git repository.

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)

        ## Import

        Repository files can be imported using a combination of the `repository ID` and `file`, e.g.

        ```sh
         $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore
        ```

         To import a file from a branch other than `master`, append `:` and the branch name, e.g.

        ```sh
         $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore:refs/heads/master
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
               does not already exist.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files (defaults to `false`).
        :param pulumi.Input[str] repository_id: The ID of the Git repository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GitRepositoryFileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage files within an Azure DevOps Git repository.

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Git API](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/?view=azure-devops-rest-7.0)

        ## Import

        Repository files can be imported using a combination of the `repository ID` and `file`, e.g.

        ```sh
         $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore
        ```

         To import a file from a branch other than `master`, append `:` and the branch name, e.g.

        ```sh
         $ pulumi import azuredevops:index/gitRepositoryFile:GitRepositoryFile example 00000000-0000-0000-0000-000000000000/.gitignore:refs/heads/master
        ```

        :param str resource_name: The name of the resource.
        :param GitRepositoryFileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GitRepositoryFileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            GitRepositoryFileArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GitRepositoryFileArgs.__new__(GitRepositoryFileArgs)

            __props__.__dict__["branch"] = branch
            __props__.__dict__["commit_message"] = commit_message
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            if file is None and not opts.urn:
                raise TypeError("Missing required property 'file'")
            __props__.__dict__["file"] = file
            __props__.__dict__["overwrite_on_create"] = overwrite_on_create
            if repository_id is None and not opts.urn:
                raise TypeError("Missing required property 'repository_id'")
            __props__.__dict__["repository_id"] = repository_id
        super(GitRepositoryFile, __self__).__init__(
            'azuredevops:index/gitRepositoryFile:GitRepositoryFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            commit_message: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            file: Optional[pulumi.Input[str]] = None,
            overwrite_on_create: Optional[pulumi.Input[bool]] = None,
            repository_id: Optional[pulumi.Input[str]] = None) -> 'GitRepositoryFile':
        """
        Get an existing GitRepositoryFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
               does not already exist.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files (defaults to `false`).
        :param pulumi.Input[str] repository_id: The ID of the Git repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GitRepositoryFileState.__new__(_GitRepositoryFileState)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["commit_message"] = commit_message
        __props__.__dict__["content"] = content
        __props__.__dict__["file"] = file
        __props__.__dict__["overwrite_on_create"] = overwrite_on_create
        __props__.__dict__["repository_id"] = repository_id
        return GitRepositoryFile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[Optional[str]]:
        """
        Git branch (defaults to `refs/heads/master`). The branch must already exist, it will not be created if it
        does not already exist.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> pulumi.Output[str]:
        """
        Commit message when adding or updating the managed file.
        """
        return pulumi.get(self, "commit_message")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[str]:
        """
        The path of the file to manage.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter(name="overwriteOnCreate")
    def overwrite_on_create(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable overwriting existing files (defaults to `false`).
        """
        return pulumi.get(self, "overwrite_on_create")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[str]:
        """
        The ID of the Git repository.
        """
        return pulumi.get(self, "repository_id")

