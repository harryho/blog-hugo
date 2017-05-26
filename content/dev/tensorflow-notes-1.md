+++
categories = ["dev"]
date = "2016-12-07T16:56:21+11:00"
title = "Tensorflow Notes, Part 1"
tags = ["ai"]
+++

### Assumption

* You are using Ubuntu 16
* Your machine has Nvidia GPU card installed
* If you are using Ubuntu 14, the installation of CUDA and cuDNN will be a bit different. Please check Google's instructions. 


### Install python3 and pip3

* [Please find instructions here](https://harryho.github.io/os/ubuntu-server-14)

### Install virtualenv via pip3 
    
    pip3 install virtualenv

### Create two tensorflow virtualenvs.

    mkdir ~/.envs
    virtualenv --system-site-packages ~/.envs/tf  # CPU only
    virtualenv --system-site-packages ~/.envs/tfgpu   # GPU enabled

### Install tensorflow for different virtualenvs

    source ~/.envs/tf/bin/activate
    source ~/.envs/tfgpu/bin/activate
    
    pip3 install tensorflow # CPU only
    pip3 install tensorflow-gpu # GPU enabled

### Install CUDA and cuDNN for tensorflow-gpu


#### Use following command to check you GPU information

    lspic -nn | grep '\[03' ]
    lshw -numeric -C display

    # GPU info sample
    # NVIDIA Corporation GM107M [GeForce GTX 850M]

#### Download and install Nvidia driver based on above GPU info http://www.geforce.com/drivers
    chmod +x NVIDIA-Linux-xxxx.run
    sudo ./NVIDIA-Linux-xxxx.run   

#### Download and install  CUDA from NVIDIA https://developer.nvidia.com/cuda-downloads

    sudo dpkg -i cuda-repo-xxxxx.deb
    sudo apt-get udpate
    sudo apt-get install cuda

#### Setup CUDA_HOME 

    # CUDA
    export CUDA_HOME=/usr/local/cuda-8.0 
    export LD_LIBRARY_PATH=${CUDA_HOME}/lib64 


#### Download and install cuDNN for CUDA https://developer.nvidia.com/cudnn

    # extra the cuDNN tar ball
    tar -xvf cudnn-8.0
    cd cuda 
    sudo cp lib64/* /usr/local/cuda-8.0/lib64
    sudo cp include/* /usr/local/cuda-8.0/include

### Use sample code to test Tensorflow

#### Save code below to file test.py 

```python
import numpy as np
import tensorflow as tf

# Model parameters
W = tf.Variable([.3], tf.float32)
b = tf.Variable([-.3], tf.float32)
# Model input and output
x = tf.placeholder(tf.float32)
linear_model = W * x + b
y = tf.placeholder(tf.float32)
# loss
loss = tf.reduce_sum(tf.square(linear_model - y)) # sum of the squares
# optimizer
optimizer = tf.train.GradientDescentOptimizer(0.01)
train = optimizer.minimize(loss)
# training data
x_train = [1,2,3,4]
y_train = [0,-1,-2,-3]
# training loop
init = tf.global_variables_initializer()
sess = tf.Session()
sess.run(init) # reset values to wrong
for i in range(1000):
sess.run(train, {x:x_train, y:y_train})

# evaluate training accuracy
curr_W, curr_b, curr_loss  = sess.run([W, b, loss], {x:x_train, y:y_train})
print("W: %s b: %s loss: %s"%(curr_W, curr_b, curr_loss))
```

#### Test with tensorflow-gpu (GPU enabled)

    source ~/.envs/tfgpu/bin/activate
    python3 test.py

    # You will probably see the result as follow 
    # ....
    # name: GeForce GTX 850M
    # major: 5 minor: 0 memoryClockRate (GHz) 0.9015
    # pciBusID 0000:0a:00.0
    # Total memory: 3.95GiB
    # Free memory: 3.58GiB
    # 2017-05-25 11:25:59.640621: I tensorflow/core/common_runtime/gpu/gpu_device.cc:908] DMA: 0 
    # 2017-05-25 11:25:59.640626: I tensorflow/core/common_runtime/gpu/gpu_device.cc:918] 0:   Y 
    # 2017-05-25 11:25:59.640640: I tensorflow/core/common_runtime/gpu/gpu_device.cc:977] Creating TensorFlow device (/gpu:0) -> (device: 0, name: GeForce GTX 850M, pci # bus id: 0000:0a:00.0)
    # W: [-0.9999969] b: [ 0.99999082] loss: 5.69997e-11










