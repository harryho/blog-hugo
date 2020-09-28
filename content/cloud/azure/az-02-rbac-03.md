+++
title = "Azure: RBAC - 3"
weight = 1
description="Best practices & Azure AD Privileged Identity Management"

+++

## Best practices 

### Only grant the access users need

Scope | Roles   | | | | |
------|--------|-------------------|--------|-------------|----|
-     | Reader | Resource-specific | Custom | Contributor | Owner|
Management Group | Observers | Users managing resources| Users managing resources| Users managing resources | Admins
Subscriptions | Observers | Users managing resources |  Users managing resources | Users managing resources| Admins
Resource Group | Observers | Users managing resources |  Users managing resources | Users managing resources |  Admins
Resources | Automated processes | Automated processes | Automated processes | Automated processes | Automated processes | 


### Azure AD Privileged Identity Management

Privileged Identity Management (PIM) is a service in Azure Active Directory (Azure AD) that enables you to manage, control, and monitor access to important resources in your organization. These resources include resources in Azure AD, Azure, and other Microsoft Online Services such as Microsoft 365 or Microsoft Intune.


#### What does it do

Privileged Identity Management provides time-based and approval-based role activation to mitigate the risks of excessive, unnecessary, or misused access permissions on resources that you care about.

- Provide just-in-time privileged access to Azure AD and Azure resources
- Assign time-bound access to resources using start and end dates
- Require approval to activate privileged roles
- Enforce multi-factor authentication to activate any role
- Use justification to understand why users activate
- Get notifications when privileged roles are activated
- Conduct access reviews to ensure users still need roles
- Download audit history for internal or external audit

