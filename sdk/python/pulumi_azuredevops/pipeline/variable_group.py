# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VariableGroupArgs', 'VariableGroup']

@pulumi.input_type
class VariableGroupArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 variables: pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]],
                 allow_access: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_vault: Optional[pulumi.Input['VariableGroupKeyVaultArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VariableGroup resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]] variables: One or more `variable` blocks as documented below.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input['VariableGroupKeyVaultArgs'] key_vault: A list of `key_vault` blocks as documented below.
        :param pulumi.Input[str] name: The name of the Variable Group.
        """
        VariableGroupArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            variables=variables,
            allow_access=allow_access,
            description=description,
            key_vault=key_vault,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: Optional[pulumi.Input[str]] = None,
             variables: Optional[pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]] = None,
             allow_access: Optional[pulumi.Input[bool]] = None,
             description: Optional[pulumi.Input[str]] = None,
             key_vault: Optional[pulumi.Input['VariableGroupKeyVaultArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")
        if variables is None:
            raise TypeError("Missing 'variables' argument")
        if allow_access is None and 'allowAccess' in kwargs:
            allow_access = kwargs['allowAccess']
        if key_vault is None and 'keyVault' in kwargs:
            key_vault = kwargs['keyVault']

        _setter("project_id", project_id)
        _setter("variables", variables)
        if allow_access is not None:
            _setter("allow_access", allow_access)
        if description is not None:
            _setter("description", description)
        if key_vault is not None:
            _setter("key_vault", key_vault)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]:
        """
        One or more `variable` blocks as documented below.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]):
        pulumi.set(self, "variables", value)

    @property
    @pulumi.getter(name="allowAccess")
    def allow_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean that indicate if this variable group is shared by all pipelines of this project.
        """
        return pulumi.get(self, "allow_access")

    @allow_access.setter
    def allow_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_access", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Variable Group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="keyVault")
    def key_vault(self) -> Optional[pulumi.Input['VariableGroupKeyVaultArgs']]:
        """
        A list of `key_vault` blocks as documented below.
        """
        return pulumi.get(self, "key_vault")

    @key_vault.setter
    def key_vault(self, value: Optional[pulumi.Input['VariableGroupKeyVaultArgs']]):
        pulumi.set(self, "key_vault", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Variable Group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _VariableGroupState:
    def __init__(__self__, *,
                 allow_access: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_vault: Optional[pulumi.Input['VariableGroupKeyVaultArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]] = None):
        """
        Input properties used for looking up and filtering VariableGroup resources.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input['VariableGroupKeyVaultArgs'] key_vault: A list of `key_vault` blocks as documented below.
        :param pulumi.Input[str] name: The name of the Variable Group.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]] variables: One or more `variable` blocks as documented below.
        """
        _VariableGroupState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allow_access=allow_access,
            description=description,
            key_vault=key_vault,
            name=name,
            project_id=project_id,
            variables=variables,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allow_access: Optional[pulumi.Input[bool]] = None,
             description: Optional[pulumi.Input[str]] = None,
             key_vault: Optional[pulumi.Input['VariableGroupKeyVaultArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             variables: Optional[pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if allow_access is None and 'allowAccess' in kwargs:
            allow_access = kwargs['allowAccess']
        if key_vault is None and 'keyVault' in kwargs:
            key_vault = kwargs['keyVault']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']

        if allow_access is not None:
            _setter("allow_access", allow_access)
        if description is not None:
            _setter("description", description)
        if key_vault is not None:
            _setter("key_vault", key_vault)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if variables is not None:
            _setter("variables", variables)

    @property
    @pulumi.getter(name="allowAccess")
    def allow_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean that indicate if this variable group is shared by all pipelines of this project.
        """
        return pulumi.get(self, "allow_access")

    @allow_access.setter
    def allow_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_access", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Variable Group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="keyVault")
    def key_vault(self) -> Optional[pulumi.Input['VariableGroupKeyVaultArgs']]:
        """
        A list of `key_vault` blocks as documented below.
        """
        return pulumi.get(self, "key_vault")

    @key_vault.setter
    def key_vault(self, value: Optional[pulumi.Input['VariableGroupKeyVaultArgs']]):
        pulumi.set(self, "key_vault", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Variable Group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]]:
        """
        One or more `variable` blocks as documented below.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VariableGroupVariableArgs']]]]):
        pulumi.set(self, "variables", value)


warnings.warn("""azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup""", DeprecationWarning)


class VariableGroup(pulumi.CustomResource):
    warnings.warn("""azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_access: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_vault: Optional[pulumi.Input[pulumi.InputType['VariableGroupKeyVaultArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]]] = None,
                 __props__=None):
        """
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
        - [Azure DevOps Service REST API 7.0 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Variable Groups**: Read, Create, & Manage
        - **Build**: Read & execute
        - **Project and Team**: Read
        - **Token Administration**: Read & manage
        - **Tokens**: Read & manage
        - **Work Items**: Read

        ## Import

        **Variable groups containing secret values cannot be imported.** Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.

        ```sh
         $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example "Example Project/10"
        ```

         or

        ```sh
         $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example 00000000-0000-0000-0000-000000000000/0
        ```

         _Note that for secret variables, the import command retrieve blank value in the tfstate._

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input[pulumi.InputType['VariableGroupKeyVaultArgs']] key_vault: A list of `key_vault` blocks as documented below.
        :param pulumi.Input[str] name: The name of the Variable Group.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]] variables: One or more `variable` blocks as documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VariableGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Relevant Links

        - [Azure DevOps Service REST API 7.0 - Variable Groups](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups?view=azure-devops-rest-7.0)
        - [Azure DevOps Service REST API 7.0 - Authorized Resources](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/authorizedresources?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Variable Groups**: Read, Create, & Manage
        - **Build**: Read & execute
        - **Project and Team**: Read
        - **Token Administration**: Read & manage
        - **Tokens**: Read & manage
        - **Work Items**: Read

        ## Import

        **Variable groups containing secret values cannot be imported.** Azure DevOps Variable groups can be imported using the project name/variable group ID or by the project Guid/variable group ID, e.g.

        ```sh
         $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example "Example Project/10"
        ```

         or

        ```sh
         $ pulumi import azuredevops:Pipeline/variableGroup:VariableGroup example 00000000-0000-0000-0000-000000000000/0
        ```

         _Note that for secret variables, the import command retrieve blank value in the tfstate._

        :param str resource_name: The name of the resource.
        :param VariableGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VariableGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VariableGroupArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_access: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_vault: Optional[pulumi.Input[pulumi.InputType['VariableGroupKeyVaultArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""VariableGroup is deprecated: azuredevops.pipeline.VariableGroup has been deprecated in favor of azuredevops.VariableGroup""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VariableGroupArgs.__new__(VariableGroupArgs)

            __props__.__dict__["allow_access"] = allow_access
            __props__.__dict__["description"] = description
            key_vault = _utilities.configure(key_vault, VariableGroupKeyVaultArgs, True)
            __props__.__dict__["key_vault"] = key_vault
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if variables is None and not opts.urn:
                raise TypeError("Missing required property 'variables'")
            __props__.__dict__["variables"] = variables
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
            variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]]] = None) -> 'VariableGroup':
        """
        Get an existing VariableGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_access: Boolean that indicate if this variable group is shared by all pipelines of this project.
        :param pulumi.Input[str] description: The description of the Variable Group.
        :param pulumi.Input[pulumi.InputType['VariableGroupKeyVaultArgs']] key_vault: A list of `key_vault` blocks as documented below.
        :param pulumi.Input[str] name: The name of the Variable Group.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VariableGroupVariableArgs']]]] variables: One or more `variable` blocks as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VariableGroupState.__new__(_VariableGroupState)

        __props__.__dict__["allow_access"] = allow_access
        __props__.__dict__["description"] = description
        __props__.__dict__["key_vault"] = key_vault
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["variables"] = variables
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
        """
        A list of `key_vault` blocks as documented below.
        """
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
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Sequence['outputs.VariableGroupVariable']]:
        """
        One or more `variable` blocks as documented below.
        """
        return pulumi.get(self, "variables")

