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
from . import outputs
from ._inputs import *

__all__ = ['BranchPolicyStatusCheckArgs', 'BranchPolicyStatusCheck']

@pulumi.input_type
class BranchPolicyStatusCheckArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[builtins.str],
                 settings: pulumi.Input['BranchPolicyStatusCheckSettingsArgs'],
                 blocking: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a BranchPolicyStatusCheck resource.
        :param pulumi.Input[builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input['BranchPolicyStatusCheckSettingsArgs'] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        :param pulumi.Input[builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "settings", settings)
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Input['BranchPolicyStatusCheckSettingsArgs']:
        """
        A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: pulumi.Input['BranchPolicyStatusCheckSettingsArgs']):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter
    def blocking(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @blocking.setter
    def blocking(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "blocking", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _BranchPolicyStatusCheckState:
    def __init__(__self__, *,
                 blocking: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 settings: Optional[pulumi.Input['BranchPolicyStatusCheckSettingsArgs']] = None):
        """
        Input properties used for looking up and filtering BranchPolicyStatusCheck resources.
        :param pulumi.Input[builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input['BranchPolicyStatusCheckSettingsArgs'] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
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
    def blocking(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @blocking.setter
    def blocking(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "blocking", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input['BranchPolicyStatusCheckSettingsArgs']]:
        """
        A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input['BranchPolicyStatusCheckSettingsArgs']]):
        pulumi.set(self, "settings", value)


class BranchPolicyStatusCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 settings: Optional[pulumi.Input[Union['BranchPolicyStatusCheckSettingsArgs', 'BranchPolicyStatusCheckSettingsArgsDict']]] = None,
                 __props__=None):
        """
        Manages a status check branch policy within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            features={
                "testplans": "disabled",
                "artifacts": "disabled",
            },
            description="Managed by Pulumi")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "init_type": "Clean",
            })
        example_user = azuredevops.User("example",
            principal_name="mail@email.com",
            account_license_type="basic")
        example_branch_policy_status_check = azuredevops.BranchPolicyStatusCheck("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            settings={
                "name": "Release",
                "author_id": example_user.id,
                "invalidate_on_update": True,
                "applicability": "conditional",
                "display_name": "PreCheck",
                "scopes": [
                    {
                        "repository_id": example_git.id,
                        "repository_ref": example_git.default_branch,
                        "match_type": "Exact",
                    },
                    {
                        "match_type": "DefaultBranch",
                    },
                ],
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:

        ```sh
        $ pulumi import azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Union['BranchPolicyStatusCheckSettingsArgs', 'BranchPolicyStatusCheckSettingsArgsDict']] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchPolicyStatusCheckArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a status check branch policy within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            features={
                "testplans": "disabled",
                "artifacts": "disabled",
            },
            description="Managed by Pulumi")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "init_type": "Clean",
            })
        example_user = azuredevops.User("example",
            principal_name="mail@email.com",
            account_license_type="basic")
        example_branch_policy_status_check = azuredevops.BranchPolicyStatusCheck("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            settings={
                "name": "Release",
                "author_id": example_user.id,
                "invalidate_on_update": True,
                "applicability": "conditional",
                "display_name": "PreCheck",
                "scopes": [
                    {
                        "repository_id": example_git.id,
                        "repository_ref": example_git.default_branch,
                        "match_type": "Exact",
                    },
                    {
                        "match_type": "DefaultBranch",
                    },
                ],
            })
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:

        ```sh
        $ pulumi import azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param BranchPolicyStatusCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchPolicyStatusCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 settings: Optional[pulumi.Input[Union['BranchPolicyStatusCheckSettingsArgs', 'BranchPolicyStatusCheckSettingsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchPolicyStatusCheckArgs.__new__(BranchPolicyStatusCheckArgs)

            __props__.__dict__["blocking"] = blocking
            __props__.__dict__["enabled"] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if settings is None and not opts.urn:
                raise TypeError("Missing required property 'settings'")
            __props__.__dict__["settings"] = settings
        super(BranchPolicyStatusCheck, __self__).__init__(
            'azuredevops:index/branchPolicyStatusCheck:BranchPolicyStatusCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blocking: Optional[pulumi.Input[builtins.bool]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            settings: Optional[pulumi.Input[Union['BranchPolicyStatusCheckSettingsArgs', 'BranchPolicyStatusCheckSettingsArgsDict']]] = None) -> 'BranchPolicyStatusCheck':
        """
        Get an existing BranchPolicyStatusCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Union['BranchPolicyStatusCheckSettingsArgs', 'BranchPolicyStatusCheckSettingsArgsDict']] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchPolicyStatusCheckState.__new__(_BranchPolicyStatusCheckState)

        __props__.__dict__["blocking"] = blocking
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["settings"] = settings
        return BranchPolicyStatusCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def blocking(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output['outputs.BranchPolicyStatusCheckSettings']:
        """
        A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

