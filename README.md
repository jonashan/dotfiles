# Dotfiles
These are my dotfiles which is build for using with zsh. 

# Setup
These are the steps that are required to setup the dotfiles on OSX and UNIX systems

## 01. Required software
- npm
- ripgrep
- nvim
- alacritty (Terminal Emulator)
- GNU stow (Used for handling the symlinks)
- tmp - Tmux Package Manager

### 02. Ruby gems
- Rubocop - For formatting and linting inside nvim
- Tmuxinator - For creating tmux sessions

### 03. Mason (nvim)
- Rubocop
- eslint_d
- prettier
- stylua
- html_lsp

## 04. Setup symlinks
- Clone repo to home folder
- `cd ~/dotfiles`
- `stow .` will create the symlinks to the parent folder (Home dir)

## 05. Setup Aliases
`echo ~/.aliases > ~/.zshrc`

## 06. Setup ENVvars
`echo ~/.envvars > ~(.zshrc`

## 07. Manual symlinks
- `ln -sf "$SECOND_BRAIN" ~/SecondBrain`



