# Teleprompt - remote assistance for your $PS1

# Usage
For Bourne-compatible shells, activate Teleprompt in the PS1 variable like this,

    PS1="\$(teleprompt)"

Notice the prefixed `\` that causes the shell to evaluate the PS1 variable upon
every invocation. Without this, the PS1 contents will be stale, e.g. not reflecting
change of directory or Git branches.

For Zsh, use the `PROMPT` variable like this,

    PROMPT="\$(teleprompt)"

Other shells may utilize similar invocations.
