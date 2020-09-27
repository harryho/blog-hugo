+++
title = "Azure: RBAC - 2"
weight = 1
description="Difference - Classic subscription, Azure Roles & Azure AD Roles"

+++

## Roles

- Classic subscription administrator roles
- Azure roles
- Azure Active Directory (Azure AD) roles

### History

> When Azure was initially released, access to resources was managed with just three administrator roles: Account Administrator, Service Administrator, and Co-Administrator. Later, Azure role-based access control (Azure RBAC) was added. Azure RBAC is a newer authorization system that provides fine-grained access management to Azure resources. Azure RBAC includes many built-in roles, can be assigned at different scopes, and allows you to create your own custom roles. To manage resources in Azure AD, such as users, groups, and domains, there are several Azure AD roles. To manage resources in Azure AD, such as users, groups, and domains, there are several Azure AD roles.

### High-level view 

{{<mermaid>}}
graph TB
    Root(Global Admin,User Access Admin-elevated access)
    AADT(Azure Active Directtory Tenent)
    RMG(Root Management Group)
    MG(Management Group)
    Sub(Subscriptions, Azure Account, Account Admin)
    RG(Resource Group)
    R(Resource)
    AADT --> Root
    Root --> RMG
    GAAA([Global_Admin, Application_Admin])
    OCRU([Owner,Contributor, Reader,User_Acess_Admin])
    subgraph Azure_AD_Roles
        GAAA
        AADT
    end
    Root;
    subgraph Azure_Roles
        OCRU
        RMG --> MG
        MG --> Sub
        subgraph Classic_Subscription_Admin_Roles
            Sub --> RG
            RG --> R
        end
    end
{{</mermaid >}}


### Classic subscription administrator roles

Account Administrator, Service Administrator, and Co-Administrator are the three classic subscription administrator roles in Azure. Classic subscription administrators have full access to the Azure subscription. 

Classic subscription administrators have full access to the Azure subscription. 

#### Account Administrator
* per Azure account	
* Manage billing in the Azure portal 
* Manage all subscriptions in an account
* Create new subscriptions
* Cancel subscriptions
* Change the billing for a subscription
* Change the Service Administrator
* Conceptually, the billing owner of the subscription.
* The Account Administrator has no access to the Azure portal.

#### Service Administrator	
* 1 per Azure subscription	
* Manage services in the Azure portal
* Cancel the subscription
* Assign users to the Co-Administrator role
* By default, for a new subscription, the Account Administrator is also the Service Administrator.
* The Service Administrator has the equivalent access of a user who is assigned the Owner role at the subscription scope.
* The Service Administrator has full access to the Azure portal.

#### Co-Administrator	
* 200 per subscription	
* Same access privileges as the Service Administrator, but can’t change the association of subscriptions to Azure directories
* Assign users to the Co-Administrator role, but cannot change the Service Administrator
* The Co-Administrator has the equivalent access of a user who is assigned the Owner role at the subscription scope.

### Azure account and Azure subscriptions

An Azure account represents a billing relationship. An Azure account is a user identity, one or more Azure subscriptions, and an associated set of Azure resources. The person who creates the account is the Account Administrator for all subscriptions created in that account. That person is also the default Service Administrator for the subscription.

Azure subscriptions help you organize access to Azure resources. They also help you control how resource usage is reported, billed, and paid for.

Each subscription can have a different billing and payment setup, so you can have different subscriptions and different plans by office, department, project, and so on. 

### Azure roles

Azure RBAC is an authorization system built on Azure Resource Manager that provides fine-grained access management to Azure resources.


#### Owner
* Full access to all resources
* Delegate access to others
* The Service Administrator and Co-Administrators are assigned the Owner role at the subscription scope
* Applies to all resource types.
  
#### Contributor
* Create and manage all of types of Azure resources
* Create a new tenant in Azure Active Directory
* Cannot grant access to others
* Applies to all resource types.
  
#### Reader	
* View Azure resources
* Applies to all resource types.

#### User Access Administrator	
* Manage user access to Azure resources


### Azure AD roles

Azure AD roles are used to manage Azure AD resources in a directory such as create or edit users, assign administrative roles to others, reset user passwords, manage user licenses, and manage domains.

#### Global Administrator	
* Manage access to all administrative features in Azure Active Directory, as well as services that federate to Azure Active Directory
* Assign administrator roles to others
* Reset the password for any user and all other administrators
* The person who signs up for the Azure Active Directory tenant becomes a Global Administrator.

#### User Administrator	
* Create and manage all aspects of users and groups
* Manage support tickets
* Monitor service health
* Change passwords for users, Helpdesk administrators, and other User Administrators
  
#### Billing Administrator	
* Make purchases
* Manage subscriptions
* Manage support tickets
* Monitors service health

### Azure roles VS Azure AD roles


Azure roles	| Azure AD roles
-----|------
Manage access to Azure resources	| Manage access to Azure Active Directory resources
Supports custom roles	|Supports custom roles
Scope can be specified at multiple levels (management group, subscription, resource group, resource)	| Scope is at the tenant level
Role information can be accessed in Azure portal, Azure CLI, Azure PowerShell, Azure Resource Manager templates, REST API	| Role information can be accessed in Azure admin portal, Microsoft 365 admin center, Microsoft Graph, AzureAD PowerShell


### Overlap

By default, Azure roles and Azure AD roles do not span Azure and Azure AD. However, if a Global Administrator elevates their access by choosing the Access management for Azure resources switch in the Azure portal, the Global Administrator will be granted the User Access Administrator role (an Azure role) on all subscriptions for a particular tenant.


