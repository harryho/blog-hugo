#!/bin/bash

if [ ! -z themes/docdock/layouts ]; then
    echo 'docdock does not exist.'
    git submodule add --force https://github.com/harryho/hugo-theme-docdock.git themes/docdock 
    git submodule init .
    git submodule update .
fi

hugo serve -t docdock --watch  --disableFastRender
