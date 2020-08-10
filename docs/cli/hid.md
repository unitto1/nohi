# Generation of human-friendly IDs
## Usage
```bash
$ nohi hid --help
Generate human-friendly ID

Usage:
  nohi hid [flags]

Flags:
  -d, --digit              enable digits in suffix (default true)
  -h, --help               help for hid
  -l, --length uint        suffix length (default 5)
  -s, --separator string   separator characters (default "_")

Global Flags:
  -n, --count uint      elements to generate (default 1)
  -o, --output string   output file
  -w, --workers uint    generation concurrency (default 4)
```

## Examples
### Single ID
```bash
$ nohi hid
priceless_lovelace_a6yt7

Done
```

### Disable digits, change separator and length
```bash
$ nohi hid -s - -d=false -l 10
interesting-villani-ouusblfpca

Done
```
