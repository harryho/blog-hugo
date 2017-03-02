+++
tags =  ["windows","cmd"]
categories = ["blog"]
date = "2014-05-04T10:59:31+11:00"
title = "Use Windows command & hotkey as a hacker - Part 3"
draft = true
+++



*This article will continue the topic of Windows command & hotkeys. [Part-1](/blog/use-windows-command-hotkey-as-hacker-1) shows you common hotkeys and short command lines for `Run` windnow dialog. [Part-2](/blog/use-windows-command-hotkey-as-hacker-2) advanced commands and how to create a batch script with all those commands. Here I am going to show you another secret weapon in Windows system-VBScript/JScript*

## Breif history

> VBScript/JScript is an Active Scripting language developed by Microsoft that is modeled on Visual Basic. It allows Microsoft Windows system administrators to generate powerful tools for managing computers with error handling, subroutines, and other advanced programming constructs. 

> A VBScript script must be executed within a host environment, of which there are several provided with Microsoft Windows, including: Windows Script Host (WSH), Internet Explorer (IE), and Internet Information Services (IIS). VBScript uses the Component Object Model to access elements of the environment within which it is running.

> JScript is Microsoft's dialect of the ECMAScript standard that is used in Microsoft's Internet Explorer. JScript is implemented as an Active Scripting engine. This means that it can be "plugged in" to OLE Automation applications that support Active Scripting, such as Internet Explorer, Active Server Pages, and Windows Script Host.

> With Cscript.exe, you can run scripts by typing the name of a script file at the command prompt. Like Microsoft Internet Explorer, Windows Script Host serves as a controller of Windows Script compliant scripting engines, but Windows Script Host has very low memory requirements. Windows Script Host is ideal for both interactive and non-interactive scripting needs, such as logon scripting and administrative scripting.


### Sample of VBScript file 

```ps
Set oFSO = CreateObject("Scripting.FileSystemObject"^)
Set oInput = oFSO.OpenTextFile(WScript.Arguments(0^), 1^)
sData = Replace(oInput.ReadAll, "," ^& VbCrLf, VbCrLf^)
Set oOutput = oFSO.CreateTextFile(WScript.Arguments(1^), True^)
oOutput.Write sData
oInput.Close
oOutput.Close

```

### Replace content script

#### Create a file named `repltxt.bat` and open the file with `Notepad`
#### Copy & paste the code into the file `repltxt.bat`, then script is ready to play


```powershell

@if (@X)==(@Y) @end /* Harmless hybrid line that begins a JScript comment

::************ Documentation ***********
::NEWREPL.BAT version 1.0
:::
:::NEWREPL  Search  replace
:::NEWREPL  /?
:::NEWREPL  /V


@echo off
if .%2 equ . (
  if "%~1" equ "/?" (
    echo."%~f0"
    <"%~f0" cscript //E:JScript //nologo "%~f0" "^:::" "" a
    exit /b 0
  ) else if /i "%~1" equ "/V" (
    <"%~f0" cscript //E:JScript //nologo "%~f0" "^::(NEWREPL\.BAT version)" "$1" a
    exit /b 0
  ) else (
    call :err "Insufficient arguments"
    exit /b 2
  )
)

cscript //E:JScript //nologo "%~f0" %*
exit /b %errorlevel%

:err
>&2 echo ERROR: %~1. Use newrepl /? to get help.
exit /b

************* JScript portion **********/
var rtn=1;
try {
    var env=WScript.CreateObject("WScript.Shell").Environment("Process");
    var args=WScript.Arguments;
    var search=args.Item(0);
    var replace=args.Item(1);
    var options="g";
    if (args.length>2) options+=args.Item(2).toLowerCase();
    var alterations=(options.indexOf("a")>=0);
    if (alterations) options=options.replace(/a/g,"");
    if (options.indexOf("v")>=0) {
        options=options.replace(/v/g,"");
        search=env(search);
        replace=env(replace);
    }
    var search=new RegExp(search,options);
    var str1, str2;

      while (!WScript.StdIn.AtEndOfStream) {
        str1=WScript.StdIn.ReadLine();
        str2=str1.replace(search,replace);
        if (!alterations || str1!=str2) WScript.Stdout.WriteLine(str2);
        if (str1!=str2) rtn=0;
    }
} catch(e) {
  WScript.Stderr.WriteLine("JScript runtime error: "+e.message);
  rtn=3;
}
WScript.Quit(rtn);

```

#### 