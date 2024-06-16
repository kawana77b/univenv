# -------------------- Created By univenv --------------------

# env example
set -gx RIPGREP_CONFIG_PATH $HOME/.ripgreprc

# alias example
alias fd='fd-find'

# comment example
# The following changes the result depending on the specified shell

# source example (fish)
source $HOME/configs/foo.fish

# path example
set -gx PATH $HOME/go/bin $PATH
# If you don't want to add a space above, just don't set the title, like this. Then add three line breaks at the bottom



# This is a raw item
# You can write anything here
# This is useful when you want to write a script

# command conditional example
type -q dotnet; and set -gx PATH $HOME/.dotnet/tools $PATH

# directory conditional example
test -d $HOME/.cargo; and set -gx PATH $HOME/.cargo/bin $PATH

# if-command
if type -q docker
    alias dc='docker compose'
    alias dcu='docker compose up'
    alias dcd='docker compose down'
end

# if-directory
if test -d $HOME/.local/bin
    set -gx PATH $HOME/.local/bin $PATH
end
# -------------------- End Of Created By univenv --------------------
