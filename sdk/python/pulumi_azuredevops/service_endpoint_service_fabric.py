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
from . import outputs
from ._inputs import *

__all__ = ['ServiceEndpointServiceFabricArgs', 'ServiceEndpointServiceFabric']

@pulumi.input_type
class ServiceEndpointServiceFabricArgs:
    def __init__(__self__, *,
                 cluster_endpoint: pulumi.Input[_builtins.str],
                 project_id: pulumi.Input[_builtins.str],
                 service_endpoint_name: pulumi.Input[_builtins.str],
                 azure_active_directory: Optional[pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs']] = None,
                 certificate: Optional[pulumi.Input['ServiceEndpointServiceFabricCertificateArgs']] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 none: Optional[pulumi.Input['ServiceEndpointServiceFabricNoneArgs']] = None):
        """
        The set of arguments for constructing a ServiceEndpointServiceFabric resource.
        :param pulumi.Input[_builtins.str] cluster_endpoint: Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs'] azure_active_directory: An `azure_active_directory` block as documented below.
        :param pulumi.Input['ServiceEndpointServiceFabricCertificateArgs'] certificate: A `certificate` block as documented below.
        :param pulumi.Input['ServiceEndpointServiceFabricNoneArgs'] none: A `none` block as documented below.
        """
        pulumi.set(__self__, "cluster_endpoint", cluster_endpoint)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if azure_active_directory is not None:
            pulumi.set(__self__, "azure_active_directory", azure_active_directory)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if none is not None:
            pulumi.set(__self__, "none", none)

    @_builtins.property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> pulumi.Input[_builtins.str]:
        """
        Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        """
        return pulumi.get(self, "cluster_endpoint")

    @cluster_endpoint.setter
    def cluster_endpoint(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "cluster_endpoint", value)

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
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "service_endpoint_name", value)

    @_builtins.property
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> Optional[pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs']]:
        """
        An `azure_active_directory` block as documented below.
        """
        return pulumi.get(self, "azure_active_directory")

    @azure_active_directory.setter
    def azure_active_directory(self, value: Optional[pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs']]):
        pulumi.set(self, "azure_active_directory", value)

    @_builtins.property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input['ServiceEndpointServiceFabricCertificateArgs']]:
        """
        A `certificate` block as documented below.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input['ServiceEndpointServiceFabricCertificateArgs']]):
        pulumi.set(self, "certificate", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def none(self) -> Optional[pulumi.Input['ServiceEndpointServiceFabricNoneArgs']]:
        """
        A `none` block as documented below.
        """
        return pulumi.get(self, "none")

    @none.setter
    def none(self, value: Optional[pulumi.Input['ServiceEndpointServiceFabricNoneArgs']]):
        pulumi.set(self, "none", value)


@pulumi.input_type
class _ServiceEndpointServiceFabricState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 azure_active_directory: Optional[pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs']] = None,
                 certificate: Optional[pulumi.Input['ServiceEndpointServiceFabricCertificateArgs']] = None,
                 cluster_endpoint: Optional[pulumi.Input[_builtins.str]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 none: Optional[pulumi.Input['ServiceEndpointServiceFabricNoneArgs']] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointServiceFabric resources.
        :param pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs'] azure_active_directory: An `azure_active_directory` block as documented below.
        :param pulumi.Input['ServiceEndpointServiceFabricCertificateArgs'] certificate: A `certificate` block as documented below.
        :param pulumi.Input[_builtins.str] cluster_endpoint: Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        :param pulumi.Input['ServiceEndpointServiceFabricNoneArgs'] none: A `none` block as documented below.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if azure_active_directory is not None:
            pulumi.set(__self__, "azure_active_directory", azure_active_directory)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if cluster_endpoint is not None:
            pulumi.set(__self__, "cluster_endpoint", cluster_endpoint)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if none is not None:
            pulumi.set(__self__, "none", none)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
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
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> Optional[pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs']]:
        """
        An `azure_active_directory` block as documented below.
        """
        return pulumi.get(self, "azure_active_directory")

    @azure_active_directory.setter
    def azure_active_directory(self, value: Optional[pulumi.Input['ServiceEndpointServiceFabricAzureActiveDirectoryArgs']]):
        pulumi.set(self, "azure_active_directory", value)

    @_builtins.property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input['ServiceEndpointServiceFabricCertificateArgs']]:
        """
        A `certificate` block as documented below.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input['ServiceEndpointServiceFabricCertificateArgs']]):
        pulumi.set(self, "certificate", value)

    @_builtins.property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        """
        return pulumi.get(self, "cluster_endpoint")

    @cluster_endpoint.setter
    def cluster_endpoint(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "cluster_endpoint", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def none(self) -> Optional[pulumi.Input['ServiceEndpointServiceFabricNoneArgs']]:
        """
        A `none` block as documented below.
        """
        return pulumi.get(self, "none")

    @none.setter
    def none(self, value: Optional[pulumi.Input['ServiceEndpointServiceFabricNoneArgs']]):
        pulumi.set(self, "none", value)

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
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service_endpoint_name", value)


@pulumi.type_token("azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric")
class ServiceEndpointServiceFabric(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_active_directory: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricAzureActiveDirectoryArgs', 'ServiceEndpointServiceFabricAzureActiveDirectoryArgsDict']]] = None,
                 certificate: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricCertificateArgs', 'ServiceEndpointServiceFabricCertificateArgsDict']]] = None,
                 cluster_endpoint: Optional[pulumi.Input[_builtins.str]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 none: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricNoneArgs', 'ServiceEndpointServiceFabricNoneArgsDict']]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages a Service Fabric service endpoint within Azure DevOps.

        ## Example Usage

        ### Azure Active Directory Authentication

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            name="Sample Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        test = azuredevops.ServiceEndpointServiceFabric("test",
            project_id=project.id,
            service_endpoint_name="Sample Service Fabric",
            description="Managed by Pulumi",
            cluster_endpoint="tcp://test",
            azure_active_directory={
                "server_certificate_lookup": "Thumbprint",
                "server_certificate_thumbprint": "0000000000000000000000000000000000000000",
                "username": "username",
                "password": "password",
            })
        ```

        ### Windows Authentication

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            name="Sample Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        test = azuredevops.ServiceEndpointServiceFabric("test",
            project_id=project.id,
            service_endpoint_name="Sample Service Fabric",
            description="Managed by Pulumi",
            cluster_endpoint="tcp://test",
            none={
                "unsecured": False,
                "cluster_spn": "HTTP/www.contoso.com",
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Fabric Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceEndpointServiceFabricAzureActiveDirectoryArgs', 'ServiceEndpointServiceFabricAzureActiveDirectoryArgsDict']] azure_active_directory: An `azure_active_directory` block as documented below.
        :param pulumi.Input[Union['ServiceEndpointServiceFabricCertificateArgs', 'ServiceEndpointServiceFabricCertificateArgsDict']] certificate: A `certificate` block as documented below.
        :param pulumi.Input[_builtins.str] cluster_endpoint: Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        :param pulumi.Input[Union['ServiceEndpointServiceFabricNoneArgs', 'ServiceEndpointServiceFabricNoneArgsDict']] none: A `none` block as documented below.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointServiceFabricArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Service Fabric service endpoint within Azure DevOps.

        ## Example Usage

        ### Azure Active Directory Authentication

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            name="Sample Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        test = azuredevops.ServiceEndpointServiceFabric("test",
            project_id=project.id,
            service_endpoint_name="Sample Service Fabric",
            description="Managed by Pulumi",
            cluster_endpoint="tcp://test",
            azure_active_directory={
                "server_certificate_lookup": "Thumbprint",
                "server_certificate_thumbprint": "0000000000000000000000000000000000000000",
                "username": "username",
                "password": "password",
            })
        ```

        ### Windows Authentication

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            name="Sample Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        test = azuredevops.ServiceEndpointServiceFabric("test",
            project_id=project.id,
            service_endpoint_name="Sample Service Fabric",
            description="Managed by Pulumi",
            cluster_endpoint="tcp://test",
            none={
                "unsecured": False,
                "cluster_spn": "HTTP/www.contoso.com",
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Fabric Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointServiceFabricArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointServiceFabricArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_active_directory: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricAzureActiveDirectoryArgs', 'ServiceEndpointServiceFabricAzureActiveDirectoryArgsDict']]] = None,
                 certificate: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricCertificateArgs', 'ServiceEndpointServiceFabricCertificateArgsDict']]] = None,
                 cluster_endpoint: Optional[pulumi.Input[_builtins.str]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 none: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricNoneArgs', 'ServiceEndpointServiceFabricNoneArgsDict']]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceEndpointServiceFabricArgs.__new__(ServiceEndpointServiceFabricArgs)

            __props__.__dict__["azure_active_directory"] = azure_active_directory
            __props__.__dict__["certificate"] = certificate
            if cluster_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_endpoint'")
            __props__.__dict__["cluster_endpoint"] = cluster_endpoint
            __props__.__dict__["description"] = description
            __props__.__dict__["none"] = none
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            __props__.__dict__["authorization"] = None
        super(ServiceEndpointServiceFabric, __self__).__init__(
            'azuredevops:index/serviceEndpointServiceFabric:ServiceEndpointServiceFabric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            azure_active_directory: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricAzureActiveDirectoryArgs', 'ServiceEndpointServiceFabricAzureActiveDirectoryArgsDict']]] = None,
            certificate: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricCertificateArgs', 'ServiceEndpointServiceFabricCertificateArgsDict']]] = None,
            cluster_endpoint: Optional[pulumi.Input[_builtins.str]] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            none: Optional[pulumi.Input[Union['ServiceEndpointServiceFabricNoneArgs', 'ServiceEndpointServiceFabricNoneArgsDict']]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None,
            service_endpoint_name: Optional[pulumi.Input[_builtins.str]] = None) -> 'ServiceEndpointServiceFabric':
        """
        Get an existing ServiceEndpointServiceFabric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ServiceEndpointServiceFabricAzureActiveDirectoryArgs', 'ServiceEndpointServiceFabricAzureActiveDirectoryArgsDict']] azure_active_directory: An `azure_active_directory` block as documented below.
        :param pulumi.Input[Union['ServiceEndpointServiceFabricCertificateArgs', 'ServiceEndpointServiceFabricCertificateArgsDict']] certificate: A `certificate` block as documented below.
        :param pulumi.Input[_builtins.str] cluster_endpoint: Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        :param pulumi.Input[Union['ServiceEndpointServiceFabricNoneArgs', 'ServiceEndpointServiceFabricNoneArgsDict']] none: A `none` block as documented below.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project.
        :param pulumi.Input[_builtins.str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointServiceFabricState.__new__(_ServiceEndpointServiceFabricState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["azure_active_directory"] = azure_active_directory
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["cluster_endpoint"] = cluster_endpoint
        __props__.__dict__["description"] = description
        __props__.__dict__["none"] = none
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return ServiceEndpointServiceFabric(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, _builtins.str]]:
        return pulumi.get(self, "authorization")

    @_builtins.property
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> pulumi.Output[Optional['outputs.ServiceEndpointServiceFabricAzureActiveDirectory']]:
        """
        An `azure_active_directory` block as documented below.
        """
        return pulumi.get(self, "azure_active_directory")

    @_builtins.property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional['outputs.ServiceEndpointServiceFabricCertificate']]:
        """
        A `certificate` block as documented below.
        """
        return pulumi.get(self, "certificate")

    @_builtins.property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> pulumi.Output[_builtins.str]:
        """
        Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.
        """
        return pulumi.get(self, "cluster_endpoint")

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter
    def none(self) -> pulumi.Output[Optional['outputs.ServiceEndpointServiceFabricNone']]:
        """
        A `none` block as documented below.
        """
        return pulumi.get(self, "none")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[_builtins.str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

