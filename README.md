# byte

CLI tool to create bytes (mini posts) for my website.

## Installation

```bash
go install github.com/mskelton/byte@latest
```

## Usage

### `byte`

Create a new byte.

```bash
byte
```

### `byte list`

_Alias `ls`_

List all bytes that have been created.

```bash
byte list
```

## Shell completion

To generate shell auto completion scripts, run the following command in your shell.

**bash**

Add the following to your `~/.bash_profile`:

```bash
eval "$(byte completion bash)"
```

**zsh**

```bash
byte completion zsh > /usr/local/share/zsh/site-functions/_byte
```

**fish**

```bash
byte completion fish > ~/.config/fish/completions/byte.fish
```
