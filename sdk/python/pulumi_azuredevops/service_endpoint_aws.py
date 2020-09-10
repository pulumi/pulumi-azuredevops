# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['ServiceEndpointAws']


class ServiceEndpointAws(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key_id: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 role_session_name: Optional[pulumi.Input[str]] = None,
                 role_to_assume: Optional[pulumi.Input[str]] = None,
                 secret_access_key: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 session_token: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a AWS service endpoint within Azure DevOps. Using this service endpoint requires you to first install [AWS Toolkit for Azure DevOps](https://marketplace.visualstudio.com/items?itemName=AmazonWebServices.aws-vsts-tools).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        project = azuredevops.Project("project",
            visibility="private",
            version_control="Git",
            work_item_template="Agile")
        serviceendpoint = azuredevops.ServiceEndpointAws("serviceendpoint",
            project_id=project.id,
            service_endpoint_name="aws_serviceendpoint",
            description="Managed by AzureDevOps",
            access_key_id="xxxx",
            secret_access_key="xxxx")
        ```
        ## Relevant Links

        * [aws-toolkit-azure-devops](https://github.com/aws/aws-toolkit-azure-devops)
        * [Azure DevOps Service REST API 5.1 - Agent Pools](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: The AWS access key ID for signing programmatic requests.
        :param pulumi.Input[str] external_id: A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] role_session_name: Optional identifier for the assumed role session.
        :param pulumi.Input[str] role_to_assume: The Amazon Resource Name (ARN) of the role to assume.
        :param pulumi.Input[str] secret_access_key: The AWS secret access key for signing programmatic requests.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] session_token: The AWS session token for signing programmatic requests.
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

            if access_key_id is None:
                raise TypeError("Missing required property 'access_key_id'")
            __props__['access_key_id'] = access_key_id
            __props__['authorization'] = authorization
            __props__['description'] = description
            __props__['external_id'] = external_id
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['role_session_name'] = role_session_name
            __props__['role_to_assume'] = role_to_assume
            if secret_access_key is None:
                raise TypeError("Missing required property 'secret_access_key'")
            __props__['secret_access_key'] = secret_access_key
            if service_endpoint_name is None:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__['service_endpoint_name'] = service_endpoint_name
            __props__['session_token'] = session_token
            __props__['secret_access_key_hash'] = None
            __props__['session_token_hash'] = None
        super(ServiceEndpointAws, __self__).__init__(
            'azuredevops:index/serviceEndpointAws:ServiceEndpointAws',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key_id: Optional[pulumi.Input[str]] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            role_session_name: Optional[pulumi.Input[str]] = None,
            role_to_assume: Optional[pulumi.Input[str]] = None,
            secret_access_key: Optional[pulumi.Input[str]] = None,
            secret_access_key_hash: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            session_token: Optional[pulumi.Input[str]] = None,
            session_token_hash: Optional[pulumi.Input[str]] = None) -> 'ServiceEndpointAws':
        """
        Get an existing ServiceEndpointAws resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: The AWS access key ID for signing programmatic requests.
        :param pulumi.Input[str] external_id: A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
        :param pulumi.Input[str] project_id: The project ID or project name.
        :param pulumi.Input[str] role_session_name: Optional identifier for the assumed role session.
        :param pulumi.Input[str] role_to_assume: The Amazon Resource Name (ARN) of the role to assume.
        :param pulumi.Input[str] secret_access_key: The AWS secret access key for signing programmatic requests.
        :param pulumi.Input[str] secret_access_key_hash: A bcrypted hash of the attribute 'secret_access_key'
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] session_token: The AWS session token for signing programmatic requests.
        :param pulumi.Input[str] session_token_hash: A bcrypted hash of the attribute 'session_token'
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_key_id"] = access_key_id
        __props__["authorization"] = authorization
        __props__["description"] = description
        __props__["external_id"] = external_id
        __props__["project_id"] = project_id
        __props__["role_session_name"] = role_session_name
        __props__["role_to_assume"] = role_to_assume
        __props__["secret_access_key"] = secret_access_key
        __props__["secret_access_key_hash"] = secret_access_key_hash
        __props__["service_endpoint_name"] = service_endpoint_name
        __props__["session_token"] = session_token
        __props__["session_token_hash"] = session_token_hash
        return ServiceEndpointAws(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKeyId")
    def access_key_id(self) -> pulumi.Output[str]:
        """
        The AWS access key ID for signing programmatic requests.
        """
        return pulumi.get(self, "access_key_id")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[Optional[str]]:
        """
        A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project ID or project name.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="roleSessionName")
    def role_session_name(self) -> pulumi.Output[Optional[str]]:
        """
        Optional identifier for the assumed role session.
        """
        return pulumi.get(self, "role_session_name")

    @property
    @pulumi.getter(name="roleToAssume")
    def role_to_assume(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the role to assume.
        """
        return pulumi.get(self, "role_to_assume")

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> pulumi.Output[str]:
        """
        The AWS secret access key for signing programmatic requests.
        """
        return pulumi.get(self, "secret_access_key")

    @property
    @pulumi.getter(name="secretAccessKeyHash")
    def secret_access_key_hash(self) -> pulumi.Output[str]:
        """
        A bcrypted hash of the attribute 'secret_access_key'
        """
        return pulumi.get(self, "secret_access_key_hash")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter(name="sessionToken")
    def session_token(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS session token for signing programmatic requests.
        """
        return pulumi.get(self, "session_token")

    @property
    @pulumi.getter(name="sessionTokenHash")
    def session_token_hash(self) -> pulumi.Output[str]:
        """
        A bcrypted hash of the attribute 'session_token'
        """
        return pulumi.get(self, "session_token_hash")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

