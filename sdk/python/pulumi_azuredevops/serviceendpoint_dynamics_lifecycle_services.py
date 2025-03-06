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

__all__ = ['ServiceendpointDynamicsLifecycleServicesArgs', 'ServiceendpointDynamicsLifecycleServices']

@pulumi.input_type
class ServiceendpointDynamicsLifecycleServicesArgs:
    def __init__(__self__, *,
                 authorization_endpoint: pulumi.Input[str],
                 client_id: pulumi.Input[str],
                 lifecycle_services_api_endpoint: pulumi.Input[str],
                 password: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 username: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointDynamicsLifecycleServices resource.
        :param pulumi.Input[str] authorization_endpoint: The URL of the Authentication Endpoint.
        :param pulumi.Input[str] client_id: The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        :param pulumi.Input[str] lifecycle_services_api_endpoint: The URL of the Lifecycle Services API Endpoint.
        :param pulumi.Input[str] password: The Password for the Azure Active Directory account.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] username: The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "lifecycle_services_api_endpoint", lifecycle_services_api_endpoint)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "username", username)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> pulumi.Input[str]:
        """
        The URL of the Authentication Endpoint.
        """
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="lifecycleServicesApiEndpoint")
    def lifecycle_services_api_endpoint(self) -> pulumi.Input[str]:
        """
        The URL of the Lifecycle Services API Endpoint.
        """
        return pulumi.get(self, "lifecycle_services_api_endpoint")

    @lifecycle_services_api_endpoint.setter
    def lifecycle_services_api_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "lifecycle_services_api_endpoint", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The Password for the Azure Active Directory account.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

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
    def username(self) -> pulumi.Input[str]:
        """
        The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ServiceendpointDynamicsLifecycleServicesState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lifecycle_services_api_endpoint: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointDynamicsLifecycleServices resources.
        :param pulumi.Input[str] authorization_endpoint: The URL of the Authentication Endpoint.
        :param pulumi.Input[str] client_id: The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        :param pulumi.Input[str] lifecycle_services_api_endpoint: The URL of the Lifecycle Services API Endpoint.
        :param pulumi.Input[str] password: The Password for the Azure Active Directory account.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] username: The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if authorization_endpoint is not None:
            pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lifecycle_services_api_endpoint is not None:
            pulumi.set(__self__, "lifecycle_services_api_endpoint", lifecycle_services_api_endpoint)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
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
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the Authentication Endpoint.
        """
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lifecycleServicesApiEndpoint")
    def lifecycle_services_api_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the Lifecycle Services API Endpoint.
        """
        return pulumi.get(self, "lifecycle_services_api_endpoint")

    @lifecycle_services_api_endpoint.setter
    def lifecycle_services_api_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_services_api_endpoint", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The Password for the Azure Active Directory account.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

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
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class ServiceendpointDynamicsLifecycleServices(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lifecycle_services_api_endpoint: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Dynamics Lifecycle Services service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Dynamics Lifecycle Services](https://marketplace.visualstudio.com/items?itemName=Dyn365FinOps.dynamics365-finops-tools)

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
        example_serviceendpoint_dynamics_lifecycle_services = azuredevops.ServiceendpointDynamicsLifecycleServices("example",
            project_id=example.id,
            service_endpoint_name="Example Service connection",
            authorization_endpoint="https://login.microsoftonline.com/organization",
            lifecycle_services_api_endpoint="https://lcsapi.lcs.dynamics.com",
            client_id="00000000-0000-0000-0000-000000000000",
            username="username",
            password="password",
            description="Managed by Pulumi")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Dynamics Life Cycle Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceendpointDynamicsLifecycleServices:ServiceendpointDynamicsLifecycleServices example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_endpoint: The URL of the Authentication Endpoint.
        :param pulumi.Input[str] client_id: The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        :param pulumi.Input[str] lifecycle_services_api_endpoint: The URL of the Lifecycle Services API Endpoint.
        :param pulumi.Input[str] password: The Password for the Azure Active Directory account.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] username: The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointDynamicsLifecycleServicesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Dynamics Lifecycle Services service endpoint within Azure DevOps. Using this service endpoint requires you to install: [Dynamics Lifecycle Services](https://marketplace.visualstudio.com/items?itemName=Dyn365FinOps.dynamics365-finops-tools)

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
        example_serviceendpoint_dynamics_lifecycle_services = azuredevops.ServiceendpointDynamicsLifecycleServices("example",
            project_id=example.id,
            service_endpoint_name="Example Service connection",
            authorization_endpoint="https://login.microsoftonline.com/organization",
            lifecycle_services_api_endpoint="https://lcsapi.lcs.dynamics.com",
            client_id="00000000-0000-0000-0000-000000000000",
            username="username",
            password="password",
            description="Managed by Pulumi")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Dynamics Life Cycle Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceendpointDynamicsLifecycleServices:ServiceendpointDynamicsLifecycleServices example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointDynamicsLifecycleServicesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointDynamicsLifecycleServicesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lifecycle_services_api_endpoint: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointDynamicsLifecycleServicesArgs.__new__(ServiceendpointDynamicsLifecycleServicesArgs)

            if authorization_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_endpoint'")
            __props__.__dict__["authorization_endpoint"] = authorization_endpoint
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["description"] = description
            if lifecycle_services_api_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'lifecycle_services_api_endpoint'")
            __props__.__dict__["lifecycle_services_api_endpoint"] = lifecycle_services_api_endpoint
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["authorization"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceendpointDynamicsLifecycleServices, __self__).__init__(
            'azuredevops:index/serviceendpointDynamicsLifecycleServices:ServiceendpointDynamicsLifecycleServices',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            authorization_endpoint: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lifecycle_services_api_endpoint: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointDynamicsLifecycleServices':
        """
        Get an existing ServiceendpointDynamicsLifecycleServices resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_endpoint: The URL of the Authentication Endpoint.
        :param pulumi.Input[str] client_id: The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        :param pulumi.Input[str] lifecycle_services_api_endpoint: The URL of the Lifecycle Services API Endpoint.
        :param pulumi.Input[str] password: The Password for the Azure Active Directory account.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] username: The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointDynamicsLifecycleServicesState.__new__(_ServiceendpointDynamicsLifecycleServicesState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["authorization_endpoint"] = authorization_endpoint
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["description"] = description
        __props__.__dict__["lifecycle_services_api_endpoint"] = lifecycle_services_api_endpoint
        __props__.__dict__["password"] = password
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["username"] = username
        return ServiceendpointDynamicsLifecycleServices(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> pulumi.Output[str]:
        """
        The URL of the Authentication Endpoint.
        """
        return pulumi.get(self, "authorization_endpoint")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The client ID for a native application registration in Azure Active Directory with API permissions for Dynamics Lifecycle Services.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lifecycleServicesApiEndpoint")
    def lifecycle_services_api_endpoint(self) -> pulumi.Output[str]:
        """
        The URL of the Lifecycle Services API Endpoint.
        """
        return pulumi.get(self, "lifecycle_services_api_endpoint")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The Password for the Azure Active Directory account.
        """
        return pulumi.get(self, "password")

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
    def username(self) -> pulumi.Output[str]:
        """
        The E-mail address of user with sufficient permissions to interact with LCS asset library and environments.
        """
        return pulumi.get(self, "username")

