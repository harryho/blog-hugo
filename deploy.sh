#!/bin/bash

echo "Deploying updates to GitHub ..."

# Build the project. If using a theme, replace by `hugo -t <yourtheme>`
hugo -t docdock --ignoreCache

# Go To Public folder
cd public

# Add changes to git.
pwd

git config user.name "harryho"
git config user.email "harry.ho_long@yahoo.com"

git remote show origin

git checkout master

git pull origin master

git add -A

# Commit changes.
# set msg="rebuilding site `date`"
# if  NOT "%1"=="" set msg=%1

msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi

git commit -m "$msg"
# Push source and build repos.
git push origin master

# Come Back
cd ..
