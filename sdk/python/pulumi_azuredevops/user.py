# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 account_license_type: Optional[pulumi.Input[str]] = None,
                 licensing_source: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 principal_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] account_license_type: Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        :param pulumi.Input[str] licensing_source: The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        :param pulumi.Input[str] origin: The type of source provider for the origin identifier.
        :param pulumi.Input[str] origin_id: The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        :param pulumi.Input[str] principal_name: The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        if account_license_type is not None:
            pulumi.set(__self__, "account_license_type", account_license_type)
        if licensing_source is not None:
            pulumi.set(__self__, "licensing_source", licensing_source)
        if origin is not None:
            pulumi.set(__self__, "origin", origin)
        if origin_id is not None:
            pulumi.set(__self__, "origin_id", origin_id)
        if principal_name is not None:
            pulumi.set(__self__, "principal_name", principal_name)

    @property
    @pulumi.getter(name="accountLicenseType")
    def account_license_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        """
        return pulumi.get(self, "account_license_type")

    @account_license_type.setter
    def account_license_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_license_type", value)

    @property
    @pulumi.getter(name="licensingSource")
    def licensing_source(self) -> Optional[pulumi.Input[str]]:
        """
        The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        """
        return pulumi.get(self, "licensing_source")

    @licensing_source.setter
    def licensing_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "licensing_source", value)

    @property
    @pulumi.getter
    def origin(self) -> Optional[pulumi.Input[str]]:
        """
        The type of source provider for the origin identifier.
        """
        return pulumi.get(self, "origin")

    @origin.setter
    def origin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin", value)

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        """
        return pulumi.get(self, "origin_id")

    @origin_id.setter
    def origin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_id", value)

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> Optional[pulumi.Input[str]]:
        """
        The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        return pulumi.get(self, "principal_name")

    @principal_name.setter
    def principal_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_name", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 account_license_type: Optional[pulumi.Input[str]] = None,
                 descriptor: Optional[pulumi.Input[str]] = None,
                 licensing_source: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 principal_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] account_license_type: Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        :param pulumi.Input[str] descriptor: The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
        :param pulumi.Input[str] licensing_source: The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        :param pulumi.Input[str] origin: The type of source provider for the origin identifier.
        :param pulumi.Input[str] origin_id: The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        :param pulumi.Input[str] principal_name: The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        if account_license_type is not None:
            pulumi.set(__self__, "account_license_type", account_license_type)
        if descriptor is not None:
            pulumi.set(__self__, "descriptor", descriptor)
        if licensing_source is not None:
            pulumi.set(__self__, "licensing_source", licensing_source)
        if origin is not None:
            pulumi.set(__self__, "origin", origin)
        if origin_id is not None:
            pulumi.set(__self__, "origin_id", origin_id)
        if principal_name is not None:
            pulumi.set(__self__, "principal_name", principal_name)

    @property
    @pulumi.getter(name="accountLicenseType")
    def account_license_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        """
        return pulumi.get(self, "account_license_type")

    @account_license_type.setter
    def account_license_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_license_type", value)

    @property
    @pulumi.getter
    def descriptor(self) -> Optional[pulumi.Input[str]]:
        """
        The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
        """
        return pulumi.get(self, "descriptor")

    @descriptor.setter
    def descriptor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "descriptor", value)

    @property
    @pulumi.getter(name="licensingSource")
    def licensing_source(self) -> Optional[pulumi.Input[str]]:
        """
        The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        """
        return pulumi.get(self, "licensing_source")

    @licensing_source.setter
    def licensing_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "licensing_source", value)

    @property
    @pulumi.getter
    def origin(self) -> Optional[pulumi.Input[str]]:
        """
        The type of source provider for the origin identifier.
        """
        return pulumi.get(self, "origin")

    @origin.setter
    def origin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin", value)

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        """
        return pulumi.get(self, "origin_id")

    @origin_id.setter
    def origin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_id", value)

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> Optional[pulumi.Input[str]]:
        """
        The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        return pulumi.get(self, "principal_name")

    @principal_name.setter
    def principal_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_name", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_license_type: Optional[pulumi.Input[str]] = None,
                 licensing_source: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 principal_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a user entitlement within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.User("example", principal_name="foo@contoso.com")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user-entitlements/add?view=azure-devops-rest-7.0)
        - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)

        ## PAT Permissions Required

        - **Member Entitlement Management**: Read & Write

        ## Import

        The resources allows the import via the UUID of a user entitlement or by using the principal name of a user owning an entitlement.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_license_type: Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        :param pulumi.Input[str] licensing_source: The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        :param pulumi.Input[str] origin: The type of source provider for the origin identifier.
        :param pulumi.Input[str] origin_id: The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        :param pulumi.Input[str] principal_name: The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a user entitlement within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.User("example", principal_name="foo@contoso.com")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - User Entitlements - Add](https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user-entitlements/add?view=azure-devops-rest-7.0)
        - [Programmatic mapping of access levels](https://docs.microsoft.com/en-us/azure/devops/organizations/security/access-levels?view=azure-devops#programmatic-mapping-of-access-levels)

        ## PAT Permissions Required

        - **Member Entitlement Management**: Read & Write

        ## Import

        The resources allows the import via the UUID of a user entitlement or by using the principal name of a user owning an entitlement.

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_license_type: Optional[pulumi.Input[str]] = None,
                 licensing_source: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 origin_id: Optional[pulumi.Input[str]] = None,
                 principal_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["account_license_type"] = account_license_type
            __props__.__dict__["licensing_source"] = licensing_source
            __props__.__dict__["origin"] = origin
            __props__.__dict__["origin_id"] = origin_id
            __props__.__dict__["principal_name"] = principal_name
            __props__.__dict__["descriptor"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Entitlement/user:User")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(User, __self__).__init__(
            'azuredevops:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_license_type: Optional[pulumi.Input[str]] = None,
            descriptor: Optional[pulumi.Input[str]] = None,
            licensing_source: Optional[pulumi.Input[str]] = None,
            origin: Optional[pulumi.Input[str]] = None,
            origin_id: Optional[pulumi.Input[str]] = None,
            principal_name: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_license_type: Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        :param pulumi.Input[str] descriptor: The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
        :param pulumi.Input[str] licensing_source: The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        :param pulumi.Input[str] origin: The type of source provider for the origin identifier.
        :param pulumi.Input[str] origin_id: The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        :param pulumi.Input[str] principal_name: The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["account_license_type"] = account_license_type
        __props__.__dict__["descriptor"] = descriptor
        __props__.__dict__["licensing_source"] = licensing_source
        __props__.__dict__["origin"] = origin
        __props__.__dict__["origin_id"] = origin_id
        __props__.__dict__["principal_name"] = principal_name
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountLicenseType")
    def account_license_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of Account License. Valid values: `advanced`, `earlyAdopter`, `express`, `none`, `professional`, or `stakeholder`. Defaults to `express`. In addition the value `basic` is allowed which is an alias for `express` and reflects the name of the `express` license used in the Azure DevOps web interface.
        """
        return pulumi.get(self, "account_license_type")

    @property
    @pulumi.getter
    def descriptor(self) -> pulumi.Output[str]:
        """
        The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject.
        """
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter(name="licensingSource")
    def licensing_source(self) -> pulumi.Output[Optional[str]]:
        """
        The source of the licensing (e.g. Account. MSDN etc.) Valid values: `account` (Default), `auto`, `msdn`, `none`, `profile`, `trial`**NOTE:** A user can only be referenced by it's `principal_name` or by the combination of `origin_id` and `origin`.
        """
        return pulumi.get(self, "licensing_source")

    @property
    @pulumi.getter
    def origin(self) -> pulumi.Output[str]:
        """
        The type of source provider for the origin identifier.
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> pulumi.Output[str]:
        """
        The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.
        """
        return pulumi.get(self, "origin_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> pulumi.Output[str]:
        """
        The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.
        """
        return pulumi.get(self, "principal_name")

