# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

warnings.warn("""azuredevops.identities.getGroup has been deprecated in favor of azuredevops.getGroup""", DeprecationWarning)

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, descriptor=None, id=None, name=None, origin=None, origin_id=None, project_id=None):
        if descriptor and not isinstance(descriptor, str):
            raise TypeError("Expected argument 'descriptor' to be a str")
        pulumi.set(__self__, "descriptor", descriptor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if origin and not isinstance(origin, str):
            raise TypeError("Expected argument 'origin' to be a str")
        pulumi.set(__self__, "origin", origin)
        if origin_id and not isinstance(origin_id, str):
            raise TypeError("Expected argument 'origin_id' to be a str")
        pulumi.set(__self__, "origin_id", origin_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def descriptor(self) -> str:
        """
        The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations.
        """
        return pulumi.get(self, "descriptor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def origin(self) -> str:
        """
        The type of source provider for the origin identifier (ex:AD, AAD, MSA)
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="originId")
    def origin_id(self) -> str:
        """
        The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
        """
        return pulumi.get(self, "origin_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            descriptor=self.descriptor,
            id=self.id,
            name=self.name,
            origin=self.origin,
            origin_id=self.origin_id,
            project_id=self.project_id)


def get_group(name: Optional[str] = None,
              project_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Use this data source to access information about an existing Group within Azure DevOps

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.get_project(name="Example Project")
    example_group = azuredevops.get_group(project_id=example_project.id,
        name="Example Group")
    pulumi.export("groupId", example_group.id)
    pulumi.export("groupDescriptor", example_group.descriptor)
    example_collection_group = azuredevops.get_group(name="Project Collection Administrators")
    pulumi.export("collectionGroupId", example_group.id)
    pulumi.export("collectionGroupDescriptor", example_group.descriptor)
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-7.0)


    :param str name: The Group Name.
    :param str project_id: The Project ID. If no project ID is specified the project collection groups will be searched.
    """
    pulumi.log.warn("""get_group is deprecated: azuredevops.identities.getGroup has been deprecated in favor of azuredevops.getGroup""")
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:Identities/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        descriptor=pulumi.get(__ret__, 'descriptor'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        origin=pulumi.get(__ret__, 'origin'),
        origin_id=pulumi.get(__ret__, 'origin_id'),
        project_id=pulumi.get(__ret__, 'project_id'))


@_utilities.lift_output_func(get_group)
def get_group_output(name: Optional[pulumi.Input[str]] = None,
                     project_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupResult]:
    """
    Use this data source to access information about an existing Group within Azure DevOps

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    example_project = azuredevops.get_project(name="Example Project")
    example_group = azuredevops.get_group(project_id=example_project.id,
        name="Example Group")
    pulumi.export("groupId", example_group.id)
    pulumi.export("groupDescriptor", example_group.descriptor)
    example_collection_group = azuredevops.get_group(name="Project Collection Administrators")
    pulumi.export("collectionGroupId", example_group.id)
    pulumi.export("collectionGroupDescriptor", example_group.descriptor)
    ```
    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-7.0)


    :param str name: The Group Name.
    :param str project_id: The Project ID. If no project ID is specified the project collection groups will be searched.
    """
    pulumi.log.warn("""get_group is deprecated: azuredevops.identities.getGroup has been deprecated in favor of azuredevops.getGroup""")
    ...
