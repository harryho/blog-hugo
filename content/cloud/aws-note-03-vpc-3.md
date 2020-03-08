+++
title = "AWS : VPC - 3"
description = "VPC Peering, Direct Connect, Transit Gateway"
weight=5

+++


## VPC Part 3


### Endpoint

A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by AWS PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection.

Endpoints are virtual devices. They are horizontally scaled, redundant, and highly available VPC components. They allow communication between instances in your VPC and services without imposing availability risks or bandwidth constraints on your network traffic.



* Endpoint service — Your own application in your VPC. Other AWS principals can create a connection from their VPC to your endpoint service

* Gateway endpoint — A gateway endpoint is a gateway that you specify as a target for a route in your route table for traffic destined to a supported AWS service.

> Amazon S3

> DynamoDB


* Interface endpoint — An interface endpoint is an elastic network interface with a private IP address from the IP address range of your subnet that serves as an entry point for traffic destined to a supported service.


#### Limits of Interface Endpoint

* For each interface endpoint, you can choose only one subnet per Availability Zone.
* Interface endpoints support the use of policies for services that support endpoint policies.
* An interface endpoint supports TCP traffic only.
* Endpoints are supported within the same Region only. 
* Endpoints support IPv4 traffic only.



#### Limits of Gateway Endpoint

* You cannot use a prefix list ID in an outbound rule in a network ACL to allow or deny outbound traffic to the service specified in an endpoint. If your network ACL rules restrict traffic, you must specify the CIDR block (IP address range) for the service instead. 
* Endpoints are supported within the same Region only. 
* Endpoints support IPv4 traffic only.
* Cannot transfer an endpoint from one VPC to another,
* Endpoint connections cannot be extended out of a VPC. 
* Must enable DNS resolution in your VPC, or if you're using your own DNS server, ensure that DNS requests to the required service (such as Amazon S3) are resolved correctly to the IP addresses maintained by AWS. 


### Endpoint Service

#### Steps to the connections

* Create a Network Load Balancer for your application in your VPC and configure it for each subnet (Availability Zone) in which the service should be available. 
* Create a VPC endpoint service configuration and specify your Network Load Balancer.
* Grant permissions to specific service consumers (AWS accounts, IAM users, and IAM roles) to create a connection to your endpoint service.
* A service consumer that has been granted permissions creates an interface endpoint to your service, optionally in each Availability Zone in which you configured your service.
* To activate the connection, accept the interface endpoint connection request. By default, connection requests must be manually accepted. However, you can configure the acceptance settings for your endpoint service so that any connection requests are automatically accepted.
* To help achieve high availability for service consumers that use zonal DNS hostnames to access the service, you can enable cross-zone load balancing. Cross-zone load balancing enables the load balancer to distribute traffic across the registered targets in all enabled Availability Zones.

#### Endpoint Service DNS Names

AWS generates endpoint-specific DNS hostnames that you can use to communicate with the service. These names include the VPC endpoint ID, the Availability Zone name and Region Name, for example, vpce-1234-abcdev-us-east-1.vpce-svc-123345.us-east-1.vpce.amazonaws.com. By default, your consumers access the service with that DNS name and usually need to modify the application configuration. 


#### Endpoint Service Limitations

* An endpoint service supports IPv4 traffic over TCP only.
* Service consumers can use the endpoint-specific DNS hostnames to access the endpoint service, or the private DNS name.
* If an endpoint service is associated with multiple Network Load Balancers, then for a specific Availability Zone, an interface endpoint establishes a connection with *one load balancer only.
* For the endpoint service, the associated Network Load Balancer can support 55,000 simultaneous connections or about 55,000 connections per minute to each unique *target (IP address and port). 
* Availability Zones in your account might not map to the same locations as Availability Zones in another account. 
* Review the service-specific limits for your endpoint service.



#### VPC Endpoint Policies

A VPC endpoint policy is an IAM resource policy that you attach to an endpoint when you create or modify the endpoint. If you do not attach a policy when you create an endpoint, AWS attaches a default policy for you that allows full access to the service. 




### VPC Peering

A VPC peering connection is a networking connection between two VPCs that enables you to route traffic between them privately. Instances in either VPC can communicate with each other as if they are within the same network. You can create a VPC peering connection between your own VPCs, with a VPC in another AWS account, or with a VPC in a different AWS Region.

AWS uses the existing infrastructure of a VPC to create a VPC peering connection; it is neither a gateway nor an AWS Site-to-Site VPN connection, and does not rely on a separate piece of physical hardware. There is no single point of failure for communication or a bandwidth bottleneck.

