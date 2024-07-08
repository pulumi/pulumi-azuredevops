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

__all__ = ['SecurityroleAssignmentArgs', 'SecurityroleAssignment']

@pulumi.input_type
class SecurityroleAssignmentArgs:
    def __init__(__self__, *,
                 identity_id: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 role_name: pulumi.Input[str],
                 scope: pulumi.Input[str]):
        """
        The set of arguments for constructing a SecurityroleAssignment resource.
        :param pulumi.Input[str] identity_id: The ID of the identity to authorize.
        :param pulumi.Input[str] resource_id: The ID of the resource on which the role is to be assigned.
        :param pulumi.Input[str] role_name: Name of the role to assign.
        :param pulumi.Input[str] scope: The scope in which this assignment should exist.
        """
        pulumi.set(__self__, "identity_id", identity_id)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "role_name", role_name)
        pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Input[str]:
        """
        The ID of the identity to authorize.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource on which the role is to be assigned.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        """
        Name of the role to assign.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        """
        The scope in which this assignment should exist.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class _SecurityroleAssignmentState:
    def __init__(__self__, *,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityroleAssignment resources.
        :param pulumi.Input[str] identity_id: The ID of the identity to authorize.
        :param pulumi.Input[str] resource_id: The ID of the resource on which the role is to be assigned.
        :param pulumi.Input[str] role_name: Name of the role to assign.
        :param pulumi.Input[str] scope: The scope in which this assignment should exist.
        """
        if identity_id is not None:
            pulumi.set(__self__, "identity_id", identity_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the identity to authorize.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource on which the role is to be assigned.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role to assign.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The scope in which this assignment should exist.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


class SecurityroleAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages assignment of security roles to various resources within Azure DevOps organization.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_id: The ID of the identity to authorize.
        :param pulumi.Input[str] resource_id: The ID of the resource on which the role is to be assigned.
        :param pulumi.Input[str] role_name: Name of the role to assign.
        :param pulumi.Input[str] scope: The scope in which this assignment should exist.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityroleAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages assignment of security roles to various resources within Azure DevOps organization.

        :param str resource_name: The name of the resource.
        :param SecurityroleAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityroleAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityroleAssignmentArgs.__new__(SecurityroleAssignmentArgs)

            if identity_id is None and not opts.urn:
                raise TypeError("Missing required property 'identity_id'")
            __props__.__dict__["identity_id"] = identity_id
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__.__dict__["scope"] = scope
        super(SecurityroleAssignment, __self__).__init__(
            'azuredevops:index/securityroleAssignment:SecurityroleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            identity_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'SecurityroleAssignment':
        """
        Get an existing SecurityroleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_id: The ID of the identity to authorize.
        :param pulumi.Input[str] resource_id: The ID of the resource on which the role is to be assigned.
        :param pulumi.Input[str] role_name: Name of the role to assign.
        :param pulumi.Input[str] scope: The scope in which this assignment should exist.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityroleAssignmentState.__new__(_SecurityroleAssignmentState)

        __props__.__dict__["identity_id"] = identity_id
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["role_name"] = role_name
        __props__.__dict__["scope"] = scope
        return SecurityroleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Output[str]:
        """
        The ID of the identity to authorize.
        """
        return pulumi.get(self, "identity_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource on which the role is to be assigned.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        Name of the role to assign.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The scope in which this assignment should exist.
        """
        return pulumi.get(self, "scope")

