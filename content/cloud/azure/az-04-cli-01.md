+++
title = "Azure: CLI - 1"
weight = 4
description="Introduction of Azure CLI"
+++


## Azure Cli 

The Azure Command-Line Interface (CLI) is a cross-platform command-line tool to connect to Azure and execute administrative commands on Azure resources. It allows the execution of commands through a terminal using interactive command-line prompts or a script.


### Install Azure CLI on macOS

* Mac with Intel CPU 

```sh
brew update && brew install azure-cli
```


### Install Azure Cli with Docker

* Use `docker run`

```sh
docker run -it mcr.microsoft.com/azure-cli
```

* Run with SSH key

```sh
docker run -it -v ${HOME}/.ssh:/root/.ssh mcr.microsoft.com/azure-cli
```




### Get Started


* Sign In

```
az login
```

* Show App Services

```
az appservice plan list --query-examples
```


* Show Web Apps

```
az webapp  list --query-examples
```


