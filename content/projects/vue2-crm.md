+++
date = "2016-12-07T16:56:21+11:00"
title = "Vue 2 CRM Project"

description="Vue2Crm is a reusable Vue.js CRM starter project for real-world business based on Vue 2 PWA template"
+++

## Summary

**Vue2Crm** is a reusable Vue.js CRM starter project for real-world business based on Vue 2 PWA template with Vuetify.  

The goal of this project is to create a reusable project for real-world business. To achieve this target, we need a solution which includes authentication process, restful API feature and simple but elegant UI design. 


## __Features__

* This starter project is built-in with Vue 2 PWA from scratch.
* The whole UI is built on the Vuetify
* It includes Vuex, Axios as well.
* It uses Json-Server as fake Restful API. (You can simple replace it with your own API)

## Structure of Vue2Crm

``` ini
path\to\vue2crm
+---build              <-// webpack files
+---config             <-// configuration of dev or prod environment
+---db                 <-// json files for json-server
|   +---db.json        <-// dummy db
|   \---routes.json    <-// configure fake restful api
+---screenshots
+---src                <-// vue components 
|   +---components
|   |   +---404.vue
|   |   +---About.vue
|   |   +---Customers.vue
|   |   +---Customer.vue
|   |   +---Orders.vue
|   |   +---Order.vue
|   |   +---Login.vue
|   |   \---...
|   +---router         <-// vue-router
|   +---utils
|   |   +---auth.js    <-// auth service
|   |   +---backend-api.js  <-// Axios instance 
|   |   +---store.js   <-//  Vuex
|   \---stylus         <-// Customize stylus
+---static              <-// css, fonts, image files
|   +---img
|   \---manifest.json   <-// PWA manifest file
\---test
    +---e2e
    \---unit

```


## Screenshots

> ![Screenshot1](/img/v2crm-screenshot-1.jpg)

> ![Screenshot2](/img/v2crm-screenshot-2.jpg)

> ![Screenshot3](/img/v2crm-screenshot-3.jpg)

> ![Screenshot4](/img/v2crm-screenshot-4.jpg)



## Browse [Repository](https://github.com/harryho/vue2crm.git)



## __Alternatives__

There are two similar projects respectively built on the Angular and React. If you have interests in those technical stacks. You can find those projects below.

* [Angular4 Crm](/projects/angular4-crm).
* [React Redux Crm](/projects/react-crm).
