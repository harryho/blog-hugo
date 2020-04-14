+++
date = "2018-07-04T14:59:31+11:00"
title = "Ubuntu Desktop 18 LTS note"
description = "Post-installation for Ubuntu 18 desktop"
+++

### Prelude

> This article is mainly to record the stuff to do post Ubuntu Desktop 18.04 Installation.


### Purpose 

All actions post installation is to make the Ubuntu Desktop a wonderful toolkit for developer. 

### Prerequisite

* Install all essentials

```bash
sudo apt install -y git curl
sudo apt-get install -y apt-transport-https ca-certificates      gnupg-agent
sudo apt install -y software-properties-common 
```

### Install & Setup Zsh

* Install Zsh

```bash
sudo apt install -y zsh
```

* Setup Zsh

> I prefer [Prezto Zsh](https://github.com/sorin-ionescu/prezto.git), which is the minimal version of Oh-My-Zsh. In my opinion, Oh-My-Zsh is kind of slow and sort of overblown.


- Create a script `setup_prezto.zsh`
- Save the code below to the scirpt and run `zsh ./setup_prezto.sh && source ~/.zshrc `

```bash

git clone --recursive https://github.com/sorin-ionescu/prezto.git \
    "${ZDOTDIR:-$HOME}/.zprezto"

setopt EXTENDED_GLOB
for rcfile in "${ZDOTDIR:-$HOME}"/.zprezto/runcoms/^README.md(.N); do
    ln -s "$rcfile" "${ZDOTDIR:-$HOME}/.${rcfile:t}"
done


if [ -d ~/.zprezto ];then
    cp  ubt18/z* ~/.zprezto/runcoms/
fi

[[ ! -d ~/.zprezto-contrib ]] && mkdir -p ~/.zprezto-contrib;

sudo -R $UsER:$USER .zprezto
sudo -R $UsER:$USER .zprezto-contrib

chsh -s /bin/zsh

echo "Setup finished."
echo "Please reboot or restart terminal."

```

### Install & Setup Vim with useful plugins

* Install Vim

```
sudo apt install -y vim
```

* Setup Vim with some amazing plugins

> There are many different vim plugins available online. There is a plance callded [Vim Awesome](https://vimawesome.com/), which you can find anything you want. 

> My favorite option is the [Junegunn's Vim plugins](https://github.com/junegunn/vim-plug). I like the simplicity of this solution.

- Prepare the vimrc as below 

```
set nu

colorscheme delek

:imap jj <Esc>
:imap jk <Esc>
:imap kj <Esc>
:imap ii <Esc>

" Specify a directory for plugins
" - For Neovim: ~/.local/share/nvim/plugged
" - Avoid using standard Vim directory names like 'plugin'
call plug#begin('~/.vim/plugged')

" Make sure you use single quotes

" Shorthand notation; fetches https://github.com/junegunn/vim-easy-align
Plug 'junegunn/vim-easy-align'

" Any valid git URL is allowed
Plug 'https://github.com/junegunn/vim-github-dashboard.git'

" Multiple Plug commands can be written in a single line using | separators
Plug 'SirVer/ultisnips' | Plug 'honza/vim-snippets'

" On-demand loading
Plug 'scrooloose/nerdtree', { 'on':  'NERDTreeToggle' }
Plug 'tpope/vim-fireplace', { 'for': 'clojure' }

" Using a non-master branch
Plug 'rdnetto/YCM-Generator', { 'branch': 'stable' }

" Using a tagged release; wildcard allowed (requires git 1.9.2 or above)
Plug 'fatih/vim-go', { 'tag': '*' }

" Plugin options
Plug 'nsf/gocode', { 'tag': 'v.20150303', 'rtp': 'vim' }

" Plugin outside ~/.vim/plugged with post-update hook
Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
Plug 'Shougo/vimproc.vim', { 'do': 'make' }


function! BuildYCM(info)
  " info is a dictionary with 3 fields
  " - name:   name of the plugin
  " - status: 'installed', 'updated', or 'unchanged'
  " - force:  set on PlugInstall! or PlugUpdate!
  if a:info.status == 'installed' || a:info.force
    !./install.py
  endif
endfunction

Plug 'Valloric/YouCompleteMe', { 'do': function('BuildYCM') }

" Plug 'Valloric/YouCompleteMe', { 'do': './install.py' }

Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }

" Unmanaged plugin (manually installed and updated)
Plug '~/my-prototype-plugin'

" Initialize plugin system
call plug#end()
```

### Install MySql

- Here I just install the default MySql 5.7. If you want to install the new version or MariaDB, please check out the official website. 


```bash
# Install MySql server and client
sudo apt install mysql-server mysql-client

# Check if the MySql service active and running
sudo systemctl status mysql.service

# Enable or Restart MySql service
sudo systemctl enable mysql.service
sudo systemctl restart mysql.service

# Create root account
sudo mysql_secure_installation

# After the password for root has been created
sudo mysql -u root -p

# Update the password via mysql comomand pompt
mysql>ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'your_strong_password';

# Create another admin is highly recommended in production env
mysql>GRANT ALL PRIVILEGES ON *.* TO 'admin'@'localhost' IDENTIFIED BY 'your_strong_password';
```

### Other useful tools

- Office
    - I prefer [WPS Office](http://wps-community.org). Rename the template files, because the template names contain Chinese characters and it may cause problem later.
    

- Dictionary
    - Offline Dictionary
    
    
    ```
    ## Install dict client & server (dictd)
    sudo apt install -y dict
    sudo apt install -y dictd

    ## Install dictionary libraries
    sudo apt-get install dict-gcide
    sudo apt-get install dict-wn
    sudo apt-get install dict-devil
    sudo apt-get install dict-moby-thesaurus
    ```

- Meida

    - VLC
    
    ```
    sudo apt install vlc
    ```

    - Spotify

    ```bash
    sudo snap install spotify
    ```














