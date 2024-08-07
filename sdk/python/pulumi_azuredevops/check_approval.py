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

__all__ = ['CheckApprovalArgs', 'CheckApproval']

@pulumi.input_type
class CheckApprovalArgs:
    def __init__(__self__, *,
                 approvers: pulumi.Input[Sequence[pulumi.Input[str]]],
                 project_id: pulumi.Input[str],
                 target_resource_id: pulumi.Input[str],
                 target_resource_type: pulumi.Input[str],
                 instructions: Optional[pulumi.Input[str]] = None,
                 minimum_required_approvers: Optional[pulumi.Input[int]] = None,
                 requester_can_approve: Optional[pulumi.Input[bool]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a CheckApproval resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approvers: Specifies a list of approver IDs.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[str] instructions: The instructions for the approvers.
        :param pulumi.Input[int] minimum_required_approvers: The minimum number of approvers. This property is applicable when there is more than 1 approver.
        :param pulumi.Input[bool] requester_can_approve: Can the requestor approve? Defaults to `false`.
        :param pulumi.Input[int] timeout: The timeout in minutes for the approval.  Defaults to `43200`.
        """
        pulumi.set(__self__, "approvers", approvers)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "target_resource_id", target_resource_id)
        pulumi.set(__self__, "target_resource_type", target_resource_type)
        if instructions is not None:
            pulumi.set(__self__, "instructions", instructions)
        if minimum_required_approvers is not None:
            pulumi.set(__self__, "minimum_required_approvers", minimum_required_approvers)
        if requester_can_approve is not None:
            pulumi.set(__self__, "requester_can_approve", requester_can_approve)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def approvers(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies a list of approver IDs.
        """
        return pulumi.get(self, "approvers")

    @approvers.setter
    def approvers(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "approvers", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project ID. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_id", value)

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> pulumi.Input[str]:
        """
        The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "target_resource_type")

    @target_resource_type.setter
    def target_resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_resource_type", value)

    @property
    @pulumi.getter
    def instructions(self) -> Optional[pulumi.Input[str]]:
        """
        The instructions for the approvers.
        """
        return pulumi.get(self, "instructions")

    @instructions.setter
    def instructions(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instructions", value)

    @property
    @pulumi.getter(name="minimumRequiredApprovers")
    def minimum_required_approvers(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of approvers. This property is applicable when there is more than 1 approver.
        """
        return pulumi.get(self, "minimum_required_approvers")

    @minimum_required_approvers.setter
    def minimum_required_approvers(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_required_approvers", value)

    @property
    @pulumi.getter(name="requesterCanApprove")
    def requester_can_approve(self) -> Optional[pulumi.Input[bool]]:
        """
        Can the requestor approve? Defaults to `false`.
        """
        return pulumi.get(self, "requester_can_approve")

    @requester_can_approve.setter
    def requester_can_approve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requester_can_approve", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The timeout in minutes for the approval.  Defaults to `43200`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _CheckApprovalState:
    def __init__(__self__, *,
                 approvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instructions: Optional[pulumi.Input[str]] = None,
                 minimum_required_approvers: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 requester_can_approve: Optional[pulumi.Input[bool]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering CheckApproval resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approvers: Specifies a list of approver IDs.
        :param pulumi.Input[str] instructions: The instructions for the approvers.
        :param pulumi.Input[int] minimum_required_approvers: The minimum number of approvers. This property is applicable when there is more than 1 approver.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[bool] requester_can_approve: Can the requestor approve? Defaults to `false`.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the approval.  Defaults to `43200`.
        :param pulumi.Input[int] version: The version of the check.
        """
        if approvers is not None:
            pulumi.set(__self__, "approvers", approvers)
        if instructions is not None:
            pulumi.set(__self__, "instructions", instructions)
        if minimum_required_approvers is not None:
            pulumi.set(__self__, "minimum_required_approvers", minimum_required_approvers)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if requester_can_approve is not None:
            pulumi.set(__self__, "requester_can_approve", requester_can_approve)
        if target_resource_id is not None:
            pulumi.set(__self__, "target_resource_id", target_resource_id)
        if target_resource_type is not None:
            pulumi.set(__self__, "target_resource_type", target_resource_type)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def approvers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of approver IDs.
        """
        return pulumi.get(self, "approvers")

    @approvers.setter
    def approvers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "approvers", value)

    @property
    @pulumi.getter
    def instructions(self) -> Optional[pulumi.Input[str]]:
        """
        The instructions for the approvers.
        """
        return pulumi.get(self, "instructions")

    @instructions.setter
    def instructions(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instructions", value)

    @property
    @pulumi.getter(name="minimumRequiredApprovers")
    def minimum_required_approvers(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of approvers. This property is applicable when there is more than 1 approver.
        """
        return pulumi.get(self, "minimum_required_approvers")

    @minimum_required_approvers.setter
    def minimum_required_approvers(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_required_approvers", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="requesterCanApprove")
    def requester_can_approve(self) -> Optional[pulumi.Input[bool]]:
        """
        Can the requestor approve? Defaults to `false`.
        """
        return pulumi.get(self, "requester_can_approve")

    @requester_can_approve.setter
    def requester_can_approve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "requester_can_approve", value)

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @target_resource_id.setter
    def target_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_resource_id", value)

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "target_resource_type")

    @target_resource_type.setter
    def target_resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_resource_type", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The timeout in minutes for the approval.  Defaults to `43200`.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version of the check.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class CheckApproval(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instructions: Optional[pulumi.Input[str]] = None,
                 minimum_required_approvers: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 requester_can_approve: Optional[pulumi.Input[bool]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a Approval Check.

        ## Example Usage

        ### Protect an environment

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="Example Project")
        example_environment = azuredevops.Environment("example",
            project_id=example.id,
            name="Example Environment")
        example_group = azuredevops.Group("example", display_name="some-azdo-group")
        example_check_approval = azuredevops.CheckApproval("example",
            project_id=example.id,
            target_resource_id=example_environment.id,
            target_resource_type="environment",
            requester_can_approve=True,
            approvers=[example_group.origin_id])
        ```

        ## Import

        Importing this resource is not supported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approvers: Specifies a list of approver IDs.
        :param pulumi.Input[str] instructions: The instructions for the approvers.
        :param pulumi.Input[int] minimum_required_approvers: The minimum number of approvers. This property is applicable when there is more than 1 approver.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[bool] requester_can_approve: Can the requestor approve? Defaults to `false`.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the approval.  Defaults to `43200`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CheckApprovalArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Approval Check.

        ## Example Usage

        ### Protect an environment

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="Example Project")
        example_environment = azuredevops.Environment("example",
            project_id=example.id,
            name="Example Environment")
        example_group = azuredevops.Group("example", display_name="some-azdo-group")
        example_check_approval = azuredevops.CheckApproval("example",
            project_id=example.id,
            target_resource_id=example_environment.id,
            target_resource_type="environment",
            requester_can_approve=True,
            approvers=[example_group.origin_id])
        ```

        ## Import

        Importing this resource is not supported.

        :param str resource_name: The name of the resource.
        :param CheckApprovalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CheckApprovalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instructions: Optional[pulumi.Input[str]] = None,
                 minimum_required_approvers: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 requester_can_approve: Optional[pulumi.Input[bool]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CheckApprovalArgs.__new__(CheckApprovalArgs)

            if approvers is None and not opts.urn:
                raise TypeError("Missing required property 'approvers'")
            __props__.__dict__["approvers"] = approvers
            __props__.__dict__["instructions"] = instructions
            __props__.__dict__["minimum_required_approvers"] = minimum_required_approvers
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["requester_can_approve"] = requester_can_approve
            if target_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__.__dict__["target_resource_id"] = target_resource_id
            if target_resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_type'")
            __props__.__dict__["target_resource_type"] = target_resource_type
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["version"] = None
        super(CheckApproval, __self__).__init__(
            'azuredevops:index/checkApproval:CheckApproval',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            approvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instructions: Optional[pulumi.Input[str]] = None,
            minimum_required_approvers: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            requester_can_approve: Optional[pulumi.Input[bool]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None,
            target_resource_type: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'CheckApproval':
        """
        Get an existing CheckApproval resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approvers: Specifies a list of approver IDs.
        :param pulumi.Input[str] instructions: The instructions for the approvers.
        :param pulumi.Input[int] minimum_required_approvers: The minimum number of approvers. This property is applicable when there is more than 1 approver.
        :param pulumi.Input[str] project_id: The project ID. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[bool] requester_can_approve: Can the requestor approve? Defaults to `false`.
        :param pulumi.Input[str] target_resource_id: The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[str] target_resource_type: The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        :param pulumi.Input[int] timeout: The timeout in minutes for the approval.  Defaults to `43200`.
        :param pulumi.Input[int] version: The version of the check.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CheckApprovalState.__new__(_CheckApprovalState)

        __props__.__dict__["approvers"] = approvers
        __props__.__dict__["instructions"] = instructions
        __props__.__dict__["minimum_required_approvers"] = minimum_required_approvers
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["requester_can_approve"] = requester_can_approve
        __props__.__dict__["target_resource_id"] = target_resource_id
        __props__.__dict__["target_resource_type"] = target_resource_type
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["version"] = version
        return CheckApproval(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def approvers(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies a list of approver IDs.
        """
        return pulumi.get(self, "approvers")

    @property
    @pulumi.getter
    def instructions(self) -> pulumi.Output[Optional[str]]:
        """
        The instructions for the approvers.
        """
        return pulumi.get(self, "instructions")

    @property
    @pulumi.getter(name="minimumRequiredApprovers")
    def minimum_required_approvers(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum number of approvers. This property is applicable when there is more than 1 approver.
        """
        return pulumi.get(self, "minimum_required_approvers")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="requesterCanApprove")
    def requester_can_approve(self) -> pulumi.Output[Optional[bool]]:
        """
        Can the requestor approve? Defaults to `false`.
        """
        return pulumi.get(self, "requester_can_approve")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource being protected by the check. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "target_resource_id")

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> pulumi.Output[str]:
        """
        The type of resource being protected by the check. Valid values: `endpoint`, `environment`, `queue`, `repository`, `securefile`, `variablegroup`. Changing this forces a new Approval Check to be created.
        """
        return pulumi.get(self, "target_resource_type")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The timeout in minutes for the approval.  Defaults to `43200`.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The version of the check.
        """
        return pulumi.get(self, "version")

