# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetPoolsResult',
    'AwaitableGetPoolsResult',
    'get_pools',
    'get_pools_output',
]

@pulumi.output_type
class GetPoolsResult:
    """
    A collection of values returned by getPools.
    """
    def __init__(__self__, agent_pools=None, id=None):
        if agent_pools and not isinstance(agent_pools, list):
            raise TypeError("Expected argument 'agent_pools' to be a list")
        pulumi.set(__self__, "agent_pools", agent_pools)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="agentPools")
    def agent_pools(self) -> Sequence['outputs.GetPoolsAgentPoolResult']:
        """
        A list of existing agent pools in your Azure DevOps Organization with the following details about every agent pool:
        """
        return pulumi.get(self, "agent_pools")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetPoolsResult(GetPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoolsResult(
            agent_pools=self.agent_pools,
            id=self.id)


def get_pools(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPoolsResult:
    """
    Use this data source to access information about existing Agent Pools within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_pools()
    pulumi.export("agentPoolName", [__item.name for __item in example.agent_pools])
    pulumi.export("autoProvision", [__item.auto_provision for __item in example.agent_pools])
    pulumi.export("autoUpdate", [__item.auto_update for __item in example.agent_pools])
    pulumi.export("poolType", [__item.pool_type for __item in example.agent_pools])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getPools:getPools', __args__, opts=opts, typ=GetPoolsResult).value

    return AwaitableGetPoolsResult(
        agent_pools=pulumi.get(__ret__, 'agent_pools'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_pools)
def get_pools_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPoolsResult]:
    """
    Use this data source to access information about existing Agent Pools within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_pools()
    pulumi.export("agentPoolName", [__item.name for __item in example.agent_pools])
    pulumi.export("autoProvision", [__item.auto_provision for __item in example.agent_pools])
    pulumi.export("autoUpdate", [__item.auto_update for __item in example.agent_pools])
    pulumi.export("poolType", [__item.pool_type for __item in example.agent_pools])
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-7.0)
    """
    ...
