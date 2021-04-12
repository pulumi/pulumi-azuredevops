# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetUsersUserResult',
]

@pulumi.output_type
class GetUsersUserResult(dict):
    def __init__(__self__, *,
                 descriptor: str,
                 display_name: str,
                 mail_address: str,
                 origin: str,
                 principal_name: str,
                 origin_id: Optional[str] = None):
        """
        :param str descriptor: The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
        :param str display_name: This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
        :param str mail_address: The email address of record for a given graph member. This may be different than the principal name.
        :param str origin: The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
        :param str principal_name: The PrincipalName of this graph member from the source provider.
        :param str origin_id: The unique identifier from the system of origin.
        """
        pulumi.set(__self__, "descriptor", descriptor)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "mail_address", mail_address)
        pulumi.set(__self__, "origin", origin)
        pulumi.set(__self__, "principal_name", principal_name)
        if origin_id is not None:
            pulumi.set(__self__, "origin_id", origin_id)

    @property
    @pulumi.getter
    def descriptor(self) -> str:
        """
        The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
        """
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="mailAddress")
    def mail_address(self) -> str:
        """
        The email address of record for a given graph member. This may be different than the principal name.
        """
        return pulumi.get(self, "mail_address")

    @property
    @pulumi.getter
    def origin(self) -> str:
        """
        The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> str:
        """
        The PrincipalName of this graph member from the source provider.
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[str]:
        """
        The unique identifier from the system of origin.
        """
        return pulumi.get(self, "origin_id")


