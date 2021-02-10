# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'BranchPolicyBuildValidationSettingsArgs',
    'BranchPolicyBuildValidationSettingsScopeArgs',
    'BranchPolicyMinReviewersSettingsArgs',
    'BranchPolicyMinReviewersSettingsScopeArgs',
]

@pulumi.input_type
class BranchPolicyBuildValidationSettingsArgs:
    def __init__(__self__, *,
                 build_definition_id: pulumi.Input[int],
                 display_name: pulumi.Input[str],
                 scopes: pulumi.Input[Sequence[pulumi.Input['BranchPolicyBuildValidationSettingsScopeArgs']]],
                 filename_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 manual_queue_only: Optional[pulumi.Input[bool]] = None,
                 queue_on_source_update_only: Optional[pulumi.Input[bool]] = None,
                 valid_duration: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] build_definition_id: The ID of the build to monitor for the policy.
        :param pulumi.Input[str] display_name: The display name for the policy.
        :param pulumi.Input[Sequence[pulumi.Input['BranchPolicyBuildValidationSettingsScopeArgs']]] scopes: Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] filename_patterns: If a path filter is set, the policy wil only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
        :param pulumi.Input[bool] manual_queue_only: If set to true, the build will need to be manually queued. Defaults to `false`
        :param pulumi.Input[bool] queue_on_source_update_only: True if the build should queue on source updates only. Defaults to `true`.
        :param pulumi.Input[int] valid_duration: The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
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
    def build_definition_id(self) -> pulumi.Input[int]:
        """
        The ID of the build to monitor for the policy.
        """
        return pulumi.get(self, "build_definition_id")

    @build_definition_id.setter
    def build_definition_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "build_definition_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name for the policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Input[Sequence[pulumi.Input['BranchPolicyBuildValidationSettingsScopeArgs']]]:
        """
        Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: pulumi.Input[Sequence[pulumi.Input['BranchPolicyBuildValidationSettingsScopeArgs']]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="filenamePatterns")
    def filename_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        If a path filter is set, the policy wil only apply when files which match the filter are changes. Not setting this field means that the policy will always apply. You can specify absolute paths and wildcards. Example: `["/WebApp/Models/Data.cs", "/WebApp/*", "*.cs"]`. Paths prefixed with "!" are excluded. Example: `["/WebApp/*", "!/WebApp/Tests/*"]`. Order is significant.
        """
        return pulumi.get(self, "filename_patterns")

    @filename_patterns.setter
    def filename_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "filename_patterns", value)

    @property
    @pulumi.getter(name="manualQueueOnly")
    def manual_queue_only(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the build will need to be manually queued. Defaults to `false`
        """
        return pulumi.get(self, "manual_queue_only")

    @manual_queue_only.setter
    def manual_queue_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manual_queue_only", value)

    @property
    @pulumi.getter(name="queueOnSourceUpdateOnly")
    def queue_on_source_update_only(self) -> Optional[pulumi.Input[bool]]:
        """
        True if the build should queue on source updates only. Defaults to `true`.
        """
        return pulumi.get(self, "queue_on_source_update_only")

    @queue_on_source_update_only.setter
    def queue_on_source_update_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "queue_on_source_update_only", value)

    @property
    @pulumi.getter(name="validDuration")
    def valid_duration(self) -> Optional[pulumi.Input[int]]:
        """
        The number of minutes for which the build is valid. If `0`, the build will not expire. Defaults to `720` (12 hours).
        """
        return pulumi.get(self, "valid_duration")

    @valid_duration.setter
    def valid_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "valid_duration", value)


@pulumi.input_type
class BranchPolicyBuildValidationSettingsScopeArgs:
    def __init__(__self__, *,
                 match_type: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 repository_ref: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] match_type: The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
        :param pulumi.Input[str] repository_id: The repository ID. Needed only if the scope of the policy will be limited to a single repository.
        :param pulumi.Input[str] repository_ref: The ref pattern to use for the match. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        if match_type is not None:
            pulumi.set(__self__, "match_type", match_type)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if repository_ref is not None:
            pulumi.set(__self__, "repository_ref", repository_ref)

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> Optional[pulumi.Input[str]]:
        """
        The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
        """
        return pulumi.get(self, "match_type")

    @match_type.setter
    def match_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_type", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The repository ID. Needed only if the scope of the policy will be limited to a single repository.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="repositoryRef")
    def repository_ref(self) -> Optional[pulumi.Input[str]]:
        """
        The ref pattern to use for the match. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        return pulumi.get(self, "repository_ref")

    @repository_ref.setter
    def repository_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_ref", value)


@pulumi.input_type
class BranchPolicyMinReviewersSettingsArgs:
    def __init__(__self__, *,
                 scopes: pulumi.Input[Sequence[pulumi.Input['BranchPolicyMinReviewersSettingsScopeArgs']]],
                 allow_completion_with_rejects_or_waits: Optional[pulumi.Input[bool]] = None,
                 last_pusher_cannot_approve: Optional[pulumi.Input[bool]] = None,
                 on_last_iteration_require_vote: Optional[pulumi.Input[bool]] = None,
                 on_push_reset_all_votes: Optional[pulumi.Input[bool]] = None,
                 on_push_reset_approved_votes: Optional[pulumi.Input[bool]] = None,
                 reviewer_count: Optional[pulumi.Input[int]] = None,
                 submitter_can_vote: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['BranchPolicyMinReviewersSettingsScopeArgs']]] scopes: Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        :param pulumi.Input[bool] allow_completion_with_rejects_or_waits: Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
        :param pulumi.Input[bool] last_pusher_cannot_approve: Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
        :param pulumi.Input[bool] on_last_iteration_require_vote: On last iteration require vote. Defaults to `false`.
        :param pulumi.Input[bool] on_push_reset_all_votes: When new changes are pushed reset all code reviewer votes. Defaults to `false`.
        :param pulumi.Input[bool] on_push_reset_approved_votes: When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
        :param pulumi.Input[int] reviewer_count: The number of reviewers needed to approve.
        :param pulumi.Input[bool] submitter_can_vote: Allow requesters to approve their own changes. Defaults to `false`.
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
    def scopes(self) -> pulumi.Input[Sequence[pulumi.Input['BranchPolicyMinReviewersSettingsScopeArgs']]]:
        """
        Controls which repositories and branches the policy will be enabled for. This block must be defined at least once.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: pulumi.Input[Sequence[pulumi.Input['BranchPolicyMinReviewersSettingsScopeArgs']]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="allowCompletionWithRejectsOrWaits")
    def allow_completion_with_rejects_or_waits(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow completion even if some reviewers vote to wait or reject. Defaults to `false`.
        """
        return pulumi.get(self, "allow_completion_with_rejects_or_waits")

    @allow_completion_with_rejects_or_waits.setter
    def allow_completion_with_rejects_or_waits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_completion_with_rejects_or_waits", value)

    @property
    @pulumi.getter(name="lastPusherCannotApprove")
    def last_pusher_cannot_approve(self) -> Optional[pulumi.Input[bool]]:
        """
        Prohibit the most recent pusher from approving their own changes. Defaults to `false`.
        """
        return pulumi.get(self, "last_pusher_cannot_approve")

    @last_pusher_cannot_approve.setter
    def last_pusher_cannot_approve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "last_pusher_cannot_approve", value)

    @property
    @pulumi.getter(name="onLastIterationRequireVote")
    def on_last_iteration_require_vote(self) -> Optional[pulumi.Input[bool]]:
        """
        On last iteration require vote. Defaults to `false`.
        """
        return pulumi.get(self, "on_last_iteration_require_vote")

    @on_last_iteration_require_vote.setter
    def on_last_iteration_require_vote(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_last_iteration_require_vote", value)

    @property
    @pulumi.getter(name="onPushResetAllVotes")
    def on_push_reset_all_votes(self) -> Optional[pulumi.Input[bool]]:
        """
        When new changes are pushed reset all code reviewer votes. Defaults to `false`.
        """
        return pulumi.get(self, "on_push_reset_all_votes")

    @on_push_reset_all_votes.setter
    def on_push_reset_all_votes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_push_reset_all_votes", value)

    @property
    @pulumi.getter(name="onPushResetApprovedVotes")
    def on_push_reset_approved_votes(self) -> Optional[pulumi.Input[bool]]:
        """
        When new changes are pushed reset all approval votes (does not reset votes to reject or wait). Defaults to `false`.
        """
        return pulumi.get(self, "on_push_reset_approved_votes")

    @on_push_reset_approved_votes.setter
    def on_push_reset_approved_votes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_push_reset_approved_votes", value)

    @property
    @pulumi.getter(name="reviewerCount")
    def reviewer_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of reviewers needed to approve.
        """
        return pulumi.get(self, "reviewer_count")

    @reviewer_count.setter
    def reviewer_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reviewer_count", value)

    @property
    @pulumi.getter(name="submitterCanVote")
    def submitter_can_vote(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow requesters to approve their own changes. Defaults to `false`.
        """
        return pulumi.get(self, "submitter_can_vote")

    @submitter_can_vote.setter
    def submitter_can_vote(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "submitter_can_vote", value)


@pulumi.input_type
class BranchPolicyMinReviewersSettingsScopeArgs:
    def __init__(__self__, *,
                 match_type: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 repository_ref: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] match_type: The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
        :param pulumi.Input[str] repository_id: The repository ID. Needed only if the scope of the policy will be limited to a single repository.
        :param pulumi.Input[str] repository_ref: The ref pattern to use for the match. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        if match_type is not None:
            pulumi.set(__self__, "match_type", match_type)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if repository_ref is not None:
            pulumi.set(__self__, "repository_ref", repository_ref)

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> Optional[pulumi.Input[str]]:
        """
        The match type to use when applying the policy. Supported values are `Exact` (default) or `Prefix`.
        """
        return pulumi.get(self, "match_type")

    @match_type.setter
    def match_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_type", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        The repository ID. Needed only if the scope of the policy will be limited to a single repository.
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="repositoryRef")
    def repository_ref(self) -> Optional[pulumi.Input[str]]:
        """
        The ref pattern to use for the match. If `match_type` is `Exact`, this should be a qualified ref such as `refs/heads/master`. If `match_type` is `Prefix`, this should be a ref path such as `refs/heads/releases`.
        """
        return pulumi.get(self, "repository_ref")

    @repository_ref.setter
    def repository_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_ref", value)


