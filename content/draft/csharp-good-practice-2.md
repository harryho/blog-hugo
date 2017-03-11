+++
tags = ["c#"]
categories = ["note"]
date = "2016-04-10T14:59:31+11:00"
title = "C# good practice - Part 1"
draft = true
+++

## Performance and debug

* Boxing and unboxing are time-consuming processes that might affect the performance of an application.
The boxing operation can take up to 20 times more time than the assignment operation. The right use of
the boxing operation is important when the performance is the key factor for your application. Let’s find
out how boxing and unboxing operations affect the performance of an application. 

* Use [Windbg](https://developer.microsoft.com/en-us/windows/hardware/windows-driver-kit) 

## 