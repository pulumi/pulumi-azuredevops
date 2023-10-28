# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'BranchPolicyBuildValidationSettings',
    'BranchPolicyBuildValidationSettingsScope',
    'BranchPolicyMinReviewersSettings',
    'BranchPolicyMinReviewersSettingsScope',
]

@pulumi.output_type
class BranchPolicyBuildValidationSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "buildDefinitionId":
            suggest = "build_definition_id"
        elif key == "displayName":
            suggest = "display_name"
        elif key == "filenamePatterns":
            suggest = "filename_patterns"
        elif key == "manualQueueOnly":
            suggest = "manual_queue_only"
        elif key == "queueOnSourceUpdateOnly":
            suggest = "queue_on_source_update_only"
        elif key == "validDuration":
            suggest = "valid_duration"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BranchPolicyBuildValidationSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BranchPolicyBuildValidationSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BranchPolicyBuildValidationSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 build_definition_id: int,
                 display_name: str,
                 scopes: Sequence['outputs.BranchPolicyBuildValidationSettingsScope'],
                 filename_patterns: Optional[Sequence[str]] = None,
                 manual_queue_only: Optional[bool] = None,
                 queue_on_source_update_only: Optional[bool] = None,
                 valid_duration: Optional[int] = None):
        """
        :param int build_definition_id: The ID of the build to monitor for the policy.
        :param str display_name: The display name for the policy.
        :param Sequence['BranchPolicyBuildValidationSettingsScopeArgs'] scopes: Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        :param Sequence[str] filename_patterns: If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
        :param bool manual_queue_only: If set to true, the build will need to be manually queued. Defaults to `false`
        :param bool queue_on_source_update_only: True if the build should queue on source updates only. Defaults to `true`.
        :param int valid_duration: The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
        """
        pulumi.set(__self__, "build_definition_id", build_definition_id)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "scopes", scopes)
        if filename_patterns is not None:
            pulumi.set(__self__, "filename_patterns", filename_patterns)
        if manual_queue_only is not None:
            pulumi.set(__self__, "manual_queue_only", manual_queue_only)
        if queue_on_source_update_only is not None:
            pulumi.set(__self__, "queue_on_source_update_only", queue_on_source_update_only)
        if valid_duration is not None:
            pulumi.set(__self__, "valid_duration", valid_duration)

    @property
    @pulumi.getter(name="buildDefinitionId")
    def build_definition_id(self) -> int:
        """
        The ID of the build to monitor for the policy.
        """
        return pulumi.get(self, "build_definition_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for the policy.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence['outputs.BranchPolicyBuildValidationSettingsScope']:
        """
        Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="filenamePatterns")
    def filename_patterns(self) -> Optional[Sequence[str]]:
        """
        If a path filter is set, the policy will only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
        """
        return pulumi.get(self, "filename_patterns")

    @property
    @pulumi.getter(name="manualQueueOnly")
    def manual_queue_only(self) -> Optional[bool]:
        """
        If set to true, the build will need to be manually queued. Defaults to `false`
        """
        return pulumi.get(self, "manual_queue_only")

    @property
    @pulumi.getter(name="queueOnSourceUpdateOnly")
    def queue_on_source_update_only(self) -> Optional[bool]:
        """
        True if the build should queue on source updates only. Defaults to `true`.
        """
        return pulumi.get(self, "queue_on_source_update_only")

    @property
    @pulumi.getter(name="validDuration")
    def valid_duration(self) -> Optional[int]:
        """
        The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
        """
        return pulumi.get(self, "valid_duration")


@pulumi.output_type
class BranchPolicyBuildValidationSettingsScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "matchType":
            suggest = "match_type"
        elif key == "repositoryId":
            suggest = "repository_id"
        elif key == "repositoryRef":
            suggest = "repository_ref"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BranchPolicyBuildValidationSettingsScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BranchPolicyBuildValidationSettingsScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BranchPolicyBuildValidationSettingsScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 match_type: Optional[str] = None,
                 repository_id: Optional[str] = None,
                 repository_ref: Optional[str] = None):
        """
        :param str match_type: The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
        :param str repository_id: The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
        :param str repository_ref: The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        if match_type is not None:
            pulumi.set(__self__, "match_type", match_type)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if repository_ref is not None:
            pulumi.set(__self__, "repository_ref", repository_ref)

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> Optional[str]:
        """
        The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
        """
        return pulumi.get(self, "match_type")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[str]:
        """
        The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter(name="repositoryRef")
    def repository_ref(self) -> Optional[str]:
        """
        The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        return pulumi.get(self, "repository_ref")


@pulumi.output_type
class BranchPolicyMinReviewersSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowCompletionWithRejectsOrWaits":
            suggest = "allow_completion_with_rejects_or_waits"
        elif key == "lastPusherCannotApprove":
            suggest = "last_pusher_cannot_approve"
        elif key == "onLastIterationRequireVote":
            suggest = "on_last_iteration_require_vote"
        elif key == "onPushResetAllVotes":
            suggest = "on_push_reset_all_votes"
        elif key == "onPushResetApprovedVotes":
            suggest = "on_push_reset_approved_votes"
        elif key == "reviewerCount":
            suggest = "reviewer_count"
        elif key == "submitterCanVote":
            suggest = "submitter_can_vote"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BranchPolicyMinReviewersSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BranchPolicyMinReviewersSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BranchPolicyMinReviewersSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 scopes: Sequence['outputs.BranchPolicyMinReviewersSettingsScope'],
                 allow_completion_with_rejects_or_waits: Optional[bool] = None,
                 last_pusher_cannot_approve: Optional[bool] = None,
                 on_last_iteration_require_vote: Optional[bool] = None,
                 on_push_reset_all_votes: Optional[bool] = None,
                 on_push_reset_approved_votes: Optional[bool] = None,
                 reviewer_count: Optional[int] = None,
                 submitter_can_vote: Optional[bool] = None):
        """
        :param Sequence['BranchPolicyMinReviewersSettingsScopeArgs'] scopes: A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        :param bool allow_completion_with_rejects_or_waits: Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
        :param bool last_pusher_cannot_approve: Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
        :param bool on_last_iteration_require_vote: On last iteration require vote. Defaults to `false`.
        :param bool on_push_reset_all_votes: When new changes are pushed reset all code reviewer votes. Defaults to `false`.
               
               > **Note:** If `on_push_reset_all_votes` is `true` then `on_push_reset_approved_votes` will be set to `true`. To enable `on_push_reset_approved_votes`, you need explicitly set `on_push_reset_all_votes` `false` or not configure.
        :param bool on_push_reset_approved_votes: When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
        :param int reviewer_count: The number of reviewers needed to approve.
        :param bool submitter_can_vote: Allow requesters to approve their own changes. Defaults to `false`.
        """
        pulumi.set(__self__, "scopes", scopes)
        if allow_completion_with_rejects_or_waits is not None:
            pulumi.set(__self__, "allow_completion_with_rejects_or_waits", allow_completion_with_rejects_or_waits)
        if last_pusher_cannot_approve is not None:
            pulumi.set(__self__, "last_pusher_cannot_approve", last_pusher_cannot_approve)
        if on_last_iteration_require_vote is not None:
            pulumi.set(__self__, "on_last_iteration_require_vote", on_last_iteration_require_vote)
        if on_push_reset_all_votes is not None:
            pulumi.set(__self__, "on_push_reset_all_votes", on_push_reset_all_votes)
        if on_push_reset_approved_votes is not None:
            pulumi.set(__self__, "on_push_reset_approved_votes", on_push_reset_approved_votes)
        if reviewer_count is not None:
            pulumi.set(__self__, "reviewer_count", reviewer_count)
        if submitter_can_vote is not None:
            pulumi.set(__self__, "submitter_can_vote", submitter_can_vote)

    @property
    @pulumi.getter
    def scopes(self) -> Sequence['outputs.BranchPolicyMinReviewersSettingsScope']:
        """
        A `scope` block as defined below. Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="allowCompletionWithRejectsOrWaits")
    def allow_completion_with_rejects_or_waits(self) -> Optional[bool]:
        """
        Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
        """
        return pulumi.get(self, "allow_completion_with_rejects_or_waits")

    @property
    @pulumi.getter(name="lastPusherCannotApprove")
    def last_pusher_cannot_approve(self) -> Optional[bool]:
        """
        Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
        """
        return pulumi.get(self, "last_pusher_cannot_approve")

    @property
    @pulumi.getter(name="onLastIterationRequireVote")
    def on_last_iteration_require_vote(self) -> Optional[bool]:
        """
        On last iteration require vote. Defaults to `false`.
        """
        return pulumi.get(self, "on_last_iteration_require_vote")

    @property
    @pulumi.getter(name="onPushResetAllVotes")
    def on_push_reset_all_votes(self) -> Optional[bool]:
        """
        When new changes are pushed reset all code reviewer votes. Defaults to `false`.

        > **Note:** If `on_push_reset_all_votes` is `true` then `on_push_reset_approved_votes` will be set to `true`. To enable `on_push_reset_approved_votes`, you need explicitly set `on_push_reset_all_votes` `false` or not configure.
        """
        return pulumi.get(self, "on_push_reset_all_votes")

    @property
    @pulumi.getter(name="onPushResetApprovedVotes")
    def on_push_reset_approved_votes(self) -> Optional[bool]:
        """
        When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
        """
        return pulumi.get(self, "on_push_reset_approved_votes")

    @property
    @pulumi.getter(name="reviewerCount")
    def reviewer_count(self) -> Optional[int]:
        """
        The number of reviewers needed to approve.
        """
        return pulumi.get(self, "reviewer_count")

    @property
    @pulumi.getter(name="submitterCanVote")
    def submitter_can_vote(self) -> Optional[bool]:
        """
        Allow requesters to approve their own changes. Defaults to `false`.
        """
        return pulumi.get(self, "submitter_can_vote")


@pulumi.output_type
class BranchPolicyMinReviewersSettingsScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "matchType":
            suggest = "match_type"
        elif key == "repositoryId":
            suggest = "repository_id"
        elif key == "repositoryRef":
            suggest = "repository_ref"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BranchPolicyMinReviewersSettingsScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BranchPolicyMinReviewersSettingsScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BranchPolicyMinReviewersSettingsScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 match_type: Optional[str] = None,
                 repository_id: Optional[str] = None,
                 repository_ref: Optional[str] = None):
        """
        :param str match_type: The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
        :param str repository_id: The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
        :param str repository_ref: The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        if match_type is not None:
            pulumi.set(__self__, "match_type", match_type)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if repository_ref is not None:
            pulumi.set(__self__, "repository_ref", repository_ref)

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> Optional[str]:
        """
        The match type to use when applying the policy. Supported values are `Exact` (default), `Prefix` or `DefaultBranch`.
        """
        return pulumi.get(self, "match_type")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[str]:
        """
        The repository ID. Needed only if the scope of the policy will be limited to a single repository. If `match_type` is `DefaultBranch`, this should not be defined.
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter(name="repositoryRef")
    def repository_ref(self) -> Optional[str]:
        """
        The ref pattern to use for the match when `match_type` other than `DefaultBranch`. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        return pulumi.get(self, "repository_ref")


