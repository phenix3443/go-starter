#!/bin/bash

set -e

GIT_COMPLETION_PATH="$HOME/.git-completion.bash"

echo "Checking if git-completion.bash is already available..."
if [ -f "/usr/share/bash-completion/completions/git" ]; then
	GIT_COMPLETION_PATH="/usr/share/bash-completion/completions/git"
elif [ -f "/etc/bash_completion.d/git-completion.bash" ]; then
	GIT_COMPLETION_PATH="/etc/bash_completion.d/git-completion.bash"
elif [ -f "/usr/local/etc/bash_completion.d/git-completion.bash" ]; then
	GIT_COMPLETION_PATH="/usr/local/etc/bash_completion.d/git-completion.bash"
elif [ -f "/usr/local/share/bash-completion/git-completion.bash" ]; then
	GIT_COMPLETION_PATH="/usr/local/share/bash-completion/git-completion.bash"
else
	echo "git-completion.bash not found, downloading..."
	curl -o "$HOME/.git-completion.bash" https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash
	GIT_COMPLETION_PATH="$HOME/.git-completion.bash"
fi

echo "Checking if bash-completion is installed..."
if ! command -v complete &>/dev/null; then
	echo "bash-completion not found, installing..."
	if [[ "$OSTYPE" == "linux-gnu"* ]]; then
		if command -v apt &>/dev/null; then
			sudo apt update && sudo apt install -y bash-completion
		elif command -v yum &>/dev/null; then
			sudo yum install -y bash-completion
		elif command -v dnf &>/dev/null; then
			sudo dnf install -y bash-completion
		fi
	elif [[ "$OSTYPE" == "darwin"* ]]; then
		if command -v brew &>/dev/null; then
			brew install bash-completion
			echo '[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"' >>~/.bashrc
		else
			echo "Homebrew not found. Please install Homebrew and rerun this script."
			exit 1
		fi
	fi
fi

echo "Configuring ~/.bashrc..."
if ! grep -q "source $GIT_COMPLETION_PATH" ~/.bashrc; then
	echo "source $GIT_COMPLETION_PATH" >>~/.bashrc
fi

echo "Adding Git alias completion..."
if ! grep -q '__git_complete gco _git_checkout' ~/.bashrc; then
	echo "__git_complete gco _git_checkout" >>~/.bashrc
	echo "alias gco='git checkout'" >>~/.bashrc
fi

echo "Reloading Bash configuration..."
source ~/.bashrc

echo "Git command completion setup complete! 🚀"
