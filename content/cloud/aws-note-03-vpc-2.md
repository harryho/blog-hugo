+++
title = "AWS - VPC part 2"
description = "VPC -  "
weight=3
draft=true
+++


## VPC Part 2

VPC has an implicit router (implied router), and you use route tables to control where network traffic is directed.

### Route table

* Have up to 200 route tables per VPC
* Have up to 50 route entries per route table
* Each subnet must be associated with only one route table
* The subent (when created) will be associated with main (default) VPC route table
* Can change the subnet association to another route table
* Can NOT delete the main route table
* Every route table in a VPC comes with a default rule that allows all VPC subnets to comminunicate with one another. This rule can NOT be modified or deleted.

### Security Group 


* It is a virtual firewall
* It controls traffic at the EC2ss level
* Up to 5 security groups per EC2 
* Stateful, return traffic, of allowed inbound traffic, is allowed, even if there are no rules to allow it. 
* Can only have permit rules, can NOT have deny rules
* Implicit deny rule at the end
* Security group is associated wth EC2's network interface
* Any change on security group takes effect immediately
* Default security groupd can not be deleted
* It is VPC resource, hence, different EC2 in differenet AZs within the same VPC, can have the same security group applied to them.
* It can NOT block a certain range of IP addresses from Internet from gettting to EC2 fleets

#### Default vs Customized Group

The default security group has inbound and outbound rules when created. The inboud rule allows all traffics in from the same security group. The outbound rule allows all traffics to any destination. The customized security group has outbound rule only by default.

### Network ACL

* A default NACL allows all traffic inbound and outbound.
* A cutomized NACL blocks/denies all traffic inbound / outbound by default.
* Can help you block certain ranges of IP address from a large pool(Internet address for instance)

### Security Group vs NACL

Security Group | NACL 
---------------|----------
Operates at the instance level| Operates at the subnet level
Supports allow rules only| Supports allow rules and deny rules
Is stateful: Return traffic is automatically allowed, regardless of any rules| Is stateless: Return traffic must be explicitly allowed by rules
We evaluate all rules before deciding whether to allow traffic| We process rules in number order when deciding whether to allow traffic
Applies to an instance only if someone specifies the security group when launching the instance, or associates the security group with the instance later on| Automatically applies to all instances in the subnets that it's associated with (therefore, it provides an additional layer of defense if the security group rules are too permissive)


### NAT instance



### Endpoint






### NAT Gateway





