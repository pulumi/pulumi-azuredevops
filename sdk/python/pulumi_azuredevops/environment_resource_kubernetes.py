# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['EnvironmentResourceKubernetesArgs', 'EnvironmentResourceKubernetes']

@pulumi.input_type
class EnvironmentResourceKubernetesArgs:
    def __init__(__self__, *,
                 environment_id: pulumi.Input[builtins.int],
                 namespace: pulumi.Input[builtins.str],
                 project_id: pulumi.Input[builtins.str],
                 service_endpoint_id: pulumi.Input[builtins.str],
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a EnvironmentResourceKubernetes resource.
        :param pulumi.Input[builtins.int] environment_id: The ID of the environment under which to create the Kubernetes Resource.
        :param pulumi.Input[builtins.str] namespace: The namespace for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] project_id: The ID of the project.
        :param pulumi.Input[builtins.str] service_endpoint_id: The ID of the service endpoint to associate with the Kubernetes Resource.
        :param pulumi.Input[builtins.str] cluster_name: A cluster name for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] name: The name for the Kubernetes Resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A set of tags for the Kubernetes Resource.
        """
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "service_endpoint_id", service_endpoint_id)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[builtins.int]:
        """
        The ID of the environment under which to create the Kubernetes Resource.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[builtins.str]:
        """
        The namespace for the Kubernetes Resource.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the service endpoint to associate with the Kubernetes Resource.
        """
        return pulumi.get(self, "service_endpoint_id")

    @service_endpoint_id.setter
    def service_endpoint_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_endpoint_id", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A cluster name for the Kubernetes Resource.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for the Kubernetes Resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A set of tags for the Kubernetes Resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EnvironmentResourceKubernetesState:
    def __init__(__self__, *,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_endpoint_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering EnvironmentResourceKubernetes resources.
        :param pulumi.Input[builtins.str] cluster_name: A cluster name for the Kubernetes Resource.
        :param pulumi.Input[builtins.int] environment_id: The ID of the environment under which to create the Kubernetes Resource.
        :param pulumi.Input[builtins.str] name: The name for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] namespace: The namespace for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] project_id: The ID of the project.
        :param pulumi.Input[builtins.str] service_endpoint_id: The ID of the service endpoint to associate with the Kubernetes Resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A set of tags for the Kubernetes Resource.
        """
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if service_endpoint_id is not None:
            pulumi.set(__self__, "service_endpoint_id", service_endpoint_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A cluster name for the Kubernetes Resource.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The ID of the environment under which to create the Kubernetes Resource.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for the Kubernetes Resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace for the Kubernetes Resource.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the service endpoint to associate with the Kubernetes Resource.
        """
        return pulumi.get(self, "service_endpoint_id")

    @service_endpoint_id.setter
    def service_endpoint_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_endpoint_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A set of tags for the Kubernetes Resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


class EnvironmentResourceKubernetes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_endpoint_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Manages a Kubernetes Resource for an Environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Pulumi")
        example_environment = azuredevops.Environment("example",
            project_id=example.id,
            name="Example Environment")
        example_service_endpoint_kubernetes = azuredevops.ServiceEndpointKubernetes("example",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="AzureSubscription",
            azure_subscriptions=[{
                "subscription_id": "00000000-0000-0000-0000-000000000000",
                "subscription_name": "Example",
                "tenant_id": "00000000-0000-0000-0000-000000000000",
                "resourcegroup_id": "example-rg",
                "namespace": "default",
                "cluster_name": "example-aks",
            }])
        example_environment_resource_kubernetes = azuredevops.EnvironmentResourceKubernetes("example",
            project_id=example.id,
            environment_id=example_environment.id,
            service_endpoint_id=example_service_endpoint_kubernetes.id,
            name="Example",
            namespace="default",
            cluster_name="example-aks",
            tags=[
                "tag1",
                "tag2",
            ])
        ```

        ## Relevant Links

        * [Azure DevOps Service REST API 6.0 - Kubernetes](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/kubernetes?view=azure-devops-rest-6.0)

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_name: A cluster name for the Kubernetes Resource.
        :param pulumi.Input[builtins.int] environment_id: The ID of the environment under which to create the Kubernetes Resource.
        :param pulumi.Input[builtins.str] name: The name for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] namespace: The namespace for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] project_id: The ID of the project.
        :param pulumi.Input[builtins.str] service_endpoint_id: The ID of the service endpoint to associate with the Kubernetes Resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A set of tags for the Kubernetes Resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentResourceKubernetesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kubernetes Resource for an Environment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuredevops as azuredevops

        example = azuredevops.Project("example",
            name="Example Project",
            work_item_template="Agile",
            version_control="Git",
            visibility="private",
            description="Managed by Pulumi")
        example_environment = azuredevops.Environment("example",
            project_id=example.id,
            name="Example Environment")
        example_service_endpoint_kubernetes = azuredevops.ServiceEndpointKubernetes("example",
            project_id=example.id,
            service_endpoint_name="Example Kubernetes",
            apiserver_url="https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io",
            authorization_type="AzureSubscription",
            azure_subscriptions=[{
                "subscription_id": "00000000-0000-0000-0000-000000000000",
                "subscription_name": "Example",
                "tenant_id": "00000000-0000-0000-0000-000000000000",
                "resourcegroup_id": "example-rg",
                "namespace": "default",
                "cluster_name": "example-aks",
            }])
        example_environment_resource_kubernetes = azuredevops.EnvironmentResourceKubernetes("example",
            project_id=example.id,
            environment_id=example_environment.id,
            service_endpoint_id=example_service_endpoint_kubernetes.id,
            name="Example",
            namespace="default",
            cluster_name="example-aks",
            tags=[
                "tag1",
                "tag2",
            ])
        ```

        ## Relevant Links

        * [Azure DevOps Service REST API 6.0 - Kubernetes](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/kubernetes?view=azure-devops-rest-6.0)

        ## Import

        The resource does not support import.

        :param str resource_name: The name of the resource.
        :param EnvironmentResourceKubernetesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentResourceKubernetesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_endpoint_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentResourceKubernetesArgs.__new__(EnvironmentResourceKubernetesArgs)

            __props__.__dict__["cluster_name"] = cluster_name
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["name"] = name
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if service_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_endpoint_id'")
            __props__.__dict__["service_endpoint_id"] = service_endpoint_id
            __props__.__dict__["tags"] = tags
        super(EnvironmentResourceKubernetes, __self__).__init__(
            'azuredevops:index/environmentResourceKubernetes:EnvironmentResourceKubernetes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_name: Optional[pulumi.Input[builtins.str]] = None,
            environment_id: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            namespace: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            service_endpoint_id: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'EnvironmentResourceKubernetes':
        """
        Get an existing EnvironmentResourceKubernetes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_name: A cluster name for the Kubernetes Resource.
        :param pulumi.Input[builtins.int] environment_id: The ID of the environment under which to create the Kubernetes Resource.
        :param pulumi.Input[builtins.str] name: The name for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] namespace: The namespace for the Kubernetes Resource.
        :param pulumi.Input[builtins.str] project_id: The ID of the project.
        :param pulumi.Input[builtins.str] service_endpoint_id: The ID of the service endpoint to associate with the Kubernetes Resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A set of tags for the Kubernetes Resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentResourceKubernetesState.__new__(_EnvironmentResourceKubernetesState)

        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["environment_id"] = environment_id
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["service_endpoint_id"] = service_endpoint_id
        __props__.__dict__["tags"] = tags
        return EnvironmentResourceKubernetes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A cluster name for the Kubernetes Resource.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[builtins.int]:
        """
        The ID of the environment under which to create the Kubernetes Resource.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name for the Kubernetes Resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[builtins.str]:
        """
        The namespace for the Kubernetes Resource.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serviceEndpointId")
    def service_endpoint_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the service endpoint to associate with the Kubernetes Resource.
        """
        return pulumi.get(self, "service_endpoint_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A set of tags for the Kubernetes Resource.
        """
        return pulumi.get(self, "tags")

