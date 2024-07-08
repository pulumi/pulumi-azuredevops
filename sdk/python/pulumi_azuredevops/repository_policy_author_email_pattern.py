# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['RepositoryPolicyAuthorEmailPatternArgs', 'RepositoryPolicyAuthorEmailPattern']

@pulumi.input_type
class RepositoryPolicyAuthorEmailPatternArgs:
    def __init__(__self__, *,
                 author_email_patterns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 project_id: pulumi.Input[str],
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RepositoryPolicyAuthorEmailPattern resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] author_email_patterns: Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
               Email patterns prefixed with "!" are excluded. Order is important.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        pulumi.set(__self__, "author_email_patterns", author_email_patterns)
        pulumi.set(__self__, "project_id", project_id)
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if repository_ids is not None:
            pulumi.set(__self__, "repository_ids", repository_ids)

    @property
    @pulumi.getter(name="authorEmailPatterns")
    def author_email_patterns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
        Email patterns prefixed with "!" are excluded. Order is important.
        """
        return pulumi.get(self, "author_email_patterns")

    @author_email_patterns.setter
    def author_email_patterns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "author_email_patterns", value)

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
class _RepositoryPolicyAuthorEmailPatternState:
    def __init__(__self__, *,
                 author_email_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RepositoryPolicyAuthorEmailPattern resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] author_email_patterns: Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
               Email patterns prefixed with "!" are excluded. Order is important.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        if author_email_patterns is not None:
            pulumi.set(__self__, "author_email_patterns", author_email_patterns)
        if blocking is not None:
            pulumi.set(__self__, "blocking", blocking)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repository_ids is not None:
            pulumi.set(__self__, "repository_ids", repository_ids)

    @property
    @pulumi.getter(name="authorEmailPatterns")
    def author_email_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
        Email patterns prefixed with "!" are excluded. Order is important.
        """
        return pulumi.get(self, "author_email_patterns")

    @author_email_patterns.setter
    def author_email_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "author_email_patterns", value)

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


class RepositoryPolicyAuthorEmailPattern(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author_email_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 blocking: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manage author email pattern repository policy within Azure DevOps project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "initType": "Clean",
            })
        example_repository_policy_author_email_pattern = azuredevops.RepositoryPolicyAuthorEmailPattern("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            author_email_patterns=[
                "user1@test.com",
                "user2@test.com",
            ],
            repository_ids=[example_git.id])
        ```

        ## Set project level repository policy

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_repository_policy_author_email_pattern = azuredevops.RepositoryPolicyAuthorEmailPattern("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            author_email_patterns=[
                "user1@test.com",
                "user2@test.com",
            ])
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:

        ```sh
        $ pulumi import azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] author_email_patterns: Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
               Email patterns prefixed with "!" are excluded. Order is important.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryPolicyAuthorEmailPatternArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage author email pattern repository policy within Azure DevOps project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_git = azuredevops.Git("example",
            project_id=example.id,
            name="Example Repository",
            initialization={
                "initType": "Clean",
            })
        example_repository_policy_author_email_pattern = azuredevops.RepositoryPolicyAuthorEmailPattern("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            author_email_patterns=[
                "user1@test.com",
                "user2@test.com",
            ],
            repository_ids=[example_git.id])
        ```

        ## Set project level repository policy

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_repository_policy_author_email_pattern = azuredevops.RepositoryPolicyAuthorEmailPattern("example",
            project_id=example.id,
            enabled=True,
            blocking=True,
            author_email_patterns=[
                "user1@test.com",
                "user2@test.com",
            ])
        ```

        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Policy Configurations](https://docs.microsoft.com/en-us/rest/api/azure/devops/policy/configurations?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Branch Policies can be imported using the project ID and policy configuration ID:

        ```sh
        $ pulumi import azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryPolicyAuthorEmailPatternArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryPolicyAuthorEmailPatternArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author_email_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = RepositoryPolicyAuthorEmailPatternArgs.__new__(RepositoryPolicyAuthorEmailPatternArgs)

            if author_email_patterns is None and not opts.urn:
                raise TypeError("Missing required property 'author_email_patterns'")
            __props__.__dict__["author_email_patterns"] = author_email_patterns
            __props__.__dict__["blocking"] = blocking
            __props__.__dict__["enabled"] = enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["repository_ids"] = repository_ids
        super(RepositoryPolicyAuthorEmailPattern, __self__).__init__(
            'azuredevops:index/repositoryPolicyAuthorEmailPattern:RepositoryPolicyAuthorEmailPattern',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            author_email_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            blocking: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'RepositoryPolicyAuthorEmailPattern':
        """
        Get an existing RepositoryPolicyAuthorEmailPattern resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] author_email_patterns: Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
               Email patterns prefixed with "!" are excluded. Order is important.
        :param pulumi.Input[bool] blocking: A flag indicating if the policy should be blocking. Defaults to `true`.
        :param pulumi.Input[bool] enabled: A flag indicating if the policy should be enabled. Defaults to `true`.
        :param pulumi.Input[str] project_id: The ID of the project in which the policy will be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repository_ids: Control whether the policy is enabled for the repository or the project. If `repository_ids` not configured, the policy will be set to the project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryPolicyAuthorEmailPatternState.__new__(_RepositoryPolicyAuthorEmailPatternState)

        __props__.__dict__["author_email_patterns"] = author_email_patterns
        __props__.__dict__["blocking"] = blocking
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repository_ids"] = repository_ids
        return RepositoryPolicyAuthorEmailPattern(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorEmailPatterns")
    def author_email_patterns(self) -> pulumi.Output[Sequence[str]]:
        """
        Block pushes with a commit author email that does not match the patterns. You can specify exact emails or use wildcards. 
        Email patterns prefixed with "!" are excluded. Order is important.
        """
        return pulumi.get(self, "author_email_patterns")

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

