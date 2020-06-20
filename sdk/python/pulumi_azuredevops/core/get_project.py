# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, description=None, features=None, id=None, process_template_id=None, project_name=None, version_control=None, visibility=None, work_item_template=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if features and not isinstance(features, dict):
            raise TypeError("Expected argument 'features' to be a dict")
        __self__.features = features
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if process_template_id and not isinstance(process_template_id, str):
            raise TypeError("Expected argument 'process_template_id' to be a str")
        __self__.process_template_id = process_template_id
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        __self__.project_name = project_name
        if version_control and not isinstance(version_control, str):
            raise TypeError("Expected argument 'version_control' to be a str")
        __self__.version_control = version_control
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        __self__.visibility = visibility
        if work_item_template and not isinstance(work_item_template, str):
            raise TypeError("Expected argument 'work_item_template' to be a str")
        __self__.work_item_template = work_item_template
class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            description=self.description,
            features=self.features,
            id=self.id,
            process_template_id=self.process_template_id,
            project_name=self.project_name,
            version_control=self.version_control,
            visibility=self.visibility,
            work_item_template=self.work_item_template)

def get_project(project_name=None,opts=None):
    """
    Use this data source to access information about an existing Project within Azure DevOps.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_azuredevops as azuredevops

    project = azuredevops.Core.get_project(project_name="Sample Project")
    pulumi.export("id", project.id)
    pulumi.export("projectName", project.project_name)
    pulumi.export("visibility", project.visibility)
    pulumi.export("versionControl", project.version_control)
    pulumi.export("workItemTemplate", project.work_item_template)
    pulumi.export("processTemplateId", project.process_template_id)
    ```

    ## Relevant Links

    - [Azure DevOps Service REST API 5.1 - Projects - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/get?view=azure-devops-rest-5.1)
    """
    __args__ = dict()


    __args__['projectName'] = project_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuredevops:Core/getProject:getProject', __args__, opts=opts).value

    return AwaitableGetProjectResult(
        description=__ret__.get('description'),
        features=__ret__.get('features'),
        id=__ret__.get('id'),
        process_template_id=__ret__.get('processTemplateId'),
        project_name=__ret__.get('projectName'),
        version_control=__ret__.get('versionControl'),
        visibility=__ret__.get('visibility'),
        work_item_template=__ret__.get('workItemTemplate'))
