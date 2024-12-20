## server_clipboard

A server which saves my clipboard (in memory), so I can share it between my devices.

This has both a CLI interface and a web interface -- I use [termux](https://termux.com/) on my phone to communicate with my server (with `server_clipboard <c|p>` (copy/paste)); On other devices I don't have a terminal on, this has a web interface at `/`:

<img src="https://github.com/purarue/server_clipboard/blob/main/frontend/demo.png" alt="screencap of server html page">

### Install

Install `golang` (requires `1.18`+)

You can clone and run `go build`, or:

```
go install -v "github.com/purarue/server_clipboard/cmd/server_clipboard@latest"
```

which downloads, builds and puts the binary on your `$GOBIN`.

### Usage

Run `server_clipboard server` on a remote server somewhere, update your `~/.bashrc`/`~/.zshenv` to include a password/remote address:

```
export CLIPBOARD_PASSWORD='i8nCzZnSY4hlHwUF9Ny15vqtPjfezpMHKZll0030Gn1p17Uiw7'
export CLIPBOARD_ADDRESS='http://mywebsite.com/clipboard'
```

```
NAME:
   server_clipboard - share clipboard between devices using a server

USAGE:
   server_clipboard [global options] command [command options] [arguments...]

COMMANDS:
   server, s  start server
   copy, c    copy to server clipboard
   paste, p   paste from server clipboard
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  port to listen on (default: 5025) [$CLIPBOARD_PORT]
   --password value        password to use [$CLIPBOARD_PASSWORD]
   --server_address value  server address to connect to (default: "localhost:5025") [$CLIPBOARD_ADDRESS]
   --help, -h              show help (default: false)
```

This automatically detects which operating system you're on and uses the corresponding clipboard commands, see [`clipboard.go`](clipboard.go). If this can't, set the `CLIPBOARD_COPY_COMMAND` and `CLIPBOARD_PASTE_COMMAND` environment variables (those commands should read/write from/to STDIN/STDOUT)

#### clear-after

If you want to clear the clipboard after a certain amount of time, you can use the `--clear-after` flag. For example, to clear the clipboard after 10 minutes:

```
server_clipboard server --clear-after 600
```
