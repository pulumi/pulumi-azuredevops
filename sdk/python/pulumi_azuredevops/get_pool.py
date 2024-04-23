# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetPoolResult',
    'AwaitableGetPoolResult',
    'get_pool',
    'get_pool_output',
]

@pulumi.output_type
class GetPoolResult:
    """
    A collection of values returned by getPool.
    """
    def __init__(__self__, auto_provision=None, auto_update=None, id=None, name=None, pool_type=None):
        if auto_provision and not isinstance(auto_provision, bool):
            raise TypeError("Expected argument 'auto_provision' to be a bool")
        pulumi.set(__self__, "auto_provision", auto_provision)
        if auto_update and not isinstance(auto_update, bool):
            raise TypeError("Expected argument 'auto_update' to be a bool")
        pulumi.set(__self__, "auto_update", auto_update)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pool_type and not isinstance(pool_type, str):
            raise TypeError("Expected argument 'pool_type' to be a str")
        pulumi.set(__self__, "pool_type", pool_type)

    @property
    @pulumi.getter(name="autoProvision")
    def auto_provision(self) -> bool:
        return pulumi.get(self, "auto_provision")

    @property
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> bool:
        return pulumi.get(self, "auto_update")

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
    @pulumi.getter(name="poolType")
    def pool_type(self) -> str:
        return pulumi.get(self, "pool_type")


class AwaitableGetPoolResult(GetPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoolResult(
            auto_provision=self.auto_provision,
            auto_update=self.auto_update,
            id=self.id,
            name=self.name,
            pool_type=self.pool_type)


def get_pool(name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoolResult:
    """
    Use this data source to access information about an existing Agent Pool within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_pool(name="Example Agent Pool")
    pulumi.export("name", example.name)
    pulumi.export("poolType", example.pool_type)
    pulumi.export("autoProvision", example.auto_provision)
    pulumi.export("autoUpdate", example.auto_update)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)


    :param str name: Name of the Agent Pool.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getPool:getPool', __args__, opts=opts, typ=GetPoolResult).value

    return AwaitableGetPoolResult(
        auto_provision=pulumi.get(__ret__, 'auto_provision'),
        auto_update=pulumi.get(__ret__, 'auto_update'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        pool_type=pulumi.get(__ret__, 'pool_type'))


@_utilities.lift_output_func(get_pool)
def get_pool_output(name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPoolResult]:
    """
    Use this data source to access information about an existing Agent Pool within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_pool(name="Example Agent Pool")
    pulumi.export("name", example.name)
    pulumi.export("poolType", example.pool_type)
    pulumi.export("autoProvision", example.auto_provision)
    pulumi.export("autoUpdate", example.auto_update)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)


    :param str name: Name of the Agent Pool.
    """
    ...
