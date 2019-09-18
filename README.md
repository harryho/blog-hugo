# blog-hugo

[![Join the chat at https://gitter.im/blog-hugo/Lobby](https://badges.gitter.im/blog-hugo/Lobby.svg)](https://gitter.im/blog-hugo/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
Hugo blog repo

* Use any `markdown` editor to edit the blog or post under the folder `content`
* Commit the markdown files to repo `blog-hugo` by running script the commit.bat. The commit.bat will only commit the markdown files. To commit other files please use git bash or git extension.
* Deploy posts to harryho.github.io by running script deploy.bat. After that, you can check it via browser.


The theme Docdock has been forked as customized repository 

```
git submodule add --force https://github.com/harryho/hugo-theme-docdock.git themes/docdock 

```

Update submodule 

```
git submodule init
git submodule update

cd themes/<your_theme>
git checkout hw  # only hw branch should be used
git pull

```


Run the site locally

```
hugo server

# with script 
# Linux / Mac
./start.sh

# Windows
./start.cmd
```


Edit the site with draft

```bash
# Run the start script with param d
./start.sh  d
```
