# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PoolArgs', 'Pool']

@pulumi.input_type
class PoolArgs:
    def __init__(__self__, *,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Pool resource.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[bool] auto_update: Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        PoolArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auto_provision=auto_provision,
            auto_update=auto_update,
            name=name,
            pool_type=pool_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auto_provision: Optional[pulumi.Input[bool]] = None,
             auto_update: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             pool_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if auto_provision is None and 'autoProvision' in kwargs:
            auto_provision = kwargs['autoProvision']
        if auto_update is None and 'autoUpdate' in kwargs:
            auto_update = kwargs['autoUpdate']
        if pool_type is None and 'poolType' in kwargs:
            pool_type = kwargs['poolType']

        if auto_provision is not None:
            _setter("auto_provision", auto_provision)
        if auto_update is not None:
            _setter("auto_update", auto_update)
        if name is not None:
            _setter("name", name)
        if pool_type is not None:
            _setter("pool_type", pool_type)

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
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        """
        return pulumi.get(self, "auto_update")

    @auto_update.setter
    def auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_update", value)

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
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Pool resources.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[bool] auto_update: Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        _PoolState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auto_provision=auto_provision,
            auto_update=auto_update,
            name=name,
            pool_type=pool_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auto_provision: Optional[pulumi.Input[bool]] = None,
             auto_update: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             pool_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if auto_provision is None and 'autoProvision' in kwargs:
            auto_provision = kwargs['autoProvision']
        if auto_update is None and 'autoUpdate' in kwargs:
            auto_update = kwargs['autoUpdate']
        if pool_type is None and 'poolType' in kwargs:
            pool_type = kwargs['poolType']

        if auto_provision is not None:
            _setter("auto_provision", auto_provision)
        if auto_update is not None:
            _setter("auto_update", auto_update)
        if name is not None:
            _setter("name", name)
        if pool_type is not None:
            _setter("pool_type", pool_type)

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
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        """
        return pulumi.get(self, "auto_update")

    @auto_update.setter
    def auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_update", value)

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


class Pool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an agent pool within Azure DevOps.

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Agent Pools can be imported using the agent pool ID, e.g.

        ```sh
         $ pulumi import azuredevops:index/pool:Pool example 0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[bool] auto_update: Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
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

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/pools?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Agent Pools can be imported using the agent pool ID, e.g.

        ```sh
         $ pulumi import azuredevops:index/pool:Pool example 0
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            PoolArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_provision: Optional[pulumi.Input[bool]] = None,
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PoolArgs.__new__(PoolArgs)

            __props__.__dict__["auto_provision"] = auto_provision
            __props__.__dict__["auto_update"] = auto_update
            __props__.__dict__["name"] = name
            __props__.__dict__["pool_type"] = pool_type
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Agent/pool:Pool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Pool, __self__).__init__(
            'azuredevops:index/pool:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_provision: Optional[pulumi.Input[bool]] = None,
            auto_update: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_type: Optional[pulumi.Input[str]] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_provision: Specifies whether a queue should be automatically provisioned for each project collection. Defaults to `false`.
        :param pulumi.Input[bool] auto_update: Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] pool_type: Specifies whether the agent pool type is Automation or Deployment. Defaults to `automation`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PoolState.__new__(_PoolState)

        __props__.__dict__["auto_provision"] = auto_provision
        __props__.__dict__["auto_update"] = auto_update
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
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether or not agents within the pool should be automatically updated. Defaults to `true`.
        """
        return pulumi.get(self, "auto_update")

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

