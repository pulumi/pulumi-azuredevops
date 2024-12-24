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

__all__ = ['ServiceendpointAzureServiceBusArgs', 'ServiceendpointAzureServiceBus']

@pulumi.input_type
class ServiceendpointAzureServiceBusArgs:
    def __init__(__self__, *,
                 connection_string: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 queue_name: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointAzureServiceBus resource.
        :param pulumi.Input[str] connection_string: The  Azure Service Bus Connection string.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] queue_name: The Azure Service Bus Queue Name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "queue_name", queue_name)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Input[str]:
        """
        The  Azure Service Bus Connection string.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Input[str]:
        """
        The Azure Service Bus Queue Name.
        """
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ServiceendpointAzureServiceBusState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 queue_name: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointAzureServiceBus resources.
        :param pulumi.Input[str] connection_string: The  Azure Service Bus Connection string.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] queue_name: The Azure Service Bus Queue Name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if queue_name is not None:
            pulumi.set(__self__, "queue_name", queue_name)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The  Azure Service Bus Connection string.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Service Bus Queue Name.
        """
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)


class ServiceendpointAzureServiceBus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 queue_name: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Azure Service Bus endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_serviceendpoint_azure_service_bus = azuredevops.ServiceendpointAzureServiceBus("example",
            project_id=example.id,
            service_endpoint_name="Example Azure Service Bus",
            queue_name="queue",
            connection_string="connection string",
            description="Managed by Pulumi")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Azure Service Bus Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceendpointAzureServiceBus:ServiceendpointAzureServiceBus example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string: The  Azure Service Bus Connection string.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] queue_name: The Azure Service Bus Queue Name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointAzureServiceBusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure Service Bus endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Pulumi")
        example_serviceendpoint_azure_service_bus = azuredevops.ServiceendpointAzureServiceBus("example",
            project_id=example.id,
            service_endpoint_name="Example Azure Service Bus",
            queue_name="queue",
            connection_string="connection string",
            description="Managed by Pulumi")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Azure Service Bus Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceendpointAzureServiceBus:ServiceendpointAzureServiceBus example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointAzureServiceBusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointAzureServiceBusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 queue_name: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointAzureServiceBusArgs.__new__(ServiceendpointAzureServiceBusArgs)

            __props__.__dict__["authorization"] = authorization
            if connection_string is None and not opts.urn:
                raise TypeError("Missing required property 'connection_string'")
            __props__.__dict__["connection_string"] = None if connection_string is None else pulumi.Output.secret(connection_string)
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if queue_name is None and not opts.urn:
                raise TypeError("Missing required property 'queue_name'")
            __props__.__dict__["queue_name"] = queue_name
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["connectionString"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceendpointAzureServiceBus, __self__).__init__(
            'azuredevops:index/serviceendpointAzureServiceBus:ServiceendpointAzureServiceBus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            queue_name: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointAzureServiceBus':
        """
        Get an existing ServiceendpointAzureServiceBus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string: The  Azure Service Bus Connection string.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] queue_name: The Azure Service Bus Queue Name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointAzureServiceBusState.__new__(_ServiceendpointAzureServiceBusState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["connection_string"] = connection_string
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["queue_name"] = queue_name
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return ServiceendpointAzureServiceBus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        The  Azure Service Bus Connection string.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Output[str]:
        """
        The Azure Service Bus Queue Name.
        """
        return pulumi.get(self, "queue_name")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

