+++
tags =  ["linux","ubuntu","centos"]
categories = ["note"]
date = "2016-01-10T14:59:31+11:00"
title = "Ubuntu 14 -- desktop, extra tools"
draft = true
+++

## How to install Ubunt 14
* Please find it from [Brief history of Linux](/blog/ubuntu-desktop-14/)

***Install chrome***

```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
```

***General prerequest***

```bash
sed -i "/^# deb .*partner/ s/^# //" /etc/apt/sources.list && apt-get update
sudo apt-get install geany byobu p7zip-full gimp pdfshuffler scribus \
filezilla lftp ubuntu-restricted-extras vlc pyrenamer \
imagemagick hugin darktable skype avidemux
```

***Remove Games***

```bash
sudo apt-get remove aisleriot gnome-mahjongg gnomine gnome-sudoku
```

*** Geany themes***
```bash
cd ~/Downloads
git clone https://github.com/codebrainz/geany-themes.git
mkdir ~/.config/geany/colorschemes
cp ~/Downloads/geany-themes/colorschemes/* ~/.config/geany/colorschemes/
rm -rf ~/Downloads/geany-themes
```

***Cloud***
> from: http://www.webupd8.org/2014/06/install-copycom-client-in-ubuntu-or.html

```bash
sudo add-apt-repository ppa:paolorotolo/copy
sudo apt-get update
sudo apt-get install copy
sudo /opt/copy-client/CopyAgent -installOverlay
nautilus -q
copy
```
 
***Data processing***

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys E084DAB9
sudo add-apt-repository 'deb http://star-www.st-andrews.ac.uk/cran/bin/linux/ubuntu trusty/'
sudo apt-get update
sudo apt-get install spyder python-numpy python-numpy-doc sqlite3 \
python-scipy python-matplotlib python-matplotlib-doc r-base git-core
```

***Don't forget to use your own name and email!***

```bash
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
```

***Maps and GIS software***

```bash
sudo apt-get install python-software-properties
sudo add-apt-repository 'deb http://qgis.org/debian trusty main'
gpg --keyserver keyserver.ubuntu.com --recv DD45F6C3
gpg --export --armor DD45F6C3 | sudo apt-key add -
sudo apt-get update
sudo apt-get install qgis python-qgis qgis-plugin-grass grass-gui grass-doc \
libgdal1-dev libproj-dev gpsbabel
```

***Latex type stuff***

```bash
sudo apt-get install jabref ibus-qt4 texlive texlive-latex-extra \
texlive-humanities texlive-fonts-extra latex-beamer
sudo apt-get -f install
```

***Package download and install (Texmaker and RStudio)***

```bash
wget http://www.xm1math.net/texmaker/texmaker_ubuntu_14.04_4.4.1_amd64.deb
wget http://download1.rstudio.org/rstudio-0.98.1102-amd64.deb
sudo dpkg -i *.deb
sudo rm *.deb
 
sudo apt-get update && apt-get upgrade
sudo apt-get autoremove
```


```bash
sudo nano /etc/update-manager/release-upgrades
```
