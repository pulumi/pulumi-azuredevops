# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectPermissionsArgs', 'ProjectPermissions']

@pulumi.input_type
class ProjectPermissionsArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 principal: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 replace: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ProjectPermissions resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available
               
               | Permission                   | Description                                  |
               |------------------------------|----------------------------------------------|
               | GENERIC_READ                 | View project-level information               |
               | GENERIC_WRITE                | Edit project-level information               |
               | DELETE                       | Delete team project                          |
               | PUBLISH_TEST_RESULTS         | Create test runs                             |
               | ADMINISTER_BUILD             | Administer a build                           |
               | START_BUILD                  | Start a build                                |
               | EDIT_BUILD_STATUS            | Edit build quality                           |
               | UPDATE_BUILD                 | Write to build operational store             |
               | DELETE_TEST_RESULTS          | Delete test runs                             |
               | VIEW_TEST_RESULTS            | View test runs                               |
               | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
               | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
               | WORK_ITEM_DELETE             | Delete and restore work items                |
               | WORK_ITEM_MOVE               | Move work items out of this project          |
               | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
               | RENAME                       | Rename team project                          |
               | MANAGE_PROPERTIES            | Manage project properties                    |
               | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
               | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
               | BYPASS_RULES                 | Bypass rules on work item updates            |
               | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
               | UPDATE_VISIBILITY            | Update project visibility                    |
               | CHANGE_PROCESS               | Change process of team project.              |
               | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
               | AGILETOOLS_PLANS             | Agile plans.                                 |
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        ProjectPermissionsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            permissions=permissions,
            principal=principal,
            project_id=project_id,
            replace=replace,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             principal: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             replace: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if permissions is None:
            raise TypeError("Missing 'permissions' argument")
        if principal is None:
            raise TypeError("Missing 'principal' argument")
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")

        _setter("permissions", permissions)
        _setter("principal", principal)
        _setter("project_id", project_id)
        if replace is not None:
            _setter("replace", replace)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        the permissions to assign. The following permissions are available

        | Permission                   | Description                                  |
        |------------------------------|----------------------------------------------|
        | GENERIC_READ                 | View project-level information               |
        | GENERIC_WRITE                | Edit project-level information               |
        | DELETE                       | Delete team project                          |
        | PUBLISH_TEST_RESULTS         | Create test runs                             |
        | ADMINISTER_BUILD             | Administer a build                           |
        | START_BUILD                  | Start a build                                |
        | EDIT_BUILD_STATUS            | Edit build quality                           |
        | UPDATE_BUILD                 | Write to build operational store             |
        | DELETE_TEST_RESULTS          | Delete test runs                             |
        | VIEW_TEST_RESULTS            | View test runs                               |
        | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
        | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
        | WORK_ITEM_DELETE             | Delete and restore work items                |
        | WORK_ITEM_MOVE               | Move work items out of this project          |
        | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
        | RENAME                       | Rename team project                          |
        | MANAGE_PROPERTIES            | Manage project properties                    |
        | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
        | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
        | BYPASS_RULES                 | Bypass rules on work item updates            |
        | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
        | UPDATE_VISIBILITY            | Update project visibility                    |
        | CHANGE_PROCESS               | Change process of team project.              |
        | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
        | AGILETOOLS_PLANS             | Agile plans.                                 |
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
        The ID of the project to assign the permissions.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def replace(self) -> Optional[pulumi.Input[bool]]:
        """
        Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        return pulumi.get(self, "replace")

    @replace.setter
    def replace(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace", value)


@pulumi.input_type
class _ProjectPermissionsState:
    def __init__(__self__, *,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ProjectPermissions resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available
               
               | Permission                   | Description                                  |
               |------------------------------|----------------------------------------------|
               | GENERIC_READ                 | View project-level information               |
               | GENERIC_WRITE                | Edit project-level information               |
               | DELETE                       | Delete team project                          |
               | PUBLISH_TEST_RESULTS         | Create test runs                             |
               | ADMINISTER_BUILD             | Administer a build                           |
               | START_BUILD                  | Start a build                                |
               | EDIT_BUILD_STATUS            | Edit build quality                           |
               | UPDATE_BUILD                 | Write to build operational store             |
               | DELETE_TEST_RESULTS          | Delete test runs                             |
               | VIEW_TEST_RESULTS            | View test runs                               |
               | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
               | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
               | WORK_ITEM_DELETE             | Delete and restore work items                |
               | WORK_ITEM_MOVE               | Move work items out of this project          |
               | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
               | RENAME                       | Rename team project                          |
               | MANAGE_PROPERTIES            | Manage project properties                    |
               | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
               | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
               | BYPASS_RULES                 | Bypass rules on work item updates            |
               | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
               | UPDATE_VISIBILITY            | Update project visibility                    |
               | CHANGE_PROCESS               | Change process of team project.              |
               | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
               | AGILETOOLS_PLANS             | Agile plans.                                 |
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        _ProjectPermissionsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            permissions=permissions,
            principal=principal,
            project_id=project_id,
            replace=replace,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             principal: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             replace: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']

        if permissions is not None:
            _setter("permissions", permissions)
        if principal is not None:
            _setter("principal", principal)
        if project_id is not None:
            _setter("project_id", project_id)
        if replace is not None:
            _setter("replace", replace)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        the permissions to assign. The following permissions are available

        | Permission                   | Description                                  |
        |------------------------------|----------------------------------------------|
        | GENERIC_READ                 | View project-level information               |
        | GENERIC_WRITE                | Edit project-level information               |
        | DELETE                       | Delete team project                          |
        | PUBLISH_TEST_RESULTS         | Create test runs                             |
        | ADMINISTER_BUILD             | Administer a build                           |
        | START_BUILD                  | Start a build                                |
        | EDIT_BUILD_STATUS            | Edit build quality                           |
        | UPDATE_BUILD                 | Write to build operational store             |
        | DELETE_TEST_RESULTS          | Delete test runs                             |
        | VIEW_TEST_RESULTS            | View test runs                               |
        | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
        | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
        | WORK_ITEM_DELETE             | Delete and restore work items                |
        | WORK_ITEM_MOVE               | Move work items out of this project          |
        | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
        | RENAME                       | Rename team project                          |
        | MANAGE_PROPERTIES            | Manage project properties                    |
        | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
        | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
        | BYPASS_RULES                 | Bypass rules on work item updates            |
        | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
        | UPDATE_VISIBILITY            | Update project visibility                    |
        | CHANGE_PROCESS               | Change process of team project.              |
        | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
        | AGILETOOLS_PLANS             | Agile plans.                                 |
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
        The ID of the project to assign the permissions.
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
        """
        return pulumi.get(self, "replace")

    @replace.setter
    def replace(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace", value)


class ProjectPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages permissions for a AzureDevOps project

        > **Note** Permissions can be assigned to group principals and not to single user principals.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_readers = azuredevops.get_group_output(project_id=example.id,
            name="Readers")
        example_permission = azuredevops.ProjectPermissions("example-permission",
            project_id=example.id,
            principal=example_readers.id,
            permissions={
                "DELETE": "Deny",
                "EDIT_BUILD_STATUS": "NotSet",
                "WORK_ITEM_MOVE": "Allow",
                "DELETE_TEST_RESULTS": "Deny",
            })
        ```
        ## Relevant Links

        * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available
               
               | Permission                   | Description                                  |
               |------------------------------|----------------------------------------------|
               | GENERIC_READ                 | View project-level information               |
               | GENERIC_WRITE                | Edit project-level information               |
               | DELETE                       | Delete team project                          |
               | PUBLISH_TEST_RESULTS         | Create test runs                             |
               | ADMINISTER_BUILD             | Administer a build                           |
               | START_BUILD                  | Start a build                                |
               | EDIT_BUILD_STATUS            | Edit build quality                           |
               | UPDATE_BUILD                 | Write to build operational store             |
               | DELETE_TEST_RESULTS          | Delete test runs                             |
               | VIEW_TEST_RESULTS            | View test runs                               |
               | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
               | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
               | WORK_ITEM_DELETE             | Delete and restore work items                |
               | WORK_ITEM_MOVE               | Move work items out of this project          |
               | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
               | RENAME                       | Rename team project                          |
               | MANAGE_PROPERTIES            | Manage project properties                    |
               | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
               | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
               | BYPASS_RULES                 | Bypass rules on work item updates            |
               | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
               | UPDATE_VISIBILITY            | Update project visibility                    |
               | CHANGE_PROCESS               | Change process of team project.              |
               | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
               | AGILETOOLS_PLANS             | Agile plans.                                 |
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages permissions for a AzureDevOps project

        > **Note** Permissions can be assigned to group principals and not to single user principals.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_readers = azuredevops.get_group_output(project_id=example.id,
            name="Readers")
        example_permission = azuredevops.ProjectPermissions("example-permission",
            project_id=example.id,
            principal=example_readers.id,
            permissions={
                "DELETE": "Deny",
                "EDIT_BUILD_STATUS": "NotSet",
                "WORK_ITEM_MOVE": "Allow",
                "DELETE_TEST_RESULTS": "Deny",
            })
        ```
        ## Relevant Links

        * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param ProjectPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProjectPermissionsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectPermissionsArgs.__new__(ProjectPermissionsArgs)

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
        super(ProjectPermissions, __self__).__init__(
            'azuredevops:index/projectPermissions:ProjectPermissions',
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
            replace: Optional[pulumi.Input[bool]] = None) -> 'ProjectPermissions':
        """
        Get an existing ProjectPermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available
               
               | Permission                   | Description                                  |
               |------------------------------|----------------------------------------------|
               | GENERIC_READ                 | View project-level information               |
               | GENERIC_WRITE                | Edit project-level information               |
               | DELETE                       | Delete team project                          |
               | PUBLISH_TEST_RESULTS         | Create test runs                             |
               | ADMINISTER_BUILD             | Administer a build                           |
               | START_BUILD                  | Start a build                                |
               | EDIT_BUILD_STATUS            | Edit build quality                           |
               | UPDATE_BUILD                 | Write to build operational store             |
               | DELETE_TEST_RESULTS          | Delete test runs                             |
               | VIEW_TEST_RESULTS            | View test runs                               |
               | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
               | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
               | WORK_ITEM_DELETE             | Delete and restore work items                |
               | WORK_ITEM_MOVE               | Move work items out of this project          |
               | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
               | RENAME                       | Rename team project                          |
               | MANAGE_PROPERTIES            | Manage project properties                    |
               | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
               | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
               | BYPASS_RULES                 | Bypass rules on work item updates            |
               | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
               | UPDATE_VISIBILITY            | Update project visibility                    |
               | CHANGE_PROCESS               | Change process of team project.              |
               | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
               | AGILETOOLS_PLANS             | Agile plans.                                 |
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectPermissionsState.__new__(_ProjectPermissionsState)

        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["principal"] = principal
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["replace"] = replace
        return ProjectPermissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Mapping[str, str]]:
        """
        the permissions to assign. The following permissions are available

        | Permission                   | Description                                  |
        |------------------------------|----------------------------------------------|
        | GENERIC_READ                 | View project-level information               |
        | GENERIC_WRITE                | Edit project-level information               |
        | DELETE                       | Delete team project                          |
        | PUBLISH_TEST_RESULTS         | Create test runs                             |
        | ADMINISTER_BUILD             | Administer a build                           |
        | START_BUILD                  | Start a build                                |
        | EDIT_BUILD_STATUS            | Edit build quality                           |
        | UPDATE_BUILD                 | Write to build operational store             |
        | DELETE_TEST_RESULTS          | Delete test runs                             |
        | VIEW_TEST_RESULTS            | View test runs                               |
        | MANAGE_TEST_ENVIRONMENTS     | Manage test environments                     |
        | MANAGE_TEST_CONFIGURATIONS   | Manage test configurations                   |
        | WORK_ITEM_DELETE             | Delete and restore work items                |
        | WORK_ITEM_MOVE               | Move work items out of this project          |
        | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items                |
        | RENAME                       | Rename team project                          |
        | MANAGE_PROPERTIES            | Manage project properties                    |
        | MANAGE_SYSTEM_PROPERTIES     | Manage system project properties             |
        | BYPASS_PROPERTY_CACHE        | Bypass project property cache                |
        | BYPASS_RULES                 | Bypass rules on work item updates            |
        | SUPPRESS_NOTIFICATIONS       | Suppress notifications for work item updates |
        | UPDATE_VISIBILITY            | Update project visibility                    |
        | CHANGE_PROCESS               | Change process of team project.              |
        | AGILETOOLS_BACKLOG           | Agile backlog management.                    |
        | AGILETOOLS_PLANS             | Agile plans.                                 |
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
        The ID of the project to assign the permissions.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def replace(self) -> pulumi.Output[Optional[bool]]:
        """
        Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        return pulumi.get(self, "replace")

