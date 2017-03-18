@echo OFF

echo  commit blog content to github

REM Add changes under folder content to git.
git add content/*
git add static/img/*

REM Commit changes.
set msg="commit blog content %date%"
if  NOT "%1"=="" set msg=%1
git commit -m '%msg%'

REM Push source and build repos.
git push origin master

