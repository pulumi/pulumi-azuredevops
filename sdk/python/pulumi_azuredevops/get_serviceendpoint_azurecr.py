# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetServiceendpointAzurecrResult',
    'AwaitableGetServiceendpointAzurecrResult',
    'get_serviceendpoint_azurecr',
    'get_serviceendpoint_azurecr_output',
]

@pulumi.output_type
class GetServiceendpointAzurecrResult:
    """
    A collection of values returned by getServiceendpointAzurecr.
    """
    def __init__(__self__, app_object_id=None, authorization=None, az_spn_role_assignment_id=None, az_spn_role_permissions=None, azurecr_name=None, azurecr_spn_tenantid=None, azurecr_subscription_id=None, azurecr_subscription_name=None, description=None, id=None, project_id=None, resource_group=None, service_endpoint_id=None, service_endpoint_name=None, service_principal_id=None, spn_object_id=None):
        if app_object_id and not isinstance(app_object_id, str):
            raise TypeError("Expected argument 'app_object_id' to be a str")
        pulumi.set(__self__, "app_object_id", app_object_id)
        if authorization and not isinstance(authorization, dict):
            raise TypeError("Expected argument 'authorization' to be a dict")
        pulumi.set(__self__, "authorization", authorization)
        if az_spn_role_assignment_id and not isinstance(az_spn_role_assignment_id, str):
            raise TypeError("Expected argument 'az_spn_role_assignment_id' to be a str")
        pulumi.set(__self__, "az_spn_role_assignment_id", az_spn_role_assignment_id)
        if az_spn_role_permissions and not isinstance(az_spn_role_permissions, str):
            raise TypeError("Expected argument 'az_spn_role_permissions' to be a str")
        pulumi.set(__self__, "az_spn_role_permissions", az_spn_role_permissions)
        if azurecr_name and not isinstance(azurecr_name, str):
            raise TypeError("Expected argument 'azurecr_name' to be a str")
        pulumi.set(__self__, "azurecr_name", azurecr_name)
        if azurecr_spn_tenantid and not isinstance(azurecr_spn_tenantid, str):
            raise TypeError("Expected argument 'azurecr_spn_tenantid' to be a str")
        pulumi.set(__self__, "azurecr_spn_tenantid", azurecr_spn_tenantid)
        if azurecr_subscription_id and not isinstance(azurecr_subscription_id, str):
            raise TypeError("Expected argument 'azurecr_subscription_id' to be a str")
        pulumi.set(__self__, "azurecr_subscription_id", azurecr_subscription_id)
        if azurecr_subscription_name and not isinstance(azurecr_subscription_name, str):
            raise TypeError("Expected argument 'azurecr_subscription_name' to be a str")
        pulumi.set(__self__, "azurecr_subscription_name", azurecr_subscription_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        pulumi.set(__self__, "resource_group", resource_group)
        if service_endpoint_id and not isinstance(service_endpoint_id, str):
            raise TypeError("Expected argument 'service_endpoint_id' to be a str")
        pulumi.set(__self__, "service_endpoint_id", service_endpoint_id)
        if service_endpoint_name and not isinstance(service_endpoint_name, str):
            raise TypeError("Expected argument 'service_endpoint_name' to be a str")
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if service_principal_id and not isinstance(service_principal_id, str):
            raise TypeError("Expected argument 'service_principal_id' to be a str")
        pulumi.set(__self__, "service_principal_id", service_principal_id)
        if spn_object_id and not isinstance(spn_object_id, str):
            raise TypeError("Expected argument 'spn_object_id' to be a str")
        pulumi.set(__self__, "spn_object_id", spn_object_id)

    @property
    @pulumi.getter(name="appObjectId")
    def app_object_id(self) -> str:
        """
        The Object ID of the Service Principal.
        """
        return pulumi.get(self, "app_object_id")

    @property
    @pulumi.getter
    def authorization(self) -> Mapping[str, str]:
        """
        Specifies the Authorization Scheme Map.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="azSpnRoleAssignmentId")
    def az_spn_role_assignment_id(self) -> str:
        """
        The ID of Service Principal Role Assignment.
        """
        return pulumi.get(self, "az_spn_role_assignment_id")

    @property
    @pulumi.getter(name="azSpnRolePermissions")
    def az_spn_role_permissions(self) -> str:
        """
        The Service Principal Role Permissions.
        """
        return pulumi.get(self, "az_spn_role_permissions")

    @property
    @pulumi.getter(name="azurecrName")
    def azurecr_name(self) -> str:
        """
        The Azure Container Registry name.
        """
        return pulumi.get(self, "azurecr_name")

    @property
    @pulumi.getter(name="azurecrSpnTenantid")
    def azurecr_spn_tenantid(self) -> str:
        """
        The Tenant ID of the service principal.
        """
        return pulumi.get(self, "azurecr_spn_tenantid")

    @property
    @pulumi.getter(name="azurecrSubscriptionId")
    def azurecr_subscription_id(self) -> str:
        """
        The Subscription ID of the Azure targets.
        """
        return pulumi.get(self, "azurecr_subscription_id")

    @property
    @pulumi.getter(name="azurecrSubscriptionName")
    def azurecr_subscription_name(self) -> str:
        """
        The Subscription Name of the Azure targets.
        """
        return pulumi.get(self, "azurecr_subscription_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The Service Endpoint description.
        """
        return pulumi.get(self, "description")

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
        The Resource Group to which the Container Registry belongs.
        """
        return pulumi.get(self, "resource_group")

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
    @pulumi.getter(name="spnObjectId")
    def spn_object_id(self) -> str:
        """
        The ID of the Service Principal.
        """
        return pulumi.get(self, "spn_object_id")


class AwaitableGetServiceendpointAzurecrResult(GetServiceendpointAzurecrResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceendpointAzurecrResult(
            app_object_id=self.app_object_id,
            authorization=self.authorization,
            az_spn_role_assignment_id=self.az_spn_role_assignment_id,
            az_spn_role_permissions=self.az_spn_role_permissions,
            azurecr_name=self.azurecr_name,
            azurecr_spn_tenantid=self.azurecr_spn_tenantid,
            azurecr_subscription_id=self.azurecr_subscription_id,
            azurecr_subscription_name=self.azurecr_subscription_name,
            description=self.description,
            id=self.id,
            project_id=self.project_id,
            resource_group=self.resource_group,
            service_endpoint_id=self.service_endpoint_id,
            service_endpoint_name=self.service_endpoint_name,
            service_principal_id=self.service_principal_id,
            spn_object_id=self.spn_object_id)


def get_serviceendpoint_azurecr(project_id: Optional[str] = None,
                                service_endpoint_id: Optional[str] = None,
                                service_endpoint_name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceendpointAzurecrResult:
    """
    Use this data source to access information about an existing Azure Container Registry Service Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_serviceendpoint_azurecr(project_id=azuredevops_project["example"]["id"],
        service_endpoint_name="Example Azure Container Registry")
    pulumi.export("serviceEndpointId", example.id)
    ```


    :param str project_id: The ID of the project.
    :param str service_endpoint_id: the ID of the Service Endpoint.
    :param str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['serviceEndpointId'] = service_endpoint_id
    __args__['serviceEndpointName'] = service_endpoint_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getServiceendpointAzurecr:getServiceendpointAzurecr', __args__, opts=opts, typ=GetServiceendpointAzurecrResult).value

    return AwaitableGetServiceendpointAzurecrResult(
        app_object_id=pulumi.get(__ret__, 'app_object_id'),
        authorization=pulumi.get(__ret__, 'authorization'),
        az_spn_role_assignment_id=pulumi.get(__ret__, 'az_spn_role_assignment_id'),
        az_spn_role_permissions=pulumi.get(__ret__, 'az_spn_role_permissions'),
        azurecr_name=pulumi.get(__ret__, 'azurecr_name'),
        azurecr_spn_tenantid=pulumi.get(__ret__, 'azurecr_spn_tenantid'),
        azurecr_subscription_id=pulumi.get(__ret__, 'azurecr_subscription_id'),
        azurecr_subscription_name=pulumi.get(__ret__, 'azurecr_subscription_name'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        resource_group=pulumi.get(__ret__, 'resource_group'),
        service_endpoint_id=pulumi.get(__ret__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__ret__, 'service_endpoint_name'),
        service_principal_id=pulumi.get(__ret__, 'service_principal_id'),
        spn_object_id=pulumi.get(__ret__, 'spn_object_id'))


@_utilities.lift_output_func(get_serviceendpoint_azurecr)
def get_serviceendpoint_azurecr_output(project_id: Optional[pulumi.Input[str]] = None,
                                       service_endpoint_id: Optional[pulumi.Input[Optional[str]]] = None,
                                       service_endpoint_name: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceendpointAzurecrResult]:
    """
    Use this data source to access information about an existing Azure Container Registry Service Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_serviceendpoint_azurecr(project_id=azuredevops_project["example"]["id"],
        service_endpoint_name="Example Azure Container Registry")
    pulumi.export("serviceEndpointId", example.id)
    ```


    :param str project_id: The ID of the project.
    :param str service_endpoint_id: the ID of the Service Endpoint.
    :param str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
    """
    ...
