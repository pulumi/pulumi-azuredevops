# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CheckExclusiveLockArgs', 'CheckExclusiveLock']

@pulumi.input_type
class CheckExclusiveLockArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 target_resource_id: pulumi.Input[str],
                 target_resource_type: pulumi.Input[str],
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a CheckExclusiveLock resource.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Exclusive Lock Check to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "target_resource_id", target_resource_id)
        pulumi.set(__self__, "target_resource_type", target_resource_type)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project ID. Changing this forces a new Exclusive Lock Check to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_id", value)

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> pulumi.Input[str]:
        """
        The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        """
        return pulumi.get(self, "target_resource_type")

    @target_resource_type.setter
    def target_resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_type", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _CheckExclusiveLockState:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[str]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering CheckExclusiveLock resources.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Exclusive Lock Check to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if target_resource_id is not None:
            pulumi.set(__self__, "target_resource_id", target_resource_id)
        if target_resource_type is not None:
            pulumi.set(__self__, "target_resource_type", target_resource_type)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID. Changing this forces a new Exclusive Lock Check to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_resource_id", value)

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        """
        return pulumi.get(self, "target_resource_type")

    @target_resource_type.setter
    def target_resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_resource_type", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


class CheckExclusiveLock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a Exclusive Lock Check.

        Adding an exclusive lock will only allow a single stage to utilize this resource at a time. If multiple stages are waiting on the lock, only the latest will run. All others will be canceled.

        ## Example Usage
        ### Add Exclusive Lock to an environment

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_service_endpoint_generic = azuredevops.ServiceEndpointGeneric("exampleServiceEndpointGeneric",
            project_id=example_project.id,
            server_url="https://some-server.example.com",
            username="username",
            password="password",
            service_endpoint_name="Example Generic",
            description="Managed by Terraform")
        example_check_exclusive_lock = azuredevops.CheckExclusiveLock("exampleCheckExclusiveLock",
            project_id=example_project.id,
            target_resource_id=example_service_endpoint_generic.id,
            target_resource_type="endpoint",
            timeout=43200)
        ```
        ### Protect an environment

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_environment = azuredevops.Environment("exampleEnvironment", project_id=example_project.id)
        example_check_exclusive_lock = azuredevops.CheckExclusiveLock("exampleCheckExclusiveLock",
            project_id=example_project.id,
            target_resource_id=example_environment.id,
            target_resource_type="environment",
            timeout=43200)
        ```

        ## Import

        Importing this resource is not supported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Exclusive Lock Check to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CheckExclusiveLockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Exclusive Lock Check.

        Adding an exclusive lock will only allow a single stage to utilize this resource at a time. If multiple stages are waiting on the lock, only the latest will run. All others will be canceled.

        ## Example Usage
        ### Add Exclusive Lock to an environment

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_service_endpoint_generic = azuredevops.ServiceEndpointGeneric("exampleServiceEndpointGeneric",
            project_id=example_project.id,
            server_url="https://some-server.example.com",
            username="username",
            password="password",
            service_endpoint_name="Example Generic",
            description="Managed by Terraform")
        example_check_exclusive_lock = azuredevops.CheckExclusiveLock("exampleCheckExclusiveLock",
            project_id=example_project.id,
            target_resource_id=example_service_endpoint_generic.id,
            target_resource_type="endpoint",
            timeout=43200)
        ```
        ### Protect an environment

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_environment = azuredevops.Environment("exampleEnvironment", project_id=example_project.id)
        example_check_exclusive_lock = azuredevops.CheckExclusiveLock("exampleCheckExclusiveLock",
            project_id=example_project.id,
            target_resource_id=example_environment.id,
            target_resource_type="environment",
            timeout=43200)
        ```

        ## Import

        Importing this resource is not supported.

        :param str resource_name: The name of the resource.
        :param CheckExclusiveLockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CheckExclusiveLockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CheckExclusiveLockArgs.__new__(CheckExclusiveLockArgs)

            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if target_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__.__dict__["target_resource_id"] = target_resource_id
            if target_resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_type'")
            __props__.__dict__["target_resource_type"] = target_resource_type
            __props__.__dict__["timeout"] = timeout
        super(CheckExclusiveLock, __self__).__init__(
            'azuredevops:index/checkExclusiveLock:CheckExclusiveLock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None,
            target_resource_type: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None) -> 'CheckExclusiveLock':
        """
        Get an existing CheckExclusiveLock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Exclusive Lock Check to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CheckExclusiveLockState.__new__(_CheckExclusiveLockState)

        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["target_resource_id"] = target_resource_id
        __props__.__dict__["target_resource_type"] = target_resource_type
        __props__.__dict__["timeout"] = timeout
        return CheckExclusiveLock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID. Changing this forces a new Exclusive Lock Check to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource being protected by the check. Changing this forces a new Exclusive Lock to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> pulumi.Output[str]:
        """
        The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Exclusive Lock to be created.
        """
        return pulumi.get(self, "target_resource_type")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The timeout in minutes for the exclusive lock. Defaults to `43200`.
        """
        return pulumi.get(self, "timeout")
