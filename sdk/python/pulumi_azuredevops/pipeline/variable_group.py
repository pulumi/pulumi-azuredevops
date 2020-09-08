# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VariableGroup']


class VariableGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_access: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_vault: Optional[pulumi.Input[pulumi.InputType['VariableGroupKeyVaultArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages variable groups within Azure DevOps.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.core.Project("project", project_name="Test Project")
        variablegroup = azuredevops.pipeline.VariableGroup("variablegroup",
            project_id=project.id,
            description="Test Variable Group Description",
            allow_access=True,
            variables=[
                azuredevops.pipeline.VariableGroupVariableArgs(
                    name="key",
                    value="value",
                ),
                azuredevops.pipeline.VariableGroupVariableArgs(
                    name="Account Password",
                    value="p@ssword123",
                    is_secret=True,
                ),
            ])
        ```
        ## Relevant Links

        * [Azure DevOps Service REST API 5.1 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-5.1)
        * [Azure DevOps Service REST API 5.1 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-5.1)

        ## PAT Permissions Required

        - **Variable Groups**: Read, Create, & Manage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input[str] name: The name of the Variable Group.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]] variables: One or more `variable` blocks as documented below.
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

            __props__['allow_access'] = allow_access
            __props__['description'] = description
            __props__['key_vault'] = key_vault
            __props__['name'] = name
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            if variables is None:
                raise TypeError("Missing required property 'variables'")
            __props__['variables'] = variables
        super(VariableGroup, __self__).__init__(
            'azuredevops:Pipeline/variableGroup:VariableGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_access: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            key_vault: Optional[pulumi.Input[pulumi.InputType['VariableGroupKeyVaultArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            variables: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]]] = None) -> 'VariableGroup':
        """
        Get an existing VariableGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input[str] name: The name of the Variable Group.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]] variables: One or more `variable` blocks as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_access"] = allow_access
        __props__["description"] = description
        __props__["key_vault"] = key_vault
        __props__["name"] = name
        __props__["project_id"] = project_id
        __props__["variables"] = variables
        return VariableGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowAccess")
    def allow_access(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean that indicate if this variable group is shared by all pipelines of this project.
        """
        return pulumi.get(self, "allow_access")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Variable Group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="keyVault")
    def key_vault(self) -> pulumi.Output[Optional['outputs.VariableGroupKeyVault']]:
        return pulumi.get(self, "key_vault")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Variable Group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[List['outputs.VariableGroupVariable']]:
        """
        One or more `variable` blocks as documented below.
        """
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

