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

__all__ = ['ServicehookStorageQueuePipelinesArgs', 'ServicehookStorageQueuePipelines']

@pulumi.input_type
class ServicehookStorageQueuePipelinesArgs:
    def __init__(__self__, *,
                 account_key: pulumi.Input[builtins.str],
                 account_name: pulumi.Input[builtins.str],
                 project_id: pulumi.Input[builtins.str],
                 queue_name: pulumi.Input[builtins.str],
                 run_state_changed_event: Optional[pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs']] = None,
                 stage_state_changed_event: Optional[pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs']] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 visi_timeout: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a ServicehookStorageQueuePipelines resource.
        :param pulumi.Input[builtins.str] account_key: A valid account key from the queue's storage account.
        :param pulumi.Input[builtins.str] account_name: The queue's storage account name.
        :param pulumi.Input[builtins.str] project_id: The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        :param pulumi.Input[builtins.str] queue_name: The name of the queue that will store the events.
        :param pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs'] run_state_changed_event: A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        :param pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs'] stage_state_changed_event: A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`
               
               > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        :param pulumi.Input[builtins.int] ttl: event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        :param pulumi.Input[builtins.int] visi_timeout: event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        pulumi.set(__self__, "account_key", account_key)
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "queue_name", queue_name)
        if run_state_changed_event is not None:
            pulumi.set(__self__, "run_state_changed_event", run_state_changed_event)
        if stage_state_changed_event is not None:
            pulumi.set(__self__, "stage_state_changed_event", stage_state_changed_event)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if visi_timeout is not None:
            pulumi.set(__self__, "visi_timeout", visi_timeout)

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> pulumi.Input[builtins.str]:
        """
        A valid account key from the queue's storage account.
        """
        return pulumi.get(self, "account_key")

    @account_key.setter
    def account_key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "account_key", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[builtins.str]:
        """
        The queue's storage account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the queue that will store the events.
        """
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="runStateChangedEvent")
    def run_state_changed_event(self) -> Optional[pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs']]:
        """
        A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        """
        return pulumi.get(self, "run_state_changed_event")

    @run_state_changed_event.setter
    def run_state_changed_event(self, value: Optional[pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs']]):
        pulumi.set(self, "run_state_changed_event", value)

    @property
    @pulumi.getter(name="stageStateChangedEvent")
    def stage_state_changed_event(self) -> Optional[pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs']]:
        """
        A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`

        > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        """
        return pulumi.get(self, "stage_state_changed_event")

    @stage_state_changed_event.setter
    def stage_state_changed_event(self, value: Optional[pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs']]):
        pulumi.set(self, "stage_state_changed_event", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="visiTimeout")
    def visi_timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        return pulumi.get(self, "visi_timeout")

    @visi_timeout.setter
    def visi_timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "visi_timeout", value)


@pulumi.input_type
class _ServicehookStorageQueuePipelinesState:
    def __init__(__self__, *,
                 account_key: Optional[pulumi.Input[builtins.str]] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 queue_name: Optional[pulumi.Input[builtins.str]] = None,
                 run_state_changed_event: Optional[pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs']] = None,
                 stage_state_changed_event: Optional[pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs']] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 visi_timeout: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering ServicehookStorageQueuePipelines resources.
        :param pulumi.Input[builtins.str] account_key: A valid account key from the queue's storage account.
        :param pulumi.Input[builtins.str] account_name: The queue's storage account name.
        :param pulumi.Input[builtins.str] project_id: The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        :param pulumi.Input[builtins.str] queue_name: The name of the queue that will store the events.
        :param pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs'] run_state_changed_event: A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        :param pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs'] stage_state_changed_event: A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`
               
               > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        :param pulumi.Input[builtins.int] ttl: event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        :param pulumi.Input[builtins.int] visi_timeout: event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        if account_key is not None:
            pulumi.set(__self__, "account_key", account_key)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if queue_name is not None:
            pulumi.set(__self__, "queue_name", queue_name)
        if run_state_changed_event is not None:
            pulumi.set(__self__, "run_state_changed_event", run_state_changed_event)
        if stage_state_changed_event is not None:
            pulumi.set(__self__, "stage_state_changed_event", stage_state_changed_event)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if visi_timeout is not None:
            pulumi.set(__self__, "visi_timeout", visi_timeout)

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A valid account key from the queue's storage account.
        """
        return pulumi.get(self, "account_key")

    @account_key.setter
    def account_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "account_key", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The queue's storage account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the queue that will store the events.
        """
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="runStateChangedEvent")
    def run_state_changed_event(self) -> Optional[pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs']]:
        """
        A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        """
        return pulumi.get(self, "run_state_changed_event")

    @run_state_changed_event.setter
    def run_state_changed_event(self, value: Optional[pulumi.Input['ServicehookStorageQueuePipelinesRunStateChangedEventArgs']]):
        pulumi.set(self, "run_state_changed_event", value)

    @property
    @pulumi.getter(name="stageStateChangedEvent")
    def stage_state_changed_event(self) -> Optional[pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs']]:
        """
        A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`

        > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        """
        return pulumi.get(self, "stage_state_changed_event")

    @stage_state_changed_event.setter
    def stage_state_changed_event(self, value: Optional[pulumi.Input['ServicehookStorageQueuePipelinesStageStateChangedEventArgs']]):
        pulumi.set(self, "stage_state_changed_event", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="visiTimeout")
    def visi_timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        return pulumi.get(self, "visi_timeout")

    @visi_timeout.setter
    def visi_timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "visi_timeout", value)


class ServicehookStorageQueuePipelines(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_key: Optional[pulumi.Input[builtins.str]] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 queue_name: Optional[pulumi.Input[builtins.str]] = None,
                 run_state_changed_event: Optional[pulumi.Input[Union['ServicehookStorageQueuePipelinesRunStateChangedEventArgs', 'ServicehookStorageQueuePipelinesRunStateChangedEventArgsDict']]] = None,
                 stage_state_changed_event: Optional[pulumi.Input[Union['ServicehookStorageQueuePipelinesStageStateChangedEventArgs', 'ServicehookStorageQueuePipelinesStageStateChangedEventArgsDict']]] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 visi_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Manages a Storage Queue Pipelines Service Hook .

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="example-project")
        example_resource_group = azure.core.ResourceGroup("example",
            name="example-resources",
            location="West Europe")
        example_account = azure.storage.Account("example",
            name="servicehookexamplestacc",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_queue = azure.storage.Queue("example",
            name="examplequeue",
            storage_account_name=example_account.name)
        example_servicehook_storage_queue_pipelines = azuredevops.ServicehookStorageQueuePipelines("example",
            project_id=example.id,
            account_name=example_account.name,
            account_key=example_account.primary_access_key,
            queue_name=example_queue.name,
            visi_timeout=30,
            run_state_changed_event={
                "run_state_filter": "Completed",
                "run_result_filter": "Succeeded",
            })
        ```

        An empty configuration block will occur in all events triggering the associated action.

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.ServicehookStorageQueuePipelines("example",
            project_id=example_azuredevops_project["id"],
            account_name=example_azurerm_storage_account["name"],
            account_key=example_azurerm_storage_account["primaryAccessKey"],
            queue_name=example_azurerm_storage_queue["name"],
            visi_timeout=30,
            run_state_changed_event={})
        ```

        ## Import

        Storage Queue Pipelines Service Hook can be imported using the `resource id`, e.g.

        ```sh
        $ pulumi import azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_key: A valid account key from the queue's storage account.
        :param pulumi.Input[builtins.str] account_name: The queue's storage account name.
        :param pulumi.Input[builtins.str] project_id: The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        :param pulumi.Input[builtins.str] queue_name: The name of the queue that will store the events.
        :param pulumi.Input[Union['ServicehookStorageQueuePipelinesRunStateChangedEventArgs', 'ServicehookStorageQueuePipelinesRunStateChangedEventArgsDict']] run_state_changed_event: A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        :param pulumi.Input[Union['ServicehookStorageQueuePipelinesStageStateChangedEventArgs', 'ServicehookStorageQueuePipelinesStageStateChangedEventArgsDict']] stage_state_changed_event: A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`
               
               > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        :param pulumi.Input[builtins.int] ttl: event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        :param pulumi.Input[builtins.int] visi_timeout: event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServicehookStorageQueuePipelinesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Storage Queue Pipelines Service Hook .

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example", name="example-project")
        example_resource_group = azure.core.ResourceGroup("example",
            name="example-resources",
            location="West Europe")
        example_account = azure.storage.Account("example",
            name="servicehookexamplestacc",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_queue = azure.storage.Queue("example",
            name="examplequeue",
            storage_account_name=example_account.name)
        example_servicehook_storage_queue_pipelines = azuredevops.ServicehookStorageQueuePipelines("example",
            project_id=example.id,
            account_name=example_account.name,
            account_key=example_account.primary_access_key,
            queue_name=example_queue.name,
            visi_timeout=30,
            run_state_changed_event={
                "run_state_filter": "Completed",
                "run_result_filter": "Succeeded",
            })
        ```

        An empty configuration block will occur in all events triggering the associated action.

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.ServicehookStorageQueuePipelines("example",
            project_id=example_azuredevops_project["id"],
            account_name=example_azurerm_storage_account["name"],
            account_key=example_azurerm_storage_account["primaryAccessKey"],
            queue_name=example_azurerm_storage_queue["name"],
            visi_timeout=30,
            run_state_changed_event={})
        ```

        ## Import

        Storage Queue Pipelines Service Hook can be imported using the `resource id`, e.g.

        ```sh
        $ pulumi import azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServicehookStorageQueuePipelinesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServicehookStorageQueuePipelinesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_key: Optional[pulumi.Input[builtins.str]] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 queue_name: Optional[pulumi.Input[builtins.str]] = None,
                 run_state_changed_event: Optional[pulumi.Input[Union['ServicehookStorageQueuePipelinesRunStateChangedEventArgs', 'ServicehookStorageQueuePipelinesRunStateChangedEventArgsDict']]] = None,
                 stage_state_changed_event: Optional[pulumi.Input[Union['ServicehookStorageQueuePipelinesStageStateChangedEventArgs', 'ServicehookStorageQueuePipelinesStageStateChangedEventArgsDict']]] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 visi_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServicehookStorageQueuePipelinesArgs.__new__(ServicehookStorageQueuePipelinesArgs)

            if account_key is None and not opts.urn:
                raise TypeError("Missing required property 'account_key'")
            __props__.__dict__["account_key"] = None if account_key is None else pulumi.Output.secret(account_key)
            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if queue_name is None and not opts.urn:
                raise TypeError("Missing required property 'queue_name'")
            __props__.__dict__["queue_name"] = queue_name
            __props__.__dict__["run_state_changed_event"] = run_state_changed_event
            __props__.__dict__["stage_state_changed_event"] = stage_state_changed_event
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["visi_timeout"] = visi_timeout
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accountKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServicehookStorageQueuePipelines, __self__).__init__(
            'azuredevops:index/servicehookStorageQueuePipelines:ServicehookStorageQueuePipelines',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_key: Optional[pulumi.Input[builtins.str]] = None,
            account_name: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            queue_name: Optional[pulumi.Input[builtins.str]] = None,
            run_state_changed_event: Optional[pulumi.Input[Union['ServicehookStorageQueuePipelinesRunStateChangedEventArgs', 'ServicehookStorageQueuePipelinesRunStateChangedEventArgsDict']]] = None,
            stage_state_changed_event: Optional[pulumi.Input[Union['ServicehookStorageQueuePipelinesStageStateChangedEventArgs', 'ServicehookStorageQueuePipelinesStageStateChangedEventArgsDict']]] = None,
            ttl: Optional[pulumi.Input[builtins.int]] = None,
            visi_timeout: Optional[pulumi.Input[builtins.int]] = None) -> 'ServicehookStorageQueuePipelines':
        """
        Get an existing ServicehookStorageQueuePipelines resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_key: A valid account key from the queue's storage account.
        :param pulumi.Input[builtins.str] account_name: The queue's storage account name.
        :param pulumi.Input[builtins.str] project_id: The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        :param pulumi.Input[builtins.str] queue_name: The name of the queue that will store the events.
        :param pulumi.Input[Union['ServicehookStorageQueuePipelinesRunStateChangedEventArgs', 'ServicehookStorageQueuePipelinesRunStateChangedEventArgsDict']] run_state_changed_event: A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        :param pulumi.Input[Union['ServicehookStorageQueuePipelinesStageStateChangedEventArgs', 'ServicehookStorageQueuePipelinesStageStateChangedEventArgsDict']] stage_state_changed_event: A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`
               
               > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        :param pulumi.Input[builtins.int] ttl: event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        :param pulumi.Input[builtins.int] visi_timeout: event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServicehookStorageQueuePipelinesState.__new__(_ServicehookStorageQueuePipelinesState)

        __props__.__dict__["account_key"] = account_key
        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["queue_name"] = queue_name
        __props__.__dict__["run_state_changed_event"] = run_state_changed_event
        __props__.__dict__["stage_state_changed_event"] = stage_state_changed_event
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["visi_timeout"] = visi_timeout
        return ServicehookStorageQueuePipelines(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> pulumi.Output[builtins.str]:
        """
        A valid account key from the queue's storage account.
        """
        return pulumi.get(self, "account_key")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[builtins.str]:
        """
        The queue's storage account name.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the associated project. Changing this forces a new Service Hook Storage Queue Pipelines to be created.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the queue that will store the events.
        """
        return pulumi.get(self, "queue_name")

    @property
    @pulumi.getter(name="runStateChangedEvent")
    def run_state_changed_event(self) -> pulumi.Output[Optional['outputs.ServicehookStorageQueuePipelinesRunStateChangedEvent']]:
        """
        A `run_state_changed_event` block as defined below. Conflicts with `stage_state_changed_event`
        """
        return pulumi.get(self, "run_state_changed_event")

    @property
    @pulumi.getter(name="stageStateChangedEvent")
    def stage_state_changed_event(self) -> pulumi.Output[Optional['outputs.ServicehookStorageQueuePipelinesStageStateChangedEvent']]:
        """
        A `stage_state_changed_event` block as defined below. Conflicts with `run_state_changed_event`

        > **Note** At least one of `run_state_changed_event` and `stage_state_changed_event` has to be set.
        """
        return pulumi.get(self, "stage_state_changed_event")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        event time-to-live - the duration a message can remain in the queue before it's automatically removed. Defaults to `604800`.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="visiTimeout")
    def visi_timeout(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        event visibility timout - how long a message is invisible to other consumers after it's been dequeued. Defaults to `0`.
        """
        return pulumi.get(self, "visi_timeout")

