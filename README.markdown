# Teleprompt - remote assistance for your shell prompt
Teleprompt provides an easy way of adding simple status information
to the shell prompt, without elaborate shell configuration.


# Installation
Either manually place the `teleprompt` binary somewhere in `$PATH`, or use the Go
tools to install,

    # go get -u github.com/rkk/teleprompt

This requires the $GOPATH/bin to be present in $PATH.


# Building
Teleprompt is naively implemented in Go, with no external dependencies.  
Build it using the normal Go approach,

    # go build .


# Usage
For Bourne-compatible shells, activate Teleprompt in the PS1 variable like this,

    PS1="\$(teleprompt)"

Notice the prefixed `\` that causes the shell to evaluate the PS1 variable in
every invocation. Without this, the PS1 contents will be stale, e.g. not reflecting
change of directory.

For Zsh, use the `PROMPT` variable like this,

    PROMPT="\$(teleprompt)"

Other shells may utilize similar invocations. Add this to the respective shell
configuration files such as `~/.bashrc`, `~/.profile` or the like.  
Teleprompt is tested under Bourne-compatible shells in Linux and OpenBSD.


# Features
Run Teleprompt with the help flag `-h` to see a list of features and
how to use them.  
The following features are supported,

  - Exit code marker: Displayed if non-zero exit code from previous command
  - Git status: Displayed if the current directory is a Git checkout and has changes


# License
Teleprompt is released under the BSD license.
