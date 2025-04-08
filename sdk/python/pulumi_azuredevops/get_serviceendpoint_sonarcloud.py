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
    'GetServiceendpointSonarcloudResult',
    'AwaitableGetServiceendpointSonarcloudResult',
    'get_serviceendpoint_sonarcloud',
    'get_serviceendpoint_sonarcloud_output',
]

@pulumi.output_type
class GetServiceendpointSonarcloudResult:
    """
    A collection of values returned by getServiceendpointSonarcloud.
    """
    def __init__(__self__, authorization=None, description=None, id=None, project_id=None, service_endpoint_id=None, service_endpoint_name=None):
        if authorization and not isinstance(authorization, dict):
            raise TypeError("Expected argument 'authorization' to be a dict")
        pulumi.set(__self__, "authorization", authorization)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if service_endpoint_id and not isinstance(service_endpoint_id, str):
            raise TypeError("Expected argument 'service_endpoint_id' to be a str")
        pulumi.set(__self__, "service_endpoint_id", service_endpoint_id)
        if service_endpoint_name and not isinstance(service_endpoint_name, str):
            raise TypeError("Expected argument 'service_endpoint_name' to be a str")
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)

    @property
    @pulumi.getter
    def authorization(self) -> Mapping[str, builtins.str]:
        """
        The Authorization scheme.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        The description of the Service Endpoint.
        """
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> builtins.str:
        return pulumi.get(self, "service_endpoint_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> builtins.str:
        return pulumi.get(self, "service_endpoint_name")


class AwaitableGetServiceendpointSonarcloudResult(GetServiceendpointSonarcloudResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceendpointSonarcloudResult(
            authorization=self.authorization,
            description=self.description,
            id=self.id,
            project_id=self.project_id,
            service_endpoint_id=self.service_endpoint_id,
            service_endpoint_name=self.service_endpoint_name)


def get_serviceendpoint_sonarcloud(project_id: Optional[builtins.str] = None,
                                   service_endpoint_id: Optional[builtins.str] = None,
                                   service_endpoint_name: Optional[builtins.str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceendpointSonarcloudResult:
    """
    Use this data source to access information about an existing Sonar Cloud Service Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_serviceendpoint_sonarcloud(project_id=example_azuredevops_project["id"],
        service_endpoint_name="Example Sonar Cloud")
    pulumi.export("serviceEndpointId", example.id)
    ```


    :param builtins.str project_id: The ID of the project.
    :param builtins.str service_endpoint_id: the ID of the Service Endpoint.
    :param builtins.str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['serviceEndpointId'] = service_endpoint_id
    __args__['serviceEndpointName'] = service_endpoint_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getServiceendpointSonarcloud:getServiceendpointSonarcloud', __args__, opts=opts, typ=GetServiceendpointSonarcloudResult).value

    return AwaitableGetServiceendpointSonarcloudResult(
        authorization=pulumi.get(__ret__, 'authorization'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        service_endpoint_id=pulumi.get(__ret__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__ret__, 'service_endpoint_name'))
def get_serviceendpoint_sonarcloud_output(project_id: Optional[pulumi.Input[builtins.str]] = None,
                                          service_endpoint_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          service_endpoint_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceendpointSonarcloudResult]:
    """
    Use this data source to access information about an existing Sonar Cloud Service Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_serviceendpoint_sonarcloud(project_id=example_azuredevops_project["id"],
        service_endpoint_name="Example Sonar Cloud")
    pulumi.export("serviceEndpointId", example.id)
    ```


    :param builtins.str project_id: The ID of the project.
    :param builtins.str service_endpoint_id: the ID of the Service Endpoint.
    :param builtins.str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['serviceEndpointId'] = service_endpoint_id
    __args__['serviceEndpointName'] = service_endpoint_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getServiceendpointSonarcloud:getServiceendpointSonarcloud', __args__, opts=opts, typ=GetServiceendpointSonarcloudResult)
    return __ret__.apply(lambda __response__: GetServiceendpointSonarcloudResult(
        authorization=pulumi.get(__response__, 'authorization'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        project_id=pulumi.get(__response__, 'project_id'),
        service_endpoint_id=pulumi.get(__response__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__response__, 'service_endpoint_name')))
