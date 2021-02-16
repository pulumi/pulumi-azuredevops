# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'org_service_url',
    'personal_access_token',
]

__config__ = pulumi.Config('azuredevops')

org_service_url = __config__.get('orgServiceUrl') or _utilities.get_env('AZDO_ORG_SERVICE_URL')
"""
The url of the Azure DevOps instance which should be used.
"""

personal_access_token = __config__.get('personalAccessToken')
"""
The personal access token which should be used.
"""

