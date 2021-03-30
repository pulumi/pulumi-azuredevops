# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs

__all__ = [
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, id=None, origin=None, origin_id=None, principal_name=None, subject_types=None, users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if origin and not isinstance(origin, str):
            raise TypeError("Expected argument 'origin' to be a str")
        pulumi.set(__self__, "origin", origin)
        if origin_id and not isinstance(origin_id, str):
            raise TypeError("Expected argument 'origin_id' to be a str")
        pulumi.set(__self__, "origin_id", origin_id)
        if principal_name and not isinstance(principal_name, str):
            raise TypeError("Expected argument 'principal_name' to be a str")
        pulumi.set(__self__, "principal_name", principal_name)
        if subject_types and not isinstance(subject_types, list):
            raise TypeError("Expected argument 'subject_types' to be a list")
        pulumi.set(__self__, "subject_types", subject_types)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def origin(self) -> Optional[str]:
        """
        The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> Optional[str]:
        """
        The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
        """
        return pulumi.get(self, "origin_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> Optional[str]:
        """
        This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="subjectTypes")
    def subject_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "subject_types")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        """
        A set of existing users in your Azure DevOps Organization with details about every single user which includes:
        """
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            id=self.id,
            origin=self.origin,
            origin_id=self.origin_id,
            principal_name=self.principal_name,
            subject_types=self.subject_types,
            users=self.users)


def get_users(origin: Optional[str] = None,
              origin_id: Optional[str] = None,
              principal_name: Optional[str] = None,
              subject_types: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    Use this data source to access information about an existing users within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    user = azuredevops.get_users(principal_name="contoso-user@contoso.onmicrosoft.com")
    all_users = azuredevops.get_users()
    all_from_origin = azuredevops.get_users(origin="aad")
    all_from_subject_types = azuredevops.get_users(subject_types=[
        "aad",
        "msa",
    ])
    all_from_origin_id = azuredevops.get_users(origin="aad",
        origin_id="a7ead982-8438-4cd2-b9e3-c3aa51a7b675")
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-5.1)


    :param str origin: The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
    :param str origin_id: The unique identifier from the system of origin.
    :param str principal_name: The PrincipalName of this graph member from the source provider.
    :param Sequence[str] subject_types: A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
    """
    __args__ = dict()
    __args__['origin'] = origin
    __args__['originId'] = origin_id
    __args__['principalName'] = principal_name
    __args__['subjectTypes'] = subject_types
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        id=__ret__.id,
        origin=__ret__.origin,
        origin_id=__ret__.origin_id,
        principal_name=__ret__.principal_name,
        subject_types=__ret__.subject_types,
        users=__ret__.users)
