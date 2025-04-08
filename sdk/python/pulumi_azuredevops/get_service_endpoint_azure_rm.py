# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = [
    'GetServiceEndpointAzureRMResult',
    'AwaitableGetServiceEndpointAzureRMResult',
    'get_service_endpoint_azure_rm',
    'get_service_endpoint_azure_rm_output',
]

@pulumi.output_type
class GetServiceEndpointAzureRMResult:
    """
    A collection of values returned by getServiceEndpointAzureRM.
    """
    def __init__(__self__, authorization=None, azurerm_management_group_id=None, azurerm_management_group_name=None, azurerm_spn_tenantid=None, azurerm_subscription_id=None, azurerm_subscription_name=None, description=None, environment=None, id=None, project_id=None, resource_group=None, server_url=None, service_endpoint_authentication_scheme=None, service_endpoint_id=None, service_endpoint_name=None, service_principal_id=None, workload_identity_federation_issuer=None, workload_identity_federation_subject=None):
        if authorization and not isinstance(authorization, dict):
            raise TypeError("Expected argument 'authorization' to be a dict")
        pulumi.set(__self__, "authorization", authorization)
        if azurerm_management_group_id and not isinstance(azurerm_management_group_id, str):
            raise TypeError("Expected argument 'azurerm_management_group_id' to be a str")
        pulumi.set(__self__, "azurerm_management_group_id", azurerm_management_group_id)
        if azurerm_management_group_name and not isinstance(azurerm_management_group_name, str):
            raise TypeError("Expected argument 'azurerm_management_group_name' to be a str")
        pulumi.set(__self__, "azurerm_management_group_name", azurerm_management_group_name)
        if azurerm_spn_tenantid and not isinstance(azurerm_spn_tenantid, str):
            raise TypeError("Expected argument 'azurerm_spn_tenantid' to be a str")
        pulumi.set(__self__, "azurerm_spn_tenantid", azurerm_spn_tenantid)
        if azurerm_subscription_id and not isinstance(azurerm_subscription_id, str):
            raise TypeError("Expected argument 'azurerm_subscription_id' to be a str")
        pulumi.set(__self__, "azurerm_subscription_id", azurerm_subscription_id)
        if azurerm_subscription_name and not isinstance(azurerm_subscription_name, str):
            raise TypeError("Expected argument 'azurerm_subscription_name' to be a str")
        pulumi.set(__self__, "azurerm_subscription_name", azurerm_subscription_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        pulumi.set(__self__, "resource_group", resource_group)
        if server_url and not isinstance(server_url, str):
            raise TypeError("Expected argument 'server_url' to be a str")
        pulumi.set(__self__, "server_url", server_url)
        if service_endpoint_authentication_scheme and not isinstance(service_endpoint_authentication_scheme, str):
            raise TypeError("Expected argument 'service_endpoint_authentication_scheme' to be a str")
        pulumi.set(__self__, "service_endpoint_authentication_scheme", service_endpoint_authentication_scheme)
        if service_endpoint_id and not isinstance(service_endpoint_id, str):
            raise TypeError("Expected argument 'service_endpoint_id' to be a str")
        pulumi.set(__self__, "service_endpoint_id", service_endpoint_id)
        if service_endpoint_name and not isinstance(service_endpoint_name, str):
            raise TypeError("Expected argument 'service_endpoint_name' to be a str")
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if service_principal_id and not isinstance(service_principal_id, str):
            raise TypeError("Expected argument 'service_principal_id' to be a str")
        pulumi.set(__self__, "service_principal_id", service_principal_id)
        if workload_identity_federation_issuer and not isinstance(workload_identity_federation_issuer, str):
            raise TypeError("Expected argument 'workload_identity_federation_issuer' to be a str")
        pulumi.set(__self__, "workload_identity_federation_issuer", workload_identity_federation_issuer)
        if workload_identity_federation_subject and not isinstance(workload_identity_federation_subject, str):
            raise TypeError("Expected argument 'workload_identity_federation_subject' to be a str")
        pulumi.set(__self__, "workload_identity_federation_subject", workload_identity_federation_subject)

    @property
    @pulumi.getter
    def authorization(self) -> Mapping[str, builtins.str]:
        """
        The Authorization scheme.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="azurermManagementGroupId")
    def azurerm_management_group_id(self) -> builtins.str:
        """
        The Management Group ID of the Service Endpoint is target, if available.
        """
        return pulumi.get(self, "azurerm_management_group_id")

    @property
    @pulumi.getter(name="azurermManagementGroupName")
    def azurerm_management_group_name(self) -> builtins.str:
        """
        The Management Group Name of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "azurerm_management_group_name")

    @property
    @pulumi.getter(name="azurermSpnTenantid")
    def azurerm_spn_tenantid(self) -> builtins.str:
        """
        The Tenant ID of the Azure targets.
        """
        return pulumi.get(self, "azurerm_spn_tenantid")

    @property
    @pulumi.getter(name="azurermSubscriptionId")
    def azurerm_subscription_id(self) -> builtins.str:
        """
        The Subscription ID of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "azurerm_subscription_id")

    @property
    @pulumi.getter(name="azurermSubscriptionName")
    def azurerm_subscription_name(self) -> builtins.str:
        """
        The Subscription Name of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "azurerm_subscription_name")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        The description of the Service Endpoint.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> builtins.str:
        """
        The Cloud Environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> builtins.str:
        """
        The Resource Group of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> builtins.str:
        """
        The server URL of the service Endpoint.
        """
        return pulumi.get(self, "server_url")

    @property
    @pulumi.getter(name="serviceEndpointAuthenticationScheme")
    def service_endpoint_authentication_scheme(self) -> builtins.str:
        """
        The authentication scheme of Azure Resource Management Endpoint
        """
        return pulumi.get(self, "service_endpoint_authentication_scheme")

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> builtins.str:
        return pulumi.get(self, "service_endpoint_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> builtins.str:
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> builtins.str:
        """
        The Application(Client) ID of the Service Principal.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="workloadIdentityFederationIssuer")
    def workload_identity_federation_issuer(self) -> builtins.str:
        """
        The issuer if `of the Workload Identity Federation Subject
        """
        return pulumi.get(self, "workload_identity_federation_issuer")

    @property
    @pulumi.getter(name="workloadIdentityFederationSubject")
    def workload_identity_federation_subject(self) -> builtins.str:
        """
        The subject of the Workload Identity Federation Subject.
        """
        return pulumi.get(self, "workload_identity_federation_subject")


class AwaitableGetServiceEndpointAzureRMResult(GetServiceEndpointAzureRMResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceEndpointAzureRMResult(
            authorization=self.authorization,
            azurerm_management_group_id=self.azurerm_management_group_id,
            azurerm_management_group_name=self.azurerm_management_group_name,
            azurerm_spn_tenantid=self.azurerm_spn_tenantid,
            azurerm_subscription_id=self.azurerm_subscription_id,
            azurerm_subscription_name=self.azurerm_subscription_name,
            description=self.description,
            environment=self.environment,
            id=self.id,
            project_id=self.project_id,
            resource_group=self.resource_group,
            server_url=self.server_url,
            service_endpoint_authentication_scheme=self.service_endpoint_authentication_scheme,
            service_endpoint_id=self.service_endpoint_id,
            service_endpoint_name=self.service_endpoint_name,
            service_principal_id=self.service_principal_id,
            workload_identity_federation_issuer=self.workload_identity_federation_issuer,
            workload_identity_federation_subject=self.workload_identity_federation_subject)


def get_service_endpoint_azure_rm(project_id: Optional[builtins.str] = None,
                                  service_endpoint_id: Optional[builtins.str] = None,
                                  service_endpoint_name: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceEndpointAzureRMResult:
    """
    Use this data source to access information about an existing AzureRM service Endpoint.

    ## Example Usage

    ### By Service Endpoint ID

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    sample = azuredevops.get_project(name="Sample Project")
    serviceendpoint = azuredevops.get_service_endpoint_azure_rm(project_id=sample.id,
        service_endpoint_id="00000000-0000-0000-0000-000000000000")
    pulumi.export("serviceEndpointName", serviceendpoint.service_endpoint_name)
    ```

    ### By Service Endpoint Name

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    sample = azuredevops.get_project(name="Sample Project")
    serviceendpoint = azuredevops.get_service_endpoint_azure_rm(project_id=sample.id,
        service_endpoint_name="Example-Service-Endpoint")
    pulumi.export("serviceEndpointId", serviceendpoint.id)
    ```


    :param builtins.str project_id: The ID of the project.
    :param builtins.str service_endpoint_id: the ID of the Service Endpoint.
    :param builtins.str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** 1. One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
           <br>2. When supplying `service_endpoint_name`, take care to ensure that this is a unique name.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['serviceEndpointId'] = service_endpoint_id
    __args__['serviceEndpointName'] = service_endpoint_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getServiceEndpointAzureRM:getServiceEndpointAzureRM', __args__, opts=opts, typ=GetServiceEndpointAzureRMResult).value

    return AwaitableGetServiceEndpointAzureRMResult(
        authorization=pulumi.get(__ret__, 'authorization'),
        azurerm_management_group_id=pulumi.get(__ret__, 'azurerm_management_group_id'),
        azurerm_management_group_name=pulumi.get(__ret__, 'azurerm_management_group_name'),
        azurerm_spn_tenantid=pulumi.get(__ret__, 'azurerm_spn_tenantid'),
        azurerm_subscription_id=pulumi.get(__ret__, 'azurerm_subscription_id'),
        azurerm_subscription_name=pulumi.get(__ret__, 'azurerm_subscription_name'),
        description=pulumi.get(__ret__, 'description'),
        environment=pulumi.get(__ret__, 'environment'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        resource_group=pulumi.get(__ret__, 'resource_group'),
        server_url=pulumi.get(__ret__, 'server_url'),
        service_endpoint_authentication_scheme=pulumi.get(__ret__, 'service_endpoint_authentication_scheme'),
        service_endpoint_id=pulumi.get(__ret__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__ret__, 'service_endpoint_name'),
        service_principal_id=pulumi.get(__ret__, 'service_principal_id'),
        workload_identity_federation_issuer=pulumi.get(__ret__, 'workload_identity_federation_issuer'),
        workload_identity_federation_subject=pulumi.get(__ret__, 'workload_identity_federation_subject'))
def get_service_endpoint_azure_rm_output(project_id: Optional[pulumi.Input[builtins.str]] = None,
                                         service_endpoint_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         service_endpoint_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceEndpointAzureRMResult]:
    """
    Use this data source to access information about an existing AzureRM service Endpoint.

    ## Example Usage

    ### By Service Endpoint ID

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    sample = azuredevops.get_project(name="Sample Project")
    serviceendpoint = azuredevops.get_service_endpoint_azure_rm(project_id=sample.id,
        service_endpoint_id="00000000-0000-0000-0000-000000000000")
    pulumi.export("serviceEndpointName", serviceendpoint.service_endpoint_name)
    ```

    ### By Service Endpoint Name

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    sample = azuredevops.get_project(name="Sample Project")
    serviceendpoint = azuredevops.get_service_endpoint_azure_rm(project_id=sample.id,
        service_endpoint_name="Example-Service-Endpoint")
    pulumi.export("serviceEndpointId", serviceendpoint.id)
    ```


    :param builtins.str project_id: The ID of the project.
    :param builtins.str service_endpoint_id: the ID of the Service Endpoint.
    :param builtins.str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** 1. One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
           <br>2. When supplying `service_endpoint_name`, take care to ensure that this is a unique name.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['serviceEndpointId'] = service_endpoint_id
    __args__['serviceEndpointName'] = service_endpoint_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getServiceEndpointAzureRM:getServiceEndpointAzureRM', __args__, opts=opts, typ=GetServiceEndpointAzureRMResult)
    return __ret__.apply(lambda __response__: GetServiceEndpointAzureRMResult(
        authorization=pulumi.get(__response__, 'authorization'),
        azurerm_management_group_id=pulumi.get(__response__, 'azurerm_management_group_id'),
        azurerm_management_group_name=pulumi.get(__response__, 'azurerm_management_group_name'),
        azurerm_spn_tenantid=pulumi.get(__response__, 'azurerm_spn_tenantid'),
        azurerm_subscription_id=pulumi.get(__response__, 'azurerm_subscription_id'),
        azurerm_subscription_name=pulumi.get(__response__, 'azurerm_subscription_name'),
        description=pulumi.get(__response__, 'description'),
        environment=pulumi.get(__response__, 'environment'),
        id=pulumi.get(__response__, 'id'),
        project_id=pulumi.get(__response__, 'project_id'),
        resource_group=pulumi.get(__response__, 'resource_group'),
        server_url=pulumi.get(__response__, 'server_url'),
        service_endpoint_authentication_scheme=pulumi.get(__response__, 'service_endpoint_authentication_scheme'),
        service_endpoint_id=pulumi.get(__response__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__response__, 'service_endpoint_name'),
        service_principal_id=pulumi.get(__response__, 'service_principal_id'),
        workload_identity_federation_issuer=pulumi.get(__response__, 'workload_identity_federation_issuer'),
        workload_identity_federation_subject=pulumi.get(__response__, 'workload_identity_federation_subject')))
