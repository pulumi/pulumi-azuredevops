# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = [
    'GetIdentityGroupResult',
    'AwaitableGetIdentityGroupResult',
    'get_identity_group',
    'get_identity_group_output',
]

@pulumi.output_type
class GetIdentityGroupResult:
    """
    A collection of values returned by getIdentityGroup.
    """
    def __init__(__self__, descriptor=None, id=None, name=None, project_id=None):
        if descriptor and not isinstance(descriptor, str):
            raise TypeError("Expected argument 'descriptor' to be a str")
        pulumi.set(__self__, "descriptor", descriptor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def descriptor(self) -> str:
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
        """
        This is the non-unique display name of the identity subject. To change this field, you must alter its value in the source provider.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")


class AwaitableGetIdentityGroupResult(GetIdentityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdentityGroupResult(
            descriptor=self.descriptor,
            id=self.id,
            name=self.name,
            project_id=self.project_id)


def get_identity_group(name: Optional[str] = None,
                       project_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdentityGroupResult:
    """
    Use this data source to access information about an existing Group within Azure DevOps On-Premise(Azure DevOps Server).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    # load existing group with specific name
    example_project_group = azuredevops.get_identity_group(project_id=example["id"],
        name="[Project-Name]\\\\Group-Name")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)


    :param str name: The name of the group.
    :param str project_id: The Project ID.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuredevops:index/getIdentityGroup:getIdentityGroup', __args__, opts=opts, typ=GetIdentityGroupResult).value

    return AwaitableGetIdentityGroupResult(
        descriptor=pulumi.get(__ret__, 'descriptor'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'))


@_utilities.lift_output_func(get_identity_group)
def get_identity_group_output(name: Optional[pulumi.Input[str]] = None,
                              project_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdentityGroupResult]:
    """
    Use this data source to access information about an existing Group within Azure DevOps On-Premise(Azure DevOps Server).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    # load existing group with specific name
    example_project_group = azuredevops.get_identity_group(project_id=example["id"],
        name="[Project-Name]\\\\Group-Name")
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 7.0 - Identities](https://docs.microsoft.com/en-us/rest/api/azure/devops/ims/?view=azure-devops-rest-7.2)


    :param str name: The name of the group.
    :param str project_id: The Project ID.
    """
    ...
