# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BranchPolicyMinReviewers']


class BranchPolicyMinReviewers(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[pulumi.InputType['BranchPolicyMinReviewersSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a minimum reviewer branch policy within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project")
        git = azuredevops.Git("git",
            project_id=project.id,
            initialization=azuredevops.GitInitializationArgs(
                init_type="Clean",
            ))
        branch_policy_min_reviewers = azuredevops.BranchPolicyMinReviewers("branchPolicyMinReviewers",
            project_id=project.id,
            enabled=True,
            blocking=True,
            settings=azuredevops.BranchPolicyMinReviewersSettingsArgs(
                reviewer_count=2,
                submitter_can_vote=False,
                scopes=[
                    azuredevops.BranchPolicyMinReviewersSettingsScopeArgs(
                        repository_id=git.id,
                        repository_ref=git.default_branch,
                        match_type="Exact",
                    ),
                    azuredevops.BranchPolicyMinReviewersSettingsScopeArgs(
                        repository_id=git.id,
                        repository_ref="refs/heads/releases",
                        match_type="Prefix",
                    ),
                ],
            ))
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations/create?view=azure-devops-rest-5.1)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID

        ```sh
         $ pulumi import azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers p aa4a9756-8a86-4588-86d7-b3ee2d88b033/60
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[pulumi.InputType['BranchPolicyMinReviewersSettingsArgs']] settings: Configuration for the policy. This block must be defined exactly once.
        """
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

            __props__['blocking'] = blocking
            __props__['enabled'] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            if settings is None and not opts.urn:
                raise TypeError("Missing required property 'settings'")
            __props__['settings'] = settings
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Policy/branchPolicyMinReviewers:BranchPolicyMinReviewers")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BranchPolicyMinReviewers, __self__).__init__(
            'azuredevops:index/branchPolicyMinReviewers:BranchPolicyMinReviewers',
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
            settings: Optional[pulumi.Input[pulumi.InputType['BranchPolicyMinReviewersSettingsArgs']]] = None) -> 'BranchPolicyMinReviewers':
        """
        Get an existing BranchPolicyMinReviewers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[pulumi.InputType['BranchPolicyMinReviewersSettingsArgs']] settings: Configuration for the policy. This block must be defined exactly once.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["blocking"] = blocking
        __props__["enabled"] = enabled
        __props__["project_id"] = project_id
        __props__["settings"] = settings
        return BranchPolicyMinReviewers(resource_name, opts=opts, __props__=__props__)

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
    def settings(self) -> pulumi.Output['outputs.BranchPolicyMinReviewersSettings']:
        """
        Configuration for the policy. This block must be defined exactly once.
        """
        return pulumi.get(self, "settings")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

