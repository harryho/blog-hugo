+++
title = "AWS: Mesh - 1"
description = "Getting started"
draft="true"
+++


## Mesh 

AWS App Mesh is a service mesh based on the Envoy proxy that helps you monitor and control services. App Mesh standardizes how your services communicate, giving you end-to-end visibility into and helping to ensure high-availability for your applications. 



#### Create a mesh and virtual service

A service mesh is a logical boundary for network traffic between the services that reside within it.


  - A mesh named pgMesh, since all of the services in the scenario are registered to the pgMesh.local namespace.

  - A virtual service named svc1.pgMesh.local, since the virtual service represents a service that is discoverable with that name, and you don't want to change your code to reference another name. A virtual service named svc2.pgMesh.local is added in a later step.

        aws appmesh create-mesh --mesh-name pgMesh


  - Create a virtual service with the create-virtual-service command.

        aws appmesh create-virtual-service \
                --mesh-name pgMesh \
                --virtual-service-name svc1.pgMesh.local \
                --spec {}

#### Create a virtual node
  
A virtual node acts as a logical pointer to an actual service. 

Create a virtual node named svcNode1, since one of the virtual nodes represents the actual service named svcNode1. The actual service that the virtual node represents is discoverable through DNS with a hostname of svc1.pgMesh.local. Alternately, you can discover actual services using AWS Cloud Map. The virtual node will listen for traffic using the HTTP/2 protocol on port 80. 


* Create a file named create-virtual-node-svc1.json with the following contents:

```json
{
    "meshName": "pgMesh",
    "spec": {
        "listeners": [
            {
                "portMapping": {
                    "port": 80,
                    "protocol": "http2"
                }
            }
        ],
        "serviceDiscovery": {
            "dns": {
                "hostname": "svcNode1.pgMesh.local"
            }
        }
    },
    "virtualNodeName": "svcNode1"
}
```

* Create the virtual node with the create-virtual-node command using the JSON file as input.

        aws appmesh create-virtual-node --cli-input-json file://create-virtual-node-svc1.json


  
#### Create a virtual router and route
  
Virtual routers route traffic for one or more virtual services within your mesh. 

Create the following resources:

> A virtual router named svcNode1, since the svcNode1.pgMesh.local virtual service doesn't initiate outbound communication with any other service. Remember that the virtual service that you created previously is an abstraction of your actual svc1.pgMesh.local service. The virtual service sends traffic to the virtual router. The virtual router will listen for traffic using the HTTP/2 protocol on port 80. Other protocols are also supported.

> A route named svcNode1. It will route 100 percent of its traffic to the svcNode1 virtual node. You'll change the weight in a later step once you've added the svcNode1v2 virtual node. Though not covered in this guide, you can add additional filter criteria for the route and add a retry policy to cause the Envoy proxy to make multiple attempts to send traffic to a virtual node when it experiences a communication problem.


#### Create a virtual router

Create a file named create-virtual-router.json with the following contents:

```json
{
    "meshName": "pgMesh",
    "spec": {
        "listeners": [
            {
                "portMapping": {
                    "port": 80,
                    "protocol": "http2"
                }
            }
        ]
    },
    "virtualRouterName": "svcNode1"
}
```


* Create the virtual router with the create-virtual-router command using the JSON file as input.

        aws appmesh create-virtual-router --cli-input-json file://create-virtual-router.json

* Create a route.

  - Create a file named create-route.json with the following contents:

```json
{
    "meshName" : "pgMesh",
    "routeName" : "svcNode1",
    "spec" : {
        "httpRoute" : {
            "action" : {
                "weightedTargets" : [
                    {
                        "virtualNode" : "svcNode1",
                        "weight" : 100
                    }
                ]
            },
            "match" : {
                "prefix" : "/"
            }
        }
    },
    "virtualRouterName" : "svcNode1"
}
```

* Create the route with the create-route command using the JSON file as input.

        aws appmesh create-route --cli-input-json file://create-route.json


#### Review and create

Review the settings against the previous instructions.

* Review the settings of the mesh you created with the describe-mesh command.

        aws appmesh describe-mesh --mesh-name pgMesh

* Review the settings of the virtual service that you created with the describe-virtual-service command.

        aws appmesh describe-virtual-service --mesh-name pgMesh --virtual-service-name svc1.pgMesh.local

* Review the settings of the virtual node that you created with the describe-virtual-node command.

        aws appmesh describe-virtual-node --mesh-name pgMesh --virtual-node-name svcNode1

* Review the settings of the virtual router that you created with the describe-virtual-router command.

        aws appmesh describe-virtual-router --mesh-name pgMesh --virtual-router-name svcNode1

* Review the settings of the route that you created with the describe-route command.

        aws appmesh describe-route --mesh-name pgMesh \
                --virtual-router-name svcNode1  --route-name svcNode1


#### Create additional resources

Create one virtual node named svcNode1v2 and another named svcNode2. Both virtual nodes listen for requests over HTTP/2 port 80. For the svcNode2 virtual node, configure a backend of svc1.pgMesh.local, since all outbound traffic from the svcNode2 virtual node is sent to the virtual service named svc1.pgMesh.local. Though not covered in this guide, you can also specify a file path to write access logs to for a virtual node.

Create one additional virtual service named svc2.pgMesh.local, which will send all traffic directly to the svcNode2 virtual node.

Update the svcNode1 route that you created in a previous step to send 75 percent of its traffic to the svcNode1 virtual node and 25 percent of its traffic to the svcNode1v2 virtual node. Over time, you can continue to modify the weights until svcNode1v2 receives 100 percent of the traffic. 

Once all traffic is sent to svcNode1v2, you can deprecate the svcNode1 virtual node and actual service. As you change weights, your code doesn't require any modification, because the svc1.pgMesh.local virtual and actual service names don't change. Recall that the svc1.pgMesh.local virtual service sends traffic to the virtual router, which routes the traffic to the virtual nodes. The service discovery names for the virtual nodes can be changed at any time.

* Create the svcNode1v2 virtual node.

* Create a file named create-virtual-node-svcnode1v2.json with the following contents:

```json
{
    "meshName": "pgMesh",
    "spec": {
        "listeners": [
            {
                "portMapping": {
                    "port": 80,
                    "protocol": "http2"
                }
            }
        ],
        "serviceDiscovery": {
            "dns": {
                "hostname": "svcNode1v2.pgMesh.local"
            }
        }
    },
    "virtualNodeName": "svcNode1v2"
}
```


* Create the virtual node.

        aws appmesh create-virtual-node --cli-input-json file://create-virtual-node-svcnode1v2.json

* Create the svcNode2 virtual node.

* Create a file named create-virtual-node-svc2.json with the following contents:

```json
{
   "meshName" : "pgMesh",
   "spec" : {
      "backends" : [
         {
            "virtualService" : {
               "virtualServiceName" : "svc1.pgMesh.local"
            }
         }
      ],
      "listeners" : [
         {
            "portMapping" : {
               "port" : 80,
               "protocol" : "http2"
            }
         }
      ],
      "serviceDiscovery" : {
         "dns" : {
            "hostname" : "svc2.pgMesh.local"
         }
      }
   },
   "virtualNodeName" : "svcNode2"
}

```


* Create the virtual node


        aws appmesh create-virtual-node --cli-input-json file://create-virtual-node-svc2.json


Update the svc1.pgMesh.local virtual service that you created in a previous step to send its traffic to the svcNode1 virtual router. When the virtual service was originally created, it didn't send traffic anywhere, since the svcNode1 virtual router hadn't been created yet.

Create a file named update-virtual-service.json with the following contents:

```json
{
   "meshName" : "pgMesh",
   "spec" : {
      "provider" : {
         "virtualRouter" : {
            "virtualRouterName" : "svcNode1"
         }
      }
   },
   "virtualServiceName" : "svc1.pgMesh.local"
}
```


* Update the virtual service with the update-virtual-service command.

        aws appmesh update-virtual-service --cli-input-json file://update-virtual-service.json

* Update the svcNode1 route that you created in a previous step.

  - Create a file named update-route.json with the following contents:

```json
{
   "meshName" : "pgMesh",
   "routeName" : "svcNode1",
   "spec" : {
      "http2Route" : {
         "action" : {
            "weightedTargets" : [
               {
                  "virtualNode" : "svcNode1",
                  "weight" : 75
               },
               {
                  "virtualNode" : "svcNode1v2",
                  "weight" : 25
               }
            ]
         },
         "match" : {
            "prefix" : "/"
         }
      }
   },
   "virtualRouterName" : "svcNode1"
}
```

* Update the route with the update-route command.

        aws appmesh update-route --cli-input-json file://update-route.json

* Create the svcNode2 virtual service.

  - Create a file named create-virtual-svc2.json with the following contents:


```json
{
   "meshName" : "pgMesh",
   "spec" : {
      "provider" : {
         "virtualNode" : {
            "virtualNodeName" : "svcNode2"
         }
      }
   },
   "virtualServiceName" : "svc2.pgMesh.local"
}
```

* Create the virtual service.

        aws appmesh create-virtual-service --cli-input-json file://create-virtual-svc2.json


#### Mesh summary


> Before you created the service mesh, you had three actual services named svc2.pgMesh.local, svc1.pgMesh.local, and servicebv2.pgMesh.local. In addition to the actual services, you now have a service mesh that contains the following resources that represent the actual services:

> Two virtual services. The proxy sends all traffic from the svc2.pgMesh.local virtual service to the svc1.pgMesh.local virtual service through a virtual router.

> Three virtual nodes named svcNode2, svcNode1, and svcNode1v2. The Envoy proxy uses the service discovery information configured for the virtual nodes to look up the IP addresses of the actual services.

> One virtual router with one route that instructs the Envoy proxy to route 75 percent of inbound traffic to the svcNode1 virtual node and 25 percent of the traffic to the svcNode1v2 virtual node.

