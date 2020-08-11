# nohi
> **Last Release:** v0.1.0

> **Docs URL:** [nohi.unittolabs.com](https://nohi.unittolabs.com)

_NOHI - it's an abbreviation from Not Only Human-friendly IDs_

The main idea of this library and command-line tool is to provide easy and performant
way to generate useful sensitive data such a Human-friendly ID, password, etc.

## Requirements
### Build from the source
[go >= 1.14](https://golang.org/doc/install)

make >= 4.3 (`sudo [apt|dnf|eopkg] install make`, check your distro docs)

### Pre-built docker image
[docker >= 19.03](https://docs.docker.com/engine/install/)

## Installation
### Build from the source

```bash
# clone source repository
$ git clone git@gitlab.com:unitto/nohi.git

# go to cloned dir
$ cd nohi

# use project shortcut
$ make build

# change user permissions
$ sudo chowd root:root nohi

# copy binary to the path
$ sudo mv nohi /usr/bin/

# test instalation
$ nohi -h

# to install auto-completion check the
$ nohi completion -h
```

### Pre-built docker image

```bash
# run docker image
$ docker run unitto/nohi -h

# stable version (last release)
$ docker run unitto/nohi:stable -h
```

## Example Usage
Lets generate 10 logins and passwords for our mates:

```bash
$ nohi hid -n 10 && nohi pwd -n 10
boring_swartz_6zbdj
charming_wright_w9acu
clever_sutherland_82o4j
stoic_pascal_isa3v
vibrant_black_agcyb
hopeful_jennings_3uhu8
elated_borg_cgi0h
ecstatic_swartz_2nv8f
goofy_ritchie_c2366
sharp_ardinghelli_4w3tm

Done
c66k77ojt8lu7makk4dc0e8pjrl7oukv
ihn36voltcjz89njk3omyxamr1cbrubx
sll86ixpwi9fvpldok42dj4kg7qkue6j
qd6193vcdft4hy9zufu3sqhs6a9jtsrg
pcakkelrxmt2ggnd0a2q4mfp2jf1wfxr
syig9481g4d3axb2fca2gnqiq22abbuj
xtnlv5iz3d9cd8465oqovdspv70dw2j4
x7liiwyc4hjnbfe0vmk9400l6kp2usmm
s1ie3a52fwhyr95871hh5l17u931uy72
l1sypc7z8fk4q50qg78z50sfh6b5mupb

Done
```
