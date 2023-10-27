# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetServiceendpointNpmResult',
    'AwaitableGetServiceendpointNpmResult',
    'get_serviceendpoint_npm',
    'get_serviceendpoint_npm_output',
]

@pulumi.output_type
class GetServiceendpointNpmResult:
    """
    A collection of values returned by getServiceendpointNpm.
    """
    def __init__(__self__, authorization=None, description=None, id=None, project_id=None, service_endpoint_id=None, service_endpoint_name=None, url=None):
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
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def authorization(self) -> Mapping[str, str]:
        """
        Specifies the Authorization Scheme Map.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Specifies the description of the Service Endpoint.
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
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> str:
        return pulumi.get(self, "service_endpoint_id")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> str:
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        Specifies the URL of the npm registry to connect with.
        """
        return pulumi.get(self, "url")


class AwaitableGetServiceendpointNpmResult(GetServiceendpointNpmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceendpointNpmResult(
            authorization=self.authorization,
            description=self.description,
            id=self.id,
            project_id=self.project_id,
            service_endpoint_id=self.service_endpoint_id,
            service_endpoint_name=self.service_endpoint_name,
            url=self.url)


def get_serviceendpoint_npm(project_id: Optional[str] = None,
                            service_endpoint_id: Optional[str] = None,
                            service_endpoint_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceendpointNpmResult:
    """
    Use this data source to access information about an existing NPM Service Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_serviceendpoint_npm(project_id=azuredevops_project["example"]["id"],
        service_endpoint_name="Example npm")
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
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getServiceendpointNpm:getServiceendpointNpm', __args__, opts=opts, typ=GetServiceendpointNpmResult).value

    return AwaitableGetServiceendpointNpmResult(
        authorization=pulumi.get(__ret__, 'authorization'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        service_endpoint_id=pulumi.get(__ret__, 'service_endpoint_id'),
        service_endpoint_name=pulumi.get(__ret__, 'service_endpoint_name'),
        url=pulumi.get(__ret__, 'url'))


@_utilities.lift_output_func(get_serviceendpoint_npm)
def get_serviceendpoint_npm_output(project_id: Optional[pulumi.Input[str]] = None,
                                   service_endpoint_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   service_endpoint_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceendpointNpmResult]:
    """
    Use this data source to access information about an existing NPM Service Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_serviceendpoint_npm(project_id=azuredevops_project["example"]["id"],
        service_endpoint_name="Example npm")
    pulumi.export("serviceEndpointId", example.id)
    ```


    :param str project_id: The ID of the project.
    :param str service_endpoint_id: the ID of the Service Endpoint.
    :param str service_endpoint_name: the Name of the Service Endpoint.
           
           > **NOTE:** One of either `service_endpoint_id` or `service_endpoint_name` must be specified.
    """
    ...
