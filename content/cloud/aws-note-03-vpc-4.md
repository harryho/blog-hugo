+++
title = "AWS : VPC - 4"
description = "VPC - Simple Demo "
weight=6

+++


## VPC Part 4




### Simples demo


* Diagram of customized VPC - MyDemoVPC with Internet Gatway and VPN connect

{{<mermaid>}}
graph LR
    InternetGW(Internet Gateway)
    VirtualGW(Virtual Gateway)
    INTER(Internet - Public)
    InternetGW --- INTER
    VirtualGW --- SERVER
    subgraph MyDemoVPC
        MainRouteTable --- InternetGW
        MainRouteTable --- NetworkACL
        NetworkACL --- PubSecGrp
        NetworkACL --- PrivSecGrp
        EC2_A(EC2 Instannce A) 
        EC2_B(EC2 Instannce B) 
        EC2_E(EC2 Instannce E) 
        EC2_F(EC2 Instannce F)
        EC2_C[(Database Master)]
        EC2_D[(Database Slave)]
        PrivSecGrp --- PrivSubet
        VPNSubnet --- VirtualGW
        VPNSubnet --- MainRouteTable
        PubSecGrp --- PubSubnet
        subgraph Router
            MainRouteTable(10.0.0.0/16)
        end 
        subgraph Private_Subnet
             PrivSubet(10.0.2.0/24)
             EC2_C
             EC2_D
        end
        subgraph Public_Subnet
            PubSubnet(10.0.1.0/24)
            EC2_A
            EC2_B
        end
        subgraph VPN_Subnet
             VPNSubnet(10.0.3.0/24)
             EC2_E
             EC2_F
        end
    end
    InternetGW
    subgraph Internet
        INTER
    end
    subgraph OnPremise
        SERVER
    end


{{</mermaid>}}



* Customized Route tables of Subnet Public_Subnet

Destination | Target 
----- | ------- 
10.0.1.0/16 | local 
2002:0a00:0100:0:0:0:0:0/56 | local 
0.0.0.0/0 | InternetGW 
::0/0 | InternetGW 


* Main Route tables of Subnet Private_Subnet

Destination | Target
-----|-------
10.0.1.0/16 | local
2002:0a00:0100:0:0:0:0:0/56 | local


* Route table of Subnet VPN_Subnet

Destination | Target
-----|-------
10.0.1.0/16 | local
2002:0a00:0100:0:0:0:0:0/56 | local
0.0.0.0/0 | VirtualGW
