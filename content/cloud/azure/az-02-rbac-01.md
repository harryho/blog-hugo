+++
title = "Azure: RBAC - 1"
weight = 1
description="Introduction of RBAC"

+++

## RBAC

Azure RBAC is an authorization system built on Azure Resource Manager that provides fine-grained access management of Azure resources.

### What to do with RBAC

* Allow one user to manage virtual machines in a subscription and another user to manage virtual networks
* Allow a DBA group to manage SQL databases in a subscription
* Allow a user to manage all resources in a resource group, such as virtual machines, websites, and subnets
* Allow an application to access all resources in a resource group

### How it works

The way you control access to resources using Azure RBAC is to create role assignments. This is a key concept to understand – it's how permissions are enforced. A role assignment consists of three elements: security principal, role definition, and scope.


#### Security principal

A security principal is an object that represents a user, group, service principal, or managed identity that is requesting access to Azure resources.

* User - An individual who has a profile in Azure Active Directory.

* Group - A set of users created in Azure Active Directory. 

* Service principal - A security identity used by applications or services to access specific Azure resources.

* Managed identity - An identity in Azure Active Directory that is automatically managed by Azure.


#### Role definition

A role definition is a collection of permissions. It's typically just called a role. A role definition lists the operations that can be performed, such as read, write, and delete.

Built-in roles:

* Owner - Has full access to all resources including the right to delegate access to others.
* Contributor - Can create and manage all types of Azure resources but can't grant access to others.
* Reader - Can view existing Azure resources.
* User Access Administrator - Lets you manage user access to Azure resources.


#### Scope

Scope is the set of resources that the access applies to. When you assign a role, you can further limit the actions allowed by defining a scope.


Scopes are structured in a parent-child relationship:management group, subscription, resource group, or resource.

{{<mermaid>}}
graph TB
    MG(Management Group)
    S1(Subscrition)
    S2(Subscrition)
    RG1(Resource Group)
    RG2(Resource Group)
    R1(Resource)
    R2(Resource)
    MG --- S1
    MG --- S2
    S2 --- RG1
    S2 --- RG2
    RG1 --- R1
    RG1 --- R2
{{</mermaid >}}


#### Role assignments

A role assignment is the process of attaching a role definition to a user, group, service principal, or managed identity at a particular scope for the purpose of granting access. Access is granted by creating a role assignment, and access is revoked by removing a role assignment.

{{<mermaid>}}
stateDiagram
    SP:Security Principal
    RA:Role_Assignment
    S:Scope
    RD:Role_Definition
    S --> RA
    SP --> RA
    RD --> RA

{{</mermaid >}}


#### Multiple role assignments

Azure RBAC is an additive model, so your effective permissions are the sum of your role assignments. The sum of the Contributor permissions and the Reader permissions is effectively the Contributor role for the resource group. Therefore, in this case, the Reader role assignment has no impact.

#### Deny assignments

Previously, Azure RBAC was an allow-only model with no deny, but now Azure RBAC supports deny assignments in a limited way. Similar to a role assignment, a deny assignment attaches a set of deny actions to a user, group, service principal, or managed identity at a particular scope for the purpose of denying access.

A role assignment defines a set of actions that are allowed, while a deny assignment defines a set of actions that are not allowed. In other words, deny assignments block users from performing specified actions even if a role assignment grants them access.




