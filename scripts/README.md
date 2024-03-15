# Directory of small scripts I use to automate tasks
This directory contains a few scripts I have written to make daily tasks easier.
The source directory contains the source code for the scripts, to make editing and forking possible.

## Script list
- `note` - A script to create a new note in my inbox folder inside my Obsidian vault.

## Build scripts
When building the scripts in Go, pass the argument `-o '../../'` to build to this folder. Eg.: `cd source/note ; go build -o '../../' .`

## Install scripts
To install the scripts, you can use the `install.sh` script. This script will create symlinks to the scripts in the `~/usr/local/bin` directory.

**NB: Run with sudo, and make sure it is executable with `chmod +x install.sh`**
