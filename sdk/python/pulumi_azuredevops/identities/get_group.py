# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, descriptor=None, id=None, name=None, project_id=None):
        if descriptor and not isinstance(descriptor, str):
            raise TypeError("Expected argument 'descriptor' to be a str")
        __self__.descriptor = descriptor
        """
        The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            descriptor=self.descriptor,
            id=self.id,
            name=self.name,
            project_id=self.project_id)

def get_group(name=None,project_id=None,opts=None):
    """
    Use this data source to access information about an existing Group within Azure DevOps

    ## Example Usage



    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    project = azuredevops.Core.get_project(project_name="contoso-project")
    test = azuredevops.Identities.get_group(project_id=project.id,
        name="Test Group")
    pulumi.export("groupId", test.id)
    pulumi.export("groupDescriptor", test.descriptor)
    ```

    ## Relevant Links

    * [Azure DevOps Service REST API 5.1 - Groups - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/groups/get?view=azure-devops-rest-5.1)


    :param str name: The Group Name.
    :param str project_id: The Project Id.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:Identities/getGroup:getGroup', __args__, opts=opts).value

    return AwaitableGetGroupResult(
        descriptor=__ret__.get('descriptor'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        project_id=__ret__.get('projectId'))
