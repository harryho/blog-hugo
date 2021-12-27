+++
title = "Docker Practices"
description="Useful & practical docker practices"
+++

> Docker is an open platform for developing, shipping, and running applications. Docker provides the ability to package and run an application in a loosely isolated environment called a container. The isolation and security allow you to run many containers simultaneously on a given host. Containers are lightweight because they don’t need the extra load of a hypervisor, but run directly within the host machine’s kernel. This means you can run more containers on a given hardware combination than if you were using virtual machines. You can even run Docker containers within host machines that are actually virtual machines!

> Docker enables you to separate your applications from your infrastructure so you can deliver software quickly. With Docker, you can manage your infrastructure in the same ways you manage your applications. By taking advantage of Docker’s methodologies for shipping, testing, and deploying code quickly, you can significantly reduce the delay between writing code and running it in production.


## Common Use Cases

### Test with busybox image

```
docker run -it --rm busybox echo Hello World
## output 
Hello World
```

### Test with nginx image

Create a test web page with content below

```html
<!DOCTYPE html>
<html >
<head>
   <title>Docker Nginx</title>
</head>
<body>
    <h2>Hello from docker</h2>
</body>
</html>
```

Test the web page with ngnix image

```
docker run -it --rm -d -p 8080:80 --name web -v ~/app:/usr/share/nginx nginx
```

### build 

Build an image from a Dockerfile. The presence of Dockerfile is mandatory. The file name convention is __Dockerfile__ or **your_customized_filename.Dockerfile**

    $ docker build <path_of_workspace>

Build an image with tag

    $ docker build <path_of_workspace> -t <image_tag>

Build an image with specific dockerfile

    $ docker build <path_of_workspace> -f  <path_of_dockerfile>

### tag or untag

    $ docker image tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
    
    $ docker rmi [unwanted_tag]
### push 

    $ docker push <docker_repo>:<image_tag>

### scripting

The build & push can be simplified with  some scripting. Here I recap a script from my other docker repository below.

```bash
DOCKER_REPO=$1

if [ -z "$1" ]; then 
    DOCKER_REPO='harryh00/docker-kits'
fi

echo docker repo: $DOCKER_REPO

############################################################################
## --------------------- Build images  ---------------------------
############################################################################

build_image() {
    FOLDER=$1
    echo "----- Build ${FOLDER} based image -----"
    for fname in $(ls ${FOLDER}); do
        PREFIX=${fname/".Dockerfile"/""}
        echo ":::: Build ${DOCKER_REPO}:${FOLDER}-${PREFIX}"
        docker build ${FOLDER} -f "${FOLDER}/${PREFIX}.Dockerfile" \
            -t "${DOCKER_REPO}:${FOLDER}-${PREFIX}"
    done
}

##############################################################################
## --------------------- Push images to docker hub ---------------------------
##############################################################################

## alpine basealpined image
push_image() {
    FOLDER=$1
    echo "----- Push ${FOLDER} based image -----"
    for fname in $(ls ${FOLDER}); do
        PREFIX=${fname/".Dockerfile"/""}
        echo ":::: Push ${DOCKER_REPO}:${FOLDER}-${PREFIX}"
        docker push ${DOCKER_REPO}:${FOLDER}-${PREFIX}
    done
}

# Main
main() {
    docker login
    docker info
    FOLDERS=(
        alpine
        ubuntu
        centos
    )
    for FOLDER in ${FOLDERS[@]}; do
        build_image ${FOLDER}
        push_image ${FOLDER}
    done
}

main "$@"
```

### 2-step build

For production deployment, usually we just deploy the delivery instead of the full copy of source code. To achieve that, we can the build into 2 steps. The first step is to build the source code. and the second one is to build the deliverable image. 

I recap one docker file from the my repository vue-crm here


```dockerfile
###### Build the App #####
FROM node:10.19 AS node
LABEL author="Harry Ho"
WORKDIR /
COPY . .
RUN npm install
RUN npm run build -- --prod


###### Build the Delivery #####
FROM nginx:alpine
LABEL author="Harry Ho"
WORKDIR /var/cache/nginx
COPY --from=node /dist /usr/share/nginx/html
COPY ./config/nginx.conf /etc/nginx/conf.d/default.conf

```


