@echo OFF
setlocal EnableDelayedExpansion  

set CWS=%~dp0

echo  Deploying updates to GitHub...

echo "%CWS%\n"

set bakDir="%CWS%bak\"
set pubDir="%CWS%public\"

echo  Backup Folder %bakDir%


REM  Backup the public folder for recovery
if NOT EXIST   %bakDir% (
  echo Create a new bak folder
  mkdir bak
)

if EXIST %pubDir% (
  echo Backup folder public into folder bak
  xcopy /s /e /q /y public bak

  REM Clean up the public folder except the git file
  echo Cleanup folder public 

  del /s /q  /f  public\*

  cd %pubDir%

  for /f "tokens=*" %%S in ('dir /b  /a:d .') do rd /q /s %%S 

  cd ..
)

REM Build the project. If using a theme, replace by `hugo -t <yourtheme>`
echo buiding the site ...
hugo -t docdock

REM Navigate to public folder
cd %pubDir%

REM Add changes to git.
echo git config setup

git config user.name "harryho"
git config user.email "harry.ho_long@yahoo.com"
git add -A

REM Commit changes.
echo git commit ...
set msg="rebuilding site %date%"
if  NOT "%1"=="" set msg=%1
git commit -m '%msg%'

REM Push source and build repos.
git push origin master

REM Navigate back to root folder
cd ..

REM Cleanup the bak folder
rd /s /q %bakDir%

endlocal DisableDelayedExpansion
