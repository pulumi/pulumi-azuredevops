# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VariableGroupPermissionsArgs', 'VariableGroupPermissions']

@pulumi.input_type
class VariableGroupPermissionsArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 principal: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 variable_group_id: pulumi.Input[str],
                 replace: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a VariableGroupPermissions resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] variable_group_id: The id of the variable group to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | View              | View library item                   |
               | Administer        | Administer library item             |
               | Create            | Create library item                 |
               | ViewSecrets       | View library item secrets           |
               | Use               | Use library item                    |
               | Owner             | Owner library item                  |
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "principal", principal)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "variable_group_id", variable_group_id)
        if replace is not None:
            pulumi.set(__self__, "replace", replace)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        the permissions to assign. The following permissions are available.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Input[str]:
        """
        The **group** principal to assign the permissions.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal", value)

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
    @pulumi.getter(name="variableGroupId")
    def variable_group_id(self) -> pulumi.Input[str]:
        """
        The id of the variable group to assign the permissions.
        """
        return pulumi.get(self, "variable_group_id")

    @variable_group_id.setter
    def variable_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "variable_group_id", value)

    @property
    @pulumi.getter
    def replace(self) -> Optional[pulumi.Input[bool]]:
        """
        Replace (`true`) or merge (`false`) the permissions. Default: `true`

        | Permission        | Description                         |
        | ----------------- | ----------------------------------- |
        | View              | View library item                   |
        | Administer        | Administer library item             |
        | Create            | Create library item                 |
        | ViewSecrets       | View library item secrets           |
        | Use               | Use library item                    |
        | Owner             | Owner library item                  |
        """
        return pulumi.get(self, "replace")

    @replace.setter
    def replace(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace", value)


@pulumi.input_type
class _VariableGroupPermissionsState:
    def __init__(__self__, *,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 variable_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VariableGroupPermissions resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | View              | View library item                   |
               | Administer        | Administer library item             |
               | Create            | Create library item                 |
               | ViewSecrets       | View library item secrets           |
               | Use               | Use library item                    |
               | Owner             | Owner library item                  |
        :param pulumi.Input[str] variable_group_id: The id of the variable group to assign the permissions.
        """
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if replace is not None:
            pulumi.set(__self__, "replace", replace)
        if variable_group_id is not None:
            pulumi.set(__self__, "variable_group_id", variable_group_id)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        the permissions to assign. The following permissions are available.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        """
        The **group** principal to assign the permissions.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

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
    def replace(self) -> Optional[pulumi.Input[bool]]:
        """
        Replace (`true`) or merge (`false`) the permissions. Default: `true`

        | Permission        | Description                         |
        | ----------------- | ----------------------------------- |
        | View              | View library item                   |
        | Administer        | Administer library item             |
        | Create            | Create library item                 |
        | ViewSecrets       | View library item secrets           |
        | Use               | Use library item                    |
        | Owner             | Owner library item                  |
        """
        return pulumi.get(self, "replace")

    @replace.setter
    def replace(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace", value)

    @property
    @pulumi.getter(name="variableGroupId")
    def variable_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the variable group to assign the permissions.
        """
        return pulumi.get(self, "variable_group_id")

    @variable_group_id.setter
    def variable_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_group_id", value)


class VariableGroupPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 variable_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages permissions for a Variable Group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            name="Testing",
            description="Testing-description",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example = azuredevops.VariableGroup("example",
            project_id=project.id,
            name="test",
            description="Test Description",
            allow_access=True,
            variables=[azuredevops.VariableGroupVariableArgs(
                name="key1",
                value="val1",
            )])
        tf_project_readers = azuredevops.get_group_output(project_id=project.id,
            name="Readers")
        permissions = azuredevops.VariableGroupPermissions("permissions",
            project_id=project.id,
            variable_group_id=example.id,
            principal=tf_project_readers.id,
            permissions={
                "View": "allow",
                "Administer": "allow",
                "Use": "allow",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Roles

        The Azure DevOps UI uses roles to assign permissions for variable groups.

        | Role          | Allow Permissions      |
        | ------------- | ---------------------- |
        | Reader        | View                   |
        | User          | View, Use              |
        | Administrator | View, Use, Administer  |

        ## Relevant Links

        * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)

        ## PAT Permissions Required

        - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | View              | View library item                   |
               | Administer        | Administer library item             |
               | Create            | Create library item                 |
               | ViewSecrets       | View library item secrets           |
               | Use               | Use library item                    |
               | Owner             | Owner library item                  |
        :param pulumi.Input[str] variable_group_id: The id of the variable group to assign the permissions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VariableGroupPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages permissions for a Variable Group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            name="Testing",
            description="Testing-description",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        example = azuredevops.VariableGroup("example",
            project_id=project.id,
            name="test",
            description="Test Description",
            allow_access=True,
            variables=[azuredevops.VariableGroupVariableArgs(
                name="key1",
                value="val1",
            )])
        tf_project_readers = azuredevops.get_group_output(project_id=project.id,
            name="Readers")
        permissions = azuredevops.VariableGroupPermissions("permissions",
            project_id=project.id,
            variable_group_id=example.id,
            principal=tf_project_readers.id,
            permissions={
                "View": "allow",
                "Administer": "allow",
                "Use": "allow",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Roles

        The Azure DevOps UI uses roles to assign permissions for variable groups.

        | Role          | Allow Permissions      |
        | ------------- | ---------------------- |
        | Reader        | View                   |
        | User          | View, Use              |
        | Administrator | View, Use, Administer  |

        ## Relevant Links

        * [Azure DevOps Service REST API 6.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-6.0)

        ## PAT Permissions Required

        - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param VariableGroupPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VariableGroupPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 variable_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VariableGroupPermissionsArgs.__new__(VariableGroupPermissionsArgs)

            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            if principal is None and not opts.urn:
                raise TypeError("Missing required property 'principal'")
            __props__.__dict__["principal"] = principal
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["replace"] = replace
            if variable_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'variable_group_id'")
            __props__.__dict__["variable_group_id"] = variable_group_id
        super(VariableGroupPermissions, __self__).__init__(
            'azuredevops:index/variableGroupPermissions:VariableGroupPermissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            replace: Optional[pulumi.Input[bool]] = None,
            variable_group_id: Optional[pulumi.Input[str]] = None) -> 'VariableGroupPermissions':
        """
        Get an existing VariableGroupPermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | View              | View library item                   |
               | Administer        | Administer library item             |
               | Create            | Create library item                 |
               | ViewSecrets       | View library item secrets           |
               | Use               | Use library item                    |
               | Owner             | Owner library item                  |
        :param pulumi.Input[str] variable_group_id: The id of the variable group to assign the permissions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VariableGroupPermissionsState.__new__(_VariableGroupPermissionsState)

        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["principal"] = principal
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["replace"] = replace
        __props__.__dict__["variable_group_id"] = variable_group_id
        return VariableGroupPermissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Mapping[str, str]]:
        """
        the permissions to assign. The following permissions are available.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[str]:
        """
        The **group** principal to assign the permissions.
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def replace(self) -> pulumi.Output[Optional[bool]]:
        """
        Replace (`true`) or merge (`false`) the permissions. Default: `true`

        | Permission        | Description                         |
        | ----------------- | ----------------------------------- |
        | View              | View library item                   |
        | Administer        | Administer library item             |
        | Create            | Create library item                 |
        | ViewSecrets       | View library item secrets           |
        | Use               | Use library item                    |
        | Owner             | Owner library item                  |
        """
        return pulumi.get(self, "replace")

    @property
    @pulumi.getter(name="variableGroupId")
    def variable_group_id(self) -> pulumi.Output[str]:
        """
        The id of the variable group to assign the permissions.
        """
        return pulumi.get(self, "variable_group_id")

