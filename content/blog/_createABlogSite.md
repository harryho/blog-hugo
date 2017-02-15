+++
tags =  ["hugo"]
categories = ["blog"]
date = "2017-02-15T14:59:31+11:00"
title = "Create a blog site on GitHub Pages within Windows environment"

+++


After I setup a blog site with **Hugo** on my ubuntu machine, I decided to use it to create a blog to GitHub pages on my windows machine, since I use Windows machine as workstation. If you use Unix-* system, I pretty sure you can follow the [Hugo Quick Start](https://gohugo.io/overview/quickstart/) and [Hosting on GitHub Pages](https://gohugo.io/tutorials/github-pages-blog) to create a blog to GitHub pages within 5 mins, but when I try to do the same on Windows machine, it took me over 30 mins. I hope this blog can help someone want to do sth similar within Windows environment. 


## Assumptions

* You already have **Hugo** on your computer. If not, please follow the instruction to [install hugo on Windows](https://gohugo.io/tutorials/installing-on-windows).
* You will have two repositories `blog-hugo` and `<username>.github.io` repositories to hold your hugo content and blog site respectively.
* The `blog-hugo` repository will host actual Hugo’s blog content.
* `<username>.github.io repository` repository will host the static website.
* Your hugo blog folder will be "C:\git\blog-hugo" in this example.
* Your blog site will finally sit in C driver and map to repositories as follow

    ```
    C:\>
    |--git 
        |--blog-hugo (https://github.com/<yourname>/blog-hugo.git
            |--archetypes
            |--content
            |--data
            |--layouts
            |--public (https://github.com/<yourname>/<yourname>.github.io.git)
            |--themes
            |--

    ```

## Create a blog site on GitHub pages

* Create on GitHub blog-hugo and <username>.github.io repositories  repository via GitHub website
* Clone blog-hugo via Windows command prompt

    ```
    c:\>
    c:\>cd git
    c:\git>git clone <<your-project>-hugo-url> && cd <your-project>-hugo
    ```

* Create hugo site and setup the theme you like

    ```
    C:\git>hugo new site blog-hugo
    C:\git>hugo server -t <yourtheme> -D
    ```

* Clean up the `public` folder 
* Set submodule inside the blog-hugo and map to folder `public`

    ```
    C:\>cd git
    C:\git>rm -rf public
    C:\git>git submodule add -b master https://github.com/<username>/<username>.github.io.git public
    ```

* Deploy the blog site to GitHub page with the script `deploy.bat`.
* `deploy.bat "Your optional commit message"` will commit the changes to `<username>.github.io`.
* You might want to commit the changes to `blog-hugo` repository. 

    ```
    @echo OFF

    echo  Deploying updates to GitHub...

    REM Build the project. 
    hugo -t <yourtheme> -D

    REM Go To Public folder
    cd public

    REM Add changes to git.
    git add -A

    REM Commit changes.
    set msg="rebuilding site %date%"
    if  NOT "%1"==""    set msg=%1
    git commit -m '%msg%'

    REM Push source and build repos.
    git push origin master

    REM Come Back
    cd ..
    ```
