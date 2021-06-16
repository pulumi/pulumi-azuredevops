CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to v0.1.5 of the AzureDevOps Terraform provider

---

## 2.1.0 (2021-05-27)
* Upgrade to v3.2.1 of the pulumi-terraform-bridge

## 2.0.0 (2021-04-19)
* Upgrade to v0.1.4 of the AzureDevOps Terraform provider
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 1.4.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 1.3.2 (2021-04-05)
* Upgrade to v0.1.3 of the AzureDevOps Terraform provider

## 1.3.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 1.3.0 (2021-03-15)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 1.2.2 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 1.2.1 (2021-02-10)
* Upgrade to v0.1.2 of the AzureDevOps Terraform provider

## 1.2.0 (2021-02-01)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 1.1.1 (2021-01-12)
* Upgrade to v0.1.1 of the AzureDevOps Terraform provider

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
