# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['IterativePermissions']


class IterativePermissions(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages permissions for an Iteration (Sprint)

        > **Note** Permissions can be assigned to group principals and not to single user principals.

        ## Permission levels

        Permission for Iterations within Azure DevOps can be applied on two different levels.
        Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `path`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Terraform")
        project_readers = project.id.apply(lambda id: azuredevops.get_group(project_id=id,
            name="Readers"))
        root_permissions = azuredevops.IterativePermissions("root-permissions",
            project_id=project.id,
            principal=azuredevops_group["project-readers"]["id"],
            permissions={
                "CREATE_CHILDREN": "Deny",
                "GENERIC_READ": "NotSet",
                "DELETE": "Deny",
            })
        iteration_permissions = azuredevops.IterativePermissions("iteration-permissions",
            project_id=project.id,
            principal=azuredevops_group["project-readers"]["id"],
            path="Iteration 1",
            permissions={
                "CREATE_CHILDREN": "Allow",
                "GENERIC_READ": "NotSet",
                "DELETE": "Allow",
            })
        ```
        ## Relevant Links

        * [Azure DevOps Service REST API 5.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-5.1)

        ## PAT Permissions Required

        - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] path: The name of the branch to assign the permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
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

            __props__['path'] = path
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__['permissions'] = permissions
            if principal is None and not opts.urn:
                raise TypeError("Missing required property 'principal'")
            __props__['principal'] = principal
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['replace'] = replace
        super(IterativePermissions, __self__).__init__(
            'azuredevops:index/iterativePermissions:IterativePermissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            path: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            replace: Optional[pulumi.Input[bool]] = None) -> 'IterativePermissions':
        """
        Get an existing IterativePermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] path: The name of the branch to assign the permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project to assign the permissions.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["path"] = path
        __props__["permissions"] = permissions
        __props__["principal"] = principal
        __props__["project_id"] = project_id
        __props__["replace"] = replace
        return IterativePermissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the branch to assign the permissions.
        """
        return pulumi.get(self, "path")

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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

