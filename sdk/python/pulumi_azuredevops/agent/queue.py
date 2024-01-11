# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['QueueArgs', 'Queue']

@pulumi.input_type
class QueueArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 agent_pool_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Queue resource.
        :param pulumi.Input[str] project_id: The ID of the project in which to create the resource.
        :param pulumi.Input[int] agent_pool_id: The ID of the organization agent pool. Conflicts with `name`.
               
               > **NOTE:**
               One of `name` or `agent_pool_id` must be specified, but not both.
               When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        :param pulumi.Input[str] name: The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        """
        pulumi.set(__self__, "project_id", project_id)
        if agent_pool_id is not None:
            pulumi.set(__self__, "agent_pool_id", agent_pool_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project in which to create the resource.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="agentPoolId")
    def agent_pool_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the organization agent pool. Conflicts with `name`.

        > **NOTE:**
        One of `name` or `agent_pool_id` must be specified, but not both.
        When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        """
        return pulumi.get(self, "agent_pool_id")

    @agent_pool_id.setter
    def agent_pool_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "agent_pool_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _QueueState:
    def __init__(__self__, *,
                 agent_pool_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Queue resources.
        :param pulumi.Input[int] agent_pool_id: The ID of the organization agent pool. Conflicts with `name`.
               
               > **NOTE:**
               One of `name` or `agent_pool_id` must be specified, but not both.
               When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        :param pulumi.Input[str] name: The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        :param pulumi.Input[str] project_id: The ID of the project in which to create the resource.
        """
        if agent_pool_id is not None:
            pulumi.set(__self__, "agent_pool_id", agent_pool_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="agentPoolId")
    def agent_pool_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the organization agent pool. Conflicts with `name`.

        > **NOTE:**
        One of `name` or `agent_pool_id` must be specified, but not both.
        When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        """
        return pulumi.get(self, "agent_pool_id")

    @agent_pool_id.setter
    def agent_pool_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "agent_pool_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which to create the resource.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


warnings.warn("""azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue""", DeprecationWarning)


class Queue(pulumi.CustomResource):
    warnings.warn("""azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
        Organization defined pool to a project.

        The created queue is not authorized for use by all pipelines in the project. However,
        the `ResourceAuthorization` resource can be used to grant authorization.

        ## Example Usage
        ### Creating a Queue from an organization-level pool

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_pool = azuredevops.get_pool(name="example-pool")
        example_queue = azuredevops.Queue("exampleQueue",
            project_id=example_project.id,
            agent_pool_id=example_pool.id)
        # Grant access to queue to all pipelines in the project
        example_resource_authorization = azuredevops.ResourceAuthorization("exampleResourceAuthorization",
            project_id=example_project.id,
            resource_id=example_queue.id,
            type="queue",
            authorized=True)
        ```
        ### Creating a Queue at the project level (Organization-level permissions not required)

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.get_project(name="Example Project")
        example_queue = azuredevops.Queue("exampleQueue", project_id=example_project.id)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Agent Pools can be imported using the project ID and agent queue ID, e.g.

        ```sh
         $ pulumi import azuredevops:Agent/queue:Queue example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] agent_pool_id: The ID of the organization agent pool. Conflicts with `name`.
               
               > **NOTE:**
               One of `name` or `agent_pool_id` must be specified, but not both.
               When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        :param pulumi.Input[str] name: The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        :param pulumi.Input[str] project_id: The ID of the project in which to create the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QueueArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
        Organization defined pool to a project.

        The created queue is not authorized for use by all pipelines in the project. However,
        the `ResourceAuthorization` resource can be used to grant authorization.

        ## Example Usage
        ### Creating a Queue from an organization-level pool

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject")
        example_pool = azuredevops.get_pool(name="example-pool")
        example_queue = azuredevops.Queue("exampleQueue",
            project_id=example_project.id,
            agent_pool_id=example_pool.id)
        # Grant access to queue to all pipelines in the project
        example_resource_authorization = azuredevops.ResourceAuthorization("exampleResourceAuthorization",
            project_id=example_project.id,
            resource_id=example_queue.id,
            type="queue",
            authorized=True)
        ```
        ### Creating a Queue at the project level (Organization-level permissions not required)

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.get_project(name="Example Project")
        example_queue = azuredevops.Queue("exampleQueue", project_id=example_project.id)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-7.0)

        ## Import

        Azure DevOps Agent Pools can be imported using the project ID and agent queue ID, e.g.

        ```sh
         $ pulumi import azuredevops:Agent/queue:Queue example 00000000-0000-0000-0000-000000000000/0
        ```

        :param str resource_name: The name of the resource.
        :param QueueArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QueueArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Queue is deprecated: azuredevops.agent.Queue has been deprecated in favor of azuredevops.Queue""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QueueArgs.__new__(QueueArgs)

            __props__.__dict__["agent_pool_id"] = agent_pool_id
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
        super(Queue, __self__).__init__(
            'azuredevops:Agent/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_pool_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'Queue':
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] agent_pool_id: The ID of the organization agent pool. Conflicts with `name`.
               
               > **NOTE:**
               One of `name` or `agent_pool_id` must be specified, but not both.
               When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        :param pulumi.Input[str] name: The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        :param pulumi.Input[str] project_id: The ID of the project in which to create the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QueueState.__new__(_QueueState)

        __props__.__dict__["agent_pool_id"] = agent_pool_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        return Queue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPoolId")
    def agent_pool_id(self) -> pulumi.Output[int]:
        """
        The ID of the organization agent pool. Conflicts with `name`.

        > **NOTE:**
        One of `name` or `agent_pool_id` must be specified, but not both.
        When `agent_pool_id` is specified, the agent queue name will be derived from the agent pool name.
        """
        return pulumi.get(self, "agent_pool_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the agent queue. Defaults to the ID of the agent pool. Conflicts with `agent_pool_id`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project in which to create the resource.
        """
        return pulumi.get(self, "project_id")

