# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceendpointPermissionsArgs', 'ServiceendpointPermissions']

@pulumi.input_type
class ServiceendpointPermissionsArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 principal: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 replace: Optional[pulumi.Input[bool]] = None,
                 serviceendpoint_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointPermissions resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | Use               | Use service endpoint                |
               | Administer        | Full control over service endpoints |
               | Create            | Create service endpoints            |
               | ViewAuthorization | View authorizations                 |
               | ViewEndpoint      | View service endpoint properties    |
        :param pulumi.Input[str] serviceendpoint_id: The id of the service endpoint to assign the permissions.
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "principal", principal)
        pulumi.set(__self__, "project_id", project_id)
        if replace is not None:
            pulumi.set(__self__, "replace", replace)
        if serviceendpoint_id is not None:
            pulumi.set(__self__, "serviceendpoint_id", serviceendpoint_id)

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
    @pulumi.getter
    def replace(self) -> Optional[pulumi.Input[bool]]:
        """
        Replace (`true`) or merge (`false`) the permissions. Default: `true`

        | Permission        | Description                         |
        | ----------------- | ----------------------------------- |
        | Use               | Use service endpoint                |
        | Administer        | Full control over service endpoints |
        | Create            | Create service endpoints            |
        | ViewAuthorization | View authorizations                 |
        | ViewEndpoint      | View service endpoint properties    |
        """
        return pulumi.get(self, "replace")

    @replace.setter
    def replace(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace", value)

    @property
    @pulumi.getter(name="serviceendpointId")
    def serviceendpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the service endpoint to assign the permissions.
        """
        return pulumi.get(self, "serviceendpoint_id")

    @serviceendpoint_id.setter
    def serviceendpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serviceendpoint_id", value)


@pulumi.input_type
class _ServiceendpointPermissionsState:
    def __init__(__self__, *,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 serviceendpoint_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointPermissions resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | Use               | Use service endpoint                |
               | Administer        | Full control over service endpoints |
               | Create            | Create service endpoints            |
               | ViewAuthorization | View authorizations                 |
               | ViewEndpoint      | View service endpoint properties    |
        :param pulumi.Input[str] serviceendpoint_id: The id of the service endpoint to assign the permissions.
        """
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if replace is not None:
            pulumi.set(__self__, "replace", replace)
        if serviceendpoint_id is not None:
            pulumi.set(__self__, "serviceendpoint_id", serviceendpoint_id)

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
        | Use               | Use service endpoint                |
        | Administer        | Full control over service endpoints |
        | Create            | Create service endpoints            |
        | ViewAuthorization | View authorizations                 |
        | ViewEndpoint      | View service endpoint properties    |
        """
        return pulumi.get(self, "replace")

    @replace.setter
    def replace(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replace", value)

    @property
    @pulumi.getter(name="serviceendpointId")
    def serviceendpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the service endpoint to assign the permissions.
        """
        return pulumi.get(self, "serviceendpoint_id")

    @serviceendpoint_id.setter
    def serviceendpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serviceendpoint_id", value)


class ServiceendpointPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 replace: Optional[pulumi.Input[bool]] = None,
                 serviceendpoint_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages permissions for a Service Endpoint

        > **Note** Permissions can be assigned to group principals and not to single user principals.

        ## Permission levels

        Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
        Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `serviceendpoint_id`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Terraform")
        example_readers = azuredevops.get_group_output(project_id=example.id,
            name="Readers")
        example_root_permissions = azuredevops.ServiceendpointPermissions("example-root-permissions",
            project_id=example.id,
            principal=example_readers.id,
            permissions={
                "Use": "allow",
                "Administer": "allow",
                "Create": "allow",
                "ViewAuthorization": "allow",
                "ViewEndpoint": "allow",
            })
        example_service_endpoint_docker_registry = azuredevops.ServiceEndpointDockerRegistry("example",
            project_id=example.id,
            service_endpoint_name="Example Docker Hub",
            docker_username="username",
            docker_email="email@example.com",
            docker_password="password",
            registry_type="DockerHub")
        example_permissions = azuredevops.ServiceendpointPermissions("example-permissions",
            project_id=example.id,
            principal=example_readers.id,
            serviceendpoint_id=example_service_endpoint_docker_registry.id,
            permissions={
                "Use": "allow",
                "Administer": "deny",
                "Create": "deny",
                "ViewAuthorization": "allow",
                "ViewEndpoint": "allow",
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
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] permissions: the permissions to assign. The following permissions are available.
        :param pulumi.Input[str] principal: The **group** principal to assign the permissions.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[bool] replace: Replace (`true`) or merge (`false`) the permissions. Default: `true`
               
               | Permission        | Description                         |
               | ----------------- | ----------------------------------- |
               | Use               | Use service endpoint                |
               | Administer        | Full control over service endpoints |
               | Create            | Create service endpoints            |
               | ViewAuthorization | View authorizations                 |
               | ViewEndpoint      | View service endpoint properties    |
        :param pulumi.Input[str] serviceendpoint_id: The id of the service endpoint to assign the permissions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages permissions for a Service Endpoint

        > **Note** Permissions can be assigned to group principals and not to single user principals.

        ## Permission levels

        Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
        Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `serviceendpoint_id`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Terraform")
        example_readers = azuredevops.get_group_output(project_id=example.id,
            name="Readers")
        example_root_permissions = azuredevops.ServiceendpointPermissions("example-root-permissions",
            project_id=example.id,
            principal=example_readers.id,
            permissions={
                "Use": "allow",
                "Administer": "allow",
                "Create": "allow",
                "ViewAuthorization": "allow",
                "ViewEndpoint": "allow",
            })
        example_service_endpoint_docker_registry = azuredevops.ServiceEndpointDockerRegistry("example",
            project_id=example.id,
            service_endpoint_name="Example Docker Hub",
            docker_username="username",
            docker_email="email@example.com",
            docker_password="password",
            registry_type="DockerHub")
        example_permissions = azuredevops.ServiceendpointPermissions("example-permissions",
            project_id=example.id,
            principal=example_readers.id,
            serviceendpoint_id=example_service_endpoint_docker_registry.id,
            permissions={
                "Use": "allow",
                "Administer": "deny",
                "Create": "deny",
                "ViewAuthorization": "allow",
                "ViewEndpoint": "allow",
            })
        ```

        ## Relevant Links

        * [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)

        ## PAT Permissions Required

        - **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param ServiceendpointPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 serviceendpoint_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointPermissionsArgs.__new__(ServiceendpointPermissionsArgs)

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
            __props__.__dict__["serviceendpoint_id"] = serviceendpoint_id
        super(ServiceendpointPermissions, __self__).__init__(
            'azuredevops:index/serviceendpointPermissions:ServiceendpointPermissions',
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
            serviceendpoint_id: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointPermissions':
        """
        Get an existing ServiceendpointPermissions resource's state with the given name, id, and optional extra
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
               | Use               | Use service endpoint                |
               | Administer        | Full control over service endpoints |
               | Create            | Create service endpoints            |
               | ViewAuthorization | View authorizations                 |
               | ViewEndpoint      | View service endpoint properties    |
        :param pulumi.Input[str] serviceendpoint_id: The id of the service endpoint to assign the permissions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointPermissionsState.__new__(_ServiceendpointPermissionsState)

        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["principal"] = principal
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["replace"] = replace
        __props__.__dict__["serviceendpoint_id"] = serviceendpoint_id
        return ServiceendpointPermissions(resource_name, opts=opts, __props__=__props__)

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
        | Use               | Use service endpoint                |
        | Administer        | Full control over service endpoints |
        | Create            | Create service endpoints            |
        | ViewAuthorization | View authorizations                 |
        | ViewEndpoint      | View service endpoint properties    |
        """
        return pulumi.get(self, "replace")

    @property
    @pulumi.getter(name="serviceendpointId")
    def serviceendpoint_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of the service endpoint to assign the permissions.
        """
        return pulumi.get(self, "serviceendpoint_id")

