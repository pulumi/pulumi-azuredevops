# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceendpointNexusArgs', 'ServiceendpointNexus']

@pulumi.input_type
class ServiceendpointNexusArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 username: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointNexus resource.
        :param pulumi.Input[str] password: The Service Endpoint password to authenticate at the Nexus IQ Instance.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] url: The Service Endpoint url.
        :param pulumi.Input[str] username: The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "url", url)
        pulumi.set(__self__, "username", username)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The Service Endpoint password to authenticate at the Nexus IQ Instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[str]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The Service Endpoint url.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

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
class _ServiceendpointNexusState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointNexus resources.
        :param pulumi.Input[str] password: The Service Endpoint password to authenticate at the Nexus IQ Instance.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] url: The Service Endpoint url.
        :param pulumi.Input[str] username: The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            pulumi.set(__self__, "username", username)

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
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint password to authenticate at the Nexus IQ Instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint url.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class ServiceendpointNexus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Nexus IQ service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Nexus IQ instance.
        Nexus IQ is not supported by default, to manage a nexus service connection resource, it is necessary to install the [Nexus Extension](https://marketplace.visualstudio.com/items?itemName=SonatypeIntegrations.nexus-iq-azure-extension) in Azure DevOps.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_nexus = azuredevops.ServiceendpointNexus("example",
            project_id=example.id,
            service_endpoint_name="nexus-example",
            description="Service Endpoint for 'Nexus IQ' (Managed by Terraform)",
            url="https://example.com",
            username="username",
            password="password")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Service Connection Nexus can be imported using the `projectId/id` or or `projectName/id`, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointNexus:ServiceendpointNexus example projectName/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] password: The Service Endpoint password to authenticate at the Nexus IQ Instance.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] url: The Service Endpoint url.
        :param pulumi.Input[str] username: The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointNexusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Nexus IQ service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Nexus IQ instance.
        Nexus IQ is not supported by default, to manage a nexus service connection resource, it is necessary to install the [Nexus Extension](https://marketplace.visualstudio.com/items?itemName=SonatypeIntegrations.nexus-iq-azure-extension) in Azure DevOps.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_nexus = azuredevops.ServiceendpointNexus("example",
            project_id=example.id,
            service_endpoint_name="nexus-example",
            description="Service Endpoint for 'Nexus IQ' (Managed by Terraform)",
            url="https://example.com",
            username="username",
            password="password")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Service Connection Nexus can be imported using the `projectId/id` or or `projectName/id`, e.g.

        ```sh
        $ pulumi import azuredevops:index/serviceendpointNexus:ServiceendpointNexus example projectName/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointNexusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointNexusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointNexusArgs.__new__(ServiceendpointNexusArgs)

            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceendpointNexus, __self__).__init__(
            'azuredevops:index/serviceendpointNexus:ServiceendpointNexus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointNexus':
        """
        Get an existing ServiceendpointNexus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] password: The Service Endpoint password to authenticate at the Nexus IQ Instance.
        :param pulumi.Input[str] project_id: The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        :param pulumi.Input[str] url: The Service Endpoint url.
        :param pulumi.Input[str] username: The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointNexusState.__new__(_ServiceendpointNexusState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["password"] = password
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        __props__.__dict__["username"] = username
        return ServiceendpointNexus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The Service Endpoint password to authenticate at the Nexus IQ Instance.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project. Changing this forces a new Service Connection Nexus to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The Service Endpoint url.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The Service Endpoint username to authenticate at the Nexus IQ Instance.
        """
        return pulumi.get(self, "username")

