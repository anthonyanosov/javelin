# Satellite 🛰️

Satellite is a lightweight Neovim plugin + Go CLI for showing Go function
cyclomatic complexity inline, directly in your editor.

It analyzes the current Go buffer and renders `🛰️ <complexity>` at end-of-line
for each function declaration.

## Features

- Inline complexity hints using Neovim virtual text
- On-demand analysis via command or keymap
- Simple CLI interface for scripting and debugging
- LazyVim-friendly setup

## How It Works

Satellite has two layers:

- `sat` CLI (`cmd/sat/main.go`) parses a Go file and returns JSON metrics.
- Neovim Lua module (`lua/satellite.lua`) runs the CLI and draws extmarks.

Complexity starts at `1` and increases for:

- `if`, `for`, and `range`
- `case` clauses in `switch`
- logical `&&` and `||`

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

If using a local checkout:

```lua
return {
  {
    dir = "~/path/to/satellite",
    config = function()
      local satellite = require("satellite")
      vim.keymap.set("n", "<leader>sc", satellite.show_complexity, { desc = "Satellite: Show complexity" })
    end,
  },
}
```

If installed from GitHub:

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

## Migration from Javelin

Backwards compatibility is currently included:

- `:Javelin` and `:JavelinClear` still work as aliases.
- `require("javelin")` still works as an alias to `require("satellite")`.
- `cmd/javelin/main.go` remains available for transition.

New installs should use `satellite` names and paths.

`satellite` CLI remains available as a compatibility alias.

## Troubleshooting

- `Satellite: failed to run analyzer command`
  - `sat` is not on `PATH` for your Neovim process.
- `Satellite: failed to parse JSON output: ...`
  - The CLI returned an error string. Run `sat -src <file.go>` in terminal.
- `Satellite: current buffer is not a Go file`
  - Switch to a `.go` buffer first.

## Project Structure

- `cmd/sat/main.go` - primary CLI entrypoint
- `cmd/satellite/main.go` - compatibility CLI entrypoint
- `pkg/` - Go parser and complexity analysis
- `lua/satellite.lua` - Neovim integration API
- `plugin/satellite.lua` - auto-registered Neovim user commands

## Roadmap

- Auto-refresh on `BufWritePost` for Go buffers
- Configurable highlight group and icon
- Package-level summary view

## License

MIT - see `LICENSE`.
