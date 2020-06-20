# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, id=None, project_name=None, projects=None, state=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        __self__.project_name = project_name
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        __self__.projects = projects
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            id=self.id,
            project_name=self.project_name,
            projects=self.projects,
            state=self.state)

def get_projects(project_name=None,state=None,opts=None):
    """
    Use this data source to access information about existing Projects within Azure DevOps.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    test = azuredevops.Core.get_projects(project_name="contoso",
        state="wellFormed")
    pulumi.export("projectId", [__item["project_id"] for __item in test.projects])
    pulumi.export("projectName", [__item["name"] for __item in test.projects])
    pulumi.export("projectUrl", [__item["projectUrl"] for __item in test.projects])
    pulumi.export("state", [__item["state"] for __item in test.projects])
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.1)
    """
    __args__ = dict()


    __args__['projectName'] = project_name
    __args__['state'] = state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:Core/getProjects:getProjects', __args__, opts=opts).value

    return AwaitableGetProjectsResult(
        id=__ret__.get('id'),
        project_name=__ret__.get('projectName'),
        projects=__ret__.get('projects'),
        state=__ret__.get('state'))
