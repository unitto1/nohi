# Generation of UUIDs (Version 4)
## Usage
```bash
$ nohi uuid -h
Generate UUID (version 4)

Usage:
  nohi uuid [flags]

Flags:
  -h, --help   help for uuid

Global Flags:
  -n, --count uint      elements to generate (default 1)
  -o, --output string   output file
  -w, --workers uint    generation concurrency (default 4)
```

## Examples
### Single ID
```bash
$ nohi uuid
9b102d74-4802-4efc-b9ce-76bd2db4bce4

Done
```
