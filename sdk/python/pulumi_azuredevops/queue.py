# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Queue']


class Queue(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an agent queue within Azure DevOps. In the UI, this is equivalent to adding an
        Organization defined pool to a project.

        The created queue is not authorized for use by all pipelines in the project. However,
        the `ResourceAuthorization` resource can be used to grant authorization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project")
        pool = azuredevops.get_pool(name="contoso-pool")
        queue = azuredevops.Queue("queue",
            project_id=project.id,
            agent_pool_id=pool.id)
        # Grant acccess to queue to all pipelines in the project
        auth = azuredevops.ResourceAuthorization("auth",
            project_id=project.id,
            resource_id=queue.id,
            type="queue",
            authorized=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Agent Queues](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues?view=azure-devops-rest-5.1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] agent_pool_id: The ID of the organization agent pool.
        :param pulumi.Input[str] project_id: The ID of the project in which to create the resource.
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

            if agent_pool_id is None:
                raise TypeError("Missing required property 'agent_pool_id'")
            __props__['agent_pool_id'] = agent_pool_id
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Agent/queue:Queue")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Queue, __self__).__init__(
            'azuredevops:index/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_pool_id: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'Queue':
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] agent_pool_id: The ID of the organization agent pool.
        :param pulumi.Input[str] project_id: The ID of the project in which to create the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["agent_pool_id"] = agent_pool_id
        __props__["project_id"] = project_id
        return Queue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPoolId")
    def agent_pool_id(self) -> pulumi.Output[int]:
        """
        The ID of the organization agent pool.
        """
        return pulumi.get(self, "agent_pool_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project in which to create the resource.
        """
        return pulumi.get(self, "project_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

