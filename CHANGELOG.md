CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to v0.1.1 of the AzureDevOps Terraform provider

---

## 1.1.0 (2020-12-23)
* Upgrade to v0.1.0 of the AzureDevOps Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.16.0
* Upgrade to Pulumi v2.16.0 
  **Please Note:**  
  The NodeJS and Go SDKs have lowercased all provider modules.
  `azuredevops.Agent` will now be `azuredevops.agent` in Go, NodeJS and Python SDKs.
  `azuredevops.Build` will now be `azuredevops.build` in Go, NodeJS and Python SDKs.
  `azuredevops.Core` will now be `azuredevops.core` in Go, NodeJS and Python SDKs.
  `azuredevops.Entitlement` will now be `azuredevops.entitlement` in Go, NodeJS and Python SDKs.
  `azuredevops.Identities` will now be `azuredevops.identities` in Go, NodeJS and Python SDKs.
  `azuredevops.Pipeline` will now be `azuredevops.pipeline` in Go, NodeJS and Python SDKs.
  `azuredevops.Policy` will now be `azuredevops.policy` in Go, NodeJS and Python SDKs.
  `azuredevops.Repository` will now be `azuredevops.repository` in Go, NodeJS and Python SDKs.
  `azuredevops.Security` will now be `azuredevops.security` in Go, NodeJS and Python SDKs.
  `azuredevops.ServiceEndpoint` will now be `azuredevops.serviceendpoint` in Go, NodeJS and Python SDKs.
* Deprecate the use of the following packages in favour of `azuredevops` main package. These packages
  will be removed in a future version:  
  ** `azuredevops.agent`  
  ** `azuredevops.build`  
  ** `azuredevops.core`  
  ** `azuredevops.entitlement`  
  ** `azuredevops.identities`  
  ** `azuredevops.pipeline`  
  ** `azuredevops.policy`  
  ** `azuredevops.repository`  
  ** `azuredevops.security`  
  ** `azuredevops.serviceendpoint`

## 1.0.0 (2020-06-20)
* Initial release of the provider against v0.0.1 of the AzureDevOps Terraform Provider
