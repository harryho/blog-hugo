+++
categories = ["code"]
date = "2015-04-10T14:59:31+11:00"
title = "C# Notes -- Part 2"
draft = false
+++

# Problem

> *Set up the some scheduled tasks running in the backgroud to take care of data update or sync for every 15 mins, or everyday or every week*

# Solution

## Options as scheduled backgroud service

### Windows Task Scheduler

* Click the Start button.
* Click Control Panel.
* Click System and Maintenance.
* Click Administrative Tools.
* Double-click Task Scheduler.

### Use Window Service as task scheduler

#### Overview of design

**The design here is a simplified version, which I built for previous projects. In the real world, I need to tailor it for different project with different purpose, but fundamental design of architect is the same. IMO, the design below can support most cases, which needs scheduled backgroud task service.**

```ini

        +----------------------------------+          Register as
        |     System.ServiceProces  ------ |--------O Window Service
        +----------------------------------+
        |  +----------------------------+  |
        |  |   Thread (Infinite)        |  | 
        |  +----------------------------+  |
        |  |  +----------------------+  |  |
        |  |  | MySchedulerService   |  |  | 
        |  |  +----------------------+  |  |
        |  |  | while ( true )       |  |  |
        |  |  | {                    |  |  |
        |  |  |    Task.Process()    |  |  |
        |  |  | }                    |  |  | 
        |  |  +----------------------+  |  | 
        |  +----------------------------+  |
        +----------------------------------+ 

            +------------+
            | ITask      |
            +------------+
            | Process()  |
            +------------+
                 /|\
                  |
        +------------------+
        |      BaseTask    |  -----------O Customized Task inherit BaseTask   
        +------------------+
        | lastProcessTime  | ------------O Last process time 
        | intervalTime     | ------------O Customize for next process time 
        | IsReadyProcess() | ------------O Check taks is ready to process 
        +------------------+

```

#### Create customized Window Service

#### Use ServiceProcess to create Window Service

* The System.ServiceProcess allow you to implement, install, and control Windows service applications. Services are long-running executables that run without a user interface. 
*  The project must have main method the entry point

```cs
namespace Scheduler
{
    using System.ServiceProcess;

    class Program
    {
        static void Main(string[] args)
        {
            #if DEBUG
                MySchedulerService debugService = new MySchedulerService();
                debugService.onDebug();
                System.Threading.Thread.Sleep(System.Threading.Timeout.Infinite);
            #else
                ServiceBase.Run(new ServiceBase[] { new MySchedulerService() });
            #endif
        }
    }
}

```

#### Create customized ServiceBase

* Implementing a service involves inheriting from the ServiceBase class and defining specific behavior to process when start, stop, pause, and continue commands are passed in, as well as custom behavior and actions to take when the system shuts down.

```cs

namespace Scheduler
{
    using System;
    using System.IO;
    using System.Collections.Generic;
    using System.Threading;

    public partial class MySchedulerService : System.ServiceProcess.ServiceBase
    {
        public const string START_SERVICE = "start.service";
        public const string STOP_SERVICE = "stop.service";

        public void onDebug()
        {
            OnStart(null);
        }

        protected override void OnStart(string[] args)
        {
            System.IO.File.Create(AppDomain.CurrentDomain.BaseDirectory + TART_SERVICE);
            ThreadStart tsTask = new ThreadStart(TaskLoop);
            Thread rtkTask = new Thread(tsTask);
            rtkTask.Start();
        }

        static void TaskLoop()
        {
            while (true)
            {
                // Exceute scheduled task
                ScheduledTask();
                // Then, wait for certain time interval
                System.Threading.Thread.Sleep(TimeSpan.FromMilliseconds(500));
            }
        }

        static void ScheduledTask()
        {
            // Task code which is executed periodically
            try
            {
                // Call customized tasks
                var types = Assembly.GetExecutingAssembly().GetExportedTypes().Where(p => typeof(ITask).IsAssignableFrom(p.BaseType));

                foreach (var t in types)
                {
                    var task = (ITask)Activator.CreateInstance(t);
                    if (taskSettings.Keys.Contains(t.Name))
                    {
                        task.TaskSetting = taskSettings[t.Name];
                    }
                    tasks.Add(task);
                }
            }
            catch (Exception e)
            {
                // TO Something here               
            }
        }

        protected override void OnStop()
        {
            // Insert code here to close all the open IO or conection.
            System.IO.File.Create(AppDomain.CurrentDomain.BaseDirectory + STOP_SERVICE);
        }

        private void InitializeComponent()
        {
            this.ServiceName = "MySchedulerService";
        }

        protected override void Dispose(bool disposing)
        {
            OnStop();
            base.Dispose(disposing);
        }
    }
}

```

#### Task interface and class 

* interface ITask has only one method 

```cs
namespace Scheduler
{
    public interface ITask
    {
        void Process();
    }
}
```

* BaseTask 

```cs
namespace Scheduler
{
    using System;

    public class BaseTask : ITask
    {
        protected DateTime? lastProcessTime = null;

        // interval time uses second as unit
        // e.g. 1 min of intervaling time is 60 
        protected int intervalTime = 0 ;  

        protected bool IsReadyToProcess()
        {
            bool isReadyToProcess = true;

            if (lastProcessTime.HasValue)
            {
                if (lastProcessTime.Value.AddSeconds(intervalTime) > DateTime.Now)
                {
                    isReadyToProcess = false;
                }
            } 
            return isReadyToProcess;
        }

        public virtual void Process()
        {
            throw new NotImplementedException();
        }
    }
}
```

* Sample Task1 and Task2

**Task1**

```cs
namespace Scheduler
{
    using System;
    using System.Linq;
    using System.Data.Entity;

    public class Task1 : BaseTask
    {
        public override void Process()
        {
            if (base.IsReadyToProcess())
            {                
                System.IO.File.Create(AppDomain.CurrentDomain.BaseDirectory + "Task-1-" + DateTime.Now.ToString("dd-MM-yyyy"));
            }
        }
    }
}
```
**Task2**

```cs
namespace Scheduler
{
    using System;
    using System.Linq;
    using System.Data.Entity;

    public class Task2 : BaseTask
    {
        public override void Process()
        {
            if (base.IsReadyToProcess())
            {                
                System.IO.File.Create(AppDomain.CurrentDomain.BaseDirectory + "Task-2-" + DateTime.Now.ToString("dd-MM-yyyy"));
            }
        }
    }
}
```

### Create Window Service installer

* In Solution Explorer, double-click MyScheduledService.cs.
* In the Code Editor window, right-click `Design View`, and then click `Properties`
* In the `Properties` pane, click the `Add Installer` link.
* In the Properties pane for MyScheduledServiceInstaller, change the `ServiceNameproperty` to `MyScheduledService`.
* In the Code Editor window in `Design view`, click `MyScheduledServiceProcessInstaller`.
* In the Properties pane, change the `Account` property to `LocalSystem` (The `LocalService` and `NetworkService` values are available only in Microsoft Windows XP).

**PreInstaller**

* It inherits from System.Configuration.Install.Installer. This is the base class for all custom installers in the .NET Framework. Installers are components that help install applications on a computer.

```cs
namespace Scheduler
{
    [RunInstaller(true)]
    public partial class ProjectInstaller : System.Configuration.Install.Installer
    {
        public ProjectInstaller()
        {
            InitializeComponent();
        }

        private void MySchedulerServiceInstaller_AfterInstall(object sender, InstallEventArgs e)
        {
        }
    }
}
```

### Install Window Service 

```bash
C:\Windows\Microsoft.Net\Framework\v4.0.30319\InstallUtil.exe /i <app_path>\NBE.Scheduler.exe 
```

### Uninstall Window Service 

```bash
C:\Windows\Microsoft.Net\Framework\v4.0.30319\InstallUtil.exe /u <app_path>\NBE.Scheduler.exe 
```