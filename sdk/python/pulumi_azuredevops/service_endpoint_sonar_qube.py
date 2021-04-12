# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['ServiceEndpointSonarQubeArgs', 'ServiceEndpointSonarQube']

@pulumi.input_type
class ServiceEndpointSonarQubeArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 token: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceEndpointSonarQube resource.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token: Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
        :param pulumi.Input[str] url: URL of the SonarQube server to connect with.
        :param pulumi.Input[str] description: The Service Endpoint description.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "url", url)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project ID or project name.
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
    def token(self) -> pulumi.Input[str]:
        """
        Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[str]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        URL of the SonarQube server to connect with.
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


class ServiceEndpointSonarQube(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a SonarQube service endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        serviceendpoint = azuredevops.ServiceEndpointSonarQube("serviceendpoint",
            project_id=project.id,
            service_endpoint_name="Sample SonarQube",
            url="https://sonarqube.my.com",
            token="0000000000000000000000000000000000000000",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        * [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/)

        ## Import

        Azure DevOps Service Endpoint SonarQube can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
         $ pulumi import azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token: Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
        :param pulumi.Input[str] url: URL of the SonarQube server to connect with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointSonarQubeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a SonarQube service endpoint within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        serviceendpoint = azuredevops.ServiceEndpointSonarQube("serviceendpoint",
            project_id=project.id,
            service_endpoint_name="Sample SonarQube",
            url="https://sonarqube.my.com",
            token="0000000000000000000000000000000000000000",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        * [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
        * [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/)

        ## Import

        Azure DevOps Service Endpoint SonarQube can be imported using the **projectID/serviceEndpointID**, e.g.

        ```sh
         $ pulumi import azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointSonarQubeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointSonarQubeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['authorization'] = authorization
            __props__['description'] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__['service_endpoint_name'] = service_endpoint_name
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__['token'] = token
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['token_hash'] = None
        super(ServiceEndpointSonarQube, __self__).__init__(
            'azuredevops:index/serviceEndpointSonarQube:ServiceEndpointSonarQube',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            token_hash: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceEndpointSonarQube':
        """
        Get an existing ServiceEndpointSonarQube resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The Service Endpoint description.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token: Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
        :param pulumi.Input[str] token_hash: A bcrypted hash of the attribute 'token'
        :param pulumi.Input[str] url: URL of the SonarQube server to connect with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorization"] = authorization
        __props__["description"] = description
        __props__["project_id"] = project_id
        __props__["service_endpoint_name"] = service_endpoint_name
        __props__["token"] = token
        __props__["token_hash"] = token_hash
        __props__["url"] = url
        return ServiceEndpointSonarQube(resource_name, opts=opts, __props__=__props__)

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
        The project ID or project name.
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
    def token(self) -> pulumi.Output[str]:
        """
        Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="tokenHash")
    def token_hash(self) -> pulumi.Output[str]:
        """
        A bcrypted hash of the attribute 'token'
        """
        return pulumi.get(self, "token_hash")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the SonarQube server to connect with.
        """
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

