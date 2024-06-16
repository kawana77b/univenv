# -------------------- Created By univenv --------------------

# env example
export RIPGREP_CONFIG_PATH=$HOME/.ripgreprc

# alias example
alias fd='fd-find'

# comment example
# The following changes the result depending on the specified shell

# source example (bash)
source $HOME/configs/foo.sh

# path example
export PATH=$HOME/go/bin:$PATH
# If you don't want to add a space above, just don't set the title, like this. Then add three line breaks at the bottom



# This is a raw item
# You can write anything here
# This is useful when you want to write a script

# command conditional example
type dotnet > /dev/null 2>&1 && export PATH=$HOME/.dotnet/tools:$PATH

# directory conditional example
[ -d $HOME/.cargo ] && export PATH=$HOME/.cargo/bin:$PATH

# if-command
if type docker > /dev/null 2>&1; then
    alias dc='docker compose'
    alias dcu='docker compose up'
    alias dcd='docker compose down'
fi

# if-directory
if [ -d $HOME/.local/bin ]; then
    export PATH=$HOME/.local/bin:$PATH
fi
# -------------------- End Of Created By univenv --------------------
