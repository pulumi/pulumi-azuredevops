# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['ServiceEndpointDockerRegistryArgs', 'ServiceEndpointDockerRegistry']

@pulumi.input_type
class ServiceEndpointDockerRegistryArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[_builtins.str],
                 service_endpoint_name: pulumi.Input[_builtins.str],
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_email: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_password: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_registry: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_username: Optional[pulumi.Input[_builtins.str]] = None,
                 registry_type: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a ServiceEndpointDockerRegistry resource.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The name you will use to refer to this service connection in task inputs.
        :param pulumi.Input[_builtins.str] docker_email: The email for Docker account user.
        :param pulumi.Input[_builtins.str] docker_password: The password for the account user identified above.
        :param pulumi.Input[_builtins.str] docker_registry: The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        :param pulumi.Input[_builtins.str] docker_username: The identifier of the Docker account user.
        :param pulumi.Input[_builtins.str] registry_type: Can be "DockerHub" or "Others" (Default "DockerHub")
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if docker_email is not None:
            pulumi.set(__self__, "docker_email", docker_email)
        if docker_password is not None:
            pulumi.set(__self__, "docker_password", docker_password)
        if docker_registry is not None:
            pulumi.set(__self__, "docker_registry", docker_registry)
        if docker_username is not None:
            pulumi.set(__self__, "docker_username", docker_username)
        if registry_type is not None:
            pulumi.set(__self__, "registry_type", registry_type)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[_builtins.str]:
        """
        The name you will use to refer to this service connection in task inputs.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "service_endpoint_name", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="dockerEmail")
    def docker_email(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The email for Docker account user.
        """
        return pulumi.get(self, "docker_email")

    @docker_email.setter
    def docker_email(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_email", value)

    @_builtins.property
    @pulumi.getter(name="dockerPassword")
    def docker_password(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The password for the account user identified above.
        """
        return pulumi.get(self, "docker_password")

    @docker_password.setter
    def docker_password(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_password", value)

    @_builtins.property
    @pulumi.getter(name="dockerRegistry")
    def docker_registry(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        """
        return pulumi.get(self, "docker_registry")

    @docker_registry.setter
    def docker_registry(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_registry", value)

    @_builtins.property
    @pulumi.getter(name="dockerUsername")
    def docker_username(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The identifier of the Docker account user.
        """
        return pulumi.get(self, "docker_username")

    @docker_username.setter
    def docker_username(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_username", value)

    @_builtins.property
    @pulumi.getter(name="registryType")
    def registry_type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Can be "DockerHub" or "Others" (Default "DockerHub")
        """
        return pulumi.get(self, "registry_type")

    @registry_type.setter
    def registry_type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "registry_type", value)


@pulumi.input_type
class _ServiceEndpointDockerRegistryState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_email: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_password: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_registry: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_username: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 registry_type: Optional[pulumi.Input[_builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointDockerRegistry resources.
        :param pulumi.Input[_builtins.str] docker_email: The email for Docker account user.
        :param pulumi.Input[_builtins.str] docker_password: The password for the account user identified above.
        :param pulumi.Input[_builtins.str] docker_registry: The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        :param pulumi.Input[_builtins.str] docker_username: The identifier of the Docker account user.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] registry_type: Can be "DockerHub" or "Others" (Default "DockerHub")
        :param pulumi.Input[_builtins.str] service_endpoint_name: The name you will use to refer to this service connection in task inputs.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if docker_email is not None:
            pulumi.set(__self__, "docker_email", docker_email)
        if docker_password is not None:
            pulumi.set(__self__, "docker_password", docker_password)
        if docker_registry is not None:
            pulumi.set(__self__, "docker_registry", docker_registry)
        if docker_username is not None:
            pulumi.set(__self__, "docker_username", docker_username)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if registry_type is not None:
            pulumi.set(__self__, "registry_type", registry_type)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)

    @_builtins.property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "authorization", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="dockerEmail")
    def docker_email(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The email for Docker account user.
        """
        return pulumi.get(self, "docker_email")

    @docker_email.setter
    def docker_email(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_email", value)

    @_builtins.property
    @pulumi.getter(name="dockerPassword")
    def docker_password(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The password for the account user identified above.
        """
        return pulumi.get(self, "docker_password")

    @docker_password.setter
    def docker_password(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_password", value)

    @_builtins.property
    @pulumi.getter(name="dockerRegistry")
    def docker_registry(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        """
        return pulumi.get(self, "docker_registry")

    @docker_registry.setter
    def docker_registry(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_registry", value)

    @_builtins.property
    @pulumi.getter(name="dockerUsername")
    def docker_username(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The identifier of the Docker account user.
        """
        return pulumi.get(self, "docker_username")

    @docker_username.setter
    def docker_username(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "docker_username", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter(name="registryType")
    def registry_type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Can be "DockerHub" or "Others" (Default "DockerHub")
        """
        return pulumi.get(self, "registry_type")

    @registry_type.setter
    def registry_type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "registry_type", value)

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name you will use to refer to this service connection in task inputs.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service_endpoint_name", value)


@pulumi.type_token("azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry")
class ServiceEndpointDockerRegistry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_email: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_password: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_registry: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_username: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 registry_type: Optional[pulumi.Input[_builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages a Docker Registry service endpoint within Azure DevOps.

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
        # dockerhub registry service connection
        example_service_endpoint_docker_registry = azuredevops.ServiceEndpointDockerRegistry("example",
            project_id=example.id,
            service_endpoint_name="Example Docker Hub",
            docker_username="example",
            docker_email="email@example.com",
            docker_password="12345",
            registry_type="DockerHub")
        # other docker registry service connection
        example_other = azuredevops.ServiceEndpointDockerRegistry("example-other",
            project_id=example.id,
            service_endpoint_name="Example Docker Registry",
            docker_registry="https://sample.azurecr.io/v1",
            docker_username="sample",
            docker_password="12345",
            registry_type="Others")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
        - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)

        ## Import

        Azure DevOps Docker Registry Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] docker_email: The email for Docker account user.
        :param pulumi.Input[_builtins.str] docker_password: The password for the account user identified above.
        :param pulumi.Input[_builtins.str] docker_registry: The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        :param pulumi.Input[_builtins.str] docker_username: The identifier of the Docker account user.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] registry_type: Can be "DockerHub" or "Others" (Default "DockerHub")
        :param pulumi.Input[_builtins.str] service_endpoint_name: The name you will use to refer to this service connection in task inputs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointDockerRegistryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Docker Registry service endpoint within Azure DevOps.

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
        # dockerhub registry service connection
        example_service_endpoint_docker_registry = azuredevops.ServiceEndpointDockerRegistry("example",
            project_id=example.id,
            service_endpoint_name="Example Docker Hub",
            docker_username="example",
            docker_email="email@example.com",
            docker_password="12345",
            registry_type="DockerHub")
        # other docker registry service connection
        example_other = azuredevops.ServiceEndpointDockerRegistry("example-other",
            project_id=example.id,
            service_endpoint_name="Example Docker Registry",
            docker_registry="https://sample.azurecr.io/v1",
            docker_username="sample",
            docker_password="12345",
            registry_type="Others")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
        - [Docker Registry Service Connection](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml#sep-docreg)

        ## Import

        Azure DevOps Docker Registry Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointDockerRegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointDockerRegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_email: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_password: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_registry: Optional[pulumi.Input[_builtins.str]] = None,
                 docker_username: Optional[pulumi.Input[_builtins.str]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 registry_type: Optional[pulumi.Input[_builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceEndpointDockerRegistryArgs.__new__(ServiceEndpointDockerRegistryArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["docker_email"] = docker_email
            __props__.__dict__["docker_password"] = None if docker_password is None else pulumi.Output.secret(docker_password)
            __props__.__dict__["docker_registry"] = docker_registry
            __props__.__dict__["docker_username"] = docker_username
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["registry_type"] = registry_type
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            __props__.__dict__["authorization"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["dockerPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceEndpointDockerRegistry, __self__).__init__(
            'azuredevops:index/serviceEndpointDockerRegistry:ServiceEndpointDockerRegistry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            docker_email: Optional[pulumi.Input[_builtins.str]] = None,
            docker_password: Optional[pulumi.Input[_builtins.str]] = None,
            docker_registry: Optional[pulumi.Input[_builtins.str]] = None,
            docker_username: Optional[pulumi.Input[_builtins.str]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None,
            registry_type: Optional[pulumi.Input[_builtins.str]] = None,
            service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None) -> 'ServiceEndpointDockerRegistry':
        """
        Get an existing ServiceEndpointDockerRegistry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] docker_email: The email for Docker account user.
        :param pulumi.Input[_builtins.str] docker_password: The password for the account user identified above.
        :param pulumi.Input[_builtins.str] docker_registry: The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        :param pulumi.Input[_builtins.str] docker_username: The identifier of the Docker account user.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] registry_type: Can be "DockerHub" or "Others" (Default "DockerHub")
        :param pulumi.Input[_builtins.str] service_endpoint_name: The name you will use to refer to this service connection in task inputs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointDockerRegistryState.__new__(_ServiceEndpointDockerRegistryState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["docker_email"] = docker_email
        __props__.__dict__["docker_password"] = docker_password
        __props__.__dict__["docker_registry"] = docker_registry
        __props__.__dict__["docker_username"] = docker_username
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["registry_type"] = registry_type
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return ServiceEndpointDockerRegistry(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, _builtins.str]]:
        return pulumi.get(self, "authorization")

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter(name="dockerEmail")
    def docker_email(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The email for Docker account user.
        """
        return pulumi.get(self, "docker_email")

    @_builtins.property
    @pulumi.getter(name="dockerPassword")
    def docker_password(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The password for the account user identified above.
        """
        return pulumi.get(self, "docker_password")

    @_builtins.property
    @pulumi.getter(name="dockerRegistry")
    def docker_registry(self) -> pulumi.Output[_builtins.str]:
        """
        The URL of the Docker registry. (Default: "https://index.docker.io/v1/")
        """
        return pulumi.get(self, "docker_registry")

    @_builtins.property
    @pulumi.getter(name="dockerUsername")
    def docker_username(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The identifier of the Docker account user.
        """
        return pulumi.get(self, "docker_username")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter(name="registryType")
    def registry_type(self) -> pulumi.Output[_builtins.str]:
        """
        Can be "DockerHub" or "Others" (Default "DockerHub")
        """
        return pulumi.get(self, "registry_type")

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[_builtins.str]:
        """
        The name you will use to refer to this service connection in task inputs.
        """
        return pulumi.get(self, "service_endpoint_name")

