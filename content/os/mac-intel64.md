+++


date = "2015-07-12T14:59:31+11:00"
title = "Macbook Notes - Intel X64"
description = "Environment setup"
+++


### Prerequisites

* Mac OS 10.12+
* Intel X64 CPU


### Ownership issue

* If you have Homebrew or other software installed by someone else, you need to change ownership

```
sudo chown -R $(whoami) /usr/local/brew
sudo chown -R $(whoami) /usr/local/etc
sudo chown -R $(whoami) /usr/local/share
sudo chown -R $(whoami) /usr/local/lib

```

### Install Homebrew 

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

## Brew commands

brew update
brew upgrade 
brew list

```
{{% notice info %}}
The cask is not longer a brew command from 2021. The command has changed as below
{{% /notice %}}

```
# Before 2021
brew cask install XXXX

# After 2021
brew install --cask XXXX
```
### Install Zsh Prezto

- Clone Zsh Prezto

```sh
git clone --recursive https://github.com/sorin-ionescu/prezto.git "${ZDOTDIR:-$HOME}/.zprezto"

```

- Remove default `zshrc`

```sh
rm -f ~/.zshrc
```

- Initialize our Prezto configuration files.

```sh
setopt EXTENDED_GLOB
for rcfile in "${ZDOTDIR:-$HOME}"/.zprezto/runcoms/^README.md(.N); do
  ln -s "$rcfile" "${ZDOTDIR:-$HOME}/.${rcfile:t}"
done
```

- Setup Prezto Style

  - Open up ~/.zpreztorc and find where it says:
  - change “sorin” to “steeef.”

```sh
zstyle ':prezto:module:prompt' theme 'steeef'
```
  - Add Some Prezto Modules

```sh
'environment' \
  'terminal' \
  'editor' \
  'history' \
  'directory' \
  'spectrum' \
  'utility' \
  'completion' \
  'prompt' \
  'git' \
  'completion' \
  'syntax-highlighting' \
  'history-substring-search'
```


### Install NVM

```sh
brew install nvm
```


* Add following setting to file `.zshrc`

```sh
# NVM
export NVM_DIR="$HOME/.nvm"
  [ -s "/usr/local/opt/nvm/nvm.sh" ] && \. "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
  [ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && \. "/usr/local/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion
```

### Install python 3

```
brew install python@3.x
```


### Install MySql
{{% notice info %}}
Only the MySql 5.6 or 5.7 doesn't work well on Mac Big Sur or later. If you need old MySql, please consider to se docker to host mysql db.
{{% /notice %}}

```
brew install mysql
```



### Install JDK 


* Tap adoptopenjdk to brew

```sh
brew tap adoptopenjdk/openjdk

brew search openjdk
==> Formulae
openjdk                                     openjdk@11                                  openjdk@8                                   openj9                                      openvdb

==> Casks
adoptopenjdk-jre                     adoptopenjdk11-openj9                adoptopenjdk12-openj9-jre-large      adoptopenjdk14                       adoptopenjdk15-openj9                adoptopenjdk8
adoptopenjdk-openj9                  adoptopenjdk11-openj9-jre            adoptopenjdk12-openj9-large          adoptopenjdk14-jre                   adoptopenjdk15-openj9-jre            adoptopenjdk8-jre
adoptopenjdk-openj9-jre              adoptopenjdk11-openj9-jre-large      adoptopenjdk13                       adoptopenjdk14-openj9                adoptopenjdk15-openj9-jre-large      adoptopenjdk8-openj9
adoptopenjdk-openj9-jre-large        adoptopenjdk11-openj9-large          adoptopenjdk13-jre                   adoptopenjdk14-openj9-jre            adoptopenjdk15-openj9-large          adoptopenjdk8-openj9-jre
adoptopenjdk-openj9-large            adoptopenjdk12                       adoptopenjdk13-openj9                adoptopenjdk14-openj9-jre-large      adoptopenjdk16                       adoptopenjdk8-openj9-jre-large
adoptopenjdk10                       adoptopenjdk12-jre                   adoptopenjdk13-openj9-jre            adoptopenjdk14-openj9-large          adoptopenjdk16-jre                   adoptopenjdk8-openj9-large
adoptopenjdk11                       adoptopenjdk12-openj9                adoptopenjdk13-openj9-jre-large      adoptopenjdk15                       adoptopenjdk16-openj9                adoptopenjdk9
adoptopenjdk11-jre                   adoptopenjdk12-openj9-jre            adoptopenjdk13-openj9-large          adoptopenjdk15-jre                   adoptopenjdk16-openj9-jre
```

* Install multiple versions of JDK 

```
brew install --cask adoptopenjdk # Latest version is 16
brew install --cask adoptopenjdk8
brew install --cask adoptopenjdk11
```

* LTS version info

Java Version | First Release 		| Next Release 	| End of Availability 
------|--------|-----------|-----
Java 8 (LTS) 	 | Mar 2014 	|	jdk8u302 20th Jul 2021 |	At Least May 2026 
Java 11 (LTS)  |September 2018 | 	jdk-11.0.12 20th Jul 2021 	|At Least Oct 2024
Java 17 (LTS) 	|Sep 2021 	| jdk-17 14th Sep 2021 |	TBC


* Inspect all available JDKs

```
ls  /Library/Java/JavaVirtualMachines
adoptopenjdk-11.jdk adoptopenjdk-16.jdk adoptopenjdk-8.jdk

```

#### Switch Java version with alias
- Add following lines to file .zshrc

```sh
# Java
export JAVA_8_HOME=$(/usr/libexec/java_home -v1.8)
export JAVA_11_HOME=$(/usr/libexec/java_home -v11)
export JAVA_17_HOME=$(/usr/libexec/java_home -v16)

alias java8='export JAVA_HOME=$JAVA_8_HOME'
alias java11='export JAVA_HOME=$JAVA_11_HOME'
alias java17='export JAVA_HOME=$JAVA_16_HOME'

```

- Switch JDK

```sh
source ~/.zshrc

java8
java -version
java11 
java -version
```


#### Switch Java version with function

- Add following lines to file .zshrc

```sh
jdk() {
      if [[ ! -z $1 ]]; then
        version=$1
        unset JAVA_HOME;
        export JAVA_HOME=$(/usr/libexec/java_home -v"$version");
        java -version
      else
        echo Argument version is required. e.g. 1.8, 11, 16
        echo Example: jdk 1.8  or jdk 11
      fi
}
```

- Switch JDK

```sh
source ~/.zshrc

jdk
Argument version is required. e.g. 1.8, 11, 16
Example: jdk 1.8 or jdk 11

jdk 1.8

jdk 11

```


### Install Vim plugins


```
[[ ! -d ~/.vim/autoload ]] && mkdir -p ~/.vim/autoload
[[ ! -d ~/.vim/plugged ]] && mkdir -p ~/.vim/plugged

curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

cp ./.vimrc $HOME/

echo 'Launch vi and install plugins'
############################################################
# launch vim & install plugins by typing :PlugInstall
vi


```

### Install KubeCtl

```
brew unlink kubernetes-cli
brew install https://raw.githubusercontent.com/Homebrew/homebrew-core/f25e36259eaa8bcf9b9add2c599aa6d8b15f437b/Formula/kubernetes-cli.rb

```

  
### Install Hugo

```
## Version: hugo_extended_0.54.0_macOS-64bit
tar -xvzf ~/Downloads/hugo_X.Y_osx-64bit.tgz
cp hugo_X.Y_osx-64bit.tgz
```


### Install AWS Cli 2.x

```
curl "https://awscli.amazonaws.com/AWSCLIV2-2.0.30.pkg" -o "AWSCLIV2.pkg"
sudo installer -pkg AWSCLIV2.pkg -target /
```


### Install Docker Desktop for mac

- Download & Install [Docker for Mac](https://docs.docker.com/desktop/mac/install/)

- Setup cli completion for zsh prezto

```
curl -fLo ~/.zprezto/modules/completion/external/src/_docker \
https://raw.githubusercontent.com/docker/cli/master/contrib/completion/zsh/_docker
```

- Add following line to zshrc to enable completion

```
autoload -Uz compinit; compinit
```


### Use Docker to launch databases


- Launch MySql 6.6 with docker


```
docker run -d --name mysql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=password \
mysql:5.6.51

```


- Access the MySql via bash

```
docker exec -it mysql bash

# After the bash into the docker container
mysql -u root -p
```

- Access the MySql via MySql client

```
mysql -h 0.0.0.0 \
      -u root  -p 
```

