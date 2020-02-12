# Teleprompt - remote assistance for your shell prompt

# Installation
Either place the `teleprompt` binary somewhere in `$PATH`, or use the Go
tools to install,

    # go get -u github.com/rkk/teleprompt

This requires the $GOPATH/bin to be present in $PATH.


# Usage
For Bourne-compatible shells, activate Teleprompt in the PS1 variable like this,

    PS1="\$(teleprompt)"

Notice the prefixed `\` that causes the shell to evaluate the PS1 variable upon
every invocation. Without this, the PS1 contents will be stale, e.g. not reflecting
change of directory.

For Zsh, use the `PROMPT` variable like this,

    PROMPT="\$(teleprompt)"

Other shells may utilize similar invocations.
Teleprompt is tested under Bourne-compatible shells in Linux and OpenBSD.


# License
Teleprompt is released under the BSD license.
