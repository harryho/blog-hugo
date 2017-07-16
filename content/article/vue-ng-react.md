+++
categories = ["article"]
date = "2017-06-07T16:56:21+11:00"
title = "Angular vs React vs Vue"
tags = ["javascript"]
+++

>Angular, React, Vue are most popular JavaScript frameworks at present. Here we just discuss Angular 1.x, Angular 2 / 4, React 15+ (Redux), and Vue 2+. 


## Client side is a battle field

In past 6-8 years, the Restful API has been accepted as one of standard web interface for most web applications, solution architect can simply add REST API on the top of existing web layer to provide REST API and support multiple client devices. So the developer can continue to maintain or develop system with their favorite framework or technical stacks. 

On the client-side, it is a completely different story, there are tons of variant JavaScript frameworks emerged in the last 10 years. It is good to have more options, but it is a nightmare for web developer who is working in such battle field. Especially, when the development team try to adopt new client-side framework for the new project, it is very difficult to make a decision which framework we should choose. 

There is a big risk to adopt new programming language(ES 6 or Typescript), new  framework, as well as new development, build and test tool. As a solution architect, he or she needs to think it through for development team, and also consider if the team can really pick it up quickly. That is why we have to compare those frameworks here before we make a decision. 

## Performance is not a priority or crucial criterion for choosing a framework

We can find lots of comparison between those frameworks, and so many of them are related to performance, programming language, design pattern, etc. Actually many web applications in the world are just small to median size web application, we don't need to build the web application as Google, Facebook or Twitter. In my opinion, the performance of framework is the not critical benchmark, at lease it is not first priority which we need to consider if it is right for the team. Except performance, we have more concern on tech stacks, community and ecosystem involved with the framework, which have more impact on team's productivity and system's maintainability. 


## The difference of those framework

Let's look into the frameworks and list the difference of these frameworks.

### Basic tech stacks


Tech Stacks   |    Angular 1.x   | Angular 2 / 4 | React 1.5 (Redux)  | Vue 2   
----------------|-------------------| ---------------|---------------------|---------
Vanilla JS     |     Yes       |  Supported  |  Supported       | Supported
ES 6          |     1.5+          |  Supported  |  Yes            | Yes
TypeScript    |               |  Yes        |                | Supported |
MVC           |   1.2-1.4 |            | |
Component-Based |  1.5+ |  Yes   | Yes |  Yes 
Shadow DOM | | Yes | |
Virtual DOM| | | Yes | Yes
Immutable state  | | | Yes | Yes



* Yes:  Programming language which the framework uses.
* Supported: Programming language which the framework supports.


### Cool stuff is not always the best

We have seen so many cool stuff which are finally abandoned in the past. We shouldn't choose new framework because it looks cool. We choose the new one because it really solve our existing problem, improve our productivity and quality in the long run. Don't forget there is always some cost to adopt new stuff. We need to balance the cost and outcome of such technical investment, and we can work it out if it is right time to do it. 


### Programming language is the barrier

Programming language is still the barrier we have to evaluate with the existing development team. Even ES6 or TypeScript (TS) claims it is compatible to Vanilla JS, but when you start to look into new framework or sample project, which are coded with ES6 or TS, it still makes you so confused if you are not familiar with such syntax. It will significantly impact the efficiency of learning new framework. So there is always a learning curve, which we cannot ignore, to code something in a new programming. If your team comes from .Net Web Form or Java MVC background, it would be a steep curve for the team to pick up ES6 or TypeScript and Component-based framework, not mentioned new build and test tools. 

No wonder some .Net teams were struggling with Node.js integration on Visual Studio, especially when the team members have no Node.js experience. So we need the whole team to discuss the difficulties before we adopt new technology and framework. It is helpful to make sure the team has the same view, and it is good for us to plan our training and decide how to transform development team step by step. 


### What to start 

For the team which comes with Web Form, with Vanilla JS background, we can start with Angular 1.x (Up to 1.4) on some small projects, or we can build something training project, because the MVC pattern is very similar to their previous coding experience. 

For the team which has experience of Angular 1.2 ~ 1.4, they can choose to stay on later version of Angular 1.x, e.g. Angular 1.5+, and they can start to convert coding pattern from MVC to Component-based. After that, if the team is planning to move to Angular 2 / 4, it is better to do some TypeScript training. In my view, so far the ecosystem for Angular 2 / 4 is still under development. It is a bit risk to use Angular 2 / 4 to real-world production. There are quite many gotchas which you have to figure out on your own. 

For the team which has TypeScript or ES6 experience, they can choose what they prefer. They can spend more time on UI integration. There are a few customized UI package for bespoke framework. That is what we are going to discuss in the next. 



### Responsive UI library support

To build a real-world application, we need to integrate some popular responsive UI libraries instead of building all styles on our own. Let's take a look the support of Bootstrap or Material-Design for different frameworks. 


 UI library   |    Angular 1.x     | Angular 2 / 4 | React 1.5 (Redux)  | Vue 2   
--------------|------------------- | ---------------|---------------------|---------
Bootstrap 3  | ui-bootstrap (Very Good) |            | react-bootstrap(Very Good) | VueStrap* (Very Good)
Bootstrap 4  |                          | ui-bootstrap (Alpha) | In progress | BootstrapVue (Good) |
Material Design | Materialize (Good)   | Angular Material(Basic)  |  Material-UI (Good) | Vuetify (Very Good)



* VueStrap: Please use the [Willen's fork](https://github.com/wffranco/vue-strap) for Vue 2.
* Libraries in the table above has been tested or used in some projects.


From what we can see now, the Bootstrap 4 is similar to Material-Desgin. So it is good news for developer. They just need to pick their favor, and they will always get analogical effect. 


### Stable API 

Against to Angular 1.x, the Angular 2 is a complete new animal. Angular 4 comes with breaking changes, which breaks many Angular 2 dependencies. Since the API is still changing, we cannot use it for production. According to Angular team's announcement, they want to fix all Angular-2's bugs and issues on Angular 4 and keep all built-in libraries sync to Angular 4. It will take a long while to get things ready. If your project uses Angular 1.2-1.4, I'd like to suggest you to keep it. 

React-Redux is much more popular than Flux, but it doesn't means it is better than React-Flux pattern. In my opinion, React-Flux is much more straight and close to original React design. If you ready use React-Flux, you have better to stick with it. 

Vue 2 comes with some breaking changes. There is migration guide for Vue 1.x to Vue 2. It doesn't seem very different. Vue 2 is ready for production. 


### How to compare  

In order to compare those frameworks properly, I use those frameworks to create a small real-world web application, which has built-in authentication support for the back-end API service, and integrated with some responsive UI framework, e.g. Bootstrap or Material-Design. 



### Following are the projects and related screenshots


[Angular 4 CRM](/project/reetek-angular4-crm/)

![Screenshot-Angular4Crm](/img/ng4crm-screenshot-2.jpg)

[React CRM](/project/reetek-react-crm/)

![Screenshot-React-Crm](/img/rrcrm-screenshot-2.jpg)

[Vue 2 CRM](/project/reetek-vue2-crm/)

![Screenshot-Vue2crm](/img/v2crm-screenshot-2.jpg)



### Comparison of projects for different framework

Let's go back to projects above and take a look. They implement almost the same features as real-world application. 

__Features__


* Authentication & Token support for Restful API
* Customer CRUD functions
* Order CRUD functions
* Dashboard including two charts (Bar/Line/Doughnut)
* Integrate with Material Design (Angular project include)


|             |Angular 4 CRM   |  React Redux CRM  | Vue 2 CRM
--------------|------------------- |------------------|-------------
Dependencies |       22           |      13           |  9
Code Size    |     135KB          |     113KB         |  49KB
Effort       |     72 hrs         |    80 hrs         |    48hrs


* Effort: The effort for learning curve has been eliminated.


According to above the dependencies, code size, we can see the project based on Vue.js is much more simple than other two projects. 

Vue.js makes the switch between immutable and mutable much easier than React. It combines advantages of Angular and React. Vue.js template is very handy and straight. It is the same as normal HTML, it is very easy to convert the mock-up HTML into Vue template. 

Vue.js is not just cool, it is elegant and simple. It is really a good combination of Angular and React. I am pretty if you have Angular or React background, you can pick it up in a couple days. Once you start to use it, you won't want to go back. Its official routing system is quite stable and easy to use. Compare with Angular-Router and React-Router, it is much more reliable. 

Generally, Material-Design libraries for React is not handy as other customized version for Angular or Vue. The special coding style of JSX needs to convert all CSS and HTML into JSX format. To be honest, I am not so convinced by React's JSX, because it is not straight as final HTML or CSS. Compare to other framework, it is a bit verbose and inconvenient. We cannot simply copy the style code from the browser's dev tool when we debug it on the browser. 

Angular's Material-Design library has very limit components. To build a real-world application you need to add another UI library to supplement the former missing components. Last but not least, the Vuetify is the best Material-Design so far we have found and tested. 


### Summary

Before we make any conclusion, we have to be aware the world keeps changing. Perhaps after I finish this article, some problem of framework have been solved, or some resolved problem comes back again. We need to review the decision we made and correct them ASAP if we find the cost is overweight the outcome. 

* Team with Web Form and Vanilla Js background should starts with Angular 1.4 and be familiar with Node.js
* Team with Web Form and Vanilla Js background should start to learn ES 6 or TypeScript.
* Team with ES6 /TypeScript background can start should any framework you prefer
* Angular 4 and its ecosystem is under active development. 
* Vue.js framework is a very nice. Give a go on your next project.