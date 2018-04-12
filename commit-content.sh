#!/bin/bash

echo "commit blog content to github"

git config user.name "Harry Ho"
git config user.email "harry.ho_long@yahoo.com"

# Add changes under folder content to git.
git add content/*
git add static/img/*

# Commit changes.
msg="commit blog content `date`"
# if  NOT "%1"=="" set msg=%1
if [ $# -eq 1 ] 
    then msg=$1
fi

git commit -m "$msg"

# Push source and build repos.
git push origin master