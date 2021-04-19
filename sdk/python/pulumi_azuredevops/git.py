# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GitArgs', 'Git']

@pulumi.input_type
class GitArgs:
    def __init__(__self__, *,
                 initialization: pulumi.Input['GitInitializationArgs'],
                 project_id: pulumi.Input[str],
                 default_branch: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_repository_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Git resource.
        :param pulumi.Input['GitInitializationArgs'] initialization: An `initialization` block as documented below.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] default_branch: The ref of the default branch. Will be used as the branch name for initialized repositories.
        :param pulumi.Input[str] name: The name of the git repository.
        :param pulumi.Input[str] parent_repository_id: The ID of a Git project from which a fork is to be created.
        """
        pulumi.set(__self__, "initialization", initialization)
        pulumi.set(__self__, "project_id", project_id)
        if default_branch is not None:
            pulumi.set(__self__, "default_branch", default_branch)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_repository_id is not None:
            pulumi.set(__self__, "parent_repository_id", parent_repository_id)

    @property
    @pulumi.getter
    def initialization(self) -> pulumi.Input['GitInitializationArgs']:
        """
        An `initialization` block as documented below.
        """
        return pulumi.get(self, "initialization")

    @initialization.setter
    def initialization(self, value: pulumi.Input['GitInitializationArgs']):
        pulumi.set(self, "initialization", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> Optional[pulumi.Input[str]]:
        """
        The ref of the default branch. Will be used as the branch name for initialized repositories.
        """
        return pulumi.get(self, "default_branch")

    @default_branch.setter
    def default_branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_branch", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the git repository.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentRepositoryId")
    def parent_repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a Git project from which a fork is to be created.
        """
        return pulumi.get(self, "parent_repository_id")

    @parent_repository_id.setter
    def parent_repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_repository_id", value)


@pulumi.input_type
class _GitState:
    def __init__(__self__, *,
                 default_branch: Optional[pulumi.Input[str]] = None,
                 initialization: Optional[pulumi.Input['GitInitializationArgs']] = None,
                 is_fork: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_repository_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 remote_url: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 ssh_url: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 web_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Git resources.
        :param pulumi.Input[str] default_branch: The ref of the default branch. Will be used as the branch name for initialized repositories.
        :param pulumi.Input['GitInitializationArgs'] initialization: An `initialization` block as documented below.
        :param pulumi.Input[bool] is_fork: True if the repository was created as a fork.
        :param pulumi.Input[str] name: The name of the git repository.
        :param pulumi.Input[str] parent_repository_id: The ID of a Git project from which a fork is to be created.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] remote_url: Git HTTPS URL of the repository
        :param pulumi.Input[int] size: Size in bytes.
        :param pulumi.Input[str] ssh_url: Git SSH URL of the repository.
        :param pulumi.Input[str] url: REST API URL of the repository.
        :param pulumi.Input[str] web_url: Web link to the repository.
        """
        if default_branch is not None:
            pulumi.set(__self__, "default_branch", default_branch)
        if initialization is not None:
            pulumi.set(__self__, "initialization", initialization)
        if is_fork is not None:
            pulumi.set(__self__, "is_fork", is_fork)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_repository_id is not None:
            pulumi.set(__self__, "parent_repository_id", parent_repository_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if remote_url is not None:
            pulumi.set(__self__, "remote_url", remote_url)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if ssh_url is not None:
            pulumi.set(__self__, "ssh_url", ssh_url)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if web_url is not None:
            pulumi.set(__self__, "web_url", web_url)

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> Optional[pulumi.Input[str]]:
        """
        The ref of the default branch. Will be used as the branch name for initialized repositories.
        """
        return pulumi.get(self, "default_branch")

    @default_branch.setter
    def default_branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_branch", value)

    @property
    @pulumi.getter
    def initialization(self) -> Optional[pulumi.Input['GitInitializationArgs']]:
        """
        An `initialization` block as documented below.
        """
        return pulumi.get(self, "initialization")

    @initialization.setter
    def initialization(self, value: Optional[pulumi.Input['GitInitializationArgs']]):
        pulumi.set(self, "initialization", value)

    @property
    @pulumi.getter(name="isFork")
    def is_fork(self) -> Optional[pulumi.Input[bool]]:
        """
        True if the repository was created as a fork.
        """
        return pulumi.get(self, "is_fork")

    @is_fork.setter
    def is_fork(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_fork", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the git repository.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentRepositoryId")
    def parent_repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a Git project from which a fork is to be created.
        """
        return pulumi.get(self, "parent_repository_id")

    @parent_repository_id.setter
    def parent_repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_repository_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> Optional[pulumi.Input[str]]:
        """
        Git HTTPS URL of the repository
        """
        return pulumi.get(self, "remote_url")

    @remote_url.setter
    def remote_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_url", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Size in bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="sshUrl")
    def ssh_url(self) -> Optional[pulumi.Input[str]]:
        """
        Git SSH URL of the repository.
        """
        return pulumi.get(self, "ssh_url")

    @ssh_url.setter
    def ssh_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_url", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        REST API URL of the repository.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> Optional[pulumi.Input[str]]:
        """
        Web link to the repository.
        """
        return pulumi.get(self, "web_url")

    @web_url.setter
    def web_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_url", value)


class Git(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_branch: Optional[pulumi.Input[str]] = None,
                 initialization: Optional[pulumi.Input[pulumi.InputType['GitInitializationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_repository_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a git repository within Azure DevOps.

        ## Example Usage
        ### Create Git repository

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        repo = azuredevops.Git("repo",
            project_id=project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        ```
        ### Create Fork of another Azure DevOps Git repository

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        repo = azuredevops.Git("repo",
            project_id=azuredevops_project["project"]["id"],
            parent_repository_id=azuredevops_git_repository["parent"]["id"],
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        ```
        ### Create Import from another Git repository

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        repo = azuredevops.Git("repo",
            project_id=azuredevops_project["project"]["id"],
            initialization=azuredevops.GitInitializationArgs(
                init_type="Import",
                source_type="Git",
                source_url="https://github.com/microsoft/terraform-provider-azuredevops.git",
            ))
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Git Repositories](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Repositories can be imported using the repo name or by the repo Guid e.g.

        ```sh
         $ pulumi import azuredevops:index/git:Git repository projectName/repoName
        ```

         or

        ```sh
         $ pulumi import azuredevops:index/git:Git repository projectName/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_branch: The ref of the default branch. Will be used as the branch name for initialized repositories.
        :param pulumi.Input[pulumi.InputType['GitInitializationArgs']] initialization: An `initialization` block as documented below.
        :param pulumi.Input[str] name: The name of the git repository.
        :param pulumi.Input[str] parent_repository_id: The ID of a Git project from which a fork is to be created.
        :param pulumi.Input[str] project_id: The project ID or project name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a git repository within Azure DevOps.

        ## Example Usage
        ### Create Git repository

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        repo = azuredevops.Git("repo",
            project_id=project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        ```
        ### Create Fork of another Azure DevOps Git repository

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        repo = azuredevops.Git("repo",
            project_id=azuredevops_project["project"]["id"],
            parent_repository_id=azuredevops_git_repository["parent"]["id"],
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        ```
        ### Create Import from another Git repository

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        repo = azuredevops.Git("repo",
            project_id=azuredevops_project["project"]["id"],
            initialization=azuredevops.GitInitializationArgs(
                init_type="Import",
                source_type="Git",
                source_url="https://github.com/microsoft/terraform-provider-azuredevops.git",
            ))
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Git Repositories](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Repositories can be imported using the repo name or by the repo Guid e.g.

        ```sh
         $ pulumi import azuredevops:index/git:Git repository projectName/repoName
        ```

         or

        ```sh
         $ pulumi import azuredevops:index/git:Git repository projectName/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param GitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_branch: Optional[pulumi.Input[str]] = None,
                 initialization: Optional[pulumi.Input[pulumi.InputType['GitInitializationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_repository_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = GitArgs.__new__(GitArgs)

            __props__.__dict__["default_branch"] = default_branch
            if initialization is None and not opts.urn:
                raise TypeError("Missing required property 'initialization'")
            __props__.__dict__["initialization"] = initialization
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_repository_id"] = parent_repository_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["is_fork"] = None
            __props__.__dict__["remote_url"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["ssh_url"] = None
            __props__.__dict__["url"] = None
            __props__.__dict__["web_url"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Repository/git:Git")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Git, __self__).__init__(
            'azuredevops:index/git:Git',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_branch: Optional[pulumi.Input[str]] = None,
            initialization: Optional[pulumi.Input[pulumi.InputType['GitInitializationArgs']]] = None,
            is_fork: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_repository_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            remote_url: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            ssh_url: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            web_url: Optional[pulumi.Input[str]] = None) -> 'Git':
        """
        Get an existing Git resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_branch: The ref of the default branch. Will be used as the branch name for initialized repositories.
        :param pulumi.Input[pulumi.InputType['GitInitializationArgs']] initialization: An `initialization` block as documented below.
        :param pulumi.Input[bool] is_fork: True if the repository was created as a fork.
        :param pulumi.Input[str] name: The name of the git repository.
        :param pulumi.Input[str] parent_repository_id: The ID of a Git project from which a fork is to be created.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] remote_url: Git HTTPS URL of the repository
        :param pulumi.Input[int] size: Size in bytes.
        :param pulumi.Input[str] ssh_url: Git SSH URL of the repository.
        :param pulumi.Input[str] url: REST API URL of the repository.
        :param pulumi.Input[str] web_url: Web link to the repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GitState.__new__(_GitState)

        __props__.__dict__["default_branch"] = default_branch
        __props__.__dict__["initialization"] = initialization
        __props__.__dict__["is_fork"] = is_fork
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_repository_id"] = parent_repository_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["remote_url"] = remote_url
        __props__.__dict__["size"] = size
        __props__.__dict__["ssh_url"] = ssh_url
        __props__.__dict__["url"] = url
        __props__.__dict__["web_url"] = web_url
        return Git(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> pulumi.Output[str]:
        """
        The ref of the default branch. Will be used as the branch name for initialized repositories.
        """
        return pulumi.get(self, "default_branch")

    @property
    @pulumi.getter
    def initialization(self) -> pulumi.Output['outputs.GitInitialization']:
        """
        An `initialization` block as documented below.
        """
        return pulumi.get(self, "initialization")

    @property
    @pulumi.getter(name="isFork")
    def is_fork(self) -> pulumi.Output[bool]:
        """
        True if the repository was created as a fork.
        """
        return pulumi.get(self, "is_fork")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the git repository.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentRepositoryId")
    def parent_repository_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of a Git project from which a fork is to be created.
        """
        return pulumi.get(self, "parent_repository_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> pulumi.Output[str]:
        """
        Git HTTPS URL of the repository
        """
        return pulumi.get(self, "remote_url")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Size in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sshUrl")
    def ssh_url(self) -> pulumi.Output[str]:
        """
        Git SSH URL of the repository.
        """
        return pulumi.get(self, "ssh_url")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        REST API URL of the repository.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> pulumi.Output[str]:
        """
        Web link to the repository.
        """
        return pulumi.get(self, "web_url")

