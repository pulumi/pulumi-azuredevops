# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['BranchPolicyCommentResolutionArgs', 'BranchPolicyCommentResolution']

@pulumi.input_type
class BranchPolicyCommentResolutionArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[_builtins.str],
                 settings: pulumi.Input['BranchPolicyCommentResolutionSettingsArgs'],
                 blocking: Optional[pulumi.Input[_builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None):
        """
        The set of arguments for constructing a BranchPolicyCommentResolution resource.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input['BranchPolicyCommentResolutionSettingsArgs'] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        :param pulumi.Input[_builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[_builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "settings", settings)
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter
    def settings(self) -> pulumi.Input['BranchPolicyCommentResolutionSettingsArgs']:
        """
        A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: pulumi.Input['BranchPolicyCommentResolutionSettingsArgs']):
        pulumi.set(self, "settings", value)

    @_builtins.property
    @pulumi.getter
    def blocking(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @blocking.setter
    def blocking(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "blocking", value)

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _BranchPolicyCommentResolutionState:
    def __init__(__self__, *,
                 blocking: Optional[pulumi.Input[_builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 settings: Optional[pulumi.Input['BranchPolicyCommentResolutionSettingsArgs']] = None):
        """
        Input properties used for looking up and filtering BranchPolicyCommentResolution resources.
        :param pulumi.Input[_builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[_builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input['BranchPolicyCommentResolutionSettingsArgs'] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)

    @_builtins.property
    @pulumi.getter
    def blocking(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @blocking.setter
    def blocking(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "blocking", value)

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project_id", value)

    @_builtins.property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input['BranchPolicyCommentResolutionSettingsArgs']]:
        """
        A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input['BranchPolicyCommentResolutionSettingsArgs']]):
        pulumi.set(self, "settings", value)


@pulumi.type_token("azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution")
class BranchPolicyCommentResolution(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[_builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 settings: Optional[pulumi.Input[Union['BranchPolicyCommentResolutionSettingsArgs', 'BranchPolicyCommentResolutionSettingsArgsDict']]] = None,
                 __props__=None):
        """
        Configure a comment resolution policy for your branch within Azure DevOps project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="Example Project")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "init_type": "Clean",
            })
        example_branch_policy_comment_resolution = azuredevops.BranchPolicyCommentResolution("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            settings={
                "scopes": [
                    {
                        "repository_id": example_git.id,
                        "repository_ref": example_git.default_branch,
                        "match_type": "Exact",
                    },
                    {
                        "repository_id": example_git.id,
                        "repository_ref": "refs/heads/releases",
                        "match_type": "Prefix",
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
        $ pulumi import azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[_builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Union['BranchPolicyCommentResolutionSettingsArgs', 'BranchPolicyCommentResolutionSettingsArgsDict']] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchPolicyCommentResolutionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure a comment resolution policy for your branch within Azure DevOps project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="Example Project")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "init_type": "Clean",
            })
        example_branch_policy_comment_resolution = azuredevops.BranchPolicyCommentResolution("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            settings={
                "scopes": [
                    {
                        "repository_id": example_git.id,
                        "repository_ref": example_git.default_branch,
                        "match_type": "Exact",
                    },
                    {
                        "repository_id": example_git.id,
                        "repository_ref": "refs/heads/releases",
                        "match_type": "Prefix",
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
        $ pulumi import azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param BranchPolicyCommentResolutionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchPolicyCommentResolutionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[_builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[_builtins.str]] = None,
                 settings: Optional[pulumi.Input[Union['BranchPolicyCommentResolutionSettingsArgs', 'BranchPolicyCommentResolutionSettingsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchPolicyCommentResolutionArgs.__new__(BranchPolicyCommentResolutionArgs)

            __props__.__dict__["blocking"] = blocking
            __props__.__dict__["enabled"] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if settings is None and not opts.urn:
                raise TypeError("Missing required property 'settings'")
            __props__.__dict__["settings"] = settings
        super(BranchPolicyCommentResolution, __self__).__init__(
            'azuredevops:index/branchPolicyCommentResolution:BranchPolicyCommentResolution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blocking: Optional[pulumi.Input[_builtins.bool]] = None,
            enabled: Optional[pulumi.Input[_builtins.bool]] = None,
            project_id: Optional[pulumi.Input[_builtins.str]] = None,
            settings: Optional[pulumi.Input[Union['BranchPolicyCommentResolutionSettingsArgs', 'BranchPolicyCommentResolutionSettingsArgsDict']]] = None) -> 'BranchPolicyCommentResolution':
        """
        Get an existing BranchPolicyCommentResolution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[_builtins.bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[_builtins.str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Union['BranchPolicyCommentResolutionSettingsArgs', 'BranchPolicyCommentResolutionSettingsArgsDict']] settings: A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchPolicyCommentResolutionState.__new__(_BranchPolicyCommentResolutionState)

        __props__.__dict__["blocking"] = blocking
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["settings"] = settings
        return BranchPolicyCommentResolution(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def blocking(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        A flag indicating if the policy should be blocking. Defaults to `true`.
        """
        return pulumi.get(self, "blocking")

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        A flag indicating if the policy should be enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the project in which the policy will be created.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter
    def settings(self) -> pulumi.Output['outputs.BranchPolicyCommentResolutionSettings']:
        """
        A `settings` block as defined below. Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

