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
    'GetEnvironmentResult',
    'AwaitableGetEnvironmentResult',
    'get_environment',
    'get_environment_output',
]

@pulumi.output_type
class GetEnvironmentResult:
    """
    A collection of values returned by getEnvironment.
    """
    def __init__(__self__, description=None, environment_id=None, id=None, name=None, project_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if environment_id and not isinstance(environment_id, int):
            raise TypeError("Expected argument 'environment_id' to be a int")
        pulumi.set(__self__, "environment_id", environment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description for the Environment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[int]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Environment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")


class AwaitableGetEnvironmentResult(GetEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentResult(
            description=self.description,
            environment_id=self.environment_id,
            id=self.id,
            name=self.name,
            project_id=self.project_id)


def get_environment(environment_id: Optional[int] = None,
                    name: Optional[str] = None,
                    project_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentResult:
    """
    Use this data source to access information about an Environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.Project("example",
        name="Example Project",
        work_item_template="Agile",
        version_control="Git",
        visibility="private",
        description="Managed by Pulumi")
    example_environment = azuredevops.Environment("example",
        project_id=example_project.id,
        name="Example Environment",
        description="Managed by Pulumi")
    example = pulumi.Output.all(
        exampleProjectId=example_project.id,
        exampleEnvironmentId=example_environment.id
    ).apply(lambda resolved_outputs: azuredevops.get_environment_output(project_id=resolved_outputs['exampleProjectId'],
        environment_id=%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference)))
    ```

    ## Relevant Links

    * [Azure DevOps Service REST API 7.0 - Environments](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/environments?view=azure-devops-rest-7.0)


    :param int environment_id: The ID of the Environment.
    :param str name: Name of the Environment.
           
           > **NOTE:** One of either `environment_id` or `name` must be specified.
    :param str project_id: The ID of the project.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getEnvironment:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult).value

    return AwaitableGetEnvironmentResult(
        description=pulumi.get(__ret__, 'description'),
        environment_id=pulumi.get(__ret__, 'environment_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'))
def get_environment_output(environment_id: Optional[pulumi.Input[Optional[int]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           project_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentResult]:
    """
    Use this data source to access information about an Environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.Project("example",
        name="Example Project",
        work_item_template="Agile",
        version_control="Git",
        visibility="private",
        description="Managed by Pulumi")
    example_environment = azuredevops.Environment("example",
        project_id=example_project.id,
        name="Example Environment",
        description="Managed by Pulumi")
    example = pulumi.Output.all(
        exampleProjectId=example_project.id,
        exampleEnvironmentId=example_environment.id
    ).apply(lambda resolved_outputs: azuredevops.get_environment_output(project_id=resolved_outputs['exampleProjectId'],
        environment_id=%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference)))
    ```

    ## Relevant Links

    * [Azure DevOps Service REST API 7.0 - Environments](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/environments?view=azure-devops-rest-7.0)


    :param int environment_id: The ID of the Environment.
    :param str name: Name of the Environment.
           
           > **NOTE:** One of either `environment_id` or `name` must be specified.
    :param str project_id: The ID of the project.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getEnvironment:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult)
    return __ret__.apply(lambda __response__: GetEnvironmentResult(
        description=pulumi.get(__response__, 'description'),
        environment_id=pulumi.get(__response__, 'environment_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id')))
