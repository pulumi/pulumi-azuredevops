# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AzureRMArgs', 'AzureRM']

@pulumi.input_type
class AzureRMArgs:
    def __init__(__self__, *,
                 azurerm_spn_tenantid: pulumi.Input[str],
                 azurerm_subscription_id: pulumi.Input[str],
                 azurerm_subscription_name: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 credentials: Optional[pulumi.Input['AzureRMCredentialsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AzureRM resource.
        :param pulumi.Input[str] azurerm_spn_tenantid: The tenant id if the service principal.
        :param pulumi.Input[str] azurerm_subscription_id: The subscription Id of the Azure targets.
        :param pulumi.Input[str] azurerm_subscription_name: The subscription Name of the targets.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input['AzureRMCredentialsArgs'] credentials: A `credentials` block.
        :param pulumi.Input[str] description: Service connection description.
        :param pulumi.Input[str] resource_group: The resource group used for scope of automatic service endpoint.
        """
        pulumi.set(__self__, "azurerm_spn_tenantid", azurerm_spn_tenantid)
        pulumi.set(__self__, "azurerm_subscription_id", azurerm_subscription_id)
        pulumi.set(__self__, "azurerm_subscription_name", azurerm_subscription_name)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if resource_group is not None:
            pulumi.set(__self__, "resource_group", resource_group)

    @property
    @pulumi.getter(name="azurermSpnTenantid")
    def azurerm_spn_tenantid(self) -> pulumi.Input[str]:
        """
        The tenant id if the service principal.
        """
        return pulumi.get(self, "azurerm_spn_tenantid")

    @azurerm_spn_tenantid.setter
    def azurerm_spn_tenantid(self, value: pulumi.Input[str]):
        pulumi.set(self, "azurerm_spn_tenantid", value)

    @property
    @pulumi.getter(name="azurermSubscriptionId")
    def azurerm_subscription_id(self) -> pulumi.Input[str]:
        """
        The subscription Id of the Azure targets.
        """
        return pulumi.get(self, "azurerm_subscription_id")

    @azurerm_subscription_id.setter
    def azurerm_subscription_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "azurerm_subscription_id", value)

    @property
    @pulumi.getter(name="azurermSubscriptionName")
    def azurerm_subscription_name(self) -> pulumi.Input[str]:
        """
        The subscription Name of the targets.
        """
        return pulumi.get(self, "azurerm_subscription_name")

    @azurerm_subscription_name.setter
    def azurerm_subscription_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "azurerm_subscription_name", value)

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
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input['AzureRMCredentialsArgs']]:
        """
        A `credentials` block.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input['AzureRMCredentialsArgs']]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Service connection description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group used for scope of automatic service endpoint.
        """
        return pulumi.get(self, "resource_group")

    @resource_group.setter
    def resource_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group", value)


@pulumi.input_type
class _AzureRMState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 azurerm_spn_tenantid: Optional[pulumi.Input[str]] = None,
                 azurerm_subscription_id: Optional[pulumi.Input[str]] = None,
                 azurerm_subscription_name: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input['AzureRMCredentialsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_group: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AzureRM resources.
        :param pulumi.Input[str] azurerm_spn_tenantid: The tenant id if the service principal.
        :param pulumi.Input[str] azurerm_subscription_id: The subscription Id of the Azure targets.
        :param pulumi.Input[str] azurerm_subscription_name: The subscription Name of the targets.
        :param pulumi.Input['AzureRMCredentialsArgs'] credentials: A `credentials` block.
        :param pulumi.Input[str] description: Service connection description.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] resource_group: The resource group used for scope of automatic service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if azurerm_spn_tenantid is not None:
            pulumi.set(__self__, "azurerm_spn_tenantid", azurerm_spn_tenantid)
        if azurerm_subscription_id is not None:
            pulumi.set(__self__, "azurerm_subscription_id", azurerm_subscription_id)
        if azurerm_subscription_name is not None:
            pulumi.set(__self__, "azurerm_subscription_name", azurerm_subscription_name)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if resource_group is not None:
            pulumi.set(__self__, "resource_group", resource_group)
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
    @pulumi.getter(name="azurermSpnTenantid")
    def azurerm_spn_tenantid(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant id if the service principal.
        """
        return pulumi.get(self, "azurerm_spn_tenantid")

    @azurerm_spn_tenantid.setter
    def azurerm_spn_tenantid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azurerm_spn_tenantid", value)

    @property
    @pulumi.getter(name="azurermSubscriptionId")
    def azurerm_subscription_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription Id of the Azure targets.
        """
        return pulumi.get(self, "azurerm_subscription_id")

    @azurerm_subscription_id.setter
    def azurerm_subscription_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azurerm_subscription_id", value)

    @property
    @pulumi.getter(name="azurermSubscriptionName")
    def azurerm_subscription_name(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription Name of the targets.
        """
        return pulumi.get(self, "azurerm_subscription_name")

    @azurerm_subscription_name.setter
    def azurerm_subscription_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azurerm_subscription_name", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input['AzureRMCredentialsArgs']]:
        """
        A `credentials` block.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input['AzureRMCredentialsArgs']]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Service connection description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group used for scope of automatic service endpoint.
        """
        return pulumi.get(self, "resource_group")

    @resource_group.setter
    def resource_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group", value)

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


warnings.warn("""azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM""", DeprecationWarning)


class AzureRM(pulumi.CustomResource):
    warnings.warn("""azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 azurerm_spn_tenantid: Optional[pulumi.Input[str]] = None,
                 azurerm_subscription_id: Optional[pulumi.Input[str]] = None,
                 azurerm_subscription_name: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['AzureRMCredentialsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_group: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.

        ## Requirements (Manual AzureRM Service Endpoint)

        Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.

        For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)

        ## Example Usage
        ### Manual AzureRM Service Endpoint

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        endpointazure = azuredevops.ServiceEndpointAzureRM("endpointazure",
            project_id=project.id,
            service_endpoint_name="Sample AzureRM",
            description="Managed by Terraform",
            credentials=azuredevops.ServiceEndpointAzureRMCredentialsArgs(
                serviceprincipalid="00000000-0000-0000-0000-000000000000",
                serviceprincipalkey="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            ),
            azurerm_spn_tenantid="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_id="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_name="Sample Subscription")
        ```
        ### Automatic AzureRM Service Endpoint

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        endpointazure = azuredevops.ServiceEndpointAzureRM("endpointazure",
            project_id=project.id,
            service_endpoint_name="Sample AzureRM",
            description="Managed by Terraform",
            azurerm_spn_tenantid="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_id="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_name="Microsoft Azure DEMO")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] azurerm_spn_tenantid: The tenant id if the service principal.
        :param pulumi.Input[str] azurerm_subscription_id: The subscription Id of the Azure targets.
        :param pulumi.Input[str] azurerm_subscription_name: The subscription Name of the targets.
        :param pulumi.Input[pulumi.InputType['AzureRMCredentialsArgs']] credentials: A `credentials` block.
        :param pulumi.Input[str] description: Service connection description.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] resource_group: The resource group used for scope of automatic service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureRMArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Manual or Automatic AzureRM service endpoint within Azure DevOps.

        ## Requirements (Manual AzureRM Service Endpoint)

        Before to create a service end point in Azure DevOps, you need to create a Service Principal in your Azure subscription.

        For detailed steps to create a service principal with Azure cli see the [documentation](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli?view=azure-cli-latest)

        ## Example Usage
        ### Manual AzureRM Service Endpoint

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        endpointazure = azuredevops.ServiceEndpointAzureRM("endpointazure",
            project_id=project.id,
            service_endpoint_name="Sample AzureRM",
            description="Managed by Terraform",
            credentials=azuredevops.ServiceEndpointAzureRMCredentialsArgs(
                serviceprincipalid="00000000-0000-0000-0000-000000000000",
                serviceprincipalkey="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            ),
            azurerm_spn_tenantid="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_id="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_name="Sample Subscription")
        ```
        ### Automatic AzureRM Service Endpoint

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        endpointazure = azuredevops.ServiceEndpointAzureRM("endpointazure",
            project_id=project.id,
            service_endpoint_name="Sample AzureRM",
            description="Managed by Terraform",
            azurerm_spn_tenantid="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_id="00000000-0000-0000-0000-000000000000",
            azurerm_subscription_name="Microsoft Azure DEMO")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Service End points](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Service Endpoint Azure Resource Manage can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:ServiceEndpoint/azureRM:AzureRM serviceendpoint 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AzureRMArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureRMArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 azurerm_spn_tenantid: Optional[pulumi.Input[str]] = None,
                 azurerm_subscription_id: Optional[pulumi.Input[str]] = None,
                 azurerm_subscription_name: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['AzureRMCredentialsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_group: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        pulumi.log.warn("""AzureRM is deprecated: azuredevops.serviceendpoint.AzureRM has been deprecated in favor of azuredevops.ServiceEndpointAzureRM""")
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
            __props__ = AzureRMArgs.__new__(AzureRMArgs)

            __props__.__dict__["authorization"] = authorization
            if azurerm_spn_tenantid is None and not opts.urn:
                raise TypeError("Missing required property 'azurerm_spn_tenantid'")
            __props__.__dict__["azurerm_spn_tenantid"] = azurerm_spn_tenantid
            if azurerm_subscription_id is None and not opts.urn:
                raise TypeError("Missing required property 'azurerm_subscription_id'")
            __props__.__dict__["azurerm_subscription_id"] = azurerm_subscription_id
            if azurerm_subscription_name is None and not opts.urn:
                raise TypeError("Missing required property 'azurerm_subscription_name'")
            __props__.__dict__["azurerm_subscription_name"] = azurerm_subscription_name
            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["description"] = description
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["resource_group"] = resource_group
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        super(AzureRM, __self__).__init__(
            'azuredevops:ServiceEndpoint/azureRM:AzureRM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            azurerm_spn_tenantid: Optional[pulumi.Input[str]] = None,
            azurerm_subscription_id: Optional[pulumi.Input[str]] = None,
            azurerm_subscription_name: Optional[pulumi.Input[str]] = None,
            credentials: Optional[pulumi.Input[pulumi.InputType['AzureRMCredentialsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            resource_group: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None) -> 'AzureRM':
        """
        Get an existing AzureRM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] azurerm_spn_tenantid: The tenant id if the service principal.
        :param pulumi.Input[str] azurerm_subscription_id: The subscription Id of the Azure targets.
        :param pulumi.Input[str] azurerm_subscription_name: The subscription Name of the targets.
        :param pulumi.Input[pulumi.InputType['AzureRMCredentialsArgs']] credentials: A `credentials` block.
        :param pulumi.Input[str] description: Service connection description.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] resource_group: The resource group used for scope of automatic service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AzureRMState.__new__(_AzureRMState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["azurerm_spn_tenantid"] = azurerm_spn_tenantid
        __props__.__dict__["azurerm_subscription_id"] = azurerm_subscription_id
        __props__.__dict__["azurerm_subscription_name"] = azurerm_subscription_name
        __props__.__dict__["credentials"] = credentials
        __props__.__dict__["description"] = description
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["resource_group"] = resource_group
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        return AzureRM(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="azurermSpnTenantid")
    def azurerm_spn_tenantid(self) -> pulumi.Output[str]:
        """
        The tenant id if the service principal.
        """
        return pulumi.get(self, "azurerm_spn_tenantid")

    @property
    @pulumi.getter(name="azurermSubscriptionId")
    def azurerm_subscription_id(self) -> pulumi.Output[str]:
        """
        The subscription Id of the Azure targets.
        """
        return pulumi.get(self, "azurerm_subscription_id")

    @property
    @pulumi.getter(name="azurermSubscriptionName")
    def azurerm_subscription_name(self) -> pulumi.Output[str]:
        """
        The subscription Name of the targets.
        """
        return pulumi.get(self, "azurerm_subscription_name")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional['outputs.AzureRMCredentials']]:
        """
        A `credentials` block.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Service connection description.
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
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> pulumi.Output[Optional[str]]:
        """
        The resource group used for scope of automatic service endpoint.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

