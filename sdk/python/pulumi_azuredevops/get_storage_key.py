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
    'GetStorageKeyResult',
    'AwaitableGetStorageKeyResult',
    'get_storage_key',
    'get_storage_key_output',
]

@pulumi.output_type
class GetStorageKeyResult:
    """
    A collection of values returned by getStorageKey.
    """
    def __init__(__self__, descriptor=None, id=None, storage_key=None):
        if descriptor and not isinstance(descriptor, str):
            raise TypeError("Expected argument 'descriptor' to be a str")
        pulumi.set(__self__, "descriptor", descriptor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if storage_key and not isinstance(storage_key, str):
            raise TypeError("Expected argument 'storage_key' to be a str")
        pulumi.set(__self__, "storage_key", storage_key)

    @property
    @pulumi.getter
    def descriptor(self) -> str:
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="storageKey")
    def storage_key(self) -> str:
        """
        The Storage Key of the descriptor.
        """
        return pulumi.get(self, "storage_key")


class AwaitableGetStorageKeyResult(GetStorageKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageKeyResult(
            descriptor=self.descriptor,
            id=self.id,
            storage_key=self.storage_key)


def get_storage_key(descriptor: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageKeyResult:
    """
    Use this data source to access information about an existing Storage Key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_storage_key(descriptor="aad.000000000000000000000000000000000000")
    pulumi.export("id", example.id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.1 - Storage Key - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/graph/storage-keys/get?view=azure-devops-rest-7.1)


    :param str descriptor: The descriptor that will be resolved to a storage key.
    """
    __args__ = dict()
    __args__['descriptor'] = descriptor
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getStorageKey:getStorageKey', __args__, opts=opts, typ=GetStorageKeyResult).value

    return AwaitableGetStorageKeyResult(
        descriptor=pulumi.get(__ret__, 'descriptor'),
        id=pulumi.get(__ret__, 'id'),
        storage_key=pulumi.get(__ret__, 'storage_key'))
def get_storage_key_output(descriptor: Optional[pulumi.Input[str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStorageKeyResult]:
    """
    Use this data source to access information about an existing Storage Key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_storage_key(descriptor="aad.000000000000000000000000000000000000")
    pulumi.export("id", example.id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.1 - Storage Key - Get](https://learn.microsoft.com/en-us/rest/api/azure/devops/graph/storage-keys/get?view=azure-devops-rest-7.1)


    :param str descriptor: The descriptor that will be resolved to a storage key.
    """
    __args__ = dict()
    __args__['descriptor'] = descriptor
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getStorageKey:getStorageKey', __args__, opts=opts, typ=GetStorageKeyResult)
    return __ret__.apply(lambda __response__: GetStorageKeyResult(
        descriptor=pulumi.get(__response__, 'descriptor'),
        id=pulumi.get(__response__, 'id'),
        storage_key=pulumi.get(__response__, 'storage_key')))
