+++
title = "AWS: IAM"
weight = 1
description="IAM - Idenity and Acces Management: IAM Identity, Dos & Don'ts, Federation Integration"
+++


## IAM

AWS Identity and Access Management (IAM) enables you to manage access to AWS services and resources securely. Using IAM, you can create and manage AWS users and groups, and use permissions to allow and deny their access to AWS resources.


### Root User

* Every account has a root user in AWS. A root user is something that's created automatically for you whenever you create an AWS account.
* Every single AWS account has a root user. 
* The trouble is that root users have unrestricted access to every service and resource that is in AWS inside of your account. 
* The permissions of root user can't be restricted in any way. 


### Dos and Don'ts

* You should not be accessing the root account on a regular basis, whether that's daily, weekly, or whatever. 
* Make sure that you turn on multi-factor authentication on the root account. Multi-factor authentication used to be called two-factor authentication. It really just means that we know the password and we have some sort of a token that we will get a number generated. It's something that you might even use your smartphone for. But it now means that I have to know the username and password and I have to have this token that's going to generate a code. We'll see more about how you'll do that later. 
* Make sure that you've disabled your root access keys. This isn't the interactive login for root, it has to do with how we can access the account programmatically. 
* Make sure that you rotate the credentials. Just because we say don't log in doesn't mean set the password and then forget it. 
* Don't share the root user credentials. password. And all that the audit logs show is that root logged in and did the job. Kind of dangerous. 
* Make sure that you create a user that has administrative privileges that's assigned to you and that you know the password only.


### Features & Functions

* Allows user to have very secure access including through the use of multi-factor authentication and federation. 
* Grant user a lot of granular control over the specific resources.
* Grant temporary access to different people. 
* Simplify the number of logins by using federating identities
* Integrate the IAM solution is with all of the different products that AWS offers. 


### MFA

* MFA stands for multi-factor authentication
* Extra layer of security.
* Prevent against imposters, somebody who just happened to guess the right password or happened to actually shoulder surf and watch somebody key in their username and password.


### IAM User

___It is an entity that you create in AWS. The IAM user represents the person or service who uses the IAM user to interact with AWS. A primary use for IAM users is to give people the ability to sign in to the AWS Management Console for interactive tasks and to make programmatic requests to AWS services using the API or CLI. A user in AWS consists of a name, a password to sign into the AWS Management Console, and up to two access keys that can be used with the API or CLI.___

* Receive its own unique identifier, aka, ARN, which is an acronym, of course, which stands for Amazon Resource Name. Now, each one of these user accounts will have its own set of unique credentials, which again might be a common-sense factor, but still everybody needs their own username and password.
* Access the appropriate resources by using policies. 
* very easy to remove.
* set up specific permissions. It's a good practice.

### IAM Policy



* A policy is just a JSON-formatted document. It details what you can do, and it includes information from four different categories. 
* We can attach the policy to IAM identities. And those identities could be IAM user accounts, group accounts, or IAM roles.


### Federation

#### IAM Role

___It is very similar to a user, in that it is an identity with permission policies that determine what the identity can and cannot do in AWS. However, a role does not have any credentials (password or access keys) associated with it.___

* A role is used to assisgn temporary permissions
* Polices are associated to IAM Roles
* Roles are assumed by people or applications

#### IAM Group

___It is a collection of IAM users. You can use groups to specify permissions for a collection of users, which can make those permissions easier to manage for those users.___


#### Cross Accounts

* Flexible management allows me to take multiple accounts

* It's all token based, so it really optimizes the whole process of setting up security. 

* Manage a lot of different services between multiple accounts, as we as between business units

* Steps:
    * There are two different AWS accounts. There is one user in one account and we have some resources in another. 
    * Create a role that exists within the account with the resources.
    * create a role, that permits access to the resources, can be then applied to the person or group over in the account by the virtue of what we call the sts:AssumeRole, aka, temporary security credential, aka, .

* Flow chart

{{<mermaid>}}
graph BT
    EC2(Elastice Computer Cloud)
    S3(Simple Storage Service)
    RDS(Rational Database Service)
    temporay_security_credential-->IAM_User
    IAM_Role --> temporay_security_credential([sts:AssumeRole])
    subgraph Accuont_2 
        IAM_Role(IAM Role with IAM Policies attached)
        IAM_Role --> EC2
        IAM_Role --> S3
        IAM_Role --> RDS
    end
   temporay_security_credential;
   subgraph Account_1     
        IAM_User(IAM User or Group ABC)
        IAM_User --> EC2
        IAM_User --> S3
        IAM_User --> RDS
    end 
{{</mermaid >}}




#### SAML

* SAML 2.0 is the one that you'll use for things like Active Directory. It stands for Security Assertion Markup Language. 

* It's an open standard. The idea is to create an open-standards-based approach that works with as many identity providers (idPs) as possible, things like Auth0, Microsoft Active Directory, and Shibboleth, etc. 

* Use existing corporate credentials for authentication and authorizaion in AWS console, CL, or API calls.

* Steps:

    * set up a SAML 2.0 -based federation,
    * Permit given identity to log in to the console, use CLI commands, or API calls. Now, this requires you to have an identity provider, such as Active Directory. 
    * when the user authenticates with Active Directory, Active Directory issues the SAML assertion document that says yep you are who you said you are and passes that over so that now when you go to access resources AWS says oh yeah
    * User is going to log in again to this identity provider. This checks against the local LDAP identity store. so again, Active Directory. 
    * Get confirmation, A token will be issue, that allows us to talk to Amazon's security token service, or STS, that allows us to assume a particular role within AWS. 
    * This passes back the temporary credentials. User is able to authenticate and gain access to acces AWS resources.
    * don't even have an AWS account.

#### AWS DS

* There are three different varieties. 
* Simple Active Directory for smaller, limited use cases.
* AWS Directory Service for Microsoft Active Directory, it's a mouthful. That's a fully-featured Active Directory. 
* Active Directory Connector that actually just facilitates or front ends the authentication request. And then the actual authentication is sent back across your VPN to your on-premises.




