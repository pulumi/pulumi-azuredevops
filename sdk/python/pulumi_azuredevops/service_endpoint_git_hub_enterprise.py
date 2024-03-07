# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceEndpointGitHubEnterpriseArgs', 'ServiceEndpointGitHubEnterprise']

@pulumi.input_type
class ServiceEndpointGitHubEnterpriseArgs:
    def __init__(__self__, *,
                 auth_personal: pulumi.Input['ServiceEndpointGitHubEnterpriseAuthPersonalArgs'],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceEndpointGitHubEnterprise resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: GitHub Enterprise Server Url.
        """
        pulumi.set(__self__, "auth_personal", auth_personal)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "url", url)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="authPersonal")
    def auth_personal(self) -> pulumi.Input['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']:
        return pulumi.get(self, "auth_personal")

    @auth_personal.setter
    def auth_personal(self, value: pulumi.Input['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']):
        pulumi.set(self, "auth_personal", value)

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
        GitHub Enterprise Server Url.
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


@pulumi.input_type
class _ServiceEndpointGitHubEnterpriseState:
    def __init__(__self__, *,
                 auth_personal: Optional[pulumi.Input['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointGitHubEnterprise resources.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: GitHub Enterprise Server Url.
        """
        if auth_personal is not None:
            pulumi.set(__self__, "auth_personal", auth_personal)
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
    @pulumi.getter(name="authPersonal")
    def auth_personal(self) -> Optional[pulumi.Input['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']]:
        return pulumi.get(self, "auth_personal")

    @auth_personal.setter
    def auth_personal(self, value: Optional[pulumi.Input['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']]):
        pulumi.set(self, "auth_personal", value)

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
        GitHub Enterprise Server Url.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ServiceEndpointGitHubEnterprise(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_personal: Optional[pulumi.Input[pulumi.InputType['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a GitHub Enterprise Server service endpoint within Azure DevOps.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_git_hub_enterprise = azuredevops.ServiceEndpointGitHubEnterprise("exampleServiceEndpointGitHubEnterprise",
            project_id=example_project.id,
            service_endpoint_name="Example GitHub Enterprise",
            url="https://github.contoso.com",
            description="Managed by Terraform",
            auth_personal=azuredevops.ServiceEndpointGitHubEnterpriseAuthPersonalArgs(
                personal_access_token="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint GitHub Enterprise Server can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: GitHub Enterprise Server Url.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointGitHubEnterpriseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a GitHub Enterprise Server service endpoint within Azure DevOps.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_git_hub_enterprise = azuredevops.ServiceEndpointGitHubEnterprise("exampleServiceEndpointGitHubEnterprise",
            project_id=example_project.id,
            service_endpoint_name="Example GitHub Enterprise",
            url="https://github.contoso.com",
            description="Managed by Terraform",
            auth_personal=azuredevops.ServiceEndpointGitHubEnterpriseAuthPersonalArgs(
                personal_access_token="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint GitHub Enterprise Server can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointGitHubEnterpriseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointGitHubEnterpriseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_personal: Optional[pulumi.Input[pulumi.InputType['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceEndpointGitHubEnterpriseArgs.__new__(ServiceEndpointGitHubEnterpriseArgs)

            if auth_personal is None and not opts.urn:
                raise TypeError("Missing required property 'auth_personal'")
            __props__.__dict__["auth_personal"] = auth_personal
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
        super(ServiceEndpointGitHubEnterprise, __self__).__init__(
            'azuredevops:index/serviceEndpointGitHubEnterprise:ServiceEndpointGitHubEnterprise',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_personal: Optional[pulumi.Input[pulumi.InputType['ServiceEndpointGitHubEnterpriseAuthPersonalArgs']]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ServiceEndpointGitHubEnterprise':
        """
        Get an existing ServiceEndpointGitHubEnterprise resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] url: GitHub Enterprise Server Url.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointGitHubEnterpriseState.__new__(_ServiceEndpointGitHubEnterpriseState)

        __props__.__dict__["auth_personal"] = auth_personal
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["url"] = url
        return ServiceEndpointGitHubEnterprise(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authPersonal")
    def auth_personal(self) -> pulumi.Output['outputs.ServiceEndpointGitHubEnterpriseAuthPersonal']:
        return pulumi.get(self, "auth_personal")

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
        GitHub Enterprise Server Url.
        """
        return pulumi.get(self, "url")

