# logpipe

A CLI tool for filtering and transforming log streams with pattern matching and output formatting.

## Installation

```bash
go install github.com/yourusername/logpipe@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/logpipe.git && cd logpipe && go build
```

## Usage

Pipe any log stream into `logpipe` to filter and format output:

```bash
# Filter logs by pattern
tail -f app.log | logpipe --filter "ERROR"

# Transform output format
cat service.log | logpipe --format json --filter "WARN|ERROR"

# Highlight matches and suppress non-matching lines
journalctl -f | logpipe --filter "panic" --highlight --quiet
```

### Flags

| Flag | Description |
|------|-------------|
| `--filter` | Regex pattern to match log lines |
| `--format` | Output format: `text`, `json`, `csv` |
| `--highlight` | Highlight matched patterns in output |
| `--quiet` | Suppress non-matching lines |
| `--timestamp` | Prepend timestamps to each line |

## Example

```bash
$ echo -e "INFO starting\nERROR failed to connect\nINFO done" | logpipe --filter "ERROR"
ERROR failed to connect
```

## Contributing

Pull requests are welcome. Please open an issue first to discuss any major changes.

## License

MIT © 2024 yourusername