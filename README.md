# Dotfiles
These are my dotfiles which is build for using with zsh. 

# Setup
These are the steps that are required to setup the dotfiles on OSX and UNIX systems

## Required software
- npm
- ripgrep
- nvim
- alacritty (Terminal Emulator)
- GNU stow (Used for handling the symlinks)

### Ruby gems
- Rubocop - For formatting and linting inside nvim

### Mason (nvim)
- Rubocop
- eslint_d
- prettier
- stylua

## Setup symlinks
- Clone repo to home folder
- `cd ~/dotfiles`
- `stow .` will create the symlinks to the parent folder (Home dir)

## Setup Aliases
`echo ~/.aliases > .zshrc`


