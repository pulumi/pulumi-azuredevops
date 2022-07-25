# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectPipelineSettingsArgs', 'ProjectPipelineSettings']

@pulumi.input_type
class ProjectPipelineSettingsArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 enforce_job_scope: Optional[pulumi.Input[bool]] = None,
                 enforce_referenced_repo_scoped_token: Optional[pulumi.Input[bool]] = None,
                 enforce_settable_var: Optional[pulumi.Input[bool]] = None,
                 publish_pipeline_metadata: Optional[pulumi.Input[bool]] = None,
                 status_badges_are_private: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ProjectPipelineSettings resource.
        :param pulumi.Input[str] project_id: The `id` of the project for which the project pipeline settings will be managed.
        :param pulumi.Input[bool] enforce_job_scope: Limit job authorization scope to current project for non-release pipelines.
        :param pulumi.Input[bool] enforce_referenced_repo_scoped_token: Protect access to repositories in YAML pipelines.
        :param pulumi.Input[bool] enforce_settable_var: Limit variables that can be set at queue time.
        :param pulumi.Input[bool] publish_pipeline_metadata: Publish metadata from pipelines.
        :param pulumi.Input[bool] status_badges_are_private: Disable anonymous access to badges.
        """
        pulumi.set(__self__, "project_id", project_id)
        if enforce_job_scope is not None:
            pulumi.set(__self__, "enforce_job_scope", enforce_job_scope)
        if enforce_referenced_repo_scoped_token is not None:
            pulumi.set(__self__, "enforce_referenced_repo_scoped_token", enforce_referenced_repo_scoped_token)
        if enforce_settable_var is not None:
            pulumi.set(__self__, "enforce_settable_var", enforce_settable_var)
        if publish_pipeline_metadata is not None:
            pulumi.set(__self__, "publish_pipeline_metadata", publish_pipeline_metadata)
        if status_badges_are_private is not None:
            pulumi.set(__self__, "status_badges_are_private", status_badges_are_private)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The `id` of the project for which the project pipeline settings will be managed.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="enforceJobScope")
    def enforce_job_scope(self) -> Optional[pulumi.Input[bool]]:
        """
        Limit job authorization scope to current project for non-release pipelines.
        """
        return pulumi.get(self, "enforce_job_scope")

    @enforce_job_scope.setter
    def enforce_job_scope(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_job_scope", value)

    @property
    @pulumi.getter(name="enforceReferencedRepoScopedToken")
    def enforce_referenced_repo_scoped_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Protect access to repositories in YAML pipelines.
        """
        return pulumi.get(self, "enforce_referenced_repo_scoped_token")

    @enforce_referenced_repo_scoped_token.setter
    def enforce_referenced_repo_scoped_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_referenced_repo_scoped_token", value)

    @property
    @pulumi.getter(name="enforceSettableVar")
    def enforce_settable_var(self) -> Optional[pulumi.Input[bool]]:
        """
        Limit variables that can be set at queue time.
        """
        return pulumi.get(self, "enforce_settable_var")

    @enforce_settable_var.setter
    def enforce_settable_var(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_settable_var", value)

    @property
    @pulumi.getter(name="publishPipelineMetadata")
    def publish_pipeline_metadata(self) -> Optional[pulumi.Input[bool]]:
        """
        Publish metadata from pipelines.
        """
        return pulumi.get(self, "publish_pipeline_metadata")

    @publish_pipeline_metadata.setter
    def publish_pipeline_metadata(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_pipeline_metadata", value)

    @property
    @pulumi.getter(name="statusBadgesArePrivate")
    def status_badges_are_private(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable anonymous access to badges.
        """
        return pulumi.get(self, "status_badges_are_private")

    @status_badges_are_private.setter
    def status_badges_are_private(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status_badges_are_private", value)


@pulumi.input_type
class _ProjectPipelineSettingsState:
    def __init__(__self__, *,
                 enforce_job_scope: Optional[pulumi.Input[bool]] = None,
                 enforce_referenced_repo_scoped_token: Optional[pulumi.Input[bool]] = None,
                 enforce_settable_var: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 publish_pipeline_metadata: Optional[pulumi.Input[bool]] = None,
                 status_badges_are_private: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ProjectPipelineSettings resources.
        :param pulumi.Input[bool] enforce_job_scope: Limit job authorization scope to current project for non-release pipelines.
        :param pulumi.Input[bool] enforce_referenced_repo_scoped_token: Protect access to repositories in YAML pipelines.
        :param pulumi.Input[bool] enforce_settable_var: Limit variables that can be set at queue time.
        :param pulumi.Input[str] project_id: The `id` of the project for which the project pipeline settings will be managed.
        :param pulumi.Input[bool] publish_pipeline_metadata: Publish metadata from pipelines.
        :param pulumi.Input[bool] status_badges_are_private: Disable anonymous access to badges.
        """
        if enforce_job_scope is not None:
            pulumi.set(__self__, "enforce_job_scope", enforce_job_scope)
        if enforce_referenced_repo_scoped_token is not None:
            pulumi.set(__self__, "enforce_referenced_repo_scoped_token", enforce_referenced_repo_scoped_token)
        if enforce_settable_var is not None:
            pulumi.set(__self__, "enforce_settable_var", enforce_settable_var)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if publish_pipeline_metadata is not None:
            pulumi.set(__self__, "publish_pipeline_metadata", publish_pipeline_metadata)
        if status_badges_are_private is not None:
            pulumi.set(__self__, "status_badges_are_private", status_badges_are_private)

    @property
    @pulumi.getter(name="enforceJobScope")
    def enforce_job_scope(self) -> Optional[pulumi.Input[bool]]:
        """
        Limit job authorization scope to current project for non-release pipelines.
        """
        return pulumi.get(self, "enforce_job_scope")

    @enforce_job_scope.setter
    def enforce_job_scope(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_job_scope", value)

    @property
    @pulumi.getter(name="enforceReferencedRepoScopedToken")
    def enforce_referenced_repo_scoped_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Protect access to repositories in YAML pipelines.
        """
        return pulumi.get(self, "enforce_referenced_repo_scoped_token")

    @enforce_referenced_repo_scoped_token.setter
    def enforce_referenced_repo_scoped_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_referenced_repo_scoped_token", value)

    @property
    @pulumi.getter(name="enforceSettableVar")
    def enforce_settable_var(self) -> Optional[pulumi.Input[bool]]:
        """
        Limit variables that can be set at queue time.
        """
        return pulumi.get(self, "enforce_settable_var")

    @enforce_settable_var.setter
    def enforce_settable_var(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_settable_var", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The `id` of the project for which the project pipeline settings will be managed.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="publishPipelineMetadata")
    def publish_pipeline_metadata(self) -> Optional[pulumi.Input[bool]]:
        """
        Publish metadata from pipelines.
        """
        return pulumi.get(self, "publish_pipeline_metadata")

    @publish_pipeline_metadata.setter
    def publish_pipeline_metadata(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_pipeline_metadata", value)

    @property
    @pulumi.getter(name="statusBadgesArePrivate")
    def status_badges_are_private(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable anonymous access to badges.
        """
        return pulumi.get(self, "status_badges_are_private")

    @status_badges_are_private.setter
    def status_badges_are_private(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status_badges_are_private", value)


class ProjectPipelineSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforce_job_scope: Optional[pulumi.Input[bool]] = None,
                 enforce_referenced_repo_scoped_token: Optional[pulumi.Input[bool]] = None,
                 enforce_settable_var: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 publish_pipeline_metadata: Optional[pulumi.Input[bool]] = None,
                 status_badges_are_private: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages Pipeline Settings for Azure DevOps projects

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_project_pipeline_settings = azuredevops.ProjectPipelineSettings("exampleProjectPipelineSettings",
            project_id=example_project.id,
            enforce_job_scope=True,
            enforce_referenced_repo_scoped_token=False,
            enforce_settable_var=True,
            publish_pipeline_metadata=False,
            status_badges_are_private=True)
        ```
        ## Relevant Links

        No official documentation available

        ## PAT Permissions Required

        - Full Access

        ## Import

        Azure DevOps feature settings can be imported using the project id, e.g.

        ```sh
         $ pulumi import azuredevops:index/projectPipelineSettings:ProjectPipelineSettings example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enforce_job_scope: Limit job authorization scope to current project for non-release pipelines.
        :param pulumi.Input[bool] enforce_referenced_repo_scoped_token: Protect access to repositories in YAML pipelines.
        :param pulumi.Input[bool] enforce_settable_var: Limit variables that can be set at queue time.
        :param pulumi.Input[str] project_id: The `id` of the project for which the project pipeline settings will be managed.
        :param pulumi.Input[bool] publish_pipeline_metadata: Publish metadata from pipelines.
        :param pulumi.Input[bool] status_badges_are_private: Disable anonymous access to badges.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectPipelineSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Pipeline Settings for Azure DevOps projects

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_project_pipeline_settings = azuredevops.ProjectPipelineSettings("exampleProjectPipelineSettings",
            project_id=example_project.id,
            enforce_job_scope=True,
            enforce_referenced_repo_scoped_token=False,
            enforce_settable_var=True,
            publish_pipeline_metadata=False,
            status_badges_are_private=True)
        ```
        ## Relevant Links

        No official documentation available

        ## PAT Permissions Required

        - Full Access

        ## Import

        Azure DevOps feature settings can be imported using the project id, e.g.

        ```sh
         $ pulumi import azuredevops:index/projectPipelineSettings:ProjectPipelineSettings example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ProjectPipelineSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectPipelineSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforce_job_scope: Optional[pulumi.Input[bool]] = None,
                 enforce_referenced_repo_scoped_token: Optional[pulumi.Input[bool]] = None,
                 enforce_settable_var: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 publish_pipeline_metadata: Optional[pulumi.Input[bool]] = None,
                 status_badges_are_private: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectPipelineSettingsArgs.__new__(ProjectPipelineSettingsArgs)

            __props__.__dict__["enforce_job_scope"] = enforce_job_scope
            __props__.__dict__["enforce_referenced_repo_scoped_token"] = enforce_referenced_repo_scoped_token
            __props__.__dict__["enforce_settable_var"] = enforce_settable_var
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["publish_pipeline_metadata"] = publish_pipeline_metadata
            __props__.__dict__["status_badges_are_private"] = status_badges_are_private
        super(ProjectPipelineSettings, __self__).__init__(
            'azuredevops:index/projectPipelineSettings:ProjectPipelineSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enforce_job_scope: Optional[pulumi.Input[bool]] = None,
            enforce_referenced_repo_scoped_token: Optional[pulumi.Input[bool]] = None,
            enforce_settable_var: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            publish_pipeline_metadata: Optional[pulumi.Input[bool]] = None,
            status_badges_are_private: Optional[pulumi.Input[bool]] = None) -> 'ProjectPipelineSettings':
        """
        Get an existing ProjectPipelineSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enforce_job_scope: Limit job authorization scope to current project for non-release pipelines.
        :param pulumi.Input[bool] enforce_referenced_repo_scoped_token: Protect access to repositories in YAML pipelines.
        :param pulumi.Input[bool] enforce_settable_var: Limit variables that can be set at queue time.
        :param pulumi.Input[str] project_id: The `id` of the project for which the project pipeline settings will be managed.
        :param pulumi.Input[bool] publish_pipeline_metadata: Publish metadata from pipelines.
        :param pulumi.Input[bool] status_badges_are_private: Disable anonymous access to badges.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectPipelineSettingsState.__new__(_ProjectPipelineSettingsState)

        __props__.__dict__["enforce_job_scope"] = enforce_job_scope
        __props__.__dict__["enforce_referenced_repo_scoped_token"] = enforce_referenced_repo_scoped_token
        __props__.__dict__["enforce_settable_var"] = enforce_settable_var
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["publish_pipeline_metadata"] = publish_pipeline_metadata
        __props__.__dict__["status_badges_are_private"] = status_badges_are_private
        return ProjectPipelineSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enforceJobScope")
    def enforce_job_scope(self) -> pulumi.Output[bool]:
        """
        Limit job authorization scope to current project for non-release pipelines.
        """
        return pulumi.get(self, "enforce_job_scope")

    @property
    @pulumi.getter(name="enforceReferencedRepoScopedToken")
    def enforce_referenced_repo_scoped_token(self) -> pulumi.Output[bool]:
        """
        Protect access to repositories in YAML pipelines.
        """
        return pulumi.get(self, "enforce_referenced_repo_scoped_token")

    @property
    @pulumi.getter(name="enforceSettableVar")
    def enforce_settable_var(self) -> pulumi.Output[bool]:
        """
        Limit variables that can be set at queue time.
        """
        return pulumi.get(self, "enforce_settable_var")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The `id` of the project for which the project pipeline settings will be managed.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publishPipelineMetadata")
    def publish_pipeline_metadata(self) -> pulumi.Output[bool]:
        """
        Publish metadata from pipelines.
        """
        return pulumi.get(self, "publish_pipeline_metadata")

    @property
    @pulumi.getter(name="statusBadgesArePrivate")
    def status_badges_are_private(self) -> pulumi.Output[bool]:
        """
        Disable anonymous access to badges.
        """
        return pulumi.get(self, "status_badges_are_private")
