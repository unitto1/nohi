# Command-line interface
## Usage
```bash
$ nohi -h
Fast command-line tool and library for generation human ids, uuids and passwords.

Usage:
  nohi [command]

Available Commands:
  completion  Generate auto-completion script
  help        Help about any command
  hid         Generate human-friendly ID
  pwd         Generate password
  uuid        Generate UUID (version 4)
  version     Print the version of NOHI

Flags:
  -n, --count uint      elements to generate (default 1)
  -h, --help            help for nohi
  -o, --output string   output file
  -w, --workers uint    generation concurrency (default 4)

Use "nohi [command] --help" for more information about a command.
```

## Auto-completion
```bash
$ nohi completion -h
To load auto-completion:

Bash:
Linux:
$ nohi completion bash > /etc/bash_completion.d/nohi
MacOS:
$ nohi completion bash > /usr/local/etc/bash_completion.d/nohi

Zsh:
$ nohi completion zsh > "${fpath[1]}/nohi"

Fish:
$ nohi completion fish > ~/.config/fish/completions/nohi.fish

Usage:
  nohi completion [bash|zsh|fish]

Flags:
  -h, --help   help for completion

Global Flags:
  -n, --count uint      elements to generate (default 1)
  -o, --output string   output file
  -w, --workers uint    generation concurrency (default 4)
```

## Getting Version
```bash
$ nohi version
Version:        undefined
Commit:         57bcdfc
```

## Examples
### Single element
```bash
$ nohi pwd
lrp5jj7v4bae3m81ljd2eyt4te2xfq2y

Done
```

### Multiple elements
```bash
$ nohi uuid -n 10
f63ecc26-99dd-461e-8d80-1c96a2209db9
927cb1bc-d279-4dbb-9403-fc7d79bf5e65
4c9ecd89-0382-4d88-9b6f-4c12de9b601f
ebf1f395-0958-4c8b-8180-90474024df3d
ea4f5c01-2a4e-4277-a378-46f3fd45716b
b7a640a3-16b8-436f-ace3-9a1ff6fe4b47
87feab95-013b-44bd-a408-33a221f26171
c2d54c1f-eac9-4d9c-829a-bb27e788efef
2acc8e19-842a-4f4d-9213-4d5a885ff73a
f30a4219-7f85-4877-b8cc-c65d46e61ab3

Done
```

### Store to file
```bash
$ nohi hid -n 5 -o hid.txt

Done
$ cat hid.txt 
gifted_dhawan_gb9cx
priceless_engelbart_yqhrx
sweet_rhodes_4p0vg
hardcore_bose_0dtgy
trusting_montalcini_65zw1
```

### Change concurrency
> **NOTE:** by default workers count equals to num of vCPU, but sometimes you can
> get better performance via changing this value (ex. to count of real CPUs)

```bash
$ nohi pwd -n 7 -w 32
w3ob6j7skx7jxl2hajpcwd034964s20p
i6ucx1l539ufhdi4p02pmwit6bnp632n
lrrye7wgbg4j64r1sebi75qivg4440j0
10r588hjfn9ofp7djt61ghrs0gzg1l81
41mf0401r9hk4gji9a1lhbmyqt5wujm6
4nk1ytnae31wjmucjq1g3eilnw7n0i92
o8xaxbasqp0njez3zmstjrl6or1hgj1k

Done
```
