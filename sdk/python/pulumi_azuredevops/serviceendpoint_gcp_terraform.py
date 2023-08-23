# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceendpointGcpTerraformArgs', 'ServiceendpointGcpTerraform']

@pulumi.input_type
class ServiceendpointGcpTerraformArgs:
    def __init__(__self__, *,
                 gcp_project_id: pulumi.Input[str],
                 private_key: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 service_endpoint_name: pulumi.Input[str],
                 token_uri: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 client_email: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceendpointGcpTerraform resource.
        :param pulumi.Input[str] gcp_project_id: GCP project associated with the Service Connection.
        :param pulumi.Input[str] private_key: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token_uri: The token uri field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] client_email: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] scope: Scope to be provided.
        """
        pulumi.set(__self__, "gcp_project_id", gcp_project_id)
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        pulumi.set(__self__, "token_uri", token_uri)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if client_email is not None:
            pulumi.set(__self__, "client_email", client_email)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="gcpProjectId")
    def gcp_project_id(self) -> pulumi.Input[str]:
        """
        GCP project associated with the Service Connection.
        """
        return pulumi.get(self, "gcp_project_id")

    @gcp_project_id.setter
    def gcp_project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gcp_project_id", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        """
        The client email field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)

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
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Input[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter(name="tokenUri")
    def token_uri(self) -> pulumi.Input[str]:
        """
        The token uri field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "token_uri")

    @token_uri.setter
    def token_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "token_uri", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="clientEmail")
    def client_email(self) -> Optional[pulumi.Input[str]]:
        """
        The client email field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "client_email")

    @client_email.setter
    def client_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_email", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Scope to be provided.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class _ServiceendpointGcpTerraformState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 client_email: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gcp_project_id: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 token_uri: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceendpointGcpTerraform resources.
        :param pulumi.Input[str] client_email: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] gcp_project_id: GCP project associated with the Service Connection.
        :param pulumi.Input[str] private_key: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] scope: Scope to be provided.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token_uri: The token uri field in the JSON key file for creating the JSON Web Token.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if client_email is not None:
            pulumi.set(__self__, "client_email", client_email)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if gcp_project_id is not None:
            pulumi.set(__self__, "gcp_project_id", gcp_project_id)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if service_endpoint_name is not None:
            pulumi.set(__self__, "service_endpoint_name", service_endpoint_name)
        if token_uri is not None:
            pulumi.set(__self__, "token_uri", token_uri)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="clientEmail")
    def client_email(self) -> Optional[pulumi.Input[str]]:
        """
        The client email field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "client_email")

    @client_email.setter
    def client_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_email", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="gcpProjectId")
    def gcp_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        GCP project associated with the Service Connection.
        """
        return pulumi.get(self, "gcp_project_id")

    @gcp_project_id.setter
    def gcp_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcp_project_id", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The client email field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

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
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Scope to be provided.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @service_endpoint_name.setter
    def service_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_endpoint_name", value)

    @property
    @pulumi.getter(name="tokenUri")
    def token_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The token uri field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "token_uri")

    @token_uri.setter
    def token_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_uri", value)


class ServiceendpointGcpTerraform(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 client_email: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gcp_project_id: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 token_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_gcp_terraform = azuredevops.ServiceendpointGcpTerraform("exampleServiceendpointGcpTerraform",
            project_id=example_project.id,
            token_uri="https://oauth2.example.com/token",
            client_email="gcp-sa-example@example.iam.gserviceaccount.com",
            private_key="0000000000000000000000000000000000000",
            service_endpoint_name="Example GCP Terraform extension",
            gcp_project_id="Example GCP Project",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.1)

        ## Import

        Azure DevOps Service Endpoint GCP can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform azuredevops_serviceendpoint_gcp_terraform.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_email: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] gcp_project_id: GCP project associated with the Service Connection.
        :param pulumi.Input[str] private_key: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] scope: Scope to be provided.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token_uri: The token uri field in the JSON key file for creating the JSON Web Token.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceendpointGcpTerraformArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example_project = azuredevops.Project("exampleProject",
            visibility="private",
            version_control="Git",
            work_item_template="Agile",
            description="Managed by Terraform")
        example_serviceendpoint_gcp_terraform = azuredevops.ServiceendpointGcpTerraform("exampleServiceendpointGcpTerraform",
            project_id=example_project.id,
            token_uri="https://oauth2.example.com/token",
            client_email="gcp-sa-example@example.iam.gserviceaccount.com",
            private_key="0000000000000000000000000000000000000",
            service_endpoint_name="Example GCP Terraform extension",
            gcp_project_id="Example GCP Project",
            description="Managed by Terraform")
        ```
        ## Relevant Links

        - [Azure DevOps Service REST API 7.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.1)

        ## Import

        Azure DevOps Service Endpoint GCP can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

        ```sh
         $ pulumi import azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform azuredevops_serviceendpoint_gcp_terraform.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ServiceendpointGcpTerraformArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceendpointGcpTerraformArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 client_email: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gcp_project_id: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 service_endpoint_name: Optional[pulumi.Input[str]] = None,
                 token_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceendpointGcpTerraformArgs.__new__(ServiceendpointGcpTerraformArgs)

            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["client_email"] = client_email
            __props__.__dict__["description"] = description
            if gcp_project_id is None and not opts.urn:
                raise TypeError("Missing required property 'gcp_project_id'")
            __props__.__dict__["gcp_project_id"] = gcp_project_id
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["scope"] = scope
            if service_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_name'")
            __props__.__dict__["service_endpoint_name"] = service_endpoint_name
            if token_uri is None and not opts.urn:
                raise TypeError("Missing required property 'token_uri'")
            __props__.__dict__["token_uri"] = token_uri
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServiceendpointGcpTerraform, __self__).__init__(
            'azuredevops:index/serviceendpointGcpTerraform:ServiceendpointGcpTerraform',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            client_email: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            gcp_project_id: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            service_endpoint_name: Optional[pulumi.Input[str]] = None,
            token_uri: Optional[pulumi.Input[str]] = None) -> 'ServiceendpointGcpTerraform':
        """
        Get an existing ServiceendpointGcpTerraform resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_email: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] gcp_project_id: GCP project associated with the Service Connection.
        :param pulumi.Input[str] private_key: The client email field in the JSON key file for creating the JSON Web Token.
        :param pulumi.Input[str] project_id: The ID of the project.
        :param pulumi.Input[str] scope: Scope to be provided.
        :param pulumi.Input[str] service_endpoint_name: The Service Endpoint name.
        :param pulumi.Input[str] token_uri: The token uri field in the JSON key file for creating the JSON Web Token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceendpointGcpTerraformState.__new__(_ServiceendpointGcpTerraformState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["client_email"] = client_email
        __props__.__dict__["description"] = description
        __props__.__dict__["gcp_project_id"] = gcp_project_id
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["scope"] = scope
        __props__.__dict__["service_endpoint_name"] = service_endpoint_name
        __props__.__dict__["token_uri"] = token_uri
        return ServiceendpointGcpTerraform(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="clientEmail")
    def client_email(self) -> pulumi.Output[Optional[str]]:
        """
        The client email field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "client_email")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="gcpProjectId")
    def gcp_project_id(self) -> pulumi.Output[str]:
        """
        GCP project associated with the Service Connection.
        """
        return pulumi.get(self, "gcp_project_id")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The client email field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        Scope to be provided.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="serviceEndpointName")
    def service_endpoint_name(self) -> pulumi.Output[str]:
        """
        The Service Endpoint name.
        """
        return pulumi.get(self, "service_endpoint_name")

    @property
    @pulumi.getter(name="tokenUri")
    def token_uri(self) -> pulumi.Output[str]:
        """
        The token uri field in the JSON key file for creating the JSON Web Token.
        """
        return pulumi.get(self, "token_uri")
