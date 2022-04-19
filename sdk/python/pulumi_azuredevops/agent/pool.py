# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PoolArgs', 'Pool']

@pulumi.input_type
class PoolArgs:
    def __init__(__self__, *,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Pool resource.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        if auto_provision is not None:
            pulumi.set(__self__, "auto_provision", auto_provision)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pool_type is not None:
            pulumi.set(__self__, "pool_type", pool_type)

    @property
    @pulumi.getter(name="autoProvision")
    def auto_provision(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        """
        return pulumi.get(self, "auto_provision")

    @auto_provision.setter
    def auto_provision(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_provision", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the agent pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="poolType")
    def pool_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        return pulumi.get(self, "pool_type")

    @pool_type.setter
    def pool_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_type", value)


@pulumi.input_type
class _PoolState:
    def __init__(__self__, *,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Pool resources.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        if auto_provision is not None:
            pulumi.set(__self__, "auto_provision", auto_provision)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pool_type is not None:
            pulumi.set(__self__, "pool_type", pool_type)

    @property
    @pulumi.getter(name="autoProvision")
    def auto_provision(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        """
        return pulumi.get(self, "auto_provision")

    @auto_provision.setter
    def auto_provision(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_provision", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the agent pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="poolType")
    def pool_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        return pulumi.get(self, "pool_type")

    @pool_type.setter
    def pool_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_type", value)


warnings.warn("""azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool""", DeprecationWarning)


class Pool(pulumi.CustomResource):
    warnings.warn("""azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an agent pool within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Pool("example", auto_provision=False)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-6.0)

        ## Import

        Azure DevOps Agent Pools can be imported using the agent pool ID, e.g.

        ```sh
         $ pulumi import azuredevops:Agent/pool:Pool example 0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PoolArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an agent pool within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Pool("example", auto_provision=False)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-6.0)

        ## Import

        Azure DevOps Agent Pools can be imported using the agent pool ID, e.g.

        ```sh
         $ pulumi import azuredevops:Agent/pool:Pool example 0
        ```

        :param str resource_name: The name of the resource.
        :param PoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Pool is deprecated: azuredevops.agent.Pool has been deprecated in favor of azuredevops.Pool""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PoolArgs.__new__(PoolArgs)

            __props__.__dict__["auto_provision"] = auto_provision
            __props__.__dict__["name"] = name
            __props__.__dict__["pool_type"] = pool_type
        super(Pool, __self__).__init__(
            'azuredevops:Agent/pool:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_provision: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_type: Optional[pulumi.Input[str]] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PoolState.__new__(_PoolState)

        __props__.__dict__["auto_provision"] = auto_provision
        __props__.__dict__["name"] = name
        __props__.__dict__["pool_type"] = pool_type
        return Pool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoProvision")
    def auto_provision(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        """
        return pulumi.get(self, "auto_provision")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the agent pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolType")
    def pool_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        return pulumi.get(self, "pool_type")

