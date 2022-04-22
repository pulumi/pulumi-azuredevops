# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'BuildDefinitionCiTriggerArgs',
    'BuildDefinitionCiTriggerOverrideArgs',
    'BuildDefinitionCiTriggerOverrideBranchFilterArgs',
    'BuildDefinitionCiTriggerOverridePathFilterArgs',
    'BuildDefinitionPullRequestTriggerArgs',
    'BuildDefinitionPullRequestTriggerForksArgs',
    'BuildDefinitionPullRequestTriggerOverrideArgs',
    'BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs',
    'BuildDefinitionPullRequestTriggerOverridePathFilterArgs',
    'BuildDefinitionRepositoryArgs',
    'BuildDefinitionScheduleArgs',
    'BuildDefinitionScheduleBranchFilterArgs',
    'BuildDefinitionVariableArgs',
]

@pulumi.input_type
class BuildDefinitionCiTriggerArgs:
    def __init__(__self__, *,
                 override: Optional[pulumi.Input['BuildDefinitionCiTriggerOverrideArgs']] = None,
                 use_yaml: Optional[pulumi.Input[bool]] = None):
        if override is not None:
            pulumi.set(__self__, "override", override)
        if use_yaml is not None:
            pulumi.set(__self__, "use_yaml", use_yaml)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input['BuildDefinitionCiTriggerOverrideArgs']]:
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input['BuildDefinitionCiTriggerOverrideArgs']]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter(name="useYaml")
    def use_yaml(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_yaml")

    @use_yaml.setter
    def use_yaml(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_yaml", value)


@pulumi.input_type
class BuildDefinitionCiTriggerOverrideArgs:
    def __init__(__self__, *,
                 batch: Optional[pulumi.Input[bool]] = None,
                 branch_filters: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionCiTriggerOverrideBranchFilterArgs']]]] = None,
                 max_concurrent_builds_per_branch: Optional[pulumi.Input[int]] = None,
                 path_filters: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionCiTriggerOverridePathFilterArgs']]]] = None,
                 polling_interval: Optional[pulumi.Input[int]] = None,
                 polling_job_id: Optional[pulumi.Input[str]] = None):
        if batch is not None:
            pulumi.set(__self__, "batch", batch)
        if branch_filters is not None:
            pulumi.set(__self__, "branch_filters", branch_filters)
        if max_concurrent_builds_per_branch is not None:
            pulumi.set(__self__, "max_concurrent_builds_per_branch", max_concurrent_builds_per_branch)
        if path_filters is not None:
            pulumi.set(__self__, "path_filters", path_filters)
        if polling_interval is not None:
            pulumi.set(__self__, "polling_interval", polling_interval)
        if polling_job_id is not None:
            pulumi.set(__self__, "polling_job_id", polling_job_id)

    @property
    @pulumi.getter
    def batch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "batch")

    @batch.setter
    def batch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "batch", value)

    @property
    @pulumi.getter(name="branchFilters")
    def branch_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionCiTriggerOverrideBranchFilterArgs']]]]:
        return pulumi.get(self, "branch_filters")

    @branch_filters.setter
    def branch_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionCiTriggerOverrideBranchFilterArgs']]]]):
        pulumi.set(self, "branch_filters", value)

    @property
    @pulumi.getter(name="maxConcurrentBuildsPerBranch")
    def max_concurrent_builds_per_branch(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_concurrent_builds_per_branch")

    @max_concurrent_builds_per_branch.setter
    def max_concurrent_builds_per_branch(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_builds_per_branch", value)

    @property
    @pulumi.getter(name="pathFilters")
    def path_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionCiTriggerOverridePathFilterArgs']]]]:
        return pulumi.get(self, "path_filters")

    @path_filters.setter
    def path_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionCiTriggerOverridePathFilterArgs']]]]):
        pulumi.set(self, "path_filters", value)

    @property
    @pulumi.getter(name="pollingInterval")
    def polling_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "polling_interval")

    @polling_interval.setter
    def polling_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "polling_interval", value)

    @property
    @pulumi.getter(name="pollingJobId")
    def polling_job_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "polling_job_id")

    @polling_job_id.setter
    def polling_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "polling_job_id", value)


@pulumi.input_type
class BuildDefinitionCiTriggerOverrideBranchFilterArgs:
    def __init__(__self__, *,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "excludes")

    @excludes.setter
    def excludes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excludes", value)

    @property
    @pulumi.getter
    def includes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "includes")

    @includes.setter
    def includes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "includes", value)


@pulumi.input_type
class BuildDefinitionCiTriggerOverridePathFilterArgs:
    def __init__(__self__, *,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "excludes")

    @excludes.setter
    def excludes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excludes", value)

    @property
    @pulumi.getter
    def includes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "includes")

    @includes.setter
    def includes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "includes", value)


@pulumi.input_type
class BuildDefinitionPullRequestTriggerArgs:
    def __init__(__self__, *,
                 forks: pulumi.Input['BuildDefinitionPullRequestTriggerForksArgs'],
                 comment_required: Optional[pulumi.Input[str]] = None,
                 initial_branch: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input['BuildDefinitionPullRequestTriggerOverrideArgs']] = None,
                 use_yaml: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "forks", forks)
        if comment_required is not None:
            pulumi.set(__self__, "comment_required", comment_required)
        if initial_branch is not None:
            pulumi.set(__self__, "initial_branch", initial_branch)
        if override is not None:
            pulumi.set(__self__, "override", override)
        if use_yaml is not None:
            pulumi.set(__self__, "use_yaml", use_yaml)

    @property
    @pulumi.getter
    def forks(self) -> pulumi.Input['BuildDefinitionPullRequestTriggerForksArgs']:
        return pulumi.get(self, "forks")

    @forks.setter
    def forks(self, value: pulumi.Input['BuildDefinitionPullRequestTriggerForksArgs']):
        pulumi.set(self, "forks", value)

    @property
    @pulumi.getter(name="commentRequired")
    def comment_required(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment_required")

    @comment_required.setter
    def comment_required(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment_required", value)

    @property
    @pulumi.getter(name="initialBranch")
    def initial_branch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "initial_branch")

    @initial_branch.setter
    def initial_branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initial_branch", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input['BuildDefinitionPullRequestTriggerOverrideArgs']]:
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input['BuildDefinitionPullRequestTriggerOverrideArgs']]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter(name="useYaml")
    def use_yaml(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_yaml")

    @use_yaml.setter
    def use_yaml(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_yaml", value)


@pulumi.input_type
class BuildDefinitionPullRequestTriggerForksArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 share_secrets: pulumi.Input[bool]):
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "share_secrets", share_secrets)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="shareSecrets")
    def share_secrets(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "share_secrets")

    @share_secrets.setter
    def share_secrets(self, value: pulumi.Input[bool]):
        pulumi.set(self, "share_secrets", value)


@pulumi.input_type
class BuildDefinitionPullRequestTriggerOverrideArgs:
    def __init__(__self__, *,
                 auto_cancel: Optional[pulumi.Input[bool]] = None,
                 branch_filters: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs']]]] = None,
                 path_filters: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionPullRequestTriggerOverridePathFilterArgs']]]] = None):
        if auto_cancel is not None:
            pulumi.set(__self__, "auto_cancel", auto_cancel)
        if branch_filters is not None:
            pulumi.set(__self__, "branch_filters", branch_filters)
        if path_filters is not None:
            pulumi.set(__self__, "path_filters", path_filters)

    @property
    @pulumi.getter(name="autoCancel")
    def auto_cancel(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_cancel")

    @auto_cancel.setter
    def auto_cancel(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_cancel", value)

    @property
    @pulumi.getter(name="branchFilters")
    def branch_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs']]]]:
        return pulumi.get(self, "branch_filters")

    @branch_filters.setter
    def branch_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs']]]]):
        pulumi.set(self, "branch_filters", value)

    @property
    @pulumi.getter(name="pathFilters")
    def path_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionPullRequestTriggerOverridePathFilterArgs']]]]:
        return pulumi.get(self, "path_filters")

    @path_filters.setter
    def path_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionPullRequestTriggerOverridePathFilterArgs']]]]):
        pulumi.set(self, "path_filters", value)


@pulumi.input_type
class BuildDefinitionPullRequestTriggerOverrideBranchFilterArgs:
    def __init__(__self__, *,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "excludes")

    @excludes.setter
    def excludes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excludes", value)

    @property
    @pulumi.getter
    def includes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "includes")

    @includes.setter
    def includes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "includes", value)


@pulumi.input_type
class BuildDefinitionPullRequestTriggerOverridePathFilterArgs:
    def __init__(__self__, *,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "excludes")

    @excludes.setter
    def excludes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excludes", value)

    @property
    @pulumi.getter
    def includes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "includes")

    @includes.setter
    def includes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "includes", value)


@pulumi.input_type
class BuildDefinitionRepositoryArgs:
    def __init__(__self__, *,
                 repo_id: pulumi.Input[str],
                 repo_type: pulumi.Input[str],
                 yml_path: pulumi.Input[str],
                 branch_name: Optional[pulumi.Input[str]] = None,
                 github_enterprise_url: Optional[pulumi.Input[str]] = None,
                 report_build_status: Optional[pulumi.Input[bool]] = None,
                 service_connection_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "repo_id", repo_id)
        pulumi.set(__self__, "repo_type", repo_type)
        pulumi.set(__self__, "yml_path", yml_path)
        if branch_name is not None:
            pulumi.set(__self__, "branch_name", branch_name)
        if github_enterprise_url is not None:
            pulumi.set(__self__, "github_enterprise_url", github_enterprise_url)
        if report_build_status is not None:
            pulumi.set(__self__, "report_build_status", report_build_status)
        if service_connection_id is not None:
            pulumi.set(__self__, "service_connection_id", service_connection_id)

    @property
    @pulumi.getter(name="repoId")
    def repo_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repo_id")

    @repo_id.setter
    def repo_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "repo_id", value)

    @property
    @pulumi.getter(name="repoType")
    def repo_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repo_type")

    @repo_type.setter
    def repo_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "repo_type", value)

    @property
    @pulumi.getter(name="ymlPath")
    def yml_path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "yml_path")

    @yml_path.setter
    def yml_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "yml_path", value)

    @property
    @pulumi.getter(name="branchName")
    def branch_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "branch_name")

    @branch_name.setter
    def branch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch_name", value)

    @property
    @pulumi.getter(name="githubEnterpriseUrl")
    def github_enterprise_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "github_enterprise_url")

    @github_enterprise_url.setter
    def github_enterprise_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "github_enterprise_url", value)

    @property
    @pulumi.getter(name="reportBuildStatus")
    def report_build_status(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "report_build_status")

    @report_build_status.setter
    def report_build_status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "report_build_status", value)

    @property
    @pulumi.getter(name="serviceConnectionId")
    def service_connection_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_connection_id")

    @service_connection_id.setter
    def service_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_connection_id", value)


@pulumi.input_type
class BuildDefinitionScheduleArgs:
    def __init__(__self__, *,
                 days_to_builds: pulumi.Input[Sequence[pulumi.Input[str]]],
                 branch_filters: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleBranchFilterArgs']]]] = None,
                 schedule_job_id: Optional[pulumi.Input[str]] = None,
                 schedule_only_with_changes: Optional[pulumi.Input[bool]] = None,
                 start_hours: Optional[pulumi.Input[int]] = None,
                 start_minutes: Optional[pulumi.Input[int]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "days_to_builds", days_to_builds)
        if branch_filters is not None:
            pulumi.set(__self__, "branch_filters", branch_filters)
        if schedule_job_id is not None:
            pulumi.set(__self__, "schedule_job_id", schedule_job_id)
        if schedule_only_with_changes is not None:
            pulumi.set(__self__, "schedule_only_with_changes", schedule_only_with_changes)
        if start_hours is not None:
            pulumi.set(__self__, "start_hours", start_hours)
        if start_minutes is not None:
            pulumi.set(__self__, "start_minutes", start_minutes)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="daysToBuilds")
    def days_to_builds(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "days_to_builds")

    @days_to_builds.setter
    def days_to_builds(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "days_to_builds", value)

    @property
    @pulumi.getter(name="branchFilters")
    def branch_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleBranchFilterArgs']]]]:
        return pulumi.get(self, "branch_filters")

    @branch_filters.setter
    def branch_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildDefinitionScheduleBranchFilterArgs']]]]):
        pulumi.set(self, "branch_filters", value)

    @property
    @pulumi.getter(name="scheduleJobId")
    def schedule_job_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule_job_id")

    @schedule_job_id.setter
    def schedule_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_job_id", value)

    @property
    @pulumi.getter(name="scheduleOnlyWithChanges")
    def schedule_only_with_changes(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "schedule_only_with_changes")

    @schedule_only_with_changes.setter
    def schedule_only_with_changes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "schedule_only_with_changes", value)

    @property
    @pulumi.getter(name="startHours")
    def start_hours(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_hours")

    @start_hours.setter
    def start_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_hours", value)

    @property
    @pulumi.getter(name="startMinutes")
    def start_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_minutes")

    @start_minutes.setter
    def start_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_minutes", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class BuildDefinitionScheduleBranchFilterArgs:
    def __init__(__self__, *,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "excludes")

    @excludes.setter
    def excludes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excludes", value)

    @property
    @pulumi.getter
    def includes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "includes")

    @includes.setter
    def includes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "includes", value)


@pulumi.input_type
class BuildDefinitionVariableArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 allow_override: Optional[pulumi.Input[bool]] = None,
                 is_secret: Optional[pulumi.Input[bool]] = None,
                 secret_value: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        if allow_override is not None:
            pulumi.set(__self__, "allow_override", allow_override)
        if is_secret is not None:
            pulumi.set(__self__, "is_secret", is_secret)
        if secret_value is not None:
            pulumi.set(__self__, "secret_value", secret_value)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="allowOverride")
    def allow_override(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_override")

    @allow_override.setter
    def allow_override(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_override", value)

    @property
    @pulumi.getter(name="isSecret")
    def is_secret(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_secret")

    @is_secret.setter
    def is_secret(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_secret", value)

    @property
    @pulumi.getter(name="secretValue")
    def secret_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_value")

    @secret_value.setter
    def secret_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


