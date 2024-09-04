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
    def __init__(__self__, authorization=None, azurerm_management_group_id=None, azurerm_management_group_name=None, azurerm_spn_tenantid=None, azurerm_subscription_id=None, azurerm_subscription_name=None, description=None, environment=None, id=None, project_id=None, resource_group=None, service_endpoint_authentication_scheme=None, service_endpoint_id=None, service_endpoint_name=None, service_principal_id=None, workload_identity_federation_issuer=None, workload_identity_federation_subject=None):
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
    def authorization(self) -> Mapping[str, str]:
        """
        Specifies the Authorization Scheme Map.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="azurermManagementGroupId")
    def azurerm_management_group_id(self) -> str:
        """
        Specified the Management Group ID of the Service Endpoint is target, if available.
        """
        return pulumi.get(self, "azurerm_management_group_id")

    @property
    @pulumi.getter(name="azurermManagementGroupName")
    def azurerm_management_group_name(self) -> str:
        """
        Specified the Management Group Name of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "azurerm_management_group_name")

    @property
    @pulumi.getter(name="azurermSpnTenantid")
    def azurerm_spn_tenantid(self) -> str:
        """
        Specifies the Tenant ID of the Azure targets.
        """
        return pulumi.get(self, "azurerm_spn_tenantid")

    @property
    @pulumi.getter(name="azurermSubscriptionId")
    def azurerm_subscription_id(self) -> str:
        """
        Specifies the Subscription ID of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "azurerm_subscription_id")

    @property
    @pulumi.getter(name="azurermSubscriptionName")
    def azurerm_subscription_name(self) -> str:
        """
        Specifies the Subscription Name of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "azurerm_subscription_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Specifies the description of the Service Endpoint.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        The Cloud Environment. Possible values are `AzureCloud`, `AzureChinaCloud`, `AzureUSGovernment`, and `AzureGermanCloud`.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Specifies the Resource Group of the Service Endpoint target, if available.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="serviceEndpointAuthenticationScheme")
    def service_endpoint_authentication_scheme(self) -> str:
        """
        Specifies the authentication scheme of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`.
        """
        return pulumi.get(self, "service_endpoint_authentication_scheme")

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> str:
        return pulumi.get(self, "service_endpoint_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> str:
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> str:
        """
        The Application(Client) ID of the Service Principal.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="workloadIdentityFederationIssuer")
    def workload_identity_federation_issuer(self) -> str:
        """
        The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/f66a4bc2-08ad-4ec0-a25e-e769d6b3b294`, where the GUID is the Organization ID of your Azure DevOps Organisation.
        """
        return pulumi.get(self, "workload_identity_federation_issuer")

    @property
    @pulumi.getter(name="workloadIdentityFederationSubject")
    def workload_identity_federation_subject(self) -> str:
        """
        The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://my-organisation/my-project/my-service-connection-name`.
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
            service_endpoint_authentication_scheme=self.service_endpoint_authentication_scheme,
            service_endpoint_id=self.service_endpoint_id,
            service_endpoint_name=self.service_endpoint_name,
            service_principal_id=self.service_principal_id,
            workload_identity_federation_issuer=self.workload_identity_federation_issuer,
            workload_identity_federation_subject=self.workload_identity_federation_subject)


def get_service_endpoint_azure_rm(project_id: Optional[str] = None,
                                  service_endpoint_id: Optional[str] = None,
                                  service_endpoint_name: Optional[str] = None,
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


    :param str project_id: The ID of the project.
    :param str service_endpoint_id: the ID of the Service Endpoint.
    :param str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
           > **NOTE:** When supplying `service_endpoint_name`, take care to ensure that this is a unique name.
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
        service_endpoint_authentication_scheme=pulumi.get(__ret__, 'service_endpoint_authentication_scheme'),
        service_endpoint_id=pulumi.get(__ret__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__ret__, 'service_endpoint_name'),
        service_principal_id=pulumi.get(__ret__, 'service_principal_id'),
        workload_identity_federation_issuer=pulumi.get(__ret__, 'workload_identity_federation_issuer'),
        workload_identity_federation_subject=pulumi.get(__ret__, 'workload_identity_federation_subject'))


@_utilities.lift_output_func(get_service_endpoint_azure_rm)
def get_service_endpoint_azure_rm_output(project_id: Optional[pulumi.Input[str]] = None,
                                         service_endpoint_id: Optional[pulumi.Input[Optional[str]]] = None,
                                         service_endpoint_name: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceEndpointAzureRMResult]:
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


    :param str project_id: The ID of the project.
    :param str service_endpoint_id: the ID of the Service Endpoint.
    :param str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
           > **NOTE:** When supplying `service_endpoint_name`, take care to ensure that this is a unique name.
    """
    ...
