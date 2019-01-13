+++
title = "C# Note - 3"
hidden="true"
+++


## C# Thread & Task

> From Net 4.0, .Net applicatoin introduced Parallel Framework Extensions (PFx), along the way it delivered an entirely new model for async processing in
.NET. In .NET 4.0 the thread pool queue was redesigned with the new requirements of PFx in mind. Instead of using a simple linked list, the queue was built with arrays of work items with the arrays connected into a linked list.

> With the release of .NET 4.0, Microsoft introduced yet another API for building asynchronous applications: the Task Parallel Library (TPL).
 

### Old way of C# thread

* Early days of .NET applications you will see many following sample codes.

## Lock keyword

```cs
    lock(stateGuard)
    {
        cash += amount;
        receivables -= amount;
    }
```

## Abort 

`Thread.Abort` should be avoid at all costs. Its behavior is much safer and predictable since .NET 2.0, but there are still some pretty serious pitfalls with it. The other ways to stop thread.
* Use `Thread.Interrupt`
* Use a stopping flag
* Use `WaitHandle` events

### New way of C# Async & Thread

## Monitor 

*  Monitor is no different from lock but the monitor class provides more control over the synchronization of various threads trying to access the same lock of code.

* Lock and Monitor sample

```cs
class Program  
{  
    static readonly object _object = new object();  

    public static void PrintNumbers()  
    {  
        Monitor.Enter(_object);  
        try  
        {  
            for (int i = 0; i < 5; i++)  
            {  
                Thread.Sleep(100);  
                Console.Write(i + ",");  
            }  
            Console.WriteLine();  
        }  
        finally  
        {  
            Monitor.Exit(_object);  
        }  
    }  

    static void TestLock()  
    {            
        lock (_object)  
        {  
            Thread.Sleep(100);  
            Console.WriteLine(Environment.TickCount);  
        }  
    }  

    static void Main(string[] args)      
    {            
        Thread[] Threads = new Thread[3];  
        for (int i = 0; i < 3; i++)  
        {  
            Threads[i] = new Thread(new ThreadStart(PrintNumbers));  
            Threads[i].Name = "Child " + i;  
        }  
        foreach (Thread t in Threads)  
            t.Start();  

        Console.ReadLine();  
    }  
}  
```

* Once you have a lock on a code region, you can use the Monitor.Wait, Monitor.Pulse, and Monitor.PulseAll methods.

* Lock and monitor are basically used for the same purpose in multithreading, the difference is that only when we want more control over synchronization with multiple threads running for a specific section of code.


### Monitor Semaphore


```cs
public class MonitorSemaphore
{
    private int currentCount;
    private readonly int maxCount;
    private readonly object guard = new object();
    public MonitorSemaphore(int initialCount, int maxCount)
    {
        this.currentCount = initialCount;
        this.maxCount = maxCount;
    }

    public void Enter()
    {
        lock (guard)
        {
            while (currentCount == maxCount)
            {
                Monitor.Wait(guard);
            }
            currentCount++;
        }
    }

    public void Exit()
    {
        lock (guard)
        {
            currentCount--;
            Monitor.Pulse(guard);
        }
    }
    public int CurrentCount{get { return currentCount; }}
}
```

