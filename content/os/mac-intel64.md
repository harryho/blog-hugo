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


## Install Zsh Prezto

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


## Install NVM

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

### Install visual studio code


```
brew cask install visual-studio-code

```

### Install python 3

```
brew install python@3.x
```


### Install MySql

```
brew install mysql
```



### Install JDK

* Install different version of JDK

```
brew tap AdoptOpenJDK/openjdk

brew cask install <jdk_version>

```

Java  Version |	JDK	|JRE
------|--------|----------
OpenJDK8 with Hotspot JVM	 | adoptopenjdk8	 | adoptopenjdk8-jre
OpenJDK8 with OpenJ9 JVM	 | adoptopenjdk8-openj9	 | adoptopenjdk8-openj9-jre
OpenJDK8 with OpenJ9 JVM, large heap*	 | adoptopenjdk8-openj9-large	 | adoptopenjdk8-openj9-jre-large
OpenJDK9 with Hotspot JVM	 | adoptopenjdk9	n/a
OpenJDK10 with Hotspot JVM	 | adoptopenjdk10	n/a
OpenJDK11 with Hotspot JVM	 | adoptopenjdk11	 | adoptopenjdk11-jre
OpenJDK11 with OpenJ9 JVM	 | adoptopenjdk11-openj9	 | adoptopenjdk11-openj9-jre
OpenJDK11 with OpenJ9 JVM, large heap*	 | adoptopenjdk11-openj9-large	 | adoptopenjdk11-openj9-jre-large
OpenJDK12 with Hotspot JVM	 | adoptopenjdk12	 | adoptopenjdk12-jre
OpenJDK12 with OpenJ9 JVM	 | adoptopenjdk12-openj9	 | adoptopenjdk12-openj9-jre
OpenJDK12 with OpenJ9 JVM, large heap*	 | adoptopenjdk12-openj9-large	 | adoptopenjdk12-openj9-jre-large
OpenJDK13 with Hotspot JVM	 | adoptopenjdk13	 | adoptopenjdk13-jre
OpenJDK13 with OpenJ9 JVM	 | adoptopenjdk13-openj9	 | adoptopenjdk13-openj9-jre
OpenJDK13 with OpenJ9 JVM, large heap*	 | adoptopenjdk13-openj9-large	 | adoptopenjdk13-openj9-jre-large

* Setup the profile

```
## java home variable to profile
echo "export JAVA_8_HOME=$java8_home">>$HOME/.profile
echo "export JAVA_11_HOME=$java11_home">>$HOME/.profile


java8_home=/usr/lib/jvm/adoptopenjdk-8-hotspot-amd64
java11_home=/usr/lib/jvm/adoptopenjdk-11-hotspot-amd64

echo "# Java 8 & 11">>$HOME/.bash_aliases
echo "alias java8='$java8'">>$HOME/.bash_aliases
echo "alias java11='$java11'">>$HOME/.bash_aliases
source $HOME/.bash_aliases


## Switch JDK 
java8
java -version
java11 
java -version
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


