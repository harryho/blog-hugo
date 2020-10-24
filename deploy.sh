#!/bin/bash

echo "Deploying updates to GitHub ..."

# Go To Public folder
cd public

# Add changes to git.
pwd

git config user.name "harryho"
git config user.email "harry.ho_long@yahoo.com"

git remote show origin

git checkout master

git pull origin master

## back to root of blog-hugo
cd ..

# Build the project. If using a theme, replace by `hugo -t <yourtheme>`
hugo -t docdock --ignoreCache

# Navigate into public
cd public

git add -A

msg="rebuilding site $(date)"
if [ $# -eq 1 ]; then
  msg="$1"
fi

git commit -m "$msg"
# Push source and build repos.
git push origin master

# Come Back
cd ..
