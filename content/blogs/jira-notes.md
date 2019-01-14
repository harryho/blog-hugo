+++
date = "2017-01-11T11:59:31+11:00"
title = "JIRA Practices"
description="JIRA is a proprietary issue tracking product, developed by Atlassian. Today I recap the bullet points about how I work with client within JIRA "
draft="true"
+++

## What is JIRA

> Jira is a proprietary issue tracking product, developed by Atlassian. It provides bug tracking, issue tracking, and project management functions. Although normally styled JIRA, the product name is not an acronym, but a truncation of Gojira, the Japanese name for Godzilla,[6] itself a reference to Jira's main competitor, Bugzilla. It has been developed since 2002.


## How to cooperate with client via JIRA

> The instruction below is only for the client who has independent JIRA instance. Here I am not going to discuss how to cooperate with multiple clients and projects on the same JIRA instance. IMO, the strategy for that would be case by case. 

### Prerequisites

* Your client's business name is __ABC__. You help them to manage their JIRA. You are the administrator of JIRA instance. 

* You have created a project for client, named **ABC IT Project**

* You want to allow the client to access this project on JIRA, but you want to avoid client's unintented update to mess up the project management. 

* You want to cooperate with other project team which build some product communicating with your system

* You want to integrate with CI / CD tools


### Requirement of permission control

* Product owner, ScrumMaster, and development team  should be able to track the development progress and be aware of criticle issues, etc. 

* Business owners, executives, and managers as internal stakeholders should see the progress firsthand so that they can suggest course corrections. 

* Sales, marketing, support, legal, compliance, and other Scrum and non-Scrum development teams might want to attend sprint reviews to provide area-specific feedback, etc. 

* Customers, users, partners, and regulators as External stakeholders can provide valuable feedback to the Scrum team 


### Change default permission scheme

* Default permission scheme is the built-in scheme, which you cannot delete it, but you can change it.

* Choose `Isssues` > `Permission schemes `

* Change all permissions from `Project Access - Any logged in user` to `Project role - Administrator` except the following items.
    * View Development Tools
    * View Read-Only Workflow
    * Assignable User
    * Delete Own Comments
    * Edit Own Comments
    * Delete Own Attachments
    * Delete Own Worklogs
    * Edit Own Worklogs
    * Work On Issues

* Copy the default permission scheme and rename it to `<client_business_name> Permission Scheme`. In this case it would be `ABC Permission Scheme`. 

* Create a new group named **ABC Group**, 

* Update `ABC Permission Scheme` by granting `ABC Group` to some items.
    * Browse Projects
    * Create Issues
* Update the permission scheme of the project `ABC IT Project` from to `ABC Permission Scheme`. 

### Change user's group 

* Choose `User management` > `User` 
* Add the ABC's user to group `ABC Group`
* Remove the user from default `jira-software-user` and make sure the checkbox `Access JIRA Software` is selected.

 
### Time Tracking and Estimation

#### Enable Time Tracking

* Choose `Issues` > `Time Tracking` 
    * Enable the `Time Tracking` if it is inactive.

* Choose `Boards` > `View All Boards` 
    * Choose `Board Settings` of the board
    * Choose `Estimation`
    * Choose `Remaining Estimate and Time Spent` as Time Tracking


