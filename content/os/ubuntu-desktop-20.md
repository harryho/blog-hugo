+++
date = "2021-12-20T14:59:31+11:00"
title = "Ubuntu Desktop 20 LTS note"
description = "Post-installation for Ubuntu 20 desktop"
+++

## Ubuntu 20.04.3 LTS (Focal Fossa)

Ubuntu is the world’s most popular open-source desktop operating system. Ubuntu 20.04 LTS is an enterprise-grade, secure, cost-effective operating system for organisations and home users. 



### Zsh Prezto

#### Install Prezto


```sh
clear
sudo apt-get install -y git
sudo apt-get update && sudo apt-get install -y zsh
# Get prezto
git clone --recursive https://github.com/sorin-ionescu/prezto.git ~/.zprezto

# Backup zsh config if it exists
if [ -f ~/.zshrc ];
    then
        mv ~/.zshrc ~/.zshrc.backup
fi

# Create links to zsh config files
ln -s ~/.zprezto/runcoms/zlogin ~/.zlogin
ln -s ~/.zprezto/runcoms/zlogout ~/.zlogout
ln -s ~/.zprezto/runcoms/zpreztorc ~/.zpreztorc
ln -s ~/.zprezto/runcoms/zprofile ~/.zprofile
ln -s ~/.zprezto/runcoms/zshenv ~/.zshenv
ln -s ~/.zprezto/runcoms/zshrc ~/.zshrc

```

#### Change theme & module

* Update the theme 'sorin' to 'steeef' in .zpreztorc
* Add following plugins

```
zstyle ':prezto:load' pmodule \
  'environment' \
  'terminal' \
  'editor' \
  'history' \
  'directory' \
  'spectrum' \
  'utility' \
  'completion' \
  'git' \
  'syntax-highlighting' \
  'history-substring-search' \
  'prompt'

```

#### Change shell to Zsh

```sh
chsh -s $(Which zsh)
source ~/.zshrc
```

### Install docker


#### Set up the repository


```sh
 sudo apt-get update

 sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
```

#### Add Docker’s official GPG key:

```sh
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor \
    -o /usr/share/keyrings/docker-archive-keyring.gpg

# Use the following command to set up the stable repository. 
# To add the nightly or test repository, add the word nightly or test (or both) 
# after the word stable in the commands below. Learn about nightly and test channels.

echo \
"deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```



#### Install Docker Engine

```sh
# Update the apt package index, and install the latest version of 
# Docker Engine and containerd, or go to the next step to install 
# a specific version:

 sudo apt-get update

 sudo apt-get install docker-ce docker-ce-cli containerd.io
```


#### Post installation of Docker

```sh

# Create the docker group.
sudo groupadd docker

# Add your user to the docker group.
sudo usermod -aG docker $USER

# Log out and log back in so that your group membership is re-evaluated.

# On Linux, you can also run the following command to activate the changes to groups:

newgrp docker

# Test the docker
docker ps
```