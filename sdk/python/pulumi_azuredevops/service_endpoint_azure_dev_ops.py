# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceEndpointAzureDevOpsArgs', 'ServiceEndpointAzureDevOps']

@pulumi.input_type
class ServiceEndpointAzureDevOpsArgs:
    def __init__(__self__, *,
                 org_url: pulumi.Input[str],
                 personal_access_token: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 release_api_url: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceEndpointAzureDevOps resource.
        :param pulumi.Input[str] org_url: The organization URL.
        :param pulumi.Input[str] personal_access_token: The Azure DevOps personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] release_api_url: The URL of the release API.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        pulumi.set(__self__, "org_url", org_url)
        pulumi.set(__self__, "personal_access_token", personal_access_token)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "release_api_url", release_api_url)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="orgUrl")
    def org_url(self) -> pulumi.Input[str]:
        """
        The organization URL.
        """
        return pulumi.get(self, "org_url")

    @org_url.setter
    def org_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_url", value)

    @property
    @pulumi.getter(name="personalAccessToken")
    def personal_access_token(self) -> pulumi.Input[str]:
        """
        The Azure DevOps personal access token.
        """
        return pulumi.get(self, "personal_access_token")

    @personal_access_token.setter
    def personal_access_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "personal_access_token", value)

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
    @pulumi.getter(name="releaseApiUrl")
    def release_api_url(self) -> pulumi.Input[str]:
        """
        The URL of the release API.
        """
        return pulumi.get(self, "release_api_url")

    @release_api_url.setter
    def release_api_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "release_api_url", value)

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
class _ServiceEndpointAzureDevOpsState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org_url: Optional[pulumi.Input[str]] = None,
                 personal_access_token: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 release_api_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointAzureDevOps resources.
        :param pulumi.Input[str] org_url: The organization URL.
        :param pulumi.Input[str] personal_access_token: The Azure DevOps personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] release_api_url: The URL of the release API.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if org_url is not None:
            pulumi.set(__self__, "org_url", org_url)
        if personal_access_token is not None:
            pulumi.set(__self__, "personal_access_token", personal_access_token)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if release_api_url is not None:
            pulumi.set(__self__, "release_api_url", release_api_url)
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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="orgUrl")
    def org_url(self) -> Optional[pulumi.Input[str]]:
        """
        The organization URL.
        """
        return pulumi.get(self, "org_url")

    @org_url.setter
    def org_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_url", value)

    @property
    @pulumi.getter(name="personalAccessToken")
    def personal_access_token(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure DevOps personal access token.
        """
        return pulumi.get(self, "personal_access_token")

    @personal_access_token.setter
    def personal_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "personal_access_token", value)

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
    @pulumi.getter(name="releaseApiUrl")
    def release_api_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the release API.
        """
        return pulumi.get(self, "release_api_url")

    @release_api_url.setter
    def release_api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "release_api_url", value)

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


class ServiceEndpointAzureDevOps(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org_url: Optional[pulumi.Input[str]] = None,
                 personal_access_token: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 release_api_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Azure DevOps service endpoint within Azure DevOps.

        > **Note** This resource is duplicate with `ServiceEndpointPipeline`,  will be removed in the future, use `ServiceEndpointPipeline` instead.

        > **Note** Prerequisite: Extension [Configurable Pipeline Runner](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines) has been installed for the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_azure_dev_ops = azuredevops.ServiceEndpointAzureDevOps("example",
            project_id=example.id,
            service_endpoint_name="Example Azure DevOps",
            org_url="https://dev.azure.com/testorganization",
            release_api_url="https://vsrm.dev.azure.com/testorganization",
            personal_access_token="0000000000000000000000000000000000000000000000000000",
            description="Managed by Terraform")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_url: The organization URL.
        :param pulumi.Input[str] personal_access_token: The Azure DevOps personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] release_api_url: The URL of the release API.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointAzureDevOpsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure DevOps service endpoint within Azure DevOps.

        > **Note** This resource is duplicate with `ServiceEndpointPipeline`,  will be removed in the future, use `ServiceEndpointPipeline` instead.

        > **Note** Prerequisite: Extension [Configurable Pipeline Runner](https://marketplace.visualstudio.com/items?itemName=CSE-DevOps.RunPipelines) has been installed for the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_azure_dev_ops = azuredevops.ServiceEndpointAzureDevOps("example",
            project_id=example.id,
            service_endpoint_name="Example Azure DevOps",
            org_url="https://dev.azure.com/testorganization",
            release_api_url="https://vsrm.dev.azure.com/testorganization",
            personal_access_token="0000000000000000000000000000000000000000000000000000",
            description="Managed by Terraform")
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Service Endpoint Azure DevOps can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
        $ pulumi import azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointAzureDevOpsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointAzureDevOpsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org_url: Optional[pulumi.Input[str]] = None,
                 personal_access_token: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 release_api_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceEndpointAzureDevOpsArgs.__new__(ServiceEndpointAzureDevOpsArgs)

            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            if org_url is None and not opts.urn:
                raise TypeError("Missing required property 'org_url'")
            __props__.__dict__["org_url"] = org_url
            if personal_access_token is None and not opts.urn:
                raise TypeError("Missing required property 'personal_access_token'")
            __props__.__dict__["personal_access_token"] = None if personal_access_token is None else pulumi.Output.secret(personal_access_token)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if release_api_url is None and not opts.urn:
                raise TypeError("Missing required property 'release_api_url'")
            __props__.__dict__["release_api_url"] = release_api_url
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["personalAccessToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceEndpointAzureDevOps, __self__).__init__(
            'azuredevops:index/serviceEndpointAzureDevOps:ServiceEndpointAzureDevOps',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            org_url: Optional[pulumi.Input[str]] = None,
            personal_access_token: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            release_api_url: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None) -> 'ServiceEndpointAzureDevOps':
        """
        Get an existing ServiceEndpointAzureDevOps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_url: The organization URL.
        :param pulumi.Input[str] personal_access_token: The Azure DevOps personal access token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] release_api_url: The URL of the release API.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointAzureDevOpsState.__new__(_ServiceEndpointAzureDevOpsState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["org_url"] = org_url
        __props__.__dict__["personal_access_token"] = personal_access_token
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["release_api_url"] = release_api_url
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return ServiceEndpointAzureDevOps(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="orgUrl")
    def org_url(self) -> pulumi.Output[str]:
        """
        The organization URL.
        """
        return pulumi.get(self, "org_url")

    @property
    @pulumi.getter(name="personalAccessToken")
    def personal_access_token(self) -> pulumi.Output[str]:
        """
        The Azure DevOps personal access token.
        """
        return pulumi.get(self, "personal_access_token")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="releaseApiUrl")
    def release_api_url(self) -> pulumi.Output[str]:
        """
        The URL of the release API.
        """
        return pulumi.get(self, "release_api_url")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

