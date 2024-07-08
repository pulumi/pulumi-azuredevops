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
from . import outputs

__all__ = [
    'GetBuildDefinitionResult',
    'AwaitableGetBuildDefinitionResult',
    'get_build_definition',
    'get_build_definition_output',
]

@pulumi.output_type
class GetBuildDefinitionResult:
    """
    A collection of values returned by getBuildDefinition.
    """
    def __init__(__self__, agent_pool_name=None, ci_triggers=None, id=None, name=None, path=None, project_id=None, pull_request_triggers=None, queue_status=None, repositories=None, revision=None, schedules=None, variable_groups=None, variables=None):
        if agent_pool_name and not isinstance(agent_pool_name, str):
            raise TypeError("Expected argument 'agent_pool_name' to be a str")
        pulumi.set(__self__, "agent_pool_name", agent_pool_name)
        if ci_triggers and not isinstance(ci_triggers, list):
            raise TypeError("Expected argument 'ci_triggers' to be a list")
        pulumi.set(__self__, "ci_triggers", ci_triggers)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if pull_request_triggers and not isinstance(pull_request_triggers, list):
            raise TypeError("Expected argument 'pull_request_triggers' to be a list")
        pulumi.set(__self__, "pull_request_triggers", pull_request_triggers)
        if queue_status and not isinstance(queue_status, str):
            raise TypeError("Expected argument 'queue_status' to be a str")
        pulumi.set(__self__, "queue_status", queue_status)
        if repositories and not isinstance(repositories, list):
            raise TypeError("Expected argument 'repositories' to be a list")
        pulumi.set(__self__, "repositories", repositories)
        if revision and not isinstance(revision, int):
            raise TypeError("Expected argument 'revision' to be a int")
        pulumi.set(__self__, "revision", revision)
        if schedules and not isinstance(schedules, list):
            raise TypeError("Expected argument 'schedules' to be a list")
        pulumi.set(__self__, "schedules", schedules)
        if variable_groups and not isinstance(variable_groups, list):
            raise TypeError("Expected argument 'variable_groups' to be a list")
        pulumi.set(__self__, "variable_groups", variable_groups)
        if variables and not isinstance(variables, list):
            raise TypeError("Expected argument 'variables' to be a list")
        pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> str:
        """
        The agent pool that should execute the build.
        """
        return pulumi.get(self, "agent_pool_name")

    @property
    @pulumi.getter(name="ciTriggers")
    def ci_triggers(self) -> Sequence['outputs.GetBuildDefinitionCiTriggerResult']:
        """
        A `ci_trigger` block as defined below.
        """
        return pulumi.get(self, "ci_triggers")

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
        The name of the variable.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="pullRequestTriggers")
    def pull_request_triggers(self) -> Sequence['outputs.GetBuildDefinitionPullRequestTriggerResult']:
        """
        A `pull_request_trigger` block as defined below.
        """
        return pulumi.get(self, "pull_request_triggers")

    @property
    @pulumi.getter(name="queueStatus")
    def queue_status(self) -> str:
        """
        The queue status of the build definition.
        """
        return pulumi.get(self, "queue_status")

    @property
    @pulumi.getter
    def repositories(self) -> Sequence['outputs.GetBuildDefinitionRepositoryResult']:
        """
        A `repository` block as defined below.
        """
        return pulumi.get(self, "repositories")

    @property
    @pulumi.getter
    def revision(self) -> int:
        """
        The revision of the build definition.
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def schedules(self) -> Sequence['outputs.GetBuildDefinitionScheduleResult']:
        """
        A `schedules` block as defined below.
        """
        return pulumi.get(self, "schedules")

    @property
    @pulumi.getter(name="variableGroups")
    def variable_groups(self) -> Sequence[int]:
        """
        A list of variable group IDs.
        """
        return pulumi.get(self, "variable_groups")

    @property
    @pulumi.getter
    def variables(self) -> Sequence['outputs.GetBuildDefinitionVariableResult']:
        """
        A `variable` block as defined below.
        """
        return pulumi.get(self, "variables")


class AwaitableGetBuildDefinitionResult(GetBuildDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBuildDefinitionResult(
            agent_pool_name=self.agent_pool_name,
            ci_triggers=self.ci_triggers,
            id=self.id,
            name=self.name,
            path=self.path,
            project_id=self.project_id,
            pull_request_triggers=self.pull_request_triggers,
            queue_status=self.queue_status,
            repositories=self.repositories,
            revision=self.revision,
            schedules=self.schedules,
            variable_groups=self.variable_groups,
            variables=self.variables)


def get_build_definition(name: Optional[str] = None,
                         path: Optional[str] = None,
                         project_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBuildDefinitionResult:
    """
    Use this data source to access information about an existing Build Definition.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    example_get_build_definition = azuredevops.get_build_definition(project_id=example.id,
        name="existing")
    pulumi.export("id", example_get_build_definition.id)
    ```


    :param str name: The name of this Build Definition.
    :param str path: The path of the build definition. Default to `\\`.
    :param str project_id: The ID of the project.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['path'] = path
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getBuildDefinition:getBuildDefinition', __args__, opts=opts, typ=GetBuildDefinitionResult).value

    return AwaitableGetBuildDefinitionResult(
        agent_pool_name=pulumi.get(__ret__, 'agent_pool_name'),
        ci_triggers=pulumi.get(__ret__, 'ci_triggers'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        path=pulumi.get(__ret__, 'path'),
        project_id=pulumi.get(__ret__, 'project_id'),
        pull_request_triggers=pulumi.get(__ret__, 'pull_request_triggers'),
        queue_status=pulumi.get(__ret__, 'queue_status'),
        repositories=pulumi.get(__ret__, 'repositories'),
        revision=pulumi.get(__ret__, 'revision'),
        schedules=pulumi.get(__ret__, 'schedules'),
        variable_groups=pulumi.get(__ret__, 'variable_groups'),
        variables=pulumi.get(__ret__, 'variables'))


@_utilities.lift_output_func(get_build_definition)
def get_build_definition_output(name: Optional[pulumi.Input[str]] = None,
                                path: Optional[pulumi.Input[Optional[str]]] = None,
                                project_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBuildDefinitionResult]:
    """
    Use this data source to access information about an existing Build Definition.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example = azuredevops.get_project(name="Example Project")
    example_get_build_definition = azuredevops.get_build_definition(project_id=example.id,
        name="existing")
    pulumi.export("id", example_get_build_definition.id)
    ```


    :param str name: The name of this Build Definition.
    :param str path: The path of the build definition. Default to `\\`.
    :param str project_id: The ID of the project.
    """
    ...
