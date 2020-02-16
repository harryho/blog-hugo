#!/bin/bash

if [ ! -z themes/docdock/layouts ]; then
    echo 'docdock does not exist.'
    git submodule add --force https://github.com/harryho/hugo-theme-docdock.git themes/docdock 
    git submodule init .
    git submodule update .
fi

DRAFT=$1

if [ ${DRAFT} == 'd' ]; then 
    echo 'start with draft '
    hugo serve -t docdock -D --watch  --disableFastRender
else
    hugo serve -t docdock --watch  --disableFastRender --ignoreCache
fi




