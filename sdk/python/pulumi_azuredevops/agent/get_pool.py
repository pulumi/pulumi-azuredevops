# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetPoolResult:
    """
    A collection of values returned by getPool.
    """
    def __init__(__self__, auto_provision=None, id=None, name=None, pool_type=None):
        if auto_provision and not isinstance(auto_provision, bool):
            raise TypeError("Expected argument 'auto_provision' to be a bool")
        __self__.auto_provision = auto_provision
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if pool_type and not isinstance(pool_type, str):
            raise TypeError("Expected argument 'pool_type' to be a str")
        __self__.pool_type = pool_type


class AwaitableGetPoolResult(GetPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPoolResult(
            auto_provision=self.auto_provision,
            id=self.id,
            name=self.name,
            pool_type=self.pool_type)


def get_pool(name=None, opts=None):
    """
    Use this data source to access information about an existing Agent Pool within Azure DevOps.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    pool = azuredevops.Agent.get_pool(name="Sample Agent Pool")
    pulumi.export("name", pool.name)
    pulumi.export("poolType", pool.pool_type)
    pulumi.export("autoProvision", pool.auto_provision)
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Agent Pools - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools/get?view=azure-devops-rest-5.1)


    :param str name: Name of the Agent Pool.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:Agent/getPool:getPool', __args__, opts=opts).value

    return AwaitableGetPoolResult(
        auto_provision=__ret__.get('autoProvision'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        pool_type=__ret__.get('poolType'))
