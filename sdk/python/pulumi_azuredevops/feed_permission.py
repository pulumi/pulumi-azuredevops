# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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

__all__ = ['FeedPermissionArgs', 'FeedPermission']

@pulumi.input_type
class FeedPermissionArgs:
    def __init__(__self__, *,
                 feed_id: pulumi.Input[builtins.str],
                 identity_descriptor: pulumi.Input[builtins.str],
                 role: pulumi.Input[builtins.str],
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a FeedPermission resource.
        :param pulumi.Input[builtins.str] feed_id: The ID of the Feed.
        :param pulumi.Input[builtins.str] identity_descriptor: The Descriptor of identity you want to assign a role.
        :param pulumi.Input[builtins.str] role: The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        :param pulumi.Input[builtins.str] display_name: The display name of the assignment
        :param pulumi.Input[builtins.str] project_id: The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        """
        pulumi.set(__self__, "feed_id", feed_id)
        pulumi.set(__self__, "identity_descriptor", identity_descriptor)
        pulumi.set(__self__, "role", role)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the Feed.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "feed_id", value)

    @property
    @pulumi.getter(name="identityDescriptor")
    def identity_descriptor(self) -> pulumi.Input[builtins.str]:
        """
        The Descriptor of identity you want to assign a role.
        """
        return pulumi.get(self, "identity_descriptor")

    @identity_descriptor.setter
    def identity_descriptor(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "identity_descriptor", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[builtins.str]:
        """
        The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The display name of the assignment
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _FeedPermissionState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 feed_id: Optional[pulumi.Input[builtins.str]] = None,
                 identity_descriptor: Optional[pulumi.Input[builtins.str]] = None,
                 identity_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering FeedPermission resources.
        :param pulumi.Input[builtins.str] display_name: The display name of the assignment
        :param pulumi.Input[builtins.str] feed_id: The ID of the Feed.
        :param pulumi.Input[builtins.str] identity_descriptor: The Descriptor of identity you want to assign a role.
        :param pulumi.Input[builtins.str] identity_id: The ID of the identity.
        :param pulumi.Input[builtins.str] project_id: The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        :param pulumi.Input[builtins.str] role: The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if feed_id is not None:
            pulumi.set(__self__, "feed_id", feed_id)
        if identity_descriptor is not None:
            pulumi.set(__self__, "identity_descriptor", identity_descriptor)
        if identity_id is not None:
            pulumi.set(__self__, "identity_id", identity_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The display name of the assignment
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the Feed.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "feed_id", value)

    @property
    @pulumi.getter(name="identityDescriptor")
    def identity_descriptor(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Descriptor of identity you want to assign a role.
        """
        return pulumi.get(self, "identity_descriptor")

    @identity_descriptor.setter
    def identity_descriptor(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "identity_descriptor", value)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the identity.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "identity_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role", value)


class FeedPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 feed_id: Optional[pulumi.Input[builtins.str]] = None,
                 identity_descriptor: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages creation of the Feed Permission within Azure DevOps organization.

        ## Example Usage

        ### Create Feed Permission
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="Example Project")
        example_group = azuredevops.Group("example",
            scope=example.id,
            display_name="Example group",
            description="Example description")
        example_feed = azuredevops.Feed("example", name="examplefeed")
        permission = azuredevops.FeedPermission("permission",
            feed_id=example_feed.id,
            role="reader",
            identity_descriptor=example_group.descriptor)
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] display_name: The display name of the assignment
        :param pulumi.Input[builtins.str] feed_id: The ID of the Feed.
        :param pulumi.Input[builtins.str] identity_descriptor: The Descriptor of identity you want to assign a role.
        :param pulumi.Input[builtins.str] project_id: The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        :param pulumi.Input[builtins.str] role: The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FeedPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages creation of the Feed Permission within Azure DevOps organization.

        ## Example Usage

        ### Create Feed Permission
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="Example Project")
        example_group = azuredevops.Group("example",
            scope=example.id,
            display_name="Example group",
            description="Example description")
        example_feed = azuredevops.Feed("example", name="examplefeed")
        permission = azuredevops.FeedPermission("permission",
            feed_id=example_feed.id,
            role="reader",
            identity_descriptor=example_group.descriptor)
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Feed Management](https://learn.microsoft.com/en-us/rest/api/azure/devops/artifacts/feed-management?view=azure-devops-rest-7.0)

        :param str resource_name: The name of the resource.
        :param FeedPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FeedPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 feed_id: Optional[pulumi.Input[builtins.str]] = None,
                 identity_descriptor: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FeedPermissionArgs.__new__(FeedPermissionArgs)

            __props__.__dict__["display_name"] = display_name
            if feed_id is None and not opts.urn:
                raise TypeError("Missing required property 'feed_id'")
            __props__.__dict__["feed_id"] = feed_id
            if identity_descriptor is None and not opts.urn:
                raise TypeError("Missing required property 'identity_descriptor'")
            __props__.__dict__["identity_descriptor"] = identity_descriptor
            __props__.__dict__["project_id"] = project_id
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["identity_id"] = None
        super(FeedPermission, __self__).__init__(
            'azuredevops:index/feedPermission:FeedPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[builtins.str]] = None,
            feed_id: Optional[pulumi.Input[builtins.str]] = None,
            identity_descriptor: Optional[pulumi.Input[builtins.str]] = None,
            identity_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            role: Optional[pulumi.Input[builtins.str]] = None) -> 'FeedPermission':
        """
        Get an existing FeedPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] display_name: The display name of the assignment
        :param pulumi.Input[builtins.str] feed_id: The ID of the Feed.
        :param pulumi.Input[builtins.str] identity_descriptor: The Descriptor of identity you want to assign a role.
        :param pulumi.Input[builtins.str] identity_id: The ID of the identity.
        :param pulumi.Input[builtins.str] project_id: The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        :param pulumi.Input[builtins.str] role: The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FeedPermissionState.__new__(_FeedPermissionState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["feed_id"] = feed_id
        __props__.__dict__["identity_descriptor"] = identity_descriptor
        __props__.__dict__["identity_id"] = identity_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["role"] = role
        return FeedPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The display name of the assignment
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Feed.
        """
        return pulumi.get(self, "feed_id")

    @property
    @pulumi.getter(name="identityDescriptor")
    def identity_descriptor(self) -> pulumi.Output[builtins.str]:
        """
        The Descriptor of identity you want to assign a role.
        """
        return pulumi.get(self, "identity_descriptor")

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the identity.
        """
        return pulumi.get(self, "identity_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the Project Feed is created in. If not specified, feed will be created at the organization level.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[builtins.str]:
        """
        The role to be assigned. Possible values are: `reader`, `contributor`, `collaborator`, `administrator`
        """
        return pulumi.get(self, "role")

