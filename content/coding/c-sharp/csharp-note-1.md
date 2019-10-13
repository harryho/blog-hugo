+++

title = "C# Console App"
description = "How to create C# console application without Visual Studio"
weight=10

+++

> __C# is an elegant and type-safe object-oriented language that enables developers to build a variety of secure and robust applications that run on the .NET Framework. C# syntax is highly expressive, yet it is also simple and easy to learn. The curly-brace syntax of C# will be instantly recognizable to anyone familiar with C, C++ or Java.__


Here I am going to demonstrate how to create simple .net project without Visual Studio 

### Prerequisites

* .Net Framework has been install to your PC or laptop, and the version of .Net Framework on your PC is 2.0+ or later
* Assume the path  of .Net Frameowork is __c:\Windows\Microsfot.Net\Frameowork\v2.0.50727__



### Create a project

* Create a project named `csharp-project` 

    ```bat
    md csharp-project
    cd csharp-project 
    md bin src
    echo.>csharp-project.proj
    echo.>src\helloworld.cs
    ```

### Update config file 

* Update config file `csharp-project.proj` as below

```xml
<Project DefaultTargets = "Compile"
    xmlns="http://schemas.microsoft.com/developer/msbuild/2003" >

    <!-- Set the application name as a property -->
    <PropertyGroup>
        <appname>csharp-app</appname>
    </PropertyGroup>

    <!-- Specify the inputs by type and file name -->
    <ItemGroup>
        <CSFile Include = "src\helloworld.cs"/>
    </ItemGroup>

    <Target Name = "Compile">
        <!-- Run the Visual C# compilation using input files of type CSFile -->
        <CSC
            Sources = "@(CSFile)"
            OutputAssembly = "bin\$(appname).exe">
            <!-- Set the OutputAssembly attribute of the CSC task
            to the name of the executable file that is created -->
            <Output
                TaskParameter = "OutputAssembly"
                ItemName = "EXEFile" />
        </CSC>
        <!-- Log the file name of the output file -->
        <Message Text="The output file is @(EXEFile)"/>
    </Target>
</Project>

```


### Create the main program

* Create a new file named `HelloWorld.cs` and copy following coding to the new file

    ```csharp
    public class Hello
    {
        public static void Main()
        {
            System.Console.WriteLine("Hello, World!");
        }
    }
    ```

### Run the console app

* Compile with MSBuild & Run 

    ```bat
    c:\Windows\Microsfot.Net\Frameowork\v2.0.50727\MSBuild
    bin\csharp-app.exe
    ```






