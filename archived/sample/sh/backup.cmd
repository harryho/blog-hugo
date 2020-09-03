@echo off
setlocal EnableDelayedExpansion  

set logfile=backup.log
set CWS=%~dp0

echo Start backup %date%-%time%

echo --------------------------------------------------

if "%1"=="" (
	echo First parameter must be fe or be 
	echo sample: .\backup.cmd fe 1234
	pause 
	exit 0
)
if  "%1"=="fe" (  
	set forb=FE
	set app=nsw
)
if  "%1"=="be"  ( 
	set forb=BE
	set app=nswadmin
)

echo FE or BE: %forb%
echo app: %app%


if "%2"=="" (
	echo Second parameter can not be empty. It should be a number of the ticket 
	echo sample: .\backup.cmd fe 1234
	pause
	exit 0
)
set ticket=%2



REM echo "Start backup %date%-%time%">backup.log


echo "%CWS%\n"

set folder=%date:~-4,4%%date:~-10,2%%date:~7,2%_%forb%_%ticket%

set src=d:\data\%app%
set dest=d:\backup\%folder%\%app%\

echo %src% 
echo %dest%


mkdir  %dest%

echo "Create folder  d:\backup\%folder%"

REM dir d:\backup\

set excludelist=%CWS%loglist.txt

dir /b d:\data\%app%\log>%excludelist%

echo Create exclude list %excludelist%


xcopy /s /e  /h /q /y  /EXCLUDE:%excludelist%    %src% %dest%


echo %app% has been backup to %dest% 

echo DONE %date%-%time%

echo  ########################################################

pause 

endlocal DisableDelayedExpansion