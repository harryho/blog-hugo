+++
title = "AWS: EKS - 3"
description = "Cluster Autoscaler, Horizontal Pod Autoscaler, Vertical Pod Autoscaler"
weight=13
+++


## EKS - Part 3

### Cluster Autoscaler

The Kubernetes Cluster Autoscaler automatically adjusts the number of nodes in your cluster when pods fail to launch due to lack of resources or when nodes in the cluster are underutilized and their pods can be rescheduled onto other nodes in the cluster.

### Strategy of auto     scaling

* Stateful application 
  
If you are running a stateful application across multiple Availability Zones that is backed by Amazon EBS volumes and using the Kubernetes Cluster Autoscaler, you should configure multiple node groups, each scoped to a single Availability Zone.

* Other option

Create a single node group that spans multiple Availability Zones.


#### Single managed node group

Create an Amazon EKS cluster with a single managed node group

    eksctl create cluster --name pg-smng --version 1.15 --managed --asg-access


#### Node group per AZ

Create a cluster with a dedicated managed node group for each Availability Zone

    eksctl create cluster --name pg-ngaz --version 1.15 --without-nodegroup

For each Availability Zone in your cluster, use the following eksctl command to create a node group. 

    eksctl create nodegroup \
            --cluster pg-ngaz \
            --node-zones ap-southeast-2a \
            --name ap-southeast-2a \
            --asg-access \
            --node-type t3.medium \
            --nodes-min 1 --nodes 1 \
            --nodes-max 3 --managed


* Node Group IAM Policy

The Cluster Autoscaler requires the following IAM permissions to make calls to AWS APIs on your behalf. The tool __eksctl__ automatically provides and attaches to your worker node IAM roles, when it creates the node groups.

    {
        "Version": "2012-10-17",
        "Statement": [
                {
                "Action": [
                        "autoscaling:DescribeAutoScalingGroups",
                        "autoscaling:DescribeAutoScalingInstances",
                        "autoscaling:DescribeLaunchConfigurations",
                        "autoscaling:DescribeTags",
                        "autoscaling:SetDesiredCapacity",
                        "autoscaling:TerminateInstanceInAutoScalingGroup",
                        "ec2:DescribeLaunchTemplateVersions"
                ],
                "Resource": "*",
                "Effect": "Allow"
                }
        ]
    }


#### Deploy Autoscaler


Deploy the Cluster Autoscaler to your cluster with the following command.

    kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/master/cluster-autoscaler/cloudprovider/aws/examples/cluster-autoscaler-autodiscover.yaml

Add the cluster-autoscaler.kubernetes.io/safe-to-evict annotation to the deployment with the following command.

    kubectl -n kube-system annotate deployment.apps/cluster-autoscaler cluster-autoscaler.kubernetes.io/safe-to-evict="false"

Edit the Cluster Autoscaler deployment with the following command.

    kubectl -n kube-system edit deployment.apps/cluster-autoscaler  

Edit the cluster-autoscaler container command to replace <YOUR CLUSTER NAME> with your cluster's name, and add the following options.

    --balance-similar-node-groups
    --skip-nodes-with-system-pods=false

Final change will look like below

    spec:
      containers:
      - command:
        - ./cluster-autoscaler
        - --v=4
        - --stderrthreshold=info
        - --cloud-provider=aws
        - --skip-nodes-with-local-storage=false
        - --expander=least-waste
        - --node-group-auto-discovery=asg:tag=k8s.io/cluster-autoscaler/enabled,k8s.io/cluster-autoscaler/<YOUR CLUSTER NAME>
        - --balance-similar-node-groups
        - --skip-nodes-with-system-pods=false

Set the Cluster Autoscaler image tag

    kubectl -n kube-system set image deployment.apps/cluster-autoscaler cluster-autoscaler=asia.gcr.io/k8s-artifacts-prod/autoscaling/cluster-autoscaler:v1.15.6

Log Autoscaler

    kubectl -n kube-system logs -f deployment.apps/cluster-autoscaler


Add scale policy 

    


### Horizontal Pod Autoscaler


The Kubernetes Horizontal Pod Autoscaler automatically scales the number of pods in a deployment, replication controller, or replica set based on that resource's CPU utilization. This can help your applications scale out to meet increased demand or scale in when resources are not needed, thus freeing up your worker nodes for other applications. When you set a target CPU utilization percentage, the Horizontal Pod Autoscaler scales your application in or out to try to meet that target.


The Horizontal Pod Autoscaler is a standard API resource in Kubernetes that simply requires that a metrics source (such as the Kubernetes metrics server) is installed on your Amazon EKS cluster to work. 


Install metrics-server with __curl__ and __jq__

    DOWNLOAD_URL=$(curl -Ls "https://api.github.com/repos/kubernetes-sigs/metrics-server/releases/latest" | jq -r .tarball_url)
    DOWNLOAD_VERSION=$(grep -o '[^/v]*$' <<< $DOWNLOAD_URL)
    curl -Ls $DOWNLOAD_URL -o metrics-server-$DOWNLOAD_VERSION.tar.gz
    mkdir metrics-server-$DOWNLOAD_VERSION
    tar -xzf metrics-server-$DOWNLOAD_VERSION.tar.gz --directory metrics-server-$DOWNLOAD_VERSION --strip-components 1
    kubectl apply -f metrics-server-$DOWNLOAD_VERSION/deploy/1.8+/


#### Horizontal Autoscale Test

* Install httpd pod
  
        kubectl run httpd \
        --generator=run-pod/v1 \
        --image=httpd --requests=cpu=100m \
        --limits=cpu=200m --expose --port=80

* Run benchmark test
  
        kubectl run apache-bench \
            --generator=run-pod/v1 \
            -i --tty --rm --image=httpd \
            -- ab -n 900000 \
            -c 9999 http://httpd.default.svc.cluster.local/


### Vertical Pod Autoscaler

The Kubernetes Vertical Pod Autoscaler automatically adjusts the CPU and memory reservations for your pods to help "right size" your applications. This adjustment can improve cluster resource utilization and free up CPU and memory for other pods. This topic helps you to deploy the Vertical Pod Autoscaler to your cluster and verify that it is working.

### Deploy Vertical Autoscaler

Open a terminal window and navigate to a directory where you would like to download the Vertical Pod Autoscaler source code.

* Clone the kubernetes/autoscaler GitHub repository.

        git clone https://github.com/kubernetes/autoscaler.git

* Change to the vertical-pod-autoscaler directory.

        cd autoscaler/vertical-pod-autoscaler/

* (Optional) If you have already deployed another version of the Vertical Pod Autoscaler, remove it with the following command.

        ./hack/vpa-down.sh

* Deploy the Vertical Pod Autoscaler to your cluster with the following command.

        ./hack/vpa-up.sh

* Check Vertical Autoscaler pods

        kubectl get pods -n kube-system | grep vpa



#### Test Vertical Autoscaler 

* Deploy the hamster.yaml Vertical Pod Autoscaler example with the following command.

        kubectl apply -f examples/hamster.yaml

* Get the pods from the hamster example application.

        kubectl get pods -l app=hamster
        
        # Output:

        hamster-c7d89d6db-rglf5   1/1     Running   0          48s
        hamster-c7d89d6db-znvz5   1/1     Running   0          48s

* Describe one of the pods to view its CPU and memory reservation.

        kubectl describe pod hamster-c7d89d6db-rglf5

* Describe the hamster-vpa resource to view the new recommendation.

        kubectl describe vpa/hamster-vpa

        # Output
        Status:
        Conditions:
            Last Transition Time:  2020-02-11T13:31:48Z
            Status:                True
            Type:                  RecommendationProvided
        Recommendation:
            Container Recommendations:
            Container Name:  hamster
            Lower Bound:
                Cpu:     530m
                Memory:  262144k
            Target:
                Cpu:     587m
                Memory:  262144k
            Uncapped Target:
                Cpu:     587m
                Memory:  262144k
            Upper Bound:
                Cpu:     1
                Memory:  500Mi





