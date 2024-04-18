## This is a install scripts for Linux machines
## It will setup all the necessary tools for development and also setup dotfiles
## This can be run before cloning the dotfiles repository

# Update
sudo apt update -y
sudo apt install -y \
	git curl docker.io build-essential stow \
	tmux tmuxinator neovim snapd

# Clone dotfiles
git clone https://github.com/jonashan/dotfiles.git ~/dotfiles

# Stowing the dotfiles (creates symlink)
cd ~/dotfiles
stow .
cd ~/

# Install Starship
sudo snap install starship --edge

# asdf - Ruby and Node.js
git clone https://github.com/asdf-vm/asdf.git ~/.asdf
echo '. "$HOME/.asdf/asdf.sh"' >>~/.bashrc
echo '. "$HOME/.asdf/completions/asdf.bash"' >>~/.bashrc
source ~/.bashrc

# Setting up bashrc
cat <<EOF >>~/.bashrc
# Aliases
source ~/.aliases
EOF
echo 'eval "$(starship init bash)"' >>~/.bashrc
source ~/.bashrc

# Setup LazyGit
LAZYGIT_VERSION=$(curl -s "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | grep -Po '"tag_name": "v\K[^"]*')
curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/latest/download/lazygit_${LAZYGIT_VERSION}_Linux_x86_64.tar.gz"
tar xf lazygit.tar.gz lazygit
sudo install lazygit /usr/local/bin
