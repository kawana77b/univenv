# -------------------- Created By univenv --------------------

# env example
$env:RIPGREP_CONFIG_PATH = "$HOME/.ripgreprc"

# alias example
Set-Alias -Name fd -Value "fd-find"

# comment example
# The following changes the result depending on the specified shell

# source example (pwsh)
. "$HOME/configs/foo.ps1"

# path example
$env:PATH = "$HOME/go/bin:$env:PATH"
# If you don't want to add a space above, just don't set the title, like this. Then add three line breaks at the bottom



# This is a raw item
# You can write anything here
# This is useful when you want to write a script

# command conditional example
if (Get-Command dotnet -ErrorAction SilentlyContinue) { $env:PATH = "$HOME/.dotnet/tools:$env:PATH" }

# directory conditional example
if (Test-Path $HOME/.cargo) { $env:PATH = "$HOME/.cargo/bin:$env:PATH" }

# if-command
if (Get-Command docker -ErrorAction SilentlyContinue) {
    Set-Alias -Name dc -Value "docker compose"
    Set-Alias -Name dcu -Value "docker compose up"
    Set-Alias -Name dcd -Value "docker compose down"
}

# if-directory
if (Test-Path $HOME/.local/bin) {
    $env:PATH = "$HOME/.local/bin:$env:PATH"
}
# -------------------- End Of Created By univenv --------------------
