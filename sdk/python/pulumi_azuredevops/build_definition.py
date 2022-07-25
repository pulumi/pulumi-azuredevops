# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BuildDefinitionArgs', 'BuildDefinition']

@pulumi.input_type
class BuildDefinitionArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 repository: pulumi.Input['BuildDefinitionRepositoryArgs'],
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 ci_trigger: Optional[pulumi.Input['BuildDefinitionCiTriggerArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 pull_request_trigger: Optional[pulumi.Input['BuildDefinitionPullRequestTriggerArgs']] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleArgs']]]] = None,
                 variable_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]]] = None):
        """
        The set of arguments for constructing a BuildDefinition resource.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input['BuildDefinitionRepositoryArgs'] repository: A `repository` block as documented below.
        :param pulumi.Input[str] agent_pool_name: The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        :param pulumi.Input['BuildDefinitionCiTriggerArgs'] ci_trigger: Continuous Integration trigger.
        :param pulumi.Input[str] name: The name of the build definition.
        :param pulumi.Input[str] path: The folder path of the build definition.
        :param pulumi.Input['BuildDefinitionPullRequestTriggerArgs'] pull_request_trigger: Pull Request Integration Integration trigger.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] variable_groups: A list of variable group IDs (integers) to link to the build definition.
        :param pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]] variables: A list of `variable` blocks, as documented below.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "repository", repository)
        if agent_pool_name is not None:
            pulumi.set(__self__, "agent_pool_name", agent_pool_name)
        if ci_trigger is not None:
            pulumi.set(__self__, "ci_trigger", ci_trigger)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if pull_request_trigger is not None:
            pulumi.set(__self__, "pull_request_trigger", pull_request_trigger)
        if schedules is not None:
            pulumi.set(__self__, "schedules", schedules)
        if variable_groups is not None:
            pulumi.set(__self__, "variable_groups", variable_groups)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

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
    @pulumi.getter
    def repository(self) -> pulumi.Input['BuildDefinitionRepositoryArgs']:
        """
        A `repository` block as documented below.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input['BuildDefinitionRepositoryArgs']):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> Optional[pulumi.Input[str]]:
        """
        The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        """
        return pulumi.get(self, "agent_pool_name")

    @agent_pool_name.setter
    def agent_pool_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_pool_name", value)

    @property
    @pulumi.getter(name="ciTrigger")
    def ci_trigger(self) -> Optional[pulumi.Input['BuildDefinitionCiTriggerArgs']]:
        """
        Continuous Integration trigger.
        """
        return pulumi.get(self, "ci_trigger")

    @ci_trigger.setter
    def ci_trigger(self, value: Optional[pulumi.Input['BuildDefinitionCiTriggerArgs']]):
        pulumi.set(self, "ci_trigger", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the build definition.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The folder path of the build definition.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="pullRequestTrigger")
    def pull_request_trigger(self) -> Optional[pulumi.Input['BuildDefinitionPullRequestTriggerArgs']]:
        """
        Pull Request Integration Integration trigger.
        """
        return pulumi.get(self, "pull_request_trigger")

    @pull_request_trigger.setter
    def pull_request_trigger(self, value: Optional[pulumi.Input['BuildDefinitionPullRequestTriggerArgs']]):
        pulumi.set(self, "pull_request_trigger", value)

    @property
    @pulumi.getter
    def schedules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleArgs']]]]:
        return pulumi.get(self, "schedules")

    @schedules.setter
    def schedules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleArgs']]]]):
        pulumi.set(self, "schedules", value)

    @property
    @pulumi.getter(name="variableGroups")
    def variable_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        A list of variable group IDs (integers) to link to the build definition.
        """
        return pulumi.get(self, "variable_groups")

    @variable_groups.setter
    def variable_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "variable_groups", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]]]:
        """
        A list of `variable` blocks, as documented below.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]]]):
        pulumi.set(self, "variables", value)


@pulumi.input_type
class _BuildDefinitionState:
    def __init__(__self__, *,
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 ci_trigger: Optional[pulumi.Input['BuildDefinitionCiTriggerArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 pull_request_trigger: Optional[pulumi.Input['BuildDefinitionPullRequestTriggerArgs']] = None,
                 repository: Optional[pulumi.Input['BuildDefinitionRepositoryArgs']] = None,
                 revision: Optional[pulumi.Input[int]] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleArgs']]]] = None,
                 variable_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]]] = None):
        """
        Input properties used for looking up and filtering BuildDefinition resources.
        :param pulumi.Input[str] agent_pool_name: The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        :param pulumi.Input['BuildDefinitionCiTriggerArgs'] ci_trigger: Continuous Integration trigger.
        :param pulumi.Input[str] name: The name of the build definition.
        :param pulumi.Input[str] path: The folder path of the build definition.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input['BuildDefinitionPullRequestTriggerArgs'] pull_request_trigger: Pull Request Integration Integration trigger.
        :param pulumi.Input['BuildDefinitionRepositoryArgs'] repository: A `repository` block as documented below.
        :param pulumi.Input[int] revision: The revision of the build definition
        :param pulumi.Input[Sequence[pulumi.Input[int]]] variable_groups: A list of variable group IDs (integers) to link to the build definition.
        :param pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]] variables: A list of `variable` blocks, as documented below.
        """
        if agent_pool_name is not None:
            pulumi.set(__self__, "agent_pool_name", agent_pool_name)
        if ci_trigger is not None:
            pulumi.set(__self__, "ci_trigger", ci_trigger)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if pull_request_trigger is not None:
            pulumi.set(__self__, "pull_request_trigger", pull_request_trigger)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)
        if schedules is not None:
            pulumi.set(__self__, "schedules", schedules)
        if variable_groups is not None:
            pulumi.set(__self__, "variable_groups", variable_groups)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> Optional[pulumi.Input[str]]:
        """
        The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        """
        return pulumi.get(self, "agent_pool_name")

    @agent_pool_name.setter
    def agent_pool_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_pool_name", value)

    @property
    @pulumi.getter(name="ciTrigger")
    def ci_trigger(self) -> Optional[pulumi.Input['BuildDefinitionCiTriggerArgs']]:
        """
        Continuous Integration trigger.
        """
        return pulumi.get(self, "ci_trigger")

    @ci_trigger.setter
    def ci_trigger(self, value: Optional[pulumi.Input['BuildDefinitionCiTriggerArgs']]):
        pulumi.set(self, "ci_trigger", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the build definition.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The folder path of the build definition.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

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
    @pulumi.getter(name="pullRequestTrigger")
    def pull_request_trigger(self) -> Optional[pulumi.Input['BuildDefinitionPullRequestTriggerArgs']]:
        """
        Pull Request Integration Integration trigger.
        """
        return pulumi.get(self, "pull_request_trigger")

    @pull_request_trigger.setter
    def pull_request_trigger(self, value: Optional[pulumi.Input['BuildDefinitionPullRequestTriggerArgs']]):
        pulumi.set(self, "pull_request_trigger", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input['BuildDefinitionRepositoryArgs']]:
        """
        A `repository` block as documented below.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input['BuildDefinitionRepositoryArgs']]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[int]]:
        """
        The revision of the build definition
        """
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision", value)

    @property
    @pulumi.getter
    def schedules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleArgs']]]]:
        return pulumi.get(self, "schedules")

    @schedules.setter
    def schedules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleArgs']]]]):
        pulumi.set(self, "schedules", value)

    @property
    @pulumi.getter(name="variableGroups")
    def variable_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        A list of variable group IDs (integers) to link to the build definition.
        """
        return pulumi.get(self, "variable_groups")

    @variable_groups.setter
    def variable_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "variable_groups", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]]]:
        """
        A list of `variable` blocks, as documented below.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionVariableArgs']]]]):
        pulumi.set(self, "variables", value)


class BuildDefinition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 ci_trigger: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 pull_request_trigger: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']]] = None,
                 repository: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']]] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionScheduleArgs']]]]] = None,
                 variable_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]]] = None,
                 __props__=None):
        """
        Manages a Build Definition within Azure DevOps.

        ## Example Usage
        ### Tfs
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example_git = azuredevops.Git("exampleGit",
            project_id=example_project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        example_variable_group = azuredevops.VariableGroup("exampleVariableGroup",
            project_id=example_project.id,
            description="Managed by Terraform",
            allow_access=True,
            variables=[azuredevops.VariableGroupVariableArgs(
                name="FOO",
                value="BAR",
            )])
        example_build_definition = azuredevops.BuildDefinition("exampleBuildDefinition",
            project_id=example_project.id,
            path="\\\\ExampleFolder",
            ci_trigger=azuredevops.BuildDefinitionCiTriggerArgs(
                use_yaml=True,
            ),
            schedules=[azuredevops.BuildDefinitionScheduleArgs(
                branch_filters=[azuredevops.BuildDefinitionScheduleBranchFilterArgs(
                    includes=["master"],
                    excludes=[
                        "test",
                        "regression",
                    ],
                )],
                days_to_builds=[
                    "Wed",
                    "Sun",
                ],
                schedule_only_with_changes=True,
                start_hours=10,
                start_minutes=59,
                time_zone="(UTC) Coordinated Universal Time",
            )],
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="TfsGit",
                repo_id=example_git.id,
                branch_name=example_git.default_branch,
                yml_path="azure-pipelines.yml",
            ),
            variable_groups=[example_variable_group.id],
            variables=[
                azuredevops.BuildDefinitionVariableArgs(
                    name="PipelineVariable",
                    value="Go Microsoft!",
                ),
                azuredevops.BuildDefinitionVariableArgs(
                    name="PipelineSecret",
                    secret_value="ZGV2cw",
                    is_secret=True,
                ),
            ])
        ```
        ### GitHub Enterprise
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example_service_endpoint_git_hub_enterprise = azuredevops.ServiceEndpointGitHubEnterprise("exampleServiceEndpointGitHubEnterprise",
            project_id=example_project.id,
            service_endpoint_name="Example GitHub Enterprise",
            url="https://github.contoso.com",
            description="Managed by Terraform",
            auth_personal=azuredevops.ServiceEndpointGitHubEnterpriseAuthPersonalArgs(
                personal_access_token="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            ))
        example_build_definition = azuredevops.BuildDefinition("exampleBuildDefinition",
            project_id=example_project.id,
            path="\\\\ExampleFolder",
            ci_trigger=azuredevops.BuildDefinitionCiTriggerArgs(
                use_yaml=True,
            ),
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="GitHubEnterprise",
                repo_id="<GitHub Org>/<Repo Name>",
                github_enterprise_url="https://github.company.com",
                branch_name="master",
                yml_path="azure-pipelines.yml",
                service_connection_id=example_service_endpoint_git_hub_enterprise.id,
            ),
            schedules=[azuredevops.BuildDefinitionScheduleArgs(
                branch_filters=[azuredevops.BuildDefinitionScheduleBranchFilterArgs(
                    includes=["main"],
                    excludes=[
                        "test",
                        "regression",
                    ],
                )],
                days_to_builds=[
                    "Wed",
                    "Sun",
                ],
                schedule_only_with_changes=True,
                start_hours=10,
                start_minutes=59,
                time_zone="(UTC) Coordinated Universal Time",
            )])
        ```
        ## Remarks

        The path attribute can not end in `\\` unless the path is the root value of `\\`.

        Valid path values (yaml encoded) include:
        - `\\\\`
        - `\\\\ExampleFolder`
        - `\\\\Nested\\\\Example Folder`

        The value of `\\\\ExampleFolder\\\\` would be invalid.

        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-6.0)

        ## Import

        Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.

        ```sh
         $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example "Example Project"/10
        ```

         or

        ```sh
         $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_pool_name: The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']] ci_trigger: Continuous Integration trigger.
        :param pulumi.Input[str] name: The name of the build definition.
        :param pulumi.Input[str] path: The folder path of the build definition.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']] pull_request_trigger: Pull Request Integration Integration trigger.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']] repository: A `repository` block as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] variable_groups: A list of variable group IDs (integers) to link to the build definition.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]] variables: A list of `variable` blocks, as documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BuildDefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Build Definition within Azure DevOps.

        ## Example Usage
        ### Tfs
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example_git = azuredevops.Git("exampleGit",
            project_id=example_project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        example_variable_group = azuredevops.VariableGroup("exampleVariableGroup",
            project_id=example_project.id,
            description="Managed by Terraform",
            allow_access=True,
            variables=[azuredevops.VariableGroupVariableArgs(
                name="FOO",
                value="BAR",
            )])
        example_build_definition = azuredevops.BuildDefinition("exampleBuildDefinition",
            project_id=example_project.id,
            path="\\\\ExampleFolder",
            ci_trigger=azuredevops.BuildDefinitionCiTriggerArgs(
                use_yaml=True,
            ),
            schedules=[azuredevops.BuildDefinitionScheduleArgs(
                branch_filters=[azuredevops.BuildDefinitionScheduleBranchFilterArgs(
                    includes=["master"],
                    excludes=[
                        "test",
                        "regression",
                    ],
                )],
                days_to_builds=[
                    "Wed",
                    "Sun",
                ],
                schedule_only_with_changes=True,
                start_hours=10,
                start_minutes=59,
                time_zone="(UTC) Coordinated Universal Time",
            )],
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="TfsGit",
                repo_id=example_git.id,
                branch_name=example_git.default_branch,
                yml_path="azure-pipelines.yml",
            ),
            variable_groups=[example_variable_group.id],
            variables=[
                azuredevops.BuildDefinitionVariableArgs(
                    name="PipelineVariable",
                    value="Go Microsoft!",
                ),
                azuredevops.BuildDefinitionVariableArgs(
                    name="PipelineSecret",
                    secret_value="ZGV2cw",
                    is_secret=True,
                ),
            ])
        ```
        ### GitHub Enterprise
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example_service_endpoint_git_hub_enterprise = azuredevops.ServiceEndpointGitHubEnterprise("exampleServiceEndpointGitHubEnterprise",
            project_id=example_project.id,
            service_endpoint_name="Example GitHub Enterprise",
            url="https://github.contoso.com",
            description="Managed by Terraform",
            auth_personal=azuredevops.ServiceEndpointGitHubEnterpriseAuthPersonalArgs(
                personal_access_token="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            ))
        example_build_definition = azuredevops.BuildDefinition("exampleBuildDefinition",
            project_id=example_project.id,
            path="\\\\ExampleFolder",
            ci_trigger=azuredevops.BuildDefinitionCiTriggerArgs(
                use_yaml=True,
            ),
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="GitHubEnterprise",
                repo_id="<GitHub Org>/<Repo Name>",
                github_enterprise_url="https://github.company.com",
                branch_name="master",
                yml_path="azure-pipelines.yml",
                service_connection_id=example_service_endpoint_git_hub_enterprise.id,
            ),
            schedules=[azuredevops.BuildDefinitionScheduleArgs(
                branch_filters=[azuredevops.BuildDefinitionScheduleBranchFilterArgs(
                    includes=["main"],
                    excludes=[
                        "test",
                        "regression",
                    ],
                )],
                days_to_builds=[
                    "Wed",
                    "Sun",
                ],
                schedule_only_with_changes=True,
                start_hours=10,
                start_minutes=59,
                time_zone="(UTC) Coordinated Universal Time",
            )])
        ```
        ## Remarks

        The path attribute can not end in `\\` unless the path is the root value of `\\`.

        Valid path values (yaml encoded) include:
        - `\\\\`
        - `\\\\ExampleFolder`
        - `\\\\Nested\\\\Example Folder`

        The value of `\\\\ExampleFolder\\\\` would be invalid.

        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-6.0)

        ## Import

        Azure DevOps Build Definitions can be imported using the project name/definitions Id or by the project Guid/definitions Id, e.g.

        ```sh
         $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example "Example Project"/10
        ```

         or

        ```sh
         $ pulumi import azuredevops:index/buildDefinition:BuildDefinition example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param BuildDefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BuildDefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 ci_trigger: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 pull_request_trigger: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']]] = None,
                 repository: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']]] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionScheduleArgs']]]]] = None,
                 variable_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BuildDefinitionArgs.__new__(BuildDefinitionArgs)

            __props__.__dict__["agent_pool_name"] = agent_pool_name
            __props__.__dict__["ci_trigger"] = ci_trigger
            __props__.__dict__["name"] = name
            __props__.__dict__["path"] = path
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["pull_request_trigger"] = pull_request_trigger
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["schedules"] = schedules
            __props__.__dict__["variable_groups"] = variable_groups
            __props__.__dict__["variables"] = variables
            __props__.__dict__["revision"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Build/buildDefinition:BuildDefinition")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BuildDefinition, __self__).__init__(
            'azuredevops:index/buildDefinition:BuildDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_pool_name: Optional[pulumi.Input[str]] = None,
            ci_trigger: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            pull_request_trigger: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']]] = None,
            repository: Optional[pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']]] = None,
            revision: Optional[pulumi.Input[int]] = None,
            schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionScheduleArgs']]]]] = None,
            variable_groups: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]]] = None) -> 'BuildDefinition':
        """
        Get an existing BuildDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_pool_name: The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']] ci_trigger: Continuous Integration trigger.
        :param pulumi.Input[str] name: The name of the build definition.
        :param pulumi.Input[str] path: The folder path of the build definition.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']] pull_request_trigger: Pull Request Integration Integration trigger.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']] repository: A `repository` block as documented below.
        :param pulumi.Input[int] revision: The revision of the build definition
        :param pulumi.Input[Sequence[pulumi.Input[int]]] variable_groups: A list of variable group IDs (integers) to link to the build definition.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]] variables: A list of `variable` blocks, as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BuildDefinitionState.__new__(_BuildDefinitionState)

        __props__.__dict__["agent_pool_name"] = agent_pool_name
        __props__.__dict__["ci_trigger"] = ci_trigger
        __props__.__dict__["name"] = name
        __props__.__dict__["path"] = path
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["pull_request_trigger"] = pull_request_trigger
        __props__.__dict__["repository"] = repository
        __props__.__dict__["revision"] = revision
        __props__.__dict__["schedules"] = schedules
        __props__.__dict__["variable_groups"] = variable_groups
        __props__.__dict__["variables"] = variables
        return BuildDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> pulumi.Output[Optional[str]]:
        """
        The agent pool that should execute the build. Defaults to `Azure Pipelines`.
        """
        return pulumi.get(self, "agent_pool_name")

    @property
    @pulumi.getter(name="ciTrigger")
    def ci_trigger(self) -> pulumi.Output[Optional['outputs.BuildDefinitionCiTrigger']]:
        """
        Continuous Integration trigger.
        """
        return pulumi.get(self, "ci_trigger")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the build definition.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The folder path of the build definition.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="pullRequestTrigger")
    def pull_request_trigger(self) -> pulumi.Output[Optional['outputs.BuildDefinitionPullRequestTrigger']]:
        """
        Pull Request Integration Integration trigger.
        """
        return pulumi.get(self, "pull_request_trigger")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output['outputs.BuildDefinitionRepository']:
        """
        A `repository` block as documented below.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Output[int]:
        """
        The revision of the build definition
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def schedules(self) -> pulumi.Output[Optional[Sequence['outputs.BuildDefinitionSchedule']]]:
        return pulumi.get(self, "schedules")

    @property
    @pulumi.getter(name="variableGroups")
    def variable_groups(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        A list of variable group IDs (integers) to link to the build definition.
        """
        return pulumi.get(self, "variable_groups")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Sequence['outputs.BuildDefinitionVariable']]]:
        """
        A list of `variable` blocks, as documented below.
        """
        return pulumi.get(self, "variables")

