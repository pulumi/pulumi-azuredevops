# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceEndpointNpmArgs', 'ServiceEndpointNpm']

@pulumi.input_type
class ServiceEndpointNpmArgs:
    def __init__(__self__, *,
                 access_token: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceEndpointNpm resource.
        :param pulumi.Input[str] access_token: The access token for npm registry.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the npm registry to connect with.
        :param pulumi.Input[str] description: The Service Endpoint description.
        """
        pulumi.set(__self__, "access_token", access_token)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "url", url)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Input[str]:
        """
        The access token for npm registry.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_token", value)

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
        URL of the npm registry to connect with.
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
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ServiceEndpointNpmState:
    def __init__(__self__, *,
                 access_token: Optional[pulumi.Input[str]] = None,
                 access_token_hash: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointNpm resources.
        :param pulumi.Input[str] access_token: The access token for npm registry.
        :param pulumi.Input[str] access_token_hash: A bcrypted hash of the attribute 'access_token'
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the npm registry to connect with.
        """
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if access_token_hash is not None:
            pulumi.set(__self__, "access_token_hash", access_token_hash)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        The access token for npm registry.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter(name="accessTokenHash")
    def access_token_hash(self) -> Optional[pulumi.Input[str]]:
        """
        A bcrypted hash of the attribute 'access_token'
        """
        return pulumi.get(self, "access_token_hash")

    @access_token_hash.setter
    def access_token_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token_hash", value)

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
        """
        The Service Endpoint description.
        """
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
        URL of the npm registry to connect with.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServiceEndpointNpm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a npm service endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_npm = azuredevops.ServiceEndpointNpm("exampleServiceEndpointNpm",
            project_id=example_project.id,
            service_endpoint_name="Example npm",
            url="https://registry.npmjs.org",
            access_token="00000000-0000-0000-0000-000000000000",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
        - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        - [npm User Token](https://docs.npmjs.com/about-access-tokens)

        ## Import

        Azure DevOps Service Endpoint npm can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
         $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: The access token for npm registry.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the npm registry to connect with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointNpmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a npm service endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_npm = azuredevops.ServiceEndpointNpm("exampleServiceEndpointNpm",
            project_id=example_project.id,
            service_endpoint_name="Example npm",
            url="https://registry.npmjs.org",
            access_token="00000000-0000-0000-0000-000000000000",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
        - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        - [npm User Token](https://docs.npmjs.com/about-access-tokens)

        ## Import

        Azure DevOps Service Endpoint npm can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
         $ pulumi import azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointNpmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointNpmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceEndpointNpmArgs.__new__(ServiceEndpointNpmArgs)

            if access_token is None and not opts.urn:
                raise TypeError("Missing required property 'access_token'")
            __props__.__dict__["access_token"] = access_token
            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["access_token_hash"] = None
        super(ServiceEndpointNpm, __self__).__init__(
            'azuredevops:index/serviceEndpointNpm:ServiceEndpointNpm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_token: Optional[pulumi.Input[str]] = None,
            access_token_hash: Optional[pulumi.Input[str]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceEndpointNpm':
        """
        Get an existing ServiceEndpointNpm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: The access token for npm registry.
        :param pulumi.Input[str] access_token_hash: A bcrypted hash of the attribute 'access_token'
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: URL of the npm registry to connect with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointNpmState.__new__(_ServiceEndpointNpmState)

        __props__.__dict__["access_token"] = access_token
        __props__.__dict__["access_token_hash"] = access_token_hash
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceEndpointNpm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Output[str]:
        """
        The access token for npm registry.
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="accessTokenHash")
    def access_token_hash(self) -> pulumi.Output[str]:
        """
        A bcrypted hash of the attribute 'access_token'
        """
        return pulumi.get(self, "access_token_hash")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "description")

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
        URL of the npm registry to connect with.
        """
        return pulumi.get(self, "url")

