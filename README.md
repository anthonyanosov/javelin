# Satellite 🛰️

<img src="/assets/satellite.png" alt="satellite" width="200" />

Satellite is a lightweight Neovim plugin + Go CLI for showing Go function
cyclomatic complexity inline, directly in your editor.

It analyzes the current Go buffer and renders `🛰️ <complexity>` at end-of-line
for each function declaration.

## Features

- Inline complexity hints using Neovim virtual text
- On-demand analysis via command or keymap
- LazyVim-friendly setup

## How It Works

Satellite has two layers:

- `sat` CLI (`cmd/sat/main.go`) parses a Go file and returns JSON metrics.
- Neovim Lua module (`lua/satellite.lua`) runs the CLI and draws extmarks.

Complexity starts at `1` and increases for:

- `if`
- `for`
- `range`
- `case` clauses in `switch`
- logical `&&` and `||`

Here's how satellite 🛰️ looks in Neovim when running `<leader>sc`:

<img src="/assets/demo.png" alt="demo" width="500"/>

## Requirements

- Go `1.22+`
- Neovim `0.9+`
- `sat` available on your `PATH`

## Installation

### 1) Install the CLI

From the project root:

```bash
go install ./cmd/sat
```

Ensure your Go bin directory is on `PATH` (commonly `$HOME/go/bin`):

```bash
export PATH="$HOME/go/bin:$PATH"
```

### 2) Add plugin in LazyVim

For LazyVim users, this is often under `~/.config/nvim/lua/plugins/`, make a new file `satellite.lua`:

```lua
return {
  {
    "anthonyanosov/satellite",
    config = function()
      local satellite = require("satellite")
      vim.keymap.set("n", "<leader>sc", satellite.show_complexity, { desc = "Satellite: Show complexity" })
    end,
  },
}
```

## Usage

- `:Satellite` - Analyze current Go buffer and show inline complexity
- `:SatelliteClear` - Clear all Satellite virtual text in current buffer
- Optional keymap: `<leader>sc` (from config above)

## CLI Usage

```bash
sat -src ./path/to/file.go
```

Example output:

```json
[{"Name":"process","Complexity":4,"StartLine":10,"EndLine":37}]
```

## Roadmap

- Auto-refresh on `BufWritePost` for Go buffers
- Package-level summary view for `sat`
- Explicit flag to include/exclude anonymous functions
- Maybe a custom parser to generate ASTs? 🧪

## License

MIT - see `LICENSE`.
