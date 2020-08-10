# Generation of passwords
## Usage
```bash
$ nohi pwd -h
Generate password

Usage:
  nohi pwd [flags]

Flags:
  -d, --digit         enable digit characters (default true)
  -h, --help          help for pwd
  -l, --length uint   password length (default 32)
  -s, --special       enable special characters
  -u, --upper         enable upper characters

Global Flags:
  -n, --count uint      elements to generate (default 1)
  -o, --output string   output file
  -w, --workers uint    generation concurrency (default 4)
```

## Examples
### Single ID
```bash
$ nohi pwd
gwu80hcmiq4fefhfi0sjuj2h7z4xzk5u

Done
```

### Disable digits, enable upper-case letters, and change length
```bash
$ nohi pwd -d=false -u -l 20
zrPvbcHdeRqEfPwisVhH

Done
```
