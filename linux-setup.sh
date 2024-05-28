## This is a install scripts for Linux machines
## It will setup all the necessary tools for development and also setup dotfiles
## This can be run before cloning the dotfiles repository

# Updating and setting up system tools and dependencies
sudo apt update -y
sudo apt install -y \
	git curl docker.io build-essential pkg-config stow \
	tmux tmuxinator snapd \
	libreadline-dev zlib1g-dev libyaml-dev libreadline-dev libncurses5-dev libffi-dev libgdbm-dev libjemalloc2 \
	libvips imagemagick libmagickwand-dev \
	libssl-dev clang \
	imagemagick libmagickwand-dev \
	stow fzf eza btop

# Clone dotfiles
echo "Cloning dotfiles"
git clone https://github.com/jonashan/dotfiles.git ~/dotfiles

# Install TPM for Tmux
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm

# Stowing the dotfiles (creates symlink)
echo "Stowing dotfiles"
cd ~/dotfiles
stow .
cd ~/

echo "Installing Ruby and Nodejs"
# asdf - Ruby and Node.js
curl https://mise.run | sh
echo 'eval "$(~/.local/bin/mise activate bash)"' >>~/.bashrc
eval "$(~/.local/bin/mise activate bash)"
source ~/.bashrc
mise install

# Setting up aliases
echo "Setting up aliases"
cat <<EOF >>~/.bashrc
# Aliases
source ~/.aliases
EOF
source ~/.bashrc

# Setup neovim
echo "Setting up neovim"
curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz
sudo rm -rf /opt/nvim
sudo tar -C /opt -xzf nvim-linux64.tar.gz
echo 'export PATH="$PATH:/opt/nvim-linux64/bin"' >>~/.bashrc
source ~/.bashrc
git clone https://github.com/LazyVim/starter ~/.config/nvim
source ~/.bashrc

# Setup LazyGit
echo "Setting up LazyGit"
LAZYGIT_VERSION=$(curl -s "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | grep -Po '"tag_name": "v\K[^"]*')
curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/latest/download/lazygit_${LAZYGIT_VERSION}_Linux_x86_64.tar.gz"
tar xf lazygit.tar.gz lazygit
sudo install lazygit /usr/local/bin

# Setting up tmuxinator and rubocop
gem install tmuxinator rubocop

# Starting docker containers
echo "Setup docker user and start the containers"
sudo usermod -aG docker ${USER}
# Needs to sudo at the first run because usermod is not effective before reboot
sudo docker run -d --restart unless-stopped -p 5432:5432 --name=postgres -e POSTGRES_PASSWORD=postgres postgres
sudo docker run -d --restart unless-stopped -p 6379:6379 --name=redis redis
# docker run -d --restart unless-stopped -p 3306:3306 --name=mysql5.7 -e MYSQL_ROOT_PASSWORD= -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql:5.7

# Install Starship
echo "Installing Starship"
sudo snap install starship --edge
echo "Launching starship"
echo 'eval "$(starship init bash)"' >>~/.bashrc
eval "$(starship init bash)"
source ~/.bashrc
