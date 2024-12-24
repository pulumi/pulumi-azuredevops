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

__all__ = ['ServiceendpointSnykArgs', 'ServiceendpointSnyk']

@pulumi.input_type
class ServiceendpointSnykArgs:
    def __init__(__self__, *,
                 api_token: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 server_url: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointSnyk resource.
        :param pulumi.Input[str] api_token: The API token of the Snyk Security Scan.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] server_url: The server URL of the Snyk Security Scan.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        pulumi.set(__self__, "api_token", api_token)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "server_url", server_url)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Input[str]:
        """
        The API token of the Snyk Security Scan.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_token", value)

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
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> pulumi.Input[str]:
        """
        The server URL of the Snyk Security Scan.
        """
        return pulumi.get(self, "server_url")

    @server_url.setter
    def server_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_url", value)

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
class _ServiceendpointSnykState:
    def __init__(__self__, *,
                 api_token: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 server_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointSnyk resources.
        :param pulumi.Input[str] api_token: The API token of the Snyk Security Scan.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] server_url: The server URL of the Snyk Security Scan.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if server_url is not None:
            pulumi.set(__self__, "server_url", server_url)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[str]]:
        """
        The API token of the Snyk Security Scan.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_token", value)

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
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> Optional[pulumi.Input[str]]:
        """
        The server URL of the Snyk Security Scan.
        """
        return pulumi.get(self, "server_url")

    @server_url.setter
    def server_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_url", value)

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


class ServiceendpointSnyk(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 server_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Snyk Security Scan service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Snyk Security Scan](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan)

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
        example_serviceendpoint_snyk = azuredevops.ServiceendpointSnyk("example",
            project_id=example.id,
            server_url="https://snyk.io/",
            api_token="00000000-0000-0000-0000-000000000000",
            service_endpoint_name="Example Snyk",
            description="Managed by Pulumi")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint Snyk can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: The API token of the Snyk Security Scan.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] server_url: The server URL of the Snyk Security Scan.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointSnykArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Snyk Security Scan service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Snyk Security Scan](https://marketplace.visualstudio.com/items?itemName=Snyk.snyk-security-scan)

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
        example_serviceendpoint_snyk = azuredevops.ServiceendpointSnyk("example",
            project_id=example.id,
            server_url="https://snyk.io/",
            api_token="00000000-0000-0000-0000-000000000000",
            service_endpoint_name="Example Snyk",
            description="Managed by Pulumi")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint Snyk can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointSnykArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointSnykArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 server_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointSnykArgs.__new__(ServiceendpointSnykArgs)

            if api_token is None and not opts.urn:
                raise TypeError("Missing required property 'api_token'")
            __props__.__dict__["api_token"] = None if api_token is None else pulumi.Output.secret(api_token)
            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if server_url is None and not opts.urn:
                raise TypeError("Missing required property 'server_url'")
            __props__.__dict__["server_url"] = server_url
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["apiToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceendpointSnyk, __self__).__init__(
            'azuredevops:index/serviceendpointSnyk:ServiceendpointSnyk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_token: Optional[pulumi.Input[str]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            server_url: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointSnyk':
        """
        Get an existing ServiceendpointSnyk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: The API token of the Snyk Security Scan.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] server_url: The server URL of the Snyk Security Scan.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointSnykState.__new__(_ServiceendpointSnykState)

        __props__.__dict__["api_token"] = api_token
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["server_url"] = server_url
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return ServiceendpointSnyk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[str]:
        """
        The API token of the Snyk Security Scan.
        """
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

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
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> pulumi.Output[str]:
        """
        The server URL of the Snyk Security Scan.
        """
        return pulumi.get(self, "server_url")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")
