+++
title = "Azure: CAF - 1"
weight = 1
description="Introduction of Cloud Adoption Framework"
+++

## Cloud Adoption Framework

The Microsoft Cloud Adoption Framework for Azure is proven guidance that's designed to help you create and implement the business and technology strategies necessary for your organization to succeed in the cloud. It provides best practices, documentation, and tools that cloud architects, IT professionals, and business decision makers need to successfully achieve short-term and long-term objectives.

The Cloud Adoption Framework brings together cloud adoption best practices from Microsoft employees, partners, and customers. It provides a set of tools, guidance, and narratives that help shape technology, business, and people strategies for driving desired business outcomes during your cloud adoption effort. Review the guidance for each methodology below, providing you with easy access to the right guidance at the right time.

### Structure

methodology |  descripton |
---|----
Strategy| define business justification and expected outcomes of adoption.		
Plan| align actionable adoption plans to business outcomes.
Ready| Prepare the cloud environment for the planned changes.		
Migrate| Migrate and modernize existing workloads.
Innovate| Develop new cloud-native or hybrid solutions.	
Govern| Govern the environment and workloads.
Manage| Operations management for cloud and hybrid solutions.	
Organize| Govern the environment and workloads.


### RACI

RACI - Responsible, Accountable,Consulted,and Informed


### Cloud Adoption Capability

Capability	| Description
----|----
Cloud Adoption	|Deliver technical solutions
Cloud Strategy	|Align technical change to business needs
Cloud Operations	|Support and operate adopted solutions
Cloud Center of Excellence	|Improve quality, speed, and resiliency of adoption
Cloud Governance	|Manage risk, drive consistency, governance, and compliance
Cloud Platform	|Operate and mature the platform
Cloud Automation	| Accelerate adoption and innovation


### SME of Capability

Capability	| Expertise
----|-----
Cloud Strategy  |   Finance
Cloud Strategy	|   Project Management
Cloud Governance|	IT Security
Cloud Governance|	IT Governance
Cloud Platform|	Network
Cloud Platform|	Identity
Cloud Platform|	Virtualization
Cloud Platform|	Disaster Recovery
Cloud Operations|	IT Operations
Cloud Operations|	Application Owner
Cloud Adoption|	Project Lead(s)
Cloud Adoption|	Architect(s)


### Lifecycle


{{<mermaid>}}
graph LR
    S(Define Strategy)
    P(Plan)
    R(Ready)
    A(Adopt)
   subgraph Layer_1
        S-->P
        P-->R
        R-->A
    end 
{{</mermaid >}}
{{<mermaid>}}
graph BT
    G(Govern)
    M(Management)
    subgraph Layer_2
        G
        M
    end
{{</mermaid >}}

