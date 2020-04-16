+++
title = "AWS: EKS - 6"
description = "Labs"
weight=16

draft=true      
+++


## EKS - Part 6

### Guestbook

* Create the Redis master replication controller.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook-go/redis-master-controller.json


* Create the Redis master service.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook-go/redis-master-service.json



* Create the Redis slave replication controller.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook-go/redis-slave-controller.json


* Create the Redis slave service.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook-go/redis-slave-service.json

* Create the guestbook replication controller.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook-go/guestbook-controller.json


* Create the guestbook service.

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook-go/guestbook-service.json



* Query the services in your cluster and wait until the External IP column for the guestbook service is populated.

        kubectl get services -o wide

After your external IP address is available, point a web browser to that address at port 3000 to view your guest book. For example, 

        http://a2444a44644eb431ca8b9f7617d85aad-1238693525.ap-southeast-2.elb.amazonaws.com:3000/



### Mesh

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

* Create a file named create-virtual-node-servicebv2.json with the following contents:

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

        aws appmesh create-virtual-node --cli-input-json file://create-virtual-node-servicebv2.json

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


#### Update services
  
After creating your mesh, you need to complete the following tasks:

- Authorize the Envoy proxy that you deploy with each Kubernetes pod to read the configuration of one or more virtual nodes. For more information about how to authorize the proxy, see Proxy authorization.

- Update each of your existing Kubernetes pod specs to use the Envoy proxy.

App Mesh vends the following custom container images that you must add to your Kubernetes pod specifications:

- Specify one of the following App Mesh Envoy container images, depending on which region you want to pull the image from.

- All supported Regions other than me-south-1 and ap-east-1. You can replace us-west-2 with any Region other than me-south-1 and ap-east-1.

        840364872350.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-envoy:v1.12.2.1-prod

- me-south-1 Region:

        772975370895.dkr.ecr.me-south-1.amazonaws.com/aws-appmesh-envoy:v1.12.2.1-prod

- ap-east-1 Region:

        856666278305.dkr.ecr.ap-east-1.amazonaws.com/aws-appmesh-envoy:v1.12.2.1-prod


Envoy uses the configuration defined in the App Mesh control plane to determine where to send your application traffic.

You must use the App Mesh Envoy container image until the Envoy project team merges changes that support App Mesh. For additional details, see the GitHub roadmap issue.

App Mesh proxy route manager – 111345817488.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-proxy-route-manager:v2. The route manager sets up a pod’s network namespace with iptables rules that route ingress and egress traffic through Envoy.

Update each pod specification in your application to include these containers, as shown in the following example. Once updated, deploy the new specifications to update your services and start using App Mesh with your Kubernetes application. 

The following example shows updating the svcNode1 pod specification, that aligns to the scenario. To complete the scenario, you also need to update the svcNode1v2 and svcNode2 pod specifications by changing the values appropriately. 

For your own applications, substitute your mesh name and virtual node name for the APPMESH_VIRTUAL_NODE_NAME value, and add a list of ports that your application listens on for the APPMESH_APP_PORTS value. Substitute the Amazon EC2 instance AWS Region for the AWS_REGION value.

Example Kubernetes pod spec

```yaml
spec:
  containers:
    - name: envoy
      image: 840364872350.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-envoy:v1.12.2.1-prod
      securityContext:
        runAsUser: 1337
      env:
        - name: "APPMESH_VIRTUAL_NODE_NAME"
          value: "mesh/pgMesh/virtualNode/svcNode1"
        - name: "ENVOY_LOG_LEVEL"
          value: "info"
        - name: "AWS_REGION"
          value: "aws_region_name"
  initContainers:
    - name: proxyinit
      image: 111345817488.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-proxy-route-manager:v2
      securityContext:
        capabilities:
          add: 
            - NET_ADMIN
      env:
        - name: "APPMESH_START_ENABLED"
          value: "1"
        - name: "APPMESH_IGNORE_UID"
          value: "1337"
        - name: "APPMESH_ENVOY_INGRESS_PORT"
          value: "15000"
        - name: "APPMESH_ENVOY_EGRESS_PORT"
          value: "15001"
        - name: "APPMESH_APP_PORTS"
          value: "application_port_list"
        - name: "APPMESH_EGRESS_IGNORED_IP"
          value: "169.254.169.254"
        - name: "APPMESH_EGRESS_IGNORED_PORTS"
          value: "22"
```