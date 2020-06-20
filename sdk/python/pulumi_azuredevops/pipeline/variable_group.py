# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VariableGroup(pulumi.CustomResource):
    allow_access: pulumi.Output[bool]
    """
    Boolean that indicate if this variable group is shared by all pipelines of this project.
    """
    description: pulumi.Output[str]
    """
    The description of the Variable Group.
    """
    key_vault: pulumi.Output[dict]
    name: pulumi.Output[str]
    """
    The name of the Variable Group.
    """
    project_id: pulumi.Output[str]
    """
    The project ID or project name.
    """
    variables: pulumi.Output[list]
    """
    One or more `variable` blocks as documented below.

      * `contentType` (`str`)
      * `enabled` (`bool`)
      * `expires` (`str`)
      * `isSecret` (`bool`) - A boolean flag describing if the variable value is sensitive. Defaults to `false`.
      * `name` (`str`) - The key value used for the variable. Must be unique within the Variable Group.
      * `secretValue` (`str`) - The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
      * `value` (`str`) - The value of the variable. If omitted, it will default to empty string.
    """
    def __init__(__self__, resource_name, opts=None, allow_access=None, description=None, key_vault=None, name=None, project_id=None, variables=None, __props__=None, __name__=None, __opts__=None):
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
            variable=[
                {
                    "name": "key",
                    "value": "value",
                },
                {
                    "name": "Account Password",
                    "value": "p@ssword123",
                    "isSecret": True,
                },
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
        :param pulumi.Input[list] variables: One or more `variable` blocks as documented below.

        The **key_vault** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the Variable Group.
          * `serviceEndpointId` (`pulumi.Input[str]`)

        The **variables** object supports the following:

          * `contentType` (`pulumi.Input[str]`)
          * `enabled` (`pulumi.Input[bool]`)
          * `expires` (`pulumi.Input[str]`)
          * `isSecret` (`pulumi.Input[bool]`) - A boolean flag describing if the variable value is sensitive. Defaults to `false`.
          * `name` (`pulumi.Input[str]`) - The key value used for the variable. Must be unique within the Variable Group.
          * `secretValue` (`pulumi.Input[str]`) - The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
          * `value` (`pulumi.Input[str]`) - The value of the variable. If omitted, it will default to empty string.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, allow_access=None, description=None, key_vault=None, name=None, project_id=None, variables=None):
        """
        Get an existing VariableGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input[str] name: The name of the Variable Group.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[list] variables: One or more `variable` blocks as documented below.

        The **key_vault** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the Variable Group.
          * `serviceEndpointId` (`pulumi.Input[str]`)

        The **variables** object supports the following:

          * `contentType` (`pulumi.Input[str]`)
          * `enabled` (`pulumi.Input[bool]`)
          * `expires` (`pulumi.Input[str]`)
          * `isSecret` (`pulumi.Input[bool]`) - A boolean flag describing if the variable value is sensitive. Defaults to `false`.
          * `name` (`pulumi.Input[str]`) - The key value used for the variable. Must be unique within the Variable Group.
          * `secretValue` (`pulumi.Input[str]`) - The secret value of the variable. If omitted, it will default to empty string. Used when `is_secret` set to `true`.
          * `value` (`pulumi.Input[str]`) - The value of the variable. If omitted, it will default to empty string.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

