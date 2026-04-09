# Javelin 🛰️

Javelin is a lightweight Neovim plugin + Go CLI for showing Go function
cyclomatic complexity inline, directly in your editor.

It analyzes the current Go buffer and renders `🛰️ <complexity>` at end-of-line
for each function declaration.

## Features

- Inline complexity hints using Neovim virtual text
- On-demand analysis via command or keymap
- Simple CLI interface for scripting and debugging
- LazyVim-friendly setup

## How It Works

Javelin has two layers:

- `javelin` CLI (`cmd/javelin/main.go`) parses a Go file and returns JSON metrics.
- Neovim Lua module (`lua/javelin.lua`) runs the CLI and draws extmarks.

Complexity starts at `1` and increases for:

- `if`, `for`, and `range`
- `case` clauses in `switch`
- logical `&&` and `||`

## Requirements

- Go `1.22+`
- Neovim `0.9+`
- `javelin` available on your `PATH`

## Installation

### 1) Install the CLI

From the project root:

```bash
go install ./cmd/javelin
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
    dir = "~/path/to/javelin",
    config = function()
      local javelin = require("javelin")
      vim.keymap.set("n", "<leader>jc", javelin.show_complexity, { desc = "Javelin: Show complexity" })
    end,
  },
}
```

If installed from GitHub:

```lua
return {
  {
    "anthonyanosov/javelin",
    config = function()
      local javelin = require("javelin")
      vim.keymap.set("n", "<leader>jc", javelin.show_complexity, { desc = "Javelin: Show complexity" })
    end,
  },
}
```

## Usage

- `:Javelin` - Analyze current Go buffer and show inline complexity
- `:JavelinClear` - Clear all Javelin virtual text in current buffer
- Optional keymap: `<leader>jc` (from config above)

## CLI Usage

```bash
javelin -src ./path/to/file.go
```

Example output:

```json
[{"Name":"process","Complexity":4,"StartLine":10,"EndLine":37}]
```

## Troubleshooting

- `Javelin: failed to run analyzer command`
  - `javelin` is not on `PATH` for your Neovim process.
- `Javelin: failed to parse JSON output: ...`
  - The CLI returned an error string. Run `javelin -src <file.go>` in terminal.
- `Javelin: current buffer is not a Go file`
  - Switch to a `.go` buffer first.

## Project Structure

- `cmd/javelin/main.go` - CLI entrypoint
- `pkg/` - Go parser and complexity analysis
- `lua/javelin.lua` - Neovim integration API
- `plugin/javelin.lua` - auto-registered Neovim user commands

## Roadmap

- Auto-refresh on `BufWritePost` for Go buffers
- Configurable highlight group and icon
- Package-level summary view
- Custom AST parser? 🧪

## License

MIT - see `LICENSE`.
