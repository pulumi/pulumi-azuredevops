# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryPolicyReservedNamesArgs', 'RepositoryPolicyReservedNames']

@pulumi.input_type
class RepositoryPolicyReservedNamesArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RepositoryPolicyReservedNames resource.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        pulumi.set(__self__, "project_id", project_id)
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if repository_ids is not None:
            pulumi.set(__self__, "repository_ids", repository_ids)

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
class _RepositoryPolicyReservedNamesState:
    def __init__(__self__, *,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RepositoryPolicyReservedNames resources.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repository_ids is not None:
            pulumi.set(__self__, "repository_ids", repository_ids)

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


class RepositoryPolicyReservedNames(pulumi.CustomResource):
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
        Manage a reserved names repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.

        > If both project and project policy are enabled, the project policy has high priority.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            description="Managed by Terraform",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        git = azuredevops.Git("git",
            project_id=project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        repository_policy_reserved_names = azuredevops.RepositoryPolicyReservedNames("repositoryPolicyReservedNames",
            project_id=project.id,
            enabled=True,
            blocking=True,
            repository_ids=[git.id])
        ```

        # Set project level repository policy
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        repository_policy_reserved_names = azuredevops.RepositoryPolicyReservedNames("repositoryPolicyReservedNames",
            project_id=azuredevops_project["p"]["id"],
            enabled=True,
            blocking=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID

        ```sh
         $ pulumi import azuredevops:index/repositoryPolicyReservedNames:RepositoryPolicyReservedNames p 00000000-0000-0000-0000-000000000000/0
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
                 args: RepositoryPolicyReservedNamesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a reserved names repository policy within Azure DevOps project. Block pushes that introduce files, folders, or branch names that include platform reserved names or incompatible characters.

        > If both project and project policy are enabled, the project policy has high priority.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            description="Managed by Terraform",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        git = azuredevops.Git("git",
            project_id=project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        repository_policy_reserved_names = azuredevops.RepositoryPolicyReservedNames("repositoryPolicyReservedNames",
            project_id=project.id,
            enabled=True,
            blocking=True,
            repository_ids=[git.id])
        ```

        # Set project level repository policy
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        repository_policy_reserved_names = azuredevops.RepositoryPolicyReservedNames("repositoryPolicyReservedNames",
            project_id=azuredevops_project["p"]["id"],
            enabled=True,
            blocking=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps repository policies can be imported using the projectID/policyID or projectName/policyID

        ```sh
         $ pulumi import azuredevops:index/repositoryPolicyReservedNames:RepositoryPolicyReservedNames p 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryPolicyReservedNamesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryPolicyReservedNamesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryPolicyReservedNamesArgs.__new__(RepositoryPolicyReservedNamesArgs)

            __props__.__dict__["blocking"] = blocking
            __props__.__dict__["enabled"] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["repository_ids"] = repository_ids
        super(RepositoryPolicyReservedNames, __self__).__init__(
            'azuredevops:index/repositoryPolicyReservedNames:RepositoryPolicyReservedNames',
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
            repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'RepositoryPolicyReservedNames':
        """
        Get an existing RepositoryPolicyReservedNames resource's state with the given name, id, and optional extra
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

        __props__ = _RepositoryPolicyReservedNamesState.__new__(_RepositoryPolicyReservedNamesState)

        __props__.__dict__["blocking"] = blocking
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repository_ids"] = repository_ids
        return RepositoryPolicyReservedNames(resource_name, opts=opts, __props__=__props__)

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

