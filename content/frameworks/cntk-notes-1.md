+++
date = "2016-12-07T16:56:21+11:00"
title = "CNTK Note - 1"
description="AI Framework from Microsoft"

+++


##  Prerequisites

* You are using Windows 7 or higher version
* You are using Anaconda to setup the environment

##   Create CNTK virtual environment

###   use follow command to remove existing virtual environment
    
        conda remove -n cntk --all
        conda create -n cntk 

###  Activate virtual environment and install CNTK

    activate cntk
    pip install https://cntk.ai/PythonWheel/CPU-Only/cntk-2.0rc3-cp36-cp36m-win_amd64.whl

##  Test CNTK

    python
    >>> import cntk
    >>> cntk.__version__
    '2.0rc3'



