# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceendpointOctopusdeployArgs', 'ServiceendpointOctopusdeploy']

@pulumi.input_type
class ServiceendpointOctopusdeployArgs:
    def __init__(__self__, *,
                 api_key: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_ssl_error: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ServiceendpointOctopusdeploy resource.
        :param pulumi.Input[str] api_key: API key to connect to Octopus Deploy.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: Octopus Server url.
        :param pulumi.Input[bool] ignore_ssl_error: Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        """
        ServiceendpointOctopusdeployArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_key=api_key,
            project_id=project_id,
            service_endpoint_name=service_endpoint_name,
            url=url,
            authorization=authorization,
            description=description,
            ignore_ssl_error=ignore_ssl_error,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_key: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             service_endpoint_name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             ignore_ssl_error: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if api_key is None and 'apiKey' in kwargs:
            api_key = kwargs['apiKey']
        if api_key is None:
            raise TypeError("Missing 'api_key' argument")
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")
        if service_endpoint_name is None and 'serviceEndpointName' in kwargs:
            service_endpoint_name = kwargs['serviceEndpointName']
        if service_endpoint_name is None:
            raise TypeError("Missing 'service_endpoint_name' argument")
        if url is None:
            raise TypeError("Missing 'url' argument")
        if ignore_ssl_error is None and 'ignoreSslError' in kwargs:
            ignore_ssl_error = kwargs['ignoreSslError']

        _setter("api_key", api_key)
        _setter("project_id", project_id)
        _setter("service_endpoint_name", service_endpoint_name)
        _setter("url", url)
        if authorization is not None:
            _setter("authorization", authorization)
        if description is not None:
            _setter("description", description)
        if ignore_ssl_error is not None:
            _setter("ignore_ssl_error", ignore_ssl_error)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Input[str]:
        """
        API key to connect to Octopus Deploy.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_key", value)

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
    def url(self) -> pulumi.Input[str]:
        """
        Octopus Server url.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

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

    @property
    @pulumi.getter(name="ignoreSslError")
    def ignore_ssl_error(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        """
        return pulumi.get(self, "ignore_ssl_error")

    @ignore_ssl_error.setter
    def ignore_ssl_error(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_ssl_error", value)


@pulumi.input_type
class _ServiceendpointOctopusdeployState:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_ssl_error: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointOctopusdeploy resources.
        :param pulumi.Input[str] api_key: API key to connect to Octopus Deploy.
        :param pulumi.Input[bool] ignore_ssl_error: Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: Octopus Server url.
        """
        _ServiceendpointOctopusdeployState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_key=api_key,
            authorization=authorization,
            description=description,
            ignore_ssl_error=ignore_ssl_error,
            project_id=project_id,
            service_endpoint_name=service_endpoint_name,
            url=url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_key: Optional[pulumi.Input[str]] = None,
             authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             ignore_ssl_error: Optional[pulumi.Input[bool]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             service_endpoint_name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if api_key is None and 'apiKey' in kwargs:
            api_key = kwargs['apiKey']
        if ignore_ssl_error is None and 'ignoreSslError' in kwargs:
            ignore_ssl_error = kwargs['ignoreSslError']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if service_endpoint_name is None and 'serviceEndpointName' in kwargs:
            service_endpoint_name = kwargs['serviceEndpointName']

        if api_key is not None:
            _setter("api_key", api_key)
        if authorization is not None:
            _setter("authorization", authorization)
        if description is not None:
            _setter("description", description)
        if ignore_ssl_error is not None:
            _setter("ignore_ssl_error", ignore_ssl_error)
        if project_id is not None:
            _setter("project_id", project_id)
        if service_endpoint_name is not None:
            _setter("service_endpoint_name", service_endpoint_name)
        if url is not None:
            _setter("url", url)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        """
        API key to connect to Octopus Deploy.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

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

    @property
    @pulumi.getter(name="ignoreSslError")
    def ignore_ssl_error(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        """
        return pulumi.get(self, "ignore_ssl_error")

    @ignore_ssl_error.setter
    def ignore_ssl_error(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_ssl_error", value)

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
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Octopus Server url.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServiceendpointOctopusdeploy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_ssl_error: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Octopus Deploy service endpoint within Azure DevOps. Using this service endpoint requires you to install [Octopus Deploy](https://marketplace.visualstudio.com/items?itemName=octopusdeploy.octopus-deploy-build-release-tasks).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_octopusdeploy = azuredevops.ServiceendpointOctopusdeploy("exampleServiceendpointOctopusdeploy",
            project_id=example_project.id,
            url="https://octopus.com",
            api_key="000000000000000000000000000000000000",
            service_endpoint_name="Example Octopus Deploy",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint Octopus Deploy can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: API key to connect to Octopus Deploy.
        :param pulumi.Input[bool] ignore_ssl_error: Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: Octopus Server url.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointOctopusdeployArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Octopus Deploy service endpoint within Azure DevOps. Using this service endpoint requires you to install [Octopus Deploy](https://marketplace.visualstudio.com/items?itemName=octopusdeploy.octopus-deploy-build-release-tasks).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_octopusdeploy = azuredevops.ServiceendpointOctopusdeploy("exampleServiceendpointOctopusdeploy",
            project_id=example_project.id,
            url="https://octopus.com",
            api_key="000000000000000000000000000000000000",
            service_endpoint_name="Example Octopus Deploy",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint Octopus Deploy can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointOctopusdeployArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointOctopusdeployArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ServiceendpointOctopusdeployArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_ssl_error: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointOctopusdeployArgs.__new__(ServiceendpointOctopusdeployArgs)

            if api_key is None and not opts.urn:
                raise TypeError("Missing required property 'api_key'")
            __props__.__dict__["api_key"] = api_key
            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            __props__.__dict__["ignore_ssl_error"] = ignore_ssl_error
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
        super(ServiceendpointOctopusdeploy, __self__).__init__(
            'azuredevops:index/serviceendpointOctopusdeploy:ServiceendpointOctopusdeploy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ignore_ssl_error: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointOctopusdeploy':
        """
        Get an existing ServiceendpointOctopusdeploy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: API key to connect to Octopus Deploy.
        :param pulumi.Input[bool] ignore_ssl_error: Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: Octopus Server url.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointOctopusdeployState.__new__(_ServiceendpointOctopusdeployState)

        __props__.__dict__["api_key"] = api_key
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["ignore_ssl_error"] = ignore_ssl_error
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceendpointOctopusdeploy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[str]:
        """
        API key to connect to Octopus Deploy.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ignoreSslError")
    def ignore_ssl_error(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to ignore SSL errors when connecting to the Octopus server from the agent. Default to `false`.
        """
        return pulumi.get(self, "ignore_ssl_error")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        Octopus Server url.
        """
        return pulumi.get(self, "url")

