+++
tags =  ["agile"]
categories = ["blog"]
date = "2017-01-11T11:59:31+11:00"
title = "JIRA Tricks"
draft = true
+++

## JIRA

> Jira is a proprietary issue tracking product, developed by Atlassian. It provides bug tracking, issue tracking, and project management functions. Although normally styled JIRA, the product name is not an acronym, but a truncation of Gojira, the Japanese name for Godzilla,[6] itself a reference to Jira's main competitor, Bugzilla. It has been developed since 2002.

> 

## How to work with client within JIRA

> The instruction below is only for the client who has independent JIRA instance. Here I am not going to discuss how to cooperate with multiple clients and projects on the same JIRA instance. IMO, the strategy for that would be case by case. 

### Assumption

* Your client's business name is ABC. You help them to manage their JIRA. You are the administrator of JIRA instance. 

* You have created a project for client, named **ABC IT Project**

* You want to allow the client to access this project on JIRA, but you want to avoid client's unintented update to mess up the project management. 

### Requirement of permission control

* In Agile practice, product owner should attend the 


### Change default permission scheme

* Default permission scheme is the built-in scheme, which you cannot delete it, but you can change it.

* Access it by choosing [gear icon]` > Isssues > Permission schemes `

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

* Copy the default permission scheme and rename it to `<client_business_name> permission scheme`

* Create a new group named **ABC Group**, 

* 

 



