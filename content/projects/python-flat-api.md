+++

date = "2017-08-03T16:56:21+11:00"
title = "FlatApi - Restful API for python dev"

description = "FlatApi is a zero coding and zero configuration restful API server inspired by Json-Server and Eve"
+++

## Summary

FlatApi is a **zero coding** and **zero configuration** restful API server inspired by Json-Server_ and Eve_. It is designed to be used as fake restful api for development, especially for people want to use Python stack. Setup process is **less than 10 seconds**. 


## FlatApi

- **Zero coding and configuration to setup Restful API** FlatApi is designed to use without coding and configuration by default. You just need one config to setup all endpoints you need, then you can use it immediately. 

- **Flask based web server** FlatApi is built on the top of _Flask

- **Json flat file database** FlatApi uses FlatApi_ to manage the Json flat file database. FlatApi is a document oriented database. 

- **Caching memory storage availble** FlatApi supports caching momery storage after version 4.0.0. 


### Install Package


```bash
$ pip uninstall flatapi
$ pip install --no-cache-dir flatapi
```

### Quick Start


- Launch FlatApi without configuration

```bash
## Start the FlatApi - Sample 1 
$ python3 /<path_to_package>/flatapi -S MEMORY -G NO
## Start the FlatApi - Sample 2
$ python3 /<path_to_package>/flatapi --storage MEMORY -cfgfile NO
## Start the FlatApi with prefix - Sample 3
$ python3 /<path_to_package>/flatapi --storage memory -cfgfile no -X api



\(^_^)/ Hi

Loading  is done.

There is no config file found. Flat Api uses internal configuration.

Resource :
/<string:doc> -- The doc is the collection name
                    you want to post or put the object.
/<string:doc>/<int:id> --The id is the unique id for query or delete.

Database: Memory

* Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
```

- Test api via postman 

    It would be a much more handy and easy way to play around with the API immediately.

```bash
GET /posts       --> Get all posts
POST /posts      --> Add new post
PUT /posts/1     --> Update existing post which id is 1
DELETE /posts/1  --> Delete a post which id is 1
DELETE /posts    --> Delete all posts
```

- Test api via curl 

```bash
## Add a new post
$ curl -d "{\"text\":\"post 1\",\"author\":\"harry\"}" -H "Content-Type: application/json" -X POST http://localhost:5000/posts
{"author": "harry", "text": "post 1", "id": 1}

## Get post by Id
$ curl -X GET http://localhost:5000/posts/1
{"author": "harry", "text": "post 1", "id": 1}

## Get all posts
$ curl -X GET http://localhost:5000/posts
[{"author": "harry", "text": "post 1", "id": 1}]

## Update  the post
$ curl -d "{\"text\":\"post updated\",\"author\":\"harry\"}" -H "Content-Type: application/json" -X PUT http://localhost:5000/posts/1
[{"author": "harry", "text": "post updated", "id": 1}]

## Delete 
$ curl -X DELETE http://localhost:5000/posts 
```
  
### Browse [Repository](https://github.com/harryho/flat-api.git)