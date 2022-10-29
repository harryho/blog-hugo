+++
title = "NeoVim"
description = "Empower the vi with NeoVimn"
+++

__This article will show you how to empower your vi with NeoVim.__

## NeoVim

> [Neovim](https://github.com/neovim/neovim) is a refactor, and sometimes redactor, in the tradition of Vim (which itself derives from Stevie). It is not a rewrite but a continuation and extension of Vim. Many clones and derivatives exist, some very clever—but none are Vim. Neovim is built for users who want the good parts of Vim, and more.


### Install Neovim

Please follow the installation instructions [here](https://github.com/neovim/neovim/wiki/Installing-Neovim) 

### Vim Plugin

Please follow the installation instructions from [Vim-Plug's README](https://github.com/neovim/neovim)https://github.com/junegunn/vim-plug


### Init


- Create a file `init.vim`

```lua
:set number
:set relativenumber
:set autoindent
:set tabstop=4
:set shiftwidth=4
:set smarttab
:set softtabstop=4
:set mouse=a

call plug#begin()

Plug 'http://github.com/tpope/vim-surround' " Surrounding ysw)
Plug 'https://github.com/preservim/nerdtree' " NerdTree
Plug 'https://github.com/tpope/vim-commentary' " For Commenting gcc & gc
Plug 'https://github.com/vim-airline/vim-airline' " Status bar
Plug 'https://github.com/lifepillar/pgsql.vim' " PSQL Pluging needs :SQLSetType pgsql.vim
Plug 'https://github.com/ap/vim-css-color' " CSS Color Preview
Plug 'https://github.com/rafi/awesome-vim-colorschemes' " Retro Scheme
Plug 'https://github.com/neoclide/coc.nvim'  " Auto Completion
Plug 'https://github.com/ryanoasis/vim-devicons' " Developer Icons
Plug 'https://github.com/tc50cal/vim-terminal' " Vim Terminal
Plug 'https://github.com/preservim/tagbar' " Tagbar for code navigation
Plug 'https://github.com/terryma/vim-multiple-cursors' " CTRL + N for multiple cursors
Plug 'https://github.com/EdenEast/nightfox.nvim' " A highly customizable theme
Plug 'ctrlpvim/ctrlp.vim'
Plug 'kassio/neoterm'



set encoding=UTF-8

call plug#end()

nnoremap <C-f> :NERDTreeFocus<CR>
nnoremap <C-n> :NERDTree<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
nnoremap <C-l> :call CocActionAsync('jumpDefinition')<CR>

nmap <F8> :TagbarToggle<CR>

:set completeopt-=preview " For No Previews

:colorscheme nightfox

let g:NERDTreeDirArrowExpandable="+"
let g:NERDTreeDirArrowCollapsible="~"

" --- Just Some Notes ---
" :PlugClean :PlugInstall :UpdateRemotePlugins
"
" :CocInstall coc-python
" :CocInstall coc-clangd
" :CocInstall coc-snippets
" :CocCommand snippets.edit... FOR EACH FILE TYPE

" air-line
let g:airline_powerline_fonts = 1

if !exists('g:airline_symbols')
    let g:airline_symbols = {}
endif

" airline symbols
let g:airline_left_sep = ''
let g:airline_left_alt_sep = ''
let g:airline_right_sep = ''
let g:airline_right_alt_sep = ''
let g:airline_symbols.branch = ''
let g:airline_symbols.readonly = ''
let g:airline_symbols.linenr = ''

inoremap <expr> <Tab> pumvisible() ? coc#_select_confirm() : "<Tab>"

```


- Install plugins

  - Open nvim
  - Enter command mode and enter `PlugInstall`





### Trouble shooting

#### Fonts

- Mac
  - Install Nerd fonts 
    - [Follow the instruction to install Powerline Nerd Font](https://webinstall.dev/nerdfont/)  
  - Download other fonts, e.g. Caskaydia Cove Nerd Font 
```
cd ~/Library/Fonts
open . 
# Copy font files to fonts folder
# Caskaydia Cove Nerd Font Complete Regular.otf
# Caskaydia Cove Nerd Font Complete Mono Regular.otf
```

#### coc.nvim 

- Windows pwsh

```powershell
cd <coc.nvim_folder>
sudo nvm use 14.19.3
npm install -g yarn 
yarn install

```


- mac 

```
cd ~/.local/share/nvim/plugged/coc.nvim
yarn install
```