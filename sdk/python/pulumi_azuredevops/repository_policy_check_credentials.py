# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryPolicyCheckCredentialsArgs', 'RepositoryPolicyCheckCredentials']

@pulumi.input_type
class RepositoryPolicyCheckCredentialsArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RepositoryPolicyCheckCredentials resource.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        RepositoryPolicyCheckCredentialsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            blocking=blocking,
            enabled=enabled,
            repository_ids=repository_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: Optional[pulumi.Input[str]] = None,
             blocking: Optional[pulumi.Input[bool]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")
        if repository_ids is None and 'repositoryIds' in kwargs:
            repository_ids = kwargs['repositoryIds']

        _setter("project_id", project_id)
        if blocking is not None:
            _setter("blocking", blocking)
        if enabled is not None:
            _setter("enabled", enabled)
        if repository_ids is not None:
            _setter("repository_ids", repository_ids)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def blocking(self) -> Optional[pulumi.Input[bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @blocking.setter
    def blocking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blocking", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="repositoryIds")
    def repository_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        return pulumi.get(self, "repository_ids")

    @repository_ids.setter
    def repository_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repository_ids", value)


@pulumi.input_type
class _RepositoryPolicyCheckCredentialsState:
    def __init__(__self__, *,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RepositoryPolicyCheckCredentials resources.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        _RepositoryPolicyCheckCredentialsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            blocking=blocking,
            enabled=enabled,
            project_id=project_id,
            repository_ids=repository_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             blocking: Optional[pulumi.Input[bool]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if repository_ids is None and 'repositoryIds' in kwargs:
            repository_ids = kwargs['repositoryIds']

        if blocking is not None:
            _setter("blocking", blocking)
        if enabled is not None:
            _setter("enabled", enabled)
        if project_id is not None:
            _setter("project_id", project_id)
        if repository_ids is not None:
            _setter("repository_ids", repository_ids)

    @property
    @pulumi.getter
    def blocking(self) -> Optional[pulumi.Input[bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @blocking.setter
    def blocking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blocking", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="repositoryIds")
    def repository_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        return pulumi.get(self, "repository_ids")

    @repository_ids.setter
    def repository_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repository_ids", value)


class RepositoryPolicyCheckCredentials(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manage a credentials check repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.

        > If both project and project policy are enabled, the project policy has high priority.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_git = azuredevops.Git("exampleGit",
            project_id=example_project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        example_repository_policy_check_credentials = azuredevops.RepositoryPolicyCheckCredentials("exampleRepositoryPolicyCheckCredentials",
            project_id=example_project.id,
            enabled=True,
            blocking=True,
            repository_ids=[example_git.id])
        ```

        # Set project level repository policy
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_repository_policy_check_credentials = azuredevops.RepositoryPolicyCheckCredentials("exampleRepositoryPolicyCheckCredentials",
            project_id=example_project.id,
            enabled=True,
            blocking=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID

        ```sh
         $ pulumi import azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryPolicyCheckCredentialsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a credentials check repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.

        > If both project and project policy are enabled, the project policy has high priority.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_git = azuredevops.Git("exampleGit",
            project_id=example_project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        example_repository_policy_check_credentials = azuredevops.RepositoryPolicyCheckCredentials("exampleRepositoryPolicyCheckCredentials",
            project_id=example_project.id,
            enabled=True,
            blocking=True,
            repository_ids=[example_git.id])
        ```

        # Set project level repository policy
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_repository_policy_check_credentials = azuredevops.RepositoryPolicyCheckCredentials("exampleRepositoryPolicyCheckCredentials",
            project_id=example_project.id,
            enabled=True,
            blocking=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID

        ```sh
         $ pulumi import azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryPolicyCheckCredentialsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryPolicyCheckCredentialsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RepositoryPolicyCheckCredentialsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryPolicyCheckCredentialsArgs.__new__(RepositoryPolicyCheckCredentialsArgs)

            __props__.__dict__["blocking"] = blocking
            __props__.__dict__["enabled"] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["repository_ids"] = repository_ids
        super(RepositoryPolicyCheckCredentials, __self__).__init__(
            'azuredevops:index/repositoryPolicyCheckCredentials:RepositoryPolicyCheckCredentials',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blocking: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'RepositoryPolicyCheckCredentials':
        """
        Get an existing RepositoryPolicyCheckCredentials resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryPolicyCheckCredentialsState.__new__(_RepositoryPolicyCheckCredentialsState)

        __props__.__dict__["blocking"] = blocking
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repository_ids"] = repository_ids
        return RepositoryPolicyCheckCredentials(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def blocking(self) -> pulumi.Output[Optional[bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="repositoryIds")
    def repository_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        return pulumi.get(self, "repository_ids")

