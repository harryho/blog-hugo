+++

date = "2017-02-15T14:59:31+11:00"
title = "Create a blog site on GitHub Pages"
description="After I setup a blog site with Hugo on my ubuntu machine, I decided to use it to create a blog to GitHub pages on my windows machine "
+++


> If you use Unix-style system, I recommend you to follow the [Hugo Quick Start](https://gohugo.io/overview/quickstart/) and [Hosting on GitHub Pages](https://gohugo.io/tutorials/github-pages-blog) to create a blog to GitHub pages within 5 mins.

> When I decided to use hugo to create a blog on GitHub pages from my windows machine, it took me over 30 mins. I hope this blog can help someone want to do sth similar within Windows environment.*


### Step 1 - Plan and prepare

#### Prerequisite

* You already have *Hugo* on your computer. If not, please follow the instruction to [install hugo on Windows](https://gohugo.io/tutorials/installing-on-windows).

#### Manage your github repositories

* You will have two repositories **blog-hugo** and `<username>.github.io` repositories to hold your hugo content and blog site respectively.

* The **blog-hugo** repository will host actual Hugo’s blog content.

* `<username>.github.io repository` repository will host the static website.

#### Manage your blog site 

* Your *Hugo* blog folder will be "C:\git\blog-hugo" in this example.

* Your blog site will finally sit in C driver and map to repositories as follow

    ```bash
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


### Step -2: Create a blog site


#### Create github repositories

* Create on GitHub blog-hugo and <username>.github.io repositories  repository via GitHub website

#### Create a bloodily good blog site 

* Clone blog-hugo via Windows command prompt

    ```bat
    c:\>
    c:\>cd git
    c:\git>git clone <<your-project>-hugo-url> && cd <your-project>-hugo
    ```

* Create hugo site and setup the theme you like

    ```bat   
    C:\git>hugo new site blog-hugo
    C:\git>hugo server -t <yourtheme> -D
    ```

#### Setup a sub module for publish

* Clean up the `public` folder 
* Set submodule inside the blog-hugo and map to folder `public`

    ```bat
    C:\>cd git/blog-hugo
    C:\git>del /s /q  /f  public\*
    C:\git>rd /s /q public
    C:\git>git submodule add -b master https://github.com/<username>/<username>.github.io.git public
    ```

### Deploy to Github page

* Deploy the blog site to GitHub page with the script `deploy.bat`.
* `deploy.bat "Your optional commit message"` will commit the changes to `<username>.github.io`. You can use and tailor the script below as your `deploy.bat`

    ```batch
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

* You might want to commit the changes to **blog-hugo** repository. Please don't forget to add `public` into the `.gitignore`.

