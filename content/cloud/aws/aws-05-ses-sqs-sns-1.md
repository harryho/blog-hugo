+++
title = "AWS: SQS,SNS,SES - 1"
description = "Introduction of SQS, SNS, SES"
weight=7
draft=true
+++

## SQS


Amazon Simple Queue Service (SQS) is a fully managed message queuing service that enables you to decouple and scale microservices, distributed systems, and serverless applications. SQS eliminates the complexity and overhead associated with managing and operating message oriented middleware, and empowers developers to focus on differentiating work. Using SQS, you can send, store, and receive messages between software components at any volume, without losing messages or requiring other services to be available. Get started with SQS in minutes using the AWS console, Command Line Interface or SDK of your choice, and three simple commands.



__Queue types__

### Standard Queues

* Unlimited Throughput: Standard queues support a nearly unlimited number of transactions per second (TPS) per API action.

* At-Least-Once Delivery: A message is delivered at least once, but occasionally more than one copy of a message is delivered.

* Best-Effort Ordering: Occasionally, messages might be delivered in an order different from which they were sent.

### FIFO Queues

* High Throughput: By default, FIFO queues support up to 300 messages per second (300 send, receive, or delete operations per second). When you batch 10 messages per operation (maximum), FIFO queues can support up to 3,000 messages per second. To request a quota increase, file a support request.

* Exactly-Once Processing: A message is delivered once and remains available until a consumer processes and deletes it. Duplicates aren't introduced into the queue.

* First-In-First-Out Delivery: The order in which messages are sent and received is strictly preserved (i.e. First-In-First-Out).


## SNS

Amazon Simple Notification Service (SNS) is a fully managed messaging service for both system-to-system and app-to-person (A2P) communication. It enables you to communicate between systems through publish/subscribe (pub/sub) patterns that enable messaging between decoupled microservice applications or to communicate directly to users via SMS, mobile push and email.

> There are two main uses for SNS. First, you can use it as a traditional pub/sub messaging system. An example here is a microservice architecture where Service A may be interested in updates to objects in Service B. Rather than Service B directly notifying Service A about the update, Service B can send a message to an SNS Topic with details about the update. Service A can subscribe to the topic and process the messages as they arrive. The main alternatives to using SNS in this manner are tools like RabbitMQ in pub/sub mode or NATS.

> The second core use case for SNS is to deliver messages to large numbers of end users, such as via mobile push notifications or SMS messages. SNS allows for extremely high fan-out in these use cases, as you can have up to 12.5 million subscribers to a single topic. This is great for blasting your users with updates about a new sale or news in your application.

## SES

Amazon Simple Email Service (SES) is a cost-effective, flexible, and scalable email service that enables developers to send mail from within any application. You can configure Amazon SES quickly to support several email use cases, including transactional, marketing, or mass email communications. Amazon SES's flexible IP deployment and email authentication options help drive higher deliverability and protect sender reputation, while sending analytics measure the impact of each email. With Amazon SES, you can send email securely, globally, and at scale.


> The core main use case for SES is sending email. There are two email types. First is transactional email. Transaction email is when you send an automated message to a specific person. For example, you could send a customer an email that their order has shipped, or you can alert a user about the subscription  ending.

> The second bucket of outgoing email is marketing email. This is when you send the same email to a large number of users at the same time. This could be notifying your email list of a large upcoming sale or making an announcement of a new product.