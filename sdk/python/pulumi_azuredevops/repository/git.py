# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Git']

warnings.warn("azuredevops.repository.Git has been deprecated in favor of azuredevops.Git", DeprecationWarning)


class Git(pulumi.CustomResource):
    warnings.warn("azuredevops.repository.Git has been deprecated in favor of azuredevops.Git", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_branch: Optional[pulumi.Input[str]] = None,
                 initialization: Optional[pulumi.Input[pulumi.InputType['GitInitializationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_repository_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

        - [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/repositories?view=azure-devops-rest-5.1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_branch: The ref of the default branch. Will be used as the branch name for initialized repositories.
        :param pulumi.Input[pulumi.InputType['GitInitializationArgs']] initialization: An `initialization` block as documented below.
        :param pulumi.Input[str] name: The name of the git repository.
        :param pulumi.Input[str] parent_repository_id: The ID of a Git project from which a fork is to be created.
        :param pulumi.Input[str] project_id: The project ID or project name.
        """
        pulumi.log.warn("Git is deprecated: azuredevops.repository.Git has been deprecated in favor of azuredevops.Git")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['default_branch'] = default_branch
            if initialization is None:
                raise TypeError("Missing required property 'initialization'")
            __props__['initialization'] = initialization
            __props__['name'] = name
            __props__['parent_repository_id'] = parent_repository_id
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['is_fork'] = None
            __props__['remote_url'] = None
            __props__['size'] = None
            __props__['ssh_url'] = None
            __props__['url'] = None
            __props__['web_url'] = None
        super(Git, __self__).__init__(
            'azuredevops:Repository/git:Git',
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

        __props__ = dict()

        __props__["default_branch"] = default_branch
        __props__["initialization"] = initialization
        __props__["is_fork"] = is_fork
        __props__["name"] = name
        __props__["parent_repository_id"] = parent_repository_id
        __props__["project_id"] = project_id
        __props__["remote_url"] = remote_url
        __props__["size"] = size
        __props__["ssh_url"] = ssh_url
        __props__["url"] = url
        __props__["web_url"] = web_url
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

