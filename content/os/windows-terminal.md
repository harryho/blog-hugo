+++
title = "Windows Terminal"
description = "Empower Windows with Terminal"
+++


### Windows Terminal

> The Windows Terminal is a modern, fast, efficient, powerful, and productive terminal application for users of command-line tools and shells like Command Prompt, PowerShell, and WSL. Its main features include multiple tabs, panes, Unicode and UTF-8 character support, a GPU accelerated text rendering engine, and custom themes, styles, and configurations.

#### Prerequisite
- PowerShell 7
  - Install latest PowerShell (version 7.x) from Windows Store
- Scoop 
  - [Scoop](https://scoop.sh) can be found here or simply execute the command below
    
    ```powershell
    Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')

    # or shorter
    iwr -useb get.scoop.sh | iex
    ```

#### Install Windows Terminal

This is an open source project and we welcome community participation. To participate please visit https://github.com/microsoft/terminal

#### Install Nerd font

- [Nerd Font](https://www.nerdfonts.com/font-downloads)
- Download Meslo Nerd Font
- Install font file



#### Update Setting

- Settings > Startups
    - Profile -> PowerShell ( Default is Windows PowerShell )
    - Launch mode -> Maximised

- Update Setting file

    - Hide the proflle Windows PowerShell & move the PowerShell to the top

    ```json
    "list": [
        {
                "guid": "{574e775e-4f2a-5b96-ac1e-a2962a402336}",
                "hidden": false,
                "name": "PowerShell",
                "source": "Windows.Terminal.PowershellCore"
        },
        ...
        {
                "commandline": "powershell.exe",
                "font": 
                {
                    "face": "MesloLGM NF"
                },
                "guid": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}",
                "hidden": true,
                "name": "Windows PowerShell"
        } 
    ```

    - Create a customzed color scheme

    ```json
    {
        "background": "#001D26",
        "black": "#282C34",
        "blue": "#61AFEF",
        "brightBlack": "#9aabc5",
        "brightBlue": "#61AFEF",
        "brightCyan": "#56B6C2",
        "brightGreen": "#98C379",
        "brightPurple": "#C678DD",
        "brightRed": "#E06C75",
        "brightWhite": "#DCDFE4",
        "brightYellow": "#E5C07B",
        "cursorColor": "#FFFFFF",
        "cyan": "#56B6C2",
        "foreground": "#DCDFE4",
        "green": "#98C379",
        "name": "One Half Dark (mod)",
        "purple": "#C678DD",
        "red": "#E06C75",
        "selectionBackground": "#FFFFFF",
        "white": "#DCDFE4",
        "yellow": "#E5C07B"
    }
    ```

- Settings > Default > Appearence
  - Color scheme : One Half Dark (mod)
  - Font face:  MesloLGM NF
  - Acrylic opacity: 60% 


#### Git

```
winget install -e --id Git.Git
```

#### Other tools

```
scoop install curl sudo jq yq
scoop install neovim gcc 
scoop install 7zip bat 
```

#### Setup User Profile

- Create user profile folder

```cmd
mkdir ~\.config\powershell
```

- Create a launch script

```powershell
type nul > $env:USERPROFILE\.config\powershell\profile.ps1
```

- Add alias to `$env:USERPROFILE\.config\powershell\profile.ps1`

```powershell
# Alias
Set-Alias vi nvim
Set-Alias ll ls
Set-Alias g git
Set-Alias grep findstr
Set-Alias tig $env:USERPROFILE\app\git\usr\bin\tig.exe
Set-Alias less $env:USERPROFILE\app\git\usr\bin\less.exe
```

- Update built-in profile

```powershell
# Get all profiles
$PROFILE | Get-Member -Type NoteProperty

# If the file of $PROFILE.CurrentUserCurrentHost does not exist
type nul > $PROFILE.CurrentUserCurrentHost
```

- Copy content below to `$PROFILE.CurrentUserCurrentHost`

```
. $env:USERPROFILE\.config\powershell\profile.ps1
```


#### Oh My Posh

- Installation

  [Please follow latest instruction](https://ohmyposh.dev/docs/migrating)


- Create customized theme file `sample.omp.json`

```
{
  "$schema": "https://raw.githubusercontent.com/JanDeDobbeleer/oh-my-posh/main/themes/schema.json",
  "final_space": false,
  "osc99": true,
  "blocks": [
    {
      "type": "prompt",
      "alignment": "left",
      "segments": [
        {
          "type": "shell",
          "style": "diamond",
          "leading_diamond": "╭─",
          "trailing_diamond": "",
          "foreground": "#ffffff",
          "background": "#0077c2",
          "properties": {
          }
        },
        {
          "type": "root",
          "style": "diamond",
          "leading_diamond": "",
          "trailing_diamond": "",
          "foreground": "#FFFB38",
          "background": "#ef5350",
          "properties": {
            "root_icon": "\uf292",
            "prefix": "<parentBackground>\uE0B0</> "
          }
        },
        {
          "type": "path",
          "style": "powerline",
          "powerline_symbol": "\uE0B0",
          "foreground": "#E4E4E4",
          "background": "#444444",
          "properties": {
            "style": "full",
            "enable_hyperlink": true
          }
        },
        {
          "type": "git",
          "style": "powerline",
          "powerline_symbol": "\uE0B0",
          "foreground": "#011627",
          "background": "#FFFB38",
          "background_templates": [
            "{{ if or (.Working.Changed) (.Staging.Changed) }}#ffeb95{{ end }}",
            "{{ if and (gt .Ahead 0) (gt .Behind 0) }}#c5e478{{ end }}",
            "{{ if gt .Ahead 0 }}#C792EA{{ end }}",
            "{{ if gt .Behind 0 }}#C792EA{{ end }}"
          ],
          "properties": {
            "branch_icon": "\ue725 ",
            "fetch_status": true,
            "fetch_upstream_icon": true,
            "template": "{{ .HEAD }} {{ if .Working.Changed }}{{ .Working.String }}{{ end }}{{ if and (.Working.Changed) (.Staging.Changed) }} |{{ end }}{{ if .Staging.Changed }}<#ef5350> \uF046 {{ .Staging.String }}</>{{ end }}"
          }
        }
      ]
    },
    {
      "type": "prompt",
      "alignment": "right",
      "segments": [
        {
          "type": "node",
          "style": "diamond",
          "leading_diamond": " \uE0B6",
          "trailing_diamond": "\uE0B4",
          "foreground": "#3C873A",
          "background": "#303030",
          "properties": {
            "prefix": "\uE718 ",
            "postfix": "",
            "display_package_manager": true,
            "yarn_icon": " <#348cba></>",
            "npm_icon": " <#cc3a3a></> "
          }
        },
        {
          "type": "time",
          "style": "diamond",
          "invert_powerline": true,
          "leading_diamond": " \uE0B6",
          "trailing_diamond": "\uE0B4",
          "background": "#40c4ff",
          "foreground": "#ffffff",
          "properties": {
            "prefix": " \uf5ef ",
            "postfix": " "
        }
        }
      ]
    },
    {
      "type": "prompt",
      "alignment": "left",
      "newline": true,
      "segments": [
        {
          "type": "text",
          "style": "plain",
          "foreground": "#21c7c7",
          "properties": {
            "prefix": "",
            "postfix": "",
            "text": "╰─"
          }
        },
        {
          "type": "exit",
          "style": "plain",
          "foreground": "#e0f8ff",
          "properties": {
            "prefix": "\u276F",
            "display_exit_code": false,
            "always_enabled": true,
            "error_color": "#ef5350"
          }
        }
      ]
    }
  ]
}

```


- Update user profile

  - Add content to `$env:USERPROFILE\.config\powershell\profile.ps1`

  ```powershell

  # Prompt
  Import-Module posh-git
  $omp_config = "$env:USERPROFILE\sample.omp.json"
  oh-my-posh --init --shell pwsh --config $omp_config | Invoke-Expression
  ```


#### Posh-Git 

```
Install-Module posh-git -Repository PSGallery -Force
```

#### Terminal Icons

- Installation

```powershell
 Install-Module -Name Terminal-Icons -Repository PSGallery -Force
```

- Update user profile
  - Add content to `$env:USERPROFILE\.config\powershell\profile.ps1`

  ```powershell
  Import-Module -Name Terminal-Icons
  ```

#### PSReadLine

- Installation


```powershell
Install-Module -Name PSReadLine  -AllowPrerelease -Scope CurrentUser -Force -SkipPublisherCheck
```

- Update user profile
  - Add content to `$env:USERPROFILE\.config\powershell\profile.ps1`

  ```powershell
  # PSReadLine
  Set-PSReadLineOption -EditMode Emacs
  Set-PSReadLineOption -BellStyle None
  Set-PSReadLineKeyHandler -Chord 'Ctrl+d' -Function DeleteChar
  Set-PSReadLineOption -PredictionSource History
  ```

#### FZF

- Installation

```powershell
scoop install fzf
Install-Module -Name PSFzf -Scope CurrentUser -Force
```

- Update user profile
  - Add content to `$env:USERPROFILE\.config\powershell\profile.ps1`
  ```powershell
  # Fzf
  Import-Module PSFzf
  Set-PsFzfOption -PSReadlineChordProvider 'Ctrl+f' -PSReadlineChordReverseHistory 'Ctrl+r'
  ```
#### Z

```
Install-Module -Name z --Respository SGallery -Force
```

#### Exmple of launch script

```powershell
# Alias
Set-Alias vi nvim
Set-Alias ll ls
Set-Alias g git
Set-Alias grep findstr
Set-Alias tig $env:USERPROFILE\app\git\usr\bin\tig.exe
Set-Alias less $env:USERPROFILE\app\git\usr\bin\less.exe

# Prompt
Import-Module posh-git
$omp_config = "$env:USERPROFILE\sample.omp.json"
oh-my-posh --init --shell pwsh --config $omp_config | Invoke-Expression

# Terminal Icons
Import-Module -Name Terminal-Icons

# PSReadLine
Set-PSReadLineOption -EditMode Emacs
Set-PSReadLineOption -BellStyle None
Set-PSReadLineKeyHandler -Chord 'Ctrl+d' -Function DeleteChar
Set-PSReadLineOption -PredictionSource History

# Fzf
Import-Module PSFzf
Set-PsFzfOption -PSReadlineChordProvider 'Ctrl+f' -PSReadlineChordReverseHistory 'Ctrl+r'

# Customized functions 
# Add your functions here

```