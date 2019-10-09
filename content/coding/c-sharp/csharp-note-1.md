+++

title = "Console App"
description = "How to create C# console application without Visual Studio"
weight=10

+++

### Prerequisites

> *.Net Framework has been install to your PC or laptop*


### Create simple .net project without Visual Studio

#### Prerequisites 

* The version of .Net Framework on your PC is 2.0+ or later
* Assume the path  of .Net Frameowork is `c:\Windows\Microsfot.Net\Frameowork\v2.0.50727`

#### Simple C# project

### Create a proejct named `csharp-project` 

```bash
md csharp-project
cd csharp-project 
md bin src
echo.>csharp-project.proj
echo.>src\helloworld.cs
```

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

* Compile with MSBuild & Run 

```bash
c:\Windows\Microsfot.Net\Frameowork\v2.0.50727\MSBuild
bin\csharp-app.exe
```






