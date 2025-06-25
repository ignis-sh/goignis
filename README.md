# goignis

`goignis` is an optional _CLI_ for the [ignis](https://github.com/ignis-sh/ignis) widget framework.

## Installation

### Prerequisites

Install the latest [go tool chain](https://go.dev/dl/).

Or install with a package manager:

```sh
pacman -S go
```

### Quick start

Install with _Go_ tool chain:

```sh
# install into $GOPATH/bin, usually ~/go/bin
go install github.com/ignis-sh/goignis@latest
```

Optionally, manage packages with [gup](https://github.com/nao1215/gup):

```sh
# install gup
go install github.com/nao1215/gup@latest
# manage packages
gup list
gup check
gup update
```

### Shell completions

Bash:

```sh
# prerequisite: install bash-completion
pacman -S bash-completion
# current shell session only
source <(goignis completion bash)
# install permanently
mkdir -p ~/.local/share/bash-completion/completions/
goignis completion bash >~/.local/share/bash-completion/completions/goignis.bash
```

Fish:

```sh
# current shell session only
goignis completion fish | source
# install permanently
mkdir -p ~/.config/fish/completions/
goignis completion fish >~/.config/fish/completions/goignis.fish
```

### Build manually

Clone and build:

```sh
git clone https://github.com/ignis-sh/goignis.git
cd goignis
go build
```

## Usage

```txt
An optional CLI for ignis

Usage:
  goignis [command]

Available Commands:
  close-window  Close a window
  completion    Generate the autocompletion script for the specified shell
  help          Help about any command
  init          Initialize Ignis
  inspector     Open GTK Inspector
  list-windows  List names of all windows
  open-window   Open a window
  quit          Quit Ignis
  reload        Reload Ignis
  systeminfo    Print system information
  toggle-window Toggle a window

Flags:
  -h, --help   help for goignis
  -j, --json   Print results in json

Use "goignis [command] --help" for more information about a command.
```

Read the manual:

```sh
goignis help
goignis help init
goignis list-windows -h
```

Initialize an _ignis_ instance:

```sh
goignis init -d
```

List all _ignis_ windows:

```sh
goignis list-windows -j | jq
```

## Development

Install [gopls](https://github.com/golang/tools/blob/master/gopls/README.md), the _Go_ language server:

```sh
pacman -S gopls
```

### Contribution

- Use _gopls_ or _gofumpt_ formatters.
- Write docs and add examples or unit tests if new features are introduced.
- Keep package `pkg` and `cmd` their own roles:
  - Package `pkg` provides exported APIs for third-party derivations.
  - Package `cmd` provides APIs for customizing subcommands.

```sh
# clone and install dependencies
git clone https://github.com/ignis-sh/goignis.git
cd goignis
go mod tidy
# start coding here
nvim
```

### Derivation

- Import `github.com/ignis-sh/goignis/pkg` if you're willing to make use of those APIs.
- Import `github.com/ignis-sh/goignis/cmd` if you want to add custom subcommands.
- If your features can benefit other users, please consider contributing here :).

```sh
# create a new project and add goignis as a dependency
mkdir myowncli && cd myowncli
go mod init github.com/username/myowncli
go get github.com/ignis-sh/goignis
# then you can write your features with APIs from goignis
nvim main.go
```
