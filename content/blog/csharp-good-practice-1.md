+++
categories = ["blog"]
date = "2014-04-10T14:59:31+11:00"
title = "C# good practice -- Part 1"
draft = true
+++

## Prelude

> *As the title says, this is first article of the series, which cover C# good practice and common mistakes made by many developers. As we known, .net framework and its most important player changes every year. So some good practice will be obsolete or replaced by some advanced solution, because real world keeps change. The most ridiculous thing is some good practice becomes common mistake in some special context. That is why we keep looking for new ways of good practice to support emerging chanlleges.*  

## Performance and debug

* Boxing and unboxing are time-consuming processes that might affect the performance of an application.
The boxing operation can take up to 20 times more time than the assignment operation. The right use of
the boxing operation is important when the performance is the key factor for your application. Let’s find
out how boxing and unboxing operations affect the performance of an application. 

* Use [Windbg](https://developer.microsoft.com/en-us/windows/hardware/windows-driver-kit) 

## 

