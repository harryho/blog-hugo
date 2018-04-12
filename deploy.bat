@echo OFF

echo  Deploying updates to GitHub...

REM Build the project. If using a theme, replace by `hugo -t <yourtheme>`
hugo -t bootie-docs

REM Go To Public folder
cd public

REM Add changes to git.
git config user.name "harryho"
git config user.email "harry.ho_long@yahoo.com"
git add -A

REM Commit changes.
set msg="rebuilding site %date%"
if  NOT "%1"=="" set msg=%1
git commit -m '%msg%'

REM Push source and build repos.
git push origin master

REM Come Back
cd ..

