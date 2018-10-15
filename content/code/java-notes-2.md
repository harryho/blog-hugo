+++
tags = ["java"]
categories = ["code"]
type   = "java"
date = "2013-07-08T14:59:31+11:00"
title = "Java Notes - Part 2: Currency"
draft = false
+++

# Thread

## Join
* The join method allows one thread to wait for the completion of another. join responds to an interrupt by exiting with an InterruptedException.

* Demo code of thread join

```java
public class JoinDemo implements Runnable {
    private Random rand = new Random(System.currentTimeMillis());

    public void run() {
        //simulate some CPU expensive task
        for (int i = 0; i < 100000000; i++) {
            rand.nextInt();
        }
        System.out.println("[" + Thread.currentThread().getName() + "] finished  .");
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[5];
        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Thread(new JoinDemo(), "joinThread " + i);
            threads[i].start();
        }
        for (int i = 0; i < threads.length; i++) {
            threads[i].join();
        }
        System.out.println("[" + Thread.currentThread().getName() + "] All  -threads have finished.");
    }
}
```

## Common problem of multithred program 

* When there are many threads running, the exact sequence in which all running threads are executed depends next to the thread
configuration like priority also on the available CPU resources and the way the scheduler chooses the next thread to execute.
Although the behavior of the scheduler is completely deterministic, it is hard to predict which threads execute in which moment
at a given point in time. This makes access to shared resources critical as it is hard to predict which thread will be the first thread
that tries to access it.

* Sample code without sync can show you what the problem is. If you run the following sample code, you may get different output from mine here. It is also common Thread safe issue for multiple threads program. 

```java
public class NotSyncCounter implements Runnable {
    private static int counter = 0;

    public void run() {
        while (counter < 10) {
            System.out.println("[" + Thread.currentThread().getName() + "]  -    before: " + counter);
            counter++;
            System.out.println("[" + Thread.currentThread().getName() + "]  -    after: " + counter);
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[5];
        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Thread(new NotSyncCounter(), "  -    thread-" + i);
            threads[i].start();
        }
        for (int i = 0; i < threads.length; i++) {
            threads[i].join();
        }
    }
}


//Possible output:
// [  -    thread-2]  -    before: 0
// [  -    thread-1]  -    before: 0
// [  -    thread-4]  -    before: 0
// [  -    thread-3]  -    before: 0
// [  -    thread-0]  -    before: 0
// [  -    thread-3]  -    after: 4
// [  -    thread-3]  -    before: 5
// [  -    thread-4]  -    after: 3
// [  -    thread-1]  -    after: 2
// [  -    thread-1]  -    before: 6
// [  -    thread-1]  -    after: 7
// [  -    thread-2]  -    after: 1
// [  -    thread-1]  -    before: 7
// [  -    thread-4]  -    before: 6
// [  -    thread-4]  -    after: 9
// [  -    thread-4]  -    before: 9
// [  -    thread-3]  -    after: 6
// [  -    thread-0]  -    after: 5
// [  -    thread-4]  -    after: 10
// [  -    thread-1]  -    after: 8
// [  -    thread-2]  -    before: 7
// [  -    thread-2]  -    after: 11
```

* To solve the problme, there is `synchronized` keyword in Java available for us to handle the multiple threads program. 

* Demo code of `synchronized` to solve the problem on above sample.

```java

public class SyncCounter implements Runnable {
    private static int counter = 0;

    public void run() {
        while (counter < 10) {
            synchronized (SyncCounter.class) {
                System.out.println("[" + Thread.currentThread().getName() + "] - before: " + counter);
                counter++;
                System.out.println("[" + Thread.currentThread().getName() + "] - after: " + counter);
            }
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[5];
        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Thread(new SyncCounter(), "  -    thread-" + i);
            threads[i].start();
        }
        for (int i = 0; i < threads.length; i++) {
            threads[i].join();
        }
    }
}

```




### Deadlock

In general the following requirements for a deadlock can be identified:
* Mutual exclusion: There is a resource which can be accessed only by one thread at any point in time.
* Resource holding: While having locked one resource, the thread tries to acquire another lock on some other exclusive resource.
* No preemption: There is no mechanism, which frees the resource if one threads holds the lock for a specific period of time.
* Circular wait: During runtime a constellation occurs in which two (or more) threads are each waiting on the other thread to free a resource that it has locked.


* Monitor with wait and notify



```java

import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;
public class SyncWaitNotfiyDemo {
    private static final Queue <Integer>queue = new ConcurrentLinkedQueue<Integer>();

    public Integer getNextInt() {
        Integer retVal = null;
        synchronized (queue) {
            try {
                while (queue.isEmpty()) {
                    queue.wait();
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            retVal = queue.poll();
        }
        return retVal;
    }

    public synchronized void putInt(Integer value) {
        synchronized (queue) {
            queue.add(value);
            queue.notify();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        final SyncWaitNotfiyDemo queue = new SyncWaitNotfiyDemo();

        Thread thread1 = new Thread(new Runnable() {
            public void run() {
                for (int i = 0; i < 10; i++) {
                    queue.putInt(i);
                }
            }
        });

        Thread thread2 = new Thread(new Runnable() {
            public void run() {
                for (int i = 0; i < 10; i++) {
                    Integer nextInt = queue.getNextInt();
                    System.out.println("Next int: " + nextInt);
                }
            }
        });

        thread1.start();
        thread2.start();
        thread1.join();
        thread2.join();
    }
}
```

### Useful concurrent collections

#### ConcurrentHashMap

> `ConcurrentHashMap` is undoubtedly most popular collection class introduced in Java 5 and most of us are already using it. It provides a concurrent alternative of Hashtable or Synchronized Map classes with aim to support higher level of concurrency by implementing fined grained locking. Multiple reader can access the Map concurrently  while a portion of Map gets locked for write operation depends upon concurrency level of Map. Also it provides better scalability than there synchronized counter part. Iterator of `ConcurrentHashMap` are fail-safe iterators which doesn't throw ConcurrencModificationException thus eliminates another requirement of locking during iteration which result in further scalability and performance.


#### CopyOnWriteArrayList and CopyOnWriteArraySet

> `CopyOnWriteArrayList` is a concurrent alternative of synchronized List. It provides better concurrency than synchronized List by allowing multiple concurrent reader and replacing the whole list on write operation. Yes, write operation is costly on `CopyOnWriteArrayList` but it performs better when there are multiple reader and requirement of iteration is more than writing. Since `CopyOnWriteArrayList` Iterator also don't throw ConcurrencModificationException it eliminates need to lock the collection during iteration. Remember both `ConcurrentHashMap` and `CopyOnWriteArrayList` doesn't provides same level of locking as Synchronized Collection and achieves thread-safety by there locking and mutability strategy. So they perform better if requirements suits there nature. Similarly, `CopyOnWriteArraySet` is a concurrent replacement to Synchronized Set.

#### BlockingQueue and `Deque`

> `BlockingQueue` makes it easy to implement producer-consumer design pattern by providing inbuilt blocking support for put() and take() method. put() method will block if Queue is full while take() method will block if Queue is empty. Java 5 API provides two concrete implementation of `BlockingQueue` in form of `ArrayBlockingQueue` and `LinkedBlockingQueue`, both of them implement FIFO ordering of element. ArrayBlockingQueue is backed by Array and its bounded in nature while `LinkedBlockingQueue` is optionally bounded. Consider using `BlockingQueue` to solve producer Consumer problem in Java instead of writing your won wait-notify code. Java 5 also provides `PriorityBlockingQueue`, another implementation of `BlockingQueue` which is ordered on priority and useful if you want to process elements on order other than FIFO.

> `Deque` interface is added in Java 6 and it extends Queue interface to support insertion and removal from both end of Queue referred as head and tail. Java6 also provides concurrent implementation of `Deque` like ArrayDeque and LinkedBlockingDeque. `Deque` Can be used efficiently to increase parallelism in program by allowing set of worker thread to help each other by taking some of work load from other thread by utilizing `Deque` double end consumption property. So if all Thread has there own set of task Queue and they are consuming from head; helper thread can also share some work load via consumption from tail.


#### ConcurrentSkipListMap and ConcurrentSkipListSet

> Just like `ConcurrentHashMap` provides a concurrent alternative of synchronized HashMap. ConcurrentSkipListMap and ConcurrentSkipListSet provide concurrent alternative for synchronized version of SortedMap and SortedSet. For example instead of using TreeMap or TreeSet wrapped inside synchronized Collection, You can consider using ConcurrentSkipListMap or ConcurrentSkipListSet from java.util.concurrent package. They also implement NavigableMap and NavigableSet to add additional navigation method.


## Synchronizer

### Counter Semaphore

* Counting Semaphore in Java maintains specified number of pass or permits, In order to access a shared resource, Current Thread must acquire a permit. If permit is already exhausted by other thread than it can wait until a permit is available due to release of permit from different thread. This concurrency utility can be very useful to implement producer consumer design pattern or implement bounded pool or resources like Thread Pool, DB Connection pool etc.



```java

import java.util.concurrent.Semaphore;

public class SemaphoreDemo {

    Semaphore binary = new Semaphore(1);
  
    public static void main(String args[]) {
        final SemaphoreDemo test = new SemaphoreDemo();
        new Thread(){
            @Override
            public void run(){
              test.mutualExclusion(); 
            }
        }.start();
      
        new Thread(){
            @Override
            public void run(){
              test.mutualExclusion(); 
            }
        }.start();
      
    }
  
    private void mutualExclusion() {
        try {
            binary.acquire();

            //mutual exclusive region
            System.out.println(Thread.currentThread().getName() + " inside mutual exclusive region");
            Thread.sleep(1000);

        } catch (InterruptedException i.e.) {
            ie.printStackTrace();
        } finally {
            binary.release();
            System.out.println(Thread.currentThread().getName() + " outside of mutual exclusive region");
        }
    } 
  
}

// Output:
// Thread-0 inside mutual exclusive region
// Thread-0 outside of mutual exclusive region
// Thread-1 inside mutual exclusive region
// Thread-1 outside of mutual exclusive region

```


### CountDownLatch

* `CountDownLatch` in Java is a kind of synchronizer which allows one Thread  to wait for one or more Threads before starts processing. You can also implement same functionality using  wait and notify mechanism in Java but it requires lot of code and getting it write in first attempt is tricky,  With `CountDownLatch` it can  be done in just few lines. `CountDownLatch` also allows flexibility on number of thread for which main thread should wait, It can wait for one thread or n number of thread, there is not much change on code. 

* The difficulty to use it properly is where to use `CountDownLatch`. First, let us figour out how `CountDownLatch` works. usaullly main thread of application,  which calls CountDownLatch.await() will wait until count reaches zero or its interrupted by another Thread. All other thread are required to do count down by calling CountDownLatch.countDown() once they are completed. One disadvantage of `CountDownLatch` is not reusable, once its count reaches zero. 

* Sample program requires 3 services namely CacheService, AlertService  and ValidationService  to be started and ready before application can handle any request.

```java
import java.util.Date;
import java.util.concurrent.CountDownLatch;
import java.util.logging.Level;
import java.util.logging.Logger;

public class CountDownLatchDemo {

    public static void main(String args[]) {
       final CountDownLatch latch = new CountDownLatch(3);
       Thread cacheService = new Thread(new Service("CacheService", 1000, latch));
       Thread alertService = new Thread(new Service("AlertService", 1000, latch));
       Thread validationService = new Thread(new Service("ValidationService", 1000, latch));
      
       cacheService.start(); //separate thread will initialize CacheService
       alertService.start(); //another thread for AlertService initialization
       validationService.start();
      
      
       //count is 3 since we have 3 Threads (Services)
      
       try{
            latch.await();  //main thread is waiting on CountDownLatch to finish
            System.out.println("All services are up, Application is starting now");
       }catch(InterruptedException ie){
           ie.printStackTrace();
       }
      
    }
  
}

/**
 * Service class which will be executed by Thread using CountDownLatch synchronizer.
 */
class Service implements Runnable{
    private final String name;
    private final int timeToStart;
    private final CountDownLatch latch;
  
    public Service(String name, int timeToStart, CountDownLatch latch){
        this.name = name;
        this.timeToStart = timeToStart;
        this.latch = latch;
    }
  
    @Override
    public void run() {
        try {
            Thread.sleep(timeToStart);
        } catch (InterruptedException ex) {
            Logger.getLogger(Service.class.getName()).log(Level.SEVERE, null, ex);
        }
        System.out.println( name + " is Up");
        latch.countDown(); //reduce count of CountDownLatch by 1
    }  
}



```

### CylicBarrier

* `CyclicBarrier` is similar to `CountDownLatch` which we have seen in the last article  What is `CountDownLatch` in Java and allows multiple threads to wait for each other (barrier) before proceeding. The difference between `CountDownLatch` and `CyclicBarrier` is an also very popular multi-threading interview question in Java. `CyclicBarrier` is a natural requirement for a concurrent program because it can be used to perform final part of the task once individual tasks  are completed.

* The demo of `CyclicBarrier` on which we initialize `CyclicBarrier` with 3 parties, means in order to cross barrier, 3 thread needs to call await() method. each thread calls await method in short duration but they don't proceed until all 3 threads reached the barrier, once all thread reach the barrier, barrier gets broker and each thread started their execution from that point.

```java
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.`CyclicBarrier`;
import java.util.logging.Level;
import java.util.logging.Logger;

public class CyclicBarrierDemo {

    //Runnable task for each thread
    private static class Task implements Runnable {

        private CyclicBarrier barrier;

        public Task(CyclicBarrier barrier) {
            this.barrier = barrier;
        }

        @Override
        public void run() {
            try {
                System.out.println(Thread.currentThread().getName() + " is waiting on barrier");
                barrier.await();
                System.out.println(Thread.currentThread().getName() + " has crossed the barrier");
            } catch (InterruptedException ex) {
                Logger.getLogger(CyclicBarrierDemo.class.getName()).log(Level.SEVERE, null, ex);
            } catch (BrokenBarrierException ex) {
                Logger.getLogger(CyclicBarrierDemo.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
    }

    public static void main(String args[]) {

        //creating CyclicBarrier with 3 parties i.e. 3 Threads needs to call await()
        final CyclicBarrier cb = new CyclicBarrier(3, new Runnable(){
            @Override
            public void run(){
                //This task will be executed once all thread reaches barrier
                System.out.println("All parties are arrived at barrier, lets play");
            }
        });

        //starting each of thread
        Thread t1 = new Thread(new Task(cb), "Thread 1");
        Thread t2 = new Thread(new Task(cb), "Thread 2");
        Thread t3 = new Thread(new Task(cb), "Thread 3");

        t1.start();
        t2.start();
        t3.start();
      
    }
}

// Output:
// Thread 1 is waiting on barrier
// Thread 3 is waiting on barrier
// Thread 2 is waiting on barrier
// All parties have arrived at barrier, lets play
// Thread 3 has crossed the barrier
// Thread 1 has crossed the barrier
// Thread 2 has crossed the barrier

```


### Producer / Consumer pattern 

* Producer Consumer Design pattern is a classic concurrency or threading pattern which reduces coupling between
Producer and Consumer by separating Identification of work with Execution of Work. In producer consumer design pattern a shared queue is used to control the flow and this separation allows you to code producer and consumer separately.

* It is everywhere in real life and depict coordination and collaboration. Like one person is preparing food (Producer) while other one is serving food (Consumer), both will use shared table for putting food plates and taking food plates. Producer which is the person preparing food will wait if table is full and Consumer (Person who is serving food) will wait if table is empty. table is a shared object here. On Java library Executor framework itself implement Producer Consumer design pattern be separating responsibility of addition and execution of task.

```java
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.logging.Level;
import java.util.logging.Logger;

public class ProducerConsumerPattern {

    public static void main(String args[]){
  
     //Creating shared object
     BlockingQueue sharedQueue = new LinkedBlockingQueue();
 
     //Creating Producer and Consumer Thread
     Thread prodThread = new Thread(new Producer(sharedQueue));
     Thread consThread = new Thread(new Consumer(sharedQueue));

     //Starting producer and Consumer thread
     prodThread.start();
     consThread.start();
    }
 
}

//Producer Class in java
class Producer implements Runnable {

    private final BlockingQueue sharedQueue;

    public Producer(BlockingQueue sharedQueue) {
        this.sharedQueue = sharedQueue;
    }

    @Override
    public void run() {
        for(int i=0; i<10; i++){
            try {
                System.out.println("Produced: " + i);
                sharedQueue.put(i);
            } catch (InterruptedException ex) {
                Logger.getLogger(Producer.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
    }

}

//Consumer Class in Java
class Consumer implements Runnable{

    private final BlockingQueue sharedQueue;

    public Consumer (BlockingQueue sharedQueue) {
        this.sharedQueue = sharedQueue;
    }
  
    @Override
    public void run() {
        while(true){
            try {
                System.out.println("Consumed: "+ sharedQueue.take());
            } catch (InterruptedException ex) {
                Logger.getLogger(Consumer.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
    }
  
  
}

// Output:
// Produced: 0
// Produced: 1
// Consumed: 0
// Produced: 2
// Consumed: 1
// Produced: 3
// Consumed: 2
// Produced: 4
// Consumed: 3
// Produced: 5
// Consumed: 4
// Produced: 6
// Consumed: 5
// Produced: 7
// Consumed: 6
// Produced: 8
// Consumed: 7
// Produced: 9
// Consumed: 8
// Consumed: 9

```


### Executor -- Thread Pool

* Java 1.5 introduced Thread pool in Java in the form of Executor framework, which allows Java programmer to decouple submission of a task to execution of the task. It also introduced a full feature built-in Thread Pool framework commonly known as Executor framework. Executor framework also provides different kind of Thread Pool e.g. `SingleThreadExecutor` which creates just one worker thread or `CachedThreadPool` which creates worker threads as and when necessary. 

* Demo of thread pool

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class ThreadPoolDemo {

    public static void main(String args[]) {
       ExecutorService service = Executors.newFixedThreadPool(10);
       for (int i =0; i<100; i++){
           service.submit(new Task(i));
       }
    }
  
}

final class Task implements Runnable{
    private int taskId;
  
    public Task(int id){
        this.taskId = id;
    }
  
    @Override
    public void run() {
        System.out.println("Task ID : " + this.taskId +" performed by " + Thread.currentThread().getName());
    }  
}

// Output:
// Task ID : 0 performed by pool-1-thread-1
// Task ID : 7 performed by pool-1-thread-8
// Task ID : 8 performed by pool-1-thread-9
// Task ID : 6 performed by pool-1-thread-7
// Task ID : 4 performed by pool-1-thread-5
// Task ID : 5 performed by pool-1-thread-6
// Task ID : 3 performed by pool-1-thread-4
// Task ID : 1 performed by pool-1-thread-2
// ...
```

### Submit(Runnable)

> The submit(Runnable) method also takes a Runnable implementation, but returns a Future object. This Future object can be used to check if the Runnable as finished executing.

### Submit(Callable)

> The submit(Callable) method is similar to the submit(Runnable) method except for the type of parameter it takes. The Callable instance is very similar to a Runnable except that its call() method can return a result.

### InvokeAny()

> The invokeAny() method takes a collection of Callable objects, or subinterfaces of Callable. If one of the tasks complete (or throws an exception), the rest of the Callable's are cancelled.

### InvokeAll()

> The invokeAll() method invokes all of the Callable objects and returns a list of Future objects. Keep in mind that a task might finish due to an exception, so it may not have "succeeded". There is no way on a Future to tell the difference.

### Demo of submit, InvokeAny

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Callable;
import java.util.concurrent.Future;
import java.util.concurrent.ExecutionException;
import java.util.HashSet;
import java.util.Set;

public class SubmitInvokeDemo {

    public static void main(String args[]) throws InterruptedException, ExecutionException {


        ExecutorService executorService = Executors.newSingleThreadExecutor();

        Future future = executorService.submit(new Runnable() {
            public void run() {
                System.out.println("Asynchronous task");
            }
        });

        future.get();  //returns null if the task has finished correctly.

        Set<Callable<String>> callables = new HashSet<Callable<String>>();

        callables.add(new Callable<String>() {
            public String call() throws Exception {
                return "Task 1";
            }
        });
        callables.add(new Callable<String>() {
            public String call() throws Exception {
                return "Task 2";
            }
        });
        callables.add(new Callable<String>() {
            public String call() throws Exception {
                return "Task 3";
            }
        });

        String result = executorService.invokeAny(callables);

        System.out.println("result = " + result);

        executorService.shutdown();

    }
}

```




## Join and Fork 

* Here is an introduction into the Fork/Join Framework that is part of the JDK since version 1.7.

### Join and Fork with Executor Service

* The demo code submit() our tasks to the ExecutorService and then use the returned instance of `Future` to wait() for the result. The normal
`ExecutorService` where you would have to block the current thread while waiting for a result. If we would only provide as many threads to the pool as we have CPUs available, the program would run out of resources and hang indefinitely.

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Callable;
import java.util.concurrent.Future;
import java.util.concurrent.ExecutionException;
import java.util.Random;

public class FindMinTask implements Callable<Integer> {
    private int[] numbers;
    private int startIndex;
    private int endIndex;
    private ExecutorService executorService;

    public FindMinTask(ExecutorService executorService, int[] numbers, int startIndex, int endIndex) {
        this.executorService = executorService;
        this.numbers = numbers;
        this.startIndex = startIndex;
        this.endIndex = endIndex;
    }

    public Integer call() throws Exception {

        int sliceLength = (endIndex - startIndex) + 1;
        if (sliceLength > 2) {
            FindMinTask lowerFindMin = new FindMinTask(executorService, numbers, startIndex,
                    startIndex + (sliceLength / 2) - 1);
            Future<Integer> futureLowerFindMin = executorService.submit(lowerFindMin);
            FindMinTask upperFindMin = new FindMinTask(executorService, numbers, startIndex + (sliceLength / 2),
                    endIndex);
            Future<Integer> futureUpperFindMin = executorService.submit(upperFindMin);
            return Math.min(futureLowerFindMin.get(), futureUpperFindMin.get());
        } else {
            return Math.min(numbers[startIndex], numbers[endIndex]);
        }
    }

    public static void main(String[] args) throws InterruptedException, ExecutionException {
        int[] numbers = new int[100];
        Random random = new Random(System.currentTimeMillis());
        for (int i = 0; i < numbers.length; i++) {
            numbers[i] = random.nextInt(100);
        }
        ExecutorService executorService = Executors.newFixedThreadPool(64);
        Future<Integer> futureResult = executorService
                .submit(new FindMinTask(executorService, numbers, 0, numbers.length - 1));
        System.out.println(futureResult.get());
        executorService.shutdown();
    }
}
```


### Join and Fork with JoinForkPool

* The `ForkJoinPool` implements the already mentioned work-stealing strategy, i.e. every time a running thread has to wait for some result; the thread removes the current task from the work queue and executes some other task ready to run. This way the current thread is not blocked and can be used to execute other tasks. Once the result for the originally suspended task has been computed the task gets executed again and the join() method returns the result. This is an important difference between `JoinForkPool` and `ExecutorService`.

* Demo of JoinForkPool

```java
import java.awt.*;
import java.awt.image.*;
import java.io.*;
import java.util.concurrent.ForkJoinPool;
import javax.imageio.*;
import java.util.concurrent.RecursiveAction;

public class GrayscaleImageAction extends RecursiveAction {
    private static final long serialVersionUID = 1L;
    private int row;
    private BufferedImage bufferedImage;

    public GrayscaleImageAction(int row, BufferedImage bufferedImage) {
        this.row = row;
        this.bufferedImage = bufferedImage;
    }

    @Override
    protected void compute() {
        for (int column = 0; column < bufferedImage.getWidth(); column++) {
            int rgb = bufferedImage.getRGB(column, row);
            int r = (rgb >> 16) & 0xFF;
            int g = (rgb >> 8) & 0xFF;
            int b = (rgb & 0xFF);
            int gray = (int) (0.2126 * (float) r + 0.7152 * (float) g + 0.0722 * (float) b);
            gray = (gray << 16) + (gray << 8) + gray;
            bufferedImage.setRGB(column, row, gray);
        }
    }

    public static void main(String[] args) throws IOException {
        ForkJoinPool pool = new ForkJoinPool(Runtime.getRuntime().availableProcessors());
        BufferedImage bufferedImage = ImageIO.read(new File(args[0]));
        for (int row = 0; row < bufferedImage.getHeight(); row++) {
            GrayscaleImageAction action = new GrayscaleImageAction(row, bufferedImage);
            pool.execute(action);
        }
        pool.shutdown();
        ImageIO.write(bufferedImage, "jpg", new File(args[1]));
    }
}

```