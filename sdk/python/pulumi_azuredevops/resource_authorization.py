# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['ResourceAuthorizationArgs', 'ResourceAuthorization']

@pulumi.input_type
class ResourceAuthorizationArgs:
    def __init__(__self__, *,
                 authorized: pulumi.Input[bool],
                 project_id: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 definition_id: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ResourceAuthorization resource.
        :param pulumi.Input[bool] authorized: Set to true to allow public access in the project. Type: boolean.
        :param pulumi.Input[str] project_id: The project ID or project name. Type: string.
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Type: string.
        :param pulumi.Input[int] definition_id: The ID of the build definition to authorize. Type: string.
        :param pulumi.Input[str] type: The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        """
        pulumi.set(__self__, "authorized", authorized)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "resource_id", resource_id)
        if definition_id is not None:
            pulumi.set(__self__, "definition_id", definition_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Input[bool]:
        """
        Set to true to allow public access in the project. Type: boolean.
        """
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: pulumi.Input[bool]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project ID or project name. Type: string.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource to authorize. Type: string.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="definitionId")
    def definition_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the build definition to authorize. Type: string.
        """
        return pulumi.get(self, "definition_id")

    @definition_id.setter
    def definition_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "definition_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class ResourceAuthorization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[bool]] = None,
                 definition_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages authorization of resources, e.g. for access in build pipelines.

        Currently supported resources: service endpoint (aka service connection, endpoint).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project")
        bitbucket_account = azuredevops.ServiceEndpointBitBucket("bitbucketAccount",
            project_id=project.id,
            username="xxxx",
            password="xxxx",
            service_endpoint_name="test-bitbucket",
            description="test")
        auth = azuredevops.ResourceAuthorization("auth",
            project_id=project.id,
            resource_id=bitbucket_account.id,
            authorized=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Authorize Definition Resource](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/resources/authorize%20definition%20resources?view=azure-devops-rest-5.1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] authorized: Set to true to allow public access in the project. Type: boolean.
        :param pulumi.Input[int] definition_id: The ID of the build definition to authorize. Type: string.
        :param pulumi.Input[str] project_id: The project ID or project name. Type: string.
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Type: string.
        :param pulumi.Input[str] type: The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceAuthorizationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages authorization of resources, e.g. for access in build pipelines.

        Currently supported resources: service endpoint (aka service connection, endpoint).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project")
        bitbucket_account = azuredevops.ServiceEndpointBitBucket("bitbucketAccount",
            project_id=project.id,
            username="xxxx",
            password="xxxx",
            service_endpoint_name="test-bitbucket",
            description="test")
        auth = azuredevops.ResourceAuthorization("auth",
            project_id=project.id,
            resource_id=bitbucket_account.id,
            authorized=True)
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 5.1 - Authorize Definition Resource](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/resources/authorize%20definition%20resources?view=azure-devops-rest-5.1)

        :param str resource_name: The name of the resource.
        :param ResourceAuthorizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceAuthorizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[bool]] = None,
                 definition_id: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if authorized is None and not opts.urn:
                raise TypeError("Missing required property 'authorized'")
            __props__['authorized'] = authorized
            __props__['definition_id'] = definition_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['type'] = type
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azuredevops:Security/resourceAuthorization:ResourceAuthorization")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ResourceAuthorization, __self__).__init__(
            'azuredevops:index/resourceAuthorization:ResourceAuthorization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized: Optional[pulumi.Input[bool]] = None,
            definition_id: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ResourceAuthorization':
        """
        Get an existing ResourceAuthorization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] authorized: Set to true to allow public access in the project. Type: boolean.
        :param pulumi.Input[int] definition_id: The ID of the build definition to authorize. Type: string.
        :param pulumi.Input[str] project_id: The project ID or project name. Type: string.
        :param pulumi.Input[str] resource_id: The ID of the resource to authorize. Type: string.
        :param pulumi.Input[str] type: The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized"] = authorized
        __props__["definition_id"] = definition_id
        __props__["project_id"] = project_id
        __props__["resource_id"] = resource_id
        __props__["type"] = type
        return ResourceAuthorization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Output[bool]:
        """
        Set to true to allow public access in the project. Type: boolean.
        """
        return pulumi.get(self, "authorized")

    @property
    @pulumi.getter(name="definitionId")
    def definition_id(self) -> pulumi.Output[Optional[int]]:
        """
        The ID of the build definition to authorize. Type: string.
        """
        return pulumi.get(self, "definition_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID or project name. Type: string.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource to authorize. Type: string.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the resource to authorize. Type: string. Valid values: `endpoint`, `queue`, `variablegroup`. Default value: `endpoint`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

