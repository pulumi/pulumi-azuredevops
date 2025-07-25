# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['WikiArgs', 'Wiki']

@pulumi.input_type
class WikiArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[_builtins.str],
                 mapped_path: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[_builtins.str]] = None,
                 version: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Wiki resource.
        :param pulumi.Input[_builtins.str] type: The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        :param pulumi.Input[_builtins.str] mapped_path: Folder path inside repository which is shown as Wiki.
        :param pulumi.Input[_builtins.str] name: The name of the Wiki.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project.
        :param pulumi.Input[_builtins.str] repository_id: The ID of the repository.
        :param pulumi.Input[_builtins.str] version: Version of the wiki.
        """
        pulumi.set(__self__, "type", type)
        if mapped_path is not None:
            pulumi.set(__self__, "mapped_path", mapped_path)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Input[_builtins.str]:
        """
        The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter(name="mappedPath")
    def mapped_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Folder path inside repository which is shown as Wiki.
        """
        return pulumi.get(self, "mapped_path")

    @mapped_path.setter
    def mapped_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mapped_path", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the Wiki.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the Project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the repository.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "repository_id", value)

    @_builtins.property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Version of the wiki.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _WikiState:
    def __init__(__self__, *,
                 mapped_path: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 remote_url: Optional[pulumi.Input[_builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 url: Optional[pulumi.Input[_builtins.str]] = None,
                 version: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering Wiki resources.
        :param pulumi.Input[_builtins.str] mapped_path: Folder path inside repository which is shown as Wiki.
        :param pulumi.Input[_builtins.str] name: The name of the Wiki.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project.
        :param pulumi.Input[_builtins.str] remote_url: The remote web url to the wiki.
        :param pulumi.Input[_builtins.str] repository_id: The ID of the repository.
        :param pulumi.Input[_builtins.str] type: The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        :param pulumi.Input[_builtins.str] url: The REST url for this wiki.
        :param pulumi.Input[_builtins.str] version: Version of the wiki.
        """
        if mapped_path is not None:
            pulumi.set(__self__, "mapped_path", mapped_path)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if remote_url is not None:
            pulumi.set(__self__, "remote_url", remote_url)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @_builtins.property
    @pulumi.getter(name="mappedPath")
    def mapped_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Folder path inside repository which is shown as Wiki.
        """
        return pulumi.get(self, "mapped_path")

    @mapped_path.setter
    def mapped_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mapped_path", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the Wiki.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the Project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The remote web url to the wiki.
        """
        return pulumi.get(self, "remote_url")

    @remote_url.setter
    def remote_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "remote_url", value)

    @_builtins.property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the repository.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "repository_id", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The REST url for this wiki.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "url", value)

    @_builtins.property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Version of the wiki.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "version", value)


@pulumi.type_token("azuredevops:index/wiki:Wiki")
class Wiki(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mapped_path: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 version: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages Wikis within Azure DevOps project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            description="Managed by Pulumi")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "init_type": "Clean",
            })
        example_wiki = azuredevops.Wiki("example",
            name="Example project wiki ",
            project_id=example.id,
            type="projectWiki")
        example2 = azuredevops.Wiki("example2",
            name="Example wiki in repository",
            project_id=example.id,
            repository_id=example_git.id,
            version="main",
            type="codeWiki",
            mapped_path="/")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.1 - Wiki ](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/wikis?view=azure-devops-rest-7.1)

        ## Import

        Azure DevOps Wiki can be imported using the `id`

        ```sh
        $ pulumi import azuredevops:index/wiki:Wiki wiki 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] mapped_path: Folder path inside repository which is shown as Wiki.
        :param pulumi.Input[_builtins.str] name: The name of the Wiki.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project.
        :param pulumi.Input[_builtins.str] repository_id: The ID of the repository.
        :param pulumi.Input[_builtins.str] type: The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        :param pulumi.Input[_builtins.str] version: Version of the wiki.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WikiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Wikis within Azure DevOps project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            description="Managed by Pulumi")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "init_type": "Clean",
            })
        example_wiki = azuredevops.Wiki("example",
            name="Example project wiki ",
            project_id=example.id,
            type="projectWiki")
        example2 = azuredevops.Wiki("example2",
            name="Example wiki in repository",
            project_id=example.id,
            repository_id=example_git.id,
            version="main",
            type="codeWiki",
            mapped_path="/")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.1 - Wiki ](https://learn.microsoft.com/en-us/rest/api/azure/devops/wiki/wikis?view=azure-devops-rest-7.1)

        ## Import

        Azure DevOps Wiki can be imported using the `id`

        ```sh
        $ pulumi import azuredevops:index/wiki:Wiki wiki 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param WikiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WikiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mapped_path: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 repository_id: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 version: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WikiArgs.__new__(WikiArgs)

            __props__.__dict__["mapped_path"] = mapped_path
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["repository_id"] = repository_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["version"] = version
            __props__.__dict__["remote_url"] = None
            __props__.__dict__["url"] = None
        super(Wiki, __self__).__init__(
            'azuredevops:index/wiki:Wiki',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            mapped_path: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None,
            remote_url: Optional[pulumi.Input[_builtins.str]] = None,
            repository_id: Optional[pulumi.Input[_builtins.str]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None,
            url: Optional[pulumi.Input[_builtins.str]] = None,
            version: Optional[pulumi.Input[_builtins.str]] = None) -> 'Wiki':
        """
        Get an existing Wiki resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] mapped_path: Folder path inside repository which is shown as Wiki.
        :param pulumi.Input[_builtins.str] name: The name of the Wiki.
        :param pulumi.Input[_builtins.str] project_id: The ID of the Project.
        :param pulumi.Input[_builtins.str] remote_url: The remote web url to the wiki.
        :param pulumi.Input[_builtins.str] repository_id: The ID of the repository.
        :param pulumi.Input[_builtins.str] type: The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        :param pulumi.Input[_builtins.str] url: The REST url for this wiki.
        :param pulumi.Input[_builtins.str] version: Version of the wiki.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WikiState.__new__(_WikiState)

        __props__.__dict__["mapped_path"] = mapped_path
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["remote_url"] = remote_url
        __props__.__dict__["repository_id"] = repository_id
        __props__.__dict__["type"] = type
        __props__.__dict__["url"] = url
        __props__.__dict__["version"] = version
        return Wiki(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="mappedPath")
    def mapped_path(self) -> pulumi.Output[_builtins.str]:
        """
        Folder path inside repository which is shown as Wiki.
        """
        return pulumi.get(self, "mapped_path")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the Wiki.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The ID of the Project.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> pulumi.Output[_builtins.str]:
        """
        The remote web url to the wiki.
        """
        return pulumi.get(self, "remote_url")

    @_builtins.property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the repository.
        """
        return pulumi.get(self, "repository_id")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        The type of the wiki. Possible values are `codeWiki`, `projectWiki`.
        """
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter
    def url(self) -> pulumi.Output[_builtins.str]:
        """
        The REST url for this wiki.
        """
        return pulumi.get(self, "url")

    @_builtins.property
    @pulumi.getter
    def version(self) -> pulumi.Output[_builtins.str]:
        """
        Version of the wiki.
        """
        return pulumi.get(self, "version")

