# byte

CLI tool to create bytes (mini posts) for my website.

## Installation

You can install url by running the install script which will download
the [latest release](https://github.com/mskelton/byte/releases/latest).

```bash
curl -LSfs https://mskelton.dev/byte/install | sh
```

Or you can build from source.

```bash
git clone git@github.com:mskelton/byte.git
cd byte
go install .
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

### `tag list`

List all tags associated with bytes.

```bash
byte tag list
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
