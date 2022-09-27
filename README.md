# zet

CLI to create and manage Zettelkasten.

_Inspired by [rwxrob/zet](https://github.com/rwxrob/zet), but written from scratch because why not!_

## Installation

```bash
go install github.com/mskelton/zet@latest
```

## Usage

### `zet`

Create a new zettel.

```bash
zet
```

### `zet id`

Create a new zettel id. This is typically not used on it's own.

```bash
zet id create
```

### Shell completion

To generate shell auto completion scripts, run the following command in your shell.

```bash
zet completion fish > ~/.config/fish/completions/zet.fish
```
