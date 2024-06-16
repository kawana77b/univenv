# univenv

## Description

`univenv` is a tool for universal shell platforms that allows you to output your `yml` as `bash`, `fish` or `powershell`.

- Managed in a common `yml` to absorb OS and shell differences
- Multiple `yml` files can be managed and read as you wish
- Conditioning of OS, commands and shells.
- They will be reflected in the configuration order you wrote them in
- Run in a shell configuration file, which you can load whenever you want to start working with it

## Motivation

Each shell has its own syntax.
If you can work only with your favorite shell, no problem, but if you work in many different places, you have to consider the differences between operating systems and shells.
For example, if you are suddenly faced with a new environment, be it Windows or Mac, it will fluctuate, and there will be situations where you are comfortable with Bash but have to deal with it in PowerShell.

Do I have to maintain my own `.bashrc`, `config.fish`, and `Profile.ps1` and translate and rewrite the `PATH` and `alias` in all of them? It was very complicated for me!

This tool does a simple language translation and makes it easy to accommodate each of them.

## Install

This tool assumes Windows, Mac OS, and Linux as the operating system to be used.  
Binaries are available on the [Release page](https://github.com/kawana77b/univenv/releases).

If you have a Go execution environment, the following is also available.

```bash
go install github.com/kawana77b/univenv@latest
```

## Getting Started

Create folders and files.

```bash
mkdir -p ~/.config/univenv && touch ~/.config/univenv/config.yml
```

Edit File by Your Editor.

```bash
vim ~/.config/univenv/config.yml
```

Paste the following yaml.

```yaml
items:
  - type: alias
    name: pip
    value: pip3
```

Let's run it. The argument shell name can be `bash`, `fish`, or `pwsh`.

```bash
univenv bash
```

If you see the following, you have succeeded! Comments can also be hidden with `--no-comment/-n`.

```bash
# -------------------- Created By univenv --------------------
alias pip='pip3'
# -------------------- End Of Created By univenv --------------------
```

You can check the default status at `univenv status`

### Output script

Output and load the script as follows.

```bash
univenv bash > script.sh
```

### Specify a file

To change the referenced file:

```bash
univenv bash --file ~/foo/baz/config.yml
```

### Called at shell startup

If you wish, this script can be set to run in the configuration file as is.

#### bash

Add the following to `~/.bashrc`

```bash
eval "$(univenv bash)"
```

#### fish

Add the following to `~/.config/fish/config.fish`

```fish
univenv fish | source
```

#### powershell

Add the following to your `$PROFILE`

```powershell
Invoke-Expression (& {
    (univenv pwsh | Out-String)
})
```

## Configuration file location and target

- Directory is the default, `~/.config/univenv`. This is common regardless of OS. However, you can change the directory to read by setting the environment variable `$UNIVENV_CONFIG_DIR` before using the tool.
- The file name is fixed to `config.yml` unless you try to load it using the `--file/-f` option.
- The `--target/-t` option allows you to switch which `yml` to use. That is, you can have more than one file to place in `~/.config/univenv`. For example, if you do not specify a target, it is `config.yml`, but if you do `--target homepc`, for example, it will try to read `config.homepc.yml`

> [!NOTE]
> In `pwsh`, if `~` is prefixed, the value is replaced by `$HOME`. Also, paths or commas are cleaned in some configurations.
> This is a specification.

## Configuration

You can see concrete examples on the [Examples Page](https://github.com/kawana77b/univenv/tree/main/examples)

The definition of yml is as follows:

```yaml
items:
  - title: string # If specified, a comment will be added to this item. Also, a new line will be added above
    type: env | path | alias | comment | source | raw | if-command | if-directory # Select the type of this item
    name: string # Specify key names for environment variables and aliases
    value: string # Various value. This is required
    directory: string # Directory location. If specified, adds a script that checks if the location exists.
    command: string # Command Name. If specified, adds a script that checks for the presence of the command.
    shell: [bash | fish | pwsh] # If a shell name is specified, the script will be output only for the shell selected in the argument
    os: [windows | darwin | linux] # If an OS name is specified, the script will be output only when the running OS corresponds to it.
    lf: number # Add a line break with a description of 1 or more
    disabled: bool # Specify with true to suppress output of this item
    items: # Nests items when if-** is specified in type
```

### If Nested

`type: if-command` etc. can create if-branches.

```yaml
- title: Docker
  type: if-command
  value: docker
  items:
    - type: alias
      name: dc
      value: docker compose
    - type: alias
      name: dcu
      value: docker compose up
```

> [!NOTE]
> Only one hierarchy can be nested.

### Raw Script

You can add raw scripts by specifying `type: raw`.

```yaml
items:
  - title: test func
    type: raw
    value: |-
      function foo {
        echo "foo"
      }
    shell: [bash]
```

## NOTE

Bug reports, etc. can be sent to Issue. I would be very happy to receive your feedback or a star if you like it!  
Due to my current circumstances, I may not be able to answer pull requests or other replies. Please use the fork. However, I would be happy to see some sort of indication that I am the original author.
