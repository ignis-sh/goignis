# goignis

`goignis` is an optional _CLI_ for the [ignis](https://github.com/ignis-sh/ignis) widget framework.

## Installation

```sh
go install github.com/ignis-sh/goignis@latest
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
