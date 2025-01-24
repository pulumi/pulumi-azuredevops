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

__all__ = ['ProjectTagsArgs', 'ProjectTags']

@pulumi.input_type
class ProjectTagsArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 tags: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a ProjectTags resource.
        :param pulumi.Input[str] project_id: The ID of the Project. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A mapping of tags assigned to the Project.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the Project. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A mapping of tags assigned to the Project.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ProjectTagsState:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ProjectTags resources.
        :param pulumi.Input[str] project_id: The ID of the Project. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A mapping of tags assigned to the Project.
        """
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Project. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A mapping of tags assigned to the Project.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class ProjectTags(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages Project Tags within Azure DevOps organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_project_tags = azuredevops.ProjectTags("example",
            project_id=example.id,
            tags=[
                "tag1",
                "tag2",
            ])
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Project Properties](https://learn.microsoft.com/en-us/rest/api/azure/devops/core/projects/get-project-properties?view=azure-devops-rest-7.1&tabs=HTTP)

        ## Import

        Azure DevOps Project Tags can be imported using the Project ID e.g.:

        ```sh
        $ pulumi import azuredevops:index/projectTags:ProjectTags example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The ID of the Project. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A mapping of tags assigned to the Project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectTagsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Project Tags within Azure DevOps organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_project_tags = azuredevops.ProjectTags("example",
            project_id=example.id,
            tags=[
                "tag1",
                "tag2",
            ])
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Project Properties](https://learn.microsoft.com/en-us/rest/api/azure/devops/core/projects/get-project-properties?view=azure-devops-rest-7.1&tabs=HTTP)

        ## Import

        Azure DevOps Project Tags can be imported using the Project ID e.g.:

        ```sh
        $ pulumi import azuredevops:index/projectTags:ProjectTags example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ProjectTagsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectTagsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectTagsArgs.__new__(ProjectTagsArgs)

            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if tags is None and not opts.urn:
                raise TypeError("Missing required property 'tags'")
            __props__.__dict__["tags"] = tags
        super(ProjectTags, __self__).__init__(
            'azuredevops:index/projectTags:ProjectTags',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ProjectTags':
        """
        Get an existing ProjectTags resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The ID of the Project. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A mapping of tags assigned to the Project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectTagsState.__new__(_ProjectTagsState)

        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["tags"] = tags
        return ProjectTags(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the Project. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        A mapping of tags assigned to the Project.
        """
        return pulumi.get(self, "tags")

