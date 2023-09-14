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

__all__ = ['BranchPolicyWorkItemLinkingArgs', 'BranchPolicyWorkItemLinking']

@pulumi.input_type
class BranchPolicyWorkItemLinkingArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 settings: pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs'],
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BranchPolicyWorkItemLinking resource.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs'] settings: Configuration for the policy. This block must be defined exactly once.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "settings", settings)
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

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
    def settings(self) -> pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs']:
        """
        Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs']):
        pulumi.set(self, "settings", value)

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


@pulumi.input_type
class _BranchPolicyWorkItemLinkingState:
    def __init__(__self__, *,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs']] = None):
        """
        Input properties used for looking up and filtering BranchPolicyWorkItemLinking resources.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs'] settings: Configuration for the policy. This block must be defined exactly once.
        """
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)

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
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs']]:
        """
        Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input['BranchPolicyWorkItemLinkingSettingsArgs']]):
        pulumi.set(self, "settings", value)


class BranchPolicyWorkItemLinking(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[pulumi.InputType['BranchPolicyWorkItemLinkingSettingsArgs']]] = None,
                 __props__=None):
        """
        Require associations between branches and a work item within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_git = azuredevops.Git("exampleGit",
            project_id=example_project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        example_branch_policy_work_item_linking = azuredevops.BranchPolicyWorkItemLinking("exampleBranchPolicyWorkItemLinking",
            project_id=example_project.id,
            enabled=True,
            blocking=True,
            settings=azuredevops.BranchPolicyWorkItemLinkingSettingsArgs(
                scopes=[
                    azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs(
                        repository_id=example_git.id,
                        repository_ref=example_git.default_branch,
                        match_type="Exact",
                    ),
                    azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs(
                        repository_id=example_git.id,
                        repository_ref="refs/heads/releases",
                        match_type="Prefix",
                    ),
                    azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs(
                        match_type="DefaultBranch",
                    ),
                ],
            ))
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID

        ```sh
         $ pulumi import azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[pulumi.InputType['BranchPolicyWorkItemLinkingSettingsArgs']] settings: Configuration for the policy. This block must be defined exactly once.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchPolicyWorkItemLinkingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Require associations between branches and a work item within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_git = azuredevops.Git("exampleGit",
            project_id=example_project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        example_branch_policy_work_item_linking = azuredevops.BranchPolicyWorkItemLinking("exampleBranchPolicyWorkItemLinking",
            project_id=example_project.id,
            enabled=True,
            blocking=True,
            settings=azuredevops.BranchPolicyWorkItemLinkingSettingsArgs(
                scopes=[
                    azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs(
                        repository_id=example_git.id,
                        repository_ref=example_git.default_branch,
                        match_type="Exact",
                    ),
                    azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs(
                        repository_id=example_git.id,
                        repository_ref="refs/heads/releases",
                        match_type="Prefix",
                    ),
                    azuredevops.BranchPolicyWorkItemLinkingSettingsScopeArgs(
                        match_type="DefaultBranch",
                    ),
                ],
            ))
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID

        ```sh
         $ pulumi import azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param BranchPolicyWorkItemLinkingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchPolicyWorkItemLinkingArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 settings: Optional[pulumi.Input[pulumi.InputType['BranchPolicyWorkItemLinkingSettingsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchPolicyWorkItemLinkingArgs.__new__(BranchPolicyWorkItemLinkingArgs)

            __props__.__dict__["blocking"] = blocking
            __props__.__dict__["enabled"] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if settings is None and not opts.urn:
                raise TypeError("Missing required property 'settings'")
            __props__.__dict__["settings"] = settings
        super(BranchPolicyWorkItemLinking, __self__).__init__(
            'azuredevops:index/branchPolicyWorkItemLinking:BranchPolicyWorkItemLinking',
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
            settings: Optional[pulumi.Input[pulumi.InputType['BranchPolicyWorkItemLinkingSettingsArgs']]] = None) -> 'BranchPolicyWorkItemLinking':
        """
        Get an existing BranchPolicyWorkItemLinking resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[pulumi.InputType['BranchPolicyWorkItemLinkingSettingsArgs']] settings: Configuration for the policy. This block must be defined exactly once.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchPolicyWorkItemLinkingState.__new__(_BranchPolicyWorkItemLinkingState)

        __props__.__dict__["blocking"] = blocking
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["settings"] = settings
        return BranchPolicyWorkItemLinking(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter
    def settings(self) -> pulumi.Output['outputs.BranchPolicyWorkItemLinkingSettings']:
        """
        Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

