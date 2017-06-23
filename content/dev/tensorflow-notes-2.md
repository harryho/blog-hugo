+++
categories = ["dev"]
date = "2016-12-07T16:56:21+11:00"
title = "Tensorflow Notes, Part 2"
tags = ["ai"]
+++

### Assumption

* You are using Windows 7 or higher version
* You are using Anaconda to setup the environment


### Install Anaconda 

* [Please download Anaconda from the official site](https://www.continuum.io/downloads)
  
    

### Create tensorflow virtualenv with python 3.5
* Anaconda uses python 3.6 by default. Tensorflow only supports python 3.5. 

    cd /path/to/envs
    conda create -n tensorflow
    

### Install tensorflow 

    activate tensorflow
    # For CPU
    pip install --ignore-installed --upgrade https://storage.googleapis.com/tensorflow/windows/cpu/tensorflow-1.1.0-cp35-cp35m-win_amd64.whl

    # Or for GPU
    pip install --ignore-installed --upgrade https://storage.googleapis.com/tensorflow/windows/gpu/tensorflow_gpu-1.1.0-cp35-cp35m-win_amd64.whl


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

    activate tensorflow
    cd /ws/python/tf
    python3 test.py

    # You will probably see the result as follow 
    # ....
    # name: GeForce GTX 850M
    # major: 5 minor: 0 memoryClockRate (GHz) 0.9015
    # pciBusID 0000:0a:00.0
    # Total memory: 3.95GiB
    # Free memory: 3.58GiB
    # 2017-04-25 10:25:59.640621: I tensorflow/core/common_runtime/gpu/gpu_device.cc:908] DMA: 0 
    # 2017-04-25 10:25:59.640626: I tensorflow/core/common_runtime/gpu/gpu_device.cc:918] 0:   Y 
    # 2017-04-25 10:25:59.640640: I tensorflow/core/common_runtime/gpu/gpu_device.cc:977] Creating TensorFlow device (/gpu:0) -> (device: 0, name: GeForce GTX 850M, pci # bus id: 0000:0a:00.0)
    # W: [-0.9999969] b: [ 0.99999082] loss: 5.69997e-11










