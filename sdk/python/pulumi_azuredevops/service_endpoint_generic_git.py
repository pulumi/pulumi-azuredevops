# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceEndpointGenericGitArgs', 'ServiceEndpointGenericGit']

@pulumi.input_type
class ServiceEndpointGenericGitArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 repository_url: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_pipelines_access: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceEndpointGenericGit resource.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] repository_url: The URL of the repository associated with the service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint.
        :param pulumi.Input[bool] enable_pipelines_access: A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        :param pulumi.Input[str] password: The PAT or password used to authenticate to the git repository.
        :param pulumi.Input[str] username: The username used to authenticate to the git repository.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "repository_url", repository_url)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_pipelines_access is not None:
            pulumi.set(__self__, "enable_pipelines_access", enable_pipelines_access)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

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
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Input[str]:
        """
        The URL of the repository associated with the service endpoint.
        """
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[str]:
        """
        The name of the service endpoint.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enablePipelinesAccess")
    def enable_pipelines_access(self) -> Optional[pulumi.Input[bool]]:
        """
        A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        """
        return pulumi.get(self, "enable_pipelines_access")

    @enable_pipelines_access.setter
    def enable_pipelines_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_pipelines_access", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The PAT or password used to authenticate to the git repository.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username used to authenticate to the git repository.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _ServiceEndpointGenericGitState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_pipelines_access: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_hash: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceEndpointGenericGit resources.
        :param pulumi.Input[bool] enable_pipelines_access: A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        :param pulumi.Input[str] password: The PAT or password used to authenticate to the git repository.
        :param pulumi.Input[str] password_hash: A bcrypted hash of the attribute 'password'
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] repository_url: The URL of the repository associated with the service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint.
        :param pulumi.Input[str] username: The username used to authenticate to the git repository.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_pipelines_access is not None:
            pulumi.set(__self__, "enable_pipelines_access", enable_pipelines_access)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_hash is not None:
            pulumi.set(__self__, "password_hash", password_hash)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repository_url is not None:
            pulumi.set(__self__, "repository_url", repository_url)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enablePipelinesAccess")
    def enable_pipelines_access(self) -> Optional[pulumi.Input[bool]]:
        """
        A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        """
        return pulumi.get(self, "enable_pipelines_access")

    @enable_pipelines_access.setter
    def enable_pipelines_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_pipelines_access", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The PAT or password used to authenticate to the git repository.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordHash")
    def password_hash(self) -> Optional[pulumi.Input[str]]:
        """
        A bcrypted hash of the attribute 'password'
        """
        return pulumi.get(self, "password_hash")

    @password_hash.setter
    def password_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_hash", value)

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
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the repository associated with the service endpoint.
        """
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service endpoint.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username used to authenticate to the git repository.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class ServiceEndpointGenericGit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_pipelines_access: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external git service
        using basic authentication via a username and password. This is mostly useful for importing private git repositories.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_generic_git = azuredevops.ServiceEndpointGenericGit("exampleServiceEndpointGenericGit",
            project_id=example_project.id,
            repository_url="https://dev.azure.com/org/project/_git/repository",
            username="username",
            password="password",
            service_endpoint_name="Example Generic Git",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)

        ## Import

        Azure DevOps Service Endpoint Generic Git can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_pipelines_access: A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        :param pulumi.Input[str] password: The PAT or password used to authenticate to the git repository.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] repository_url: The URL of the repository associated with the service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint.
        :param pulumi.Input[str] username: The username used to authenticate to the git repository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceEndpointGenericGitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external git service
        using basic authentication via a username and password. This is mostly useful for importing private git repositories.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_service_endpoint_generic_git = azuredevops.ServiceEndpointGenericGit("exampleServiceEndpointGenericGit",
            project_id=example_project.id,
            repository_url="https://dev.azure.com/org/project/_git/repository",
            username="username",
            password="password",
            service_endpoint_name="Example Generic Git",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 6.0 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)

        ## Import

        Azure DevOps Service Endpoint Generic Git can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceEndpointGenericGitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceEndpointGenericGitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_pipelines_access: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceEndpointGenericGitArgs.__new__(ServiceEndpointGenericGitArgs)

            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_pipelines_access"] = enable_pipelines_access
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if repository_url is None and not opts.urn:
                raise TypeError("Missing required property 'repository_url'")
            __props__.__dict__["repository_url"] = repository_url
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            __props__.__dict__["username"] = username
            __props__.__dict__["password_hash"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password", "passwordHash"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceEndpointGenericGit, __self__).__init__(
            'azuredevops:index/serviceEndpointGenericGit:ServiceEndpointGenericGit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_pipelines_access: Optional[pulumi.Input[bool]] = None,
            password: Optional[pulumi.Input[str]] = None,
            password_hash: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            repository_url: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'ServiceEndpointGenericGit':
        """
        Get an existing ServiceEndpointGenericGit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_pipelines_access: A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        :param pulumi.Input[str] password: The PAT or password used to authenticate to the git repository.
        :param pulumi.Input[str] password_hash: A bcrypted hash of the attribute 'password'
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] repository_url: The URL of the repository associated with the service endpoint.
        :param pulumi.Input[str] service_endpoint_name: The name of the service endpoint.
        :param pulumi.Input[str] username: The username used to authenticate to the git repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceEndpointGenericGitState.__new__(_ServiceEndpointGenericGitState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["description"] = description
        __props__.__dict__["enable_pipelines_access"] = enable_pipelines_access
        __props__.__dict__["password"] = password
        __props__.__dict__["password_hash"] = password_hash
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repository_url"] = repository_url
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["username"] = username
        return ServiceEndpointGenericGit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enablePipelinesAccess")
    def enable_pipelines_access(self) -> pulumi.Output[Optional[bool]]:
        """
        A value indicating whether or not to attempt accessing this git server from Azure Pipelines.
        """
        return pulumi.get(self, "enable_pipelines_access")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The PAT or password used to authenticate to the git repository.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordHash")
    def password_hash(self) -> pulumi.Output[str]:
        """
        A bcrypted hash of the attribute 'password'
        """
        return pulumi.get(self, "password_hash")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Output[str]:
        """
        The URL of the repository associated with the service endpoint.
        """
        return pulumi.get(self, "repository_url")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The name of the service endpoint.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The username used to authenticate to the git repository.
        """
        return pulumi.get(self, "username")

