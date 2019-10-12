+++

date = "2016-12-07T16:56:21+11:00"
title = "Angular 4 CRM Project"
description="Ng4Crm is reusable CRM project for real-world business based on Angular 4"
+++

## Summary

**Ng4Crm** is reusable CRM project for real-world business based on Angular 4, Angular-Material & Bootstrap 3.

 This project starts from a popular starter project [AngularClass/AngularStarter](https://github.com/AngularClass/angular-starter). The goal of this project is to create reusable project for real-world business. To achieve this target, we need a solution which should include authentication process, restful API feature with token support and simple but elegant UI design. 




## __Features__

* This project is built on the top of AngularClass/Angular-Starter. 

* The UI part of this project combine Angular-Material and Bootstrap 3. Because the controls from Angular-Material is very limited, there are a few extra components require Bootstrap 3 for help. 

* This project includes ng-charts, pagination, progress-bar, confirmation dialog, etc. features.

* It uses Json-Server as fake Restful API. (You can simple replace it with your own API)


## Structure of Ng4Crm

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

> ![Screenshot1](/img/ng4crm-screenshot-1.jpg)

> ![Screenshot2](/img/ng4crm-screenshot-2.jpg)

> ![Screenshot3](/img/ng4crm-screenshot-3.jpg)

> ![Screenshot4](/img/ng4crm-screenshot-4.jpg)



## Browse [Repository](https://github.com/harryho/ng4crm.git)


## __Alternatives__

There are two similar projects respectively built on the Vue.js and React. If you have interests in those technical stacks. You can find and clone those projects below.

* [Vue2 Crm](/projects/vue2-crm).
* [React Redux Crm](/projects/react-crm).
