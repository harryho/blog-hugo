
+++

date = "2016-12-07T16:56:21+11:00"
title = "React Redux CRM Project"
tags = ["javascript"]
description="React-Crm is reusable CRM starter project for real-world business based on React 15.4"
+++


## Summary

**React-Crm** is reusable CRM starter project for real-world business based on React 15.4, React-Redux & Material-UI.


The goal of this starter project is to create reusable project for real-world business. To achieve this target, we need a solution which should include authentication process, restful API feature with token support and simple but elegant UI design. 


## __Features__

* This project is built on the top of React/Redux. 
* The UI part of this project uses Material-UI. 
* This project uses Redux-Thunk to support backend API.
* It uses Json-Server as fake Restful API. (You can simple replace it with your own API)

## Structure of React-Crm

``` ini
path\to\ng4crm
+---config             <-// configuration of dev or prod environment
+---db                 <-// json files for json-server
|   +---db.json        <-// dummy db
|   \---routes.json    <-// configure fake restful api
+---screenshots
+---src                <-// vue components 
|   +---app
|   |   +---_gurad             <-// auth guard for authentication
|   |   +---_models             <-// common models for whole app
|   |   +---_services            <-// common services for whole app
|   |   +---about                <-// about component   
|   |   +---customer              <-// customer component
|   |   +---dashboard            <-// dashboard component  
|   |   +---notfoundpage         <-// notfoundpage component  
|   |   +---login                <-// login component  
|   |   +---order                <-// customer component 
|   |   +---root                <-// root component 
|   |   +---shared                <-// common component for whole app
|   |   +---app.component.ts
|   |   +---app.module.ts
|   |   +---app.routes.ts
|   |   +---app.services.ts
|   |   +---environment.ts
|   |   \---...
|   +---assets         <-//   images  and css from third parties
|   +---styles         <-//   customized css files
|   +---main.browser.aot.ts     
|   +---main.browser.ts  
|   +---polyfills.browser.ts  
|   \---...
\...

```


## Screenshots

> ![Screenshot1](/img/rrcrm-screenshot-1.jpg)

> ![Screenshot2](/img/rrcrm-screenshot-2.jpg)

> ![Screenshot3](/img/rrcrm-screenshot-3.jpg)

> ![Screenshot4](/img/rrcrm-screenshot-4.jpg)


## Browse [Repository](https://github.com/harryho/react-crm.git)



__Alternatives__

There are two similar projects respectively built on the Vue.js and Angular. If you have interests in those technical stacks. You can find and clone those projects below.

* [Vue2 Crm](/projects/vue2-crm/).
* [Angular4 Crm](/projects/angular4-crm).

