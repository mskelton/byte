# zet

CLI to create and manage Zettlekasten.

_Inspired by [rwxrob/zet](https://github.com/rwxrob/zet), but written from scratch because why not!_

## Installation

```bash
go install github.com/mskelton/zet@latest
```

## Usage

#### `id create`

Create a new Zettlekasten id. This is typically not used on it's own.

```bash
zet id create
```

### `completion`

To generate shell auto completion scripts, run the following command in your shell.

```bash
zet completion fish > ~/.config/fish/completions/zet.fish
```
