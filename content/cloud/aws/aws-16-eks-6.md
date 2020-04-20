+++
title = "AWS: EKS - 6"
description = "Labs: GuestBook"
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


