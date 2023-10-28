# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PipelineAuthorizationArgs', 'PipelineAuthorization']

@pulumi.input_type
class PipelineAuthorizationArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 pipeline_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PipelineAuthorization resource.
        :param pulumi.Input[str] project_id: The  ID of the project. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Changing this forces a new resource to be created
        :param pulumi.Input[str] type: The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        :param pulumi.Input[int] pipeline_id: The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "type", type)
        if pipeline_id is not None:
            pulumi.set(__self__, "pipeline_id", pipeline_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The  ID of the project. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource to authorize. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pipeline_id", value)


@pulumi.input_type
class _PipelineAuthorizationState:
    def __init__(__self__, *,
                 pipeline_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PipelineAuthorization resources.
        :param pulumi.Input[int] pipeline_id: The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        :param pulumi.Input[str] project_id: The  ID of the project. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Changing this forces a new resource to be created
        :param pulumi.Input[str] type: The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        """
        if pipeline_id is not None:
            pulumi.set(__self__, "pipeline_id", pipeline_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The  ID of the project. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource to authorize. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class PipelineAuthorization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pipeline_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage pipeline access permissions to resources.

        > **Note** This resource is a replacement for `ResourceAuthorization`.  Pipeline authorizations managed by `ResourceAuthorization` can also be managed by this resource.

        > **Note** If both "All Pipeline Authorization" and "Custom Pipeline Authorization" are configured, "All Pipeline Authorization" has higher priority.

        ## Example Usage
        ### Authorization for all pipelines

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_pool = azuredevops.Pool("examplePool",
            auto_provision=False,
            auto_update=False)
        example_queue = azuredevops.Queue("exampleQueue",
            project_id=example_project.id,
            agent_pool_id=example_pool.id)
        example_pipeline_authorization = azuredevops.PipelineAuthorization("examplePipelineAuthorization",
            project_id=example_project.id,
            resource_id=example_queue.id,
            type="queue")
        ```
        ### Authorization for specific pipeline

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_pool = azuredevops.Pool("examplePool",
            auto_provision=False,
            auto_update=False)
        example_queue = azuredevops.Queue("exampleQueue",
            project_id=example_project.id,
            agent_pool_id=example_pool.id)
        example_git_repository = azuredevops.get_git_repository_output(project_id=example_project.id,
            name="Example Project")
        example_build_definition = azuredevops.BuildDefinition("exampleBuildDefinition",
            project_id=example_project.id,
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="TfsGit",
                repo_id=example_git_repository.id,
                yml_path="azure-pipelines.yml",
            ))
        example_pipeline_authorization = azuredevops.PipelineAuthorization("examplePipelineAuthorization",
            project_id=example_project.id,
            resource_id=example_queue.id,
            type="queue",
            pipeline_id=example_build_definition.id)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.1 - Pipeline Permissions](https://learn.microsoft.com/en-us/rest/api/azure/devops/approvalsandchecks/pipeline-permissions?view=azure-devops-rest-7.1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] pipeline_id: The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        :param pulumi.Input[str] project_id: The  ID of the project. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Changing this forces a new resource to be created
        :param pulumi.Input[str] type: The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineAuthorizationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage pipeline access permissions to resources.

        > **Note** This resource is a replacement for `ResourceAuthorization`.  Pipeline authorizations managed by `ResourceAuthorization` can also be managed by this resource.

        > **Note** If both "All Pipeline Authorization" and "Custom Pipeline Authorization" are configured, "All Pipeline Authorization" has higher priority.

        ## Example Usage
        ### Authorization for all pipelines

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_pool = azuredevops.Pool("examplePool",
            auto_provision=False,
            auto_update=False)
        example_queue = azuredevops.Queue("exampleQueue",
            project_id=example_project.id,
            agent_pool_id=example_pool.id)
        example_pipeline_authorization = azuredevops.PipelineAuthorization("examplePipelineAuthorization",
            project_id=example_project.id,
            resource_id=example_queue.id,
            type="queue")
        ```
        ### Authorization for specific pipeline

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_pool = azuredevops.Pool("examplePool",
            auto_provision=False,
            auto_update=False)
        example_queue = azuredevops.Queue("exampleQueue",
            project_id=example_project.id,
            agent_pool_id=example_pool.id)
        example_git_repository = azuredevops.get_git_repository_output(project_id=example_project.id,
            name="Example Project")
        example_build_definition = azuredevops.BuildDefinition("exampleBuildDefinition",
            project_id=example_project.id,
            repository=azuredevops.BuildDefinitionRepositoryArgs(
                repo_type="TfsGit",
                repo_id=example_git_repository.id,
                yml_path="azure-pipelines.yml",
            ))
        example_pipeline_authorization = azuredevops.PipelineAuthorization("examplePipelineAuthorization",
            project_id=example_project.id,
            resource_id=example_queue.id,
            type="queue",
            pipeline_id=example_build_definition.id)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.1 - Pipeline Permissions](https://learn.microsoft.com/en-us/rest/api/azure/devops/approvalsandchecks/pipeline-permissions?view=azure-devops-rest-7.1)

        :param str resource_name: The name of the resource.
        :param PipelineAuthorizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineAuthorizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pipeline_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineAuthorizationArgs.__new__(PipelineAuthorizationArgs)

            __props__.__dict__["pipeline_id"] = pipeline_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(PipelineAuthorization, __self__).__init__(
            'azuredevops:index/pipelineAuthorization:PipelineAuthorization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            pipeline_id: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'PipelineAuthorization':
        """
        Get an existing PipelineAuthorization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] pipeline_id: The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        :param pulumi.Input[str] project_id: The  ID of the project. Changing this forces a new resource to be created
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Changing this forces a new resource to be created
        :param pulumi.Input[str] type: The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineAuthorizationState.__new__(_PipelineAuthorizationState)

        __props__.__dict__["pipeline_id"] = pipeline_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["type"] = type
        return PipelineAuthorization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Output[Optional[int]]:
        """
        The ID of the pipeline. If not configured, all pipelines will be authorized. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The  ID of the project. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource to authorize. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource to authorize. Valid values: `endpoint`, `queue`, `variablegroup`, `environment`. Changing this forces a new resource to be created
        """
        return pulumi.get(self, "type")

