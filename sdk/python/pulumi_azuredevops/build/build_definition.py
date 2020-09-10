# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BuildDefinition']

warnings.warn("azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition", DeprecationWarning)


class BuildDefinition(pulumi.CustomResource):
    warnings.warn("azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition", DeprecationWarning)

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
                 variable_groups: Optional[pulumi.Input[List[pulumi.Input[float]]]] = None,
                 variables: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Build Definition within Azure DevOps.

        ## Example Usage
        ### Tfs
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        repository = azuredevops.Git("repository",
            project_id=project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        vars = azuredevops.VariableGroup("vars",
            project_id=project.id,
            description="Managed by Terraform",
            allow_access=True,
            variables=[azuredevops.VariableGroupVariableArgs(
                name="FOO",
                value="BAR",
            )])
        build = azuredevops.BuildDefinition("build",
            project_id=project.id,
            path="\\ExampleFolder",
            ci_trigger=azuredevops.BuildDefinitionCiTriggerArgs(
                use_yaml=True,
            ),
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="TfsGit",
                repo_id=repository.id,
                branch_name=repository.default_branch,
                yml_path="azure-pipelines.yml",
            ),
            variable_groups=[vars.id],
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

        sample_dotnetcore_app_release = azuredevops.BuildDefinition("sampleDotnetcoreAppRelease",
            project_id=azuredevops_project["project"]["id"],
            path="\\ExampleFolder",
            ci_trigger=azuredevops.BuildDefinitionCiTriggerArgs(
                use_yaml=True,
            ),
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="GitHubEnterprise",
                repo_id="<GitHub Org>/<Repo Name>",
                github_enterprise_url="https://github.company.com",
                branch_name="master",
                yml_path="azure-pipelines.yml",
                service_connection_id="...",
            ))
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Build Definitions](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/definitions?view=azure-devops-rest-5.1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_pool_name: The agent pool that should execute the build.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']] ci_trigger: Continuous Integration trigger.
        :param pulumi.Input[str] name: The name of the build definition.
        :param pulumi.Input[str] path: The folder path of the build definition.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']] pull_request_trigger: Pull Request Integration Integration trigger.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']] repository: A `repository` block as documented below.
        :param pulumi.Input[List[pulumi.Input[float]]] variable_groups: A list of variable group IDs (integers) to link to the build definition.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]] variables: A list of `variable` blocks, as documented below.
        """
        pulumi.log.warn("BuildDefinition is deprecated: azuredevops.build.BuildDefinition has been deprecated in favor of azuredevops.BuildDefinition")
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
            __props__ = dict()

            __props__['agent_pool_name'] = agent_pool_name
            __props__['ci_trigger'] = ci_trigger
            __props__['name'] = name
            __props__['path'] = path
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['pull_request_trigger'] = pull_request_trigger
            if repository is None:
                raise TypeError("Missing required property 'repository'")
            __props__['repository'] = repository
            __props__['variable_groups'] = variable_groups
            __props__['variables'] = variables
            __props__['revision'] = None
        super(BuildDefinition, __self__).__init__(
            'azuredevops:Build/buildDefinition:BuildDefinition',
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
            revision: Optional[pulumi.Input[float]] = None,
            variable_groups: Optional[pulumi.Input[List[pulumi.Input[float]]]] = None,
            variables: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]]] = None) -> 'BuildDefinition':
        """
        Get an existing BuildDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_pool_name: The agent pool that should execute the build.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionCiTriggerArgs']] ci_trigger: Continuous Integration trigger.
        :param pulumi.Input[str] name: The name of the build definition.
        :param pulumi.Input[str] path: The folder path of the build definition.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionPullRequestTriggerArgs']] pull_request_trigger: Pull Request Integration Integration trigger.
        :param pulumi.Input[pulumi.InputType['BuildDefinitionRepositoryArgs']] repository: A `repository` block as documented below.
        :param pulumi.Input[float] revision: The revision of the build definition
        :param pulumi.Input[List[pulumi.Input[float]]] variable_groups: A list of variable group IDs (integers) to link to the build definition.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['BuildDefinitionVariableArgs']]]] variables: A list of `variable` blocks, as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["agent_pool_name"] = agent_pool_name
        __props__["ci_trigger"] = ci_trigger
        __props__["name"] = name
        __props__["path"] = path
        __props__["project_id"] = project_id
        __props__["pull_request_trigger"] = pull_request_trigger
        __props__["repository"] = repository
        __props__["revision"] = revision
        __props__["variable_groups"] = variable_groups
        __props__["variables"] = variables
        return BuildDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPoolName")
    def agent_pool_name(self) -> pulumi.Output[Optional[str]]:
        """
        The agent pool that should execute the build.
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
    def revision(self) -> pulumi.Output[float]:
        """
        The revision of the build definition
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter(name="variableGroups")
    def variable_groups(self) -> pulumi.Output[Optional[List[float]]]:
        """
        A list of variable group IDs (integers) to link to the build definition.
        """
        return pulumi.get(self, "variable_groups")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[List['outputs.BuildDefinitionVariable']]]:
        """
        A list of `variable` blocks, as documented below.
        """
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

