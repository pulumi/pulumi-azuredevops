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

__all__ = [
    'GetIdentityUsersResult',
    'AwaitableGetIdentityUsersResult',
    'get_identity_users',
    'get_identity_users_output',
]

@pulumi.output_type
class GetIdentityUsersResult:
    """
    A collection of values returned by getIdentityUsers.
    """
    def __init__(__self__, descriptor=None, id=None, name=None, search_filter=None):
        if descriptor and not isinstance(descriptor, str):
            raise TypeError("Expected argument 'descriptor' to be a str")
        pulumi.set(__self__, "descriptor", descriptor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if search_filter and not isinstance(search_filter, str):
            raise TypeError("Expected argument 'search_filter' to be a str")
        pulumi.set(__self__, "search_filter", search_filter)

    @property
    @pulumi.getter
    def descriptor(self) -> str:
        """
        The descriptor of the user.
        """
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="searchFilter")
    def search_filter(self) -> Optional[str]:
        return pulumi.get(self, "search_filter")


class AwaitableGetIdentityUsersResult(GetIdentityUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdentityUsersResult(
            descriptor=self.descriptor,
            id=self.id,
            name=self.name,
            search_filter=self.search_filter)


def get_identity_users(name: Optional[str] = None,
                       search_filter: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdentityUsersResult:
    """
    Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).


    :param str name: The PrincipalName of this identity member from the source provider.
    :param str search_filter: The type of search to perform. Possible values are: `AccountName`, `DisplayName`, and `MailAddress`. Default is `General`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['searchFilter'] = search_filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getIdentityUsers:getIdentityUsers', __args__, opts=opts, typ=GetIdentityUsersResult).value

    return AwaitableGetIdentityUsersResult(
        descriptor=pulumi.get(__ret__, 'descriptor'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        search_filter=pulumi.get(__ret__, 'search_filter'))
def get_identity_users_output(name: Optional[pulumi.Input[str]] = None,
                              search_filter: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIdentityUsersResult]:
    """
    Use this data source to access information about an existing users within Azure DevOps On-Premise(Azure DevOps Server).


    :param str name: The PrincipalName of this identity member from the source provider.
    :param str search_filter: The type of search to perform. Possible values are: `AccountName`, `DisplayName`, and `MailAddress`. Default is `General`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['searchFilter'] = search_filter
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getIdentityUsers:getIdentityUsers', __args__, opts=opts, typ=GetIdentityUsersResult)
    return __ret__.apply(lambda __response__: GetIdentityUsersResult(
        descriptor=pulumi.get(__response__, 'descriptor'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        search_filter=pulumi.get(__response__, 'search_filter')))
