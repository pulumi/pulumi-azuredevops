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
    'GetAgentQueueResult',
    'AwaitableGetAgentQueueResult',
    'get_agent_queue',
    'get_agent_queue_output',
]

@pulumi.output_type
class GetAgentQueueResult:
    """
    A collection of values returned by getAgentQueue.
    """
    def __init__(__self__, agent_pool_id=None, id=None, name=None, project_id=None):
        if agent_pool_id and not isinstance(agent_pool_id, int):
            raise TypeError("Expected argument 'agent_pool_id' to be a int")
        pulumi.set(__self__, "agent_pool_id", agent_pool_id)
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
    @pulumi.getter(name="agentPoolId")
    def agent_pool_id(self) -> builtins.int:
        """
        The ID of the Agent pool to which the agent queue belongs.
        """
        return pulumi.get(self, "agent_pool_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the agent queue.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.str:
        """
        The ID of the Project to which the agent queue belongs.
        """
        return pulumi.get(self, "project_id")


class AwaitableGetAgentQueueResult(GetAgentQueueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgentQueueResult(
            agent_pool_id=self.agent_pool_id,
            id=self.id,
            name=self.name,
            project_id=self.project_id)


def get_agent_queue(name: Optional[builtins.str] = None,
                    project_id: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgentQueueResult:
    """
    Use this data source to access information about an existing Agent Queue within Azure DevOps.

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
    example = azuredevops.get_agent_queue_output(project_id=example_project.id,
        name="Example Agent Queue")
    pulumi.export("name", example.name)
    pulumi.export("poolId", example.agent_pool_id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Agent Queues - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues/get?view=azure-devops-rest-7.0)


    :param builtins.str name: Name of the Agent Queue.
    :param builtins.str project_id: The Project Id.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getAgentQueue:getAgentQueue', __args__, opts=opts, typ=GetAgentQueueResult).value

    return AwaitableGetAgentQueueResult(
        agent_pool_id=pulumi.get(__ret__, 'agent_pool_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'))
def get_agent_queue_output(name: Optional[pulumi.Input[builtins.str]] = None,
                           project_id: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAgentQueueResult]:
    """
    Use this data source to access information about an existing Agent Queue within Azure DevOps.

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
    example = azuredevops.get_agent_queue_output(project_id=example_project.id,
        name="Example Agent Queue")
    pulumi.export("name", example.name)
    pulumi.export("poolId", example.agent_pool_id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Agent Queues - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues/get?view=azure-devops-rest-7.0)


    :param builtins.str name: Name of the Agent Queue.
    :param builtins.str project_id: The Project Id.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuredevops:index/getAgentQueue:getAgentQueue', __args__, opts=opts, typ=GetAgentQueueResult)
    return __ret__.apply(lambda __response__: GetAgentQueueResult(
        agent_pool_id=pulumi.get(__response__, 'agent_pool_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id')))
